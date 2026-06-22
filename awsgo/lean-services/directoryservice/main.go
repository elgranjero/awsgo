package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/directoryservice"
)

var fields_accept_shared_directory = []leanruntime.Field{
	{Name: "SharedDirectoryId", Flag: "shared-directory-id", Type: "*string", Required: true},
}

var fields_add_ip_routes = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "IpRoutes", Flag: "ip-routes", Type: "[]types.IpRoute", Required: true},
	{Name: "UpdateSecurityGroupForDirectoryControllers", Flag: "update-security-group-for-directory-controllers", Type: "bool", Required: false},
}

var fields_add_region = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: true},
	{Name: "VPCSettings", Flag: "vpc-settings", Type: "*types.DirectoryVpcSettings", Required: true},
}

var fields_add_tags_to_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_cancel_schema_extension = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "SchemaExtensionId", Flag: "schema-extension-id", Type: "*string", Required: true},
}

var fields_connect_directory = []leanruntime.Field{
	{Name: "ConnectSettings", Flag: "connect-settings", Type: "*types.DirectoryConnectSettings", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "ShortName", Flag: "short-name", Type: "*string", Required: false},
	{Name: "Size", Flag: "size", Type: "types.DirectorySize", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_alias = []leanruntime.Field{
	{Name: "Alias", Flag: "alias", Type: "*string", Required: true},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_create_computer = []leanruntime.Field{
	{Name: "ComputerAttributes", Flag: "computer-attributes", Type: "[]types.Attribute", Required: false},
	{Name: "ComputerName", Flag: "computer-name", Type: "*string", Required: true},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "OrganizationalUnitDistinguishedName", Flag: "organizational-unit-distinguished-name", Type: "*string", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
}

var fields_create_conditional_forwarder = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "DnsIpAddrs", Flag: "dns-ip-addrs", Type: "[]string", Required: false},
	{Name: "DnsIpv6Addrs", Flag: "dns-ipv6-addrs", Type: "[]string", Required: false},
	{Name: "RemoteDomainName", Flag: "remote-domain-name", Type: "*string", Required: true},
}

var fields_create_directory = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "ShortName", Flag: "short-name", Type: "*string", Required: false},
	{Name: "Size", Flag: "size", Type: "types.DirectorySize", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcSettings", Flag: "vpc-settings", Type: "*types.DirectoryVpcSettings", Required: false},
}

var fields_create_hybrid_ad = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "SecretArn", Flag: "secret-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_log_subscription = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "LogGroupName", Flag: "log-group-name", Type: "*string", Required: true},
}

var fields_create_microsoft_ad = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Edition", Flag: "edition", Type: "types.DirectoryEdition", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.NetworkType", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "ShortName", Flag: "short-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VpcSettings", Flag: "vpc-settings", Type: "*types.DirectoryVpcSettings", Required: true},
}

var fields_create_snapshot = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_create_trust = []leanruntime.Field{
	{Name: "ConditionalForwarderIpAddrs", Flag: "conditional-forwarder-ip-addrs", Type: "[]string", Required: false},
	{Name: "ConditionalForwarderIpv6Addrs", Flag: "conditional-forwarder-ipv6-addrs", Type: "[]string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "RemoteDomainName", Flag: "remote-domain-name", Type: "*string", Required: true},
	{Name: "SelectiveAuth", Flag: "selective-auth", Type: "types.SelectiveAuth", Required: false},
	{Name: "TrustDirection", Flag: "trust-direction", Type: "types.TrustDirection", Required: true},
	{Name: "TrustPassword", Flag: "trust-password", Type: "*string", Required: true},
	{Name: "TrustType", Flag: "trust-type", Type: "types.TrustType", Required: false},
}

var fields_delete_ad_assessment = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
}

var fields_delete_conditional_forwarder = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "RemoteDomainName", Flag: "remote-domain-name", Type: "*string", Required: true},
}

var fields_delete_directory = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_delete_log_subscription = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_delete_snapshot = []leanruntime.Field{
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_delete_trust = []leanruntime.Field{
	{Name: "DeleteAssociatedConditionalForwarder", Flag: "delete-associated-conditional-forwarder", Type: "bool", Required: false},
	{Name: "TrustId", Flag: "trust-id", Type: "*string", Required: true},
}

var fields_deregister_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_deregister_event_topic = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "TopicName", Flag: "topic-name", Type: "*string", Required: true},
}

var fields_describe_ad_assessment = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
}

var fields_describe_ca_enrollment_policy = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_describe_certificate = []leanruntime.Field{
	{Name: "CertificateId", Flag: "certificate-id", Type: "*string", Required: true},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_describe_client_authentication_settings = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ClientAuthenticationType", Required: false},
}

var fields_describe_conditional_forwarders = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "RemoteDomainNames", Flag: "remote-domain-names", Type: "[]string", Required: false},
}

var fields_describe_directories = []leanruntime.Field{
	{Name: "DirectoryIds", Flag: "directory-ids", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_directory_data_access = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_describe_domain_controllers = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "DomainControllerIds", Flag: "domain-controller-ids", Type: "[]string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_event_topics = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "TopicNames", Flag: "topic-names", Type: "[]string", Required: false},
}

var fields_describe_hybrid_ad_update = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UpdateType", Flag: "update-type", Type: "types.HybridUpdateType", Required: false},
}

var fields_describe_ldaps_settings = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.LDAPSType", Required: false},
}

var fields_describe_regions = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: false},
}

var fields_describe_settings = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.DirectoryConfigurationStatus", Required: false},
}

var fields_describe_shared_directories = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwnerDirectoryId", Flag: "owner-directory-id", Type: "*string", Required: true},
	{Name: "SharedDirectoryIds", Flag: "shared-directory-ids", Type: "[]string", Required: false},
}

var fields_describe_snapshots = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SnapshotIds", Flag: "snapshot-ids", Type: "[]string", Required: false},
}

var fields_describe_trusts = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrustIds", Flag: "trust-ids", Type: "[]string", Required: false},
}

var fields_describe_update_directory = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RegionName", Flag: "region-name", Type: "*string", Required: false},
	{Name: "UpdateType", Flag: "update-type", Type: "types.UpdateType", Required: true},
}

var fields_disable_ca_enrollment_policy = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_disable_client_authentication = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ClientAuthenticationType", Required: true},
}

var fields_disable_directory_data_access = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_disable_ldaps = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.LDAPSType", Required: true},
}

var fields_disable_radius = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_disable_sso = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_enable_ca_enrollment_policy = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "PcaConnectorArn", Flag: "pca-connector-arn", Type: "*string", Required: true},
}

var fields_enable_client_authentication = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ClientAuthenticationType", Required: true},
}

var fields_enable_directory_data_access = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_enable_ldaps = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.LDAPSType", Required: true},
}

var fields_enable_radius = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "RadiusSettings", Flag: "radius-settings", Type: "*types.RadiusSettings", Required: true},
}

var fields_enable_sso = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: false},
}

var fields_get_directory_limits = []leanruntime.Field{}

var fields_get_snapshot_limits = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_list_ad_assessments = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_certificates = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ip_routes = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_log_subscriptions = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_schema_extensions = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_register_certificate = []leanruntime.Field{
	{Name: "CertificateData", Flag: "certificate-data", Type: "*string", Required: true},
	{Name: "ClientCertAuthSettings", Flag: "client-cert-auth-settings", Type: "*types.ClientCertAuthSettings", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.CertificateType", Required: false},
}

var fields_register_event_topic = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "TopicName", Flag: "topic-name", Type: "*string", Required: true},
}

var fields_reject_shared_directory = []leanruntime.Field{
	{Name: "SharedDirectoryId", Flag: "shared-directory-id", Type: "*string", Required: true},
}

var fields_remove_ip_routes = []leanruntime.Field{
	{Name: "CidrIps", Flag: "cidr-ips", Type: "[]string", Required: false},
	{Name: "CidrIpv6s", Flag: "cidr-ipv6s", Type: "[]string", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_remove_region = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_remove_tags_from_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_reset_user_password = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "NewPassword", Flag: "new-password", Type: "*string", Required: true},
	{Name: "UserName", Flag: "user-name", Type: "*string", Required: true},
}

var fields_restore_from_snapshot = []leanruntime.Field{
	{Name: "SnapshotId", Flag: "snapshot-id", Type: "*string", Required: true},
}

var fields_share_directory = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "ShareMethod", Flag: "share-method", Type: "types.ShareMethod", Required: true},
	{Name: "ShareNotes", Flag: "share-notes", Type: "*string", Required: false},
	{Name: "ShareTarget", Flag: "share-target", Type: "*types.ShareTarget", Required: true},
}

var fields_start_ad_assessment = []leanruntime.Field{
	{Name: "AssessmentConfiguration", Flag: "assessment-configuration", Type: "*types.AssessmentConfiguration", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: false},
}

var fields_start_schema_extension = []leanruntime.Field{
	{Name: "CreateSnapshotBeforeSchemaExtension", Flag: "create-snapshot-before-schema-extension", Type: "bool", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "LdifContent", Flag: "ldif-content", Type: "*string", Required: true},
}

var fields_unshare_directory = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "UnshareTarget", Flag: "unshare-target", Type: "*types.UnshareTarget", Required: true},
}

var fields_update_conditional_forwarder = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "DnsIpAddrs", Flag: "dns-ip-addrs", Type: "[]string", Required: false},
	{Name: "DnsIpv6Addrs", Flag: "dns-ipv6-addrs", Type: "[]string", Required: false},
	{Name: "RemoteDomainName", Flag: "remote-domain-name", Type: "*string", Required: true},
}

var fields_update_directory_setup = []leanruntime.Field{
	{Name: "CreateSnapshotBeforeUpdate", Flag: "create-snapshot-before-update", Type: "*bool", Required: false},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "DirectorySizeUpdateSettings", Flag: "directory-size-update-settings", Type: "*types.DirectorySizeUpdateSettings", Required: false},
	{Name: "NetworkUpdateSettings", Flag: "network-update-settings", Type: "*types.NetworkUpdateSettings", Required: false},
	{Name: "OSUpdateSettings", Flag: "os-update-settings", Type: "*types.OSUpdateSettings", Required: false},
	{Name: "UpdateType", Flag: "update-type", Type: "types.UpdateType", Required: true},
}

var fields_update_hybrid_ad = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "HybridAdministratorAccountUpdate", Flag: "hybrid-administrator-account-update", Type: "*types.HybridAdministratorAccountUpdate", Required: false},
	{Name: "SelfManagedInstancesSettings", Flag: "self-managed-instances-settings", Type: "*types.HybridCustomerInstancesSettings", Required: false},
}

var fields_update_number_of_domain_controllers = []leanruntime.Field{
	{Name: "DesiredNumber", Flag: "desired-number", Type: "*int32", Required: true},
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
}

var fields_update_radius = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "RadiusSettings", Flag: "radius-settings", Type: "*types.RadiusSettings", Required: true},
}

var fields_update_settings = []leanruntime.Field{
	{Name: "DirectoryId", Flag: "directory-id", Type: "*string", Required: true},
	{Name: "Settings", Flag: "settings", Type: "[]types.Setting", Required: true},
}

var fields_update_trust = []leanruntime.Field{
	{Name: "SelectiveAuth", Flag: "selective-auth", Type: "types.SelectiveAuth", Required: false},
	{Name: "TrustId", Flag: "trust-id", Type: "*string", Required: true},
}

var fields_verify_trust = []leanruntime.Field{
	{Name: "TrustId", Flag: "trust-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-shared-directory": {
			Name:   "accept-shared-directory",
			Fields: fields_accept_shared_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptSharedDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_shared_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptSharedDirectory(ctx, input)
			},
		},
		"add-ip-routes": {
			Name:   "add-ip-routes",
			Fields: fields_add_ip_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddIpRoutesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_ip_routes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddIpRoutes(ctx, input)
			},
		},
		"add-region": {
			Name:   "add-region",
			Fields: fields_add_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddRegion(ctx, input)
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
		"cancel-schema-extension": {
			Name:   "cancel-schema-extension",
			Fields: fields_cancel_schema_extension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelSchemaExtensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_schema_extension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelSchemaExtension(ctx, input)
			},
		},
		"connect-directory": {
			Name:   "connect-directory",
			Fields: fields_connect_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ConnectDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_connect_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ConnectDirectory(ctx, input)
			},
		},
		"create-alias": {
			Name:   "create-alias",
			Fields: fields_create_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAlias(ctx, input)
			},
		},
		"create-computer": {
			Name:   "create-computer",
			Fields: fields_create_computer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateComputerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_computer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateComputer(ctx, input)
			},
		},
		"create-conditional-forwarder": {
			Name:   "create-conditional-forwarder",
			Fields: fields_create_conditional_forwarder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConditionalForwarderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_conditional_forwarder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConditionalForwarder(ctx, input)
			},
		},
		"create-directory": {
			Name:   "create-directory",
			Fields: fields_create_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDirectory(ctx, input)
			},
		},
		"create-hybrid-ad": {
			Name:   "create-hybrid-ad",
			Fields: fields_create_hybrid_ad,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHybridADInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hybrid_ad, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHybridAD(ctx, input)
			},
		},
		"create-log-subscription": {
			Name:   "create-log-subscription",
			Fields: fields_create_log_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLogSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_log_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLogSubscription(ctx, input)
			},
		},
		"create-microsoft-ad": {
			Name:   "create-microsoft-ad",
			Fields: fields_create_microsoft_ad,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMicrosoftADInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_microsoft_ad, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMicrosoftAD(ctx, input)
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
		"create-trust": {
			Name:   "create-trust",
			Fields: fields_create_trust,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrustInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trust, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrust(ctx, input)
			},
		},
		"delete-ad-assessment": {
			Name:   "delete-ad-assessment",
			Fields: fields_delete_ad_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteADAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ad_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteADAssessment(ctx, input)
			},
		},
		"delete-conditional-forwarder": {
			Name:   "delete-conditional-forwarder",
			Fields: fields_delete_conditional_forwarder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConditionalForwarderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_conditional_forwarder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConditionalForwarder(ctx, input)
			},
		},
		"delete-directory": {
			Name:   "delete-directory",
			Fields: fields_delete_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDirectory(ctx, input)
			},
		},
		"delete-log-subscription": {
			Name:   "delete-log-subscription",
			Fields: fields_delete_log_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLogSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_log_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLogSubscription(ctx, input)
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
		"delete-trust": {
			Name:   "delete-trust",
			Fields: fields_delete_trust,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrustInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trust, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrust(ctx, input)
			},
		},
		"deregister-certificate": {
			Name:   "deregister-certificate",
			Fields: fields_deregister_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterCertificate(ctx, input)
			},
		},
		"deregister-event-topic": {
			Name:   "deregister-event-topic",
			Fields: fields_deregister_event_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterEventTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_event_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterEventTopic(ctx, input)
			},
		},
		"describe-ad-assessment": {
			Name:   "describe-ad-assessment",
			Fields: fields_describe_ad_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeADAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ad_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeADAssessment(ctx, input)
			},
		},
		"describe-ca-enrollment-policy": {
			Name:   "describe-ca-enrollment-policy",
			Fields: fields_describe_ca_enrollment_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCAEnrollmentPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_ca_enrollment_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCAEnrollmentPolicy(ctx, input)
			},
		},
		"describe-certificate": {
			Name:   "describe-certificate",
			Fields: fields_describe_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCertificate(ctx, input)
			},
		},
		"describe-client-authentication-settings": {
			Name:   "describe-client-authentication-settings",
			Fields: fields_describe_client_authentication_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClientAuthenticationSettingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_client_authentication_settings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClientAuthenticationSettings(ctx, input)
				}
				var results []*svc.DescribeClientAuthenticationSettingsOutput
				p := svc.NewDescribeClientAuthenticationSettingsPaginator(client, input)
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
		"describe-conditional-forwarders": {
			Name:   "describe-conditional-forwarders",
			Fields: fields_describe_conditional_forwarders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConditionalForwardersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_conditional_forwarders, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConditionalForwarders(ctx, input)
			},
		},
		"describe-directories": {
			Name:   "describe-directories",
			Fields: fields_describe_directories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDirectoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_directories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDirectories(ctx, input)
				}
				var results []*svc.DescribeDirectoriesOutput
				p := svc.NewDescribeDirectoriesPaginator(client, input)
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
		"describe-directory-data-access": {
			Name:   "describe-directory-data-access",
			Fields: fields_describe_directory_data_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDirectoryDataAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_directory_data_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDirectoryDataAccess(ctx, input)
			},
		},
		"describe-domain-controllers": {
			Name:   "describe-domain-controllers",
			Fields: fields_describe_domain_controllers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainControllersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_domain_controllers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDomainControllers(ctx, input)
				}
				var results []*svc.DescribeDomainControllersOutput
				p := svc.NewDescribeDomainControllersPaginator(client, input)
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
		"describe-event-topics": {
			Name:   "describe-event-topics",
			Fields: fields_describe_event_topics,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEventTopicsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_event_topics, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEventTopics(ctx, input)
			},
		},
		"describe-hybrid-ad-update": {
			Name:   "describe-hybrid-ad-update",
			Fields: fields_describe_hybrid_ad_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHybridADUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_hybrid_ad_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHybridADUpdate(ctx, input)
			},
		},
		"describe-ldaps-settings": {
			Name:   "describe-ldaps-settings",
			Fields: fields_describe_ldaps_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeLDAPSSettingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_ldaps_settings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeLDAPSSettings(ctx, input)
				}
				var results []*svc.DescribeLDAPSSettingsOutput
				p := svc.NewDescribeLDAPSSettingsPaginator(client, input)
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
		"describe-regions": {
			Name:   "describe-regions",
			Fields: fields_describe_regions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRegionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_regions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeRegions(ctx, input)
				}
				var results []*svc.DescribeRegionsOutput
				p := svc.NewDescribeRegionsPaginator(client, input)
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
		"describe-settings": {
			Name:   "describe-settings",
			Fields: fields_describe_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSettings(ctx, input)
			},
		},
		"describe-shared-directories": {
			Name:   "describe-shared-directories",
			Fields: fields_describe_shared_directories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSharedDirectoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_shared_directories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeSharedDirectories(ctx, input)
				}
				var results []*svc.DescribeSharedDirectoriesOutput
				p := svc.NewDescribeSharedDirectoriesPaginator(client, input)
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
		"describe-trusts": {
			Name:   "describe-trusts",
			Fields: fields_describe_trusts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrustsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_trusts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeTrusts(ctx, input)
				}
				var results []*svc.DescribeTrustsOutput
				p := svc.NewDescribeTrustsPaginator(client, input)
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
		"describe-update-directory": {
			Name:   "describe-update-directory",
			Fields: fields_describe_update_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUpdateDirectoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_update_directory, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeUpdateDirectory(ctx, input)
				}
				var results []*svc.DescribeUpdateDirectoryOutput
				p := svc.NewDescribeUpdateDirectoryPaginator(client, input)
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
		"disable-ca-enrollment-policy": {
			Name:   "disable-ca-enrollment-policy",
			Fields: fields_disable_ca_enrollment_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableCAEnrollmentPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_ca_enrollment_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableCAEnrollmentPolicy(ctx, input)
			},
		},
		"disable-client-authentication": {
			Name:   "disable-client-authentication",
			Fields: fields_disable_client_authentication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableClientAuthenticationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_client_authentication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableClientAuthentication(ctx, input)
			},
		},
		"disable-directory-data-access": {
			Name:   "disable-directory-data-access",
			Fields: fields_disable_directory_data_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableDirectoryDataAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_directory_data_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableDirectoryDataAccess(ctx, input)
			},
		},
		"disable-ldaps": {
			Name:   "disable-ldaps",
			Fields: fields_disable_ldaps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableLDAPSInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_ldaps, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableLDAPS(ctx, input)
			},
		},
		"disable-radius": {
			Name:   "disable-radius",
			Fields: fields_disable_radius,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableRadiusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_radius, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableRadius(ctx, input)
			},
		},
		"disable-sso": {
			Name:   "disable-sso",
			Fields: fields_disable_sso,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableSsoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_sso, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableSso(ctx, input)
			},
		},
		"enable-ca-enrollment-policy": {
			Name:   "enable-ca-enrollment-policy",
			Fields: fields_enable_ca_enrollment_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableCAEnrollmentPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_ca_enrollment_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableCAEnrollmentPolicy(ctx, input)
			},
		},
		"enable-client-authentication": {
			Name:   "enable-client-authentication",
			Fields: fields_enable_client_authentication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableClientAuthenticationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_client_authentication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableClientAuthentication(ctx, input)
			},
		},
		"enable-directory-data-access": {
			Name:   "enable-directory-data-access",
			Fields: fields_enable_directory_data_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableDirectoryDataAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_directory_data_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableDirectoryDataAccess(ctx, input)
			},
		},
		"enable-ldaps": {
			Name:   "enable-ldaps",
			Fields: fields_enable_ldaps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableLDAPSInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_ldaps, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableLDAPS(ctx, input)
			},
		},
		"enable-radius": {
			Name:   "enable-radius",
			Fields: fields_enable_radius,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableRadiusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_radius, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableRadius(ctx, input)
			},
		},
		"enable-sso": {
			Name:   "enable-sso",
			Fields: fields_enable_sso,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableSsoInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_sso, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableSso(ctx, input)
			},
		},
		"get-directory-limits": {
			Name:   "get-directory-limits",
			Fields: fields_get_directory_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDirectoryLimitsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_directory_limits, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDirectoryLimits(ctx, input)
			},
		},
		"get-snapshot-limits": {
			Name:   "get-snapshot-limits",
			Fields: fields_get_snapshot_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSnapshotLimitsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_snapshot_limits, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSnapshotLimits(ctx, input)
			},
		},
		"list-ad-assessments": {
			Name:   "list-ad-assessments",
			Fields: fields_list_ad_assessments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListADAssessmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ad_assessments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListADAssessments(ctx, input)
				}
				var results []*svc.ListADAssessmentsOutput
				p := svc.NewListADAssessmentsPaginator(client, input)
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
		"list-certificates": {
			Name:   "list-certificates",
			Fields: fields_list_certificates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCertificatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_certificates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCertificates(ctx, input)
				}
				var results []*svc.ListCertificatesOutput
				p := svc.NewListCertificatesPaginator(client, input)
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
		"list-ip-routes": {
			Name:   "list-ip-routes",
			Fields: fields_list_ip_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIpRoutesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ip_routes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIpRoutes(ctx, input)
				}
				var results []*svc.ListIpRoutesOutput
				p := svc.NewListIpRoutesPaginator(client, input)
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
		"list-log-subscriptions": {
			Name:   "list-log-subscriptions",
			Fields: fields_list_log_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLogSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_log_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLogSubscriptions(ctx, input)
				}
				var results []*svc.ListLogSubscriptionsOutput
				p := svc.NewListLogSubscriptionsPaginator(client, input)
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
		"list-schema-extensions": {
			Name:   "list-schema-extensions",
			Fields: fields_list_schema_extensions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSchemaExtensionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_schema_extensions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSchemaExtensions(ctx, input)
				}
				var results []*svc.ListSchemaExtensionsOutput
				p := svc.NewListSchemaExtensionsPaginator(client, input)
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
		"register-certificate": {
			Name:   "register-certificate",
			Fields: fields_register_certificate,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterCertificateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_certificate, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterCertificate(ctx, input)
			},
		},
		"register-event-topic": {
			Name:   "register-event-topic",
			Fields: fields_register_event_topic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterEventTopicInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_event_topic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterEventTopic(ctx, input)
			},
		},
		"reject-shared-directory": {
			Name:   "reject-shared-directory",
			Fields: fields_reject_shared_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectSharedDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_shared_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectSharedDirectory(ctx, input)
			},
		},
		"remove-ip-routes": {
			Name:   "remove-ip-routes",
			Fields: fields_remove_ip_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveIpRoutesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_ip_routes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveIpRoutes(ctx, input)
			},
		},
		"remove-region": {
			Name:   "remove-region",
			Fields: fields_remove_region,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveRegionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_region, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveRegion(ctx, input)
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
		"reset-user-password": {
			Name:   "reset-user-password",
			Fields: fields_reset_user_password,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetUserPasswordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_user_password, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetUserPassword(ctx, input)
			},
		},
		"restore-from-snapshot": {
			Name:   "restore-from-snapshot",
			Fields: fields_restore_from_snapshot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestoreFromSnapshotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_from_snapshot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestoreFromSnapshot(ctx, input)
			},
		},
		"share-directory": {
			Name:   "share-directory",
			Fields: fields_share_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ShareDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_share_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ShareDirectory(ctx, input)
			},
		},
		"start-ad-assessment": {
			Name:   "start-ad-assessment",
			Fields: fields_start_ad_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartADAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_ad_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartADAssessment(ctx, input)
			},
		},
		"start-schema-extension": {
			Name:   "start-schema-extension",
			Fields: fields_start_schema_extension,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSchemaExtensionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_schema_extension, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSchemaExtension(ctx, input)
			},
		},
		"unshare-directory": {
			Name:   "unshare-directory",
			Fields: fields_unshare_directory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UnshareDirectoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_unshare_directory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UnshareDirectory(ctx, input)
			},
		},
		"update-conditional-forwarder": {
			Name:   "update-conditional-forwarder",
			Fields: fields_update_conditional_forwarder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConditionalForwarderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_conditional_forwarder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConditionalForwarder(ctx, input)
			},
		},
		"update-directory-setup": {
			Name:   "update-directory-setup",
			Fields: fields_update_directory_setup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDirectorySetupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_directory_setup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDirectorySetup(ctx, input)
			},
		},
		"update-hybrid-ad": {
			Name:   "update-hybrid-ad",
			Fields: fields_update_hybrid_ad,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHybridADInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hybrid_ad, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHybridAD(ctx, input)
			},
		},
		"update-number-of-domain-controllers": {
			Name:   "update-number-of-domain-controllers",
			Fields: fields_update_number_of_domain_controllers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNumberOfDomainControllersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_number_of_domain_controllers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNumberOfDomainControllers(ctx, input)
			},
		},
		"update-radius": {
			Name:   "update-radius",
			Fields: fields_update_radius,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRadiusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_radius, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRadius(ctx, input)
			},
		},
		"update-settings": {
			Name:   "update-settings",
			Fields: fields_update_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSettings(ctx, input)
			},
		},
		"update-trust": {
			Name:   "update-trust",
			Fields: fields_update_trust,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrustInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trust, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrust(ctx, input)
			},
		},
		"verify-trust": {
			Name:   "verify-trust",
			Fields: fields_verify_trust,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyTrustInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_trust, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyTrust(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("directoryservice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
