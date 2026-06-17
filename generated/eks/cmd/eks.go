package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// eksCmd represents the eks command
var _eksCmd = &cobra.Command{
	Use:   "eks",
	Short: "AWS eks CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := eks.NewFromConfig(cfg)
		if _eksAssociateAccessPolicy {
			eks_AssociateAccessPolicy(cfg, client)
			return
		}
		if _eksAssociateEncryptionConfig {
			eks_AssociateEncryptionConfig(cfg, client)
			return
		}
		if _eksAssociateIdentityProviderConfig {
			eks_AssociateIdentityProviderConfig(cfg, client)
			return
		}
		if _eksCreateAccessEntry {
			eks_CreateAccessEntry(cfg, client)
			return
		}
		if _eksCreateAddon {
			eks_CreateAddon(cfg, client)
			return
		}
		if _eksCreateCapability {
			eks_CreateCapability(cfg, client)
			return
		}
		if _eksCreateCluster {
			eks_CreateCluster(cfg, client)
			return
		}
		if _eksCreateEksAnywhereSubscription {
			eks_CreateEksAnywhereSubscription(cfg, client)
			return
		}
		if _eksCreateFargateProfile {
			eks_CreateFargateProfile(cfg, client)
			return
		}
		if _eksCreateNodegroup {
			eks_CreateNodegroup(cfg, client)
			return
		}
		if _eksCreatePodIdentityAssociation {
			eks_CreatePodIdentityAssociation(cfg, client)
			return
		}
		if _eksDeleteAccessEntry {
			eks_DeleteAccessEntry(cfg, client)
			return
		}
		if _eksDeleteAddon {
			eks_DeleteAddon(cfg, client)
			return
		}
		if _eksDeleteCapability {
			eks_DeleteCapability(cfg, client)
			return
		}
		if _eksDeleteCluster {
			eks_DeleteCluster(cfg, client)
			return
		}
		if _eksDeleteEksAnywhereSubscription {
			eks_DeleteEksAnywhereSubscription(cfg, client)
			return
		}
		if _eksDeleteFargateProfile {
			eks_DeleteFargateProfile(cfg, client)
			return
		}
		if _eksDeleteNodegroup {
			eks_DeleteNodegroup(cfg, client)
			return
		}
		if _eksDeletePodIdentityAssociation {
			eks_DeletePodIdentityAssociation(cfg, client)
			return
		}
		if _eksDeregisterCluster {
			eks_DeregisterCluster(cfg, client)
			return
		}
		if _eksDescribeAccessEntry {
			eks_DescribeAccessEntry(cfg, client)
			return
		}
		if _eksDescribeAddon {
			eks_DescribeAddon(cfg, client)
			return
		}
		if _eksDescribeAddonConfiguration {
			eks_DescribeAddonConfiguration(cfg, client)
			return
		}
		if _eksDescribeAddonVersions {
			eks_DescribeAddonVersions(cfg, client)
			return
		}
		if _eksDescribeCapability {
			eks_DescribeCapability(cfg, client)
			return
		}
		if _eksDescribeCluster {
			eks_DescribeCluster(cfg, client)
			return
		}
		if _eksDescribeClusterVersions {
			eks_DescribeClusterVersions(cfg, client)
			return
		}
		if _eksDescribeEksAnywhereSubscription {
			eks_DescribeEksAnywhereSubscription(cfg, client)
			return
		}
		if _eksDescribeFargateProfile {
			eks_DescribeFargateProfile(cfg, client)
			return
		}
		if _eksDescribeIdentityProviderConfig {
			eks_DescribeIdentityProviderConfig(cfg, client)
			return
		}
		if _eksDescribeInsight {
			eks_DescribeInsight(cfg, client)
			return
		}
		if _eksDescribeInsightsRefresh {
			eks_DescribeInsightsRefresh(cfg, client)
			return
		}
		if _eksDescribeNodegroup {
			eks_DescribeNodegroup(cfg, client)
			return
		}
		if _eksDescribePodIdentityAssociation {
			eks_DescribePodIdentityAssociation(cfg, client)
			return
		}
		if _eksDescribeUpdate {
			eks_DescribeUpdate(cfg, client)
			return
		}
		if _eksDisassociateAccessPolicy {
			eks_DisassociateAccessPolicy(cfg, client)
			return
		}
		if _eksDisassociateIdentityProviderConfig {
			eks_DisassociateIdentityProviderConfig(cfg, client)
			return
		}
		if _eksListAccessEntries {
			eks_ListAccessEntries(cfg, client)
			return
		}
		if _eksListAccessPolicies {
			eks_ListAccessPolicies(cfg, client)
			return
		}
		if _eksListAddons {
			eks_ListAddons(cfg, client)
			return
		}
		if _eksListAssociatedAccessPolicies {
			eks_ListAssociatedAccessPolicies(cfg, client)
			return
		}
		if _eksListCapabilities {
			eks_ListCapabilities(cfg, client)
			return
		}
		if _eksListClusters {
			eks_ListClusters(cfg, client)
			return
		}
		if _eksListEksAnywhereSubscriptions {
			eks_ListEksAnywhereSubscriptions(cfg, client)
			return
		}
		if _eksListFargateProfiles {
			eks_ListFargateProfiles(cfg, client)
			return
		}
		if _eksListIdentityProviderConfigs {
			eks_ListIdentityProviderConfigs(cfg, client)
			return
		}
		if _eksListInsights {
			eks_ListInsights(cfg, client)
			return
		}
		if _eksListNodegroups {
			eks_ListNodegroups(cfg, client)
			return
		}
		if _eksListPodIdentityAssociations {
			eks_ListPodIdentityAssociations(cfg, client)
			return
		}
		if _eksListTagsForResource {
			eks_ListTagsForResource(cfg, client)
			return
		}
		if _eksListUpdates {
			eks_ListUpdates(cfg, client)
			return
		}
		if _eksRegisterCluster {
			eks_RegisterCluster(cfg, client)
			return
		}
		if _eksStartInsightsRefresh {
			eks_StartInsightsRefresh(cfg, client)
			return
		}
		if _eksTagResource {
			eks_TagResource(cfg, client)
			return
		}
		if _eksUntagResource {
			eks_UntagResource(cfg, client)
			return
		}
		if _eksUpdateAccessEntry {
			eks_UpdateAccessEntry(cfg, client)
			return
		}
		if _eksUpdateAddon {
			eks_UpdateAddon(cfg, client)
			return
		}
		if _eksUpdateCapability {
			eks_UpdateCapability(cfg, client)
			return
		}
		if _eksUpdateClusterConfig {
			eks_UpdateClusterConfig(cfg, client)
			return
		}
		if _eksUpdateClusterVersion {
			eks_UpdateClusterVersion(cfg, client)
			return
		}
		if _eksUpdateEksAnywhereSubscription {
			eks_UpdateEksAnywhereSubscription(cfg, client)
			return
		}
		if _eksUpdateNodegroupConfig {
			eks_UpdateNodegroupConfig(cfg, client)
			return
		}
		if _eksUpdateNodegroupVersion {
			eks_UpdateNodegroupVersion(cfg, client)
			return
		}
		if _eksUpdatePodIdentityAssociation {
			eks_UpdatePodIdentityAssociation(cfg, client)
			return
		}

	},
}

var (
	_eksAssociateAccessPolicy              bool
	_eksAssociateEncryptionConfig          bool
	_eksAssociateIdentityProviderConfig    bool
	_eksCreateAccessEntry                  bool
	_eksCreateAddon                        bool
	_eksCreateCapability                   bool
	_eksCreateCluster                      bool
	_eksCreateEksAnywhereSubscription      bool
	_eksCreateFargateProfile               bool
	_eksCreateNodegroup                    bool
	_eksCreatePodIdentityAssociation       bool
	_eksDeleteAccessEntry                  bool
	_eksDeleteAddon                        bool
	_eksDeleteCapability                   bool
	_eksDeleteCluster                      bool
	_eksDeleteEksAnywhereSubscription      bool
	_eksDeleteFargateProfile               bool
	_eksDeleteNodegroup                    bool
	_eksDeletePodIdentityAssociation       bool
	_eksDeregisterCluster                  bool
	_eksDescribeAccessEntry                bool
	_eksDescribeAddon                      bool
	_eksDescribeAddonConfiguration         bool
	_eksDescribeAddonVersions              bool
	_eksDescribeCapability                 bool
	_eksDescribeCluster                    bool
	_eksDescribeClusterVersions            bool
	_eksDescribeEksAnywhereSubscription    bool
	_eksDescribeFargateProfile             bool
	_eksDescribeIdentityProviderConfig     bool
	_eksDescribeInsight                    bool
	_eksDescribeInsightsRefresh            bool
	_eksDescribeNodegroup                  bool
	_eksDescribePodIdentityAssociation     bool
	_eksDescribeUpdate                     bool
	_eksDisassociateAccessPolicy           bool
	_eksDisassociateIdentityProviderConfig bool
	_eksListAccessEntries                  bool
	_eksListAccessPolicies                 bool
	_eksListAddons                         bool
	_eksListAssociatedAccessPolicies       bool
	_eksListCapabilities                   bool
	_eksListClusters                       bool
	_eksListEksAnywhereSubscriptions       bool
	_eksListFargateProfiles                bool
	_eksListIdentityProviderConfigs        bool
	_eksListInsights                       bool
	_eksListNodegroups                     bool
	_eksListPodIdentityAssociations        bool
	_eksListTagsForResource                bool
	_eksListUpdates                        bool
	_eksRegisterCluster                    bool
	_eksStartInsightsRefresh               bool
	_eksTagResource                        bool
	_eksUntagResource                      bool
	_eksUpdateAccessEntry                  bool
	_eksUpdateAddon                        bool
	_eksUpdateCapability                   bool
	_eksUpdateClusterConfig                bool
	_eksUpdateClusterVersion               bool
	_eksUpdateEksAnywhereSubscription      bool
	_eksUpdateNodegroupConfig              bool
	_eksUpdateNodegroupVersion             bool
	_eksUpdatePodIdentityAssociation       bool

	_eksAccessConfig               string
	_eksAccessScope                string
	_eksAddonName                  string
	_eksAddonVersion               string
	_eksAmiType                    string
	_eksAssociatedPolicyArn        string
	_eksAssociationId              string
	_eksAutoRenew                  string
	_eksBootstrapSelfManagedAddons string
	_eksCapabilityName             string
	_eksCapacityType               string
	_eksClientRequestToken         string
	_eksClusterName                string
	_eksClusterType                string
	_eksClusterVersions            []string
	_eksComputeConfig              string
	_eksConfiguration              string
	_eksConfigurationValues        string
	_eksConnectorConfig            string
	_eksControlPlaneScalingConfig  string
	_eksDefaultOnly                string
	_eksDeletePropagationPolicy    string
	_eksDeletionProtection         string
	_eksDisableSessionTags         string
	_eksDiskSize                   string
	_eksEncryptionConfig           string
	_eksFargateProfileName         string
	_eksFilter                     string
	_eksForce                      string
	_eksId                         string
	_eksIdentityProviderConfig     string
	_eksInclude                    []string
	_eksIncludeAll                 string
	_eksIncludeStatus              string
	_eksInstanceTypes              []string
	_eksKubernetesGroups           []string
	_eksKubernetesNetworkConfig    string
	_eksKubernetesVersion          string
	_eksLabels                     string
	_eksLaunchTemplate             string
	_eksLicenseQuantity            string
	_eksLicenseType                string
	_eksLogging                    string
	_eksMaxResults                 string
	_eksName                       string
	_eksNamespace                  string
	_eksNamespaceConfig            string
	_eksNextToken                  string
	_eksNodeRepairConfig           string
	_eksNodeRole                   string
	_eksNodegroupName              string
	_eksOidc                       string
	_eksOutpostConfig              string
	_eksOwners                     []string
	_eksPodExecutionRoleArn        string
	_eksPodIdentityAssociations    string
	_eksPolicy                     string
	_eksPolicyArn                  string
	_eksPreserve                   string
	_eksPrincipalArn               string
	_eksPublishers                 []string
	_eksReleaseVersion             string
	_eksRemoteAccess               string
	_eksRemoteNetworkConfig        string
	_eksResolveConflicts           string
	_eksResourceArn                string
	_eksResourcesVpcConfig         string
	_eksRoleArn                    string
	_eksScalingConfig              string
	_eksSelectors                  string
	_eksServiceAccount             string
	_eksServiceAccountRoleArn      string
	_eksStatus                     string
	_eksStorageConfig              string
	_eksSubnets                    []string
	_eksTagKeys                    []string
	_eksTags                       string
	_eksTaints                     string
	_eksTargetRoleArn              string
	_eksTerm                       string
	_eksType                       string
	_eksTypes                      []string
	_eksUpdateConfig               string
	_eksUpdateId                   string
	_eksUpgradePolicy              string
	_eksUsername                   string
	_eksVersion                    string
	_eksVersionStatus              string
	_eksZonalShiftConfig           string
)

// Associates an access policy and its scope to an access entry. For more
// information about associating access policies, see [Associating and disassociating access policies to and from access entries]in the Amazon EKS User Guide.
//
// [Associating and disassociating access policies to and from access entries]: https://docs.aws.amazon.com/eks/latest/userguide/access-policies.html
func eks_AssociateAccessPolicy(cfg aws.Config, client *eks.Client) {
	input := &eks.AssociateAccessPolicyInput{
		// AccessScope: *types.AccessScope, // Required
		// ClusterName: *string, // Required
		// PolicyArn: *string, // Required
		// PrincipalArn: *string, // Required
	}

	if len(_eksAccessScope) > 0 {
		if err := assignInputField(input, "AccessScope", _eksAccessScope); err != nil {
			log.Errorf("invalid --access-scope: %s", err.Error())
			return
		}
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksPolicyArn) > 0 {
		input.PolicyArn = aws.String(_eksPolicyArn)
	}
	if len(_eksPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_eksPrincipalArn)
	}

	if resp, err := client.AssociateAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an encryption configuration to an existing cluster.
// Use this API to enable encryption on existing clusters that don't already have
// encryption enabled. This allows you to implement a defense-in-depth security
// strategy without migrating applications to new Amazon EKS clusters.
func eks_AssociateEncryptionConfig(cfg aws.Config, client *eks.Client) {
	input := &eks.AssociateEncryptionConfigInput{
		// ClusterName: *string, // Required
		// EncryptionConfig: []types.EncryptionConfig, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksEncryptionConfig) > 0 {
		if err := assignInputField(input, "EncryptionConfig", _eksEncryptionConfig); err != nil {
			log.Errorf("invalid --encryption-config: %s", err.Error())
			return
		}
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}

	if resp, err := client.AssociateEncryptionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an identity provider configuration to a cluster.
// If you want to authenticate identities using an identity provider, you can
// create an identity provider configuration and associate it to your cluster.
// After configuring authentication to your cluster you can create Kubernetes Role
// and ClusterRole objects, assign permissions to them, and then bind them to the
// identities using Kubernetes RoleBinding and ClusterRoleBinding objects. For
// more information see [Using RBAC Authorization]in the Kubernetes documentation.
//
// [Using RBAC Authorization]: https://kubernetes.io/docs/reference/access-authn-authz/rbac/
func eks_AssociateIdentityProviderConfig(cfg aws.Config, client *eks.Client) {
	input := &eks.AssociateIdentityProviderConfigInput{
		// ClusterName: *string, // Required
		// Oidc: *types.OidcIdentityProviderConfigRequest, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksOidc) > 0 {
		if err := assignInputField(input, "Oidc", _eksOidc); err != nil {
			log.Errorf("invalid --oidc: %s", err.Error())
			return
		}
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateIdentityProviderConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an access entry.
// An access entry allows an IAM principal to access your cluster. Access entries
// can replace the need to maintain entries in the aws-auth ConfigMap for
// authentication. You have the following options for authorizing an IAM principal
// to access Kubernetes objects on your cluster: Kubernetes role-based access
// control (RBAC), Amazon EKS, or both. Kubernetes RBAC authorization requires you
// to create and manage Kubernetes Role , ClusterRole , RoleBinding , and
// ClusterRoleBinding objects, in addition to managing access entries. If you use
// Amazon EKS authorization exclusively, you don't need to create and manage
// Kubernetes Role , ClusterRole , RoleBinding , and ClusterRoleBinding objects.
//
// For more information about access entries, see [Access entries] in the Amazon EKS User Guide.
//
// [Access entries]: https://docs.aws.amazon.com/eks/latest/userguide/access-entries.html
func eks_CreateAccessEntry(cfg aws.Config, client *eks.Client) {
	input := &eks.CreateAccessEntryInput{
		// ClusterName: *string, // Required
		// PrincipalArn: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_eksPrincipalArn)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksKubernetesGroups) > 0 {
		input.KubernetesGroups = append([]string(nil), _eksKubernetesGroups...)
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_eksType) > 0 {
		input.Type = aws.String(_eksType)
	}
	if len(_eksUsername) > 0 {
		input.Username = aws.String(_eksUsername)
	}

	if resp, err := client.CreateAccessEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon EKS add-on.
// Amazon EKS add-ons help to automate the provisioning and lifecycle management
// of common operational software for Amazon EKS clusters. For more information,
// see [Amazon EKS add-ons]in the Amazon EKS User Guide.
//
// [Amazon EKS add-ons]: https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html
func eks_CreateAddon(cfg aws.Config, client *eks.Client) {
	input := &eks.CreateAddonInput{
		// AddonName: *string, // Required
		// ClusterName: *string, // Required
	}

	if len(_eksAddonName) > 0 {
		input.AddonName = aws.String(_eksAddonName)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksAddonVersion) > 0 {
		input.AddonVersion = aws.String(_eksAddonVersion)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksConfigurationValues) > 0 {
		input.ConfigurationValues = aws.String(_eksConfigurationValues)
	}
	if len(_eksNamespaceConfig) > 0 {
		if err := assignInputField(input, "NamespaceConfig", _eksNamespaceConfig); err != nil {
			log.Errorf("invalid --namespace-config: %s", err.Error())
			return
		}
	}
	if len(_eksPodIdentityAssociations) > 0 {
		if err := assignInputField(input, "PodIdentityAssociations", _eksPodIdentityAssociations); err != nil {
			log.Errorf("invalid --pod-identity-associations: %s", err.Error())
			return
		}
	}
	if len(_eksResolveConflicts) > 0 {
		if err := assignInputField(input, "ResolveConflicts", _eksResolveConflicts); err != nil {
			log.Errorf("invalid --resolve-conflicts: %s", err.Error())
			return
		}
	}
	if len(_eksServiceAccountRoleArn) > 0 {
		input.ServiceAccountRoleArn = aws.String(_eksServiceAccountRoleArn)
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAddon(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a managed capability resource for an Amazon EKS cluster.
// Capabilities provide fully managed capabilities to build and scale with
// Kubernetes. When you create a capability, Amazon EKSprovisions and manages the
// infrastructure required to run the capability outside of your cluster. This
// approach reduces operational overhead and preserves cluster resources.
//
// You can only create one Capability of each type on a given Amazon EKS cluster.
// Valid types are Argo CD for declarative GitOps deployment, Amazon Web Services
// Controllers for Kubernetes (ACK) for resource management, and Kube Resource
// Orchestrator (KRO) for Kubernetes custom resource orchestration.
//
// For more information, see [EKS Capabilities] in the Amazon EKS User Guide.
//
// [EKS Capabilities]: https://docs.aws.amazon.com/eks/latest/userguide/capabilities.html
func eks_CreateCapability(cfg aws.Config, client *eks.Client) {
	input := &eks.CreateCapabilityInput{
		// CapabilityName: *string, // Required
		// ClusterName: *string, // Required
		// DeletePropagationPolicy: types.CapabilityDeletePropagationPolicy, // Required
		// RoleArn: *string, // Required
		// Type: types.CapabilityType, // Required
	}

	if len(_eksCapabilityName) > 0 {
		input.CapabilityName = aws.String(_eksCapabilityName)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksDeletePropagationPolicy) > 0 {
		if err := assignInputField(input, "DeletePropagationPolicy", _eksDeletePropagationPolicy); err != nil {
			log.Errorf("invalid --delete-propagation-policy: %s", err.Error())
			return
		}
	}
	if len(_eksRoleArn) > 0 {
		input.RoleArn = aws.String(_eksRoleArn)
	}
	if len(_eksType) > 0 {
		if err := assignInputField(input, "Type", _eksType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _eksConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCapability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon EKS control plane.
// The Amazon EKS control plane consists of control plane instances that run the
// Kubernetes software, such as etcd and the API server. The control plane runs in
// an account managed by Amazon Web Services, and the Kubernetes API is exposed by
// the Amazon EKS API server endpoint. Each Amazon EKS cluster control plane is
// single tenant and unique. It runs on its own set of Amazon EC2 instances.
//
// The cluster control plane is provisioned across multiple Availability Zones and
// fronted by an Elastic Load Balancing Network Load Balancer. Amazon EKS also
// provisions elastic network interfaces in your VPC subnets to provide
// connectivity from the control plane instances to the nodes (for example, to
// support kubectl exec , logs , and proxy data flows).
//
// Amazon EKS nodes run in your Amazon Web Services account and connect to your
// cluster's control plane over the Kubernetes API server endpoint and a
// certificate file that is created for your cluster.
//
// You can use the endpointPublicAccess and endpointPrivateAccess parameters to
// enable or disable public and private access to your cluster's Kubernetes API
// server endpoint. By default, public access is enabled, and private access is
// disabled. The endpoint domain name and IP address family depends on the value of
// the ipFamily for the cluster. For more information, see [Amazon EKS Cluster Endpoint Access Control] in the Amazon EKS User
// Guide .
//
// You can use the logging parameter to enable or disable exporting the Kubernetes
// control plane logs for your cluster to CloudWatch Logs. By default, cluster
// control plane logs aren't exported to CloudWatch Logs. For more information, see
// [Amazon EKS Cluster Control Plane Logs]in the Amazon EKS User Guide .
//
// CloudWatch Logs ingestion, archive storage, and data scanning rates apply to
// exported control plane logs. For more information, see [CloudWatch Pricing].
//
// In most cases, it takes several minutes to create a cluster. After you create
// an Amazon EKS cluster, you must configure your Kubernetes tooling to communicate
// with the API server and launch nodes into your cluster. For more information,
// see [Allowing users to access your cluster]and [Launching Amazon EKS nodes] in the Amazon EKS User Guide.
//
// [Allowing users to access your cluster]: https://docs.aws.amazon.com/eks/latest/userguide/cluster-auth.html
// [CloudWatch Pricing]: http://aws.amazon.com/cloudwatch/pricing/
// [Amazon EKS Cluster Control Plane Logs]: https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html
// [Amazon EKS Cluster Endpoint Access Control]: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
// [Launching Amazon EKS nodes]: https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html
func eks_CreateCluster(cfg aws.Config, client *eks.Client) {
	input := &eks.CreateClusterInput{
		// Name: *string, // Required
		// ResourcesVpcConfig: *types.VpcConfigRequest, // Required
		// RoleArn: *string, // Required
	}

	if len(_eksName) > 0 {
		input.Name = aws.String(_eksName)
	}
	if len(_eksResourcesVpcConfig) > 0 {
		if err := assignInputField(input, "ResourcesVpcConfig", _eksResourcesVpcConfig); err != nil {
			log.Errorf("invalid --resources-vpc-config: %s", err.Error())
			return
		}
	}
	if len(_eksRoleArn) > 0 {
		input.RoleArn = aws.String(_eksRoleArn)
	}
	if len(_eksAccessConfig) > 0 {
		if err := assignInputField(input, "AccessConfig", _eksAccessConfig); err != nil {
			log.Errorf("invalid --access-config: %s", err.Error())
			return
		}
	}
	if len(_eksBootstrapSelfManagedAddons) > 0 {
		if err := assignInputField(input, "BootstrapSelfManagedAddons", _eksBootstrapSelfManagedAddons); err != nil {
			log.Errorf("invalid --bootstrap-self-managed-addons: %s", err.Error())
			return
		}
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksComputeConfig) > 0 {
		if err := assignInputField(input, "ComputeConfig", _eksComputeConfig); err != nil {
			log.Errorf("invalid --compute-config: %s", err.Error())
			return
		}
	}
	if len(_eksControlPlaneScalingConfig) > 0 {
		if err := assignInputField(input, "ControlPlaneScalingConfig", _eksControlPlaneScalingConfig); err != nil {
			log.Errorf("invalid --control-plane-scaling-config: %s", err.Error())
			return
		}
	}
	if len(_eksDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _eksDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_eksEncryptionConfig) > 0 {
		if err := assignInputField(input, "EncryptionConfig", _eksEncryptionConfig); err != nil {
			log.Errorf("invalid --encryption-config: %s", err.Error())
			return
		}
	}
	if len(_eksKubernetesNetworkConfig) > 0 {
		if err := assignInputField(input, "KubernetesNetworkConfig", _eksKubernetesNetworkConfig); err != nil {
			log.Errorf("invalid --kubernetes-network-config: %s", err.Error())
			return
		}
	}
	if len(_eksLogging) > 0 {
		if err := assignInputField(input, "Logging", _eksLogging); err != nil {
			log.Errorf("invalid --logging: %s", err.Error())
			return
		}
	}
	if len(_eksOutpostConfig) > 0 {
		if err := assignInputField(input, "OutpostConfig", _eksOutpostConfig); err != nil {
			log.Errorf("invalid --outpost-config: %s", err.Error())
			return
		}
	}
	if len(_eksRemoteNetworkConfig) > 0 {
		if err := assignInputField(input, "RemoteNetworkConfig", _eksRemoteNetworkConfig); err != nil {
			log.Errorf("invalid --remote-network-config: %s", err.Error())
			return
		}
	}
	if len(_eksStorageConfig) > 0 {
		if err := assignInputField(input, "StorageConfig", _eksStorageConfig); err != nil {
			log.Errorf("invalid --storage-config: %s", err.Error())
			return
		}
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_eksUpgradePolicy) > 0 {
		if err := assignInputField(input, "UpgradePolicy", _eksUpgradePolicy); err != nil {
			log.Errorf("invalid --upgrade-policy: %s", err.Error())
			return
		}
	}
	if len(_eksVersion) > 0 {
		input.Version = aws.String(_eksVersion)
	}
	if len(_eksZonalShiftConfig) > 0 {
		if err := assignInputField(input, "ZonalShiftConfig", _eksZonalShiftConfig); err != nil {
			log.Errorf("invalid --zonal-shift-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an EKS Anywhere subscription. When a subscription is created, it is a
// contract agreement for the length of the term specified in the request. Licenses
// that are used to validate support are provisioned in Amazon Web Services License
// Manager and the caller account is granted access to EKS Anywhere Curated
// Packages.
func eks_CreateEksAnywhereSubscription(cfg aws.Config, client *eks.Client) {
	input := &eks.CreateEksAnywhereSubscriptionInput{
		// Name: *string, // Required
		// Term: *types.EksAnywhereSubscriptionTerm, // Required
	}

	if len(_eksName) > 0 {
		input.Name = aws.String(_eksName)
	}
	if len(_eksTerm) > 0 {
		if err := assignInputField(input, "Term", _eksTerm); err != nil {
			log.Errorf("invalid --term: %s", err.Error())
			return
		}
	}
	if len(_eksAutoRenew) > 0 {
		if err := assignInputField(input, "AutoRenew", _eksAutoRenew); err != nil {
			log.Errorf("invalid --auto-renew: %s", err.Error())
			return
		}
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksLicenseQuantity) > 0 {
		if err := assignInputField(input, "LicenseQuantity", _eksLicenseQuantity); err != nil {
			log.Errorf("invalid --license-quantity: %s", err.Error())
			return
		}
	}
	if len(_eksLicenseType) > 0 {
		if err := assignInputField(input, "LicenseType", _eksLicenseType); err != nil {
			log.Errorf("invalid --license-type: %s", err.Error())
			return
		}
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEksAnywhereSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Fargate profile for your Amazon EKS cluster. You must have at least
// one Fargate profile in a cluster to be able to run pods on Fargate.
//
// The Fargate profile allows an administrator to declare which pods run on
// Fargate and specify which pods run on which Fargate profile. This declaration is
// done through the profile's selectors. Each profile can have up to five selectors
// that contain a namespace and labels. A namespace is required for every selector.
// The label field consists of multiple optional key-value pairs. Pods that match
// the selectors are scheduled on Fargate. If a to-be-scheduled pod matches any of
// the selectors in the Fargate profile, then that pod is run on Fargate.
//
// When you create a Fargate profile, you must specify a pod execution role to use
// with the pods that are scheduled with the profile. This role is added to the
// cluster's Kubernetes [Role Based Access Control](RBAC) for authorization so that the kubelet that is
// running on the Fargate infrastructure can register with your Amazon EKS cluster
// so that it can appear in your cluster as a node. The pod execution role also
// provides IAM permissions to the Fargate infrastructure to allow read access to
// Amazon ECR image repositories. For more information, see [Pod Execution Role]in the Amazon EKS User
// Guide.
//
// Fargate profiles are immutable. However, you can create a new updated profile
// to replace an existing profile and then delete the original after the updated
// profile has finished creating.
//
// If any Fargate profiles in a cluster are in the DELETING status, you must wait
// for that Fargate profile to finish deleting before you can create any other
// profiles in that cluster.
//
// For more information, see [Fargate profile] in the Amazon EKS User Guide.
//
// [Role Based Access Control]: https://kubernetes.io/docs/reference/access-authn-authz/rbac/
// [Fargate profile]: https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html
// [Pod Execution Role]: https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html
func eks_CreateFargateProfile(cfg aws.Config, client *eks.Client) {
	input := &eks.CreateFargateProfileInput{
		// ClusterName: *string, // Required
		// FargateProfileName: *string, // Required
		// PodExecutionRoleArn: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksFargateProfileName) > 0 {
		input.FargateProfileName = aws.String(_eksFargateProfileName)
	}
	if len(_eksPodExecutionRoleArn) > 0 {
		input.PodExecutionRoleArn = aws.String(_eksPodExecutionRoleArn)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksSelectors) > 0 {
		if err := assignInputField(input, "Selectors", _eksSelectors); err != nil {
			log.Errorf("invalid --selectors: %s", err.Error())
			return
		}
	}
	if len(_eksSubnets) > 0 {
		input.Subnets = append([]string(nil), _eksSubnets...)
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFargateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a managed node group for an Amazon EKS cluster.
// You can only create a node group for your cluster that is equal to the current
// Kubernetes version for the cluster. All node groups are created with the latest
// AMI release version for the respective minor Kubernetes version of the cluster,
// unless you deploy a custom AMI using a launch template.
//
// For later updates, you will only be able to update a node group using a launch
// template only if it was originally deployed with a launch template.
// Additionally, the launch template ID or name must match what was used when the
// node group was created. You can update the launch template version with
// necessary changes. For more information about using launch templates, see [Customizing managed nodes with launch templates].
//
// An Amazon EKS managed node group is an Amazon EC2 Auto Scaling group and
// associated Amazon EC2 instances that are managed by Amazon Web Services for an
// Amazon EKS cluster. For more information, see [Managed node groups]in the Amazon EKS User Guide.
//
// Windows AMI types are only supported for commercial Amazon Web Services Regions
// that support Windows on Amazon EKS.
//
// [Customizing managed nodes with launch templates]: https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html
// [Managed node groups]: https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html
func eks_CreateNodegroup(cfg aws.Config, client *eks.Client) {
	input := &eks.CreateNodegroupInput{
		// ClusterName: *string, // Required
		// NodeRole: *string, // Required
		// NodegroupName: *string, // Required
		// Subnets: []string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksNodeRole) > 0 {
		input.NodeRole = aws.String(_eksNodeRole)
	}
	if len(_eksNodegroupName) > 0 {
		input.NodegroupName = aws.String(_eksNodegroupName)
	}
	if len(_eksSubnets) > 0 {
		input.Subnets = append([]string(nil), _eksSubnets...)
	}
	if len(_eksAmiType) > 0 {
		if err := assignInputField(input, "AmiType", _eksAmiType); err != nil {
			log.Errorf("invalid --ami-type: %s", err.Error())
			return
		}
	}
	if len(_eksCapacityType) > 0 {
		if err := assignInputField(input, "CapacityType", _eksCapacityType); err != nil {
			log.Errorf("invalid --capacity-type: %s", err.Error())
			return
		}
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksDiskSize) > 0 {
		if err := assignInputField(input, "DiskSize", _eksDiskSize); err != nil {
			log.Errorf("invalid --disk-size: %s", err.Error())
			return
		}
	}
	if len(_eksInstanceTypes) > 0 {
		input.InstanceTypes = append([]string(nil), _eksInstanceTypes...)
	}
	if len(_eksLabels) > 0 {
		if err := assignInputField(input, "Labels", _eksLabels); err != nil {
			log.Errorf("invalid --labels: %s", err.Error())
			return
		}
	}
	if len(_eksLaunchTemplate) > 0 {
		if err := assignInputField(input, "LaunchTemplate", _eksLaunchTemplate); err != nil {
			log.Errorf("invalid --launch-template: %s", err.Error())
			return
		}
	}
	if len(_eksNodeRepairConfig) > 0 {
		if err := assignInputField(input, "NodeRepairConfig", _eksNodeRepairConfig); err != nil {
			log.Errorf("invalid --node-repair-config: %s", err.Error())
			return
		}
	}
	if len(_eksReleaseVersion) > 0 {
		input.ReleaseVersion = aws.String(_eksReleaseVersion)
	}
	if len(_eksRemoteAccess) > 0 {
		if err := assignInputField(input, "RemoteAccess", _eksRemoteAccess); err != nil {
			log.Errorf("invalid --remote-access: %s", err.Error())
			return
		}
	}
	if len(_eksScalingConfig) > 0 {
		if err := assignInputField(input, "ScalingConfig", _eksScalingConfig); err != nil {
			log.Errorf("invalid --scaling-config: %s", err.Error())
			return
		}
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_eksTaints) > 0 {
		if err := assignInputField(input, "Taints", _eksTaints); err != nil {
			log.Errorf("invalid --taints: %s", err.Error())
			return
		}
	}
	if len(_eksUpdateConfig) > 0 {
		if err := assignInputField(input, "UpdateConfig", _eksUpdateConfig); err != nil {
			log.Errorf("invalid --update-config: %s", err.Error())
			return
		}
	}
	if len(_eksVersion) > 0 {
		input.Version = aws.String(_eksVersion)
	}

	if resp, err := client.CreateNodegroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an EKS Pod Identity association between a service account in an Amazon
// EKS cluster and an IAM role with EKS Pod Identity. Use EKS Pod Identity to give
// temporary IAM credentials to Pods and the credentials are rotated automatically.
//
// Amazon EKS Pod Identity associations provide the ability to manage credentials
// for your applications, similar to the way that Amazon EC2 instance profiles
// provide credentials to Amazon EC2 instances.
//
// If a Pod uses a service account that has an association, Amazon EKS sets
// environment variables in the containers of the Pod. The environment variables
// configure the Amazon Web Services SDKs, including the Command Line Interface, to
// use the EKS Pod Identity credentials.
//
// EKS Pod Identity is a simpler method than IAM roles for service accounts, as
// this method doesn't use OIDC identity providers. Additionally, you can configure
// a role for EKS Pod Identity once, and reuse it across clusters.
//
// Similar to Amazon Web Services IAM behavior, EKS Pod Identity associations are
// eventually consistent, and may take several seconds to be effective after the
// initial API call returns successfully. You must design your applications to
// account for these potential delays. We recommend that you don’t include
// association create/updates in the critical, high-availability code paths of your
// application. Instead, make changes in a separate initialization or setup routine
// that you run less frequently.
//
// You can set a target IAM role in the same or a different account for advanced
// scenarios. With a target role, EKS Pod Identity automatically performs two role
// assumptions in sequence: first assuming the role in the association that is in
// this account, then using those credentials to assume the target IAM role. This
// process provides your Pod with temporary credentials that have the permissions
// defined in the target role, allowing secure access to resources in another
// Amazon Web Services account.
func eks_CreatePodIdentityAssociation(cfg aws.Config, client *eks.Client) {
	input := &eks.CreatePodIdentityAssociationInput{
		// ClusterName: *string, // Required
		// Namespace: *string, // Required
		// RoleArn: *string, // Required
		// ServiceAccount: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksNamespace) > 0 {
		input.Namespace = aws.String(_eksNamespace)
	}
	if len(_eksRoleArn) > 0 {
		input.RoleArn = aws.String(_eksRoleArn)
	}
	if len(_eksServiceAccount) > 0 {
		input.ServiceAccount = aws.String(_eksServiceAccount)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksDisableSessionTags) > 0 {
		if err := assignInputField(input, "DisableSessionTags", _eksDisableSessionTags); err != nil {
			log.Errorf("invalid --disable-session-tags: %s", err.Error())
			return
		}
	}
	if len(_eksPolicy) > 0 {
		input.Policy = aws.String(_eksPolicy)
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_eksTargetRoleArn) > 0 {
		input.TargetRoleArn = aws.String(_eksTargetRoleArn)
	}

	if resp, err := client.CreatePodIdentityAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an access entry.
// Deleting an access entry of a type other than Standard can cause your cluster
// to function improperly. If you delete an access entry in error, you can recreate
// it.
func eks_DeleteAccessEntry(cfg aws.Config, client *eks.Client) {
	input := &eks.DeleteAccessEntryInput{
		// ClusterName: *string, // Required
		// PrincipalArn: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_eksPrincipalArn)
	}

	if resp, err := client.DeleteAccessEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon EKS add-on.
// When you remove an add-on, it's deleted from the cluster. You can always
// manually start an add-on on the cluster using the Kubernetes API.
func eks_DeleteAddon(cfg aws.Config, client *eks.Client) {
	input := &eks.DeleteAddonInput{
		// AddonName: *string, // Required
		// ClusterName: *string, // Required
	}

	if len(_eksAddonName) > 0 {
		input.AddonName = aws.String(_eksAddonName)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksPreserve) > 0 {
		if err := assignInputField(input, "Preserve", _eksPreserve); err != nil {
			log.Errorf("invalid --preserve: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAddon(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a managed capability from your Amazon EKS cluster. When you delete a
// capability, Amazon EKS removes the capability infrastructure but retains all
// resources that were managed by the capability.
//
// Before deleting a capability, you should delete all Kubernetes resources that
// were created by the capability. After the capability is deleted, these resources
// become difficult to manage because the controller that managed them is no longer
// available. To delete resources before removing the capability, use kubectl
// delete or remove them through your GitOps workflow.
func eks_DeleteCapability(cfg aws.Config, client *eks.Client) {
	input := &eks.DeleteCapabilityInput{
		// CapabilityName: *string, // Required
		// ClusterName: *string, // Required
	}

	if len(_eksCapabilityName) > 0 {
		input.CapabilityName = aws.String(_eksCapabilityName)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}

	if resp, err := client.DeleteCapability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon EKS cluster control plane.
// If you have active services and ingress resources in your cluster that are
// associated with a load balancer, you must delete those services before deleting
// the cluster so that the load balancers are deleted properly. Otherwise, you can
// have orphaned resources in your VPC that prevent you from being able to delete
// the VPC. For more information, see [Deleting a cluster]in the Amazon EKS User Guide.
//
// If you have managed node groups or Fargate profiles attached to the cluster,
// you must delete them first. For more information, see DeleteNodgroup and
// DeleteFargateProfile .
//
// [Deleting a cluster]: https://docs.aws.amazon.com/eks/latest/userguide/delete-cluster.html
func eks_DeleteCluster(cfg aws.Config, client *eks.Client) {
	input := &eks.DeleteClusterInput{
		// Name: *string, // Required
	}

	if len(_eksName) > 0 {
		input.Name = aws.String(_eksName)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an expired or inactive subscription. Deleting inactive subscriptions
// removes them from the Amazon Web Services Management Console view and from
// list/describe API responses. Subscriptions can only be cancelled within 7 days
// of creation and are cancelled by creating a ticket in the Amazon Web Services
// Support Center.
func eks_DeleteEksAnywhereSubscription(cfg aws.Config, client *eks.Client) {
	input := &eks.DeleteEksAnywhereSubscriptionInput{
		// Id: *string, // Required
	}

	if len(_eksId) > 0 {
		input.Id = aws.String(_eksId)
	}

	if resp, err := client.DeleteEksAnywhereSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Fargate profile.
// When you delete a Fargate profile, any Pod running on Fargate that was created
// with the profile is deleted. If the Pod matches another Fargate profile, then
// it is scheduled on Fargate with that profile. If it no longer matches any
// Fargate profiles, then it's not scheduled on Fargate and may remain in a pending
// state.
//
// Only one Fargate profile in a cluster can be in the DELETING status at a time.
// You must wait for a Fargate profile to finish deleting before you can delete any
// other profiles in that cluster.
func eks_DeleteFargateProfile(cfg aws.Config, client *eks.Client) {
	input := &eks.DeleteFargateProfileInput{
		// ClusterName: *string, // Required
		// FargateProfileName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksFargateProfileName) > 0 {
		input.FargateProfileName = aws.String(_eksFargateProfileName)
	}

	if resp, err := client.DeleteFargateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a managed node group.
func eks_DeleteNodegroup(cfg aws.Config, client *eks.Client) {
	input := &eks.DeleteNodegroupInput{
		// ClusterName: *string, // Required
		// NodegroupName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksNodegroupName) > 0 {
		input.NodegroupName = aws.String(_eksNodegroupName)
	}

	if resp, err := client.DeleteNodegroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a EKS Pod Identity association.
// The temporary Amazon Web Services credentials from the previous IAM role
// session might still be valid until the session expiry. If you need to
// immediately revoke the temporary session credentials, then go to the role in the
// IAM console.
func eks_DeletePodIdentityAssociation(cfg aws.Config, client *eks.Client) {
	input := &eks.DeletePodIdentityAssociationInput{
		// AssociationId: *string, // Required
		// ClusterName: *string, // Required
	}

	if len(_eksAssociationId) > 0 {
		input.AssociationId = aws.String(_eksAssociationId)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}

	if resp, err := client.DeletePodIdentityAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters a connected cluster to remove it from the Amazon EKS control plane.
// A connected cluster is a Kubernetes cluster that you've connected to your
// control plane using the [Amazon EKS Connector].
//
// [Amazon EKS Connector]: https://docs.aws.amazon.com/eks/latest/userguide/eks-connector.html
func eks_DeregisterCluster(cfg aws.Config, client *eks.Client) {
	input := &eks.DeregisterClusterInput{
		// Name: *string, // Required
	}

	if len(_eksName) > 0 {
		input.Name = aws.String(_eksName)
	}

	if resp, err := client.DeregisterCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an access entry.
func eks_DescribeAccessEntry(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeAccessEntryInput{
		// ClusterName: *string, // Required
		// PrincipalArn: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_eksPrincipalArn)
	}

	if resp, err := client.DescribeAccessEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an Amazon EKS add-on.
func eks_DescribeAddon(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeAddonInput{
		// AddonName: *string, // Required
		// ClusterName: *string, // Required
	}

	if len(_eksAddonName) > 0 {
		input.AddonName = aws.String(_eksAddonName)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}

	if resp, err := client.DescribeAddon(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns configuration options.
func eks_DescribeAddonConfiguration(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeAddonConfigurationInput{
		// AddonName: *string, // Required
		// AddonVersion: *string, // Required
	}

	if len(_eksAddonName) > 0 {
		input.AddonName = aws.String(_eksAddonName)
	}
	if len(_eksAddonVersion) > 0 {
		input.AddonVersion = aws.String(_eksAddonVersion)
	}

	if resp, err := client.DescribeAddonConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the versions for an add-on.
// Information such as the Kubernetes versions that you can use the add-on with,
// the owner , publisher , and the type of the add-on are returned.
func eks_DescribeAddonVersions(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeAddonVersionsInput{}

	if len(_eksAddonName) > 0 {
		input.AddonName = aws.String(_eksAddonName)
	}
	if len(_eksKubernetesVersion) > 0 {
		input.KubernetesVersion = aws.String(_eksKubernetesVersion)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}
	if len(_eksOwners) > 0 {
		input.Owners = append([]string(nil), _eksOwners...)
	}
	if len(_eksPublishers) > 0 {
		input.Publishers = append([]string(nil), _eksPublishers...)
	}
	if len(_eksTypes) > 0 {
		input.Types = append([]string(nil), _eksTypes...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAddonVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.DescribeAddonVersionsOutput
	p := eks.NewDescribeAddonVersionsPaginator(client, input)
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

// Returns detailed information about a specific managed capability in your Amazon
// EKS cluster, including its current status, configuration, health information,
// and any issues that may be affecting its operation.
func eks_DescribeCapability(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeCapabilityInput{
		// CapabilityName: *string, // Required
		// ClusterName: *string, // Required
	}

	if len(_eksCapabilityName) > 0 {
		input.CapabilityName = aws.String(_eksCapabilityName)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}

	if resp, err := client.DescribeCapability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an Amazon EKS cluster.
// The API server endpoint and certificate authority data returned by this
// operation are required for kubelet and kubectl to communicate with your
// Kubernetes API server. For more information, see [Creating or updating a kubeconfig file for an Amazon EKS cluster]kubeconfig .
//
// The API server endpoint and certificate authority data aren't available until
// the cluster reaches the ACTIVE state.
//
// [Creating or updating a kubeconfig file for an Amazon EKS cluster]: https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html
func eks_DescribeCluster(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeClusterInput{
		// Name: *string, // Required
	}

	if len(_eksName) > 0 {
		input.Name = aws.String(_eksName)
	}

	if resp, err := client.DescribeCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists available Kubernetes versions for Amazon EKS clusters.
func eks_DescribeClusterVersions(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeClusterVersionsInput{}

	if len(_eksClusterType) > 0 {
		input.ClusterType = aws.String(_eksClusterType)
	}
	if len(_eksClusterVersions) > 0 {
		input.ClusterVersions = append([]string(nil), _eksClusterVersions...)
	}
	if len(_eksDefaultOnly) > 0 {
		if err := assignInputField(input, "DefaultOnly", _eksDefaultOnly); err != nil {
			log.Errorf("invalid --default-only: %s", err.Error())
			return
		}
	}
	if len(_eksIncludeAll) > 0 {
		if err := assignInputField(input, "IncludeAll", _eksIncludeAll); err != nil {
			log.Errorf("invalid --include-all: %s", err.Error())
			return
		}
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}
	if len(_eksStatus) > 0 {
		if err := assignInputField(input, "Status", _eksStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_eksVersionStatus) > 0 {
		if err := assignInputField(input, "VersionStatus", _eksVersionStatus); err != nil {
			log.Errorf("invalid --version-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusterVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.DescribeClusterVersionsOutput
	p := eks.NewDescribeClusterVersionsPaginator(client, input)
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

// Returns descriptive information about a subscription.
func eks_DescribeEksAnywhereSubscription(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeEksAnywhereSubscriptionInput{
		// Id: *string, // Required
	}

	if len(_eksId) > 0 {
		input.Id = aws.String(_eksId)
	}

	if resp, err := client.DescribeEksAnywhereSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an Fargate profile.
func eks_DescribeFargateProfile(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeFargateProfileInput{
		// ClusterName: *string, // Required
		// FargateProfileName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksFargateProfileName) > 0 {
		input.FargateProfileName = aws.String(_eksFargateProfileName)
	}

	if resp, err := client.DescribeFargateProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an identity provider configuration.
func eks_DescribeIdentityProviderConfig(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeIdentityProviderConfigInput{
		// ClusterName: *string, // Required
		// IdentityProviderConfig: *types.IdentityProviderConfig, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksIdentityProviderConfig) > 0 {
		if err := assignInputField(input, "IdentityProviderConfig", _eksIdentityProviderConfig); err != nil {
			log.Errorf("invalid --identity-provider-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeIdentityProviderConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about an insight that you specify using its ID.
func eks_DescribeInsight(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeInsightInput{
		// ClusterName: *string, // Required
		// Id: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksId) > 0 {
		input.Id = aws.String(_eksId)
	}

	if resp, err := client.DescribeInsight(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the status of the latest on-demand cluster insights refresh operation.
func eks_DescribeInsightsRefresh(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeInsightsRefreshInput{
		// ClusterName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}

	if resp, err := client.DescribeInsightsRefresh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a managed node group.
func eks_DescribeNodegroup(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeNodegroupInput{
		// ClusterName: *string, // Required
		// NodegroupName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksNodegroupName) > 0 {
		input.NodegroupName = aws.String(_eksNodegroupName)
	}

	if resp, err := client.DescribeNodegroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns descriptive information about an EKS Pod Identity association.
// This action requires the ID of the association. You can get the ID from the
// response to the CreatePodIdentityAssocation for newly created associations. Or,
// you can list the IDs for associations with ListPodIdentityAssociations and
// filter the list by namespace or service account.
func eks_DescribePodIdentityAssociation(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribePodIdentityAssociationInput{
		// AssociationId: *string, // Required
		// ClusterName: *string, // Required
	}

	if len(_eksAssociationId) > 0 {
		input.AssociationId = aws.String(_eksAssociationId)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}

	if resp, err := client.DescribePodIdentityAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an update to an Amazon EKS resource.
// When the status of the update is Successful , the update is complete. If an
// update fails, the status is Failed , and an error detail explains the reason for
// the failure.
func eks_DescribeUpdate(cfg aws.Config, client *eks.Client) {
	input := &eks.DescribeUpdateInput{
		// Name: *string, // Required
		// UpdateId: *string, // Required
	}

	if len(_eksName) > 0 {
		input.Name = aws.String(_eksName)
	}
	if len(_eksUpdateId) > 0 {
		input.UpdateId = aws.String(_eksUpdateId)
	}
	if len(_eksAddonName) > 0 {
		input.AddonName = aws.String(_eksAddonName)
	}
	if len(_eksCapabilityName) > 0 {
		input.CapabilityName = aws.String(_eksCapabilityName)
	}
	if len(_eksNodegroupName) > 0 {
		input.NodegroupName = aws.String(_eksNodegroupName)
	}

	if resp, err := client.DescribeUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an access policy from an access entry.
func eks_DisassociateAccessPolicy(cfg aws.Config, client *eks.Client) {
	input := &eks.DisassociateAccessPolicyInput{
		// ClusterName: *string, // Required
		// PolicyArn: *string, // Required
		// PrincipalArn: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksPolicyArn) > 0 {
		input.PolicyArn = aws.String(_eksPolicyArn)
	}
	if len(_eksPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_eksPrincipalArn)
	}

	if resp, err := client.DisassociateAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an identity provider configuration from a cluster.
// If you disassociate an identity provider from your cluster, users included in
// the provider can no longer access the cluster. However, you can still access the
// cluster with IAM principals.
func eks_DisassociateIdentityProviderConfig(cfg aws.Config, client *eks.Client) {
	input := &eks.DisassociateIdentityProviderConfigInput{
		// ClusterName: *string, // Required
		// IdentityProviderConfig: *types.IdentityProviderConfig, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksIdentityProviderConfig) > 0 {
		if err := assignInputField(input, "IdentityProviderConfig", _eksIdentityProviderConfig); err != nil {
			log.Errorf("invalid --identity-provider-config: %s", err.Error())
			return
		}
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}

	if resp, err := client.DisassociateIdentityProviderConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the access entries for your cluster.
func eks_ListAccessEntries(cfg aws.Config, client *eks.Client) {
	input := &eks.ListAccessEntriesInput{
		// ClusterName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksAssociatedPolicyArn) > 0 {
		input.AssociatedPolicyArn = aws.String(_eksAssociatedPolicyArn)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessEntries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListAccessEntriesOutput
	p := eks.NewListAccessEntriesPaginator(client, input)
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

// Lists the available access policies.
func eks_ListAccessPolicies(cfg aws.Config, client *eks.Client) {
	input := &eks.ListAccessPoliciesInput{}

	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListAccessPoliciesOutput
	p := eks.NewListAccessPoliciesPaginator(client, input)
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

// Lists the installed add-ons.
func eks_ListAddons(cfg aws.Config, client *eks.Client) {
	input := &eks.ListAddonsInput{
		// ClusterName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAddons(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListAddonsOutput
	p := eks.NewListAddonsPaginator(client, input)
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

// Lists the access policies associated with an access entry.
func eks_ListAssociatedAccessPolicies(cfg aws.Config, client *eks.Client) {
	input := &eks.ListAssociatedAccessPoliciesInput{
		// ClusterName: *string, // Required
		// PrincipalArn: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_eksPrincipalArn)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssociatedAccessPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListAssociatedAccessPoliciesOutput
	p := eks.NewListAssociatedAccessPoliciesPaginator(client, input)
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

// Lists all managed capabilities in your Amazon EKS cluster. You can use this
// operation to get an overview of all capabilities and their current status.
func eks_ListCapabilities(cfg aws.Config, client *eks.Client) {
	input := &eks.ListCapabilitiesInput{
		// ClusterName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCapabilities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListCapabilitiesOutput
	p := eks.NewListCapabilitiesPaginator(client, input)
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

// Lists the Amazon EKS clusters in your Amazon Web Services account in the
// specified Amazon Web Services Region.
func eks_ListClusters(cfg aws.Config, client *eks.Client) {
	input := &eks.ListClustersInput{}

	if len(_eksInclude) > 0 {
		input.Include = append([]string(nil), _eksInclude...)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListClustersOutput
	p := eks.NewListClustersPaginator(client, input)
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

// Displays the full description of the subscription.
func eks_ListEksAnywhereSubscriptions(cfg aws.Config, client *eks.Client) {
	input := &eks.ListEksAnywhereSubscriptionsInput{}

	if len(_eksIncludeStatus) > 0 {
		if err := assignInputField(input, "IncludeStatus", _eksIncludeStatus); err != nil {
			log.Errorf("invalid --include-status: %s", err.Error())
			return
		}
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEksAnywhereSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListEksAnywhereSubscriptionsOutput
	p := eks.NewListEksAnywhereSubscriptionsPaginator(client, input)
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

// Lists the Fargate profiles associated with the specified cluster in your Amazon
// Web Services account in the specified Amazon Web Services Region.
func eks_ListFargateProfiles(cfg aws.Config, client *eks.Client) {
	input := &eks.ListFargateProfilesInput{
		// ClusterName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFargateProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListFargateProfilesOutput
	p := eks.NewListFargateProfilesPaginator(client, input)
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

// Lists the identity provider configurations for your cluster.
func eks_ListIdentityProviderConfigs(cfg aws.Config, client *eks.Client) {
	input := &eks.ListIdentityProviderConfigsInput{
		// ClusterName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdentityProviderConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListIdentityProviderConfigsOutput
	p := eks.NewListIdentityProviderConfigsPaginator(client, input)
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

// Returns a list of all insights checked for against the specified cluster. You
// can filter which insights are returned by category, associated Kubernetes
// version, and status. The default filter lists all categories and every status.
//
// The following lists the available categories:
//
// - UPGRADE_READINESS : Amazon EKS identifies issues that could impact your
// ability to upgrade to new versions of Kubernetes. These are called upgrade
// insights.
//
// - MISCONFIGURATION : Amazon EKS identifies misconfiguration in your EKS Hybrid
// Nodes setup that could impair functionality of your cluster or workloads. These
// are called configuration insights.
func eks_ListInsights(cfg aws.Config, client *eks.Client) {
	input := &eks.ListInsightsInput{
		// ClusterName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksFilter) > 0 {
		if err := assignInputField(input, "Filter", _eksFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInsights(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListInsightsOutput
	p := eks.NewListInsightsPaginator(client, input)
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

// Lists the managed node groups associated with the specified cluster in your
// Amazon Web Services account in the specified Amazon Web Services Region.
// Self-managed node groups aren't listed.
func eks_ListNodegroups(cfg aws.Config, client *eks.Client) {
	input := &eks.ListNodegroupsInput{
		// ClusterName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListNodegroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListNodegroupsOutput
	p := eks.NewListNodegroupsPaginator(client, input)
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

// List the EKS Pod Identity associations in a cluster. You can filter the list by
// the namespace that the association is in or the service account that the
// association uses.
func eks_ListPodIdentityAssociations(cfg aws.Config, client *eks.Client) {
	input := &eks.ListPodIdentityAssociationsInput{
		// ClusterName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNamespace) > 0 {
		input.Namespace = aws.String(_eksNamespace)
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}
	if len(_eksServiceAccount) > 0 {
		input.ServiceAccount = aws.String(_eksServiceAccount)
	}

	if disablePaginator() {
		if resp, err := client.ListPodIdentityAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListPodIdentityAssociationsOutput
	p := eks.NewListPodIdentityAssociationsPaginator(client, input)
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

// List the tags for an Amazon EKS resource.
func eks_ListTagsForResource(cfg aws.Config, client *eks.Client) {
	input := &eks.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_eksResourceArn) > 0 {
		input.ResourceArn = aws.String(_eksResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the updates associated with an Amazon EKS resource in your Amazon Web
// Services account, in the specified Amazon Web Services Region.
func eks_ListUpdates(cfg aws.Config, client *eks.Client) {
	input := &eks.ListUpdatesInput{
		// Name: *string, // Required
	}

	if len(_eksName) > 0 {
		input.Name = aws.String(_eksName)
	}
	if len(_eksAddonName) > 0 {
		input.AddonName = aws.String(_eksAddonName)
	}
	if len(_eksCapabilityName) > 0 {
		input.CapabilityName = aws.String(_eksCapabilityName)
	}
	if len(_eksMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _eksMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_eksNextToken) > 0 {
		input.NextToken = aws.String(_eksNextToken)
	}
	if len(_eksNodegroupName) > 0 {
		input.NodegroupName = aws.String(_eksNodegroupName)
	}

	if disablePaginator() {
		if resp, err := client.ListUpdates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*eks.ListUpdatesOutput
	p := eks.NewListUpdatesPaginator(client, input)
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

// Connects a Kubernetes cluster to the Amazon EKS control plane.
// Any Kubernetes cluster can be connected to the Amazon EKS control plane to view
// current information about the cluster and its nodes.
//
// Cluster connection requires two steps. First, send a [RegisterClusterRequest]RegisterClusterRequest to
// add it to the Amazon EKS control plane.
//
// Second, a [Manifest] containing the activationID and activationCode must be applied to
// the Kubernetes cluster through it's native provider to provide visibility.
//
// After the manifest is updated and applied, the connected cluster is visible to
// the Amazon EKS control plane. If the manifest isn't applied within three days,
// the connected cluster will no longer be visible and must be deregistered using
// DeregisterCluster .
//
// [RegisterClusterRequest]: https://docs.aws.amazon.com/eks/latest/APIReference/API_RegisterClusterRequest.html
// [Manifest]: https://amazon-eks.s3.us-west-2.amazonaws.com/eks-connector/manifests/eks-connector/latest/eks-connector.yaml
func eks_RegisterCluster(cfg aws.Config, client *eks.Client) {
	input := &eks.RegisterClusterInput{
		// ConnectorConfig: *types.ConnectorConfigRequest, // Required
		// Name: *string, // Required
	}

	if len(_eksConnectorConfig) > 0 {
		if err := assignInputField(input, "ConnectorConfig", _eksConnectorConfig); err != nil {
			log.Errorf("invalid --connector-config: %s", err.Error())
			return
		}
	}
	if len(_eksName) > 0 {
		input.Name = aws.String(_eksName)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates an on-demand refresh operation for cluster insights, getting the
// latest analysis outside of the standard refresh schedule.
func eks_StartInsightsRefresh(cfg aws.Config, client *eks.Client) {
	input := &eks.StartInsightsRefreshInput{
		// ClusterName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}

	if resp, err := client.StartInsightsRefresh(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified tags to an Amazon EKS resource with the specified
// resourceArn . If existing tags on a resource are not specified in the request
// parameters, they aren't changed. When a resource is deleted, the tags associated
// with that resource are also deleted. Tags that you create for Amazon EKS
// resources don't propagate to any other resources associated with the cluster.
// For example, if you tag a cluster with this operation, that tag doesn't
// automatically propagate to the subnets and nodes associated with the cluster.
func eks_TagResource(cfg aws.Config, client *eks.Client) {
	input := &eks.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_eksResourceArn) > 0 {
		input.ResourceArn = aws.String(_eksResourceArn)
	}
	if len(_eksTags) > 0 {
		if err := assignInputField(input, "Tags", _eksTags); err != nil {
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

// Deletes specified tags from an Amazon EKS resource.
func eks_UntagResource(cfg aws.Config, client *eks.Client) {
	input := &eks.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_eksResourceArn) > 0 {
		input.ResourceArn = aws.String(_eksResourceArn)
	}
	if len(_eksTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _eksTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an access entry.
func eks_UpdateAccessEntry(cfg aws.Config, client *eks.Client) {
	input := &eks.UpdateAccessEntryInput{
		// ClusterName: *string, // Required
		// PrincipalArn: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_eksPrincipalArn)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksKubernetesGroups) > 0 {
		input.KubernetesGroups = append([]string(nil), _eksKubernetesGroups...)
	}
	if len(_eksUsername) > 0 {
		input.Username = aws.String(_eksUsername)
	}

	if resp, err := client.UpdateAccessEntry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon EKS add-on.
func eks_UpdateAddon(cfg aws.Config, client *eks.Client) {
	input := &eks.UpdateAddonInput{
		// AddonName: *string, // Required
		// ClusterName: *string, // Required
	}

	if len(_eksAddonName) > 0 {
		input.AddonName = aws.String(_eksAddonName)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksAddonVersion) > 0 {
		input.AddonVersion = aws.String(_eksAddonVersion)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksConfigurationValues) > 0 {
		input.ConfigurationValues = aws.String(_eksConfigurationValues)
	}
	if len(_eksPodIdentityAssociations) > 0 {
		if err := assignInputField(input, "PodIdentityAssociations", _eksPodIdentityAssociations); err != nil {
			log.Errorf("invalid --pod-identity-associations: %s", err.Error())
			return
		}
	}
	if len(_eksResolveConflicts) > 0 {
		if err := assignInputField(input, "ResolveConflicts", _eksResolveConflicts); err != nil {
			log.Errorf("invalid --resolve-conflicts: %s", err.Error())
			return
		}
	}
	if len(_eksServiceAccountRoleArn) > 0 {
		input.ServiceAccountRoleArn = aws.String(_eksServiceAccountRoleArn)
	}

	if resp, err := client.UpdateAddon(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of a managed capability in your Amazon EKS cluster.
// You can update the IAM role, configuration settings, and delete propagation
// policy for a capability.
//
// When you update a capability, Amazon EKS applies the changes and may restart
// capability components as needed. The capability remains available during the
// update process, but some operations may be temporarily unavailable.
func eks_UpdateCapability(cfg aws.Config, client *eks.Client) {
	input := &eks.UpdateCapabilityInput{
		// CapabilityName: *string, // Required
		// ClusterName: *string, // Required
	}

	if len(_eksCapabilityName) > 0 {
		input.CapabilityName = aws.String(_eksCapabilityName)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _eksConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_eksDeletePropagationPolicy) > 0 {
		if err := assignInputField(input, "DeletePropagationPolicy", _eksDeletePropagationPolicy); err != nil {
			log.Errorf("invalid --delete-propagation-policy: %s", err.Error())
			return
		}
	}
	if len(_eksRoleArn) > 0 {
		input.RoleArn = aws.String(_eksRoleArn)
	}

	if resp, err := client.UpdateCapability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon EKS cluster configuration. Your cluster continues to function
// during the update. The response output includes an update ID that you can use to
// track the status of your cluster update with DescribeUpdate .
//
// You can use this operation to do the following actions:
//
// - You can use this API operation to enable or disable exporting the
// Kubernetes control plane logs for your cluster to CloudWatch Logs. By default,
// cluster control plane logs aren't exported to CloudWatch Logs. For more
// information, see [Amazon EKS Cluster control plane logs]in the Amazon EKS User Guide .
//
// # CloudWatch Logs ingestion, archive storage, and data scanning rates apply to
//
// exported control plane logs. For more information, see [CloudWatch Pricing].
//
// - You can also use this API operation to enable or disable public and private
// access to your cluster's Kubernetes API server endpoint. By default, public
// access is enabled, and private access is disabled. For more information, see [Cluster API server endpoint]
// in the Amazon EKS User Guide .
//
// - You can also use this API operation to choose different subnets and
// security groups for the cluster. You must specify at least two subnets that are
// in different Availability Zones. You can't change which VPC the subnets are
// from, the subnets must be in the same VPC as the subnets that the cluster was
// created with. For more information about the VPC requirements, see [https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html]in the
// Amazon EKS User Guide .
//
// - You can also use this API operation to enable or disable ARC zonal shift.
// If zonal shift is enabled, Amazon Web Services configures zonal autoshift for
// the cluster.
//
// - You can also use this API operation to add, change, or remove the
// configuration in the cluster for EKS Hybrid Nodes. To remove the configuration,
// use the remoteNetworkConfig key with an object containing both subkeys with
// empty arrays for each. Here is an inline example: "remoteNetworkConfig": {
// "remoteNodeNetworks": [], "remotePodNetworks": [] } .
//
// Cluster updates are asynchronous, and they should finish within a few minutes.
// During an update, the cluster status moves to UPDATING (this status transition
// is eventually consistent). When the update is complete (either Failed or
// Successful ), the cluster status moves to Active .
//
// [Amazon EKS Cluster control plane logs]: https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html
// [Cluster API server endpoint]: https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html
// [CloudWatch Pricing]: http://aws.amazon.com/cloudwatch/pricing/
// [https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html]: https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html
func eks_UpdateClusterConfig(cfg aws.Config, client *eks.Client) {
	input := &eks.UpdateClusterConfigInput{
		// Name: *string, // Required
	}

	if len(_eksName) > 0 {
		input.Name = aws.String(_eksName)
	}
	if len(_eksAccessConfig) > 0 {
		if err := assignInputField(input, "AccessConfig", _eksAccessConfig); err != nil {
			log.Errorf("invalid --access-config: %s", err.Error())
			return
		}
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksComputeConfig) > 0 {
		if err := assignInputField(input, "ComputeConfig", _eksComputeConfig); err != nil {
			log.Errorf("invalid --compute-config: %s", err.Error())
			return
		}
	}
	if len(_eksControlPlaneScalingConfig) > 0 {
		if err := assignInputField(input, "ControlPlaneScalingConfig", _eksControlPlaneScalingConfig); err != nil {
			log.Errorf("invalid --control-plane-scaling-config: %s", err.Error())
			return
		}
	}
	if len(_eksDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _eksDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_eksKubernetesNetworkConfig) > 0 {
		if err := assignInputField(input, "KubernetesNetworkConfig", _eksKubernetesNetworkConfig); err != nil {
			log.Errorf("invalid --kubernetes-network-config: %s", err.Error())
			return
		}
	}
	if len(_eksLogging) > 0 {
		if err := assignInputField(input, "Logging", _eksLogging); err != nil {
			log.Errorf("invalid --logging: %s", err.Error())
			return
		}
	}
	if len(_eksRemoteNetworkConfig) > 0 {
		if err := assignInputField(input, "RemoteNetworkConfig", _eksRemoteNetworkConfig); err != nil {
			log.Errorf("invalid --remote-network-config: %s", err.Error())
			return
		}
	}
	if len(_eksResourcesVpcConfig) > 0 {
		if err := assignInputField(input, "ResourcesVpcConfig", _eksResourcesVpcConfig); err != nil {
			log.Errorf("invalid --resources-vpc-config: %s", err.Error())
			return
		}
	}
	if len(_eksStorageConfig) > 0 {
		if err := assignInputField(input, "StorageConfig", _eksStorageConfig); err != nil {
			log.Errorf("invalid --storage-config: %s", err.Error())
			return
		}
	}
	if len(_eksUpgradePolicy) > 0 {
		if err := assignInputField(input, "UpgradePolicy", _eksUpgradePolicy); err != nil {
			log.Errorf("invalid --upgrade-policy: %s", err.Error())
			return
		}
	}
	if len(_eksZonalShiftConfig) > 0 {
		if err := assignInputField(input, "ZonalShiftConfig", _eksZonalShiftConfig); err != nil {
			log.Errorf("invalid --zonal-shift-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateClusterConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon EKS cluster to the specified Kubernetes version. Your cluster
// continues to function during the update. The response output includes an update
// ID that you can use to track the status of your cluster update with the [DescribeUpdate]
// DescribeUpdate API operation.
//
// Cluster updates are asynchronous, and they should finish within a few minutes.
// During an update, the cluster status moves to UPDATING (this status transition
// is eventually consistent). When the update is complete (either Failed or
// Successful ), the cluster status moves to Active .
//
// If your cluster has managed node groups attached to it, all of your node
// groups' Kubernetes versions must match the cluster's Kubernetes version in order
// to update the cluster to a new Kubernetes version.
//
// [DescribeUpdate]: https://docs.aws.amazon.com/eks/latest/APIReference/API_DescribeUpdate.html
func eks_UpdateClusterVersion(cfg aws.Config, client *eks.Client) {
	input := &eks.UpdateClusterVersionInput{
		// Name: *string, // Required
		// Version: *string, // Required
	}

	if len(_eksName) > 0 {
		input.Name = aws.String(_eksName)
	}
	if len(_eksVersion) > 0 {
		input.Version = aws.String(_eksVersion)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksForce) > 0 {
		if err := assignInputField(input, "Force", _eksForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateClusterVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an EKS Anywhere Subscription. Only auto renewal and tags can be updated
// after subscription creation.
func eks_UpdateEksAnywhereSubscription(cfg aws.Config, client *eks.Client) {
	input := &eks.UpdateEksAnywhereSubscriptionInput{
		// AutoRenew: bool, // Required
		// Id: *string, // Required
	}

	if len(_eksAutoRenew) > 0 {
		if err := assignInputField(input, "AutoRenew", _eksAutoRenew); err != nil {
			log.Errorf("invalid --auto-renew: %s", err.Error())
			return
		}
	}
	if len(_eksId) > 0 {
		input.Id = aws.String(_eksId)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}

	if resp, err := client.UpdateEksAnywhereSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon EKS managed node group configuration. Your node group
// continues to function during the update. The response output includes an update
// ID that you can use to track the status of your node group update with the [DescribeUpdate]
// DescribeUpdate API operation. You can update the Kubernetes labels and taints
// for a node group and the scaling and version update configuration.
//
// [DescribeUpdate]: https://docs.aws.amazon.com/eks/latest/APIReference/API_DescribeUpdate.html
func eks_UpdateNodegroupConfig(cfg aws.Config, client *eks.Client) {
	input := &eks.UpdateNodegroupConfigInput{
		// ClusterName: *string, // Required
		// NodegroupName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksNodegroupName) > 0 {
		input.NodegroupName = aws.String(_eksNodegroupName)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksLabels) > 0 {
		if err := assignInputField(input, "Labels", _eksLabels); err != nil {
			log.Errorf("invalid --labels: %s", err.Error())
			return
		}
	}
	if len(_eksNodeRepairConfig) > 0 {
		if err := assignInputField(input, "NodeRepairConfig", _eksNodeRepairConfig); err != nil {
			log.Errorf("invalid --node-repair-config: %s", err.Error())
			return
		}
	}
	if len(_eksScalingConfig) > 0 {
		if err := assignInputField(input, "ScalingConfig", _eksScalingConfig); err != nil {
			log.Errorf("invalid --scaling-config: %s", err.Error())
			return
		}
	}
	if len(_eksTaints) > 0 {
		if err := assignInputField(input, "Taints", _eksTaints); err != nil {
			log.Errorf("invalid --taints: %s", err.Error())
			return
		}
	}
	if len(_eksUpdateConfig) > 0 {
		if err := assignInputField(input, "UpdateConfig", _eksUpdateConfig); err != nil {
			log.Errorf("invalid --update-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNodegroupConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Kubernetes version or AMI version of an Amazon EKS managed node
// group.
//
// You can update a node group using a launch template only if the node group was
// originally deployed with a launch template. Additionally, the launch template ID
// or name must match what was used when the node group was created. You can update
// the launch template version with necessary changes.
//
// If you need to update a custom AMI in a node group that was deployed with a
// launch template, then update your custom AMI, specify the new ID in a new
// version of the launch template, and then update the node group to the new
// version of the launch template.
//
// If you update without a launch template, then you can update to the latest
// available AMI version of a node group's current Kubernetes version by not
// specifying a Kubernetes version in the request. You can update to the latest AMI
// version of your cluster's current Kubernetes version by specifying your
// cluster's Kubernetes version in the request. For information about Linux
// versions, see [Amazon EKS optimized Amazon Linux AMI versions]in the Amazon EKS User Guide. For information about Windows
// versions, see [Amazon EKS optimized Windows AMI versions]in the Amazon EKS User Guide.
//
// You cannot roll back a node group to an earlier Kubernetes version or AMI
// version.
//
// When a node in a managed node group is terminated due to a scaling action or
// update, every Pod on that node is drained first. Amazon EKS attempts to drain
// the nodes gracefully and will fail if it is unable to do so. You can force the
// update if Amazon EKS is unable to drain the nodes as a result of a Pod
// disruption budget issue.
//
// [Amazon EKS optimized Amazon Linux AMI versions]: https://docs.aws.amazon.com/eks/latest/userguide/eks-linux-ami-versions.html
// [Amazon EKS optimized Windows AMI versions]: https://docs.aws.amazon.com/eks/latest/userguide/eks-ami-versions-windows.html
func eks_UpdateNodegroupVersion(cfg aws.Config, client *eks.Client) {
	input := &eks.UpdateNodegroupVersionInput{
		// ClusterName: *string, // Required
		// NodegroupName: *string, // Required
	}

	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksNodegroupName) > 0 {
		input.NodegroupName = aws.String(_eksNodegroupName)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksForce) > 0 {
		if err := assignInputField(input, "Force", _eksForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_eksLaunchTemplate) > 0 {
		if err := assignInputField(input, "LaunchTemplate", _eksLaunchTemplate); err != nil {
			log.Errorf("invalid --launch-template: %s", err.Error())
			return
		}
	}
	if len(_eksReleaseVersion) > 0 {
		input.ReleaseVersion = aws.String(_eksReleaseVersion)
	}
	if len(_eksVersion) > 0 {
		input.Version = aws.String(_eksVersion)
	}

	if resp, err := client.UpdateNodegroupVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a EKS Pod Identity association. In an update, you can change the IAM
// role, the target IAM role, or disableSessionTags . You must change at least one
// of these in an update. An association can't be moved between clusters,
// namespaces, or service accounts. If you need to edit the namespace or service
// account, you need to delete the association and then create a new association
// with your desired settings.
//
// Similar to Amazon Web Services IAM behavior, EKS Pod Identity associations are
// eventually consistent, and may take several seconds to be effective after the
// initial API call returns successfully. You must design your applications to
// account for these potential delays. We recommend that you don’t include
// association create/updates in the critical, high-availability code paths of your
// application. Instead, make changes in a separate initialization or setup routine
// that you run less frequently.
//
// You can set a target IAM role in the same or a different account for advanced
// scenarios. With a target role, EKS Pod Identity automatically performs two role
// assumptions in sequence: first assuming the role in the association that is in
// this account, then using those credentials to assume the target IAM role. This
// process provides your Pod with temporary credentials that have the permissions
// defined in the target role, allowing secure access to resources in another
// Amazon Web Services account.
func eks_UpdatePodIdentityAssociation(cfg aws.Config, client *eks.Client) {
	input := &eks.UpdatePodIdentityAssociationInput{
		// AssociationId: *string, // Required
		// ClusterName: *string, // Required
	}

	if len(_eksAssociationId) > 0 {
		input.AssociationId = aws.String(_eksAssociationId)
	}
	if len(_eksClusterName) > 0 {
		input.ClusterName = aws.String(_eksClusterName)
	}
	if len(_eksClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_eksClientRequestToken)
	}
	if len(_eksDisableSessionTags) > 0 {
		if err := assignInputField(input, "DisableSessionTags", _eksDisableSessionTags); err != nil {
			log.Errorf("invalid --disable-session-tags: %s", err.Error())
			return
		}
	}
	if len(_eksPolicy) > 0 {
		input.Policy = aws.String(_eksPolicy)
	}
	if len(_eksRoleArn) > 0 {
		input.RoleArn = aws.String(_eksRoleArn)
	}
	if len(_eksTargetRoleArn) > 0 {
		input.TargetRoleArn = aws.String(_eksTargetRoleArn)
	}

	if resp, err := client.UpdatePodIdentityAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_eksCmd)
	_eksCmd.Flags().SortFlags = false

	_eksCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_eksCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_eksCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_eksCmd.Flags().StringVarP(&_eksAccessConfig, "access-config", "", "", "Access Config")
	_eksCmd.Flags().StringVarP(&_eksAccessScope, "access-scope", "", "", "Access Scope")
	_eksCmd.Flags().StringVarP(&_eksAddonName, "addon-name", "", "", "Addon Name")
	_eksCmd.Flags().StringVarP(&_eksAddonVersion, "addon-version", "", "", "Addon Version")
	_eksCmd.Flags().StringVarP(&_eksAmiType, "ami-type", "", "", "AMI Type")
	_eksCmd.Flags().StringVarP(&_eksAssociatedPolicyArn, "associated-policy-arn", "", "", "Associated Policy ARN")
	_eksCmd.Flags().StringVarP(&_eksAssociationId, "association-id", "", "", "Association ID")
	_eksCmd.Flags().StringVarP(&_eksAutoRenew, "auto-renew", "", "", "Auto Renew")
	_eksCmd.Flags().StringVarP(&_eksBootstrapSelfManagedAddons, "bootstrap-self-managed-addons", "", "", "Bootstrap Self Managed Addons")
	_eksCmd.Flags().StringVarP(&_eksCapabilityName, "capability-name", "", "", "Capability Name")
	_eksCmd.Flags().StringVarP(&_eksCapacityType, "capacity-type", "", "", "Capacity Type")
	_eksCmd.Flags().StringVarP(&_eksClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_eksCmd.Flags().StringVarP(&_eksClusterName, "cluster-name", "", "", "Cluster Name")
	_eksCmd.Flags().StringVarP(&_eksClusterType, "cluster-type", "", "", "Cluster Type")
	_eksCmd.Flags().StringSliceVarP(&_eksClusterVersions, "cluster-versions", "", nil, "Cluster Versions")
	_eksCmd.Flags().StringVarP(&_eksComputeConfig, "compute-config", "", "", "Compute Config")
	_eksCmd.Flags().StringVarP(&_eksConfiguration, "configuration", "", "", "Configuration")
	_eksCmd.Flags().StringVarP(&_eksConfigurationValues, "configuration-values", "", "", "Configuration Values")
	_eksCmd.Flags().StringVarP(&_eksConnectorConfig, "connector-config", "", "", "Connector Config")
	_eksCmd.Flags().StringVarP(&_eksControlPlaneScalingConfig, "control-plane-scaling-config", "", "", "Control Plane Scaling Config")
	_eksCmd.Flags().StringVarP(&_eksDefaultOnly, "default-only", "", "", "Default Only")
	_eksCmd.Flags().StringVarP(&_eksDeletePropagationPolicy, "delete-propagation-policy", "", "", "Delete Propagation Policy")
	_eksCmd.Flags().StringVarP(&_eksDeletionProtection, "deletion-protection", "", "", "Deletion Protection")
	_eksCmd.Flags().StringVarP(&_eksDisableSessionTags, "disable-session-tags", "", "", "Disable Session Tags")
	_eksCmd.Flags().StringVarP(&_eksDiskSize, "disk-size", "", "", "Disk Size")
	_eksCmd.Flags().StringVarP(&_eksEncryptionConfig, "encryption-config", "", "", "Encryption Config")
	_eksCmd.Flags().StringVarP(&_eksFargateProfileName, "fargate-profile-name", "", "", "Fargate Profile Name")
	_eksCmd.Flags().StringVarP(&_eksFilter, "filter", "", "", "Filter")
	_eksCmd.Flags().StringVarP(&_eksForce, "force", "", "", "Force")
	_eksCmd.Flags().StringVarP(&_eksId, "id", "", "", "ID")
	_eksCmd.Flags().StringVarP(&_eksIdentityProviderConfig, "identity-provider-config", "", "", "Identity Provider Config")
	_eksCmd.Flags().StringSliceVarP(&_eksInclude, "include", "", nil, "Include")
	_eksCmd.Flags().StringVarP(&_eksIncludeAll, "include-all", "", "", "Include All")
	_eksCmd.Flags().StringVarP(&_eksIncludeStatus, "include-status", "", "", "Include Status")
	_eksCmd.Flags().StringSliceVarP(&_eksInstanceTypes, "instance-types", "", nil, "Instance Types")
	_eksCmd.Flags().StringSliceVarP(&_eksKubernetesGroups, "kubernetes-groups", "", nil, "Kubernetes Groups")
	_eksCmd.Flags().StringVarP(&_eksKubernetesNetworkConfig, "kubernetes-network-config", "", "", "Kubernetes Network Config")
	_eksCmd.Flags().StringVarP(&_eksKubernetesVersion, "kubernetes-version", "", "", "Kubernetes Version")
	_eksCmd.Flags().StringVarP(&_eksLabels, "labels", "", "", "Labels")
	_eksCmd.Flags().StringVarP(&_eksLaunchTemplate, "launch-template", "", "", "Launch Template")
	_eksCmd.Flags().StringVarP(&_eksLicenseQuantity, "license-quantity", "", "", "License Quantity")
	_eksCmd.Flags().StringVarP(&_eksLicenseType, "license-type", "", "", "License Type")
	_eksCmd.Flags().StringVarP(&_eksLogging, "logging", "", "", "Logging")
	_eksCmd.Flags().StringVarP(&_eksMaxResults, "max-results", "", "", "Max Results")
	_eksCmd.Flags().StringVarP(&_eksName, "name", "", "", "Name")
	_eksCmd.Flags().StringVarP(&_eksNamespace, "namespace", "", "", "Namespace")
	_eksCmd.Flags().StringVarP(&_eksNamespaceConfig, "namespace-config", "", "", "Namespace Config")
	_eksCmd.Flags().StringVarP(&_eksNextToken, "next-token", "", "", "Next Token")
	_eksCmd.Flags().StringVarP(&_eksNodeRepairConfig, "node-repair-config", "", "", "Node Repair Config")
	_eksCmd.Flags().StringVarP(&_eksNodeRole, "node-role", "", "", "Node Role")
	_eksCmd.Flags().StringVarP(&_eksNodegroupName, "nodegroup-name", "", "", "Nodegroup Name")
	_eksCmd.Flags().StringVarP(&_eksOidc, "oidc", "", "", "OIDC")
	_eksCmd.Flags().StringVarP(&_eksOutpostConfig, "outpost-config", "", "", "Outpost Config")
	_eksCmd.Flags().StringSliceVarP(&_eksOwners, "owners", "", nil, "Owners")
	_eksCmd.Flags().StringVarP(&_eksPodExecutionRoleArn, "pod-execution-role-arn", "", "", "Pod Execution Role ARN")
	_eksCmd.Flags().StringVarP(&_eksPodIdentityAssociations, "pod-identity-associations", "", "", "Pod Identity Associations")
	_eksCmd.Flags().StringVarP(&_eksPolicy, "policy", "", "", "Policy")
	_eksCmd.Flags().StringVarP(&_eksPolicyArn, "policy-arn", "", "", "Policy ARN")
	_eksCmd.Flags().StringVarP(&_eksPreserve, "preserve", "", "", "Preserve")
	_eksCmd.Flags().StringVarP(&_eksPrincipalArn, "principal-arn", "", "", "Principal ARN")
	_eksCmd.Flags().StringSliceVarP(&_eksPublishers, "publishers", "", nil, "Publishers")
	_eksCmd.Flags().StringVarP(&_eksReleaseVersion, "release-version", "", "", "Release Version")
	_eksCmd.Flags().StringVarP(&_eksRemoteAccess, "remote-access", "", "", "Remote Access")
	_eksCmd.Flags().StringVarP(&_eksRemoteNetworkConfig, "remote-network-config", "", "", "Remote Network Config")
	_eksCmd.Flags().StringVarP(&_eksResolveConflicts, "resolve-conflicts", "", "", "Resolve Conflicts")
	_eksCmd.Flags().StringVarP(&_eksResourceArn, "resource-arn", "", "", "Resource ARN")
	_eksCmd.Flags().StringVarP(&_eksResourcesVpcConfig, "resources-vpc-config", "", "", "Resources VPC Config")
	_eksCmd.Flags().StringVarP(&_eksRoleArn, "role-arn", "", "", "Role ARN")
	_eksCmd.Flags().StringVarP(&_eksScalingConfig, "scaling-config", "", "", "Scaling Config")
	_eksCmd.Flags().StringVarP(&_eksSelectors, "selectors", "", "", "Selectors")
	_eksCmd.Flags().StringVarP(&_eksServiceAccount, "service-account", "", "", "Service Account")
	_eksCmd.Flags().StringVarP(&_eksServiceAccountRoleArn, "service-account-role-arn", "", "", "Service Account Role ARN")
	_eksCmd.Flags().StringVarP(&_eksStatus, "status", "", "", "Status")
	_eksCmd.Flags().StringVarP(&_eksStorageConfig, "storage-config", "", "", "Storage Config")
	_eksCmd.Flags().StringSliceVarP(&_eksSubnets, "subnets", "", nil, "Subnets")
	_eksCmd.Flags().StringSliceVarP(&_eksTagKeys, "tag-keys", "", nil, "Tag Keys")
	_eksCmd.Flags().StringVarP(&_eksTags, "tags", "", "", "Tags")
	_eksCmd.Flags().StringVarP(&_eksTaints, "taints", "", "", "Taints")
	_eksCmd.Flags().StringVarP(&_eksTargetRoleArn, "target-role-arn", "", "", "Target Role ARN")
	_eksCmd.Flags().StringVarP(&_eksTerm, "term", "", "", "Term")
	_eksCmd.Flags().StringVarP(&_eksType, "type", "", "", "Type")
	_eksCmd.Flags().StringSliceVarP(&_eksTypes, "types", "", nil, "Types")
	_eksCmd.Flags().StringVarP(&_eksUpdateConfig, "update-config", "", "", "Update Config")
	_eksCmd.Flags().StringVarP(&_eksUpdateId, "update-id", "", "", "Update ID")
	_eksCmd.Flags().StringVarP(&_eksUpgradePolicy, "upgrade-policy", "", "", "Upgrade Policy")
	_eksCmd.Flags().StringVarP(&_eksUsername, "username", "", "", "Username")
	_eksCmd.Flags().StringVarP(&_eksVersion, "version", "", "", "Version")
	_eksCmd.Flags().StringVarP(&_eksVersionStatus, "version-status", "", "", "Version Status")
	_eksCmd.Flags().StringVarP(&_eksZonalShiftConfig, "zonal-shift-config", "", "", "Zonal Shift Config")

	_eksCmd.Flags().BoolVarP(&_eksAssociateAccessPolicy, "associate-access-policy", "", false, "Associate Access Policy")
	_eksCmd.Flags().BoolVarP(&_eksAssociateEncryptionConfig, "associate-encryption-config", "", false, "Associate Encryption Config")
	_eksCmd.Flags().BoolVarP(&_eksAssociateIdentityProviderConfig, "associate-identity-provider-config", "", false, "Associate Identity Provider Config")
	_eksCmd.Flags().BoolVarP(&_eksCreateAccessEntry, "create-access-entry", "", false, "Create Access Entry")
	_eksCmd.Flags().BoolVarP(&_eksCreateAddon, "create-addon", "", false, "Create Addon")
	_eksCmd.Flags().BoolVarP(&_eksCreateCapability, "create-capability", "", false, "Create Capability")
	_eksCmd.Flags().BoolVarP(&_eksCreateCluster, "create-cluster", "", false, "Create Cluster")
	_eksCmd.Flags().BoolVarP(&_eksCreateEksAnywhereSubscription, "create-eks-anywhere-subscription", "", false, "Create Eks Anywhere Subscription")
	_eksCmd.Flags().BoolVarP(&_eksCreateFargateProfile, "create-fargate-profile", "", false, "Create Fargate Profile")
	_eksCmd.Flags().BoolVarP(&_eksCreateNodegroup, "create-nodegroup", "", false, "Create Nodegroup")
	_eksCmd.Flags().BoolVarP(&_eksCreatePodIdentityAssociation, "create-pod-identity-association", "", false, "Create Pod Identity Association")
	_eksCmd.Flags().BoolVarP(&_eksDeleteAccessEntry, "delete-access-entry", "", false, "Delete Access Entry")
	_eksCmd.Flags().BoolVarP(&_eksDeleteAddon, "delete-addon", "", false, "Delete Addon")
	_eksCmd.Flags().BoolVarP(&_eksDeleteCapability, "delete-capability", "", false, "Delete Capability")
	_eksCmd.Flags().BoolVarP(&_eksDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_eksCmd.Flags().BoolVarP(&_eksDeleteEksAnywhereSubscription, "delete-eks-anywhere-subscription", "", false, "Delete Eks Anywhere Subscription")
	_eksCmd.Flags().BoolVarP(&_eksDeleteFargateProfile, "delete-fargate-profile", "", false, "Delete Fargate Profile")
	_eksCmd.Flags().BoolVarP(&_eksDeleteNodegroup, "delete-nodegroup", "", false, "Delete Nodegroup")
	_eksCmd.Flags().BoolVarP(&_eksDeletePodIdentityAssociation, "delete-pod-identity-association", "", false, "Delete Pod Identity Association")
	_eksCmd.Flags().BoolVarP(&_eksDeregisterCluster, "deregister-cluster", "", false, "Deregister Cluster")
	_eksCmd.Flags().BoolVarP(&_eksDescribeAccessEntry, "describe-access-entry", "", false, "Describe Access Entry")
	_eksCmd.Flags().BoolVarP(&_eksDescribeAddon, "describe-addon", "", false, "Describe Addon")
	_eksCmd.Flags().BoolVarP(&_eksDescribeAddonConfiguration, "describe-addon-configuration", "", false, "Describe Addon Configuration")
	_eksCmd.Flags().BoolVarP(&_eksDescribeAddonVersions, "describe-addon-versions", "", false, "Describe Addon Versions")
	_eksCmd.Flags().BoolVarP(&_eksDescribeCapability, "describe-capability", "", false, "Describe Capability")
	_eksCmd.Flags().BoolVarP(&_eksDescribeCluster, "describe-cluster", "", false, "Describe Cluster")
	_eksCmd.Flags().BoolVarP(&_eksDescribeClusterVersions, "describe-cluster-versions", "", false, "Describe Cluster Versions")
	_eksCmd.Flags().BoolVarP(&_eksDescribeEksAnywhereSubscription, "describe-eks-anywhere-subscription", "", false, "Describe Eks Anywhere Subscription")
	_eksCmd.Flags().BoolVarP(&_eksDescribeFargateProfile, "describe-fargate-profile", "", false, "Describe Fargate Profile")
	_eksCmd.Flags().BoolVarP(&_eksDescribeIdentityProviderConfig, "describe-identity-provider-config", "", false, "Describe Identity Provider Config")
	_eksCmd.Flags().BoolVarP(&_eksDescribeInsight, "describe-insight", "", false, "Describe Insight")
	_eksCmd.Flags().BoolVarP(&_eksDescribeInsightsRefresh, "describe-insights-refresh", "", false, "Describe Insights Refresh")
	_eksCmd.Flags().BoolVarP(&_eksDescribeNodegroup, "describe-nodegroup", "", false, "Describe Nodegroup")
	_eksCmd.Flags().BoolVarP(&_eksDescribePodIdentityAssociation, "describe-pod-identity-association", "", false, "Describe Pod Identity Association")
	_eksCmd.Flags().BoolVarP(&_eksDescribeUpdate, "describe-update", "", false, "Describe Update")
	_eksCmd.Flags().BoolVarP(&_eksDisassociateAccessPolicy, "disassociate-access-policy", "", false, "Disassociate Access Policy")
	_eksCmd.Flags().BoolVarP(&_eksDisassociateIdentityProviderConfig, "disassociate-identity-provider-config", "", false, "Disassociate Identity Provider Config")
	_eksCmd.Flags().BoolVarP(&_eksListAccessEntries, "list-access-entries", "", false, "List Access Entries")
	_eksCmd.Flags().BoolVarP(&_eksListAccessPolicies, "list-access-policies", "", false, "List Access Policies")
	_eksCmd.Flags().BoolVarP(&_eksListAddons, "list-addons", "", false, "List Addons")
	_eksCmd.Flags().BoolVarP(&_eksListAssociatedAccessPolicies, "list-associated-access-policies", "", false, "List Associated Access Policies")
	_eksCmd.Flags().BoolVarP(&_eksListCapabilities, "list-capabilities", "", false, "List Capabilities")
	_eksCmd.Flags().BoolVarP(&_eksListClusters, "list-clusters", "", false, "List Clusters")
	_eksCmd.Flags().BoolVarP(&_eksListEksAnywhereSubscriptions, "list-eks-anywhere-subscriptions", "", false, "List Eks Anywhere Subscriptions")
	_eksCmd.Flags().BoolVarP(&_eksListFargateProfiles, "list-fargate-profiles", "", false, "List Fargate Profiles")
	_eksCmd.Flags().BoolVarP(&_eksListIdentityProviderConfigs, "list-identity-provider-configs", "", false, "List Identity Provider Configs")
	_eksCmd.Flags().BoolVarP(&_eksListInsights, "list-insights", "", false, "List Insights")
	_eksCmd.Flags().BoolVarP(&_eksListNodegroups, "list-nodegroups", "", false, "List Nodegroups")
	_eksCmd.Flags().BoolVarP(&_eksListPodIdentityAssociations, "list-pod-identity-associations", "", false, "List Pod Identity Associations")
	_eksCmd.Flags().BoolVarP(&_eksListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_eksCmd.Flags().BoolVarP(&_eksListUpdates, "list-updates", "", false, "List Updates")
	_eksCmd.Flags().BoolVarP(&_eksRegisterCluster, "register-cluster", "", false, "Register Cluster")
	_eksCmd.Flags().BoolVarP(&_eksStartInsightsRefresh, "start-insights-refresh", "", false, "Start Insights Refresh")
	_eksCmd.Flags().BoolVarP(&_eksTagResource, "tag-resource", "", false, "Tag Resource")
	_eksCmd.Flags().BoolVarP(&_eksUntagResource, "untag-resource", "", false, "Untag Resource")
	_eksCmd.Flags().BoolVarP(&_eksUpdateAccessEntry, "update-access-entry", "", false, "Update Access Entry")
	_eksCmd.Flags().BoolVarP(&_eksUpdateAddon, "update-addon", "", false, "Update Addon")
	_eksCmd.Flags().BoolVarP(&_eksUpdateCapability, "update-capability", "", false, "Update Capability")
	_eksCmd.Flags().BoolVarP(&_eksUpdateClusterConfig, "update-cluster-config", "", false, "Update Cluster Config")
	_eksCmd.Flags().BoolVarP(&_eksUpdateClusterVersion, "update-cluster-version", "", false, "Update Cluster Version")
	_eksCmd.Flags().BoolVarP(&_eksUpdateEksAnywhereSubscription, "update-eks-anywhere-subscription", "", false, "Update Eks Anywhere Subscription")
	_eksCmd.Flags().BoolVarP(&_eksUpdateNodegroupConfig, "update-nodegroup-config", "", false, "Update Nodegroup Config")
	_eksCmd.Flags().BoolVarP(&_eksUpdateNodegroupVersion, "update-nodegroup-version", "", false, "Update Nodegroup Version")
	_eksCmd.Flags().BoolVarP(&_eksUpdatePodIdentityAssociation, "update-pod-identity-association", "", false, "Update Pod Identity Association")

}
