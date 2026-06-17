package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// fsxCmd represents the fsx command
var _fsxCmd = &cobra.Command{
	Use:   "fsx",
	Short: "AWS fsx CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := fsx.NewFromConfig(cfg)
		if _fsxAssociateFileSystemAliases {
			fsx_AssociateFileSystemAliases(cfg, client)
			return
		}
		if _fsxCancelDataRepositoryTask {
			fsx_CancelDataRepositoryTask(cfg, client)
			return
		}
		if _fsxCopyBackup {
			fsx_CopyBackup(cfg, client)
			return
		}
		if _fsxCopySnapshotAndUpdateVolume {
			fsx_CopySnapshotAndUpdateVolume(cfg, client)
			return
		}
		if _fsxCreateAndAttachS3AccessPoint {
			fsx_CreateAndAttachS3AccessPoint(cfg, client)
			return
		}
		if _fsxCreateBackup {
			fsx_CreateBackup(cfg, client)
			return
		}
		if _fsxCreateDataRepositoryAssociation {
			fsx_CreateDataRepositoryAssociation(cfg, client)
			return
		}
		if _fsxCreateDataRepositoryTask {
			fsx_CreateDataRepositoryTask(cfg, client)
			return
		}
		if _fsxCreateFileCache {
			fsx_CreateFileCache(cfg, client)
			return
		}
		if _fsxCreateFileSystem {
			fsx_CreateFileSystem(cfg, client)
			return
		}
		if _fsxCreateFileSystemFromBackup {
			fsx_CreateFileSystemFromBackup(cfg, client)
			return
		}
		if _fsxCreateSnapshot {
			fsx_CreateSnapshot(cfg, client)
			return
		}
		if _fsxCreateStorageVirtualMachine {
			fsx_CreateStorageVirtualMachine(cfg, client)
			return
		}
		if _fsxCreateVolume {
			fsx_CreateVolume(cfg, client)
			return
		}
		if _fsxCreateVolumeFromBackup {
			fsx_CreateVolumeFromBackup(cfg, client)
			return
		}
		if _fsxDeleteBackup {
			fsx_DeleteBackup(cfg, client)
			return
		}
		if _fsxDeleteDataRepositoryAssociation {
			fsx_DeleteDataRepositoryAssociation(cfg, client)
			return
		}
		if _fsxDeleteFileCache {
			fsx_DeleteFileCache(cfg, client)
			return
		}
		if _fsxDeleteFileSystem {
			fsx_DeleteFileSystem(cfg, client)
			return
		}
		if _fsxDeleteSnapshot {
			fsx_DeleteSnapshot(cfg, client)
			return
		}
		if _fsxDeleteStorageVirtualMachine {
			fsx_DeleteStorageVirtualMachine(cfg, client)
			return
		}
		if _fsxDeleteVolume {
			fsx_DeleteVolume(cfg, client)
			return
		}
		if _fsxDescribeBackups {
			fsx_DescribeBackups(cfg, client)
			return
		}
		if _fsxDescribeDataRepositoryAssociations {
			fsx_DescribeDataRepositoryAssociations(cfg, client)
			return
		}
		if _fsxDescribeDataRepositoryTasks {
			fsx_DescribeDataRepositoryTasks(cfg, client)
			return
		}
		if _fsxDescribeFileCaches {
			fsx_DescribeFileCaches(cfg, client)
			return
		}
		if _fsxDescribeFileSystemAliases {
			fsx_DescribeFileSystemAliases(cfg, client)
			return
		}
		if _fsxDescribeFileSystems {
			fsx_DescribeFileSystems(cfg, client)
			return
		}
		if _fsxDescribeS3AccessPointAttachments {
			fsx_DescribeS3AccessPointAttachments(cfg, client)
			return
		}
		if _fsxDescribeSharedVpcConfiguration {
			fsx_DescribeSharedVpcConfiguration(cfg, client)
			return
		}
		if _fsxDescribeSnapshots {
			fsx_DescribeSnapshots(cfg, client)
			return
		}
		if _fsxDescribeStorageVirtualMachines {
			fsx_DescribeStorageVirtualMachines(cfg, client)
			return
		}
		if _fsxDescribeVolumes {
			fsx_DescribeVolumes(cfg, client)
			return
		}
		if _fsxDetachAndDeleteS3AccessPoint {
			fsx_DetachAndDeleteS3AccessPoint(cfg, client)
			return
		}
		if _fsxDisassociateFileSystemAliases {
			fsx_DisassociateFileSystemAliases(cfg, client)
			return
		}
		if _fsxListTagsForResource {
			fsx_ListTagsForResource(cfg, client)
			return
		}
		if _fsxReleaseFileSystemNfsV3Locks {
			fsx_ReleaseFileSystemNfsV3Locks(cfg, client)
			return
		}
		if _fsxRestoreVolumeFromSnapshot {
			fsx_RestoreVolumeFromSnapshot(cfg, client)
			return
		}
		if _fsxStartMisconfiguredStateRecovery {
			fsx_StartMisconfiguredStateRecovery(cfg, client)
			return
		}
		if _fsxTagResource {
			fsx_TagResource(cfg, client)
			return
		}
		if _fsxUntagResource {
			fsx_UntagResource(cfg, client)
			return
		}
		if _fsxUpdateDataRepositoryAssociation {
			fsx_UpdateDataRepositoryAssociation(cfg, client)
			return
		}
		if _fsxUpdateFileCache {
			fsx_UpdateFileCache(cfg, client)
			return
		}
		if _fsxUpdateFileSystem {
			fsx_UpdateFileSystem(cfg, client)
			return
		}
		if _fsxUpdateSharedVpcConfiguration {
			fsx_UpdateSharedVpcConfiguration(cfg, client)
			return
		}
		if _fsxUpdateSnapshot {
			fsx_UpdateSnapshot(cfg, client)
			return
		}
		if _fsxUpdateStorageVirtualMachine {
			fsx_UpdateStorageVirtualMachine(cfg, client)
			return
		}
		if _fsxUpdateVolume {
			fsx_UpdateVolume(cfg, client)
			return
		}

	},
}

var (
	_fsxAssociateFileSystemAliases         bool
	_fsxCancelDataRepositoryTask           bool
	_fsxCopyBackup                         bool
	_fsxCopySnapshotAndUpdateVolume        bool
	_fsxCreateAndAttachS3AccessPoint       bool
	_fsxCreateBackup                       bool
	_fsxCreateDataRepositoryAssociation    bool
	_fsxCreateDataRepositoryTask           bool
	_fsxCreateFileCache                    bool
	_fsxCreateFileSystem                   bool
	_fsxCreateFileSystemFromBackup         bool
	_fsxCreateSnapshot                     bool
	_fsxCreateStorageVirtualMachine        bool
	_fsxCreateVolume                       bool
	_fsxCreateVolumeFromBackup             bool
	_fsxDeleteBackup                       bool
	_fsxDeleteDataRepositoryAssociation    bool
	_fsxDeleteFileCache                    bool
	_fsxDeleteFileSystem                   bool
	_fsxDeleteSnapshot                     bool
	_fsxDeleteStorageVirtualMachine        bool
	_fsxDeleteVolume                       bool
	_fsxDescribeBackups                    bool
	_fsxDescribeDataRepositoryAssociations bool
	_fsxDescribeDataRepositoryTasks        bool
	_fsxDescribeFileCaches                 bool
	_fsxDescribeFileSystemAliases          bool
	_fsxDescribeFileSystems                bool
	_fsxDescribeS3AccessPointAttachments   bool
	_fsxDescribeSharedVpcConfiguration     bool
	_fsxDescribeSnapshots                  bool
	_fsxDescribeStorageVirtualMachines     bool
	_fsxDescribeVolumes                    bool
	_fsxDetachAndDeleteS3AccessPoint       bool
	_fsxDisassociateFileSystemAliases      bool
	_fsxListTagsForResource                bool
	_fsxReleaseFileSystemNfsV3Locks        bool
	_fsxRestoreVolumeFromSnapshot          bool
	_fsxStartMisconfiguredStateRecovery    bool
	_fsxTagResource                        bool
	_fsxUntagResource                      bool
	_fsxUpdateDataRepositoryAssociation    bool
	_fsxUpdateFileCache                    bool
	_fsxUpdateFileSystem                   bool
	_fsxUpdateSharedVpcConfiguration       bool
	_fsxUpdateSnapshot                     bool
	_fsxUpdateStorageVirtualMachine        bool
	_fsxUpdateVolume                       bool

	_fsxActiveDirectoryConfiguration                      string
	_fsxAliases                                           []string
	_fsxAssociationId                                     string
	_fsxAssociationIds                                    []string
	_fsxBackupId                                          string
	_fsxBackupIds                                         []string
	_fsxBatchImportMetaDataOnCreate                       string
	_fsxCapacityToRelease                                 string
	_fsxClientRequestToken                                string
	_fsxCopyStrategy                                      string
	_fsxCopyTags                                          string
	_fsxCopyTagsToDataRepositoryAssociations              string
	_fsxDataRepositoryAssociations                        string
	_fsxDataRepositoryPath                                string
	_fsxDeleteDataInFileSystem                            string
	_fsxEnableFsxRouteTableUpdatesFromParticipantAccounts string
	_fsxFileCacheId                                       string
	_fsxFileCacheIds                                      []string
	_fsxFileCacheType                                     string
	_fsxFileCacheTypeVersion                              string
	_fsxFileSystemId                                      string
	_fsxFileSystemIds                                     []string
	_fsxFileSystemPath                                    string
	_fsxFileSystemType                                    string
	_fsxFileSystemTypeVersion                             string
	_fsxFilters                                           string
	_fsxImportedFileChunkSize                             string
	_fsxIncludeShared                                     string
	_fsxKmsKeyId                                          string
	_fsxLustreConfiguration                               string
	_fsxMaxResults                                        string
	_fsxName                                              string
	_fsxNames                                             []string
	_fsxNetworkType                                       string
	_fsxNextToken                                         string
	_fsxOntapConfiguration                                string
	_fsxOpenZFSConfiguration                              string
	_fsxOptions                                           string
	_fsxPaths                                             []string
	_fsxReleaseConfiguration                              string
	_fsxReport                                            string
	_fsxResourceARN                                       string
	_fsxRootVolumeSecurityStyle                           string
	_fsxS3                                                string
	_fsxS3AccessPoint                                     string
	_fsxSecurityGroupIds                                  []string
	_fsxSnapshotId                                        string
	_fsxSnapshotIds                                       []string
	_fsxSourceBackupId                                    string
	_fsxSourceRegion                                      string
	_fsxSourceSnapshotARN                                 string
	_fsxStorageCapacity                                   string
	_fsxStorageType                                       string
	_fsxStorageVirtualMachineId                           string
	_fsxStorageVirtualMachineIds                          []string
	_fsxSubnetIds                                         []string
	_fsxSvmAdminPassword                                  string
	_fsxTagKeys                                           []string
	_fsxTags                                              string
	_fsxTaskId                                            string
	_fsxTaskIds                                           []string
	_fsxType                                              string
	_fsxVolumeId                                          string
	_fsxVolumeIds                                         []string
	_fsxVolumeType                                        string
	_fsxWindowsConfiguration                              string
)

// Use this action to associate one or more Domain Name Server (DNS) aliases with
// an existing Amazon FSx for Windows File Server file system. A file system can
// have a maximum of 50 DNS aliases associated with it at any one time. If you try
// to associate a DNS alias that is already associated with the file system, FSx
// takes no action on that alias in the request. For more information, see [Working with DNS Aliases]and [Walkthrough 5: Using DNS aliases to access your file system],
// including additional steps you must take to be able to access your file system
// using a DNS alias.
//
// The system response shows the DNS aliases that Amazon FSx is attempting to
// associate with the file system. Use the API operation to monitor the status of
// the aliases Amazon FSx is associating with the file system.
//
// [Walkthrough 5: Using DNS aliases to access your file system]: https://docs.aws.amazon.com/fsx/latest/WindowsGuide/walkthrough05-file-system-custom-CNAME.html
// [Working with DNS Aliases]: https://docs.aws.amazon.com/fsx/latest/WindowsGuide/managing-dns-aliases.html
func fsx_AssociateFileSystemAliases(cfg aws.Config, client *fsx.Client) {
	input := &fsx.AssociateFileSystemAliasesInput{
		// Aliases: []string, // Required
		// FileSystemId: *string, // Required
	}

	if len(_fsxAliases) > 0 {
		input.Aliases = append([]string(nil), _fsxAliases...)
	}
	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}

	if resp, err := client.AssociateFileSystemAliases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an existing Amazon FSx for Lustre data repository task if that task is
// in either the PENDING or EXECUTING state. When you cancel an export task,
// Amazon FSx does the following.
//
// - Any files that FSx has already exported are not reverted.
//
// - FSx continues to export any files that are in-flight when the cancel
// operation is received.
//
// - FSx does not export any files that have not yet been exported.
//
// For a release task, Amazon FSx will stop releasing files upon cancellation. Any
// files that have already been released will remain in the released state.
func fsx_CancelDataRepositoryTask(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CancelDataRepositoryTaskInput{
		// TaskId: *string, // Required
	}

	if len(_fsxTaskId) > 0 {
		input.TaskId = aws.String(_fsxTaskId)
	}

	if resp, err := client.CancelDataRepositoryTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies an existing backup within the same Amazon Web Services account to
// another Amazon Web Services Region (cross-Region copy) or within the same Amazon
// Web Services Region (in-Region copy). You can have up to five backup copy
// requests in progress to a single destination Region per account.
//
// You can use cross-Region backup copies for cross-Region disaster recovery. You
// can periodically take backups and copy them to another Region so that in the
// event of a disaster in the primary Region, you can restore from backup and
// recover availability quickly in the other Region. You can make cross-Region
// copies only within your Amazon Web Services partition. A partition is a grouping
// of Regions. Amazon Web Services currently has three partitions: aws (Standard
// Regions), aws-cn (China Regions), and aws-us-gov (Amazon Web Services GovCloud
// [US] Regions).
//
// You can also use backup copies to clone your file dataset to another Region or
// within the same Region.
//
// You can use the SourceRegion parameter to specify the Amazon Web Services
// Region from which the backup will be copied. For example, if you make the call
// from the us-west-1 Region and want to copy a backup from the us-east-2 Region,
// you specify us-east-2 in the SourceRegion parameter to make a cross-Region
// copy. If you don't specify a Region, the backup copy is created in the same
// Region where the request is sent from (in-Region copy).
//
// For more information about creating backup copies, see [Copying backups] in the Amazon FSx for
// Windows User Guide, [Copying backups]in the Amazon FSx for Lustre User Guide, and [Copying backups] in the Amazon
// FSx for OpenZFS User Guide.
//
// [Copying backups]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/using-backups.html#copy-backups
func fsx_CopyBackup(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CopyBackupInput{
		// SourceBackupId: *string, // Required
	}

	if len(_fsxSourceBackupId) > 0 {
		input.SourceBackupId = aws.String(_fsxSourceBackupId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _fsxCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_fsxKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_fsxKmsKeyId)
	}
	if len(_fsxSourceRegion) > 0 {
		input.SourceRegion = aws.String(_fsxSourceRegion)
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing volume by using a snapshot from another Amazon FSx for
// OpenZFS file system. For more information, see [on-demand data replication]in the Amazon FSx for OpenZFS
// User Guide.
//
// [on-demand data replication]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/on-demand-replication.html
func fsx_CopySnapshotAndUpdateVolume(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CopySnapshotAndUpdateVolumeInput{
		// SourceSnapshotARN: *string, // Required
		// VolumeId: *string, // Required
	}

	if len(_fsxSourceSnapshotARN) > 0 {
		input.SourceSnapshotARN = aws.String(_fsxSourceSnapshotARN)
	}
	if len(_fsxVolumeId) > 0 {
		input.VolumeId = aws.String(_fsxVolumeId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxCopyStrategy) > 0 {
		if err := assignInputField(input, "CopyStrategy", _fsxCopyStrategy); err != nil {
			log.Errorf("invalid --copy-strategy: %s", err.Error())
			return
		}
	}
	if len(_fsxOptions) > 0 {
		if err := assignInputField(input, "Options", _fsxOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopySnapshotAndUpdateVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an S3 access point and attaches it to an Amazon FSx volume. For FSx for
// OpenZFS file systems, the volume must be hosted on a high-availability file
// system, either Single-AZ or Multi-AZ. For more information, see [Accessing your data using Amazon S3 access points]. in the Amazon
// FSx for OpenZFS User Guide.
//
// The requester requires the following permissions to perform these actions:
//
// - fsx:CreateAndAttachS3AccessPoint
//
// - s3:CreateAccessPoint
//
// - s3:GetAccessPoint
//
// - s3:PutAccessPointPolicy
//
// - s3:DeleteAccessPoint
//
// The following actions are related to CreateAndAttachS3AccessPoint :
//
// # DescribeS3AccessPointAttachments
//
// # DetachAndDeleteS3AccessPoint
//
// [Accessing your data using Amazon S3 access points]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/s3accesspoints-for-FSx.html
func fsx_CreateAndAttachS3AccessPoint(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateAndAttachS3AccessPointInput{
		// Name: *string, // Required
		// Type: types.S3AccessPointAttachmentType, // Required
	}

	if len(_fsxName) > 0 {
		input.Name = aws.String(_fsxName)
	}
	if len(_fsxType) > 0 {
		if err := assignInputField(input, "Type", _fsxType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxOntapConfiguration) > 0 {
		if err := assignInputField(input, "OntapConfiguration", _fsxOntapConfiguration); err != nil {
			log.Errorf("invalid --ontap-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxOpenZFSConfiguration) > 0 {
		if err := assignInputField(input, "OpenZFSConfiguration", _fsxOpenZFSConfiguration); err != nil {
			log.Errorf("invalid --open-zfs-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxS3AccessPoint) > 0 {
		if err := assignInputField(input, "S3AccessPoint", _fsxS3AccessPoint); err != nil {
			log.Errorf("invalid --s3-access-point: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAndAttachS3AccessPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a backup of an existing Amazon FSx for Windows File Server file system,
// Amazon FSx for Lustre file system, Amazon FSx for NetApp ONTAP volume, or Amazon
// FSx for OpenZFS file system. We recommend creating regular backups so that you
// can restore a file system or volume from a backup if an issue arises with the
// original file system or volume.
//
// For Amazon FSx for Lustre file systems, you can create a backup only for file
// systems that have the following configuration:
//
// - A Persistent deployment type
//
// - Are not linked to a data repository
//
// For more information about backups, see the following:
//
// - For Amazon FSx for Lustre, see [Working with FSx for Lustre backups].
//
// - For Amazon FSx for Windows, see [Working with FSx for Windows backups].
//
// - For Amazon FSx for NetApp ONTAP, see [Working with FSx for NetApp ONTAP backups].
//
// - For Amazon FSx for OpenZFS, see [Working with FSx for OpenZFS backups].
//
// If a backup with the specified client request token exists and the parameters
// match, this operation returns the description of the existing backup. If a
// backup with the specified client request token exists and the parameters don't
// match, this operation returns IncompatibleParameterError . If a backup with the
// specified client request token doesn't exist, CreateBackup does the following:
//
// - Creates a new Amazon FSx backup with an assigned ID, and an initial
// lifecycle state of CREATING .
//
// - Returns the description of the backup.
//
// By using the idempotent operation, you can retry a CreateBackup operation
// without the risk of creating an extra backup. This approach can be useful when
// an initial call fails in a way that makes it unclear whether a backup was
// created. If you use the same client request token and the initial call created a
// backup, the operation returns a successful result because all the parameters are
// the same.
//
// The CreateBackup operation returns while the backup's lifecycle state is still
// CREATING . You can check the backup creation status by calling the [DescribeBackups] operation,
// which returns the backup state along with other information.
//
// [Working with FSx for OpenZFS backups]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/using-backups.html
// [Working with FSx for NetApp ONTAP backups]: https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/using-backups.html
// [Working with FSx for Lustre backups]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-backups-fsx.html
// [Working with FSx for Windows backups]: https://docs.aws.amazon.com/fsx/latest/WindowsGuide/using-backups.html
// [DescribeBackups]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeBackups.html
func fsx_CreateBackup(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateBackupInput{}

	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_fsxVolumeId) > 0 {
		input.VolumeId = aws.String(_fsxVolumeId)
	}

	if resp, err := client.CreateBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon FSx for Lustre data repository association (DRA). A data
// repository association is a link between a directory on the file system and an
// Amazon S3 bucket or prefix. You can have a maximum of 8 data repository
// associations on a file system. Data repository associations are supported on all
// FSx for Lustre 2.12 and 2.15 file systems, excluding scratch_1 deployment type.
//
// Each data repository association must have a unique Amazon FSx file system
// directory and a unique S3 bucket or prefix associated with it. You can configure
// a data repository association for automatic import only, for automatic export
// only, or for both. To learn more about linking a data repository to your file
// system, see [Linking your file system to an S3 bucket].
//
// CreateDataRepositoryAssociation isn't supported on Amazon File Cache resources.
// To create a DRA on Amazon File Cache, use the CreateFileCache operation.
//
// [Linking your file system to an S3 bucket]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html
func fsx_CreateDataRepositoryAssociation(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateDataRepositoryAssociationInput{
		// DataRepositoryPath: *string, // Required
		// FileSystemId: *string, // Required
	}

	if len(_fsxDataRepositoryPath) > 0 {
		input.DataRepositoryPath = aws.String(_fsxDataRepositoryPath)
	}
	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxBatchImportMetaDataOnCreate) > 0 {
		if err := assignInputField(input, "BatchImportMetaDataOnCreate", _fsxBatchImportMetaDataOnCreate); err != nil {
			log.Errorf("invalid --batch-import-meta-data-on-create: %s", err.Error())
			return
		}
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxFileSystemPath) > 0 {
		input.FileSystemPath = aws.String(_fsxFileSystemPath)
	}
	if len(_fsxImportedFileChunkSize) > 0 {
		if err := assignInputField(input, "ImportedFileChunkSize", _fsxImportedFileChunkSize); err != nil {
			log.Errorf("invalid --imported-file-chunk-size: %s", err.Error())
			return
		}
	}
	if len(_fsxS3) > 0 {
		if err := assignInputField(input, "S3", _fsxS3); err != nil {
			log.Errorf("invalid --s3: %s", err.Error())
			return
		}
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataRepositoryAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon FSx for Lustre data repository task. A
// CreateDataRepositoryTask operation will fail if a data repository is not linked
// to the FSx file system.
//
// You use import and export data repository tasks to perform bulk operations
// between your FSx for Lustre file system and its linked data repositories. An
// example of a data repository task is exporting any data and metadata changes,
// including POSIX metadata, to files, directories, and symbolic links (symlinks)
// from your FSx file system to a linked data repository.
//
// You use release data repository tasks to release data from your file system for
// files that are exported to S3. The metadata of released files remains on the
// file system so users or applications can still access released files by reading
// the files again, which will restore data from Amazon S3 to the FSx for Lustre
// file system.
//
// To learn more about data repository tasks, see [Data Repository Tasks]. To learn more about linking a
// data repository to your file system, see [Linking your file system to an S3 bucket].
//
// [Data Repository Tasks]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/data-repository-tasks.html
// [Linking your file system to an S3 bucket]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html
func fsx_CreateDataRepositoryTask(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateDataRepositoryTaskInput{
		// FileSystemId: *string, // Required
		// Report: *types.CompletionReport, // Required
		// Type: types.DataRepositoryTaskType, // Required
	}

	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxReport) > 0 {
		if err := assignInputField(input, "Report", _fsxReport); err != nil {
			log.Errorf("invalid --report: %s", err.Error())
			return
		}
	}
	if len(_fsxType) > 0 {
		if err := assignInputField(input, "Type", _fsxType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_fsxCapacityToRelease) > 0 {
		if err := assignInputField(input, "CapacityToRelease", _fsxCapacityToRelease); err != nil {
			log.Errorf("invalid --capacity-to-release: %s", err.Error())
			return
		}
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxPaths) > 0 {
		input.Paths = append([]string(nil), _fsxPaths...)
	}
	if len(_fsxReleaseConfiguration) > 0 {
		if err := assignInputField(input, "ReleaseConfiguration", _fsxReleaseConfiguration); err != nil {
			log.Errorf("invalid --release-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataRepositoryTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon File Cache resource.
// You can use this operation with a client request token in the request that
// Amazon File Cache uses to ensure idempotent creation. If a cache with the
// specified client request token exists and the parameters match, CreateFileCache
// returns the description of the existing cache. If a cache with the specified
// client request token exists and the parameters don't match, this call returns
// IncompatibleParameterError . If a file cache with the specified client request
// token doesn't exist, CreateFileCache does the following:
//
// - Creates a new, empty Amazon File Cache resource with an assigned ID, and an
// initial lifecycle state of CREATING .
//
// - Returns the description of the cache in JSON format.
//
// The CreateFileCache call returns while the cache's lifecycle state is still
// CREATING . You can check the cache creation status by calling the [DescribeFileCaches] operation,
// which returns the cache state along with other information.
//
// [DescribeFileCaches]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileCaches.html
func fsx_CreateFileCache(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateFileCacheInput{
		// FileCacheType: types.FileCacheType, // Required
		// FileCacheTypeVersion: *string, // Required
		// StorageCapacity: *int32, // Required
		// SubnetIds: []string, // Required
	}

	if len(_fsxFileCacheType) > 0 {
		if err := assignInputField(input, "FileCacheType", _fsxFileCacheType); err != nil {
			log.Errorf("invalid --file-cache-type: %s", err.Error())
			return
		}
	}
	if len(_fsxFileCacheTypeVersion) > 0 {
		input.FileCacheTypeVersion = aws.String(_fsxFileCacheTypeVersion)
	}
	if len(_fsxStorageCapacity) > 0 {
		if err := assignInputField(input, "StorageCapacity", _fsxStorageCapacity); err != nil {
			log.Errorf("invalid --storage-capacity: %s", err.Error())
			return
		}
	}
	if len(_fsxSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _fsxSubnetIds...)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxCopyTagsToDataRepositoryAssociations) > 0 {
		if err := assignInputField(input, "CopyTagsToDataRepositoryAssociations", _fsxCopyTagsToDataRepositoryAssociations); err != nil {
			log.Errorf("invalid --copy-tags-to-data-repository-associations: %s", err.Error())
			return
		}
	}
	if len(_fsxDataRepositoryAssociations) > 0 {
		if err := assignInputField(input, "DataRepositoryAssociations", _fsxDataRepositoryAssociations); err != nil {
			log.Errorf("invalid --data-repository-associations: %s", err.Error())
			return
		}
	}
	if len(_fsxKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_fsxKmsKeyId)
	}
	if len(_fsxLustreConfiguration) > 0 {
		if err := assignInputField(input, "LustreConfiguration", _fsxLustreConfiguration); err != nil {
			log.Errorf("invalid --lustre-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _fsxSecurityGroupIds...)
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFileCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new, empty Amazon FSx file system. You can create the following
// supported Amazon FSx file systems using the CreateFileSystem API operation:
//
// - Amazon FSx for Lustre
//
// - Amazon FSx for NetApp ONTAP
//
// - Amazon FSx for OpenZFS
//
// - Amazon FSx for Windows File Server
//
// This operation requires a client request token in the request that Amazon FSx
// uses to ensure idempotent creation. This means that calling the operation
// multiple times with the same client request token has no effect. By using the
// idempotent operation, you can retry a CreateFileSystem operation without the
// risk of creating an extra file system. This approach can be useful when an
// initial call fails in a way that makes it unclear whether a file system was
// created. Examples are if a transport level timeout occurred, or your connection
// was reset. If you use the same client request token and the initial call created
// a file system, the client receives success as long as the parameters are the
// same.
//
// If a file system with the specified client request token exists and the
// parameters match, CreateFileSystem returns the description of the existing file
// system. If a file system with the specified client request token exists and the
// parameters don't match, this call returns IncompatibleParameterError . If a file
// system with the specified client request token doesn't exist, CreateFileSystem
// does the following:
//
// - Creates a new, empty Amazon FSx file system with an assigned ID, and an
// initial lifecycle state of CREATING .
//
// - Returns the description of the file system in JSON format.
//
// The CreateFileSystem call returns while the file system's lifecycle state is
// still CREATING . You can check the file-system creation status by calling the [DescribeFileSystems]
// operation, which returns the file system state along with other information.
//
// [DescribeFileSystems]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileSystems.html
func fsx_CreateFileSystem(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateFileSystemInput{
		// FileSystemType: types.FileSystemType, // Required
		// SubnetIds: []string, // Required
	}

	if len(_fsxFileSystemType) > 0 {
		if err := assignInputField(input, "FileSystemType", _fsxFileSystemType); err != nil {
			log.Errorf("invalid --file-system-type: %s", err.Error())
			return
		}
	}
	if len(_fsxSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _fsxSubnetIds...)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxFileSystemTypeVersion) > 0 {
		input.FileSystemTypeVersion = aws.String(_fsxFileSystemTypeVersion)
	}
	if len(_fsxKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_fsxKmsKeyId)
	}
	if len(_fsxLustreConfiguration) > 0 {
		if err := assignInputField(input, "LustreConfiguration", _fsxLustreConfiguration); err != nil {
			log.Errorf("invalid --lustre-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _fsxNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_fsxOntapConfiguration) > 0 {
		if err := assignInputField(input, "OntapConfiguration", _fsxOntapConfiguration); err != nil {
			log.Errorf("invalid --ontap-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxOpenZFSConfiguration) > 0 {
		if err := assignInputField(input, "OpenZFSConfiguration", _fsxOpenZFSConfiguration); err != nil {
			log.Errorf("invalid --open-zfs-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _fsxSecurityGroupIds...)
	}
	if len(_fsxStorageCapacity) > 0 {
		if err := assignInputField(input, "StorageCapacity", _fsxStorageCapacity); err != nil {
			log.Errorf("invalid --storage-capacity: %s", err.Error())
			return
		}
	}
	if len(_fsxStorageType) > 0 {
		if err := assignInputField(input, "StorageType", _fsxStorageType); err != nil {
			log.Errorf("invalid --storage-type: %s", err.Error())
			return
		}
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_fsxWindowsConfiguration) > 0 {
		if err := assignInputField(input, "WindowsConfiguration", _fsxWindowsConfiguration); err != nil {
			log.Errorf("invalid --windows-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFileSystem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon FSx for Lustre, Amazon FSx for Windows File Server, or
// Amazon FSx for OpenZFS file system from an existing Amazon FSx backup.
//
// If a file system with the specified client request token exists and the
// parameters match, this operation returns the description of the file system. If
// a file system with the specified client request token exists but the parameters
// don't match, this call returns IncompatibleParameterError . If a file system
// with the specified client request token doesn't exist, this operation does the
// following:
//
// - Creates a new Amazon FSx file system from backup with an assigned ID, and
// an initial lifecycle state of CREATING .
//
// - Returns the description of the file system.
//
// Parameters like the Active Directory, default share name, automatic backup, and
// backup settings default to the parameters of the file system that was backed up,
// unless overridden. You can explicitly supply other settings.
//
// By using the idempotent operation, you can retry a CreateFileSystemFromBackup
// call without the risk of creating an extra file system. This approach can be
// useful when an initial call fails in a way that makes it unclear whether a file
// system was created. Examples are if a transport level timeout occurred, or your
// connection was reset. If you use the same client request token and the initial
// call created a file system, the client receives a success message as long as the
// parameters are the same.
//
// The CreateFileSystemFromBackup call returns while the file system's lifecycle
// state is still CREATING . You can check the file-system creation status by
// calling the [DescribeFileSystems]operation, which returns the file system state along with other
// information.
//
// [DescribeFileSystems]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileSystems.html
func fsx_CreateFileSystemFromBackup(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateFileSystemFromBackupInput{
		// BackupId: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_fsxBackupId) > 0 {
		input.BackupId = aws.String(_fsxBackupId)
	}
	if len(_fsxSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _fsxSubnetIds...)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxFileSystemTypeVersion) > 0 {
		input.FileSystemTypeVersion = aws.String(_fsxFileSystemTypeVersion)
	}
	if len(_fsxKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_fsxKmsKeyId)
	}
	if len(_fsxLustreConfiguration) > 0 {
		if err := assignInputField(input, "LustreConfiguration", _fsxLustreConfiguration); err != nil {
			log.Errorf("invalid --lustre-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _fsxNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_fsxOpenZFSConfiguration) > 0 {
		if err := assignInputField(input, "OpenZFSConfiguration", _fsxOpenZFSConfiguration); err != nil {
			log.Errorf("invalid --open-zfs-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _fsxSecurityGroupIds...)
	}
	if len(_fsxStorageCapacity) > 0 {
		if err := assignInputField(input, "StorageCapacity", _fsxStorageCapacity); err != nil {
			log.Errorf("invalid --storage-capacity: %s", err.Error())
			return
		}
	}
	if len(_fsxStorageType) > 0 {
		if err := assignInputField(input, "StorageType", _fsxStorageType); err != nil {
			log.Errorf("invalid --storage-type: %s", err.Error())
			return
		}
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_fsxWindowsConfiguration) > 0 {
		if err := assignInputField(input, "WindowsConfiguration", _fsxWindowsConfiguration); err != nil {
			log.Errorf("invalid --windows-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFileSystemFromBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of an existing Amazon FSx for OpenZFS volume. With
// snapshots, you can easily undo file changes and compare file versions by
// restoring the volume to a previous version.
//
// If a snapshot with the specified client request token exists, and the
// parameters match, this operation returns the description of the existing
// snapshot. If a snapshot with the specified client request token exists, and the
// parameters don't match, this operation returns IncompatibleParameterError . If a
// snapshot with the specified client request token doesn't exist, CreateSnapshot
// does the following:
//
// - Creates a new OpenZFS snapshot with an assigned ID, and an initial
// lifecycle state of CREATING .
//
// - Returns the description of the snapshot.
//
// By using the idempotent operation, you can retry a CreateSnapshot operation
// without the risk of creating an extra snapshot. This approach can be useful when
// an initial call fails in a way that makes it unclear whether a snapshot was
// created. If you use the same client request token and the initial call created a
// snapshot, the operation returns a successful result because all the parameters
// are the same.
//
// The CreateSnapshot operation returns while the snapshot's lifecycle state is
// still CREATING . You can check the snapshot creation status by calling the [DescribeSnapshots]
// operation, which returns the snapshot state along with other information.
//
// [DescribeSnapshots]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeSnapshots.html
func fsx_CreateSnapshot(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateSnapshotInput{
		// Name: *string, // Required
		// VolumeId: *string, // Required
	}

	if len(_fsxName) > 0 {
		input.Name = aws.String(_fsxName)
	}
	if len(_fsxVolumeId) > 0 {
		input.VolumeId = aws.String(_fsxVolumeId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a storage virtual machine (SVM) for an Amazon FSx for ONTAP file system.
func fsx_CreateStorageVirtualMachine(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateStorageVirtualMachineInput{
		// FileSystemId: *string, // Required
		// Name: *string, // Required
	}

	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxName) > 0 {
		input.Name = aws.String(_fsxName)
	}
	if len(_fsxActiveDirectoryConfiguration) > 0 {
		if err := assignInputField(input, "ActiveDirectoryConfiguration", _fsxActiveDirectoryConfiguration); err != nil {
			log.Errorf("invalid --active-directory-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxRootVolumeSecurityStyle) > 0 {
		if err := assignInputField(input, "RootVolumeSecurityStyle", _fsxRootVolumeSecurityStyle); err != nil {
			log.Errorf("invalid --root-volume-security-style: %s", err.Error())
			return
		}
	}
	if len(_fsxSvmAdminPassword) > 0 {
		input.SvmAdminPassword = aws.String(_fsxSvmAdminPassword)
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStorageVirtualMachine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an FSx for ONTAP or Amazon FSx for OpenZFS storage volume.
func fsx_CreateVolume(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateVolumeInput{
		// Name: *string, // Required
		// VolumeType: types.VolumeType, // Required
	}

	if len(_fsxName) > 0 {
		input.Name = aws.String(_fsxName)
	}
	if len(_fsxVolumeType) > 0 {
		if err := assignInputField(input, "VolumeType", _fsxVolumeType); err != nil {
			log.Errorf("invalid --volume-type: %s", err.Error())
			return
		}
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxOntapConfiguration) > 0 {
		if err := assignInputField(input, "OntapConfiguration", _fsxOntapConfiguration); err != nil {
			log.Errorf("invalid --ontap-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxOpenZFSConfiguration) > 0 {
		if err := assignInputField(input, "OpenZFSConfiguration", _fsxOpenZFSConfiguration); err != nil {
			log.Errorf("invalid --open-zfs-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Creates a new Amazon FSx for NetApp ONTAP volume from an existing Amazon FSx
// volume backup.
func fsx_CreateVolumeFromBackup(cfg aws.Config, client *fsx.Client) {
	input := &fsx.CreateVolumeFromBackupInput{
		// BackupId: *string, // Required
		// Name: *string, // Required
	}

	if len(_fsxBackupId) > 0 {
		input.BackupId = aws.String(_fsxBackupId)
	}
	if len(_fsxName) > 0 {
		input.Name = aws.String(_fsxName)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxOntapConfiguration) > 0 {
		if err := assignInputField(input, "OntapConfiguration", _fsxOntapConfiguration); err != nil {
			log.Errorf("invalid --ontap-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVolumeFromBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon FSx backup. After deletion, the backup no longer exists, and
// its data is gone.
//
// The DeleteBackup call returns instantly. The backup won't show up in later
// DescribeBackups calls.
//
// The data in a deleted backup is also deleted and can't be recovered by any
// means.
func fsx_DeleteBackup(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DeleteBackupInput{
		// BackupId: *string, // Required
	}

	if len(_fsxBackupId) > 0 {
		input.BackupId = aws.String(_fsxBackupId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}

	if resp, err := client.DeleteBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a data repository association on an Amazon FSx for Lustre file system.
// Deleting the data repository association unlinks the file system from the Amazon
// S3 bucket. When deleting a data repository association, you have the option of
// deleting the data in the file system that corresponds to the data repository
// association. Data repository associations are supported on all FSx for Lustre
// 2.12 and 2.15 file systems, excluding scratch_1 deployment type.
func fsx_DeleteDataRepositoryAssociation(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DeleteDataRepositoryAssociationInput{
		// AssociationId: *string, // Required
	}

	if len(_fsxAssociationId) > 0 {
		input.AssociationId = aws.String(_fsxAssociationId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxDeleteDataInFileSystem) > 0 {
		if err := assignInputField(input, "DeleteDataInFileSystem", _fsxDeleteDataInFileSystem); err != nil {
			log.Errorf("invalid --delete-data-in-file-system: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDataRepositoryAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon File Cache resource. After deletion, the cache no longer
// exists, and its data is gone.
//
// The DeleteFileCache operation returns while the cache has the DELETING status.
// You can check the cache deletion status by calling the [DescribeFileCaches]operation, which returns
// a list of caches in your account. If you pass the cache ID for a deleted cache,
// the DescribeFileCaches operation returns a FileCacheNotFound error.
//
// The data in a deleted cache is also deleted and can't be recovered by any means.
//
// [DescribeFileCaches]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileCaches.html
func fsx_DeleteFileCache(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DeleteFileCacheInput{
		// FileCacheId: *string, // Required
	}

	if len(_fsxFileCacheId) > 0 {
		input.FileCacheId = aws.String(_fsxFileCacheId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}

	if resp, err := client.DeleteFileCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a file system. After deletion, the file system no longer exists, and
// its data is gone. Any existing automatic backups and snapshots are also deleted.
//
// To delete an Amazon FSx for NetApp ONTAP file system, first delete all the
// volumes and storage virtual machines (SVMs) on the file system. Then provide a
// FileSystemId value to the DeleteFileSystem operation.
//
// Before deleting an Amazon FSx for OpenZFS file system, make sure that there
// aren't any Amazon S3 access points attached to any volume. For more information
// on how to list S3 access points that are attached to volumes, see [Listing S3 access point attachments]. For more
// information on how to delete S3 access points, see [Deleting an S3 access point attachment].
//
// By default, when you delete an Amazon FSx for Windows File Server file system,
// a final backup is created upon deletion. This final backup isn't subject to the
// file system's retention policy, and must be manually deleted.
//
// To delete an Amazon FSx for Lustre file system, first [unmount] it from every connected
// Amazon EC2 instance, then provide a FileSystemId value to the DeleteFileSystem
// operation. By default, Amazon FSx will not take a final backup when the
// DeleteFileSystem operation is invoked. On file systems not linked to an Amazon
// S3 bucket, set SkipFinalBackup to false to take a final backup of the file
// system you are deleting. Backups cannot be enabled on S3-linked file systems. To
// ensure all of your data is written back to S3 before deleting your file system,
// you can either monitor for the [AgeOfOldestQueuedMessage]metric to be zero (if using automatic export) or
// you can run an [export data repository task]. If you have automatic export enabled and want to use an export
// data repository task, you have to disable automatic export before executing the
// export data repository task.
//
// The DeleteFileSystem operation returns while the file system has the DELETING
// status. You can check the file system deletion status by calling the [DescribeFileSystems]operation,
// which returns a list of file systems in your account. If you pass the file
// system ID for a deleted file system, the DescribeFileSystems operation returns
// a FileSystemNotFound error.
//
// If a data repository task is in a PENDING or EXECUTING state, deleting an
// Amazon FSx for Lustre file system will fail with an HTTP status code 400 (Bad
// Request).
//
// The data in a deleted file system is also deleted and can't be recovered by any
// means.
//
// [Deleting an S3 access point attachment]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/delete-access-point.html
// [unmount]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/unmounting-fs.html
// [AgeOfOldestQueuedMessage]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/monitoring-cloudwatch.html#auto-import-export-metrics
// [DescribeFileSystems]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileSystems.html
// [export data repository task]: https://docs.aws.amazon.com/fsx/latest/LustreGuide/export-data-repo-task-dra.html
// [Listing S3 access point attachments]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/access-points-list.html
func fsx_DeleteFileSystem(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DeleteFileSystemInput{
		// FileSystemId: *string, // Required
	}

	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxLustreConfiguration) > 0 {
		if err := assignInputField(input, "LustreConfiguration", _fsxLustreConfiguration); err != nil {
			log.Errorf("invalid --lustre-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxOpenZFSConfiguration) > 0 {
		if err := assignInputField(input, "OpenZFSConfiguration", _fsxOpenZFSConfiguration); err != nil {
			log.Errorf("invalid --open-zfs-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxWindowsConfiguration) > 0 {
		if err := assignInputField(input, "WindowsConfiguration", _fsxWindowsConfiguration); err != nil {
			log.Errorf("invalid --windows-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteFileSystem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon FSx for OpenZFS snapshot. After deletion, the snapshot no
// longer exists, and its data is gone. Deleting a snapshot doesn't affect
// snapshots stored in a file system backup.
//
// The DeleteSnapshot operation returns instantly. The snapshot appears with the
// lifecycle status of DELETING until the deletion is complete.
func fsx_DeleteSnapshot(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DeleteSnapshotInput{
		// SnapshotId: *string, // Required
	}

	if len(_fsxSnapshotId) > 0 {
		input.SnapshotId = aws.String(_fsxSnapshotId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}

	if resp, err := client.DeleteSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing Amazon FSx for ONTAP storage virtual machine (SVM). Prior
// to deleting an SVM, you must delete all non-root volumes in the SVM, otherwise
// the operation will fail.
func fsx_DeleteStorageVirtualMachine(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DeleteStorageVirtualMachineInput{
		// StorageVirtualMachineId: *string, // Required
	}

	if len(_fsxStorageVirtualMachineId) > 0 {
		input.StorageVirtualMachineId = aws.String(_fsxStorageVirtualMachineId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}

	if resp, err := client.DeleteStorageVirtualMachine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS volume.
func fsx_DeleteVolume(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DeleteVolumeInput{
		// VolumeId: *string, // Required
	}

	if len(_fsxVolumeId) > 0 {
		input.VolumeId = aws.String(_fsxVolumeId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxOntapConfiguration) > 0 {
		if err := assignInputField(input, "OntapConfiguration", _fsxOntapConfiguration); err != nil {
			log.Errorf("invalid --ontap-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxOpenZFSConfiguration) > 0 {
		if err := assignInputField(input, "OpenZFSConfiguration", _fsxOpenZFSConfiguration); err != nil {
			log.Errorf("invalid --open-zfs-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description of a specific Amazon FSx backup, if a BackupIds value
// is provided for that backup. Otherwise, it returns all backups owned by your
// Amazon Web Services account in the Amazon Web Services Region of the endpoint
// that you're calling.
//
// When retrieving all backups, you can optionally specify the MaxResults
// parameter to limit the number of backups in a response. If more backups remain,
// Amazon FSx returns a NextToken value in the response. In this case, send a
// later request with the NextToken request parameter set to the value of the
// NextToken value from the last response.
//
// This operation is used in an iterative process to retrieve a list of your
// backups. DescribeBackups is called first without a NextToken value. Then the
// operation continues to be called with the NextToken parameter set to the value
// of the last NextToken value until a response has no NextToken value.
//
// When using this operation, keep the following in mind:
//
// - The operation might return fewer than the MaxResults value of backup
// descriptions while still including a NextToken value.
//
// - The order of the backups returned in the response of one DescribeBackups
// call and the order of the backups returned across the responses of a multi-call
// iteration is unspecified.
func fsx_DescribeBackups(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeBackupsInput{}

	if len(_fsxBackupIds) > 0 {
		input.BackupIds = append([]string(nil), _fsxBackupIds...)
	}
	if len(_fsxFilters) > 0 {
		if err := assignInputField(input, "Filters", _fsxFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeBackups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fsx.DescribeBackupsOutput
	p := fsx.NewDescribeBackupsPaginator(client, input)
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

// Returns the description of specific Amazon FSx for Lustre or Amazon File Cache
// data repository associations, if one or more AssociationIds values are provided
// in the request, or if filters are used in the request. Data repository
// associations are supported on Amazon File Cache resources and all FSx for Lustre
// 2.12 and 2,15 file systems, excluding scratch_1 deployment type.
//
// You can use filters to narrow the response to include just data repository
// associations for specific file systems (use the file-system-id filter with the
// ID of the file system) or caches (use the file-cache-id filter with the ID of
// the cache), or data repository associations for a specific repository type (use
// the data-repository-type filter with a value of S3 or NFS ). If you don't use
// filters, the response returns all data repository associations owned by your
// Amazon Web Services account in the Amazon Web Services Region of the endpoint
// that you're calling.
//
// When retrieving all data repository associations, you can paginate the response
// by using the optional MaxResults parameter to limit the number of data
// repository associations returned in a response. If more data repository
// associations remain, a NextToken value is returned in the response. In this
// case, send a later request with the NextToken request parameter set to the
// value of NextToken from the last response.
func fsx_DescribeDataRepositoryAssociations(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeDataRepositoryAssociationsInput{}

	if len(_fsxAssociationIds) > 0 {
		input.AssociationIds = append([]string(nil), _fsxAssociationIds...)
	}
	if len(_fsxFilters) > 0 {
		if err := assignInputField(input, "Filters", _fsxFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDataRepositoryAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fsx.DescribeDataRepositoryAssociationsOutput
	p := fsx.NewDescribeDataRepositoryAssociationsPaginator(client, input)
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

// Returns the description of specific Amazon FSx for Lustre or Amazon File Cache
// data repository tasks, if one or more TaskIds values are provided in the
// request, or if filters are used in the request. You can use filters to narrow
// the response to include just tasks for specific file systems or caches, or tasks
// in a specific lifecycle state. Otherwise, it returns all data repository tasks
// owned by your Amazon Web Services account in the Amazon Web Services Region of
// the endpoint that you're calling.
//
// When retrieving all tasks, you can paginate the response by using the optional
// MaxResults parameter to limit the number of tasks returned in a response. If
// more tasks remain, a NextToken value is returned in the response. In this case,
// send a later request with the NextToken request parameter set to the value of
// NextToken from the last response.
func fsx_DescribeDataRepositoryTasks(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeDataRepositoryTasksInput{}

	if len(_fsxFilters) > 0 {
		if err := assignInputField(input, "Filters", _fsxFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
	}
	if len(_fsxTaskIds) > 0 {
		input.TaskIds = append([]string(nil), _fsxTaskIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDataRepositoryTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fsx.DescribeDataRepositoryTasksOutput
	p := fsx.NewDescribeDataRepositoryTasksPaginator(client, input)
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

// Returns the description of a specific Amazon File Cache resource, if a
// FileCacheIds value is provided for that cache. Otherwise, it returns
// descriptions of all caches owned by your Amazon Web Services account in the
// Amazon Web Services Region of the endpoint that you're calling.
//
// When retrieving all cache descriptions, you can optionally specify the
// MaxResults parameter to limit the number of descriptions in a response. If more
// cache descriptions remain, the operation returns a NextToken value in the
// response. In this case, send a later request with the NextToken request
// parameter set to the value of NextToken from the last response.
//
// This operation is used in an iterative process to retrieve a list of your cache
// descriptions. DescribeFileCaches is called first without a NextToken value. Then
// the operation continues to be called with the NextToken parameter set to the
// value of the last NextToken value until a response has no NextToken .
//
// When using this operation, keep the following in mind:
//
// - The implementation might return fewer than MaxResults cache descriptions
// while still including a NextToken value.
//
// - The order of caches returned in the response of one DescribeFileCaches call
// and the order of caches returned across the responses of a multicall iteration
// is unspecified.
func fsx_DescribeFileCaches(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeFileCachesInput{}

	if len(_fsxFileCacheIds) > 0 {
		input.FileCacheIds = append([]string(nil), _fsxFileCacheIds...)
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFileCaches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fsx.DescribeFileCachesOutput
	p := fsx.NewDescribeFileCachesPaginator(client, input)
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

// Returns the DNS aliases that are associated with the specified Amazon FSx for
// Windows File Server file system. A history of all DNS aliases that have been
// associated with and disassociated from the file system is available in the list
// of AdministrativeActionprovided in the DescribeFileSystems operation response.
func fsx_DescribeFileSystemAliases(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeFileSystemAliasesInput{
		// FileSystemId: *string, // Required
	}

	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFileSystemAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fsx.DescribeFileSystemAliasesOutput
	p := fsx.NewDescribeFileSystemAliasesPaginator(client, input)
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

// Returns the description of specific Amazon FSx file systems, if a FileSystemIds
// value is provided for that file system. Otherwise, it returns descriptions of
// all file systems owned by your Amazon Web Services account in the Amazon Web
// Services Region of the endpoint that you're calling.
//
// When retrieving all file system descriptions, you can optionally specify the
// MaxResults parameter to limit the number of descriptions in a response. If more
// file system descriptions remain, Amazon FSx returns a NextToken value in the
// response. In this case, send a later request with the NextToken request
// parameter set to the value of NextToken from the last response.
//
// This operation is used in an iterative process to retrieve a list of your file
// system descriptions. DescribeFileSystems is called first without a NextToken
// value. Then the operation continues to be called with the NextToken parameter
// set to the value of the last NextToken value until a response has no NextToken .
//
// When using this operation, keep the following in mind:
//
// - The implementation might return fewer than MaxResults file system
// descriptions while still including a NextToken value.
//
// - The order of file systems returned in the response of one
// DescribeFileSystems call and the order of file systems returned across the
// responses of a multicall iteration is unspecified.
func fsx_DescribeFileSystems(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeFileSystemsInput{}

	if len(_fsxFileSystemIds) > 0 {
		input.FileSystemIds = append([]string(nil), _fsxFileSystemIds...)
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeFileSystems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fsx.DescribeFileSystemsOutput
	p := fsx.NewDescribeFileSystemsPaginator(client, input)
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

// Describes one or more S3 access points attached to Amazon FSx volumes.
// The requester requires the following permission to perform this action:
//
// - fsx:DescribeS3AccessPointAttachments
func fsx_DescribeS3AccessPointAttachments(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeS3AccessPointAttachmentsInput{}

	if len(_fsxFilters) > 0 {
		if err := assignInputField(input, "Filters", _fsxFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNames) > 0 {
		input.Names = append([]string(nil), _fsxNames...)
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeS3AccessPointAttachments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fsx.DescribeS3AccessPointAttachmentsOutput
	p := fsx.NewDescribeS3AccessPointAttachmentsPaginator(client, input)
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

// Indicates whether participant accounts in your organization can create Amazon
// FSx for NetApp ONTAP Multi-AZ file systems in subnets that are shared by a
// virtual private cloud (VPC) owner. For more information, see [Creating FSx for ONTAP file systems in shared subnets].
//
// [Creating FSx for ONTAP file systems in shared subnets]: https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/creating-file-systems.html#fsxn-vpc-shared-subnets
func fsx_DescribeSharedVpcConfiguration(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeSharedVpcConfigurationInput{}

	if resp, err := client.DescribeSharedVpcConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description of specific Amazon FSx for OpenZFS snapshots, if a
// SnapshotIds value is provided. Otherwise, this operation returns all snapshots
// owned by your Amazon Web Services account in the Amazon Web Services Region of
// the endpoint that you're calling.
//
// When retrieving all snapshots, you can optionally specify the MaxResults
// parameter to limit the number of snapshots in a response. If more backups
// remain, Amazon FSx returns a NextToken value in the response. In this case,
// send a later request with the NextToken request parameter set to the value of
// NextToken from the last response.
//
// Use this operation in an iterative process to retrieve a list of your
// snapshots. DescribeSnapshots is called first without a NextToken value. Then
// the operation continues to be called with the NextToken parameter set to the
// value of the last NextToken value until a response has no NextToken value.
//
// When using this operation, keep the following in mind:
//
// - The operation might return fewer than the MaxResults value of snapshot
// descriptions while still including a NextToken value.
//
// - The order of snapshots returned in the response of one DescribeSnapshots
// call and the order of backups returned across the responses of a multi-call
// iteration is unspecified.
func fsx_DescribeSnapshots(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeSnapshotsInput{}

	if len(_fsxFilters) > 0 {
		if err := assignInputField(input, "Filters", _fsxFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_fsxIncludeShared) > 0 {
		if err := assignInputField(input, "IncludeShared", _fsxIncludeShared); err != nil {
			log.Errorf("invalid --include-shared: %s", err.Error())
			return
		}
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
	}
	if len(_fsxSnapshotIds) > 0 {
		input.SnapshotIds = append([]string(nil), _fsxSnapshotIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fsx.DescribeSnapshotsOutput
	p := fsx.NewDescribeSnapshotsPaginator(client, input)
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

// Describes one or more Amazon FSx for NetApp ONTAP storage virtual machines
// (SVMs).
func fsx_DescribeStorageVirtualMachines(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeStorageVirtualMachinesInput{}

	if len(_fsxFilters) > 0 {
		if err := assignInputField(input, "Filters", _fsxFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
	}
	if len(_fsxStorageVirtualMachineIds) > 0 {
		input.StorageVirtualMachineIds = append([]string(nil), _fsxStorageVirtualMachineIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeStorageVirtualMachines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fsx.DescribeStorageVirtualMachinesOutput
	p := fsx.NewDescribeStorageVirtualMachinesPaginator(client, input)
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

// Describes one or more Amazon FSx for NetApp ONTAP or Amazon FSx for OpenZFS
// volumes.
func fsx_DescribeVolumes(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DescribeVolumesInput{}

	if len(_fsxFilters) > 0 {
		if err := assignInputField(input, "Filters", _fsxFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
	}
	if len(_fsxVolumeIds) > 0 {
		input.VolumeIds = append([]string(nil), _fsxVolumeIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeVolumes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*fsx.DescribeVolumesOutput
	p := fsx.NewDescribeVolumesPaginator(client, input)
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

// Detaches an S3 access point from an Amazon FSx volume and deletes the S3 access
// point.
//
// The requester requires the following permission to perform this action:
//
// - fsx:DetachAndDeleteS3AccessPoint
//
// - s3:DeleteAccessPoint
func fsx_DetachAndDeleteS3AccessPoint(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DetachAndDeleteS3AccessPointInput{
		// Name: *string, // Required
	}

	if len(_fsxName) > 0 {
		input.Name = aws.String(_fsxName)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}

	if resp, err := client.DetachAndDeleteS3AccessPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this action to disassociate, or remove, one or more Domain Name Service
// (DNS) aliases from an Amazon FSx for Windows File Server file system. If you
// attempt to disassociate a DNS alias that is not associated with the file system,
// Amazon FSx responds with an HTTP status code 400 (Bad Request). For more
// information, see [Working with DNS Aliases].
//
// The system generated response showing the DNS aliases that Amazon FSx is
// attempting to disassociate from the file system. Use the API operation to
// monitor the status of the aliases Amazon FSx is disassociating with the file
// system.
//
// [Working with DNS Aliases]: https://docs.aws.amazon.com/fsx/latest/WindowsGuide/managing-dns-aliases.html
func fsx_DisassociateFileSystemAliases(cfg aws.Config, client *fsx.Client) {
	input := &fsx.DisassociateFileSystemAliasesInput{
		// Aliases: []string, // Required
		// FileSystemId: *string, // Required
	}

	if len(_fsxAliases) > 0 {
		input.Aliases = append([]string(nil), _fsxAliases...)
	}
	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}

	if resp, err := client.DisassociateFileSystemAliases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists tags for Amazon FSx resources.
// When retrieving all tags, you can optionally specify the MaxResults parameter
// to limit the number of tags in a response. If more tags remain, Amazon FSx
// returns a NextToken value in the response. In this case, send a later request
// with the NextToken request parameter set to the value of NextToken from the
// last response.
//
// This action is used in an iterative process to retrieve a list of your tags.
// ListTagsForResource is called first without a NextToken value. Then the action
// continues to be called with the NextToken parameter set to the value of the
// last NextToken value until a response has no NextToken .
//
// When using this action, keep the following in mind:
//
// - The implementation might return fewer than MaxResults file system
// descriptions while still including a NextToken value.
//
// - The order of tags returned in the response of one ListTagsForResource call
// and the order of tags returned across the responses of a multi-call iteration is
// unspecified.
func fsx_ListTagsForResource(cfg aws.Config, client *fsx.Client) {
	input := &fsx.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_fsxResourceARN) > 0 {
		input.ResourceARN = aws.String(_fsxResourceARN)
	}
	if len(_fsxMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _fsxMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_fsxNextToken) > 0 {
		input.NextToken = aws.String(_fsxNextToken)
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

	var results []*fsx.ListTagsForResourceOutput
	p := fsx.NewListTagsForResourcePaginator(client, input)
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

// Releases the file system lock from an Amazon FSx for OpenZFS file system.
func fsx_ReleaseFileSystemNfsV3Locks(cfg aws.Config, client *fsx.Client) {
	input := &fsx.ReleaseFileSystemNfsV3LocksInput{
		// FileSystemId: *string, // Required
	}

	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}

	if resp, err := client.ReleaseFileSystemNfsV3Locks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an Amazon FSx for OpenZFS volume to the state saved by the specified
// snapshot.
func fsx_RestoreVolumeFromSnapshot(cfg aws.Config, client *fsx.Client) {
	input := &fsx.RestoreVolumeFromSnapshotInput{
		// SnapshotId: *string, // Required
		// VolumeId: *string, // Required
	}

	if len(_fsxSnapshotId) > 0 {
		input.SnapshotId = aws.String(_fsxSnapshotId)
	}
	if len(_fsxVolumeId) > 0 {
		input.VolumeId = aws.String(_fsxVolumeId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxOptions) > 0 {
		if err := assignInputField(input, "Options", _fsxOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}

	if resp, err := client.RestoreVolumeFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// After performing steps to repair the Active Directory configuration of an FSx
// for Windows File Server file system, use this action to initiate the process of
// Amazon FSx attempting to reconnect to the file system.
func fsx_StartMisconfiguredStateRecovery(cfg aws.Config, client *fsx.Client) {
	input := &fsx.StartMisconfiguredStateRecoveryInput{
		// FileSystemId: *string, // Required
	}

	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}

	if resp, err := client.StartMisconfiguredStateRecovery(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags an Amazon FSx resource.
func fsx_TagResource(cfg aws.Config, client *fsx.Client) {
	input := &fsx.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_fsxResourceARN) > 0 {
		input.ResourceARN = aws.String(_fsxResourceARN)
	}
	if len(_fsxTags) > 0 {
		if err := assignInputField(input, "Tags", _fsxTags); err != nil {
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

// This action removes a tag from an Amazon FSx resource.
func fsx_UntagResource(cfg aws.Config, client *fsx.Client) {
	input := &fsx.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_fsxResourceARN) > 0 {
		input.ResourceARN = aws.String(_fsxResourceARN)
	}
	if len(_fsxTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _fsxTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing data repository association on an
// Amazon FSx for Lustre file system. Data repository associations are supported on
// all FSx for Lustre 2.12 and 2.15 file systems, excluding scratch_1 deployment
// type.
func fsx_UpdateDataRepositoryAssociation(cfg aws.Config, client *fsx.Client) {
	input := &fsx.UpdateDataRepositoryAssociationInput{
		// AssociationId: *string, // Required
	}

	if len(_fsxAssociationId) > 0 {
		input.AssociationId = aws.String(_fsxAssociationId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxImportedFileChunkSize) > 0 {
		if err := assignInputField(input, "ImportedFileChunkSize", _fsxImportedFileChunkSize); err != nil {
			log.Errorf("invalid --imported-file-chunk-size: %s", err.Error())
			return
		}
	}
	if len(_fsxS3) > 0 {
		if err := assignInputField(input, "S3", _fsxS3); err != nil {
			log.Errorf("invalid --s3: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataRepositoryAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing Amazon File Cache resource. You can
// update multiple properties in a single request.
func fsx_UpdateFileCache(cfg aws.Config, client *fsx.Client) {
	input := &fsx.UpdateFileCacheInput{
		// FileCacheId: *string, // Required
	}

	if len(_fsxFileCacheId) > 0 {
		input.FileCacheId = aws.String(_fsxFileCacheId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxLustreConfiguration) > 0 {
		if err := assignInputField(input, "LustreConfiguration", _fsxLustreConfiguration); err != nil {
			log.Errorf("invalid --lustre-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFileCache(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to update the configuration of an existing Amazon FSx file
// system. You can update multiple properties in a single request.
//
// For FSx for Windows File Server file systems, you can update the following
// properties:
//
// - AuditLogConfiguration
//
// - AutomaticBackupRetentionDays
//
// - DailyAutomaticBackupStartTime
//
// - DiskIopsConfiguration
//
// - SelfManagedActiveDirectoryConfiguration
//
// - StorageCapacity
//
// - StorageType
//
// - ThroughputCapacity
//
// - WeeklyMaintenanceStartTime
//
// For FSx for Lustre file systems, you can update the following properties:
//
// - AutoImportPolicy
//
// - AutomaticBackupRetentionDays
//
// - DailyAutomaticBackupStartTime
//
// - DataCompressionType
//
// - FileSystemTypeVersion
//
// - LogConfiguration
//
// - LustreReadCacheConfiguration
//
// - LustreRootSquashConfiguration
//
// - MetadataConfiguration
//
// - PerUnitStorageThroughput
//
// - StorageCapacity
//
// - ThroughputCapacity
//
// - WeeklyMaintenanceStartTime
//
// For FSx for ONTAP file systems, you can update the following properties:
//
// - AddRouteTableIds
//
// - AutomaticBackupRetentionDays
//
// - DailyAutomaticBackupStartTime
//
// - DiskIopsConfiguration
//
// - EndpointIpv6AddressRange
//
// - FsxAdminPassword
//
// - HAPairs
//
// - RemoveRouteTableIds
//
// - StorageCapacity
//
// - ThroughputCapacity
//
// - ThroughputCapacityPerHAPair
//
// - WeeklyMaintenanceStartTime
//
// For FSx for OpenZFS file systems, you can update the following properties:
//
// - AddRouteTableIds
//
// - AutomaticBackupRetentionDays
//
// - CopyTagsToBackups
//
// - CopyTagsToVolumes
//
// - DailyAutomaticBackupStartTime
//
// - DiskIopsConfiguration
//
// - EndpointIpv6AddressRange
//
// - ReadCacheConfiguration
//
// - RemoveRouteTableIds
//
// - StorageCapacity
//
// - ThroughputCapacity
//
// - WeeklyMaintenanceStartTime
func fsx_UpdateFileSystem(cfg aws.Config, client *fsx.Client) {
	input := &fsx.UpdateFileSystemInput{
		// FileSystemId: *string, // Required
	}

	if len(_fsxFileSystemId) > 0 {
		input.FileSystemId = aws.String(_fsxFileSystemId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxFileSystemTypeVersion) > 0 {
		input.FileSystemTypeVersion = aws.String(_fsxFileSystemTypeVersion)
	}
	if len(_fsxLustreConfiguration) > 0 {
		if err := assignInputField(input, "LustreConfiguration", _fsxLustreConfiguration); err != nil {
			log.Errorf("invalid --lustre-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _fsxNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_fsxOntapConfiguration) > 0 {
		if err := assignInputField(input, "OntapConfiguration", _fsxOntapConfiguration); err != nil {
			log.Errorf("invalid --ontap-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxOpenZFSConfiguration) > 0 {
		if err := assignInputField(input, "OpenZFSConfiguration", _fsxOpenZFSConfiguration); err != nil {
			log.Errorf("invalid --open-zfs-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxStorageCapacity) > 0 {
		if err := assignInputField(input, "StorageCapacity", _fsxStorageCapacity); err != nil {
			log.Errorf("invalid --storage-capacity: %s", err.Error())
			return
		}
	}
	if len(_fsxStorageType) > 0 {
		if err := assignInputField(input, "StorageType", _fsxStorageType); err != nil {
			log.Errorf("invalid --storage-type: %s", err.Error())
			return
		}
	}
	if len(_fsxWindowsConfiguration) > 0 {
		if err := assignInputField(input, "WindowsConfiguration", _fsxWindowsConfiguration); err != nil {
			log.Errorf("invalid --windows-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFileSystem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures whether participant accounts in your organization can create Amazon
// FSx for NetApp ONTAP Multi-AZ file systems in subnets that are shared by a
// virtual private cloud (VPC) owner. For more information, see the [Amazon FSx for NetApp ONTAP User Guide].
//
// We strongly recommend that participant-created Multi-AZ file systems in the
// shared VPC are deleted before you disable this feature. Once the feature is
// disabled, these file systems will enter a MISCONFIGURED state and behave like
// Single-AZ file systems. For more information, see [Important considerations before disabling shared VPC support for Multi-AZ file systems].
//
// [Amazon FSx for NetApp ONTAP User Guide]: https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/maz-shared-vpc.html
// [Important considerations before disabling shared VPC support for Multi-AZ file systems]: https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/maz-shared-vpc.html#disabling-maz-vpc-sharing
func fsx_UpdateSharedVpcConfiguration(cfg aws.Config, client *fsx.Client) {
	input := &fsx.UpdateSharedVpcConfigurationInput{}

	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxEnableFsxRouteTableUpdatesFromParticipantAccounts) > 0 {
		input.EnableFsxRouteTableUpdatesFromParticipantAccounts = aws.String(_fsxEnableFsxRouteTableUpdatesFromParticipantAccounts)
	}

	if resp, err := client.UpdateSharedVpcConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name of an Amazon FSx for OpenZFS snapshot.
func fsx_UpdateSnapshot(cfg aws.Config, client *fsx.Client) {
	input := &fsx.UpdateSnapshotInput{
		// Name: *string, // Required
		// SnapshotId: *string, // Required
	}

	if len(_fsxName) > 0 {
		input.Name = aws.String(_fsxName)
	}
	if len(_fsxSnapshotId) > 0 {
		input.SnapshotId = aws.String(_fsxSnapshotId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}

	if resp, err := client.UpdateSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an FSx for ONTAP storage virtual machine (SVM).
func fsx_UpdateStorageVirtualMachine(cfg aws.Config, client *fsx.Client) {
	input := &fsx.UpdateStorageVirtualMachineInput{
		// StorageVirtualMachineId: *string, // Required
	}

	if len(_fsxStorageVirtualMachineId) > 0 {
		input.StorageVirtualMachineId = aws.String(_fsxStorageVirtualMachineId)
	}
	if len(_fsxActiveDirectoryConfiguration) > 0 {
		if err := assignInputField(input, "ActiveDirectoryConfiguration", _fsxActiveDirectoryConfiguration); err != nil {
			log.Errorf("invalid --active-directory-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxSvmAdminPassword) > 0 {
		input.SvmAdminPassword = aws.String(_fsxSvmAdminPassword)
	}

	if resp, err := client.UpdateStorageVirtualMachine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an Amazon FSx for NetApp ONTAP or Amazon FSx for
// OpenZFS volume.
func fsx_UpdateVolume(cfg aws.Config, client *fsx.Client) {
	input := &fsx.UpdateVolumeInput{
		// VolumeId: *string, // Required
	}

	if len(_fsxVolumeId) > 0 {
		input.VolumeId = aws.String(_fsxVolumeId)
	}
	if len(_fsxClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_fsxClientRequestToken)
	}
	if len(_fsxName) > 0 {
		input.Name = aws.String(_fsxName)
	}
	if len(_fsxOntapConfiguration) > 0 {
		if err := assignInputField(input, "OntapConfiguration", _fsxOntapConfiguration); err != nil {
			log.Errorf("invalid --ontap-configuration: %s", err.Error())
			return
		}
	}
	if len(_fsxOpenZFSConfiguration) > 0 {
		if err := assignInputField(input, "OpenZFSConfiguration", _fsxOpenZFSConfiguration); err != nil {
			log.Errorf("invalid --open-zfs-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateVolume(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_fsxCmd)
	_fsxCmd.Flags().SortFlags = false

	_fsxCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_fsxCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_fsxCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_fsxCmd.Flags().StringVarP(&_fsxActiveDirectoryConfiguration, "active-directory-configuration", "", "", "Active Directory Configuration")
	_fsxCmd.Flags().StringSliceVarP(&_fsxAliases, "aliases", "", nil, "Aliases")
	_fsxCmd.Flags().StringVarP(&_fsxAssociationId, "association-id", "", "", "Association ID")
	_fsxCmd.Flags().StringSliceVarP(&_fsxAssociationIds, "association-ids", "", nil, "Association Ids")
	_fsxCmd.Flags().StringVarP(&_fsxBackupId, "backup-id", "", "", "Backup ID")
	_fsxCmd.Flags().StringSliceVarP(&_fsxBackupIds, "backup-ids", "", nil, "Backup Ids")
	_fsxCmd.Flags().StringVarP(&_fsxBatchImportMetaDataOnCreate, "batch-import-meta-data-on-create", "", "", "Batch Import Meta Data On Create")
	_fsxCmd.Flags().StringVarP(&_fsxCapacityToRelease, "capacity-to-release", "", "", "Capacity To Release")
	_fsxCmd.Flags().StringVarP(&_fsxClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_fsxCmd.Flags().StringVarP(&_fsxCopyStrategy, "copy-strategy", "", "", "Copy Strategy")
	_fsxCmd.Flags().StringVarP(&_fsxCopyTags, "copy-tags", "", "", "Copy Tags")
	_fsxCmd.Flags().StringVarP(&_fsxCopyTagsToDataRepositoryAssociations, "copy-tags-to-data-repository-associations", "", "", "Copy Tags To Data Repository Associations")
	_fsxCmd.Flags().StringVarP(&_fsxDataRepositoryAssociations, "data-repository-associations", "", "", "Data Repository Associations")
	_fsxCmd.Flags().StringVarP(&_fsxDataRepositoryPath, "data-repository-path", "", "", "Data Repository Path")
	_fsxCmd.Flags().StringVarP(&_fsxDeleteDataInFileSystem, "delete-data-in-file-system", "", "", "Delete Data In File System")
	_fsxCmd.Flags().StringVarP(&_fsxEnableFsxRouteTableUpdatesFromParticipantAccounts, "enable-fsx-route-table-updates-from-participant-accounts", "", "", "Enable Fsx Route Table Updates From Participant Accounts")
	_fsxCmd.Flags().StringVarP(&_fsxFileCacheId, "file-cache-id", "", "", "File Cache ID")
	_fsxCmd.Flags().StringSliceVarP(&_fsxFileCacheIds, "file-cache-ids", "", nil, "File Cache Ids")
	_fsxCmd.Flags().StringVarP(&_fsxFileCacheType, "file-cache-type", "", "", "File Cache Type")
	_fsxCmd.Flags().StringVarP(&_fsxFileCacheTypeVersion, "file-cache-type-version", "", "", "File Cache Type Version")
	_fsxCmd.Flags().StringVarP(&_fsxFileSystemId, "file-system-id", "", "", "File System ID")
	_fsxCmd.Flags().StringSliceVarP(&_fsxFileSystemIds, "file-system-ids", "", nil, "File System Ids")
	_fsxCmd.Flags().StringVarP(&_fsxFileSystemPath, "file-system-path", "", "", "File System Path")
	_fsxCmd.Flags().StringVarP(&_fsxFileSystemType, "file-system-type", "", "", "File System Type")
	_fsxCmd.Flags().StringVarP(&_fsxFileSystemTypeVersion, "file-system-type-version", "", "", "File System Type Version")
	_fsxCmd.Flags().StringVarP(&_fsxFilters, "filters", "", "", "Filters")
	_fsxCmd.Flags().StringVarP(&_fsxImportedFileChunkSize, "imported-file-chunk-size", "", "", "Imported File Chunk Size")
	_fsxCmd.Flags().StringVarP(&_fsxIncludeShared, "include-shared", "", "", "Include Shared")
	_fsxCmd.Flags().StringVarP(&_fsxKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_fsxCmd.Flags().StringVarP(&_fsxLustreConfiguration, "lustre-configuration", "", "", "Lustre Configuration")
	_fsxCmd.Flags().StringVarP(&_fsxMaxResults, "max-results", "", "", "Max Results")
	_fsxCmd.Flags().StringVarP(&_fsxName, "name", "", "", "Name")
	_fsxCmd.Flags().StringSliceVarP(&_fsxNames, "names", "", nil, "Names")
	_fsxCmd.Flags().StringVarP(&_fsxNetworkType, "network-type", "", "", "Network Type")
	_fsxCmd.Flags().StringVarP(&_fsxNextToken, "next-token", "", "", "Next Token")
	_fsxCmd.Flags().StringVarP(&_fsxOntapConfiguration, "ontap-configuration", "", "", "Ontap Configuration")
	_fsxCmd.Flags().StringVarP(&_fsxOpenZFSConfiguration, "open-zfs-configuration", "", "", "Open Zfs Configuration")
	_fsxCmd.Flags().StringVarP(&_fsxOptions, "options", "", "", "Options")
	_fsxCmd.Flags().StringSliceVarP(&_fsxPaths, "paths", "", nil, "Paths")
	_fsxCmd.Flags().StringVarP(&_fsxReleaseConfiguration, "release-configuration", "", "", "Release Configuration")
	_fsxCmd.Flags().StringVarP(&_fsxReport, "report", "", "", "Report")
	_fsxCmd.Flags().StringVarP(&_fsxResourceARN, "resource-arn", "", "", "Resource ARN")
	_fsxCmd.Flags().StringVarP(&_fsxRootVolumeSecurityStyle, "root-volume-security-style", "", "", "Root Volume Security Style")
	_fsxCmd.Flags().StringVarP(&_fsxS3, "s3", "", "", "S3")
	_fsxCmd.Flags().StringVarP(&_fsxS3AccessPoint, "s3-access-point", "", "", "S3 Access Point")
	_fsxCmd.Flags().StringSliceVarP(&_fsxSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_fsxCmd.Flags().StringVarP(&_fsxSnapshotId, "snapshot-id", "", "", "Snapshot ID")
	_fsxCmd.Flags().StringSliceVarP(&_fsxSnapshotIds, "snapshot-ids", "", nil, "Snapshot Ids")
	_fsxCmd.Flags().StringVarP(&_fsxSourceBackupId, "source-backup-id", "", "", "Source Backup ID")
	_fsxCmd.Flags().StringVarP(&_fsxSourceRegion, "source-region", "", "", "Source Region")
	_fsxCmd.Flags().StringVarP(&_fsxSourceSnapshotARN, "source-snapshot-arn", "", "", "Source Snapshot ARN")
	_fsxCmd.Flags().StringVarP(&_fsxStorageCapacity, "storage-capacity", "", "", "Storage Capacity")
	_fsxCmd.Flags().StringVarP(&_fsxStorageType, "storage-type", "", "", "Storage Type")
	_fsxCmd.Flags().StringVarP(&_fsxStorageVirtualMachineId, "storage-virtual-machine-id", "", "", "Storage Virtual Machine ID")
	_fsxCmd.Flags().StringSliceVarP(&_fsxStorageVirtualMachineIds, "storage-virtual-machine-ids", "", nil, "Storage Virtual Machine Ids")
	_fsxCmd.Flags().StringSliceVarP(&_fsxSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_fsxCmd.Flags().StringVarP(&_fsxSvmAdminPassword, "svm-admin-password", "", "", "Svm Admin Password")
	_fsxCmd.Flags().StringSliceVarP(&_fsxTagKeys, "tag-keys", "", nil, "Tag Keys")
	_fsxCmd.Flags().StringVarP(&_fsxTags, "tags", "", "", "Tags")
	_fsxCmd.Flags().StringVarP(&_fsxTaskId, "task-id", "", "", "Task ID")
	_fsxCmd.Flags().StringSliceVarP(&_fsxTaskIds, "task-ids", "", nil, "Task Ids")
	_fsxCmd.Flags().StringVarP(&_fsxType, "type", "", "", "Type")
	_fsxCmd.Flags().StringVarP(&_fsxVolumeId, "volume-id", "", "", "Volume ID")
	_fsxCmd.Flags().StringSliceVarP(&_fsxVolumeIds, "volume-ids", "", nil, "Volume Ids")
	_fsxCmd.Flags().StringVarP(&_fsxVolumeType, "volume-type", "", "", "Volume Type")
	_fsxCmd.Flags().StringVarP(&_fsxWindowsConfiguration, "windows-configuration", "", "", "Windows Configuration")

	_fsxCmd.Flags().BoolVarP(&_fsxAssociateFileSystemAliases, "associate-file-system-aliases", "", false, "Associate File System Aliases")
	_fsxCmd.Flags().BoolVarP(&_fsxCancelDataRepositoryTask, "cancel-data-repository-task", "", false, "Cancel Data Repository Task")
	_fsxCmd.Flags().BoolVarP(&_fsxCopyBackup, "copy-backup", "", false, "Copy Backup")
	_fsxCmd.Flags().BoolVarP(&_fsxCopySnapshotAndUpdateVolume, "copy-snapshot-and-update-volume", "", false, "Copy Snapshot And Update Volume")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateAndAttachS3AccessPoint, "create-and-attach-s3-access-point", "", false, "Create And Attach S3 Access Point")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateBackup, "create-backup", "", false, "Create Backup")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateDataRepositoryAssociation, "create-data-repository-association", "", false, "Create Data Repository Association")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateDataRepositoryTask, "create-data-repository-task", "", false, "Create Data Repository Task")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateFileCache, "create-file-cache", "", false, "Create File Cache")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateFileSystem, "create-file-system", "", false, "Create File System")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateFileSystemFromBackup, "create-file-system-from-backup", "", false, "Create File System From Backup")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateSnapshot, "create-snapshot", "", false, "Create Snapshot")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateStorageVirtualMachine, "create-storage-virtual-machine", "", false, "Create Storage Virtual Machine")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateVolume, "create-volume", "", false, "Create Volume")
	_fsxCmd.Flags().BoolVarP(&_fsxCreateVolumeFromBackup, "create-volume-from-backup", "", false, "Create Volume From Backup")
	_fsxCmd.Flags().BoolVarP(&_fsxDeleteBackup, "delete-backup", "", false, "Delete Backup")
	_fsxCmd.Flags().BoolVarP(&_fsxDeleteDataRepositoryAssociation, "delete-data-repository-association", "", false, "Delete Data Repository Association")
	_fsxCmd.Flags().BoolVarP(&_fsxDeleteFileCache, "delete-file-cache", "", false, "Delete File Cache")
	_fsxCmd.Flags().BoolVarP(&_fsxDeleteFileSystem, "delete-file-system", "", false, "Delete File System")
	_fsxCmd.Flags().BoolVarP(&_fsxDeleteSnapshot, "delete-snapshot", "", false, "Delete Snapshot")
	_fsxCmd.Flags().BoolVarP(&_fsxDeleteStorageVirtualMachine, "delete-storage-virtual-machine", "", false, "Delete Storage Virtual Machine")
	_fsxCmd.Flags().BoolVarP(&_fsxDeleteVolume, "delete-volume", "", false, "Delete Volume")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeBackups, "describe-backups", "", false, "Describe Backups")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeDataRepositoryAssociations, "describe-data-repository-associations", "", false, "Describe Data Repository Associations")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeDataRepositoryTasks, "describe-data-repository-tasks", "", false, "Describe Data Repository Tasks")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeFileCaches, "describe-file-caches", "", false, "Describe File Caches")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeFileSystemAliases, "describe-file-system-aliases", "", false, "Describe File System Aliases")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeFileSystems, "describe-file-systems", "", false, "Describe File Systems")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeS3AccessPointAttachments, "describe-s3-access-point-attachments", "", false, "Describe S3 Access Point Attachments")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeSharedVpcConfiguration, "describe-shared-vpc-configuration", "", false, "Describe Shared VPC Configuration")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeSnapshots, "describe-snapshots", "", false, "Describe Snapshots")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeStorageVirtualMachines, "describe-storage-virtual-machines", "", false, "Describe Storage Virtual Machines")
	_fsxCmd.Flags().BoolVarP(&_fsxDescribeVolumes, "describe-volumes", "", false, "Describe Volumes")
	_fsxCmd.Flags().BoolVarP(&_fsxDetachAndDeleteS3AccessPoint, "detach-and-delete-s3-access-point", "", false, "Detach And Delete S3 Access Point")
	_fsxCmd.Flags().BoolVarP(&_fsxDisassociateFileSystemAliases, "disassociate-file-system-aliases", "", false, "Disassociate File System Aliases")
	_fsxCmd.Flags().BoolVarP(&_fsxListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_fsxCmd.Flags().BoolVarP(&_fsxReleaseFileSystemNfsV3Locks, "release-file-system-nfs-v3-locks", "", false, "Release File System Nfs V3 Locks")
	_fsxCmd.Flags().BoolVarP(&_fsxRestoreVolumeFromSnapshot, "restore-volume-from-snapshot", "", false, "Restore Volume From Snapshot")
	_fsxCmd.Flags().BoolVarP(&_fsxStartMisconfiguredStateRecovery, "start-misconfigured-state-recovery", "", false, "Start Misconfigured State Recovery")
	_fsxCmd.Flags().BoolVarP(&_fsxTagResource, "tag-resource", "", false, "Tag Resource")
	_fsxCmd.Flags().BoolVarP(&_fsxUntagResource, "untag-resource", "", false, "Untag Resource")
	_fsxCmd.Flags().BoolVarP(&_fsxUpdateDataRepositoryAssociation, "update-data-repository-association", "", false, "Update Data Repository Association")
	_fsxCmd.Flags().BoolVarP(&_fsxUpdateFileCache, "update-file-cache", "", false, "Update File Cache")
	_fsxCmd.Flags().BoolVarP(&_fsxUpdateFileSystem, "update-file-system", "", false, "Update File System")
	_fsxCmd.Flags().BoolVarP(&_fsxUpdateSharedVpcConfiguration, "update-shared-vpc-configuration", "", false, "Update Shared VPC Configuration")
	_fsxCmd.Flags().BoolVarP(&_fsxUpdateSnapshot, "update-snapshot", "", false, "Update Snapshot")
	_fsxCmd.Flags().BoolVarP(&_fsxUpdateStorageVirtualMachine, "update-storage-virtual-machine", "", false, "Update Storage Virtual Machine")
	_fsxCmd.Flags().BoolVarP(&_fsxUpdateVolume, "update-volume", "", false, "Update Volume")

}
