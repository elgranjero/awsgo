package fsx

// UpdateFileSystem is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
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
