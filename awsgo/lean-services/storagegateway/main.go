package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/storagegateway"
)

var fields_activate_gateway = []leanruntime.Field{
	{Name: "ActivationKey", Flag: "activation-key", Type: "*string", Required: true},
	{Name: "GatewayName", Flag: "gateway-name", Type: "*string", Required: true},
	{Name: "GatewayRegion", Flag: "gateway-region", Type: "*string", Required: true},
	{Name: "GatewayTimezone", Flag: "gateway-timezone", Type: "*string", Required: true},
	{Name: "GatewayType", Flag: "gateway-type", Type: "*string", Required: false},
	{Name: "MediumChangerType", Flag: "medium-changer-type", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TapeDriveType", Flag: "tape-drive-type", Type: "*string", Required: false},
}

var fields_add_cache = []leanruntime.Field{
	{Name: "DiskIds", Flag: "disk-ids", Type: "[]string", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_add_tags_to_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_add_upload_buffer = []leanruntime.Field{
	{Name: "DiskIds", Flag: "disk-ids", Type: "[]string", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_add_working_storage = []leanruntime.Field{
	{Name: "DiskIds", Flag: "disk-ids", Type: "[]string", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_assign_tape_pool = []leanruntime.Field{
	{Name: "BypassGovernanceRetention", Flag: "bypass-governance-retention", Type: "bool", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: true},
	{Name: "TapeARN", Flag: "tape-arn", Type: "*string", Required: true},
}

var fields_associate_file_system = []leanruntime.Field{
	{Name: "AuditDestinationARN", Flag: "audit-destination-arn", Type: "*string", Required: false},
	{Name: "CacheAttributes", Flag: "cache-attributes", Type: "*types.CacheAttributes", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "EndpointNetworkConfiguration", Flag: "endpoint-network-configuration", Type: "*types.EndpointNetworkConfiguration", Required: false},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "LocationARN", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_attach_volume = []leanruntime.Field{
	{Name: "DiskId", Flag: "disk-id", Type: "*string", Required: false},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
	{Name: "TargetName", Flag: "target-name", Type: "*string", Required: false},
	{Name: "VolumeARN", Flag: "volume-arn", Type: "*string", Required: true},
}

var fields_cancel_archival = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "TapeARN", Flag: "tape-arn", Type: "*string", Required: true},
}

var fields_cancel_cache_report = []leanruntime.Field{
	{Name: "CacheReportARN", Flag: "cache-report-arn", Type: "*string", Required: true},
}

var fields_cancel_retrieval = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "TapeARN", Flag: "tape-arn", Type: "*string", Required: true},
}

var fields_create_cachedi_scsi_volume = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "KMSEncrypted", Flag: "kms-encrypted", Type: "*bool", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: false},
	{Name: "SourceVolumeARN", Flag: "source-volume-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetName", Flag: "target-name", Type: "*string", Required: true},
	{Name: "VolumeSizeInBytes", Flag: "volume-size-in-bytes", Type: "int64", Required: true},
}

var fields_create_nfs_file_share = []leanruntime.Field{
	{Name: "AuditDestinationARN", Flag: "audit-destination-arn", Type: "*string", Required: false},
	{Name: "BucketRegion", Flag: "bucket-region", Type: "*string", Required: false},
	{Name: "CacheAttributes", Flag: "cache-attributes", Type: "*types.CacheAttributes", Required: false},
	{Name: "ClientList", Flag: "client-list", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DefaultStorageClass", Flag: "default-storage-class", Type: "*string", Required: false},
	{Name: "EncryptionType", Flag: "encryption-type", Type: "types.EncryptionType", Required: false},
	{Name: "FileShareName", Flag: "file-share-name", Type: "*string", Required: false},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "GuessMIMETypeEnabled", Flag: "guess-mime-type-enabled", Type: "*bool", Required: false},
	{Name: "KMSEncrypted", Flag: "kms-encrypted", Type: "*bool", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "LocationARN", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "NFSFileShareDefaults", Flag: "nfs-file-share-defaults", Type: "*types.NFSFileShareDefaults", Required: false},
	{Name: "NotificationPolicy", Flag: "notification-policy", Type: "*string", Required: false},
	{Name: "ObjectACL", Flag: "object-acl", Type: "types.ObjectACL", Required: false},
	{Name: "ReadOnly", Flag: "read-only", Type: "*bool", Required: false},
	{Name: "RequesterPays", Flag: "requester-pays", Type: "*bool", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "Squash", Flag: "squash", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VPCEndpointDNSName", Flag: "vpc-endpoint-dns-name", Type: "*string", Required: false},
}

var fields_create_smb_file_share = []leanruntime.Field{
	{Name: "AccessBasedEnumeration", Flag: "access-based-enumeration", Type: "*bool", Required: false},
	{Name: "AdminUserList", Flag: "admin-user-list", Type: "[]string", Required: false},
	{Name: "AuditDestinationARN", Flag: "audit-destination-arn", Type: "*string", Required: false},
	{Name: "Authentication", Flag: "authentication", Type: "*string", Required: false},
	{Name: "BucketRegion", Flag: "bucket-region", Type: "*string", Required: false},
	{Name: "CacheAttributes", Flag: "cache-attributes", Type: "*types.CacheAttributes", Required: false},
	{Name: "CaseSensitivity", Flag: "case-sensitivity", Type: "types.CaseSensitivity", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DefaultStorageClass", Flag: "default-storage-class", Type: "*string", Required: false},
	{Name: "EncryptionType", Flag: "encryption-type", Type: "types.EncryptionType", Required: false},
	{Name: "FileShareName", Flag: "file-share-name", Type: "*string", Required: false},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "GuessMIMETypeEnabled", Flag: "guess-mime-type-enabled", Type: "*bool", Required: false},
	{Name: "InvalidUserList", Flag: "invalid-user-list", Type: "[]string", Required: false},
	{Name: "KMSEncrypted", Flag: "kms-encrypted", Type: "*bool", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "LocationARN", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "NotificationPolicy", Flag: "notification-policy", Type: "*string", Required: false},
	{Name: "ObjectACL", Flag: "object-acl", Type: "types.ObjectACL", Required: false},
	{Name: "OplocksEnabled", Flag: "oplocks-enabled", Type: "*bool", Required: false},
	{Name: "ReadOnly", Flag: "read-only", Type: "*bool", Required: false},
	{Name: "RequesterPays", Flag: "requester-pays", Type: "*bool", Required: false},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "SMBACLEnabled", Flag: "smbacl-enabled", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VPCEndpointDNSName", Flag: "vpc-endpoint-dns-name", Type: "*string", Required: false},
	{Name: "ValidUserList", Flag: "valid-user-list", Type: "[]string", Required: false},
}

var fields_create_snapshot = []leanruntime.Field{
	{Name: "SnapshotDescription", Flag: "snapshot-description", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeARN", Flag: "volume-arn", Type: "*string", Required: true},
}

var fields_create_snapshot_from_volume_recovery_point = []leanruntime.Field{
	{Name: "SnapshotDescription", Flag: "snapshot-description", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeARN", Flag: "volume-arn", Type: "*string", Required: true},
}

var fields_create_storedi_scsi_volume = []leanruntime.Field{
	{Name: "DiskId", Flag: "disk-id", Type: "*string", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "KMSEncrypted", Flag: "kms-encrypted", Type: "*bool", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "NetworkInterfaceId", Flag: "network-interface-id", Type: "*string", Required: true},
	{Name: "PreserveExistingData", Flag: "preserve-existing-data", Type: "bool", Required: true},
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetName", Flag: "target-name", Type: "*string", Required: true},
}

var fields_create_tape_pool = []leanruntime.Field{
	{Name: "PoolName", Flag: "pool-name", Type: "*string", Required: true},
	{Name: "RetentionLockTimeInDays", Flag: "retention-lock-time-in-days", Type: "*int32", Required: false},
	{Name: "RetentionLockType", Flag: "retention-lock-type", Type: "types.RetentionLockType", Required: false},
	{Name: "StorageClass", Flag: "storage-class", Type: "types.TapeStorageClass", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_tape_with_barcode = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "KMSEncrypted", Flag: "kms-encrypted", Type: "*bool", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TapeBarcode", Flag: "tape-barcode", Type: "*string", Required: true},
	{Name: "TapeSizeInBytes", Flag: "tape-size-in-bytes", Type: "*int64", Required: true},
	{Name: "Worm", Flag: "worm", Type: "bool", Required: false},
}

var fields_create_tapes = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "KMSEncrypted", Flag: "kms-encrypted", Type: "*bool", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "NumTapesToCreate", Flag: "num-tapes-to-create", Type: "*int32", Required: true},
	{Name: "PoolId", Flag: "pool-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TapeBarcodePrefix", Flag: "tape-barcode-prefix", Type: "*string", Required: true},
	{Name: "TapeSizeInBytes", Flag: "tape-size-in-bytes", Type: "*int64", Required: true},
	{Name: "Worm", Flag: "worm", Type: "bool", Required: false},
}

var fields_delete_automatic_tape_creation_policy = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_delete_bandwidth_rate_limit = []leanruntime.Field{
	{Name: "BandwidthType", Flag: "bandwidth-type", Type: "*string", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_delete_cache_report = []leanruntime.Field{
	{Name: "CacheReportARN", Flag: "cache-report-arn", Type: "*string", Required: true},
}

var fields_delete_chap_credentials = []leanruntime.Field{
	{Name: "InitiatorName", Flag: "initiator-name", Type: "*string", Required: true},
	{Name: "TargetARN", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_delete_file_share = []leanruntime.Field{
	{Name: "FileShareARN", Flag: "file-share-arn", Type: "*string", Required: true},
	{Name: "ForceDelete", Flag: "force-delete", Type: "bool", Required: false},
}

var fields_delete_gateway = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_delete_snapshot_schedule = []leanruntime.Field{
	{Name: "VolumeARN", Flag: "volume-arn", Type: "*string", Required: true},
}

var fields_delete_tape = []leanruntime.Field{
	{Name: "BypassGovernanceRetention", Flag: "bypass-governance-retention", Type: "bool", Required: false},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "TapeARN", Flag: "tape-arn", Type: "*string", Required: true},
}

var fields_delete_tape_archive = []leanruntime.Field{
	{Name: "BypassGovernanceRetention", Flag: "bypass-governance-retention", Type: "bool", Required: false},
	{Name: "TapeARN", Flag: "tape-arn", Type: "*string", Required: true},
}

var fields_delete_tape_pool = []leanruntime.Field{
	{Name: "PoolARN", Flag: "pool-arn", Type: "*string", Required: true},
}

var fields_delete_volume = []leanruntime.Field{
	{Name: "VolumeARN", Flag: "volume-arn", Type: "*string", Required: true},
}

var fields_describe_availability_monitor_test = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_describe_bandwidth_rate_limit = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_describe_bandwidth_rate_limit_schedule = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_describe_cache = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_describe_cache_report = []leanruntime.Field{
	{Name: "CacheReportARN", Flag: "cache-report-arn", Type: "*string", Required: true},
}

var fields_describe_cachedi_scsi_volumes = []leanruntime.Field{
	{Name: "VolumeARNs", Flag: "volume-arns", Type: "[]string", Required: true},
}

var fields_describe_chap_credentials = []leanruntime.Field{
	{Name: "TargetARN", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_describe_file_system_associations = []leanruntime.Field{
	{Name: "FileSystemAssociationARNList", Flag: "file-system-association-arn-list", Type: "[]string", Required: true},
}

var fields_describe_gateway_information = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_describe_maintenance_start_time = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_describe_nfs_file_shares = []leanruntime.Field{
	{Name: "FileShareARNList", Flag: "file-share-arn-list", Type: "[]string", Required: true},
}

var fields_describe_smb_file_shares = []leanruntime.Field{
	{Name: "FileShareARNList", Flag: "file-share-arn-list", Type: "[]string", Required: true},
}

var fields_describe_smb_settings = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_describe_snapshot_schedule = []leanruntime.Field{
	{Name: "VolumeARN", Flag: "volume-arn", Type: "*string", Required: true},
}

var fields_describe_storedi_scsi_volumes = []leanruntime.Field{
	{Name: "VolumeARNs", Flag: "volume-arns", Type: "[]string", Required: true},
}

var fields_describe_tape_archives = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "TapeARNs", Flag: "tape-arns", Type: "[]string", Required: false},
}

var fields_describe_tape_recovery_points = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_describe_tapes = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "TapeARNs", Flag: "tape-arns", Type: "[]string", Required: false},
}

var fields_describe_upload_buffer = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_describe_vtl_devices = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "VTLDeviceARNs", Flag: "vtl-device-arns", Type: "[]string", Required: false},
}

var fields_describe_working_storage = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_detach_volume = []leanruntime.Field{
	{Name: "ForceDetach", Flag: "force-detach", Type: "*bool", Required: false},
	{Name: "VolumeARN", Flag: "volume-arn", Type: "*string", Required: true},
}

var fields_disable_gateway = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_disassociate_file_system = []leanruntime.Field{
	{Name: "FileSystemAssociationARN", Flag: "file-system-association-arn", Type: "*string", Required: true},
	{Name: "ForceDelete", Flag: "force-delete", Type: "bool", Required: false},
}

var fields_evict_files_failing_upload = []leanruntime.Field{
	{Name: "FileShareARN", Flag: "file-share-arn", Type: "*string", Required: true},
	{Name: "ForceRemove", Flag: "force-remove", Type: "bool", Required: false},
}

var fields_join_domain = []leanruntime.Field{
	{Name: "DomainControllers", Flag: "domain-controllers", Type: "[]string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "OrganizationalUnit", Flag: "organizational-unit", Type: "*string", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "TimeoutInSeconds", Flag: "timeout-in-seconds", Type: "*int32", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_list_automatic_tape_creation_policies = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: false},
}

var fields_list_cache_reports = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_file_shares = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_file_system_associations = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_gateways = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_list_local_disks = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tape_pools = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "PoolARNs", Flag: "pool-arns", Type: "[]string", Required: false},
}

var fields_list_tapes = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "TapeARNs", Flag: "tape-arns", Type: "[]string", Required: false},
}

var fields_list_volume_initiators = []leanruntime.Field{
	{Name: "VolumeARN", Flag: "volume-arn", Type: "*string", Required: true},
}

var fields_list_volume_recovery_points = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_list_volumes = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
}

var fields_notify_when_uploaded = []leanruntime.Field{
	{Name: "FileShareARN", Flag: "file-share-arn", Type: "*string", Required: true},
}

var fields_refresh_cache = []leanruntime.Field{
	{Name: "FileShareARN", Flag: "file-share-arn", Type: "*string", Required: true},
	{Name: "FolderList", Flag: "folder-list", Type: "[]string", Required: false},
	{Name: "Recursive", Flag: "recursive", Type: "*bool", Required: false},
}

var fields_remove_tags_from_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_reset_cache = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_retrieve_tape_archive = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "TapeARN", Flag: "tape-arn", Type: "*string", Required: true},
}

var fields_retrieve_tape_recovery_point = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "TapeARN", Flag: "tape-arn", Type: "*string", Required: true},
}

var fields_set_local_console_password = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "LocalConsolePassword", Flag: "local-console-password", Type: "*string", Required: true},
}

var fields_set_smb_guest_password = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
}

var fields_shutdown_gateway = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_start_availability_monitor_test = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_start_cache_report = []leanruntime.Field{
	{Name: "BucketRegion", Flag: "bucket-region", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ExclusionFilters", Flag: "exclusion-filters", Type: "[]types.CacheReportFilter", Required: false},
	{Name: "FileShareARN", Flag: "file-share-arn", Type: "*string", Required: true},
	{Name: "InclusionFilters", Flag: "inclusion-filters", Type: "[]types.CacheReportFilter", Required: false},
	{Name: "LocationARN", Flag: "location-arn", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VPCEndpointDNSName", Flag: "vpc-endpoint-dns-name", Type: "*string", Required: false},
}

var fields_start_gateway = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_update_automatic_tape_creation_policy = []leanruntime.Field{
	{Name: "AutomaticTapeCreationRules", Flag: "automatic-tape-creation-rules", Type: "[]types.AutomaticTapeCreationRule", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_update_bandwidth_rate_limit = []leanruntime.Field{
	{Name: "AverageDownloadRateLimitInBitsPerSec", Flag: "average-download-rate-limit-in-bits-per-sec", Type: "*int64", Required: false},
	{Name: "AverageUploadRateLimitInBitsPerSec", Flag: "average-upload-rate-limit-in-bits-per-sec", Type: "*int64", Required: false},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_update_bandwidth_rate_limit_schedule = []leanruntime.Field{
	{Name: "BandwidthRateLimitIntervals", Flag: "bandwidth-rate-limit-intervals", Type: "[]types.BandwidthRateLimitInterval", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_update_chap_credentials = []leanruntime.Field{
	{Name: "InitiatorName", Flag: "initiator-name", Type: "*string", Required: true},
	{Name: "SecretToAuthenticateInitiator", Flag: "secret-to-authenticate-initiator", Type: "*string", Required: true},
	{Name: "SecretToAuthenticateTarget", Flag: "secret-to-authenticate-target", Type: "*string", Required: false},
	{Name: "TargetARN", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_update_file_system_association = []leanruntime.Field{
	{Name: "AuditDestinationARN", Flag: "audit-destination-arn", Type: "*string", Required: false},
	{Name: "CacheAttributes", Flag: "cache-attributes", Type: "*types.CacheAttributes", Required: false},
	{Name: "FileSystemAssociationARN", Flag: "file-system-association-arn", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_update_gateway_information = []leanruntime.Field{
	{Name: "CloudWatchLogGroupARN", Flag: "cloud-watch-log-group-arn", Type: "*string", Required: false},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "GatewayCapacity", Flag: "gateway-capacity", Type: "types.GatewayCapacity", Required: false},
	{Name: "GatewayName", Flag: "gateway-name", Type: "*string", Required: false},
	{Name: "GatewayTimezone", Flag: "gateway-timezone", Type: "*string", Required: false},
}

var fields_update_gateway_software_now = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_update_maintenance_start_time = []leanruntime.Field{
	{Name: "DayOfMonth", Flag: "day-of-month", Type: "*int32", Required: false},
	{Name: "DayOfWeek", Flag: "day-of-week", Type: "*int32", Required: false},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "HourOfDay", Flag: "hour-of-day", Type: "*int32", Required: false},
	{Name: "MinuteOfHour", Flag: "minute-of-hour", Type: "*int32", Required: false},
	{Name: "SoftwareUpdatePreferences", Flag: "software-update-preferences", Type: "*types.SoftwareUpdatePreferences", Required: false},
}

var fields_update_nfs_file_share = []leanruntime.Field{
	{Name: "AuditDestinationARN", Flag: "audit-destination-arn", Type: "*string", Required: false},
	{Name: "CacheAttributes", Flag: "cache-attributes", Type: "*types.CacheAttributes", Required: false},
	{Name: "ClientList", Flag: "client-list", Type: "[]string", Required: false},
	{Name: "DefaultStorageClass", Flag: "default-storage-class", Type: "*string", Required: false},
	{Name: "EncryptionType", Flag: "encryption-type", Type: "types.EncryptionType", Required: false},
	{Name: "FileShareARN", Flag: "file-share-arn", Type: "*string", Required: true},
	{Name: "FileShareName", Flag: "file-share-name", Type: "*string", Required: false},
	{Name: "GuessMIMETypeEnabled", Flag: "guess-mime-type-enabled", Type: "*bool", Required: false},
	{Name: "KMSEncrypted", Flag: "kms-encrypted", Type: "*bool", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "NFSFileShareDefaults", Flag: "nfs-file-share-defaults", Type: "*types.NFSFileShareDefaults", Required: false},
	{Name: "NotificationPolicy", Flag: "notification-policy", Type: "*string", Required: false},
	{Name: "ObjectACL", Flag: "object-acl", Type: "types.ObjectACL", Required: false},
	{Name: "ReadOnly", Flag: "read-only", Type: "*bool", Required: false},
	{Name: "RequesterPays", Flag: "requester-pays", Type: "*bool", Required: false},
	{Name: "Squash", Flag: "squash", Type: "*string", Required: false},
}

var fields_update_smb_file_share = []leanruntime.Field{
	{Name: "AccessBasedEnumeration", Flag: "access-based-enumeration", Type: "*bool", Required: false},
	{Name: "AdminUserList", Flag: "admin-user-list", Type: "[]string", Required: false},
	{Name: "AuditDestinationARN", Flag: "audit-destination-arn", Type: "*string", Required: false},
	{Name: "CacheAttributes", Flag: "cache-attributes", Type: "*types.CacheAttributes", Required: false},
	{Name: "CaseSensitivity", Flag: "case-sensitivity", Type: "types.CaseSensitivity", Required: false},
	{Name: "DefaultStorageClass", Flag: "default-storage-class", Type: "*string", Required: false},
	{Name: "EncryptionType", Flag: "encryption-type", Type: "types.EncryptionType", Required: false},
	{Name: "FileShareARN", Flag: "file-share-arn", Type: "*string", Required: true},
	{Name: "FileShareName", Flag: "file-share-name", Type: "*string", Required: false},
	{Name: "GuessMIMETypeEnabled", Flag: "guess-mime-type-enabled", Type: "*bool", Required: false},
	{Name: "InvalidUserList", Flag: "invalid-user-list", Type: "[]string", Required: false},
	{Name: "KMSEncrypted", Flag: "kms-encrypted", Type: "*bool", Required: false},
	{Name: "KMSKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "NotificationPolicy", Flag: "notification-policy", Type: "*string", Required: false},
	{Name: "ObjectACL", Flag: "object-acl", Type: "types.ObjectACL", Required: false},
	{Name: "OplocksEnabled", Flag: "oplocks-enabled", Type: "*bool", Required: false},
	{Name: "ReadOnly", Flag: "read-only", Type: "*bool", Required: false},
	{Name: "RequesterPays", Flag: "requester-pays", Type: "*bool", Required: false},
	{Name: "SMBACLEnabled", Flag: "smbacl-enabled", Type: "*bool", Required: false},
	{Name: "ValidUserList", Flag: "valid-user-list", Type: "[]string", Required: false},
}

var fields_update_smb_file_share_visibility = []leanruntime.Field{
	{Name: "FileSharesVisible", Flag: "file-shares-visible", Type: "*bool", Required: true},
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
}

var fields_update_smb_local_groups = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "SMBLocalGroups", Flag: "smb-local-groups", Type: "*types.SMBLocalGroups", Required: true},
}

var fields_update_smb_security_strategy = []leanruntime.Field{
	{Name: "GatewayARN", Flag: "gateway-arn", Type: "*string", Required: true},
	{Name: "SMBSecurityStrategy", Flag: "smb-security-strategy", Type: "types.SMBSecurityStrategy", Required: true},
}

var fields_update_snapshot_schedule = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "RecurrenceInHours", Flag: "recurrence-in-hours", Type: "*int32", Required: true},
	{Name: "StartAt", Flag: "start-at", Type: "*int32", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VolumeARN", Flag: "volume-arn", Type: "*string", Required: true},
}

var fields_update_vtl_device_type = []leanruntime.Field{
	{Name: "DeviceType", Flag: "device-type", Type: "*string", Required: true},
	{Name: "VTLDeviceARN", Flag: "vtl-device-arn", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"activate-gateway": {
			Name:   "activate-gateway",
			Fields: fields_activate_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ActivateGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_activate_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ActivateGateway(ctx, input)
			},
		},
		"add-cache": {
			Name:   "add-cache",
			Fields: fields_add_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddCache(ctx, input)
			},
		},
		"add-tags-to-resource": {
			Name:   "add-tags-to-resource",
			Fields: fields_add_tags_to_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsToResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags_to_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTagsToResource(ctx, input)
			},
		},
		"add-upload-buffer": {
			Name:   "add-upload-buffer",
			Fields: fields_add_upload_buffer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddUploadBufferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_upload_buffer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddUploadBuffer(ctx, input)
			},
		},
		"add-working-storage": {
			Name:   "add-working-storage",
			Fields: fields_add_working_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddWorkingStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_working_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddWorkingStorage(ctx, input)
			},
		},
		"assign-tape-pool": {
			Name:   "assign-tape-pool",
			Fields: fields_assign_tape_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssignTapePoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assign_tape_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssignTapePool(ctx, input)
			},
		},
		"associate-file-system": {
			Name:   "associate-file-system",
			Fields: fields_associate_file_system,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateFileSystemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_file_system, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateFileSystem(ctx, input)
			},
		},
		"attach-volume": {
			Name:   "attach-volume",
			Fields: fields_attach_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachVolume(ctx, input)
			},
		},
		"cancel-archival": {
			Name:   "cancel-archival",
			Fields: fields_cancel_archival,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelArchivalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_archival, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelArchival(ctx, input)
			},
		},
		"cancel-cache-report": {
			Name:   "cancel-cache-report",
			Fields: fields_cancel_cache_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelCacheReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_cache_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelCacheReport(ctx, input)
			},
		},
		"cancel-retrieval": {
			Name:   "cancel-retrieval",
			Fields: fields_cancel_retrieval,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelRetrievalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_retrieval, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelRetrieval(ctx, input)
			},
		},
		"create-cachedi-scsi-volume": {
			Name:   "create-cachedi-scsi-volume",
			Fields: fields_create_cachedi_scsi_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCachediSCSIVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cachedi_scsi_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCachediSCSIVolume(ctx, input)
			},
		},
		"create-nfs-file-share": {
			Name:   "create-nfs-file-share",
			Fields: fields_create_nfs_file_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNFSFileShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_nfs_file_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNFSFileShare(ctx, input)
			},
		},
		"create-smb-file-share": {
			Name:   "create-smb-file-share",
			Fields: fields_create_smb_file_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSMBFileShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_smb_file_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSMBFileShare(ctx, input)
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
		"create-snapshot-from-volume-recovery-point": {
			Name:   "create-snapshot-from-volume-recovery-point",
			Fields: fields_create_snapshot_from_volume_recovery_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSnapshotFromVolumeRecoveryPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_snapshot_from_volume_recovery_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSnapshotFromVolumeRecoveryPoint(ctx, input)
			},
		},
		"create-storedi-scsi-volume": {
			Name:   "create-storedi-scsi-volume",
			Fields: fields_create_storedi_scsi_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStorediSCSIVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_storedi_scsi_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStorediSCSIVolume(ctx, input)
			},
		},
		"create-tape-pool": {
			Name:   "create-tape-pool",
			Fields: fields_create_tape_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTapePoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tape_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTapePool(ctx, input)
			},
		},
		"create-tape-with-barcode": {
			Name:   "create-tape-with-barcode",
			Fields: fields_create_tape_with_barcode,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTapeWithBarcodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tape_with_barcode, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTapeWithBarcode(ctx, input)
			},
		},
		"create-tapes": {
			Name:   "create-tapes",
			Fields: fields_create_tapes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTapesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tapes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTapes(ctx, input)
			},
		},
		"delete-automatic-tape-creation-policy": {
			Name:   "delete-automatic-tape-creation-policy",
			Fields: fields_delete_automatic_tape_creation_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAutomaticTapeCreationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_automatic_tape_creation_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAutomaticTapeCreationPolicy(ctx, input)
			},
		},
		"delete-bandwidth-rate-limit": {
			Name:   "delete-bandwidth-rate-limit",
			Fields: fields_delete_bandwidth_rate_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBandwidthRateLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bandwidth_rate_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBandwidthRateLimit(ctx, input)
			},
		},
		"delete-cache-report": {
			Name:   "delete-cache-report",
			Fields: fields_delete_cache_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCacheReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cache_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCacheReport(ctx, input)
			},
		},
		"delete-chap-credentials": {
			Name:   "delete-chap-credentials",
			Fields: fields_delete_chap_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteChapCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_chap_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteChapCredentials(ctx, input)
			},
		},
		"delete-file-share": {
			Name:   "delete-file-share",
			Fields: fields_delete_file_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFileShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_file_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFileShare(ctx, input)
			},
		},
		"delete-gateway": {
			Name:   "delete-gateway",
			Fields: fields_delete_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGateway(ctx, input)
			},
		},
		"delete-snapshot-schedule": {
			Name:   "delete-snapshot-schedule",
			Fields: fields_delete_snapshot_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSnapshotScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_snapshot_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSnapshotSchedule(ctx, input)
			},
		},
		"delete-tape": {
			Name:   "delete-tape",
			Fields: fields_delete_tape,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTapeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tape, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTape(ctx, input)
			},
		},
		"delete-tape-archive": {
			Name:   "delete-tape-archive",
			Fields: fields_delete_tape_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTapeArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tape_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTapeArchive(ctx, input)
			},
		},
		"delete-tape-pool": {
			Name:   "delete-tape-pool",
			Fields: fields_delete_tape_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTapePoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tape_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTapePool(ctx, input)
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
		"describe-availability-monitor-test": {
			Name:   "describe-availability-monitor-test",
			Fields: fields_describe_availability_monitor_test,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAvailabilityMonitorTestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_availability_monitor_test, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAvailabilityMonitorTest(ctx, input)
			},
		},
		"describe-bandwidth-rate-limit": {
			Name:   "describe-bandwidth-rate-limit",
			Fields: fields_describe_bandwidth_rate_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBandwidthRateLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bandwidth_rate_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBandwidthRateLimit(ctx, input)
			},
		},
		"describe-bandwidth-rate-limit-schedule": {
			Name:   "describe-bandwidth-rate-limit-schedule",
			Fields: fields_describe_bandwidth_rate_limit_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBandwidthRateLimitScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_bandwidth_rate_limit_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBandwidthRateLimitSchedule(ctx, input)
			},
		},
		"describe-cache": {
			Name:   "describe-cache",
			Fields: fields_describe_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCache(ctx, input)
			},
		},
		"describe-cache-report": {
			Name:   "describe-cache-report",
			Fields: fields_describe_cache_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCacheReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cache_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCacheReport(ctx, input)
			},
		},
		"describe-cachedi-scsi-volumes": {
			Name:   "describe-cachedi-scsi-volumes",
			Fields: fields_describe_cachedi_scsi_volumes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCachediSCSIVolumesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cachedi_scsi_volumes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCachediSCSIVolumes(ctx, input)
			},
		},
		"describe-chap-credentials": {
			Name:   "describe-chap-credentials",
			Fields: fields_describe_chap_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeChapCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_chap_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeChapCredentials(ctx, input)
			},
		},
		"describe-file-system-associations": {
			Name:   "describe-file-system-associations",
			Fields: fields_describe_file_system_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFileSystemAssociationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_file_system_associations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFileSystemAssociations(ctx, input)
			},
		},
		"describe-gateway-information": {
			Name:   "describe-gateway-information",
			Fields: fields_describe_gateway_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGatewayInformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_gateway_information, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGatewayInformation(ctx, input)
			},
		},
		"describe-maintenance-start-time": {
			Name:   "describe-maintenance-start-time",
			Fields: fields_describe_maintenance_start_time,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMaintenanceStartTimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_maintenance_start_time, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMaintenanceStartTime(ctx, input)
			},
		},
		"describe-nfs-file-shares": {
			Name:   "describe-nfs-file-shares",
			Fields: fields_describe_nfs_file_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNFSFileSharesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_nfs_file_shares, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNFSFileShares(ctx, input)
			},
		},
		"describe-smb-file-shares": {
			Name:   "describe-smb-file-shares",
			Fields: fields_describe_smb_file_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSMBFileSharesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_smb_file_shares, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSMBFileShares(ctx, input)
			},
		},
		"describe-smb-settings": {
			Name:   "describe-smb-settings",
			Fields: fields_describe_smb_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSMBSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_smb_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSMBSettings(ctx, input)
			},
		},
		"describe-snapshot-schedule": {
			Name:   "describe-snapshot-schedule",
			Fields: fields_describe_snapshot_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSnapshotScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_snapshot_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSnapshotSchedule(ctx, input)
			},
		},
		"describe-storedi-scsi-volumes": {
			Name:   "describe-storedi-scsi-volumes",
			Fields: fields_describe_storedi_scsi_volumes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeStorediSCSIVolumesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_storedi_scsi_volumes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeStorediSCSIVolumes(ctx, input)
			},
		},
		"describe-tape-archives": {
			Name:   "describe-tape-archives",
			Fields: fields_describe_tape_archives,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTapeArchivesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_tape_archives, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTapeArchives(ctx, input)
				}
				var results []*svc.DescribeTapeArchivesOutput
				p := svc.NewDescribeTapeArchivesPaginator(client, input)
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
		"describe-tape-recovery-points": {
			Name:   "describe-tape-recovery-points",
			Fields: fields_describe_tape_recovery_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTapeRecoveryPointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_tape_recovery_points, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTapeRecoveryPoints(ctx, input)
				}
				var results []*svc.DescribeTapeRecoveryPointsOutput
				p := svc.NewDescribeTapeRecoveryPointsPaginator(client, input)
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
		"describe-tapes": {
			Name:   "describe-tapes",
			Fields: fields_describe_tapes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTapesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_tapes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTapes(ctx, input)
				}
				var results []*svc.DescribeTapesOutput
				p := svc.NewDescribeTapesPaginator(client, input)
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
		"describe-upload-buffer": {
			Name:   "describe-upload-buffer",
			Fields: fields_describe_upload_buffer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUploadBufferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_upload_buffer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUploadBuffer(ctx, input)
			},
		},
		"describe-vtl-devices": {
			Name:   "describe-vtl-devices",
			Fields: fields_describe_vtl_devices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVTLDevicesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_vtl_devices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeVTLDevices(ctx, input)
				}
				var results []*svc.DescribeVTLDevicesOutput
				p := svc.NewDescribeVTLDevicesPaginator(client, input)
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
		"describe-working-storage": {
			Name:   "describe-working-storage",
			Fields: fields_describe_working_storage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWorkingStorageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_working_storage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWorkingStorage(ctx, input)
			},
		},
		"detach-volume": {
			Name:   "detach-volume",
			Fields: fields_detach_volume,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachVolumeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_volume, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachVolume(ctx, input)
			},
		},
		"disable-gateway": {
			Name:   "disable-gateway",
			Fields: fields_disable_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableGateway(ctx, input)
			},
		},
		"disassociate-file-system": {
			Name:   "disassociate-file-system",
			Fields: fields_disassociate_file_system,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFileSystemInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_file_system, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFileSystem(ctx, input)
			},
		},
		"evict-files-failing-upload": {
			Name:   "evict-files-failing-upload",
			Fields: fields_evict_files_failing_upload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EvictFilesFailingUploadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_evict_files_failing_upload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EvictFilesFailingUpload(ctx, input)
			},
		},
		"join-domain": {
			Name:   "join-domain",
			Fields: fields_join_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.JoinDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_join_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.JoinDomain(ctx, input)
			},
		},
		"list-automatic-tape-creation-policies": {
			Name:   "list-automatic-tape-creation-policies",
			Fields: fields_list_automatic_tape_creation_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAutomaticTapeCreationPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_automatic_tape_creation_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAutomaticTapeCreationPolicies(ctx, input)
			},
		},
		"list-cache-reports": {
			Name:   "list-cache-reports",
			Fields: fields_list_cache_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCacheReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cache_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCacheReports(ctx, input)
				}
				var results []*svc.ListCacheReportsOutput
				p := svc.NewListCacheReportsPaginator(client, input)
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
		"list-file-shares": {
			Name:   "list-file-shares",
			Fields: fields_list_file_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFileSharesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_file_shares, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFileShares(ctx, input)
				}
				var results []*svc.ListFileSharesOutput
				p := svc.NewListFileSharesPaginator(client, input)
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
		"list-file-system-associations": {
			Name:   "list-file-system-associations",
			Fields: fields_list_file_system_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFileSystemAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_file_system_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFileSystemAssociations(ctx, input)
				}
				var results []*svc.ListFileSystemAssociationsOutput
				p := svc.NewListFileSystemAssociationsPaginator(client, input)
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
		"list-gateways": {
			Name:   "list-gateways",
			Fields: fields_list_gateways,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGatewaysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_gateways, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGateways(ctx, input)
				}
				var results []*svc.ListGatewaysOutput
				p := svc.NewListGatewaysPaginator(client, input)
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
		"list-local-disks": {
			Name:   "list-local-disks",
			Fields: fields_list_local_disks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLocalDisksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_local_disks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLocalDisks(ctx, input)
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
		"list-tape-pools": {
			Name:   "list-tape-pools",
			Fields: fields_list_tape_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTapePoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tape_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTapePools(ctx, input)
				}
				var results []*svc.ListTapePoolsOutput
				p := svc.NewListTapePoolsPaginator(client, input)
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
		"list-tapes": {
			Name:   "list-tapes",
			Fields: fields_list_tapes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTapesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tapes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTapes(ctx, input)
				}
				var results []*svc.ListTapesOutput
				p := svc.NewListTapesPaginator(client, input)
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
		"list-volume-initiators": {
			Name:   "list-volume-initiators",
			Fields: fields_list_volume_initiators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVolumeInitiatorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_volume_initiators, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVolumeInitiators(ctx, input)
			},
		},
		"list-volume-recovery-points": {
			Name:   "list-volume-recovery-points",
			Fields: fields_list_volume_recovery_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVolumeRecoveryPointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_volume_recovery_points, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVolumeRecoveryPoints(ctx, input)
			},
		},
		"list-volumes": {
			Name:   "list-volumes",
			Fields: fields_list_volumes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVolumesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_volumes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVolumes(ctx, input)
				}
				var results []*svc.ListVolumesOutput
				p := svc.NewListVolumesPaginator(client, input)
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
		"notify-when-uploaded": {
			Name:   "notify-when-uploaded",
			Fields: fields_notify_when_uploaded,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.NotifyWhenUploadedInput{}
				if _, err := leanruntime.ApplyInput(input, fields_notify_when_uploaded, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.NotifyWhenUploaded(ctx, input)
			},
		},
		"refresh-cache": {
			Name:   "refresh-cache",
			Fields: fields_refresh_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RefreshCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_refresh_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RefreshCache(ctx, input)
			},
		},
		"remove-tags-from-resource": {
			Name:   "remove-tags-from-resource",
			Fields: fields_remove_tags_from_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsFromResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags_from_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTagsFromResource(ctx, input)
			},
		},
		"reset-cache": {
			Name:   "reset-cache",
			Fields: fields_reset_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetCache(ctx, input)
			},
		},
		"retrieve-tape-archive": {
			Name:   "retrieve-tape-archive",
			Fields: fields_retrieve_tape_archive,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetrieveTapeArchiveInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retrieve_tape_archive, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetrieveTapeArchive(ctx, input)
			},
		},
		"retrieve-tape-recovery-point": {
			Name:   "retrieve-tape-recovery-point",
			Fields: fields_retrieve_tape_recovery_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetrieveTapeRecoveryPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retrieve_tape_recovery_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetrieveTapeRecoveryPoint(ctx, input)
			},
		},
		"set-local-console-password": {
			Name:   "set-local-console-password",
			Fields: fields_set_local_console_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetLocalConsolePasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_local_console_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetLocalConsolePassword(ctx, input)
			},
		},
		"set-smb-guest-password": {
			Name:   "set-smb-guest-password",
			Fields: fields_set_smb_guest_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetSMBGuestPasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_smb_guest_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetSMBGuestPassword(ctx, input)
			},
		},
		"shutdown-gateway": {
			Name:   "shutdown-gateway",
			Fields: fields_shutdown_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ShutdownGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_shutdown_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ShutdownGateway(ctx, input)
			},
		},
		"start-availability-monitor-test": {
			Name:   "start-availability-monitor-test",
			Fields: fields_start_availability_monitor_test,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAvailabilityMonitorTestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_availability_monitor_test, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAvailabilityMonitorTest(ctx, input)
			},
		},
		"start-cache-report": {
			Name:   "start-cache-report",
			Fields: fields_start_cache_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCacheReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_cache_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCacheReport(ctx, input)
			},
		},
		"start-gateway": {
			Name:   "start-gateway",
			Fields: fields_start_gateway,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartGatewayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_gateway, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartGateway(ctx, input)
			},
		},
		"update-automatic-tape-creation-policy": {
			Name:   "update-automatic-tape-creation-policy",
			Fields: fields_update_automatic_tape_creation_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAutomaticTapeCreationPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_automatic_tape_creation_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAutomaticTapeCreationPolicy(ctx, input)
			},
		},
		"update-bandwidth-rate-limit": {
			Name:   "update-bandwidth-rate-limit",
			Fields: fields_update_bandwidth_rate_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBandwidthRateLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bandwidth_rate_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBandwidthRateLimit(ctx, input)
			},
		},
		"update-bandwidth-rate-limit-schedule": {
			Name:   "update-bandwidth-rate-limit-schedule",
			Fields: fields_update_bandwidth_rate_limit_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBandwidthRateLimitScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bandwidth_rate_limit_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBandwidthRateLimitSchedule(ctx, input)
			},
		},
		"update-chap-credentials": {
			Name:   "update-chap-credentials",
			Fields: fields_update_chap_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateChapCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_chap_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateChapCredentials(ctx, input)
			},
		},
		"update-file-system-association": {
			Name:   "update-file-system-association",
			Fields: fields_update_file_system_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFileSystemAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_file_system_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFileSystemAssociation(ctx, input)
			},
		},
		"update-gateway-information": {
			Name:   "update-gateway-information",
			Fields: fields_update_gateway_information,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewayInformationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway_information, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGatewayInformation(ctx, input)
			},
		},
		"update-gateway-software-now": {
			Name:   "update-gateway-software-now",
			Fields: fields_update_gateway_software_now,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGatewaySoftwareNowInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_gateway_software_now, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGatewaySoftwareNow(ctx, input)
			},
		},
		"update-maintenance-start-time": {
			Name:   "update-maintenance-start-time",
			Fields: fields_update_maintenance_start_time,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMaintenanceStartTimeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_maintenance_start_time, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMaintenanceStartTime(ctx, input)
			},
		},
		"update-nfs-file-share": {
			Name:   "update-nfs-file-share",
			Fields: fields_update_nfs_file_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNFSFileShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_nfs_file_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNFSFileShare(ctx, input)
			},
		},
		"update-smb-file-share": {
			Name:   "update-smb-file-share",
			Fields: fields_update_smb_file_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSMBFileShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_smb_file_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSMBFileShare(ctx, input)
			},
		},
		"update-smb-file-share-visibility": {
			Name:   "update-smb-file-share-visibility",
			Fields: fields_update_smb_file_share_visibility,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSMBFileShareVisibilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_smb_file_share_visibility, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSMBFileShareVisibility(ctx, input)
			},
		},
		"update-smb-local-groups": {
			Name:   "update-smb-local-groups",
			Fields: fields_update_smb_local_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSMBLocalGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_smb_local_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSMBLocalGroups(ctx, input)
			},
		},
		"update-smb-security-strategy": {
			Name:   "update-smb-security-strategy",
			Fields: fields_update_smb_security_strategy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSMBSecurityStrategyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_smb_security_strategy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSMBSecurityStrategy(ctx, input)
			},
		},
		"update-snapshot-schedule": {
			Name:   "update-snapshot-schedule",
			Fields: fields_update_snapshot_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSnapshotScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_snapshot_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSnapshotSchedule(ctx, input)
			},
		},
		"update-vtl-device-type": {
			Name:   "update-vtl-device-type",
			Fields: fields_update_vtl_device_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVTLDeviceTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vtl_device_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVTLDeviceType(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("storagegateway", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
