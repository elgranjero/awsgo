package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/lightsail"
)

var fields_allocate_static_ip = []leanruntime.Field{
	{Name: "StaticIpName", Flag: "static-ip-name", Type: "*string", Required: true},
}

var fields_attach_certificate_to_distribution = []leanruntime.Field{
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: true},
	{Name: "DistributionName", Flag: "distribution-name", Type: "*string", Required: true},
}

var fields_attach_disk = []leanruntime.Field{
	{Name: "AutoMounting", Flag: "auto-mounting", Type: "*bool", Required: false},
	{Name: "DiskName", Flag: "disk-name", Type: "*string", Required: true},
	{Name: "DiskPath", Flag: "disk-path", Type: "*string", Required: true},
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_attach_instances_to_load_balancer = []leanruntime.Field{
	{Name: "InstanceNames", Flag: "instance-names", Type: "[]string", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_attach_load_balancer_tls_certificate = []leanruntime.Field{
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_attach_static_ip = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
	{Name: "StaticIpName", Flag: "static-ip-name", Type: "*string", Required: true},
}

var fields_close_instance_public_ports = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
	{Name: "PortInfo", Flag: "port-info", Type: "*types.PortInfo", Required: true},
}

var fields_copy_snapshot = []leanruntime.Field{
	{Name: "RestoreDate", Flag: "restore-date", Type: "*string", Required: false},
	{Name: "SourceRegion", Flag: "source-region", Type: "types.RegionName", Required: true},
	{Name: "SourceResourceName", Flag: "source-resource-name", Type: "*string", Required: false},
	{Name: "SourceSnapshotName", Flag: "source-snapshot-name", Type: "*string", Required: false},
	{Name: "TargetSnapshotName", Flag: "target-snapshot-name", Type: "*string", Required: true},
	{Name: "UseLatestRestorableAutoSnapshot", Flag: "use-latest-restorable-auto-snapshot", Type: "*bool", Required: false},
}

var fields_create_bucket = []leanruntime.Field{
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: true},
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: true},
	{Name: "EnableObjectVersioning", Flag: "enable-object-versioning", Type: "*bool", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_bucket_access_key = []leanruntime.Field{
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: true},
}

var fields_create_certificate = []leanruntime.Field{
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "SubjectAlternativeNames", Flag: "subject-alternative-names", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_cloud_formation_stack = []leanruntime.Field{
	{Name: "Instances", Flag: "instances", Type: "[]types.InstanceEntry", Required: true},
}

var fields_create_contact_method = []leanruntime.Field{
	{Name: "ContactEndpoint", Flag: "contact-endpoint", Type: "*string", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "types.ContactProtocol", Required: true},
}

var fields_create_container_service = []leanruntime.Field{
	{Name: "Deployment", Flag: "deployment", Type: "*types.ContainerServiceDeploymentRequest", Required: false},
	{Name: "Power", Flag: "power", Type: "types.ContainerServicePowerName", Required: true},
	{Name: "PrivateRegistryAccess", Flag: "private-registry-access", Type: "*types.PrivateRegistryAccessRequest", Required: false},
	{Name: "PublicDomainNames", Flag: "public-domain-names", Type: "map[string][]string", Required: false},
	{Name: "Scale", Flag: "scale", Type: "*int32", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_container_service_deployment = []leanruntime.Field{
	{Name: "Containers", Flag: "containers", Type: "map[string]types.Container", Required: false},
	{Name: "PublicEndpoint", Flag: "public-endpoint", Type: "*types.EndpointRequest", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_create_container_service_registry_login = []leanruntime.Field{}

var fields_create_disk = []leanruntime.Field{
	{Name: "AddOns", Flag: "add-ons", Type: "[]types.AddOnRequest", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: true},
	{Name: "DiskName", Flag: "disk-name", Type: "*string", Required: true},
	{Name: "SizeInGb", Flag: "size-in-gb", Type: "*int32", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_disk_from_snapshot = []leanruntime.Field{
	{Name: "AddOns", Flag: "add-ons", Type: "[]types.AddOnRequest", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: true},
	{Name: "DiskName", Flag: "disk-name", Type: "*string", Required: true},
	{Name: "DiskSnapshotName", Flag: "disk-snapshot-name", Type: "*string", Required: false},
	{Name: "RestoreDate", Flag: "restore-date", Type: "*string", Required: false},
	{Name: "SizeInGb", Flag: "size-in-gb", Type: "*int32", Required: true},
	{Name: "SourceDiskName", Flag: "source-disk-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UseLatestRestorableAutoSnapshot", Flag: "use-latest-restorable-auto-snapshot", Type: "*bool", Required: false},
}

var fields_create_disk_snapshot = []leanruntime.Field{
	{Name: "DiskName", Flag: "disk-name", Type: "*string", Required: false},
	{Name: "DiskSnapshotName", Flag: "disk-snapshot-name", Type: "*string", Required: true},
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_distribution = []leanruntime.Field{
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: true},
	{Name: "CacheBehaviorSettings", Flag: "cache-behavior-settings", Type: "*types.CacheSettings", Required: false},
	{Name: "CacheBehaviors", Flag: "cache-behaviors", Type: "[]types.CacheBehaviorPerPath", Required: false},
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: false},
	{Name: "DefaultCacheBehavior", Flag: "default-cache-behavior", Type: "*types.CacheBehavior", Required: true},
	{Name: "DistributionName", Flag: "distribution-name", Type: "*string", Required: true},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "Origin", Flag: "origin", Type: "*types.InputOrigin", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "ViewerMinimumTlsProtocolVersion", Flag: "viewer-minimum-tls-protocol-version", Type: "types.ViewerMinimumTlsProtocolVersionEnum", Required: false},
}

var fields_create_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_domain_entry = []leanruntime.Field{
	{Name: "DomainEntry", Flag: "domain-entry", Type: "*types.DomainEntry", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_create_gui_session_access_details = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_create_instance_snapshot = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
	{Name: "InstanceSnapshotName", Flag: "instance-snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_instances = []leanruntime.Field{
	{Name: "AddOns", Flag: "add-ons", Type: "[]types.AddOnRequest", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: true},
	{Name: "BlueprintId", Flag: "blueprint-id", Type: "*string", Required: true},
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: true},
	{Name: "CustomImageName", Flag: "custom-image-name", Type: "*string", Required: false},
	{Name: "InstanceNames", Flag: "instance-names", Type: "[]string", Required: true},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "KeyPairName", Flag: "key-pair-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UserData", Flag: "user-data", Type: "*string", Required: false},
}

var fields_create_instances_from_snapshot = []leanruntime.Field{
	{Name: "AddOns", Flag: "add-ons", Type: "[]types.AddOnRequest", Required: false},
	{Name: "AttachedDiskMapping", Flag: "attached-disk-mapping", Type: "map[string][]types.DiskMap", Required: false},
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: true},
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: true},
	{Name: "InstanceNames", Flag: "instance-names", Type: "[]string", Required: true},
	{Name: "InstanceSnapshotName", Flag: "instance-snapshot-name", Type: "*string", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "KeyPairName", Flag: "key-pair-name", Type: "*string", Required: false},
	{Name: "RestoreDate", Flag: "restore-date", Type: "*string", Required: false},
	{Name: "SourceInstanceName", Flag: "source-instance-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UseLatestRestorableAutoSnapshot", Flag: "use-latest-restorable-auto-snapshot", Type: "*bool", Required: false},
	{Name: "UserData", Flag: "user-data", Type: "*string", Required: false},
}

var fields_create_key_pair = []leanruntime.Field{
	{Name: "KeyPairName", Flag: "key-pair-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_load_balancer = []leanruntime.Field{
	{Name: "CertificateAlternativeNames", Flag: "certificate-alternative-names", Type: "[]string", Required: false},
	{Name: "CertificateDomainName", Flag: "certificate-domain-name", Type: "*string", Required: false},
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: false},
	{Name: "HealthCheckPath", Flag: "health-check-path", Type: "*string", Required: false},
	{Name: "InstancePort", Flag: "instance-port", Type: "int32", Required: true},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: false},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TlsPolicyName", Flag: "tls-policy-name", Type: "*string", Required: false},
}

var fields_create_load_balancer_tls_certificate = []leanruntime.Field{
	{Name: "CertificateAlternativeNames", Flag: "certificate-alternative-names", Type: "[]string", Required: false},
	{Name: "CertificateDomainName", Flag: "certificate-domain-name", Type: "*string", Required: true},
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_relational_database = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "MasterDatabaseName", Flag: "master-database-name", Type: "*string", Required: true},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "MasterUsername", Flag: "master-username", Type: "*string", Required: true},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "RelationalDatabaseBlueprintId", Flag: "relational-database-blueprint-id", Type: "*string", Required: true},
	{Name: "RelationalDatabaseBundleId", Flag: "relational-database-bundle-id", Type: "*string", Required: true},
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_relational_database_from_snapshot = []leanruntime.Field{
	{Name: "AvailabilityZone", Flag: "availability-zone", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "RelationalDatabaseBundleId", Flag: "relational-database-bundle-id", Type: "*string", Required: false},
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
	{Name: "RelationalDatabaseSnapshotName", Flag: "relational-database-snapshot-name", Type: "*string", Required: false},
	{Name: "RestoreTime", Flag: "restore-time", Type: "*time.Time", Required: false},
	{Name: "SourceRelationalDatabaseName", Flag: "source-relational-database-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UseLatestRestorableTime", Flag: "use-latest-restorable-time", Type: "*bool", Required: false},
}

var fields_create_relational_database_snapshot = []leanruntime.Field{
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
	{Name: "RelationalDatabaseSnapshotName", Flag: "relational-database-snapshot-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_alarm = []leanruntime.Field{
	{Name: "AlarmName", Flag: "alarm-name", Type: "*string", Required: true},
}

var fields_delete_auto_snapshot = []leanruntime.Field{
	{Name: "Date", Flag: "date", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_delete_bucket = []leanruntime.Field{
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: true},
	{Name: "ForceDelete", Flag: "force-delete", Type: "*bool", Required: false},
}

var fields_delete_bucket_access_key = []leanruntime.Field{
	{Name: "AccessKeyId", Flag: "access-key-id", Type: "*string", Required: true},
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: true},
}

var fields_delete_certificate = []leanruntime.Field{
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: true},
}

var fields_delete_contact_method = []leanruntime.Field{
	{Name: "Protocol", Flag: "protocol", Type: "types.ContactProtocol", Required: true},
}

var fields_delete_container_image = []leanruntime.Field{
	{Name: "Image", Flag: "image", Type: "*string", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_delete_container_service = []leanruntime.Field{
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_delete_disk = []leanruntime.Field{
	{Name: "DiskName", Flag: "disk-name", Type: "*string", Required: true},
	{Name: "ForceDeleteAddOns", Flag: "force-delete-add-ons", Type: "*bool", Required: false},
}

var fields_delete_disk_snapshot = []leanruntime.Field{
	{Name: "DiskSnapshotName", Flag: "disk-snapshot-name", Type: "*string", Required: true},
}

var fields_delete_distribution = []leanruntime.Field{
	{Name: "DistributionName", Flag: "distribution-name", Type: "*string", Required: false},
}

var fields_delete_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_domain_entry = []leanruntime.Field{
	{Name: "DomainEntry", Flag: "domain-entry", Type: "*types.DomainEntry", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_instance = []leanruntime.Field{
	{Name: "ForceDeleteAddOns", Flag: "force-delete-add-ons", Type: "*bool", Required: false},
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_delete_instance_snapshot = []leanruntime.Field{
	{Name: "InstanceSnapshotName", Flag: "instance-snapshot-name", Type: "*string", Required: true},
}

var fields_delete_key_pair = []leanruntime.Field{
	{Name: "ExpectedFingerprint", Flag: "expected-fingerprint", Type: "*string", Required: false},
	{Name: "KeyPairName", Flag: "key-pair-name", Type: "*string", Required: true},
}

var fields_delete_known_host_keys = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_delete_load_balancer = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_delete_load_balancer_tls_certificate = []leanruntime.Field{
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_delete_relational_database = []leanruntime.Field{
	{Name: "FinalRelationalDatabaseSnapshotName", Flag: "final-relational-database-snapshot-name", Type: "*string", Required: false},
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
	{Name: "SkipFinalSnapshot", Flag: "skip-final-snapshot", Type: "*bool", Required: false},
}

var fields_delete_relational_database_snapshot = []leanruntime.Field{
	{Name: "RelationalDatabaseSnapshotName", Flag: "relational-database-snapshot-name", Type: "*string", Required: true},
}

var fields_detach_certificate_from_distribution = []leanruntime.Field{
	{Name: "DistributionName", Flag: "distribution-name", Type: "*string", Required: true},
}

var fields_detach_disk = []leanruntime.Field{
	{Name: "DiskName", Flag: "disk-name", Type: "*string", Required: true},
}

var fields_detach_instances_from_load_balancer = []leanruntime.Field{
	{Name: "InstanceNames", Flag: "instance-names", Type: "[]string", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_detach_static_ip = []leanruntime.Field{
	{Name: "StaticIpName", Flag: "static-ip-name", Type: "*string", Required: true},
}

var fields_disable_add_on = []leanruntime.Field{
	{Name: "AddOnType", Flag: "add-on-type", Type: "types.AddOnType", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_download_default_key_pair = []leanruntime.Field{}

var fields_enable_add_on = []leanruntime.Field{
	{Name: "AddOnRequest", Flag: "add-on-request", Type: "*types.AddOnRequest", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_export_snapshot = []leanruntime.Field{
	{Name: "SourceSnapshotName", Flag: "source-snapshot-name", Type: "*string", Required: true},
}

var fields_get_active_names = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_alarms = []leanruntime.Field{
	{Name: "AlarmName", Flag: "alarm-name", Type: "*string", Required: false},
	{Name: "MonitoredResourceName", Flag: "monitored-resource-name", Type: "*string", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_auto_snapshots = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_get_blueprints = []leanruntime.Field{
	{Name: "AppCategory", Flag: "app-category", Type: "types.AppCategory", Required: false},
	{Name: "IncludeInactive", Flag: "include-inactive", Type: "*bool", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_bucket_access_keys = []leanruntime.Field{
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: true},
}

var fields_get_bucket_bundles = []leanruntime.Field{
	{Name: "IncludeInactive", Flag: "include-inactive", Type: "*bool", Required: false},
}

var fields_get_bucket_metric_data = []leanruntime.Field{
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "types.BucketMetricName", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Statistics", Flag: "statistics", Type: "[]types.MetricStatistic", Required: true},
	{Name: "Unit", Flag: "unit", Type: "types.MetricUnit", Required: true},
}

var fields_get_buckets = []leanruntime.Field{
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: false},
	{Name: "IncludeConnectedResources", Flag: "include-connected-resources", Type: "*bool", Required: false},
	{Name: "IncludeCors", Flag: "include-cors", Type: "*bool", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_bundles = []leanruntime.Field{
	{Name: "AppCategory", Flag: "app-category", Type: "types.AppCategory", Required: false},
	{Name: "IncludeInactive", Flag: "include-inactive", Type: "*bool", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_certificates = []leanruntime.Field{
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: false},
	{Name: "CertificateStatuses", Flag: "certificate-statuses", Type: "[]types.CertificateStatus", Required: false},
	{Name: "IncludeCertificateDetails", Flag: "include-certificate-details", Type: "bool", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_cloud_formation_stack_records = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_contact_methods = []leanruntime.Field{
	{Name: "Protocols", Flag: "protocols", Type: "[]types.ContactProtocol", Required: false},
}

var fields_get_container_api_metadata = []leanruntime.Field{}

var fields_get_container_images = []leanruntime.Field{
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_get_container_log = []leanruntime.Field{
	{Name: "ContainerName", Flag: "container-name", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "FilterPattern", Flag: "filter-pattern", Type: "*string", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_get_container_service_deployments = []leanruntime.Field{
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_get_container_service_metric_data = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "types.ContainerServiceMetricName", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Statistics", Flag: "statistics", Type: "[]types.MetricStatistic", Required: true},
}

var fields_get_container_service_powers = []leanruntime.Field{}

var fields_get_container_services = []leanruntime.Field{
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: false},
}

var fields_get_cost_estimate = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
}

var fields_get_disk = []leanruntime.Field{
	{Name: "DiskName", Flag: "disk-name", Type: "*string", Required: true},
}

var fields_get_disk_snapshot = []leanruntime.Field{
	{Name: "DiskSnapshotName", Flag: "disk-snapshot-name", Type: "*string", Required: true},
}

var fields_get_disk_snapshots = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_disks = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_distribution_bundles = []leanruntime.Field{}

var fields_get_distribution_latest_cache_reset = []leanruntime.Field{
	{Name: "DistributionName", Flag: "distribution-name", Type: "*string", Required: false},
}

var fields_get_distribution_metric_data = []leanruntime.Field{
	{Name: "DistributionName", Flag: "distribution-name", Type: "*string", Required: true},
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "types.DistributionMetricName", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Statistics", Flag: "statistics", Type: "[]types.MetricStatistic", Required: true},
	{Name: "Unit", Flag: "unit", Type: "types.MetricUnit", Required: true},
}

var fields_get_distributions = []leanruntime.Field{
	{Name: "DistributionName", Flag: "distribution-name", Type: "*string", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_get_domains = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_export_snapshot_records = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_instance = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_get_instance_access_details = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
	{Name: "Protocol", Flag: "protocol", Type: "types.InstanceAccessProtocol", Required: false},
}

var fields_get_instance_metric_data = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "types.InstanceMetricName", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Statistics", Flag: "statistics", Type: "[]types.MetricStatistic", Required: true},
	{Name: "Unit", Flag: "unit", Type: "types.MetricUnit", Required: true},
}

var fields_get_instance_port_states = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_get_instance_snapshot = []leanruntime.Field{
	{Name: "InstanceSnapshotName", Flag: "instance-snapshot-name", Type: "*string", Required: true},
}

var fields_get_instance_snapshots = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_instance_state = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_get_instances = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_key_pair = []leanruntime.Field{
	{Name: "KeyPairName", Flag: "key-pair-name", Type: "*string", Required: true},
}

var fields_get_key_pairs = []leanruntime.Field{
	{Name: "IncludeDefaultKeyPair", Flag: "include-default-key-pair", Type: "*bool", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_load_balancer = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_get_load_balancer_metric_data = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "types.LoadBalancerMetricName", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Statistics", Flag: "statistics", Type: "[]types.MetricStatistic", Required: true},
	{Name: "Unit", Flag: "unit", Type: "types.MetricUnit", Required: true},
}

var fields_get_load_balancer_tls_certificates = []leanruntime.Field{
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_get_load_balancer_tls_policies = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_load_balancers = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_operation = []leanruntime.Field{
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
}

var fields_get_operations = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_operations_for_resource = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_get_regions = []leanruntime.Field{
	{Name: "IncludeAvailabilityZones", Flag: "include-availability-zones", Type: "*bool", Required: false},
	{Name: "IncludeRelationalDatabaseAvailabilityZones", Flag: "include-relational-database-availability-zones", Type: "*bool", Required: false},
}

var fields_get_relational_database = []leanruntime.Field{
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
}

var fields_get_relational_database_blueprints = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_relational_database_bundles = []leanruntime.Field{
	{Name: "IncludeInactive", Flag: "include-inactive", Type: "*bool", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_relational_database_events = []leanruntime.Field{
	{Name: "DurationInMinutes", Flag: "duration-in-minutes", Type: "*int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
}

var fields_get_relational_database_log_events = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: false},
	{Name: "LogStreamName", Flag: "log-stream-name", Type: "*string", Required: true},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
	{Name: "StartFromHead", Flag: "start-from-head", Type: "*bool", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
}

var fields_get_relational_database_log_streams = []leanruntime.Field{
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
}

var fields_get_relational_database_master_user_password = []leanruntime.Field{
	{Name: "PasswordVersion", Flag: "password-version", Type: "types.RelationalDatabasePasswordVersion", Required: false},
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
}

var fields_get_relational_database_metric_data = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "types.RelationalDatabaseMetricName", Required: true},
	{Name: "Period", Flag: "period", Type: "*int32", Required: true},
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Statistics", Flag: "statistics", Type: "[]types.MetricStatistic", Required: true},
	{Name: "Unit", Flag: "unit", Type: "types.MetricUnit", Required: true},
}

var fields_get_relational_database_parameters = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
}

var fields_get_relational_database_snapshot = []leanruntime.Field{
	{Name: "RelationalDatabaseSnapshotName", Flag: "relational-database-snapshot-name", Type: "*string", Required: true},
}

var fields_get_relational_database_snapshots = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_relational_databases = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_get_setup_history = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_get_static_ip = []leanruntime.Field{
	{Name: "StaticIpName", Flag: "static-ip-name", Type: "*string", Required: true},
}

var fields_get_static_ips = []leanruntime.Field{
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_import_key_pair = []leanruntime.Field{
	{Name: "KeyPairName", Flag: "key-pair-name", Type: "*string", Required: true},
	{Name: "PublicKeyBase64", Flag: "public-key-base64", Type: "*string", Required: true},
}

var fields_is_vpc_peered = []leanruntime.Field{}

var fields_open_instance_public_ports = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
	{Name: "PortInfo", Flag: "port-info", Type: "*types.PortInfo", Required: true},
}

var fields_peer_vpc = []leanruntime.Field{}

var fields_put_alarm = []leanruntime.Field{
	{Name: "AlarmName", Flag: "alarm-name", Type: "*string", Required: true},
	{Name: "ComparisonOperator", Flag: "comparison-operator", Type: "types.ComparisonOperator", Required: true},
	{Name: "ContactProtocols", Flag: "contact-protocols", Type: "[]types.ContactProtocol", Required: false},
	{Name: "DatapointsToAlarm", Flag: "datapoints-to-alarm", Type: "*int32", Required: false},
	{Name: "EvaluationPeriods", Flag: "evaluation-periods", Type: "*int32", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "types.MetricName", Required: true},
	{Name: "MonitoredResourceName", Flag: "monitored-resource-name", Type: "*string", Required: true},
	{Name: "NotificationEnabled", Flag: "notification-enabled", Type: "*bool", Required: false},
	{Name: "NotificationTriggers", Flag: "notification-triggers", Type: "[]types.AlarmState", Required: false},
	{Name: "Threshold", Flag: "threshold", Type: "*float64", Required: true},
	{Name: "TreatMissingData", Flag: "treat-missing-data", Type: "types.TreatMissingData", Required: false},
}

var fields_put_instance_public_ports = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
	{Name: "PortInfos", Flag: "port-infos", Type: "[]types.PortInfo", Required: true},
}

var fields_reboot_instance = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_reboot_relational_database = []leanruntime.Field{
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
}

var fields_register_container_image = []leanruntime.Field{
	{Name: "Digest", Flag: "digest", Type: "*string", Required: true},
	{Name: "Label", Flag: "label", Type: "*string", Required: true},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_release_static_ip = []leanruntime.Field{
	{Name: "StaticIpName", Flag: "static-ip-name", Type: "*string", Required: true},
}

var fields_reset_distribution_cache = []leanruntime.Field{
	{Name: "DistributionName", Flag: "distribution-name", Type: "*string", Required: false},
}

var fields_send_contact_method_verification = []leanruntime.Field{
	{Name: "Protocol", Flag: "protocol", Type: "types.ContactMethodVerificationProtocol", Required: true},
}

var fields_set_ip_address_type = []leanruntime.Field{
	{Name: "AcceptBundleUpdate", Flag: "accept-bundle-update", Type: "*bool", Required: false},
	{Name: "IpAddressType", Flag: "ip-address-type", Type: "types.IpAddressType", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
}

var fields_set_resource_access_for_bucket = []leanruntime.Field{
	{Name: "Access", Flag: "access", Type: "types.ResourceBucketAccess", Required: true},
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: true},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_setup_instance_https = []leanruntime.Field{
	{Name: "CertificateProvider", Flag: "certificate-provider", Type: "types.CertificateProvider", Required: true},
	{Name: "DomainNames", Flag: "domain-names", Type: "[]string", Required: true},
	{Name: "EmailAddress", Flag: "email-address", Type: "*string", Required: true},
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_start_gui_session = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_start_instance = []leanruntime.Field{
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_start_relational_database = []leanruntime.Field{
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
}

var fields_stop_gui_session = []leanruntime.Field{
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
}

var fields_stop_instance = []leanruntime.Field{
	{Name: "Force", Flag: "force", Type: "*bool", Required: false},
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_stop_relational_database = []leanruntime.Field{
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
	{Name: "RelationalDatabaseSnapshotName", Flag: "relational-database-snapshot-name", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_test_alarm = []leanruntime.Field{
	{Name: "AlarmName", Flag: "alarm-name", Type: "*string", Required: true},
	{Name: "State", Flag: "state", Type: "types.AlarmState", Required: true},
}

var fields_unpeer_vpc = []leanruntime.Field{}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: false},
	{Name: "ResourceName", Flag: "resource-name", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_bucket = []leanruntime.Field{
	{Name: "AccessLogConfig", Flag: "access-log-config", Type: "*types.BucketAccessLogConfig", Required: false},
	{Name: "AccessRules", Flag: "access-rules", Type: "*types.AccessRules", Required: false},
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: true},
	{Name: "Cors", Flag: "cors", Type: "*types.BucketCorsConfig", Required: false},
	{Name: "ReadonlyAccessAccounts", Flag: "readonly-access-accounts", Type: "[]string", Required: false},
	{Name: "Versioning", Flag: "versioning", Type: "*string", Required: false},
}

var fields_update_bucket_bundle = []leanruntime.Field{
	{Name: "BucketName", Flag: "bucket-name", Type: "*string", Required: true},
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: true},
}

var fields_update_container_service = []leanruntime.Field{
	{Name: "IsDisabled", Flag: "is-disabled", Type: "*bool", Required: false},
	{Name: "Power", Flag: "power", Type: "types.ContainerServicePowerName", Required: false},
	{Name: "PrivateRegistryAccess", Flag: "private-registry-access", Type: "*types.PrivateRegistryAccessRequest", Required: false},
	{Name: "PublicDomainNames", Flag: "public-domain-names", Type: "map[string][]string", Required: false},
	{Name: "Scale", Flag: "scale", Type: "*int32", Required: false},
	{Name: "ServiceName", Flag: "service-name", Type: "*string", Required: true},
}

var fields_update_distribution = []leanruntime.Field{
	{Name: "CacheBehaviorSettings", Flag: "cache-behavior-settings", Type: "*types.CacheSettings", Required: false},
	{Name: "CacheBehaviors", Flag: "cache-behaviors", Type: "[]types.CacheBehaviorPerPath", Required: false},
	{Name: "CertificateName", Flag: "certificate-name", Type: "*string", Required: false},
	{Name: "DefaultCacheBehavior", Flag: "default-cache-behavior", Type: "*types.CacheBehavior", Required: false},
	{Name: "DistributionName", Flag: "distribution-name", Type: "*string", Required: true},
	{Name: "IsEnabled", Flag: "is-enabled", Type: "*bool", Required: false},
	{Name: "Origin", Flag: "origin", Type: "*types.InputOrigin", Required: false},
	{Name: "UseDefaultCertificate", Flag: "use-default-certificate", Type: "*bool", Required: false},
	{Name: "ViewerMinimumTlsProtocolVersion", Flag: "viewer-minimum-tls-protocol-version", Type: "types.ViewerMinimumTlsProtocolVersionEnum", Required: false},
}

var fields_update_distribution_bundle = []leanruntime.Field{
	{Name: "BundleId", Flag: "bundle-id", Type: "*string", Required: false},
	{Name: "DistributionName", Flag: "distribution-name", Type: "*string", Required: false},
}

var fields_update_domain_entry = []leanruntime.Field{
	{Name: "DomainEntry", Flag: "domain-entry", Type: "*types.DomainEntry", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_update_instance_metadata_options = []leanruntime.Field{
	{Name: "HttpEndpoint", Flag: "http-endpoint", Type: "types.HttpEndpoint", Required: false},
	{Name: "HttpProtocolIpv6", Flag: "http-protocol-ipv6", Type: "types.HttpProtocolIpv6", Required: false},
	{Name: "HttpPutResponseHopLimit", Flag: "http-put-response-hop-limit", Type: "*int32", Required: false},
	{Name: "HttpTokens", Flag: "http-tokens", Type: "types.HttpTokens", Required: false},
	{Name: "InstanceName", Flag: "instance-name", Type: "*string", Required: true},
}

var fields_update_load_balancer_attribute = []leanruntime.Field{
	{Name: "AttributeName", Flag: "attribute-name", Type: "types.LoadBalancerAttributeName", Required: true},
	{Name: "AttributeValue", Flag: "attribute-value", Type: "*string", Required: true},
	{Name: "LoadBalancerName", Flag: "load-balancer-name", Type: "*string", Required: true},
}

var fields_update_relational_database = []leanruntime.Field{
	{Name: "ApplyImmediately", Flag: "apply-immediately", Type: "*bool", Required: false},
	{Name: "CaCertificateIdentifier", Flag: "ca-certificate-identifier", Type: "*string", Required: false},
	{Name: "DisableBackupRetention", Flag: "disable-backup-retention", Type: "*bool", Required: false},
	{Name: "EnableBackupRetention", Flag: "enable-backup-retention", Type: "*bool", Required: false},
	{Name: "MasterUserPassword", Flag: "master-user-password", Type: "*string", Required: false},
	{Name: "PreferredBackupWindow", Flag: "preferred-backup-window", Type: "*string", Required: false},
	{Name: "PreferredMaintenanceWindow", Flag: "preferred-maintenance-window", Type: "*string", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: false},
	{Name: "RelationalDatabaseBlueprintId", Flag: "relational-database-blueprint-id", Type: "*string", Required: false},
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
	{Name: "RotateMasterUserPassword", Flag: "rotate-master-user-password", Type: "*bool", Required: false},
}

var fields_update_relational_database_parameters = []leanruntime.Field{
	{Name: "Parameters", Flag: "parameters", Type: "[]types.RelationalDatabaseParameter", Required: true},
	{Name: "RelationalDatabaseName", Flag: "relational-database-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"allocate-static-ip": {
			Name:   "allocate-static-ip",
			Fields: fields_allocate_static_ip,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AllocateStaticIpInput{}
				if _, err := leanruntime.ApplyInput(input, fields_allocate_static_ip, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AllocateStaticIp(ctx, input)
			},
		},
		"attach-certificate-to-distribution": {
			Name:   "attach-certificate-to-distribution",
			Fields: fields_attach_certificate_to_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachCertificateToDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_certificate_to_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachCertificateToDistribution(ctx, input)
			},
		},
		"attach-disk": {
			Name:   "attach-disk",
			Fields: fields_attach_disk,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachDiskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_disk, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachDisk(ctx, input)
			},
		},
		"attach-instances-to-load-balancer": {
			Name:   "attach-instances-to-load-balancer",
			Fields: fields_attach_instances_to_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachInstancesToLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_instances_to_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachInstancesToLoadBalancer(ctx, input)
			},
		},
		"attach-load-balancer-tls-certificate": {
			Name:   "attach-load-balancer-tls-certificate",
			Fields: fields_attach_load_balancer_tls_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachLoadBalancerTlsCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_load_balancer_tls_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachLoadBalancerTlsCertificate(ctx, input)
			},
		},
		"attach-static-ip": {
			Name:   "attach-static-ip",
			Fields: fields_attach_static_ip,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachStaticIpInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_static_ip, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachStaticIp(ctx, input)
			},
		},
		"close-instance-public-ports": {
			Name:   "close-instance-public-ports",
			Fields: fields_close_instance_public_ports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CloseInstancePublicPortsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_close_instance_public_ports, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CloseInstancePublicPorts(ctx, input)
			},
		},
		"copy-snapshot": {
			Name:   "copy-snapshot",
			Fields: fields_copy_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopySnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopySnapshot(ctx, input)
			},
		},
		"create-bucket": {
			Name:   "create-bucket",
			Fields: fields_create_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBucket(ctx, input)
			},
		},
		"create-bucket-access-key": {
			Name:   "create-bucket-access-key",
			Fields: fields_create_bucket_access_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBucketAccessKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bucket_access_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBucketAccessKey(ctx, input)
			},
		},
		"create-certificate": {
			Name:   "create-certificate",
			Fields: fields_create_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCertificate(ctx, input)
			},
		},
		"create-cloud-formation-stack": {
			Name:   "create-cloud-formation-stack",
			Fields: fields_create_cloud_formation_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCloudFormationStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cloud_formation_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCloudFormationStack(ctx, input)
			},
		},
		"create-contact-method": {
			Name:   "create-contact-method",
			Fields: fields_create_contact_method,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContactMethodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_contact_method, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContactMethod(ctx, input)
			},
		},
		"create-container-service": {
			Name:   "create-container-service",
			Fields: fields_create_container_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContainerServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_container_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContainerService(ctx, input)
			},
		},
		"create-container-service-deployment": {
			Name:   "create-container-service-deployment",
			Fields: fields_create_container_service_deployment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContainerServiceDeploymentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_container_service_deployment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContainerServiceDeployment(ctx, input)
			},
		},
		"create-container-service-registry-login": {
			Name:   "create-container-service-registry-login",
			Fields: fields_create_container_service_registry_login,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateContainerServiceRegistryLoginInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_container_service_registry_login, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateContainerServiceRegistryLogin(ctx, input)
			},
		},
		"create-disk": {
			Name:   "create-disk",
			Fields: fields_create_disk,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDiskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_disk, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDisk(ctx, input)
			},
		},
		"create-disk-from-snapshot": {
			Name:   "create-disk-from-snapshot",
			Fields: fields_create_disk_from_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDiskFromSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_disk_from_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDiskFromSnapshot(ctx, input)
			},
		},
		"create-disk-snapshot": {
			Name:   "create-disk-snapshot",
			Fields: fields_create_disk_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDiskSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_disk_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDiskSnapshot(ctx, input)
			},
		},
		"create-distribution": {
			Name:   "create-distribution",
			Fields: fields_create_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDistribution(ctx, input)
			},
		},
		"create-domain": {
			Name:   "create-domain",
			Fields: fields_create_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomain(ctx, input)
			},
		},
		"create-domain-entry": {
			Name:   "create-domain-entry",
			Fields: fields_create_domain_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomainEntry(ctx, input)
			},
		},
		"create-gui-session-access-details": {
			Name:   "create-gui-session-access-details",
			Fields: fields_create_gui_session_access_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGUISessionAccessDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_gui_session_access_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGUISessionAccessDetails(ctx, input)
			},
		},
		"create-instance-snapshot": {
			Name:   "create-instance-snapshot",
			Fields: fields_create_instance_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstanceSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instance_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstanceSnapshot(ctx, input)
			},
		},
		"create-instances": {
			Name:   "create-instances",
			Fields: fields_create_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstances(ctx, input)
			},
		},
		"create-instances-from-snapshot": {
			Name:   "create-instances-from-snapshot",
			Fields: fields_create_instances_from_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateInstancesFromSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_instances_from_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateInstancesFromSnapshot(ctx, input)
			},
		},
		"create-key-pair": {
			Name:   "create-key-pair",
			Fields: fields_create_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKeyPair(ctx, input)
			},
		},
		"create-load-balancer": {
			Name:   "create-load-balancer",
			Fields: fields_create_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoadBalancer(ctx, input)
			},
		},
		"create-load-balancer-tls-certificate": {
			Name:   "create-load-balancer-tls-certificate",
			Fields: fields_create_load_balancer_tls_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLoadBalancerTlsCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_load_balancer_tls_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLoadBalancerTlsCertificate(ctx, input)
			},
		},
		"create-relational-database": {
			Name:   "create-relational-database",
			Fields: fields_create_relational_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRelationalDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_relational_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRelationalDatabase(ctx, input)
			},
		},
		"create-relational-database-from-snapshot": {
			Name:   "create-relational-database-from-snapshot",
			Fields: fields_create_relational_database_from_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRelationalDatabaseFromSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_relational_database_from_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRelationalDatabaseFromSnapshot(ctx, input)
			},
		},
		"create-relational-database-snapshot": {
			Name:   "create-relational-database-snapshot",
			Fields: fields_create_relational_database_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRelationalDatabaseSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_relational_database_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRelationalDatabaseSnapshot(ctx, input)
			},
		},
		"delete-alarm": {
			Name:   "delete-alarm",
			Fields: fields_delete_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAlarm(ctx, input)
			},
		},
		"delete-auto-snapshot": {
			Name:   "delete-auto-snapshot",
			Fields: fields_delete_auto_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAutoSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_auto_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAutoSnapshot(ctx, input)
			},
		},
		"delete-bucket": {
			Name:   "delete-bucket",
			Fields: fields_delete_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucket(ctx, input)
			},
		},
		"delete-bucket-access-key": {
			Name:   "delete-bucket-access-key",
			Fields: fields_delete_bucket_access_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBucketAccessKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bucket_access_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBucketAccessKey(ctx, input)
			},
		},
		"delete-certificate": {
			Name:   "delete-certificate",
			Fields: fields_delete_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCertificate(ctx, input)
			},
		},
		"delete-contact-method": {
			Name:   "delete-contact-method",
			Fields: fields_delete_contact_method,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContactMethodInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_contact_method, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContactMethod(ctx, input)
			},
		},
		"delete-container-image": {
			Name:   "delete-container-image",
			Fields: fields_delete_container_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContainerImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_container_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContainerImage(ctx, input)
			},
		},
		"delete-container-service": {
			Name:   "delete-container-service",
			Fields: fields_delete_container_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteContainerServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_container_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteContainerService(ctx, input)
			},
		},
		"delete-disk": {
			Name:   "delete-disk",
			Fields: fields_delete_disk,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDiskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_disk, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDisk(ctx, input)
			},
		},
		"delete-disk-snapshot": {
			Name:   "delete-disk-snapshot",
			Fields: fields_delete_disk_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDiskSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_disk_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDiskSnapshot(ctx, input)
			},
		},
		"delete-distribution": {
			Name:   "delete-distribution",
			Fields: fields_delete_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDistribution(ctx, input)
			},
		},
		"delete-domain": {
			Name:   "delete-domain",
			Fields: fields_delete_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomain(ctx, input)
			},
		},
		"delete-domain-entry": {
			Name:   "delete-domain-entry",
			Fields: fields_delete_domain_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomainEntry(ctx, input)
			},
		},
		"delete-instance": {
			Name:   "delete-instance",
			Fields: fields_delete_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstance(ctx, input)
			},
		},
		"delete-instance-snapshot": {
			Name:   "delete-instance-snapshot",
			Fields: fields_delete_instance_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstanceSnapshot(ctx, input)
			},
		},
		"delete-key-pair": {
			Name:   "delete-key-pair",
			Fields: fields_delete_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKeyPair(ctx, input)
			},
		},
		"delete-known-host-keys": {
			Name:   "delete-known-host-keys",
			Fields: fields_delete_known_host_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKnownHostKeysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_known_host_keys, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKnownHostKeys(ctx, input)
			},
		},
		"delete-load-balancer": {
			Name:   "delete-load-balancer",
			Fields: fields_delete_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoadBalancer(ctx, input)
			},
		},
		"delete-load-balancer-tls-certificate": {
			Name:   "delete-load-balancer-tls-certificate",
			Fields: fields_delete_load_balancer_tls_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoadBalancerTlsCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_load_balancer_tls_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoadBalancerTlsCertificate(ctx, input)
			},
		},
		"delete-relational-database": {
			Name:   "delete-relational-database",
			Fields: fields_delete_relational_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRelationalDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_relational_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRelationalDatabase(ctx, input)
			},
		},
		"delete-relational-database-snapshot": {
			Name:   "delete-relational-database-snapshot",
			Fields: fields_delete_relational_database_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRelationalDatabaseSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_relational_database_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRelationalDatabaseSnapshot(ctx, input)
			},
		},
		"detach-certificate-from-distribution": {
			Name:   "detach-certificate-from-distribution",
			Fields: fields_detach_certificate_from_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachCertificateFromDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_certificate_from_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachCertificateFromDistribution(ctx, input)
			},
		},
		"detach-disk": {
			Name:   "detach-disk",
			Fields: fields_detach_disk,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachDiskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_disk, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachDisk(ctx, input)
			},
		},
		"detach-instances-from-load-balancer": {
			Name:   "detach-instances-from-load-balancer",
			Fields: fields_detach_instances_from_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachInstancesFromLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_instances_from_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachInstancesFromLoadBalancer(ctx, input)
			},
		},
		"detach-static-ip": {
			Name:   "detach-static-ip",
			Fields: fields_detach_static_ip,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachStaticIpInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_static_ip, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachStaticIp(ctx, input)
			},
		},
		"disable-add-on": {
			Name:   "disable-add-on",
			Fields: fields_disable_add_on,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableAddOnInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_add_on, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableAddOn(ctx, input)
			},
		},
		"download-default-key-pair": {
			Name:   "download-default-key-pair",
			Fields: fields_download_default_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DownloadDefaultKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_download_default_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DownloadDefaultKeyPair(ctx, input)
			},
		},
		"enable-add-on": {
			Name:   "enable-add-on",
			Fields: fields_enable_add_on,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableAddOnInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_add_on, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableAddOn(ctx, input)
			},
		},
		"export-snapshot": {
			Name:   "export-snapshot",
			Fields: fields_export_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExportSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_export_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExportSnapshot(ctx, input)
			},
		},
		"get-active-names": {
			Name:   "get-active-names",
			Fields: fields_get_active_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetActiveNamesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_active_names, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetActiveNames(ctx, input)
			},
		},
		"get-alarms": {
			Name:   "get-alarms",
			Fields: fields_get_alarms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAlarmsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_alarms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAlarms(ctx, input)
			},
		},
		"get-auto-snapshots": {
			Name:   "get-auto-snapshots",
			Fields: fields_get_auto_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAutoSnapshotsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_auto_snapshots, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAutoSnapshots(ctx, input)
			},
		},
		"get-blueprints": {
			Name:   "get-blueprints",
			Fields: fields_get_blueprints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBlueprintsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_blueprints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBlueprints(ctx, input)
			},
		},
		"get-bucket-access-keys": {
			Name:   "get-bucket-access-keys",
			Fields: fields_get_bucket_access_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketAccessKeysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_access_keys, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketAccessKeys(ctx, input)
			},
		},
		"get-bucket-bundles": {
			Name:   "get-bucket-bundles",
			Fields: fields_get_bucket_bundles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketBundlesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_bundles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketBundles(ctx, input)
			},
		},
		"get-bucket-metric-data": {
			Name:   "get-bucket-metric-data",
			Fields: fields_get_bucket_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketMetricDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bucket_metric_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBucketMetricData(ctx, input)
			},
		},
		"get-buckets": {
			Name:   "get-buckets",
			Fields: fields_get_buckets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBucketsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_buckets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBuckets(ctx, input)
			},
		},
		"get-bundles": {
			Name:   "get-bundles",
			Fields: fields_get_bundles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBundlesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bundles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBundles(ctx, input)
			},
		},
		"get-certificates": {
			Name:   "get-certificates",
			Fields: fields_get_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCertificatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_certificates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCertificates(ctx, input)
			},
		},
		"get-cloud-formation-stack-records": {
			Name:   "get-cloud-formation-stack-records",
			Fields: fields_get_cloud_formation_stack_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCloudFormationStackRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cloud_formation_stack_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCloudFormationStackRecords(ctx, input)
			},
		},
		"get-contact-methods": {
			Name:   "get-contact-methods",
			Fields: fields_get_contact_methods,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContactMethodsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_contact_methods, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContactMethods(ctx, input)
			},
		},
		"get-container-api-metadata": {
			Name:   "get-container-api-metadata",
			Fields: fields_get_container_api_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContainerAPIMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_container_api_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContainerAPIMetadata(ctx, input)
			},
		},
		"get-container-images": {
			Name:   "get-container-images",
			Fields: fields_get_container_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContainerImagesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_container_images, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContainerImages(ctx, input)
			},
		},
		"get-container-log": {
			Name:   "get-container-log",
			Fields: fields_get_container_log,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContainerLogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_container_log, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContainerLog(ctx, input)
			},
		},
		"get-container-service-deployments": {
			Name:   "get-container-service-deployments",
			Fields: fields_get_container_service_deployments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContainerServiceDeploymentsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_container_service_deployments, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContainerServiceDeployments(ctx, input)
			},
		},
		"get-container-service-metric-data": {
			Name:   "get-container-service-metric-data",
			Fields: fields_get_container_service_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContainerServiceMetricDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_container_service_metric_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContainerServiceMetricData(ctx, input)
			},
		},
		"get-container-service-powers": {
			Name:   "get-container-service-powers",
			Fields: fields_get_container_service_powers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContainerServicePowersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_container_service_powers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContainerServicePowers(ctx, input)
			},
		},
		"get-container-services": {
			Name:   "get-container-services",
			Fields: fields_get_container_services,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContainerServicesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_container_services, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContainerServices(ctx, input)
			},
		},
		"get-cost-estimate": {
			Name:   "get-cost-estimate",
			Fields: fields_get_cost_estimate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCostEstimateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_cost_estimate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCostEstimate(ctx, input)
			},
		},
		"get-disk": {
			Name:   "get-disk",
			Fields: fields_get_disk,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDiskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_disk, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDisk(ctx, input)
			},
		},
		"get-disk-snapshot": {
			Name:   "get-disk-snapshot",
			Fields: fields_get_disk_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDiskSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_disk_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDiskSnapshot(ctx, input)
			},
		},
		"get-disk-snapshots": {
			Name:   "get-disk-snapshots",
			Fields: fields_get_disk_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDiskSnapshotsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_disk_snapshots, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDiskSnapshots(ctx, input)
			},
		},
		"get-disks": {
			Name:   "get-disks",
			Fields: fields_get_disks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDisksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_disks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDisks(ctx, input)
			},
		},
		"get-distribution-bundles": {
			Name:   "get-distribution-bundles",
			Fields: fields_get_distribution_bundles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDistributionBundlesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_distribution_bundles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDistributionBundles(ctx, input)
			},
		},
		"get-distribution-latest-cache-reset": {
			Name:   "get-distribution-latest-cache-reset",
			Fields: fields_get_distribution_latest_cache_reset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDistributionLatestCacheResetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_distribution_latest_cache_reset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDistributionLatestCacheReset(ctx, input)
			},
		},
		"get-distribution-metric-data": {
			Name:   "get-distribution-metric-data",
			Fields: fields_get_distribution_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDistributionMetricDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_distribution_metric_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDistributionMetricData(ctx, input)
			},
		},
		"get-distributions": {
			Name:   "get-distributions",
			Fields: fields_get_distributions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDistributionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_distributions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDistributions(ctx, input)
			},
		},
		"get-domain": {
			Name:   "get-domain",
			Fields: fields_get_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomain(ctx, input)
			},
		},
		"get-domains": {
			Name:   "get-domains",
			Fields: fields_get_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domains, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomains(ctx, input)
			},
		},
		"get-export-snapshot-records": {
			Name:   "get-export-snapshot-records",
			Fields: fields_get_export_snapshot_records,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExportSnapshotRecordsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_export_snapshot_records, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExportSnapshotRecords(ctx, input)
			},
		},
		"get-instance": {
			Name:   "get-instance",
			Fields: fields_get_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstance(ctx, input)
			},
		},
		"get-instance-access-details": {
			Name:   "get-instance-access-details",
			Fields: fields_get_instance_access_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceAccessDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_access_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceAccessDetails(ctx, input)
			},
		},
		"get-instance-metric-data": {
			Name:   "get-instance-metric-data",
			Fields: fields_get_instance_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceMetricDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_metric_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceMetricData(ctx, input)
			},
		},
		"get-instance-port-states": {
			Name:   "get-instance-port-states",
			Fields: fields_get_instance_port_states,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstancePortStatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_port_states, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstancePortStates(ctx, input)
			},
		},
		"get-instance-snapshot": {
			Name:   "get-instance-snapshot",
			Fields: fields_get_instance_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceSnapshot(ctx, input)
			},
		},
		"get-instance-snapshots": {
			Name:   "get-instance-snapshots",
			Fields: fields_get_instance_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceSnapshotsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_snapshots, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceSnapshots(ctx, input)
			},
		},
		"get-instance-state": {
			Name:   "get-instance-state",
			Fields: fields_get_instance_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceState(ctx, input)
			},
		},
		"get-instances": {
			Name:   "get-instances",
			Fields: fields_get_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstances(ctx, input)
			},
		},
		"get-key-pair": {
			Name:   "get-key-pair",
			Fields: fields_get_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKeyPair(ctx, input)
			},
		},
		"get-key-pairs": {
			Name:   "get-key-pairs",
			Fields: fields_get_key_pairs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetKeyPairsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_key_pairs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetKeyPairs(ctx, input)
			},
		},
		"get-load-balancer": {
			Name:   "get-load-balancer",
			Fields: fields_get_load_balancer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoadBalancerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_load_balancer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoadBalancer(ctx, input)
			},
		},
		"get-load-balancer-metric-data": {
			Name:   "get-load-balancer-metric-data",
			Fields: fields_get_load_balancer_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoadBalancerMetricDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_load_balancer_metric_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoadBalancerMetricData(ctx, input)
			},
		},
		"get-load-balancer-tls-certificates": {
			Name:   "get-load-balancer-tls-certificates",
			Fields: fields_get_load_balancer_tls_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoadBalancerTlsCertificatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_load_balancer_tls_certificates, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoadBalancerTlsCertificates(ctx, input)
			},
		},
		"get-load-balancer-tls-policies": {
			Name:   "get-load-balancer-tls-policies",
			Fields: fields_get_load_balancer_tls_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoadBalancerTlsPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_load_balancer_tls_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoadBalancerTlsPolicies(ctx, input)
			},
		},
		"get-load-balancers": {
			Name:   "get-load-balancers",
			Fields: fields_get_load_balancers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoadBalancersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_load_balancers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoadBalancers(ctx, input)
			},
		},
		"get-operation": {
			Name:   "get-operation",
			Fields: fields_get_operation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOperationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_operation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOperation(ctx, input)
			},
		},
		"get-operations": {
			Name:   "get-operations",
			Fields: fields_get_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOperationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_operations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOperations(ctx, input)
			},
		},
		"get-operations-for-resource": {
			Name:   "get-operations-for-resource",
			Fields: fields_get_operations_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOperationsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_operations_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOperationsForResource(ctx, input)
			},
		},
		"get-regions": {
			Name:   "get-regions",
			Fields: fields_get_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_regions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegions(ctx, input)
			},
		},
		"get-relational-database": {
			Name:   "get-relational-database",
			Fields: fields_get_relational_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabase(ctx, input)
			},
		},
		"get-relational-database-blueprints": {
			Name:   "get-relational-database-blueprints",
			Fields: fields_get_relational_database_blueprints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseBlueprintsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database_blueprints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabaseBlueprints(ctx, input)
			},
		},
		"get-relational-database-bundles": {
			Name:   "get-relational-database-bundles",
			Fields: fields_get_relational_database_bundles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseBundlesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database_bundles, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabaseBundles(ctx, input)
			},
		},
		"get-relational-database-events": {
			Name:   "get-relational-database-events",
			Fields: fields_get_relational_database_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabaseEvents(ctx, input)
			},
		},
		"get-relational-database-log-events": {
			Name:   "get-relational-database-log-events",
			Fields: fields_get_relational_database_log_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseLogEventsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database_log_events, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabaseLogEvents(ctx, input)
			},
		},
		"get-relational-database-log-streams": {
			Name:   "get-relational-database-log-streams",
			Fields: fields_get_relational_database_log_streams,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseLogStreamsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database_log_streams, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabaseLogStreams(ctx, input)
			},
		},
		"get-relational-database-master-user-password": {
			Name:   "get-relational-database-master-user-password",
			Fields: fields_get_relational_database_master_user_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseMasterUserPasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database_master_user_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabaseMasterUserPassword(ctx, input)
			},
		},
		"get-relational-database-metric-data": {
			Name:   "get-relational-database-metric-data",
			Fields: fields_get_relational_database_metric_data,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseMetricDataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database_metric_data, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabaseMetricData(ctx, input)
			},
		},
		"get-relational-database-parameters": {
			Name:   "get-relational-database-parameters",
			Fields: fields_get_relational_database_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabaseParameters(ctx, input)
			},
		},
		"get-relational-database-snapshot": {
			Name:   "get-relational-database-snapshot",
			Fields: fields_get_relational_database_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabaseSnapshot(ctx, input)
			},
		},
		"get-relational-database-snapshots": {
			Name:   "get-relational-database-snapshots",
			Fields: fields_get_relational_database_snapshots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabaseSnapshotsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_database_snapshots, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabaseSnapshots(ctx, input)
			},
		},
		"get-relational-databases": {
			Name:   "get-relational-databases",
			Fields: fields_get_relational_databases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRelationalDatabasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_relational_databases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRelationalDatabases(ctx, input)
			},
		},
		"get-setup-history": {
			Name:   "get-setup-history",
			Fields: fields_get_setup_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSetupHistoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_setup_history, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSetupHistory(ctx, input)
			},
		},
		"get-static-ip": {
			Name:   "get-static-ip",
			Fields: fields_get_static_ip,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStaticIpInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_static_ip, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStaticIp(ctx, input)
			},
		},
		"get-static-ips": {
			Name:   "get-static-ips",
			Fields: fields_get_static_ips,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStaticIpsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_static_ips, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStaticIps(ctx, input)
			},
		},
		"import-key-pair": {
			Name:   "import-key-pair",
			Fields: fields_import_key_pair,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportKeyPairInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_key_pair, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportKeyPair(ctx, input)
			},
		},
		"is-vpc-peered": {
			Name:   "is-vpc-peered",
			Fields: fields_is_vpc_peered,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.IsVpcPeeredInput{}
				if _, err := leanruntime.ApplyInput(input, fields_is_vpc_peered, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.IsVpcPeered(ctx, input)
			},
		},
		"open-instance-public-ports": {
			Name:   "open-instance-public-ports",
			Fields: fields_open_instance_public_ports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.OpenInstancePublicPortsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_open_instance_public_ports, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.OpenInstancePublicPorts(ctx, input)
			},
		},
		"peer-vpc": {
			Name:   "peer-vpc",
			Fields: fields_peer_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PeerVpcInput{}
				if _, err := leanruntime.ApplyInput(input, fields_peer_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PeerVpc(ctx, input)
			},
		},
		"put-alarm": {
			Name:   "put-alarm",
			Fields: fields_put_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAlarm(ctx, input)
			},
		},
		"put-instance-public-ports": {
			Name:   "put-instance-public-ports",
			Fields: fields_put_instance_public_ports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutInstancePublicPortsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_instance_public_ports, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutInstancePublicPorts(ctx, input)
			},
		},
		"reboot-instance": {
			Name:   "reboot-instance",
			Fields: fields_reboot_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootInstance(ctx, input)
			},
		},
		"reboot-relational-database": {
			Name:   "reboot-relational-database",
			Fields: fields_reboot_relational_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootRelationalDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_relational_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootRelationalDatabase(ctx, input)
			},
		},
		"register-container-image": {
			Name:   "register-container-image",
			Fields: fields_register_container_image,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterContainerImageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_container_image, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterContainerImage(ctx, input)
			},
		},
		"release-static-ip": {
			Name:   "release-static-ip",
			Fields: fields_release_static_ip,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReleaseStaticIpInput{}
				if _, err := leanruntime.ApplyInput(input, fields_release_static_ip, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReleaseStaticIp(ctx, input)
			},
		},
		"reset-distribution-cache": {
			Name:   "reset-distribution-cache",
			Fields: fields_reset_distribution_cache,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetDistributionCacheInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_distribution_cache, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetDistributionCache(ctx, input)
			},
		},
		"send-contact-method-verification": {
			Name:   "send-contact-method-verification",
			Fields: fields_send_contact_method_verification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendContactMethodVerificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_contact_method_verification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendContactMethodVerification(ctx, input)
			},
		},
		"set-ip-address-type": {
			Name:   "set-ip-address-type",
			Fields: fields_set_ip_address_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetIpAddressTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_ip_address_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetIpAddressType(ctx, input)
			},
		},
		"set-resource-access-for-bucket": {
			Name:   "set-resource-access-for-bucket",
			Fields: fields_set_resource_access_for_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetResourceAccessForBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_resource_access_for_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetResourceAccessForBucket(ctx, input)
			},
		},
		"setup-instance-https": {
			Name:   "setup-instance-https",
			Fields: fields_setup_instance_https,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetupInstanceHttpsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_setup_instance_https, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetupInstanceHttps(ctx, input)
			},
		},
		"start-gui-session": {
			Name:   "start-gui-session",
			Fields: fields_start_gui_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartGUISessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_gui_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartGUISession(ctx, input)
			},
		},
		"start-instance": {
			Name:   "start-instance",
			Fields: fields_start_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartInstance(ctx, input)
			},
		},
		"start-relational-database": {
			Name:   "start-relational-database",
			Fields: fields_start_relational_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartRelationalDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_relational_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartRelationalDatabase(ctx, input)
			},
		},
		"stop-gui-session": {
			Name:   "stop-gui-session",
			Fields: fields_stop_gui_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopGUISessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_gui_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopGUISession(ctx, input)
			},
		},
		"stop-instance": {
			Name:   "stop-instance",
			Fields: fields_stop_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopInstance(ctx, input)
			},
		},
		"stop-relational-database": {
			Name:   "stop-relational-database",
			Fields: fields_stop_relational_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopRelationalDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_relational_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopRelationalDatabase(ctx, input)
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
		"test-alarm": {
			Name:   "test-alarm",
			Fields: fields_test_alarm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestAlarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_alarm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestAlarm(ctx, input)
			},
		},
		"unpeer-vpc": {
			Name:   "unpeer-vpc",
			Fields: fields_unpeer_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnpeerVpcInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unpeer_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnpeerVpc(ctx, input)
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
		"update-bucket": {
			Name:   "update-bucket",
			Fields: fields_update_bucket,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBucketInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bucket, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBucket(ctx, input)
			},
		},
		"update-bucket-bundle": {
			Name:   "update-bucket-bundle",
			Fields: fields_update_bucket_bundle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBucketBundleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bucket_bundle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBucketBundle(ctx, input)
			},
		},
		"update-container-service": {
			Name:   "update-container-service",
			Fields: fields_update_container_service,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateContainerServiceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_container_service, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateContainerService(ctx, input)
			},
		},
		"update-distribution": {
			Name:   "update-distribution",
			Fields: fields_update_distribution,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDistributionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_distribution, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDistribution(ctx, input)
			},
		},
		"update-distribution-bundle": {
			Name:   "update-distribution-bundle",
			Fields: fields_update_distribution_bundle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDistributionBundleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_distribution_bundle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDistributionBundle(ctx, input)
			},
		},
		"update-domain-entry": {
			Name:   "update-domain-entry",
			Fields: fields_update_domain_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainEntry(ctx, input)
			},
		},
		"update-instance-metadata-options": {
			Name:   "update-instance-metadata-options",
			Fields: fields_update_instance_metadata_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateInstanceMetadataOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_instance_metadata_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateInstanceMetadataOptions(ctx, input)
			},
		},
		"update-load-balancer-attribute": {
			Name:   "update-load-balancer-attribute",
			Fields: fields_update_load_balancer_attribute,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLoadBalancerAttributeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_load_balancer_attribute, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLoadBalancerAttribute(ctx, input)
			},
		},
		"update-relational-database": {
			Name:   "update-relational-database",
			Fields: fields_update_relational_database,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRelationalDatabaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_relational_database, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRelationalDatabase(ctx, input)
			},
		},
		"update-relational-database-parameters": {
			Name:   "update-relational-database-parameters",
			Fields: fields_update_relational_database_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRelationalDatabaseParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_relational_database_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRelationalDatabaseParameters(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("lightsail", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
