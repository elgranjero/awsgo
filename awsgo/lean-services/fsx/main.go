package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/fsx"
)

var fields_associate_file_system_aliases = []leanruntime.Field{
	{Name: "Aliases", Flag: "aliases", Type: "[]string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
}

var fields_cancel_data_repository_task = []leanruntime.Field{
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_copy_backup = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CopyTags", Flag: "copy-tags", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "SourceBackupId", Flag: "source-backup-id", Type: "*string", Required: true},
	{Name: "SourceRegion", Flag: "source-region", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_copy_snapshot_and_update_volume = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CopyStrategy", Flag: "copy-strategy", Type: "types.OpenZFSCopyStrategy", Required: false},
	{Name: "Options", Flag: "options", Type: "[]types.UpdateOpenZFSVolumeOption", Required: false},
	{Name: "SourceSnapshotARN", Flag: "source-snapshot-arn", Type: "*string", Required: true},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_create_and_attach_s3_access_point = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OntapConfiguration", Flag: "ontap-configuration", Type: "*types.CreateAndAttachS3AccessPointOntapConfiguration", Required: false},
	{Name: "OpenZFSConfiguration", Flag: "open-zfs-configuration", Type: "*types.CreateAndAttachS3AccessPointOpenZFSConfiguration", Required: false},
	{Name: "S3AccessPoint", Flag: "s3-access-point", Type: "*types.CreateAndAttachS3AccessPointS3Configuration", Required: false},
	{Name: "Type", Flag: "type", Type: "types.S3AccessPointAttachmentType", Required: true},
}

var fields_create_backup = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: false},
}

var fields_create_data_repository_association = []leanruntime.Field{
	{Name: "BatchImportMetaDataOnCreate", Flag: "batch-import-meta-data-on-create", Type: "*bool", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DataRepositoryPath", Flag: "data-repository-path", Type: "*string", Required: true},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "FileSystemPath", Flag: "file-system-path", Type: "*string", Required: false},
	{Name: "ImportedFileChunkSize", Flag: "imported-file-chunk-size", Type: "*int32", Required: false},
	{Name: "S3", Flag: "s3", Type: "*types.S3DataRepositoryConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_data_repository_task = []leanruntime.Field{
	{Name: "CapacityToRelease", Flag: "capacity-to-release", Type: "*int64", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "Paths", Flag: "paths", Type: "[]string", Required: false},
	{Name: "ReleaseConfiguration", Flag: "release-configuration", Type: "*types.ReleaseConfiguration", Required: false},
	{Name: "Report", Flag: "report", Type: "*types.CompletionReport", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.DataRepositoryTaskType", Required: true},
}

var fields_create_file_cache = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "CopyTagsToDataRepositoryAssociations", Flag: "copy-tags-to-data-repository-associations", Type: "*bool", Required: false},
	{Name: "DataRepositoryAssociations", Flag: "data-repository-associations", Type: "[]types.FileCacheDataRepositoryAssociation", Required: false},
	{Name: "FileCacheType", Flag: "file-cache-type", Type: "types.FileCacheType", Required: true},
	{Name: "FileCacheTypeVersion", Flag: "file-cache-type-version", Type: "*string", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LustreConfiguration", Flag: "lustre-configuration", Type: "*types.CreateFileCacheLustreConfiguration", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "StorageCapacity", Flag: "storage-capacity", Type: "*int32", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_file_system = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemType", Flag: "file-system-type", Type: "types.FileSystemType", Required: true},
	{Name: "FileSystemTypeVersion", Flag: "file-system-type-version", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LustreConfiguration", Flag: "lustre-configuration", Type: "*types.CreateFileSystemLustreConfiguration", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "OntapConfiguration", Flag: "ontap-configuration", Type: "*types.CreateFileSystemOntapConfiguration", Required: false},
	{Name: "OpenZFSConfiguration", Flag: "open-zfs-configuration", Type: "*types.CreateFileSystemOpenZFSConfiguration", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "StorageCapacity", Flag: "storage-capacity", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "types.StorageType", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WindowsConfiguration", Flag: "windows-configuration", Type: "*types.CreateFileSystemWindowsConfiguration", Required: false},
}

var fields_create_file_system_from_backup = []leanruntime.Field{
	{Name: "BackupId", Flag: "backup-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemTypeVersion", Flag: "file-system-type-version", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "LustreConfiguration", Flag: "lustre-configuration", Type: "*types.CreateFileSystemLustreConfiguration", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "OpenZFSConfiguration", Flag: "open-zfs-configuration", Type: "*types.CreateFileSystemOpenZFSConfiguration", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: false},
	{Name: "StorageCapacity", Flag: "storage-capacity", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "types.StorageType", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "WindowsConfiguration", Flag: "windows-configuration", Type: "*types.CreateFileSystemWindowsConfiguration", Required: false},
}

var fields_create_snapshot = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_create_storage_virtual_machine = []leanruntime.Field{
	{Name: "ActiveDirectoryConfiguration", Flag: "active-directory-configuration", Type: "*types.CreateSvmActiveDirectoryConfiguration", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RootVolumeSecurityStyle", Flag: "root-volume-security-style", Type: "types.StorageVirtualMachineRootVolumeSecurityStyle", Required: false},
	{Name: "SvmAdminPassword", Flag: "svm-admin-password", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_volume = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OntapConfiguration", Flag: "ontap-configuration", Type: "*types.CreateOntapVolumeConfiguration", Required: false},
	{Name: "OpenZFSConfiguration", Flag: "open-zfs-configuration", Type: "*types.CreateOpenZFSVolumeConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeType", Flag: "volume-type", Type: "types.VolumeType", Required: true},
}

var fields_create_volume_from_backup = []leanruntime.Field{
	{Name: "BackupId", Flag: "backup-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OntapConfiguration", Flag: "ontap-configuration", Type: "*types.CreateOntapVolumeConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_backup = []leanruntime.Field{
	{Name: "BackupId", Flag: "backup-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
}

var fields_delete_data_repository_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "DeleteDataInFileSystem", Flag: "delete-data-in-file-system", Type: "*bool", Required: false},
}

var fields_delete_file_cache = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileCacheId", Flag: "file-cache-id", Type: "*string", Required: true},
}

var fields_delete_file_system = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "LustreConfiguration", Flag: "lustre-configuration", Type: "*types.DeleteFileSystemLustreConfiguration", Required: false},
	{Name: "OpenZFSConfiguration", Flag: "open-zfs-configuration", Type: "*types.DeleteFileSystemOpenZFSConfiguration", Required: false},
	{Name: "WindowsConfiguration", Flag: "windows-configuration", Type: "*types.DeleteFileSystemWindowsConfiguration", Required: false},
}

var fields_delete_snapshot = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_delete_storage_virtual_machine = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "StorageVirtualMachineId", Flag: "storage-virtual-machine-id", Type: "*string", Required: true},
}

var fields_delete_volume = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "OntapConfiguration", Flag: "ontap-configuration", Type: "*types.DeleteVolumeOntapConfiguration", Required: false},
	{Name: "OpenZFSConfiguration", Flag: "open-zfs-configuration", Type: "*types.DeleteVolumeOpenZFSConfiguration", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_describe_backups = []leanruntime.Field{
	{Name: "BackupIds", Flag: "backup-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_data_repository_associations = []leanruntime.Field{
	{Name: "AssociationIds", Flag: "association-ids", Type: "[]string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_data_repository_tasks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.DataRepositoryTaskFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TaskIds", Flag: "task-ids", Type: "[]string", Required: false},
}

var fields_describe_file_caches = []leanruntime.Field{
	{Name: "FileCacheIds", Flag: "file-cache-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_file_system_aliases = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_file_systems = []leanruntime.Field{
	{Name: "FileSystemIds", Flag: "file-system-ids", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_s3_access_point_attachments = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.S3AccessPointAttachmentsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Names", Flag: "names", Type: "[]string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_shared_vpc_configuration = []leanruntime.Field{}

var fields_describe_snapshots = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.SnapshotFilter", Required: false},
	{Name: "IncludeShared", Flag: "include-shared", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SnapshotIds", Flag: "snapshot-ids", Type: "[]string", Required: false},
}

var fields_describe_storage_virtual_machines = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.StorageVirtualMachineFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StorageVirtualMachineIds", Flag: "storage-virtual-machine-ids", Type: "[]string", Required: false},
}

var fields_describe_volumes = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.VolumeFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VolumeIds", Flag: "volume-ids", Type: "[]string", Required: false},
}

var fields_detach_and_delete_s3_access_point = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_disassociate_file_system_aliases = []leanruntime.Field{
	{Name: "Aliases", Flag: "aliases", Type: "[]string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_release_file_system_nfs_v3_locks = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
}

var fields_restore_volume_from_snapshot = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Options", Flag: "options", Type: "[]types.RestoreOpenZFSVolumeOption", Required: false},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

var fields_start_misconfigured_state_recovery = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_data_repository_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ImportedFileChunkSize", Flag: "imported-file-chunk-size", Type: "*int32", Required: false},
	{Name: "S3", Flag: "s3", Type: "*types.S3DataRepositoryConfiguration", Required: false},
}

var fields_update_file_cache = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileCacheId", Flag: "file-cache-id", Type: "*string", Required: true},
	{Name: "LustreConfiguration", Flag: "lustre-configuration", Type: "*types.UpdateFileCacheLustreConfiguration", Required: false},
}

var fields_update_file_system = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "FileSystemId", Flag: "file-system-id", Type: "*string", Required: true},
	{Name: "FileSystemTypeVersion", Flag: "file-system-type-version", Type: "*string", Required: false},
	{Name: "LustreConfiguration", Flag: "lustre-configuration", Type: "*types.UpdateFileSystemLustreConfiguration", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "OntapConfiguration", Flag: "ontap-configuration", Type: "*types.UpdateFileSystemOntapConfiguration", Required: false},
	{Name: "OpenZFSConfiguration", Flag: "open-zfs-configuration", Type: "*types.UpdateFileSystemOpenZFSConfiguration", Required: false},
	{Name: "StorageCapacity", Flag: "storage-capacity", Type: "*int32", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "types.StorageType", Required: false},
	{Name: "WindowsConfiguration", Flag: "windows-configuration", Type: "*types.UpdateFileSystemWindowsConfiguration", Required: false},
}

var fields_update_shared_vpc_configuration = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "EnableFsxRouteTableUpdatesFromParticipantAccounts", Flag: "enable-fsx-route-table-updates-from-participant-accounts", Type: "*string", Required: false},
}

var fields_update_snapshot = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_update_storage_virtual_machine = []leanruntime.Field{
	{Name: "ActiveDirectoryConfiguration", Flag: "active-directory-configuration", Type: "*types.UpdateSvmActiveDirectoryConfiguration", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "StorageVirtualMachineId", Flag: "storage-virtual-machine-id", Type: "*string", Required: true},
	{Name: "SvmAdminPassword", Flag: "svm-admin-password", Type: "*string", Required: false},
}

var fields_update_volume = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OntapConfiguration", Flag: "ontap-configuration", Type: "*types.UpdateOntapVolumeConfiguration", Required: false},
	{Name: "OpenZFSConfiguration", Flag: "open-zfs-configuration", Type: "*types.UpdateOpenZFSVolumeConfiguration", Required: false},
	{Name: "VolumeId", Flag: "volume-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-file-system-aliases": {
			Name:   "associate-file-system-aliases",
			Fields: fields_associate_file_system_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateFileSystemAliasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_file_system_aliases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateFileSystemAliases(ctx, input)
			},
		},
		"cancel-data-repository-task": {
			Name:   "cancel-data-repository-task",
			Fields: fields_cancel_data_repository_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelDataRepositoryTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_data_repository_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelDataRepositoryTask(ctx, input)
			},
		},
		"copy-backup": {
			Name:   "copy-backup",
			Fields: fields_copy_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyBackup(ctx, input)
			},
		},
		"copy-snapshot-and-update-volume": {
			Name:   "copy-snapshot-and-update-volume",
			Fields: fields_copy_snapshot_and_update_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopySnapshotAndUpdateVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_snapshot_and_update_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopySnapshotAndUpdateVolume(ctx, input)
			},
		},
		"create-and-attach-s3-access-point": {
			Name:   "create-and-attach-s3-access-point",
			Fields: fields_create_and_attach_s3_access_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAndAttachS3AccessPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_and_attach_s3_access_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAndAttachS3AccessPoint(ctx, input)
			},
		},
		"create-backup": {
			Name:   "create-backup",
			Fields: fields_create_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBackup(ctx, input)
			},
		},
		"create-data-repository-association": {
			Name:   "create-data-repository-association",
			Fields: fields_create_data_repository_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataRepositoryAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_repository_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataRepositoryAssociation(ctx, input)
			},
		},
		"create-data-repository-task": {
			Name:   "create-data-repository-task",
			Fields: fields_create_data_repository_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataRepositoryTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_repository_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataRepositoryTask(ctx, input)
			},
		},
		"create-file-cache": {
			Name:   "create-file-cache",
			Fields: fields_create_file_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFileCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_file_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFileCache(ctx, input)
			},
		},
		"create-file-system": {
			Name:   "create-file-system",
			Fields: fields_create_file_system,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFileSystemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_file_system, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFileSystem(ctx, input)
			},
		},
		"create-file-system-from-backup": {
			Name:   "create-file-system-from-backup",
			Fields: fields_create_file_system_from_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFileSystemFromBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_file_system_from_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFileSystemFromBackup(ctx, input)
			},
		},
		"create-snapshot": {
			Name:   "create-snapshot",
			Fields: fields_create_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSnapshot(ctx, input)
			},
		},
		"create-storage-virtual-machine": {
			Name:   "create-storage-virtual-machine",
			Fields: fields_create_storage_virtual_machine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStorageVirtualMachineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_storage_virtual_machine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStorageVirtualMachine(ctx, input)
			},
		},
		"create-volume": {
			Name:   "create-volume",
			Fields: fields_create_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVolume(ctx, input)
			},
		},
		"create-volume-from-backup": {
			Name:   "create-volume-from-backup",
			Fields: fields_create_volume_from_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVolumeFromBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_volume_from_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVolumeFromBackup(ctx, input)
			},
		},
		"delete-backup": {
			Name:   "delete-backup",
			Fields: fields_delete_backup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBackupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_backup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBackup(ctx, input)
			},
		},
		"delete-data-repository-association": {
			Name:   "delete-data-repository-association",
			Fields: fields_delete_data_repository_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataRepositoryAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_repository_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataRepositoryAssociation(ctx, input)
			},
		},
		"delete-file-cache": {
			Name:   "delete-file-cache",
			Fields: fields_delete_file_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFileCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_file_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFileCache(ctx, input)
			},
		},
		"delete-file-system": {
			Name:   "delete-file-system",
			Fields: fields_delete_file_system,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFileSystemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_file_system, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFileSystem(ctx, input)
			},
		},
		"delete-snapshot": {
			Name:   "delete-snapshot",
			Fields: fields_delete_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSnapshot(ctx, input)
			},
		},
		"delete-storage-virtual-machine": {
			Name:   "delete-storage-virtual-machine",
			Fields: fields_delete_storage_virtual_machine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStorageVirtualMachineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_storage_virtual_machine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStorageVirtualMachine(ctx, input)
			},
		},
		"delete-volume": {
			Name:   "delete-volume",
			Fields: fields_delete_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVolume(ctx, input)
			},
		},
		"describe-backups": {
			Name:   "describe-backups",
			Fields: fields_describe_backups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBackupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_backups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeBackups(ctx, input)
				}
				var results []*svc.DescribeBackupsOutput
				p := svc.NewDescribeBackupsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-data-repository-associations": {
			Name:   "describe-data-repository-associations",
			Fields: fields_describe_data_repository_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataRepositoryAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_data_repository_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDataRepositoryAssociations(ctx, input)
				}
				var results []*svc.DescribeDataRepositoryAssociationsOutput
				p := svc.NewDescribeDataRepositoryAssociationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-data-repository-tasks": {
			Name:   "describe-data-repository-tasks",
			Fields: fields_describe_data_repository_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDataRepositoryTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_data_repository_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDataRepositoryTasks(ctx, input)
				}
				var results []*svc.DescribeDataRepositoryTasksOutput
				p := svc.NewDescribeDataRepositoryTasksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-file-caches": {
			Name:   "describe-file-caches",
			Fields: fields_describe_file_caches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFileCachesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_file_caches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFileCaches(ctx, input)
				}
				var results []*svc.DescribeFileCachesOutput
				p := svc.NewDescribeFileCachesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-file-system-aliases": {
			Name:   "describe-file-system-aliases",
			Fields: fields_describe_file_system_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFileSystemAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_file_system_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFileSystemAliases(ctx, input)
				}
				var results []*svc.DescribeFileSystemAliasesOutput
				p := svc.NewDescribeFileSystemAliasesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-file-systems": {
			Name:   "describe-file-systems",
			Fields: fields_describe_file_systems,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFileSystemsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_file_systems, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeFileSystems(ctx, input)
				}
				var results []*svc.DescribeFileSystemsOutput
				p := svc.NewDescribeFileSystemsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-s3-access-point-attachments": {
			Name:   "describe-s3-access-point-attachments",
			Fields: fields_describe_s3_access_point_attachments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeS3AccessPointAttachmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_s3_access_point_attachments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeS3AccessPointAttachments(ctx, input)
				}
				var results []*svc.DescribeS3AccessPointAttachmentsOutput
				p := svc.NewDescribeS3AccessPointAttachmentsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-shared-vpc-configuration": {
			Name:   "describe-shared-vpc-configuration",
			Fields: fields_describe_shared_vpc_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSharedVpcConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_shared_vpc_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSharedVpcConfiguration(ctx, input)
			},
		},
		"describe-snapshots": {
			Name:   "describe-snapshots",
			Fields: fields_describe_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSnapshotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_snapshots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSnapshots(ctx, input)
				}
				var results []*svc.DescribeSnapshotsOutput
				p := svc.NewDescribeSnapshotsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-storage-virtual-machines": {
			Name:   "describe-storage-virtual-machines",
			Fields: fields_describe_storage_virtual_machines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStorageVirtualMachinesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_storage_virtual_machines, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeStorageVirtualMachines(ctx, input)
				}
				var results []*svc.DescribeStorageVirtualMachinesOutput
				p := svc.NewDescribeStorageVirtualMachinesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-volumes": {
			Name:   "describe-volumes",
			Fields: fields_describe_volumes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVolumesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_volumes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVolumes(ctx, input)
				}
				var results []*svc.DescribeVolumesOutput
				p := svc.NewDescribeVolumesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"detach-and-delete-s3-access-point": {
			Name:   "detach-and-delete-s3-access-point",
			Fields: fields_detach_and_delete_s3_access_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachAndDeleteS3AccessPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_and_delete_s3_access_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachAndDeleteS3AccessPoint(ctx, input)
			},
		},
		"disassociate-file-system-aliases": {
			Name:   "disassociate-file-system-aliases",
			Fields: fields_disassociate_file_system_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFileSystemAliasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_file_system_aliases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFileSystemAliases(ctx, input)
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"release-file-system-nfs-v3-locks": {
			Name:   "release-file-system-nfs-v3-locks",
			Fields: fields_release_file_system_nfs_v3_locks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReleaseFileSystemNfsV3LocksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_release_file_system_nfs_v3_locks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReleaseFileSystemNfsV3Locks(ctx, input)
			},
		},
		"restore-volume-from-snapshot": {
			Name:   "restore-volume-from-snapshot",
			Fields: fields_restore_volume_from_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreVolumeFromSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_volume_from_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreVolumeFromSnapshot(ctx, input)
			},
		},
		"start-misconfigured-state-recovery": {
			Name:   "start-misconfigured-state-recovery",
			Fields: fields_start_misconfigured_state_recovery,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMisconfiguredStateRecoveryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_misconfigured_state_recovery, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMisconfiguredStateRecovery(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-data-repository-association": {
			Name:   "update-data-repository-association",
			Fields: fields_update_data_repository_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataRepositoryAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_repository_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataRepositoryAssociation(ctx, input)
			},
		},
		"update-file-cache": {
			Name:   "update-file-cache",
			Fields: fields_update_file_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFileCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_file_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFileCache(ctx, input)
			},
		},
		"update-file-system": {
			Name:   "update-file-system",
			Fields: fields_update_file_system,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFileSystemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_file_system, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFileSystem(ctx, input)
			},
		},
		"update-shared-vpc-configuration": {
			Name:   "update-shared-vpc-configuration",
			Fields: fields_update_shared_vpc_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSharedVpcConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_shared_vpc_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSharedVpcConfiguration(ctx, input)
			},
		},
		"update-snapshot": {
			Name:   "update-snapshot",
			Fields: fields_update_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSnapshot(ctx, input)
			},
		},
		"update-storage-virtual-machine": {
			Name:   "update-storage-virtual-machine",
			Fields: fields_update_storage_virtual_machine,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStorageVirtualMachineInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_storage_virtual_machine, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStorageVirtualMachine(ctx, input)
			},
		},
		"update-volume": {
			Name:   "update-volume",
			Fields: fields_update_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVolume(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("fsx", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
