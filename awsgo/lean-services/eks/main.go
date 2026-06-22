package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/eks"
)

var fields_associate_access_policy = []leanruntime.Field{
	{Name: "AccessScope", Flag: "access-scope", Type: "*types.AccessScope", Required: true},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
}

var fields_associate_encryption_config = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "EncryptionConfig", Flag: "encryption-config", Type: "[]types.EncryptionConfig", Required: true},
}

var fields_associate_identity_provider_config = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Oidc", Flag: "oidc", Type: "*types.OidcIdentityProviderConfigRequest", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_access_entry = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "KubernetesGroups", Flag: "kubernetes-groups", Type: "[]string", Required: false},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_create_addon = []leanruntime.Field{
	{Name: "AddonName", Flag: "addon-name", Type: "*string", Required: true},
	{Name: "AddonVersion", Flag: "addon-version", Type: "*string", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "ConfigurationValues", Flag: "configuration-values", Type: "*string", Required: false},
	{Name: "NamespaceConfig", Flag: "namespace-config", Type: "*types.AddonNamespaceConfigRequest", Required: false},
	{Name: "PodIdentityAssociations", Flag: "pod-identity-associations", Type: "[]types.AddonPodIdentityAssociations", Required: false},
	{Name: "ResolveConflicts", Flag: "resolve-conflicts", Type: "types.ResolveConflicts", Required: false},
	{Name: "ServiceAccountRoleArn", Flag: "service-account-role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_capability = []leanruntime.Field{
	{Name: "CapabilityName", Flag: "capability-name", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.CapabilityConfigurationRequest", Required: false},
	{Name: "DeletePropagationPolicy", Flag: "delete-propagation-policy", Type: "types.CapabilityDeletePropagationPolicy", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.CapabilityType", Required: true},
}

var fields_create_cluster = []leanruntime.Field{
	{Name: "AccessConfig", Flag: "access-config", Type: "*types.CreateAccessConfigRequest", Required: false},
	{Name: "BootstrapSelfManagedAddons", Flag: "bootstrap-self-managed-addons", Type: "*bool", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ComputeConfig", Flag: "compute-config", Type: "*types.ComputeConfigRequest", Required: false},
	{Name: "ControlPlaneScalingConfig", Flag: "control-plane-scaling-config", Type: "*types.ControlPlaneScalingConfig", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "EncryptionConfig", Flag: "encryption-config", Type: "[]types.EncryptionConfig", Required: false},
	{Name: "KubernetesNetworkConfig", Flag: "kubernetes-network-config", Type: "*types.KubernetesNetworkConfigRequest", Required: false},
	{Name: "Logging", Flag: "logging", Type: "*types.Logging", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutpostConfig", Flag: "outpost-config", Type: "*types.OutpostConfigRequest", Required: false},
	{Name: "RemoteNetworkConfig", Flag: "remote-network-config", Type: "*types.RemoteNetworkConfigRequest", Required: false},
	{Name: "ResourcesVpcConfig", Flag: "resources-vpc-config", Type: "*types.VpcConfigRequest", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "StorageConfig", Flag: "storage-config", Type: "*types.StorageConfigRequest", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UpgradePolicy", Flag: "upgrade-policy", Type: "*types.UpgradePolicyRequest", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
	{Name: "ZonalShiftConfig", Flag: "zonal-shift-config", Type: "*types.ZonalShiftConfigRequest", Required: false},
}

var fields_create_eks_anywhere_subscription = []leanruntime.Field{
	{Name: "AutoRenew", Flag: "auto-renew", Type: "bool", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "LicenseQuantity", Flag: "license-quantity", Type: "int32", Required: false},
	{Name: "LicenseType", Flag: "license-type", Type: "types.EksAnywhereSubscriptionLicenseType", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Term", Flag: "term", Type: "*types.EksAnywhereSubscriptionTerm", Required: true},
}

var fields_create_fargate_profile = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "FargateProfileName", Flag: "fargate-profile-name", Type: "*string", Required: true},
	{Name: "PodExecutionRoleArn", Flag: "pod-execution-role-arn", Type: "*string", Required: true},
	{Name: "Selectors", Flag: "selectors", Type: "[]types.FargateProfileSelector", Required: false},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_nodegroup = []leanruntime.Field{
	{Name: "AmiType", Flag: "ami-type", Type: "types.AMITypes", Required: false},
	{Name: "CapacityType", Flag: "capacity-type", Type: "types.CapacityTypes", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "DiskSize", Flag: "disk-size", Type: "*int32", Required: false},
	{Name: "InstanceTypes", Flag: "instance-types", Type: "[]string", Required: false},
	{Name: "Labels", Flag: "labels", Type: "map[string]string", Required: false},
	{Name: "LaunchTemplate", Flag: "launch-template", Type: "*types.LaunchTemplateSpecification", Required: false},
	{Name: "NodeRepairConfig", Flag: "node-repair-config", Type: "*types.NodeRepairConfig", Required: false},
	{Name: "NodeRole", Flag: "node-role", Type: "*string", Required: true},
	{Name: "NodegroupName", Flag: "nodegroup-name", Type: "*string", Required: true},
	{Name: "ReleaseVersion", Flag: "release-version", Type: "*string", Required: false},
	{Name: "RemoteAccess", Flag: "remote-access", Type: "*types.RemoteAccessConfig", Required: false},
	{Name: "ScalingConfig", Flag: "scaling-config", Type: "*types.NodegroupScalingConfig", Required: false},
	{Name: "Subnets", Flag: "subnets", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Taints", Flag: "taints", Type: "[]types.Taint", Required: false},
	{Name: "UpdateConfig", Flag: "update-config", Type: "*types.NodegroupUpdateConfig", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_create_pod_identity_association = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "DisableSessionTags", Flag: "disable-session-tags", Type: "*bool", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: true},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "ServiceAccount", Flag: "service-account", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetRoleArn", Flag: "target-role-arn", Type: "*string", Required: false},
}

var fields_delete_access_entry = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
}

var fields_delete_addon = []leanruntime.Field{
	{Name: "AddonName", Flag: "addon-name", Type: "*string", Required: true},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Preserve", Flag: "preserve", Type: "bool", Required: false},
}

var fields_delete_capability = []leanruntime.Field{
	{Name: "CapabilityName", Flag: "capability-name", Type: "*string", Required: true},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_delete_cluster = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_eks_anywhere_subscription = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_fargate_profile = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "FargateProfileName", Flag: "fargate-profile-name", Type: "*string", Required: true},
}

var fields_delete_nodegroup = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NodegroupName", Flag: "nodegroup-name", Type: "*string", Required: true},
}

var fields_delete_pod_identity_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_deregister_cluster = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_access_entry = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
}

var fields_describe_addon = []leanruntime.Field{
	{Name: "AddonName", Flag: "addon-name", Type: "*string", Required: true},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_describe_addon_configuration = []leanruntime.Field{
	{Name: "AddonName", Flag: "addon-name", Type: "*string", Required: true},
	{Name: "AddonVersion", Flag: "addon-version", Type: "*string", Required: true},
}

var fields_describe_addon_versions = []leanruntime.Field{
	{Name: "AddonName", Flag: "addon-name", Type: "*string", Required: false},
	{Name: "KubernetesVersion", Flag: "kubernetes-version", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Owners", Flag: "owners", Type: "[]string", Required: false},
	{Name: "Publishers", Flag: "publishers", Type: "[]string", Required: false},
	{Name: "Types", Flag: "types", Type: "[]string", Required: false},
}

var fields_describe_capability = []leanruntime.Field{
	{Name: "CapabilityName", Flag: "capability-name", Type: "*string", Required: true},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_describe_cluster = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_describe_cluster_versions = []leanruntime.Field{
	{Name: "ClusterType", Flag: "cluster-type", Type: "*string", Required: false},
	{Name: "ClusterVersions", Flag: "cluster-versions", Type: "[]string", Required: false},
	{Name: "DefaultOnly", Flag: "default-only", Type: "*bool", Required: false},
	{Name: "IncludeAll", Flag: "include-all", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ClusterVersionStatus", Required: false},
	{Name: "VersionStatus", Flag: "version-status", Type: "types.VersionStatus", Required: false},
}

var fields_describe_eks_anywhere_subscription = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_fargate_profile = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "FargateProfileName", Flag: "fargate-profile-name", Type: "*string", Required: true},
}

var fields_describe_identity_provider_config = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "IdentityProviderConfig", Flag: "identity-provider-config", Type: "*types.IdentityProviderConfig", Required: true},
}

var fields_describe_insight = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_insights_refresh = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_describe_nodegroup = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "NodegroupName", Flag: "nodegroup-name", Type: "*string", Required: true},
}

var fields_describe_pod_identity_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_describe_update = []leanruntime.Field{
	{Name: "AddonName", Flag: "addon-name", Type: "*string", Required: false},
	{Name: "CapabilityName", Flag: "capability-name", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NodegroupName", Flag: "nodegroup-name", Type: "*string", Required: false},
	{Name: "UpdateId", Flag: "update-id", Type: "*string", Required: true},
}

var fields_disassociate_access_policy = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "PolicyArn", Flag: "policy-arn", Type: "*string", Required: true},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
}

var fields_disassociate_identity_provider_config = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "IdentityProviderConfig", Flag: "identity-provider-config", Type: "*types.IdentityProviderConfig", Required: true},
}

var fields_list_access_entries = []leanruntime.Field{
	{Name: "AssociatedPolicyArn", Flag: "associated-policy-arn", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_access_policies = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_addons = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_associated_access_policies = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
}

var fields_list_capabilities = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_clusters = []leanruntime.Field{
	{Name: "Include", Flag: "include", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_eks_anywhere_subscriptions = []leanruntime.Field{
	{Name: "IncludeStatus", Flag: "include-status", Type: "[]types.EksAnywhereSubscriptionStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_fargate_profiles = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_identity_provider_configs = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_insights = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Filter", Flag: "filter", Type: "*types.InsightsFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_nodegroups = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_pod_identity_associations = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Namespace", Flag: "namespace", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServiceAccount", Flag: "service-account", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_updates = []leanruntime.Field{
	{Name: "AddonName", Flag: "addon-name", Type: "*string", Required: false},
	{Name: "CapabilityName", Flag: "capability-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "NodegroupName", Flag: "nodegroup-name", Type: "*string", Required: false},
}

var fields_register_cluster = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ConnectorConfig", Flag: "connector-config", Type: "*types.ConnectorConfigRequest", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_insights_refresh = []leanruntime.Field{
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_access_entry = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "KubernetesGroups", Flag: "kubernetes-groups", Type: "[]string", Required: false},
	{Name: "PrincipalArn", Flag: "principal-arn", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_update_addon = []leanruntime.Field{
	{Name: "AddonName", Flag: "addon-name", Type: "*string", Required: true},
	{Name: "AddonVersion", Flag: "addon-version", Type: "*string", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "ConfigurationValues", Flag: "configuration-values", Type: "*string", Required: false},
	{Name: "PodIdentityAssociations", Flag: "pod-identity-associations", Type: "[]types.AddonPodIdentityAssociations", Required: false},
	{Name: "ResolveConflicts", Flag: "resolve-conflicts", Type: "types.ResolveConflicts", Required: false},
	{Name: "ServiceAccountRoleArn", Flag: "service-account-role-arn", Type: "*string", Required: false},
}

var fields_update_capability = []leanruntime.Field{
	{Name: "CapabilityName", Flag: "capability-name", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.UpdateCapabilityConfiguration", Required: false},
	{Name: "DeletePropagationPolicy", Flag: "delete-propagation-policy", Type: "types.CapabilityDeletePropagationPolicy", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_cluster_config = []leanruntime.Field{
	{Name: "AccessConfig", Flag: "access-config", Type: "*types.UpdateAccessConfigRequest", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ComputeConfig", Flag: "compute-config", Type: "*types.ComputeConfigRequest", Required: false},
	{Name: "ControlPlaneScalingConfig", Flag: "control-plane-scaling-config", Type: "*types.ControlPlaneScalingConfig", Required: false},
	{Name: "DeletionProtection", Flag: "deletion-protection", Type: "*bool", Required: false},
	{Name: "KubernetesNetworkConfig", Flag: "kubernetes-network-config", Type: "*types.KubernetesNetworkConfigRequest", Required: false},
	{Name: "Logging", Flag: "logging", Type: "*types.Logging", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RemoteNetworkConfig", Flag: "remote-network-config", Type: "*types.RemoteNetworkConfigRequest", Required: false},
	{Name: "ResourcesVpcConfig", Flag: "resources-vpc-config", Type: "*types.VpcConfigRequest", Required: false},
	{Name: "StorageConfig", Flag: "storage-config", Type: "*types.StorageConfigRequest", Required: false},
	{Name: "UpgradePolicy", Flag: "upgrade-policy", Type: "*types.UpgradePolicyRequest", Required: false},
	{Name: "ZonalShiftConfig", Flag: "zonal-shift-config", Type: "*types.ZonalShiftConfigRequest", Required: false},
}

var fields_update_cluster_version = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_update_eks_anywhere_subscription = []leanruntime.Field{
	{Name: "AutoRenew", Flag: "auto-renew", Type: "bool", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_nodegroup_config = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Labels", Flag: "labels", Type: "*types.UpdateLabelsPayload", Required: false},
	{Name: "NodeRepairConfig", Flag: "node-repair-config", Type: "*types.NodeRepairConfig", Required: false},
	{Name: "NodegroupName", Flag: "nodegroup-name", Type: "*string", Required: true},
	{Name: "ScalingConfig", Flag: "scaling-config", Type: "*types.NodegroupScalingConfig", Required: false},
	{Name: "Taints", Flag: "taints", Type: "*types.UpdateTaintsPayload", Required: false},
	{Name: "UpdateConfig", Flag: "update-config", Type: "*types.NodegroupUpdateConfig", Required: false},
}

var fields_update_nodegroup_version = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
	{Name: "LaunchTemplate", Flag: "launch-template", Type: "*types.LaunchTemplateSpecification", Required: false},
	{Name: "NodegroupName", Flag: "nodegroup-name", Type: "*string", Required: true},
	{Name: "ReleaseVersion", Flag: "release-version", Type: "*string", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_update_pod_identity_association = []leanruntime.Field{
	{Name: "AssociationId", Flag: "association-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "ClusterName", Flag: "cluster-name", Type: "*string", Required: true},
	{Name: "DisableSessionTags", Flag: "disable-session-tags", Type: "*bool", Required: false},
	{Name: "Policy", Flag: "policy", Type: "*string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "TargetRoleArn", Flag: "target-role-arn", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-access-policy": {
			Name:   "associate-access-policy",
			Fields: fields_associate_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAccessPolicy(ctx, input)
			},
		},
		"associate-encryption-config": {
			Name:   "associate-encryption-config",
			Fields: fields_associate_encryption_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateEncryptionConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_encryption_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateEncryptionConfig(ctx, input)
			},
		},
		"associate-identity-provider-config": {
			Name:   "associate-identity-provider-config",
			Fields: fields_associate_identity_provider_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateIdentityProviderConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_identity_provider_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateIdentityProviderConfig(ctx, input)
			},
		},
		"create-access-entry": {
			Name:   "create-access-entry",
			Fields: fields_create_access_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_access_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessEntry(ctx, input)
			},
		},
		"create-addon": {
			Name:   "create-addon",
			Fields: fields_create_addon,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAddonInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_addon, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAddon(ctx, input)
			},
		},
		"create-capability": {
			Name:   "create-capability",
			Fields: fields_create_capability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCapabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_capability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCapability(ctx, input)
			},
		},
		"create-cluster": {
			Name:   "create-cluster",
			Fields: fields_create_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCluster(ctx, input)
			},
		},
		"create-eks-anywhere-subscription": {
			Name:   "create-eks-anywhere-subscription",
			Fields: fields_create_eks_anywhere_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEksAnywhereSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_eks_anywhere_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEksAnywhereSubscription(ctx, input)
			},
		},
		"create-fargate-profile": {
			Name:   "create-fargate-profile",
			Fields: fields_create_fargate_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFargateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fargate_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFargateProfile(ctx, input)
			},
		},
		"create-nodegroup": {
			Name:   "create-nodegroup",
			Fields: fields_create_nodegroup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNodegroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_nodegroup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNodegroup(ctx, input)
			},
		},
		"create-pod-identity-association": {
			Name:   "create-pod-identity-association",
			Fields: fields_create_pod_identity_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePodIdentityAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pod_identity_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePodIdentityAssociation(ctx, input)
			},
		},
		"delete-access-entry": {
			Name:   "delete-access-entry",
			Fields: fields_delete_access_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_access_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessEntry(ctx, input)
			},
		},
		"delete-addon": {
			Name:   "delete-addon",
			Fields: fields_delete_addon,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAddonInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_addon, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAddon(ctx, input)
			},
		},
		"delete-capability": {
			Name:   "delete-capability",
			Fields: fields_delete_capability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCapabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_capability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCapability(ctx, input)
			},
		},
		"delete-cluster": {
			Name:   "delete-cluster",
			Fields: fields_delete_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCluster(ctx, input)
			},
		},
		"delete-eks-anywhere-subscription": {
			Name:   "delete-eks-anywhere-subscription",
			Fields: fields_delete_eks_anywhere_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEksAnywhereSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_eks_anywhere_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEksAnywhereSubscription(ctx, input)
			},
		},
		"delete-fargate-profile": {
			Name:   "delete-fargate-profile",
			Fields: fields_delete_fargate_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFargateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fargate_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFargateProfile(ctx, input)
			},
		},
		"delete-nodegroup": {
			Name:   "delete-nodegroup",
			Fields: fields_delete_nodegroup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNodegroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_nodegroup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNodegroup(ctx, input)
			},
		},
		"delete-pod-identity-association": {
			Name:   "delete-pod-identity-association",
			Fields: fields_delete_pod_identity_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePodIdentityAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pod_identity_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePodIdentityAssociation(ctx, input)
			},
		},
		"deregister-cluster": {
			Name:   "deregister-cluster",
			Fields: fields_deregister_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterCluster(ctx, input)
			},
		},
		"describe-access-entry": {
			Name:   "describe-access-entry",
			Fields: fields_describe_access_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccessEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_access_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccessEntry(ctx, input)
			},
		},
		"describe-addon": {
			Name:   "describe-addon",
			Fields: fields_describe_addon,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAddonInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_addon, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAddon(ctx, input)
			},
		},
		"describe-addon-configuration": {
			Name:   "describe-addon-configuration",
			Fields: fields_describe_addon_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAddonConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_addon_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAddonConfiguration(ctx, input)
			},
		},
		"describe-addon-versions": {
			Name:   "describe-addon-versions",
			Fields: fields_describe_addon_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAddonVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_addon_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeAddonVersions(ctx, input)
				}
				var results []*svc.DescribeAddonVersionsOutput
				p := svc.NewDescribeAddonVersionsPaginator(client, input)
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
		"describe-capability": {
			Name:   "describe-capability",
			Fields: fields_describe_capability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCapabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_capability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCapability(ctx, input)
			},
		},
		"describe-cluster": {
			Name:   "describe-cluster",
			Fields: fields_describe_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCluster(ctx, input)
			},
		},
		"describe-cluster-versions": {
			Name:   "describe-cluster-versions",
			Fields: fields_describe_cluster_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeClusterVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_cluster_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeClusterVersions(ctx, input)
				}
				var results []*svc.DescribeClusterVersionsOutput
				p := svc.NewDescribeClusterVersionsPaginator(client, input)
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
		"describe-eks-anywhere-subscription": {
			Name:   "describe-eks-anywhere-subscription",
			Fields: fields_describe_eks_anywhere_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEksAnywhereSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_eks_anywhere_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEksAnywhereSubscription(ctx, input)
			},
		},
		"describe-fargate-profile": {
			Name:   "describe-fargate-profile",
			Fields: fields_describe_fargate_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFargateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fargate_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFargateProfile(ctx, input)
			},
		},
		"describe-identity-provider-config": {
			Name:   "describe-identity-provider-config",
			Fields: fields_describe_identity_provider_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeIdentityProviderConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_identity_provider_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeIdentityProviderConfig(ctx, input)
			},
		},
		"describe-insight": {
			Name:   "describe-insight",
			Fields: fields_describe_insight,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInsightInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_insight, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInsight(ctx, input)
			},
		},
		"describe-insights-refresh": {
			Name:   "describe-insights-refresh",
			Fields: fields_describe_insights_refresh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInsightsRefreshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_insights_refresh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeInsightsRefresh(ctx, input)
			},
		},
		"describe-nodegroup": {
			Name:   "describe-nodegroup",
			Fields: fields_describe_nodegroup,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeNodegroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_nodegroup, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeNodegroup(ctx, input)
			},
		},
		"describe-pod-identity-association": {
			Name:   "describe-pod-identity-association",
			Fields: fields_describe_pod_identity_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePodIdentityAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_pod_identity_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePodIdentityAssociation(ctx, input)
			},
		},
		"describe-update": {
			Name:   "describe-update",
			Fields: fields_describe_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUpdate(ctx, input)
			},
		},
		"disassociate-access-policy": {
			Name:   "disassociate-access-policy",
			Fields: fields_disassociate_access_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAccessPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_access_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAccessPolicy(ctx, input)
			},
		},
		"disassociate-identity-provider-config": {
			Name:   "disassociate-identity-provider-config",
			Fields: fields_disassociate_identity_provider_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateIdentityProviderConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_identity_provider_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateIdentityProviderConfig(ctx, input)
			},
		},
		"list-access-entries": {
			Name:   "list-access-entries",
			Fields: fields_list_access_entries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessEntriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_entries, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessEntries(ctx, input)
				}
				var results []*svc.ListAccessEntriesOutput
				p := svc.NewListAccessEntriesPaginator(client, input)
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
		"list-access-policies": {
			Name:   "list-access-policies",
			Fields: fields_list_access_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_access_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessPolicies(ctx, input)
				}
				var results []*svc.ListAccessPoliciesOutput
				p := svc.NewListAccessPoliciesPaginator(client, input)
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
		"list-addons": {
			Name:   "list-addons",
			Fields: fields_list_addons,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAddonsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_addons, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAddons(ctx, input)
				}
				var results []*svc.ListAddonsOutput
				p := svc.NewListAddonsPaginator(client, input)
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
		"list-associated-access-policies": {
			Name:   "list-associated-access-policies",
			Fields: fields_list_associated_access_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedAccessPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associated_access_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociatedAccessPolicies(ctx, input)
				}
				var results []*svc.ListAssociatedAccessPoliciesOutput
				p := svc.NewListAssociatedAccessPoliciesPaginator(client, input)
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
		"list-capabilities": {
			Name:   "list-capabilities",
			Fields: fields_list_capabilities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCapabilitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_capabilities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCapabilities(ctx, input)
				}
				var results []*svc.ListCapabilitiesOutput
				p := svc.NewListCapabilitiesPaginator(client, input)
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
		"list-clusters": {
			Name:   "list-clusters",
			Fields: fields_list_clusters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListClustersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_clusters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListClusters(ctx, input)
				}
				var results []*svc.ListClustersOutput
				p := svc.NewListClustersPaginator(client, input)
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
		"list-eks-anywhere-subscriptions": {
			Name:   "list-eks-anywhere-subscriptions",
			Fields: fields_list_eks_anywhere_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEksAnywhereSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_eks_anywhere_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEksAnywhereSubscriptions(ctx, input)
				}
				var results []*svc.ListEksAnywhereSubscriptionsOutput
				p := svc.NewListEksAnywhereSubscriptionsPaginator(client, input)
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
		"list-fargate-profiles": {
			Name:   "list-fargate-profiles",
			Fields: fields_list_fargate_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFargateProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fargate_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFargateProfiles(ctx, input)
				}
				var results []*svc.ListFargateProfilesOutput
				p := svc.NewListFargateProfilesPaginator(client, input)
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
		"list-identity-provider-configs": {
			Name:   "list-identity-provider-configs",
			Fields: fields_list_identity_provider_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIdentityProviderConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_identity_provider_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIdentityProviderConfigs(ctx, input)
				}
				var results []*svc.ListIdentityProviderConfigsOutput
				p := svc.NewListIdentityProviderConfigsPaginator(client, input)
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
		"list-insights": {
			Name:   "list-insights",
			Fields: fields_list_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInsightsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_insights, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInsights(ctx, input)
				}
				var results []*svc.ListInsightsOutput
				p := svc.NewListInsightsPaginator(client, input)
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
		"list-nodegroups": {
			Name:   "list-nodegroups",
			Fields: fields_list_nodegroups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNodegroupsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_nodegroups, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNodegroups(ctx, input)
				}
				var results []*svc.ListNodegroupsOutput
				p := svc.NewListNodegroupsPaginator(client, input)
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
		"list-pod-identity-associations": {
			Name:   "list-pod-identity-associations",
			Fields: fields_list_pod_identity_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPodIdentityAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pod_identity_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPodIdentityAssociations(ctx, input)
				}
				var results []*svc.ListPodIdentityAssociationsOutput
				p := svc.NewListPodIdentityAssociationsPaginator(client, input)
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
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"list-updates": {
			Name:   "list-updates",
			Fields: fields_list_updates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUpdatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_updates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUpdates(ctx, input)
				}
				var results []*svc.ListUpdatesOutput
				p := svc.NewListUpdatesPaginator(client, input)
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
		"register-cluster": {
			Name:   "register-cluster",
			Fields: fields_register_cluster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterClusterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_cluster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterCluster(ctx, input)
			},
		},
		"start-insights-refresh": {
			Name:   "start-insights-refresh",
			Fields: fields_start_insights_refresh,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartInsightsRefreshInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_insights_refresh, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartInsightsRefresh(ctx, input)
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
		"update-access-entry": {
			Name:   "update-access-entry",
			Fields: fields_update_access_entry,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccessEntryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_access_entry, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccessEntry(ctx, input)
			},
		},
		"update-addon": {
			Name:   "update-addon",
			Fields: fields_update_addon,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAddonInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_addon, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAddon(ctx, input)
			},
		},
		"update-capability": {
			Name:   "update-capability",
			Fields: fields_update_capability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCapabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_capability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCapability(ctx, input)
			},
		},
		"update-cluster-config": {
			Name:   "update-cluster-config",
			Fields: fields_update_cluster_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClusterConfig(ctx, input)
			},
		},
		"update-cluster-version": {
			Name:   "update-cluster-version",
			Fields: fields_update_cluster_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateClusterVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_cluster_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateClusterVersion(ctx, input)
			},
		},
		"update-eks-anywhere-subscription": {
			Name:   "update-eks-anywhere-subscription",
			Fields: fields_update_eks_anywhere_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEksAnywhereSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_eks_anywhere_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEksAnywhereSubscription(ctx, input)
			},
		},
		"update-nodegroup-config": {
			Name:   "update-nodegroup-config",
			Fields: fields_update_nodegroup_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNodegroupConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_nodegroup_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNodegroupConfig(ctx, input)
			},
		},
		"update-nodegroup-version": {
			Name:   "update-nodegroup-version",
			Fields: fields_update_nodegroup_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNodegroupVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_nodegroup_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNodegroupVersion(ctx, input)
			},
		},
		"update-pod-identity-association": {
			Name:   "update-pod-identity-association",
			Fields: fields_update_pod_identity_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePodIdentityAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pod_identity_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePodIdentityAssociation(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("eks", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
