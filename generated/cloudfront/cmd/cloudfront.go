package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudfrontCmd represents the cloudfront command
var _cloudfrontCmd = &cobra.Command{
	Use:   "cloudfront",
	Short: "AWS cloudfront CLI",
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
		client := cloudfront.NewFromConfig(cfg)
		if _cloudfrontAssociateAlias {
			cloudfront_AssociateAlias(cfg, client)
			return
		}
		if _cloudfrontAssociateDistributionTenantWebACL {
			cloudfront_AssociateDistributionTenantWebACL(cfg, client)
			return
		}
		if _cloudfrontAssociateDistributionWebACL {
			cloudfront_AssociateDistributionWebACL(cfg, client)
			return
		}
		if _cloudfrontCopyDistribution {
			cloudfront_CopyDistribution(cfg, client)
			return
		}
		if _cloudfrontCreateAnycastIpList {
			cloudfront_CreateAnycastIpList(cfg, client)
			return
		}
		if _cloudfrontCreateCachePolicy {
			cloudfront_CreateCachePolicy(cfg, client)
			return
		}
		if _cloudfrontCreateCloudFrontOriginAccessIdentity {
			cloudfront_CreateCloudFrontOriginAccessIdentity(cfg, client)
			return
		}
		if _cloudfrontCreateConnectionFunction {
			cloudfront_CreateConnectionFunction(cfg, client)
			return
		}
		if _cloudfrontCreateConnectionGroup {
			cloudfront_CreateConnectionGroup(cfg, client)
			return
		}
		if _cloudfrontCreateContinuousDeploymentPolicy {
			cloudfront_CreateContinuousDeploymentPolicy(cfg, client)
			return
		}
		if _cloudfrontCreateDistribution {
			cloudfront_CreateDistribution(cfg, client)
			return
		}
		if _cloudfrontCreateDistributionTenant {
			cloudfront_CreateDistributionTenant(cfg, client)
			return
		}
		if _cloudfrontCreateDistributionWithTags {
			cloudfront_CreateDistributionWithTags(cfg, client)
			return
		}
		if _cloudfrontCreateFieldLevelEncryptionConfig {
			cloudfront_CreateFieldLevelEncryptionConfig(cfg, client)
			return
		}
		if _cloudfrontCreateFieldLevelEncryptionProfile {
			cloudfront_CreateFieldLevelEncryptionProfile(cfg, client)
			return
		}
		if _cloudfrontCreateFunction {
			cloudfront_CreateFunction(cfg, client)
			return
		}
		if _cloudfrontCreateInvalidation {
			cloudfront_CreateInvalidation(cfg, client)
			return
		}
		if _cloudfrontCreateInvalidationForDistributionTenant {
			cloudfront_CreateInvalidationForDistributionTenant(cfg, client)
			return
		}
		if _cloudfrontCreateKeyGroup {
			cloudfront_CreateKeyGroup(cfg, client)
			return
		}
		if _cloudfrontCreateKeyValueStore {
			cloudfront_CreateKeyValueStore(cfg, client)
			return
		}
		if _cloudfrontCreateMonitoringSubscription {
			cloudfront_CreateMonitoringSubscription(cfg, client)
			return
		}
		if _cloudfrontCreateOriginAccessControl {
			cloudfront_CreateOriginAccessControl(cfg, client)
			return
		}
		if _cloudfrontCreateOriginRequestPolicy {
			cloudfront_CreateOriginRequestPolicy(cfg, client)
			return
		}
		if _cloudfrontCreatePublicKey {
			cloudfront_CreatePublicKey(cfg, client)
			return
		}
		if _cloudfrontCreateRealtimeLogConfig {
			cloudfront_CreateRealtimeLogConfig(cfg, client)
			return
		}
		if _cloudfrontCreateResponseHeadersPolicy {
			cloudfront_CreateResponseHeadersPolicy(cfg, client)
			return
		}
		if _cloudfrontCreateStreamingDistribution {
			cloudfront_CreateStreamingDistribution(cfg, client)
			return
		}
		if _cloudfrontCreateStreamingDistributionWithTags {
			cloudfront_CreateStreamingDistributionWithTags(cfg, client)
			return
		}
		if _cloudfrontCreateTrustStore {
			cloudfront_CreateTrustStore(cfg, client)
			return
		}
		if _cloudfrontCreateVpcOrigin {
			cloudfront_CreateVpcOrigin(cfg, client)
			return
		}
		if _cloudfrontDeleteAnycastIpList {
			cloudfront_DeleteAnycastIpList(cfg, client)
			return
		}
		if _cloudfrontDeleteCachePolicy {
			cloudfront_DeleteCachePolicy(cfg, client)
			return
		}
		if _cloudfrontDeleteCloudFrontOriginAccessIdentity {
			cloudfront_DeleteCloudFrontOriginAccessIdentity(cfg, client)
			return
		}
		if _cloudfrontDeleteConnectionFunction {
			cloudfront_DeleteConnectionFunction(cfg, client)
			return
		}
		if _cloudfrontDeleteConnectionGroup {
			cloudfront_DeleteConnectionGroup(cfg, client)
			return
		}
		if _cloudfrontDeleteContinuousDeploymentPolicy {
			cloudfront_DeleteContinuousDeploymentPolicy(cfg, client)
			return
		}
		if _cloudfrontDeleteDistribution {
			cloudfront_DeleteDistribution(cfg, client)
			return
		}
		if _cloudfrontDeleteDistributionTenant {
			cloudfront_DeleteDistributionTenant(cfg, client)
			return
		}
		if _cloudfrontDeleteFieldLevelEncryptionConfig {
			cloudfront_DeleteFieldLevelEncryptionConfig(cfg, client)
			return
		}
		if _cloudfrontDeleteFieldLevelEncryptionProfile {
			cloudfront_DeleteFieldLevelEncryptionProfile(cfg, client)
			return
		}
		if _cloudfrontDeleteFunction {
			cloudfront_DeleteFunction(cfg, client)
			return
		}
		if _cloudfrontDeleteKeyGroup {
			cloudfront_DeleteKeyGroup(cfg, client)
			return
		}
		if _cloudfrontDeleteKeyValueStore {
			cloudfront_DeleteKeyValueStore(cfg, client)
			return
		}
		if _cloudfrontDeleteMonitoringSubscription {
			cloudfront_DeleteMonitoringSubscription(cfg, client)
			return
		}
		if _cloudfrontDeleteOriginAccessControl {
			cloudfront_DeleteOriginAccessControl(cfg, client)
			return
		}
		if _cloudfrontDeleteOriginRequestPolicy {
			cloudfront_DeleteOriginRequestPolicy(cfg, client)
			return
		}
		if _cloudfrontDeletePublicKey {
			cloudfront_DeletePublicKey(cfg, client)
			return
		}
		if _cloudfrontDeleteRealtimeLogConfig {
			cloudfront_DeleteRealtimeLogConfig(cfg, client)
			return
		}
		if _cloudfrontDeleteResourcePolicy {
			cloudfront_DeleteResourcePolicy(cfg, client)
			return
		}
		if _cloudfrontDeleteResponseHeadersPolicy {
			cloudfront_DeleteResponseHeadersPolicy(cfg, client)
			return
		}
		if _cloudfrontDeleteStreamingDistribution {
			cloudfront_DeleteStreamingDistribution(cfg, client)
			return
		}
		if _cloudfrontDeleteTrustStore {
			cloudfront_DeleteTrustStore(cfg, client)
			return
		}
		if _cloudfrontDeleteVpcOrigin {
			cloudfront_DeleteVpcOrigin(cfg, client)
			return
		}
		if _cloudfrontDescribeConnectionFunction {
			cloudfront_DescribeConnectionFunction(cfg, client)
			return
		}
		if _cloudfrontDescribeFunction {
			cloudfront_DescribeFunction(cfg, client)
			return
		}
		if _cloudfrontDescribeKeyValueStore {
			cloudfront_DescribeKeyValueStore(cfg, client)
			return
		}
		if _cloudfrontDisassociateDistributionTenantWebACL {
			cloudfront_DisassociateDistributionTenantWebACL(cfg, client)
			return
		}
		if _cloudfrontDisassociateDistributionWebACL {
			cloudfront_DisassociateDistributionWebACL(cfg, client)
			return
		}
		if _cloudfrontGetAnycastIpList {
			cloudfront_GetAnycastIpList(cfg, client)
			return
		}
		if _cloudfrontGetCachePolicy {
			cloudfront_GetCachePolicy(cfg, client)
			return
		}
		if _cloudfrontGetCachePolicyConfig {
			cloudfront_GetCachePolicyConfig(cfg, client)
			return
		}
		if _cloudfrontGetCloudFrontOriginAccessIdentity {
			cloudfront_GetCloudFrontOriginAccessIdentity(cfg, client)
			return
		}
		if _cloudfrontGetCloudFrontOriginAccessIdentityConfig {
			cloudfront_GetCloudFrontOriginAccessIdentityConfig(cfg, client)
			return
		}
		if _cloudfrontGetConnectionFunction {
			cloudfront_GetConnectionFunction(cfg, client)
			return
		}
		if _cloudfrontGetConnectionGroup {
			cloudfront_GetConnectionGroup(cfg, client)
			return
		}
		if _cloudfrontGetConnectionGroupByRoutingEndpoint {
			cloudfront_GetConnectionGroupByRoutingEndpoint(cfg, client)
			return
		}
		if _cloudfrontGetContinuousDeploymentPolicy {
			cloudfront_GetContinuousDeploymentPolicy(cfg, client)
			return
		}
		if _cloudfrontGetContinuousDeploymentPolicyConfig {
			cloudfront_GetContinuousDeploymentPolicyConfig(cfg, client)
			return
		}
		if _cloudfrontGetDistribution {
			cloudfront_GetDistribution(cfg, client)
			return
		}
		if _cloudfrontGetDistributionConfig {
			cloudfront_GetDistributionConfig(cfg, client)
			return
		}
		if _cloudfrontGetDistributionTenant {
			cloudfront_GetDistributionTenant(cfg, client)
			return
		}
		if _cloudfrontGetDistributionTenantByDomain {
			cloudfront_GetDistributionTenantByDomain(cfg, client)
			return
		}
		if _cloudfrontGetFieldLevelEncryption {
			cloudfront_GetFieldLevelEncryption(cfg, client)
			return
		}
		if _cloudfrontGetFieldLevelEncryptionConfig {
			cloudfront_GetFieldLevelEncryptionConfig(cfg, client)
			return
		}
		if _cloudfrontGetFieldLevelEncryptionProfile {
			cloudfront_GetFieldLevelEncryptionProfile(cfg, client)
			return
		}
		if _cloudfrontGetFieldLevelEncryptionProfileConfig {
			cloudfront_GetFieldLevelEncryptionProfileConfig(cfg, client)
			return
		}
		if _cloudfrontGetFunction {
			cloudfront_GetFunction(cfg, client)
			return
		}
		if _cloudfrontGetInvalidation {
			cloudfront_GetInvalidation(cfg, client)
			return
		}
		if _cloudfrontGetInvalidationForDistributionTenant {
			cloudfront_GetInvalidationForDistributionTenant(cfg, client)
			return
		}
		if _cloudfrontGetKeyGroup {
			cloudfront_GetKeyGroup(cfg, client)
			return
		}
		if _cloudfrontGetKeyGroupConfig {
			cloudfront_GetKeyGroupConfig(cfg, client)
			return
		}
		if _cloudfrontGetManagedCertificateDetails {
			cloudfront_GetManagedCertificateDetails(cfg, client)
			return
		}
		if _cloudfrontGetMonitoringSubscription {
			cloudfront_GetMonitoringSubscription(cfg, client)
			return
		}
		if _cloudfrontGetOriginAccessControl {
			cloudfront_GetOriginAccessControl(cfg, client)
			return
		}
		if _cloudfrontGetOriginAccessControlConfig {
			cloudfront_GetOriginAccessControlConfig(cfg, client)
			return
		}
		if _cloudfrontGetOriginRequestPolicy {
			cloudfront_GetOriginRequestPolicy(cfg, client)
			return
		}
		if _cloudfrontGetOriginRequestPolicyConfig {
			cloudfront_GetOriginRequestPolicyConfig(cfg, client)
			return
		}
		if _cloudfrontGetPublicKey {
			cloudfront_GetPublicKey(cfg, client)
			return
		}
		if _cloudfrontGetPublicKeyConfig {
			cloudfront_GetPublicKeyConfig(cfg, client)
			return
		}
		if _cloudfrontGetRealtimeLogConfig {
			cloudfront_GetRealtimeLogConfig(cfg, client)
			return
		}
		if _cloudfrontGetResourcePolicy {
			cloudfront_GetResourcePolicy(cfg, client)
			return
		}
		if _cloudfrontGetResponseHeadersPolicy {
			cloudfront_GetResponseHeadersPolicy(cfg, client)
			return
		}
		if _cloudfrontGetResponseHeadersPolicyConfig {
			cloudfront_GetResponseHeadersPolicyConfig(cfg, client)
			return
		}
		if _cloudfrontGetStreamingDistribution {
			cloudfront_GetStreamingDistribution(cfg, client)
			return
		}
		if _cloudfrontGetStreamingDistributionConfig {
			cloudfront_GetStreamingDistributionConfig(cfg, client)
			return
		}
		if _cloudfrontGetTrustStore {
			cloudfront_GetTrustStore(cfg, client)
			return
		}
		if _cloudfrontGetVpcOrigin {
			cloudfront_GetVpcOrigin(cfg, client)
			return
		}
		if _cloudfrontListAnycastIpLists {
			cloudfront_ListAnycastIpLists(cfg, client)
			return
		}
		if _cloudfrontListCachePolicies {
			cloudfront_ListCachePolicies(cfg, client)
			return
		}
		if _cloudfrontListCloudFrontOriginAccessIdentities {
			cloudfront_ListCloudFrontOriginAccessIdentities(cfg, client)
			return
		}
		if _cloudfrontListConflictingAliases {
			cloudfront_ListConflictingAliases(cfg, client)
			return
		}
		if _cloudfrontListConnectionFunctions {
			cloudfront_ListConnectionFunctions(cfg, client)
			return
		}
		if _cloudfrontListConnectionGroups {
			cloudfront_ListConnectionGroups(cfg, client)
			return
		}
		if _cloudfrontListContinuousDeploymentPolicies {
			cloudfront_ListContinuousDeploymentPolicies(cfg, client)
			return
		}
		if _cloudfrontListDistributionTenants {
			cloudfront_ListDistributionTenants(cfg, client)
			return
		}
		if _cloudfrontListDistributionTenantsByCustomization {
			cloudfront_ListDistributionTenantsByCustomization(cfg, client)
			return
		}
		if _cloudfrontListDistributions {
			cloudfront_ListDistributions(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByAnycastIpListId {
			cloudfront_ListDistributionsByAnycastIpListId(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByCachePolicyId {
			cloudfront_ListDistributionsByCachePolicyId(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByConnectionFunction {
			cloudfront_ListDistributionsByConnectionFunction(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByConnectionMode {
			cloudfront_ListDistributionsByConnectionMode(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByKeyGroup {
			cloudfront_ListDistributionsByKeyGroup(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByOriginRequestPolicyId {
			cloudfront_ListDistributionsByOriginRequestPolicyId(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByOwnedResource {
			cloudfront_ListDistributionsByOwnedResource(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByRealtimeLogConfig {
			cloudfront_ListDistributionsByRealtimeLogConfig(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByResponseHeadersPolicyId {
			cloudfront_ListDistributionsByResponseHeadersPolicyId(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByTrustStore {
			cloudfront_ListDistributionsByTrustStore(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByVpcOriginId {
			cloudfront_ListDistributionsByVpcOriginId(cfg, client)
			return
		}
		if _cloudfrontListDistributionsByWebACLId {
			cloudfront_ListDistributionsByWebACLId(cfg, client)
			return
		}
		if _cloudfrontListDomainConflicts {
			cloudfront_ListDomainConflicts(cfg, client)
			return
		}
		if _cloudfrontListFieldLevelEncryptionConfigs {
			cloudfront_ListFieldLevelEncryptionConfigs(cfg, client)
			return
		}
		if _cloudfrontListFieldLevelEncryptionProfiles {
			cloudfront_ListFieldLevelEncryptionProfiles(cfg, client)
			return
		}
		if _cloudfrontListFunctions {
			cloudfront_ListFunctions(cfg, client)
			return
		}
		if _cloudfrontListInvalidations {
			cloudfront_ListInvalidations(cfg, client)
			return
		}
		if _cloudfrontListInvalidationsForDistributionTenant {
			cloudfront_ListInvalidationsForDistributionTenant(cfg, client)
			return
		}
		if _cloudfrontListKeyGroups {
			cloudfront_ListKeyGroups(cfg, client)
			return
		}
		if _cloudfrontListKeyValueStores {
			cloudfront_ListKeyValueStores(cfg, client)
			return
		}
		if _cloudfrontListOriginAccessControls {
			cloudfront_ListOriginAccessControls(cfg, client)
			return
		}
		if _cloudfrontListOriginRequestPolicies {
			cloudfront_ListOriginRequestPolicies(cfg, client)
			return
		}
		if _cloudfrontListPublicKeys {
			cloudfront_ListPublicKeys(cfg, client)
			return
		}
		if _cloudfrontListRealtimeLogConfigs {
			cloudfront_ListRealtimeLogConfigs(cfg, client)
			return
		}
		if _cloudfrontListResponseHeadersPolicies {
			cloudfront_ListResponseHeadersPolicies(cfg, client)
			return
		}
		if _cloudfrontListStreamingDistributions {
			cloudfront_ListStreamingDistributions(cfg, client)
			return
		}
		if _cloudfrontListTagsForResource {
			cloudfront_ListTagsForResource(cfg, client)
			return
		}
		if _cloudfrontListTrustStores {
			cloudfront_ListTrustStores(cfg, client)
			return
		}
		if _cloudfrontListVpcOrigins {
			cloudfront_ListVpcOrigins(cfg, client)
			return
		}
		if _cloudfrontPublishConnectionFunction {
			cloudfront_PublishConnectionFunction(cfg, client)
			return
		}
		if _cloudfrontPublishFunction {
			cloudfront_PublishFunction(cfg, client)
			return
		}
		if _cloudfrontPutResourcePolicy {
			cloudfront_PutResourcePolicy(cfg, client)
			return
		}
		if _cloudfrontTagResource {
			cloudfront_TagResource(cfg, client)
			return
		}
		if _cloudfrontTestConnectionFunction {
			cloudfront_TestConnectionFunction(cfg, client)
			return
		}
		if _cloudfrontTestFunction {
			cloudfront_TestFunction(cfg, client)
			return
		}
		if _cloudfrontUntagResource {
			cloudfront_UntagResource(cfg, client)
			return
		}
		if _cloudfrontUpdateAnycastIpList {
			cloudfront_UpdateAnycastIpList(cfg, client)
			return
		}
		if _cloudfrontUpdateCachePolicy {
			cloudfront_UpdateCachePolicy(cfg, client)
			return
		}
		if _cloudfrontUpdateCloudFrontOriginAccessIdentity {
			cloudfront_UpdateCloudFrontOriginAccessIdentity(cfg, client)
			return
		}
		if _cloudfrontUpdateConnectionFunction {
			cloudfront_UpdateConnectionFunction(cfg, client)
			return
		}
		if _cloudfrontUpdateConnectionGroup {
			cloudfront_UpdateConnectionGroup(cfg, client)
			return
		}
		if _cloudfrontUpdateContinuousDeploymentPolicy {
			cloudfront_UpdateContinuousDeploymentPolicy(cfg, client)
			return
		}
		if _cloudfrontUpdateDistribution {
			cloudfront_UpdateDistribution(cfg, client)
			return
		}
		if _cloudfrontUpdateDistributionTenant {
			cloudfront_UpdateDistributionTenant(cfg, client)
			return
		}
		if _cloudfrontUpdateDistributionWithStagingConfig {
			cloudfront_UpdateDistributionWithStagingConfig(cfg, client)
			return
		}
		if _cloudfrontUpdateDomainAssociation {
			cloudfront_UpdateDomainAssociation(cfg, client)
			return
		}
		if _cloudfrontUpdateFieldLevelEncryptionConfig {
			cloudfront_UpdateFieldLevelEncryptionConfig(cfg, client)
			return
		}
		if _cloudfrontUpdateFieldLevelEncryptionProfile {
			cloudfront_UpdateFieldLevelEncryptionProfile(cfg, client)
			return
		}
		if _cloudfrontUpdateFunction {
			cloudfront_UpdateFunction(cfg, client)
			return
		}
		if _cloudfrontUpdateKeyGroup {
			cloudfront_UpdateKeyGroup(cfg, client)
			return
		}
		if _cloudfrontUpdateKeyValueStore {
			cloudfront_UpdateKeyValueStore(cfg, client)
			return
		}
		if _cloudfrontUpdateOriginAccessControl {
			cloudfront_UpdateOriginAccessControl(cfg, client)
			return
		}
		if _cloudfrontUpdateOriginRequestPolicy {
			cloudfront_UpdateOriginRequestPolicy(cfg, client)
			return
		}
		if _cloudfrontUpdatePublicKey {
			cloudfront_UpdatePublicKey(cfg, client)
			return
		}
		if _cloudfrontUpdateRealtimeLogConfig {
			cloudfront_UpdateRealtimeLogConfig(cfg, client)
			return
		}
		if _cloudfrontUpdateResponseHeadersPolicy {
			cloudfront_UpdateResponseHeadersPolicy(cfg, client)
			return
		}
		if _cloudfrontUpdateStreamingDistribution {
			cloudfront_UpdateStreamingDistribution(cfg, client)
			return
		}
		if _cloudfrontUpdateTrustStore {
			cloudfront_UpdateTrustStore(cfg, client)
			return
		}
		if _cloudfrontUpdateVpcOrigin {
			cloudfront_UpdateVpcOrigin(cfg, client)
			return
		}
		if _cloudfrontVerifyDnsConfiguration {
			cloudfront_VerifyDnsConfiguration(cfg, client)
			return
		}

	},
}

var (
	_cloudfrontAssociateAlias                             bool
	_cloudfrontAssociateDistributionTenantWebACL          bool
	_cloudfrontAssociateDistributionWebACL                bool
	_cloudfrontCopyDistribution                           bool
	_cloudfrontCreateAnycastIpList                        bool
	_cloudfrontCreateCachePolicy                          bool
	_cloudfrontCreateCloudFrontOriginAccessIdentity       bool
	_cloudfrontCreateConnectionFunction                   bool
	_cloudfrontCreateConnectionGroup                      bool
	_cloudfrontCreateContinuousDeploymentPolicy           bool
	_cloudfrontCreateDistribution                         bool
	_cloudfrontCreateDistributionTenant                   bool
	_cloudfrontCreateDistributionWithTags                 bool
	_cloudfrontCreateFieldLevelEncryptionConfig           bool
	_cloudfrontCreateFieldLevelEncryptionProfile          bool
	_cloudfrontCreateFunction                             bool
	_cloudfrontCreateInvalidation                         bool
	_cloudfrontCreateInvalidationForDistributionTenant    bool
	_cloudfrontCreateKeyGroup                             bool
	_cloudfrontCreateKeyValueStore                        bool
	_cloudfrontCreateMonitoringSubscription               bool
	_cloudfrontCreateOriginAccessControl                  bool
	_cloudfrontCreateOriginRequestPolicy                  bool
	_cloudfrontCreatePublicKey                            bool
	_cloudfrontCreateRealtimeLogConfig                    bool
	_cloudfrontCreateResponseHeadersPolicy                bool
	_cloudfrontCreateStreamingDistribution                bool
	_cloudfrontCreateStreamingDistributionWithTags        bool
	_cloudfrontCreateTrustStore                           bool
	_cloudfrontCreateVpcOrigin                            bool
	_cloudfrontDeleteAnycastIpList                        bool
	_cloudfrontDeleteCachePolicy                          bool
	_cloudfrontDeleteCloudFrontOriginAccessIdentity       bool
	_cloudfrontDeleteConnectionFunction                   bool
	_cloudfrontDeleteConnectionGroup                      bool
	_cloudfrontDeleteContinuousDeploymentPolicy           bool
	_cloudfrontDeleteDistribution                         bool
	_cloudfrontDeleteDistributionTenant                   bool
	_cloudfrontDeleteFieldLevelEncryptionConfig           bool
	_cloudfrontDeleteFieldLevelEncryptionProfile          bool
	_cloudfrontDeleteFunction                             bool
	_cloudfrontDeleteKeyGroup                             bool
	_cloudfrontDeleteKeyValueStore                        bool
	_cloudfrontDeleteMonitoringSubscription               bool
	_cloudfrontDeleteOriginAccessControl                  bool
	_cloudfrontDeleteOriginRequestPolicy                  bool
	_cloudfrontDeletePublicKey                            bool
	_cloudfrontDeleteRealtimeLogConfig                    bool
	_cloudfrontDeleteResourcePolicy                       bool
	_cloudfrontDeleteResponseHeadersPolicy                bool
	_cloudfrontDeleteStreamingDistribution                bool
	_cloudfrontDeleteTrustStore                           bool
	_cloudfrontDeleteVpcOrigin                            bool
	_cloudfrontDescribeConnectionFunction                 bool
	_cloudfrontDescribeFunction                           bool
	_cloudfrontDescribeKeyValueStore                      bool
	_cloudfrontDisassociateDistributionTenantWebACL       bool
	_cloudfrontDisassociateDistributionWebACL             bool
	_cloudfrontGetAnycastIpList                           bool
	_cloudfrontGetCachePolicy                             bool
	_cloudfrontGetCachePolicyConfig                       bool
	_cloudfrontGetCloudFrontOriginAccessIdentity          bool
	_cloudfrontGetCloudFrontOriginAccessIdentityConfig    bool
	_cloudfrontGetConnectionFunction                      bool
	_cloudfrontGetConnectionGroup                         bool
	_cloudfrontGetConnectionGroupByRoutingEndpoint        bool
	_cloudfrontGetContinuousDeploymentPolicy              bool
	_cloudfrontGetContinuousDeploymentPolicyConfig        bool
	_cloudfrontGetDistribution                            bool
	_cloudfrontGetDistributionConfig                      bool
	_cloudfrontGetDistributionTenant                      bool
	_cloudfrontGetDistributionTenantByDomain              bool
	_cloudfrontGetFieldLevelEncryption                    bool
	_cloudfrontGetFieldLevelEncryptionConfig              bool
	_cloudfrontGetFieldLevelEncryptionProfile             bool
	_cloudfrontGetFieldLevelEncryptionProfileConfig       bool
	_cloudfrontGetFunction                                bool
	_cloudfrontGetInvalidation                            bool
	_cloudfrontGetInvalidationForDistributionTenant       bool
	_cloudfrontGetKeyGroup                                bool
	_cloudfrontGetKeyGroupConfig                          bool
	_cloudfrontGetManagedCertificateDetails               bool
	_cloudfrontGetMonitoringSubscription                  bool
	_cloudfrontGetOriginAccessControl                     bool
	_cloudfrontGetOriginAccessControlConfig               bool
	_cloudfrontGetOriginRequestPolicy                     bool
	_cloudfrontGetOriginRequestPolicyConfig               bool
	_cloudfrontGetPublicKey                               bool
	_cloudfrontGetPublicKeyConfig                         bool
	_cloudfrontGetRealtimeLogConfig                       bool
	_cloudfrontGetResourcePolicy                          bool
	_cloudfrontGetResponseHeadersPolicy                   bool
	_cloudfrontGetResponseHeadersPolicyConfig             bool
	_cloudfrontGetStreamingDistribution                   bool
	_cloudfrontGetStreamingDistributionConfig             bool
	_cloudfrontGetTrustStore                              bool
	_cloudfrontGetVpcOrigin                               bool
	_cloudfrontListAnycastIpLists                         bool
	_cloudfrontListCachePolicies                          bool
	_cloudfrontListCloudFrontOriginAccessIdentities       bool
	_cloudfrontListConflictingAliases                     bool
	_cloudfrontListConnectionFunctions                    bool
	_cloudfrontListConnectionGroups                       bool
	_cloudfrontListContinuousDeploymentPolicies           bool
	_cloudfrontListDistributionTenants                    bool
	_cloudfrontListDistributionTenantsByCustomization     bool
	_cloudfrontListDistributions                          bool
	_cloudfrontListDistributionsByAnycastIpListId         bool
	_cloudfrontListDistributionsByCachePolicyId           bool
	_cloudfrontListDistributionsByConnectionFunction      bool
	_cloudfrontListDistributionsByConnectionMode          bool
	_cloudfrontListDistributionsByKeyGroup                bool
	_cloudfrontListDistributionsByOriginRequestPolicyId   bool
	_cloudfrontListDistributionsByOwnedResource           bool
	_cloudfrontListDistributionsByRealtimeLogConfig       bool
	_cloudfrontListDistributionsByResponseHeadersPolicyId bool
	_cloudfrontListDistributionsByTrustStore              bool
	_cloudfrontListDistributionsByVpcOriginId             bool
	_cloudfrontListDistributionsByWebACLId                bool
	_cloudfrontListDomainConflicts                        bool
	_cloudfrontListFieldLevelEncryptionConfigs            bool
	_cloudfrontListFieldLevelEncryptionProfiles           bool
	_cloudfrontListFunctions                              bool
	_cloudfrontListInvalidations                          bool
	_cloudfrontListInvalidationsForDistributionTenant     bool
	_cloudfrontListKeyGroups                              bool
	_cloudfrontListKeyValueStores                         bool
	_cloudfrontListOriginAccessControls                   bool
	_cloudfrontListOriginRequestPolicies                  bool
	_cloudfrontListPublicKeys                             bool
	_cloudfrontListRealtimeLogConfigs                     bool
	_cloudfrontListResponseHeadersPolicies                bool
	_cloudfrontListStreamingDistributions                 bool
	_cloudfrontListTagsForResource                        bool
	_cloudfrontListTrustStores                            bool
	_cloudfrontListVpcOrigins                             bool
	_cloudfrontPublishConnectionFunction                  bool
	_cloudfrontPublishFunction                            bool
	_cloudfrontPutResourcePolicy                          bool
	_cloudfrontTagResource                                bool
	_cloudfrontTestConnectionFunction                     bool
	_cloudfrontTestFunction                               bool
	_cloudfrontUntagResource                              bool
	_cloudfrontUpdateAnycastIpList                        bool
	_cloudfrontUpdateCachePolicy                          bool
	_cloudfrontUpdateCloudFrontOriginAccessIdentity       bool
	_cloudfrontUpdateConnectionFunction                   bool
	_cloudfrontUpdateConnectionGroup                      bool
	_cloudfrontUpdateContinuousDeploymentPolicy           bool
	_cloudfrontUpdateDistribution                         bool
	_cloudfrontUpdateDistributionTenant                   bool
	_cloudfrontUpdateDistributionWithStagingConfig        bool
	_cloudfrontUpdateDomainAssociation                    bool
	_cloudfrontUpdateFieldLevelEncryptionConfig           bool
	_cloudfrontUpdateFieldLevelEncryptionProfile          bool
	_cloudfrontUpdateFunction                             bool
	_cloudfrontUpdateKeyGroup                             bool
	_cloudfrontUpdateKeyValueStore                        bool
	_cloudfrontUpdateOriginAccessControl                  bool
	_cloudfrontUpdateOriginRequestPolicy                  bool
	_cloudfrontUpdatePublicKey                            bool
	_cloudfrontUpdateRealtimeLogConfig                    bool
	_cloudfrontUpdateResponseHeadersPolicy                bool
	_cloudfrontUpdateStreamingDistribution                bool
	_cloudfrontUpdateTrustStore                           bool
	_cloudfrontUpdateVpcOrigin                            bool
	_cloudfrontVerifyDnsConfiguration                     bool

	_cloudfrontAlias                                string
	_cloudfrontAnycastIpListId                      string
	_cloudfrontARN                                  string
	_cloudfrontAssociationFilter                    string
	_cloudfrontCaCertificatesBundleSource           string
	_cloudfrontCachePolicyConfig                    string
	_cloudfrontCachePolicyId                        string
	_cloudfrontCallerReference                      string
	_cloudfrontCertificateArn                       string
	_cloudfrontCloudFrontOriginAccessIdentityConfig string
	_cloudfrontComment                              string
	_cloudfrontConnectionFunctionCode               string
	_cloudfrontConnectionFunctionConfig             string
	_cloudfrontConnectionFunctionIdentifier         string
	_cloudfrontConnectionGroupId                    string
	_cloudfrontConnectionMode                       string
	_cloudfrontConnectionObject                     string
	_cloudfrontContinuousDeploymentPolicyConfig     string
	_cloudfrontCustomizations                       string
	_cloudfrontDistributionConfig                   string
	_cloudfrontDistributionConfigWithTags           string
	_cloudfrontDistributionId                       string
	_cloudfrontDistributionTenantId                 string
	_cloudfrontDomain                               string
	_cloudfrontDomainControlValidationResource      string
	_cloudfrontDomains                              string
	_cloudfrontEnabled                              string
	_cloudfrontEndPoints                            string
	_cloudfrontEventObject                          string
	_cloudfrontFieldLevelEncryptionConfig           string
	_cloudfrontFieldLevelEncryptionProfileConfig    string
	_cloudfrontFields                               []string
	_cloudfrontFunctionCode                         string
	_cloudfrontFunctionConfig                       string
	_cloudfrontId                                   string
	_cloudfrontIdentifier                           string
	_cloudfrontIfMatch                              string
	_cloudfrontImportSource                         string
	_cloudfrontInvalidationBatch                    string
	_cloudfrontIpAddressType                        string
	_cloudfrontIpCount                              string
	_cloudfrontIpamCidrConfigs                      string
	_cloudfrontIpv6Enabled                          string
	_cloudfrontKeyGroupConfig                       string
	_cloudfrontKeyGroupId                           string
	_cloudfrontManagedCertificateRequest            string
	_cloudfrontMarker                               string
	_cloudfrontMaxItems                             string
	_cloudfrontMonitoringSubscription               string
	_cloudfrontName                                 string
	_cloudfrontOriginAccessControlConfig            string
	_cloudfrontOriginRequestPolicyConfig            string
	_cloudfrontOriginRequestPolicyId                string
	_cloudfrontParameters                           string
	_cloudfrontPolicyDocument                       string
	_cloudfrontPrimaryDistributionId                string
	_cloudfrontPublicKeyConfig                      string
	_cloudfrontRealtimeLogConfigArn                 string
	_cloudfrontRealtimeLogConfigName                string
	_cloudfrontResource                             string
	_cloudfrontResourceArn                          string
	_cloudfrontResponseHeadersPolicyConfig          string
	_cloudfrontResponseHeadersPolicyId              string
	_cloudfrontRoutingEndpoint                      string
	_cloudfrontSamplingRate                         string
	_cloudfrontStage                                string
	_cloudfrontStaging                              string
	_cloudfrontStagingDistributionId                string
	_cloudfrontStatus                               string
	_cloudfrontStreamingDistributionConfig          string
	_cloudfrontStreamingDistributionConfigWithTags  string
	_cloudfrontTagKeys                              string
	_cloudfrontTags                                 string
	_cloudfrontTargetDistributionId                 string
	_cloudfrontTargetResource                       string
	_cloudfrontTrustStoreIdentifier                 string
	_cloudfrontType                                 string
	_cloudfrontVpcOriginEndpointConfig              string
	_cloudfrontVpcOriginId                          string
	_cloudfrontWebACLArn                            string
	_cloudfrontWebACLId                             string
)

// The AssociateAlias API operation only supports standard distributions. To move
// domains between distribution tenants and/or standard distributions, we recommend
// that you use the [UpdateDomainAssociation]API operation instead.
//
// Associates an alias with a CloudFront standard distribution. An alias is
// commonly known as a custom domain or vanity domain. It can also be called a
// CNAME or alternate domain name.
//
// With this operation, you can move an alias that's already used for a standard
// distribution to a different standard distribution. This prevents the downtime
// that could occur if you first remove the alias from one standard distribution
// and then separately add the alias to another standard distribution.
//
// To use this operation, specify the alias and the ID of the target standard
// distribution.
//
// For more information, including how to set up the target standard distribution,
// prerequisites that you must complete, and other restrictions, see [Moving an alternate domain name to a different standard distribution or distribution tenant]in the Amazon
// CloudFront Developer Guide.
//
// [UpdateDomainAssociation]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDomainAssociation.html
// [Moving an alternate domain name to a different standard distribution or distribution tenant]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html#alternate-domain-names-move
func cloudfront_AssociateAlias(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.AssociateAliasInput{
		// Alias: *string, // Required
		// TargetDistributionId: *string, // Required
	}

	if len(_cloudfrontAlias) > 0 {
		input.Alias = aws.String(_cloudfrontAlias)
	}
	if len(_cloudfrontTargetDistributionId) > 0 {
		input.TargetDistributionId = aws.String(_cloudfrontTargetDistributionId)
	}

	if resp, err := client.AssociateAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the WAF web ACL with a distribution tenant.
func cloudfront_AssociateDistributionTenantWebACL(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.AssociateDistributionTenantWebACLInput{
		// Id: *string, // Required
		// WebACLArn: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontWebACLArn) > 0 {
		input.WebACLArn = aws.String(_cloudfrontWebACLArn)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.AssociateDistributionTenantWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the WAF web ACL with a distribution.
func cloudfront_AssociateDistributionWebACL(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.AssociateDistributionWebACLInput{
		// Id: *string, // Required
		// WebACLArn: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontWebACLArn) > 0 {
		input.WebACLArn = aws.String(_cloudfrontWebACLArn)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.AssociateDistributionWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a staging distribution using the configuration of the provided primary
// distribution. A staging distribution is a copy of an existing distribution
// (called the primary distribution) that you can use in a continuous deployment
// workflow.
//
// After you create a staging distribution, you can use UpdateDistribution to
// modify the staging distribution's configuration. Then you can use
// CreateContinuousDeploymentPolicy to incrementally move traffic to the staging
// distribution.
//
// This API operation requires the following IAM permissions:
//
// [GetDistribution]
//
// [CreateDistribution]
//
// [CopyDistribution]
//
// [CopyDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CopyDistribution.html
// [GetDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistribution.html
// [CreateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html
func cloudfront_CopyDistribution(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CopyDistributionInput{
		// CallerReference: *string, // Required
		// PrimaryDistributionId: *string, // Required
	}

	if len(_cloudfrontCallerReference) > 0 {
		input.CallerReference = aws.String(_cloudfrontCallerReference)
	}
	if len(_cloudfrontPrimaryDistributionId) > 0 {
		input.PrimaryDistributionId = aws.String(_cloudfrontPrimaryDistributionId)
	}
	if len(_cloudfrontEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _cloudfrontEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontStaging) > 0 {
		if err := assignInputField(input, "Staging", _cloudfrontStaging); err != nil {
			log.Errorf("invalid --staging: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Anycast static IP list.
func cloudfront_CreateAnycastIpList(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateAnycastIpListInput{
		// IpCount: *int32, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontIpCount) > 0 {
		if err := assignInputField(input, "IpCount", _cloudfrontIpCount); err != nil {
			log.Errorf("invalid --ip-count: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _cloudfrontIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIpamCidrConfigs) > 0 {
		if err := assignInputField(input, "IpamCidrConfigs", _cloudfrontIpamCidrConfigs); err != nil {
			log.Errorf("invalid --ipam-cidr-configs: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudfrontTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAnycastIpList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a cache policy.
// After you create a cache policy, you can attach it to one or more cache
// behaviors. When it's attached to a cache behavior, the cache policy determines
// the following:
//
// - The values that CloudFront includes in the cache key. These values can
// include HTTP headers, cookies, and URL query strings. CloudFront uses the cache
// key to find an object in its cache that it can return to the viewer.
//
// - The default, minimum, and maximum time to live (TTL) values that you want
// objects to stay in the CloudFront cache.
//
// # If your minimum TTL is greater than 0, CloudFront will cache content for at
//
// least the duration specified in the cache policy's minimum TTL, even if the
// Cache-Control: no-cache , no-store , or private directives are present in the
// origin headers.
//
// The headers, cookies, and query strings that are included in the cache key are
// also included in requests that CloudFront sends to the origin. CloudFront sends
// a request when it can't find an object in its cache that matches the request's
// cache key. If you want to send values to the origin but not include them in the
// cache key, use OriginRequestPolicy .
//
// For more information about cache policies, see [Controlling the cache key] in the Amazon CloudFront
// Developer Guide.
//
// [Controlling the cache key]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html
func cloudfront_CreateCachePolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateCachePolicyInput{
		// CachePolicyConfig: *types.CachePolicyConfig, // Required
	}

	if len(_cloudfrontCachePolicyConfig) > 0 {
		if err := assignInputField(input, "CachePolicyConfig", _cloudfrontCachePolicyConfig); err != nil {
			log.Errorf("invalid --cache-policy-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCachePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new origin access identity. If you're using Amazon S3 for your
// origin, you can use an origin access identity to require users to access your
// content using a CloudFront URL instead of the Amazon S3 URL. For more
// information about how to use origin access identities, see [Serving Private Content through CloudFront]in the Amazon
// CloudFront Developer Guide.
//
// [Serving Private Content through CloudFront]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
func cloudfront_CreateCloudFrontOriginAccessIdentity(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateCloudFrontOriginAccessIdentityInput{
		// CloudFrontOriginAccessIdentityConfig: *types.CloudFrontOriginAccessIdentityConfig, // Required
	}

	if len(_cloudfrontCloudFrontOriginAccessIdentityConfig) > 0 {
		if err := assignInputField(input, "CloudFrontOriginAccessIdentityConfig", _cloudfrontCloudFrontOriginAccessIdentityConfig); err != nil {
			log.Errorf("invalid --cloud-front-origin-access-identity-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCloudFrontOriginAccessIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a connection function.
func cloudfront_CreateConnectionFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateConnectionFunctionInput{
		// ConnectionFunctionCode: []byte, // Required
		// ConnectionFunctionConfig: *types.FunctionConfig, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontConnectionFunctionCode) > 0 {
		if err := assignInputField(input, "ConnectionFunctionCode", _cloudfrontConnectionFunctionCode); err != nil {
			log.Errorf("invalid --connection-function-code: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontConnectionFunctionConfig) > 0 {
		if err := assignInputField(input, "ConnectionFunctionConfig", _cloudfrontConnectionFunctionConfig); err != nil {
			log.Errorf("invalid --connection-function-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudfrontTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnectionFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a connection group.
func cloudfront_CreateConnectionGroup(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateConnectionGroupInput{
		// Name: *string, // Required
	}

	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontAnycastIpListId) > 0 {
		input.AnycastIpListId = aws.String(_cloudfrontAnycastIpListId)
	}
	if len(_cloudfrontEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _cloudfrontEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIpv6Enabled) > 0 {
		if err := assignInputField(input, "Ipv6Enabled", _cloudfrontIpv6Enabled); err != nil {
			log.Errorf("invalid --ipv6-enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudfrontTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a continuous deployment policy that distributes traffic for a custom
// domain name to two different CloudFront distributions.
//
// To use a continuous deployment policy, first use CopyDistribution to create a
// staging distribution, then use UpdateDistribution to modify the staging
// distribution's configuration.
//
// After you create and update a staging distribution, you can use a continuous
// deployment policy to incrementally move traffic to the staging distribution.
// This workflow enables you to test changes to a distribution's configuration
// before moving all of your domain's production traffic to the new configuration.
func cloudfront_CreateContinuousDeploymentPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateContinuousDeploymentPolicyInput{
		// ContinuousDeploymentPolicyConfig: *types.ContinuousDeploymentPolicyConfig, // Required
	}

	if len(_cloudfrontContinuousDeploymentPolicyConfig) > 0 {
		if err := assignInputField(input, "ContinuousDeploymentPolicyConfig", _cloudfrontContinuousDeploymentPolicyConfig); err != nil {
			log.Errorf("invalid --continuous-deployment-policy-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContinuousDeploymentPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a CloudFront distribution.
func cloudfront_CreateDistribution(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateDistributionInput{
		// DistributionConfig: *types.DistributionConfig, // Required
	}

	if len(_cloudfrontDistributionConfig) > 0 {
		if err := assignInputField(input, "DistributionConfig", _cloudfrontDistributionConfig); err != nil {
			log.Errorf("invalid --distribution-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a distribution tenant.
func cloudfront_CreateDistributionTenant(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateDistributionTenantInput{
		// DistributionId: *string, // Required
		// Domains: []types.DomainItem, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontDistributionId) > 0 {
		input.DistributionId = aws.String(_cloudfrontDistributionId)
	}
	if len(_cloudfrontDomains) > 0 {
		if err := assignInputField(input, "Domains", _cloudfrontDomains); err != nil {
			log.Errorf("invalid --domains: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontConnectionGroupId) > 0 {
		input.ConnectionGroupId = aws.String(_cloudfrontConnectionGroupId)
	}
	if len(_cloudfrontCustomizations) > 0 {
		if err := assignInputField(input, "Customizations", _cloudfrontCustomizations); err != nil {
			log.Errorf("invalid --customizations: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _cloudfrontEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontManagedCertificateRequest) > 0 {
		if err := assignInputField(input, "ManagedCertificateRequest", _cloudfrontManagedCertificateRequest); err != nil {
			log.Errorf("invalid --managed-certificate-request: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cloudfrontParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudfrontTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDistributionTenant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new distribution with tags. This API operation requires the following
// IAM permissions:
//
// [CreateDistribution]
//
// [TagResource]
//
// [TagResource]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_TagResource.html
// [CreateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html
func cloudfront_CreateDistributionWithTags(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateDistributionWithTagsInput{
		// DistributionConfigWithTags: *types.DistributionConfigWithTags, // Required
	}

	if len(_cloudfrontDistributionConfigWithTags) > 0 {
		if err := assignInputField(input, "DistributionConfigWithTags", _cloudfrontDistributionConfigWithTags); err != nil {
			log.Errorf("invalid --distribution-config-with-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDistributionWithTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new field-level encryption configuration.
func cloudfront_CreateFieldLevelEncryptionConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateFieldLevelEncryptionConfigInput{
		// FieldLevelEncryptionConfig: *types.FieldLevelEncryptionConfig, // Required
	}

	if len(_cloudfrontFieldLevelEncryptionConfig) > 0 {
		if err := assignInputField(input, "FieldLevelEncryptionConfig", _cloudfrontFieldLevelEncryptionConfig); err != nil {
			log.Errorf("invalid --field-level-encryption-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFieldLevelEncryptionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a field-level encryption profile.
func cloudfront_CreateFieldLevelEncryptionProfile(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateFieldLevelEncryptionProfileInput{
		// FieldLevelEncryptionProfileConfig: *types.FieldLevelEncryptionProfileConfig, // Required
	}

	if len(_cloudfrontFieldLevelEncryptionProfileConfig) > 0 {
		if err := assignInputField(input, "FieldLevelEncryptionProfileConfig", _cloudfrontFieldLevelEncryptionProfileConfig); err != nil {
			log.Errorf("invalid --field-level-encryption-profile-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFieldLevelEncryptionProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a CloudFront function.
// To create a function, you provide the function code and some configuration
// information about the function. The response contains an Amazon Resource Name
// (ARN) that uniquely identifies the function.
//
// When you create a function, it's in the DEVELOPMENT stage. In this stage, you
// can test the function with TestFunction , and update it with UpdateFunction .
//
// When you're ready to use your function with a CloudFront distribution, use
// PublishFunction to copy the function from the DEVELOPMENT stage to LIVE . When
// it's live, you can attach the function to a distribution's cache behavior, using
// the function's ARN.
func cloudfront_CreateFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateFunctionInput{
		// FunctionCode: []byte, // Required
		// FunctionConfig: *types.FunctionConfig, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontFunctionCode) > 0 {
		if err := assignInputField(input, "FunctionCode", _cloudfrontFunctionCode); err != nil {
			log.Errorf("invalid --function-code: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontFunctionConfig) > 0 {
		if err := assignInputField(input, "FunctionConfig", _cloudfrontFunctionConfig); err != nil {
			log.Errorf("invalid --function-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}

	if resp, err := client.CreateFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new invalidation. For more information, see [Invalidating files] in the Amazon CloudFront
// Developer Guide.
//
// [Invalidating files]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Invalidation.html
func cloudfront_CreateInvalidation(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateInvalidationInput{
		// DistributionId: *string, // Required
		// InvalidationBatch: *types.InvalidationBatch, // Required
	}

	if len(_cloudfrontDistributionId) > 0 {
		input.DistributionId = aws.String(_cloudfrontDistributionId)
	}
	if len(_cloudfrontInvalidationBatch) > 0 {
		if err := assignInputField(input, "InvalidationBatch", _cloudfrontInvalidationBatch); err != nil {
			log.Errorf("invalid --invalidation-batch: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInvalidation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an invalidation for a distribution tenant. For more information, see [Invalidating files]
// in the Amazon CloudFront Developer Guide.
//
// [Invalidating files]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Invalidation.html
func cloudfront_CreateInvalidationForDistributionTenant(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateInvalidationForDistributionTenantInput{
		// Id: *string, // Required
		// InvalidationBatch: *types.InvalidationBatch, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontInvalidationBatch) > 0 {
		if err := assignInputField(input, "InvalidationBatch", _cloudfrontInvalidationBatch); err != nil {
			log.Errorf("invalid --invalidation-batch: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateInvalidationForDistributionTenant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a key group that you can use with [CloudFront signed URLs and signed cookies].
// To create a key group, you must specify at least one public key for the key
// group. After you create a key group, you can reference it from one or more cache
// behaviors. When you reference a key group in a cache behavior, CloudFront
// requires signed URLs or signed cookies for all requests that match the cache
// behavior. The URLs or cookies must be signed with a private key whose
// corresponding public key is in the key group. The signed URL or cookie contains
// information about which public key CloudFront should use to verify the
// signature. For more information, see [Serving private content]in the Amazon CloudFront Developer Guide.
//
// [Serving private content]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
// [CloudFront signed URLs and signed cookies]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
func cloudfront_CreateKeyGroup(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateKeyGroupInput{
		// KeyGroupConfig: *types.KeyGroupConfig, // Required
	}

	if len(_cloudfrontKeyGroupConfig) > 0 {
		if err := assignInputField(input, "KeyGroupConfig", _cloudfrontKeyGroupConfig); err != nil {
			log.Errorf("invalid --key-group-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKeyGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies the key value store resource to add to your account. In your account,
// the key value store names must be unique. You can also import key value store
// data in JSON format from an S3 bucket by providing a valid ImportSource that
// you own.
func cloudfront_CreateKeyValueStore(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateKeyValueStoreInput{
		// Name: *string, // Required
	}

	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontComment) > 0 {
		input.Comment = aws.String(_cloudfrontComment)
	}
	if len(_cloudfrontImportSource) > 0 {
		if err := assignInputField(input, "ImportSource", _cloudfrontImportSource); err != nil {
			log.Errorf("invalid --import-source: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKeyValueStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables additional Amazon CloudWatch metrics for the specified
// CloudFront distribution. The additional metrics incur an additional cost.
//
// For more information, see [Viewing additional CloudFront distribution metrics] in the Amazon CloudFront Developer Guide.
//
// [Viewing additional CloudFront distribution metrics]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/viewing-cloudfront-metrics.html#monitoring-console.distributions-additional
func cloudfront_CreateMonitoringSubscription(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateMonitoringSubscriptionInput{
		// DistributionId: *string, // Required
		// MonitoringSubscription: *types.MonitoringSubscription, // Required
	}

	if len(_cloudfrontDistributionId) > 0 {
		input.DistributionId = aws.String(_cloudfrontDistributionId)
	}
	if len(_cloudfrontMonitoringSubscription) > 0 {
		if err := assignInputField(input, "MonitoringSubscription", _cloudfrontMonitoringSubscription); err != nil {
			log.Errorf("invalid --monitoring-subscription: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMonitoringSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new origin access control in CloudFront. After you create an origin
// access control, you can add it to an origin in a CloudFront distribution so that
// CloudFront sends authenticated (signed) requests to the origin.
//
// This makes it possible to block public access to the origin, allowing viewers
// (users) to access the origin's content only through CloudFront.
//
// For more information about using a CloudFront origin access control, see [Restricting access to an Amazon Web Services origin] in
// the Amazon CloudFront Developer Guide.
//
// [Restricting access to an Amazon Web Services origin]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-origin.html
func cloudfront_CreateOriginAccessControl(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateOriginAccessControlInput{
		// OriginAccessControlConfig: *types.OriginAccessControlConfig, // Required
	}

	if len(_cloudfrontOriginAccessControlConfig) > 0 {
		if err := assignInputField(input, "OriginAccessControlConfig", _cloudfrontOriginAccessControlConfig); err != nil {
			log.Errorf("invalid --origin-access-control-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOriginAccessControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an origin request policy.
// After you create an origin request policy, you can attach it to one or more
// cache behaviors. When it's attached to a cache behavior, the origin request
// policy determines the values that CloudFront includes in requests that it sends
// to the origin. Each request that CloudFront sends to the origin includes the
// following:
//
// - The request body and the URL path (without the domain name) from the viewer
// request.
//
// - The headers that CloudFront automatically includes in every origin request,
// including Host , User-Agent , and X-Amz-Cf-Id .
//
// - All HTTP headers, cookies, and URL query strings that are specified in the
// cache policy or the origin request policy. These can include items from the
// viewer request and, in the case of headers, additional ones that are added by
// CloudFront.
//
// CloudFront sends a request when it can't find a valid object in its cache that
// matches the request. If you want to send values to the origin and also include
// them in the cache key, use CachePolicy .
//
// For more information about origin request policies, see [Controlling origin requests] in the Amazon
// CloudFront Developer Guide.
//
// [Controlling origin requests]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html
func cloudfront_CreateOriginRequestPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateOriginRequestPolicyInput{
		// OriginRequestPolicyConfig: *types.OriginRequestPolicyConfig, // Required
	}

	if len(_cloudfrontOriginRequestPolicyConfig) > 0 {
		if err := assignInputField(input, "OriginRequestPolicyConfig", _cloudfrontOriginRequestPolicyConfig); err != nil {
			log.Errorf("invalid --origin-request-policy-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOriginRequestPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads a public key to CloudFront that you can use with [signed URLs and signed cookies], or with [field-level encryption].
//
// [signed URLs and signed cookies]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html
// [field-level encryption]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html
func cloudfront_CreatePublicKey(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreatePublicKeyInput{
		// PublicKeyConfig: *types.PublicKeyConfig, // Required
	}

	if len(_cloudfrontPublicKeyConfig) > 0 {
		if err := assignInputField(input, "PublicKeyConfig", _cloudfrontPublicKeyConfig); err != nil {
			log.Errorf("invalid --public-key-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a real-time log configuration.
// After you create a real-time log configuration, you can attach it to one or
// more cache behaviors to send real-time log data to the specified Amazon Kinesis
// data stream.
//
// For more information about real-time log configurations, see [Real-time logs] in the Amazon
// CloudFront Developer Guide.
//
// [Real-time logs]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html
func cloudfront_CreateRealtimeLogConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateRealtimeLogConfigInput{
		// EndPoints: []types.EndPoint, // Required
		// Fields: []string, // Required
		// Name: *string, // Required
		// SamplingRate: *int64, // Required
	}

	if len(_cloudfrontEndPoints) > 0 {
		if err := assignInputField(input, "EndPoints", _cloudfrontEndPoints); err != nil {
			log.Errorf("invalid --end-points: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontFields) > 0 {
		input.Fields = append([]string(nil), _cloudfrontFields...)
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontSamplingRate) > 0 {
		if err := assignInputField(input, "SamplingRate", _cloudfrontSamplingRate); err != nil {
			log.Errorf("invalid --sampling-rate: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRealtimeLogConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a response headers policy.
// A response headers policy contains information about a set of HTTP headers. To
// create a response headers policy, you provide some metadata about the policy and
// a set of configurations that specify the headers.
//
// After you create a response headers policy, you can use its ID to attach it to
// one or more cache behaviors in a CloudFront distribution. When it's attached to
// a cache behavior, the response headers policy affects the HTTP headers that
// CloudFront includes in HTTP responses to requests that match the cache behavior.
// CloudFront adds or removes response headers according to the configuration of
// the response headers policy.
//
// For more information, see [Adding or removing HTTP headers in CloudFront responses] in the Amazon CloudFront Developer Guide.
//
// [Adding or removing HTTP headers in CloudFront responses]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html
func cloudfront_CreateResponseHeadersPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateResponseHeadersPolicyInput{
		// ResponseHeadersPolicyConfig: *types.ResponseHeadersPolicyConfig, // Required
	}

	if len(_cloudfrontResponseHeadersPolicyConfig) > 0 {
		if err := assignInputField(input, "ResponseHeadersPolicyConfig", _cloudfrontResponseHeadersPolicyConfig); err != nil {
			log.Errorf("invalid --response-headers-policy-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResponseHeadersPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is deprecated. Amazon CloudFront is deprecating real-time messaging
// protocol (RTMP) distributions on December 31, 2020. For more information, [read the announcement]on
// the Amazon CloudFront discussion forum.
//
// [read the announcement]: http://forums.aws.amazon.com/ann.jspa?annID=7356
func cloudfront_CreateStreamingDistribution(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateStreamingDistributionInput{
		// StreamingDistributionConfig: *types.StreamingDistributionConfig, // Required
	}

	if len(_cloudfrontStreamingDistributionConfig) > 0 {
		if err := assignInputField(input, "StreamingDistributionConfig", _cloudfrontStreamingDistributionConfig); err != nil {
			log.Errorf("invalid --streaming-distribution-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStreamingDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is deprecated. Amazon CloudFront is deprecating real-time messaging
// protocol (RTMP) distributions on December 31, 2020. For more information, [read the announcement]on
// the Amazon CloudFront discussion forum.
//
// [read the announcement]: http://forums.aws.amazon.com/ann.jspa?annID=7356
func cloudfront_CreateStreamingDistributionWithTags(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateStreamingDistributionWithTagsInput{
		// StreamingDistributionConfigWithTags: *types.StreamingDistributionConfigWithTags, // Required
	}

	if len(_cloudfrontStreamingDistributionConfigWithTags) > 0 {
		if err := assignInputField(input, "StreamingDistributionConfigWithTags", _cloudfrontStreamingDistributionConfigWithTags); err != nil {
			log.Errorf("invalid --streaming-distribution-config-with-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStreamingDistributionWithTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a trust store.
func cloudfront_CreateTrustStore(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateTrustStoreInput{
		// CaCertificatesBundleSource: types.CaCertificatesBundleSource, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontCaCertificatesBundleSource) > 0 {
		if err := assignInputField(input, "CaCertificatesBundleSource", _cloudfrontCaCertificatesBundleSource); err != nil {
			log.Errorf("invalid --ca-certificates-bundle-source: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudfrontTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an Amazon CloudFront VPC origin.
func cloudfront_CreateVpcOrigin(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.CreateVpcOriginInput{
		// VpcOriginEndpointConfig: *types.VpcOriginEndpointConfig, // Required
	}

	if len(_cloudfrontVpcOriginEndpointConfig) > 0 {
		if err := assignInputField(input, "VpcOriginEndpointConfig", _cloudfrontVpcOriginEndpointConfig); err != nil {
			log.Errorf("invalid --vpc-origin-endpoint-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudfrontTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVpcOrigin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Anycast static IP list.
func cloudfront_DeleteAnycastIpList(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteAnycastIpListInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteAnycastIpList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cache policy.
// You cannot delete a cache policy if it's attached to a cache behavior. First
// update your distributions to remove the cache policy from all cache behaviors,
// then delete the cache policy.
//
// To delete a cache policy, you must provide the policy's identifier and version.
// To get these values, you can use ListCachePolicies or GetCachePolicy .
func cloudfront_DeleteCachePolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteCachePolicyInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteCachePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an origin access identity.
func cloudfront_DeleteCloudFrontOriginAccessIdentity(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteCloudFrontOriginAccessIdentityInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteCloudFrontOriginAccessIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a connection function.
func cloudfront_DeleteConnectionFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteConnectionFunctionInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteConnectionFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a connection group.
func cloudfront_DeleteConnectionGroup(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteConnectionGroupInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteConnectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a continuous deployment policy.
// You cannot delete a continuous deployment policy that's attached to a primary
// distribution. First update your distribution to remove the continuous deployment
// policy, then you can delete the policy.
func cloudfront_DeleteContinuousDeploymentPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteContinuousDeploymentPolicyInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteContinuousDeploymentPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a distribution.
// Before you can delete a distribution, you must disable it, which requires
// permission to update the distribution. Once deleted, a distribution cannot be
// recovered.
func cloudfront_DeleteDistribution(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteDistributionInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a distribution tenant. If you use this API operation to delete a
// distribution tenant that is currently enabled, the request will fail.
//
// To delete a distribution tenant, you must first disable the distribution tenant
// by using the UpdateDistributionTenant API operation.
func cloudfront_DeleteDistributionTenant(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteDistributionTenantInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteDistributionTenant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove a field-level encryption configuration.
func cloudfront_DeleteFieldLevelEncryptionConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteFieldLevelEncryptionConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteFieldLevelEncryptionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove a field-level encryption profile.
func cloudfront_DeleteFieldLevelEncryptionProfile(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteFieldLevelEncryptionProfileInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteFieldLevelEncryptionProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a CloudFront function.
// You cannot delete a function if it's associated with a cache behavior. First,
// update your distributions to remove the function association from all cache
// behaviors, then delete the function.
//
// To delete a function, you must provide the function's name and version ( ETag
// value). To get these values, you can use ListFunctions and DescribeFunction .
func cloudfront_DeleteFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteFunctionInput{
		// IfMatch: *string, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}

	if resp, err := client.DeleteFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a key group.
// You cannot delete a key group that is referenced in a cache behavior. First
// update your distributions to remove the key group from all cache behaviors, then
// delete the key group.
//
// To delete a key group, you must provide the key group's identifier and version.
// To get these values, use ListKeyGroups followed by GetKeyGroup or
// GetKeyGroupConfig .
func cloudfront_DeleteKeyGroup(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteKeyGroupInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteKeyGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies the key value store to delete.
func cloudfront_DeleteKeyValueStore(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteKeyValueStoreInput{
		// IfMatch: *string, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}

	if resp, err := client.DeleteKeyValueStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables additional CloudWatch metrics for the specified CloudFront
// distribution.
func cloudfront_DeleteMonitoringSubscription(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteMonitoringSubscriptionInput{
		// DistributionId: *string, // Required
	}

	if len(_cloudfrontDistributionId) > 0 {
		input.DistributionId = aws.String(_cloudfrontDistributionId)
	}

	if resp, err := client.DeleteMonitoringSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a CloudFront origin access control.
// You cannot delete an origin access control if it's in use. First, update all
// distributions to remove the origin access control from all origins, then delete
// the origin access control.
func cloudfront_DeleteOriginAccessControl(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteOriginAccessControlInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteOriginAccessControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an origin request policy.
// You cannot delete an origin request policy if it's attached to any cache
// behaviors. First update your distributions to remove the origin request policy
// from all cache behaviors, then delete the origin request policy.
//
// To delete an origin request policy, you must provide the policy's identifier
// and version. To get the identifier, you can use ListOriginRequestPolicies or
// GetOriginRequestPolicy .
func cloudfront_DeleteOriginRequestPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteOriginRequestPolicyInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteOriginRequestPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove a public key you previously added to CloudFront.
func cloudfront_DeletePublicKey(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeletePublicKeyInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeletePublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a real-time log configuration.
// You cannot delete a real-time log configuration if it's attached to a cache
// behavior. First update your distributions to remove the real-time log
// configuration from all cache behaviors, then delete the real-time log
// configuration.
//
// To delete a real-time log configuration, you can provide the configuration's
// name or its Amazon Resource Name (ARN). You must provide at least one. If you
// provide both, CloudFront uses the name to identify the real-time log
// configuration to delete.
func cloudfront_DeleteRealtimeLogConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteRealtimeLogConfigInput{}

	if len(_cloudfrontARN) > 0 {
		input.ARN = aws.String(_cloudfrontARN)
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}

	if resp, err := client.DeleteRealtimeLogConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource policy attached to the CloudFront resource.
func cloudfront_DeleteResourcePolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_cloudfrontResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudfrontResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a response headers policy.
// You cannot delete a response headers policy if it's attached to a cache
// behavior. First update your distributions to remove the response headers policy
// from all cache behaviors, then delete the response headers policy.
//
// To delete a response headers policy, you must provide the policy's identifier
// and version. To get these values, you can use ListResponseHeadersPolicies or
// GetResponseHeadersPolicy .
func cloudfront_DeleteResponseHeadersPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteResponseHeadersPolicyInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteResponseHeadersPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a streaming distribution. To delete an RTMP distribution using the
// CloudFront API, perform the following steps.
//
// To delete an RTMP distribution using the CloudFront API:
//
// - Disable the RTMP distribution.
//
// - Submit a GET Streaming Distribution Config request to get the current
// configuration and the Etag header for the distribution.
//
// - Update the XML document that was returned in the response to your GET
// Streaming Distribution Config request to change the value of Enabled to false .
//
// - Submit a PUT Streaming Distribution Config request to update the
// configuration for your distribution. In the request body, include the XML
// document that you updated in Step 3. Then set the value of the HTTP If-Match
// header to the value of the ETag header that CloudFront returned when you
// submitted the GET Streaming Distribution Config request in Step 2.
//
// - Review the response to the PUT Streaming Distribution Config request to
// confirm that the distribution was successfully disabled.
//
// - Submit a GET Streaming Distribution Config request to confirm that your
// changes have propagated. When propagation is complete, the value of Status is
// Deployed .
//
// - Submit a DELETE Streaming Distribution request. Set the value of the HTTP
// If-Match header to the value of the ETag header that CloudFront returned when
// you submitted the GET Streaming Distribution Config request in Step 2.
//
// - Review the response to your DELETE Streaming Distribution request to confirm
// that the distribution was successfully deleted.
//
// For information about deleting a distribution using the CloudFront console, see [Deleting a Distribution]
// in the Amazon CloudFront Developer Guide.
//
// [Deleting a Distribution]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/HowToDeleteDistribution.html
func cloudfront_DeleteStreamingDistribution(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteStreamingDistributionInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteStreamingDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a trust store.
func cloudfront_DeleteTrustStore(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteTrustStoreInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an Amazon CloudFront VPC origin.
func cloudfront_DeleteVpcOrigin(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DeleteVpcOriginInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DeleteVpcOrigin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a connection function.
func cloudfront_DescribeConnectionFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DescribeConnectionFunctionInput{
		// Identifier: *string, // Required
	}

	if len(_cloudfrontIdentifier) > 0 {
		input.Identifier = aws.String(_cloudfrontIdentifier)
	}
	if len(_cloudfrontStage) > 0 {
		if err := assignInputField(input, "Stage", _cloudfrontStage); err != nil {
			log.Errorf("invalid --stage: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeConnectionFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets configuration information and metadata about a CloudFront function, but
// not the function's code. To get a function's code, use GetFunction .
//
// To get configuration information and metadata about a function, you must
// provide the function's name and stage. To get these values, you can use
// ListFunctions .
func cloudfront_DescribeFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DescribeFunctionInput{
		// Name: *string, // Required
	}

	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontStage) > 0 {
		if err := assignInputField(input, "Stage", _cloudfrontStage); err != nil {
			log.Errorf("invalid --stage: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies the key value store and its configuration.
func cloudfront_DescribeKeyValueStore(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DescribeKeyValueStoreInput{
		// Name: *string, // Required
	}

	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}

	if resp, err := client.DescribeKeyValueStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a distribution tenant from the WAF web ACL.
func cloudfront_DisassociateDistributionTenantWebACL(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DisassociateDistributionTenantWebACLInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DisassociateDistributionTenantWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a distribution from the WAF web ACL.
func cloudfront_DisassociateDistributionWebACL(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.DisassociateDistributionWebACLInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.DisassociateDistributionWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Anycast static IP list.
func cloudfront_GetAnycastIpList(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetAnycastIpListInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetAnycastIpList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a cache policy, including the following metadata:
// - The policy's identifier.
//
// - The date and time when the policy was last modified.
//
// To get a cache policy, you must provide the policy's identifier. If the cache
// policy is attached to a distribution's cache behavior, you can get the policy's
// identifier using ListDistributions or GetDistribution . If the cache policy is
// not attached to a cache behavior, you can get the identifier using
// ListCachePolicies .
func cloudfront_GetCachePolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetCachePolicyInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetCachePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a cache policy configuration.
// To get a cache policy configuration, you must provide the policy's identifier.
// If the cache policy is attached to a distribution's cache behavior, you can get
// the policy's identifier using ListDistributions or GetDistribution . If the
// cache policy is not attached to a cache behavior, you can get the identifier
// using ListCachePolicies .
func cloudfront_GetCachePolicyConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetCachePolicyConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetCachePolicyConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the information about an origin access identity.
func cloudfront_GetCloudFrontOriginAccessIdentity(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetCloudFrontOriginAccessIdentityInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetCloudFrontOriginAccessIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the configuration information about an origin access identity.
func cloudfront_GetCloudFrontOriginAccessIdentityConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetCloudFrontOriginAccessIdentityConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetCloudFrontOriginAccessIdentityConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a connection function.
func cloudfront_GetConnectionFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetConnectionFunctionInput{
		// Identifier: *string, // Required
	}

	if len(_cloudfrontIdentifier) > 0 {
		input.Identifier = aws.String(_cloudfrontIdentifier)
	}
	if len(_cloudfrontStage) > 0 {
		if err := assignInputField(input, "Stage", _cloudfrontStage); err != nil {
			log.Errorf("invalid --stage: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetConnectionFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a connection group.
func cloudfront_GetConnectionGroup(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetConnectionGroupInput{
		// Identifier: *string, // Required
	}

	if len(_cloudfrontIdentifier) > 0 {
		input.Identifier = aws.String(_cloudfrontIdentifier)
	}

	if resp, err := client.GetConnectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a connection group by using the endpoint that you
// specify.
func cloudfront_GetConnectionGroupByRoutingEndpoint(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetConnectionGroupByRoutingEndpointInput{
		// RoutingEndpoint: *string, // Required
	}

	if len(_cloudfrontRoutingEndpoint) > 0 {
		input.RoutingEndpoint = aws.String(_cloudfrontRoutingEndpoint)
	}

	if resp, err := client.GetConnectionGroupByRoutingEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a continuous deployment policy, including metadata (the policy's
// identifier and the date and time when the policy was last modified).
func cloudfront_GetContinuousDeploymentPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetContinuousDeploymentPolicyInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetContinuousDeploymentPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets configuration information about a continuous deployment policy.
func cloudfront_GetContinuousDeploymentPolicyConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetContinuousDeploymentPolicyConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetContinuousDeploymentPolicyConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the information about a distribution.
func cloudfront_GetDistribution(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetDistributionInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the configuration information about a distribution.
func cloudfront_GetDistributionConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetDistributionConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetDistributionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a distribution tenant.
func cloudfront_GetDistributionTenant(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetDistributionTenantInput{
		// Identifier: *string, // Required
	}

	if len(_cloudfrontIdentifier) > 0 {
		input.Identifier = aws.String(_cloudfrontIdentifier)
	}

	if resp, err := client.GetDistributionTenant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a distribution tenant by the associated domain.
func cloudfront_GetDistributionTenantByDomain(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetDistributionTenantByDomainInput{
		// Domain: *string, // Required
	}

	if len(_cloudfrontDomain) > 0 {
		input.Domain = aws.String(_cloudfrontDomain)
	}

	if resp, err := client.GetDistributionTenantByDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the field-level encryption configuration information.
func cloudfront_GetFieldLevelEncryption(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetFieldLevelEncryptionInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetFieldLevelEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the field-level encryption configuration information.
func cloudfront_GetFieldLevelEncryptionConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetFieldLevelEncryptionConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetFieldLevelEncryptionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the field-level encryption profile information.
func cloudfront_GetFieldLevelEncryptionProfile(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetFieldLevelEncryptionProfileInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetFieldLevelEncryptionProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the field-level encryption profile configuration information.
func cloudfront_GetFieldLevelEncryptionProfileConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetFieldLevelEncryptionProfileConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetFieldLevelEncryptionProfileConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the code of a CloudFront function. To get configuration information and
// metadata about a function, use DescribeFunction .
//
// To get a function's code, you must provide the function's name and stage. To
// get these values, you can use ListFunctions .
func cloudfront_GetFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetFunctionInput{
		// Name: *string, // Required
	}

	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontStage) > 0 {
		if err := assignInputField(input, "Stage", _cloudfrontStage); err != nil {
			log.Errorf("invalid --stage: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the information about an invalidation.
func cloudfront_GetInvalidation(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetInvalidationInput{
		// DistributionId: *string, // Required
		// Id: *string, // Required
	}

	if len(_cloudfrontDistributionId) > 0 {
		input.DistributionId = aws.String(_cloudfrontDistributionId)
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetInvalidation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specific invalidation for a distribution tenant.
func cloudfront_GetInvalidationForDistributionTenant(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetInvalidationForDistributionTenantInput{
		// DistributionTenantId: *string, // Required
		// Id: *string, // Required
	}

	if len(_cloudfrontDistributionTenantId) > 0 {
		input.DistributionTenantId = aws.String(_cloudfrontDistributionTenantId)
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetInvalidationForDistributionTenant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a key group, including the date and time when the key group was last
// modified.
//
// To get a key group, you must provide the key group's identifier. If the key
// group is referenced in a distribution's cache behavior, you can get the key
// group's identifier using ListDistributions or GetDistribution . If the key group
// is not referenced in a cache behavior, you can get the identifier using
// ListKeyGroups .
func cloudfront_GetKeyGroup(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetKeyGroupInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetKeyGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a key group configuration.
// To get a key group configuration, you must provide the key group's identifier.
// If the key group is referenced in a distribution's cache behavior, you can get
// the key group's identifier using ListDistributions or GetDistribution . If the
// key group is not referenced in a cache behavior, you can get the identifier
// using ListKeyGroups .
func cloudfront_GetKeyGroupConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetKeyGroupConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetKeyGroupConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about the CloudFront managed ACM certificate.
func cloudfront_GetManagedCertificateDetails(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetManagedCertificateDetailsInput{
		// Identifier: *string, // Required
	}

	if len(_cloudfrontIdentifier) > 0 {
		input.Identifier = aws.String(_cloudfrontIdentifier)
	}

	if resp, err := client.GetManagedCertificateDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about whether additional CloudWatch metrics are enabled for
// the specified CloudFront distribution.
func cloudfront_GetMonitoringSubscription(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetMonitoringSubscriptionInput{
		// DistributionId: *string, // Required
	}

	if len(_cloudfrontDistributionId) > 0 {
		input.DistributionId = aws.String(_cloudfrontDistributionId)
	}

	if resp, err := client.GetMonitoringSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a CloudFront origin access control, including its unique identifier.
func cloudfront_GetOriginAccessControl(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetOriginAccessControlInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetOriginAccessControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a CloudFront origin access control configuration.
func cloudfront_GetOriginAccessControlConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetOriginAccessControlConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetOriginAccessControlConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an origin request policy, including the following metadata:
// - The policy's identifier.
//
// - The date and time when the policy was last modified.
//
// To get an origin request policy, you must provide the policy's identifier. If
// the origin request policy is attached to a distribution's cache behavior, you
// can get the policy's identifier using ListDistributions or GetDistribution . If
// the origin request policy is not attached to a cache behavior, you can get the
// identifier using ListOriginRequestPolicies .
func cloudfront_GetOriginRequestPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetOriginRequestPolicyInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetOriginRequestPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an origin request policy configuration.
// To get an origin request policy configuration, you must provide the policy's
// identifier. If the origin request policy is attached to a distribution's cache
// behavior, you can get the policy's identifier using ListDistributions or
// GetDistribution . If the origin request policy is not attached to a cache
// behavior, you can get the identifier using ListOriginRequestPolicies .
func cloudfront_GetOriginRequestPolicyConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetOriginRequestPolicyConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetOriginRequestPolicyConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a public key.
func cloudfront_GetPublicKey(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetPublicKeyInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a public key configuration.
func cloudfront_GetPublicKeyConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetPublicKeyConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetPublicKeyConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a real-time log configuration.
// To get a real-time log configuration, you can provide the configuration's name
// or its Amazon Resource Name (ARN). You must provide at least one. If you provide
// both, CloudFront uses the name to identify the real-time log configuration to
// get.
func cloudfront_GetRealtimeLogConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetRealtimeLogConfigInput{}

	if len(_cloudfrontARN) > 0 {
		input.ARN = aws.String(_cloudfrontARN)
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}

	if resp, err := client.GetRealtimeLogConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource policy for the specified CloudFront resource that you
// own and have shared.
func cloudfront_GetResourcePolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_cloudfrontResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudfrontResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a response headers policy, including metadata (the policy's identifier and
// the date and time when the policy was last modified).
//
// To get a response headers policy, you must provide the policy's identifier. If
// the response headers policy is attached to a distribution's cache behavior, you
// can get the policy's identifier using ListDistributions or GetDistribution . If
// the response headers policy is not attached to a cache behavior, you can get the
// identifier using ListResponseHeadersPolicies .
func cloudfront_GetResponseHeadersPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetResponseHeadersPolicyInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetResponseHeadersPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a response headers policy configuration.
// To get a response headers policy configuration, you must provide the policy's
// identifier. If the response headers policy is attached to a distribution's cache
// behavior, you can get the policy's identifier using ListDistributions or
// GetDistribution . If the response headers policy is not attached to a cache
// behavior, you can get the identifier using ListResponseHeadersPolicies .
func cloudfront_GetResponseHeadersPolicyConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetResponseHeadersPolicyConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetResponseHeadersPolicyConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified RTMP distribution, including the
// distribution configuration.
func cloudfront_GetStreamingDistribution(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetStreamingDistributionInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetStreamingDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the configuration information about a streaming distribution.
func cloudfront_GetStreamingDistributionConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetStreamingDistributionConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetStreamingDistributionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a trust store.
func cloudfront_GetTrustStore(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetTrustStoreInput{
		// Identifier: *string, // Required
	}

	if len(_cloudfrontIdentifier) > 0 {
		input.Identifier = aws.String(_cloudfrontIdentifier)
	}

	if resp, err := client.GetTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the details of an Amazon CloudFront VPC origin.
func cloudfront_GetVpcOrigin(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.GetVpcOriginInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}

	if resp, err := client.GetVpcOrigin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists your Anycast static IP lists.
func cloudfront_ListAnycastIpLists(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListAnycastIpListsInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListAnycastIpLists(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of cache policies.
// You can optionally apply a filter to return only the managed policies created
// by Amazon Web Services, or only the custom policies created in your Amazon Web
// Services account.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListCachePolicies(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListCachePoliciesInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontType) > 0 {
		if err := assignInputField(input, "Type", _cloudfrontType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListCachePolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists origin access identities.
func cloudfront_ListCloudFrontOriginAccessIdentities(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListCloudFrontOriginAccessIdentitiesInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCloudFrontOriginAccessIdentities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput
	p := cloudfront.NewListCloudFrontOriginAccessIdentitiesPaginator(client, input)
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

// The ListConflictingAliases API operation only supports standard distributions.
// To list domain conflicts for both standard distributions and distribution
// tenants, we recommend that you use the [ListDomainConflicts]API operation instead.
//
// Gets a list of aliases that conflict or overlap with the provided alias, and
// the associated CloudFront standard distribution and Amazon Web Services accounts
// for each conflicting alias. An alias is commonly known as a custom domain or
// vanity domain. It can also be called a CNAME or alternate domain name.
//
// In the returned list, the standard distribution and account IDs are partially
// hidden, which allows you to identify the standard distribution and accounts that
// you own, and helps to protect the information of ones that you don't own.
//
// Use this operation to find aliases that are in use in CloudFront that conflict
// or overlap with the provided alias. For example, if you provide www.example.com
// as input, the returned list can include www.example.com and the overlapping
// wildcard alternate domain name ( .example.com ), if they exist. If you provide
// .example.com as input, the returned list can include *.example.com and any
// alternate domain names covered by that wildcard (for example, www.example.com ,
// test.example.com , dev.example.com , and so on), if they exist.
//
// To list conflicting aliases, specify the alias to search and the ID of a
// standard distribution in your account that has an attached TLS certificate that
// includes the provided alias. For more information, including how to set up the
// standard distribution and certificate, see [Moving an alternate domain name to a different standard distribution or distribution tenant]in the Amazon CloudFront Developer
// Guide.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
//
// [Moving an alternate domain name to a different standard distribution or distribution tenant]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html#alternate-domain-names-move
// [ListDomainConflicts]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListDomainConflicts.html
func cloudfront_ListConflictingAliases(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListConflictingAliasesInput{
		// Alias: *string, // Required
		// DistributionId: *string, // Required
	}

	if len(_cloudfrontAlias) > 0 {
		input.Alias = aws.String(_cloudfrontAlias)
	}
	if len(_cloudfrontDistributionId) > 0 {
		input.DistributionId = aws.String(_cloudfrontDistributionId)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListConflictingAliases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists connection functions.
func cloudfront_ListConnectionFunctions(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListConnectionFunctionsInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontStage) > 0 {
		if err := assignInputField(input, "Stage", _cloudfrontStage); err != nil {
			log.Errorf("invalid --stage: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListConnectionFunctions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListConnectionFunctionsOutput
	p := cloudfront.NewListConnectionFunctionsPaginator(client, input)
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

// Lists the connection groups in your Amazon Web Services account.
func cloudfront_ListConnectionGroups(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListConnectionGroupsInput{}

	if len(_cloudfrontAssociationFilter) > 0 {
		if err := assignInputField(input, "AssociationFilter", _cloudfrontAssociationFilter); err != nil {
			log.Errorf("invalid --association-filter: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListConnectionGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListConnectionGroupsOutput
	p := cloudfront.NewListConnectionGroupsPaginator(client, input)
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

// Gets a list of the continuous deployment policies in your Amazon Web Services
// account.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListContinuousDeploymentPolicies(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListContinuousDeploymentPoliciesInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListContinuousDeploymentPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the distribution tenants in your Amazon Web Services account.
func cloudfront_ListDistributionTenants(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionTenantsInput{}

	if len(_cloudfrontAssociationFilter) > 0 {
		if err := assignInputField(input, "AssociationFilter", _cloudfrontAssociationFilter); err != nil {
			log.Errorf("invalid --association-filter: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDistributionTenants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListDistributionTenantsOutput
	p := cloudfront.NewListDistributionTenantsPaginator(client, input)
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

// Lists distribution tenants by the customization that you specify.
// You must specify either the CertificateArn parameter or WebACLArn parameter,
// but not both in the same request.
func cloudfront_ListDistributionTenantsByCustomization(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionTenantsByCustomizationInput{}

	if len(_cloudfrontCertificateArn) > 0 {
		input.CertificateArn = aws.String(_cloudfrontCertificateArn)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontWebACLArn) > 0 {
		input.WebACLArn = aws.String(_cloudfrontWebACLArn)
	}

	if disablePaginator() {
		if resp, err := client.ListDistributionTenantsByCustomization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListDistributionTenantsByCustomizationOutput
	p := cloudfront.NewListDistributionTenantsByCustomizationPaginator(client, input)
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

// List CloudFront distributions.
func cloudfront_ListDistributions(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDistributions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListDistributionsOutput
	p := cloudfront.NewListDistributionsPaginator(client, input)
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

// Lists the distributions in your account that are associated with the specified
// AnycastIpListId .
func cloudfront_ListDistributionsByAnycastIpListId(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByAnycastIpListIdInput{
		// AnycastIpListId: *string, // Required
	}

	if len(_cloudfrontAnycastIpListId) > 0 {
		input.AnycastIpListId = aws.String(_cloudfrontAnycastIpListId)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDistributionsByAnycastIpListId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of distribution IDs for distributions that have a cache behavior
// that's associated with the specified cache policy.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListDistributionsByCachePolicyId(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByCachePolicyIdInput{
		// CachePolicyId: *string, // Required
	}

	if len(_cloudfrontCachePolicyId) > 0 {
		input.CachePolicyId = aws.String(_cloudfrontCachePolicyId)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDistributionsByCachePolicyId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists distributions by connection function.
func cloudfront_ListDistributionsByConnectionFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByConnectionFunctionInput{
		// ConnectionFunctionIdentifier: *string, // Required
	}

	if len(_cloudfrontConnectionFunctionIdentifier) > 0 {
		input.ConnectionFunctionIdentifier = aws.String(_cloudfrontConnectionFunctionIdentifier)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDistributionsByConnectionFunction(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListDistributionsByConnectionFunctionOutput
	p := cloudfront.NewListDistributionsByConnectionFunctionPaginator(client, input)
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

// Lists the distributions by the connection mode that you specify.
func cloudfront_ListDistributionsByConnectionMode(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByConnectionModeInput{
		// ConnectionMode: types.ConnectionMode, // Required
	}

	if len(_cloudfrontConnectionMode) > 0 {
		if err := assignInputField(input, "ConnectionMode", _cloudfrontConnectionMode); err != nil {
			log.Errorf("invalid --connection-mode: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDistributionsByConnectionMode(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListDistributionsByConnectionModeOutput
	p := cloudfront.NewListDistributionsByConnectionModePaginator(client, input)
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

// Gets a list of distribution IDs for distributions that have a cache behavior
// that references the specified key group.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListDistributionsByKeyGroup(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByKeyGroupInput{
		// KeyGroupId: *string, // Required
	}

	if len(_cloudfrontKeyGroupId) > 0 {
		input.KeyGroupId = aws.String(_cloudfrontKeyGroupId)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDistributionsByKeyGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of distribution IDs for distributions that have a cache behavior
// that's associated with the specified origin request policy.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListDistributionsByOriginRequestPolicyId(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByOriginRequestPolicyIdInput{
		// OriginRequestPolicyId: *string, // Required
	}

	if len(_cloudfrontOriginRequestPolicyId) > 0 {
		input.OriginRequestPolicyId = aws.String(_cloudfrontOriginRequestPolicyId)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDistributionsByOriginRequestPolicyId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the CloudFront distributions that are associated with the specified
// resource that you own.
func cloudfront_ListDistributionsByOwnedResource(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByOwnedResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_cloudfrontResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudfrontResourceArn)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDistributionsByOwnedResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of distributions that have a cache behavior that's associated with
// the specified real-time log configuration.
//
// You can specify the real-time log configuration by its name or its Amazon
// Resource Name (ARN). You must provide at least one. If you provide both,
// CloudFront uses the name to identify the real-time log configuration to list
// distributions for.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListDistributionsByRealtimeLogConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByRealtimeLogConfigInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontRealtimeLogConfigArn) > 0 {
		input.RealtimeLogConfigArn = aws.String(_cloudfrontRealtimeLogConfigArn)
	}
	if len(_cloudfrontRealtimeLogConfigName) > 0 {
		input.RealtimeLogConfigName = aws.String(_cloudfrontRealtimeLogConfigName)
	}

	if resp, err := client.ListDistributionsByRealtimeLogConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of distribution IDs for distributions that have a cache behavior
// that's associated with the specified response headers policy.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListDistributionsByResponseHeadersPolicyId(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByResponseHeadersPolicyIdInput{
		// ResponseHeadersPolicyId: *string, // Required
	}

	if len(_cloudfrontResponseHeadersPolicyId) > 0 {
		input.ResponseHeadersPolicyId = aws.String(_cloudfrontResponseHeadersPolicyId)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDistributionsByResponseHeadersPolicyId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists distributions by trust store.
func cloudfront_ListDistributionsByTrustStore(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByTrustStoreInput{
		// TrustStoreIdentifier: *string, // Required
	}

	if len(_cloudfrontTrustStoreIdentifier) > 0 {
		input.TrustStoreIdentifier = aws.String(_cloudfrontTrustStoreIdentifier)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDistributionsByTrustStore(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListDistributionsByTrustStoreOutput
	p := cloudfront.NewListDistributionsByTrustStorePaginator(client, input)
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

// List CloudFront distributions by their VPC origin ID.
func cloudfront_ListDistributionsByVpcOriginId(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByVpcOriginIdInput{
		// VpcOriginId: *string, // Required
	}

	if len(_cloudfrontVpcOriginId) > 0 {
		input.VpcOriginId = aws.String(_cloudfrontVpcOriginId)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDistributionsByVpcOriginId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the distributions that are associated with a specified WAF web ACL.
func cloudfront_ListDistributionsByWebACLId(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDistributionsByWebACLIdInput{
		// WebACLId: *string, // Required
	}

	if len(_cloudfrontWebACLId) > 0 {
		input.WebACLId = aws.String(_cloudfrontWebACLId)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDistributionsByWebACLId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// We recommend that you use the ListDomainConflicts API operation to check for
// domain conflicts, as it supports both standard distributions and distribution
// tenants. [ListConflictingAliases]performs similar checks but only supports standard distributions.
//
// Lists existing domain associations that conflict with the domain that you
// specify.
//
// You can use this API operation to identify potential domain conflicts when
// moving domains between standard distributions and/or distribution tenants.
// Domain conflicts must be resolved first before they can be moved.
//
// For example, if you provide www.example.com as input, the returned list can
// include www.example.com and the overlapping wildcard alternate domain name (
// .example.com ), if they exist. If you provide .example.com as input, the
// returned list can include *.example.com and any alternate domain names covered
// by that wildcard (for example, www.example.com , test.example.com ,
// dev.example.com , and so on), if they exist.
//
// To list conflicting domains, specify the following:
//
// - The domain to search for
//
// - The ID of a standard distribution or distribution tenant in your account
// that has an attached TLS certificate, which covers the specified domain
//
// For more information, including how to set up the standard distribution or
// distribution tenant, and the certificate, see [Moving an alternate domain name to a different standard distribution or distribution tenant]in the Amazon CloudFront
// Developer Guide.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
//
// [ListConflictingAliases]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListConflictingAliases.html
// [Moving an alternate domain name to a different standard distribution or distribution tenant]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html#alternate-domain-names-move
func cloudfront_ListDomainConflicts(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListDomainConflictsInput{
		// Domain: *string, // Required
		// DomainControlValidationResource: *types.DistributionResourceId, // Required
	}

	if len(_cloudfrontDomain) > 0 {
		input.Domain = aws.String(_cloudfrontDomain)
	}
	if len(_cloudfrontDomainControlValidationResource) > 0 {
		if err := assignInputField(input, "DomainControlValidationResource", _cloudfrontDomainControlValidationResource); err != nil {
			log.Errorf("invalid --domain-control-validation-resource: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDomainConflicts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListDomainConflictsOutput
	p := cloudfront.NewListDomainConflictsPaginator(client, input)
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

// List all field-level encryption configurations that have been created in
// CloudFront for this account.
func cloudfront_ListFieldLevelEncryptionConfigs(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListFieldLevelEncryptionConfigsInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListFieldLevelEncryptionConfigs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Request a list of field-level encryption profiles that have been created in
// CloudFront for this account.
func cloudfront_ListFieldLevelEncryptionProfiles(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListFieldLevelEncryptionProfilesInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListFieldLevelEncryptionProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of all CloudFront functions in your Amazon Web Services account.
// You can optionally apply a filter to return only the functions that are in the
// specified stage, either DEVELOPMENT or LIVE .
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListFunctions(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListFunctionsInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontStage) > 0 {
		if err := assignInputField(input, "Stage", _cloudfrontStage); err != nil {
			log.Errorf("invalid --stage: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListFunctions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists invalidation batches.
func cloudfront_ListInvalidations(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListInvalidationsInput{
		// DistributionId: *string, // Required
	}

	if len(_cloudfrontDistributionId) > 0 {
		input.DistributionId = aws.String(_cloudfrontDistributionId)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInvalidations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListInvalidationsOutput
	p := cloudfront.NewListInvalidationsPaginator(client, input)
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

// Lists the invalidations for a distribution tenant.
func cloudfront_ListInvalidationsForDistributionTenant(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListInvalidationsForDistributionTenantInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInvalidationsForDistributionTenant(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListInvalidationsForDistributionTenantOutput
	p := cloudfront.NewListInvalidationsForDistributionTenantPaginator(client, input)
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

// Gets a list of key groups.
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListKeyGroups(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListKeyGroupsInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListKeyGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies the key value stores to list.
func cloudfront_ListKeyValueStores(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListKeyValueStoresInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontStatus) > 0 {
		input.Status = aws.String(_cloudfrontStatus)
	}

	if disablePaginator() {
		if resp, err := client.ListKeyValueStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListKeyValueStoresOutput
	p := cloudfront.NewListKeyValueStoresPaginator(client, input)
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

// Gets the list of CloudFront origin access controls (OACs) in this Amazon Web
// Services account.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send another request that specifies the NextMarker value from the
// current response as the Marker value in the next request.
//
// If you're not using origin access controls for your Amazon Web Services
// account, the ListOriginAccessControls operation doesn't return the Items
// element in the response.
func cloudfront_ListOriginAccessControls(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListOriginAccessControlsInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOriginAccessControls(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListOriginAccessControlsOutput
	p := cloudfront.NewListOriginAccessControlsPaginator(client, input)
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

// Gets a list of origin request policies.
// You can optionally apply a filter to return only the managed policies created
// by Amazon Web Services, or only the custom policies created in your Amazon Web
// Services account.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListOriginRequestPolicies(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListOriginRequestPoliciesInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontType) > 0 {
		if err := assignInputField(input, "Type", _cloudfrontType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListOriginRequestPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all public keys that have been added to CloudFront for this account.
func cloudfront_ListPublicKeys(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListPublicKeysInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPublicKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListPublicKeysOutput
	p := cloudfront.NewListPublicKeysPaginator(client, input)
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

// Gets a list of real-time log configurations.
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListRealtimeLogConfigs(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListRealtimeLogConfigsInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListRealtimeLogConfigs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of response headers policies.
// You can optionally apply a filter to get only the managed policies created by
// Amazon Web Services, or only the custom policies created in your Amazon Web
// Services account.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func cloudfront_ListResponseHeadersPolicies(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListResponseHeadersPoliciesInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontType) > 0 {
		if err := assignInputField(input, "Type", _cloudfrontType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListResponseHeadersPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List streaming distributions.
func cloudfront_ListStreamingDistributions(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListStreamingDistributionsInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListStreamingDistributions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListStreamingDistributionsOutput
	p := cloudfront.NewListStreamingDistributionsPaginator(client, input)
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

// List tags for a CloudFront resource. For more information, see [Tagging a distribution] in the Amazon
// CloudFront Developer Guide.
//
// [Tagging a distribution]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/tagging.html
func cloudfront_ListTagsForResource(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListTagsForResourceInput{
		// Resource: *string, // Required
	}

	if len(_cloudfrontResource) > 0 {
		input.Resource = aws.String(_cloudfrontResource)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists trust stores.
func cloudfront_ListTrustStores(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListTrustStoresInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTrustStores(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfront.ListTrustStoresOutput
	p := cloudfront.NewListTrustStoresPaginator(client, input)
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

// List the CloudFront VPC origins in your account.
func cloudfront_ListVpcOrigins(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.ListVpcOriginsInput{}

	if len(_cloudfrontMarker) > 0 {
		input.Marker = aws.String(_cloudfrontMarker)
	}
	if len(_cloudfrontMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _cloudfrontMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListVpcOrigins(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Publishes a connection function.
func cloudfront_PublishConnectionFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.PublishConnectionFunctionInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.PublishConnectionFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Publishes a CloudFront function by copying the function code from the
// DEVELOPMENT stage to LIVE . This automatically updates all cache behaviors that
// are using this function to use the newly published copy in the LIVE stage.
//
// When a function is published to the LIVE stage, you can attach the function to
// a distribution's cache behavior, using the function's Amazon Resource Name
// (ARN).
//
// To publish a function, you must provide the function's name and version ( ETag
// value). To get these values, you can use ListFunctions and DescribeFunction .
func cloudfront_PublishFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.PublishFunctionInput{
		// IfMatch: *string, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}

	if resp, err := client.PublishFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a resource control policy for a given CloudFront resource.
func cloudfront_PutResourcePolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.PutResourcePolicyInput{
		// PolicyDocument: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_cloudfrontPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_cloudfrontPolicyDocument)
	}
	if len(_cloudfrontResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudfrontResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add tags to a CloudFront resource. For more information, see [Tagging a distribution] in the Amazon
// CloudFront Developer Guide.
//
// [Tagging a distribution]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/tagging.html
func cloudfront_TagResource(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.TagResourceInput{
		// Resource: *string, // Required
		// Tags: *types.Tags, // Required
	}

	if len(_cloudfrontResource) > 0 {
		input.Resource = aws.String(_cloudfrontResource)
	}
	if len(_cloudfrontTags) > 0 {
		if err := assignInputField(input, "Tags", _cloudfrontTags); err != nil {
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

// Tests a connection function.
func cloudfront_TestConnectionFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.TestConnectionFunctionInput{
		// ConnectionObject: []byte, // Required
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontConnectionObject) > 0 {
		if err := assignInputField(input, "ConnectionObject", _cloudfrontConnectionObject); err != nil {
			log.Errorf("invalid --connection-object: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontStage) > 0 {
		if err := assignInputField(input, "Stage", _cloudfrontStage); err != nil {
			log.Errorf("invalid --stage: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestConnectionFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tests a CloudFront function.
// To test a function, you provide an event object that represents an HTTP request
// or response that your CloudFront distribution could receive in production.
// CloudFront runs the function, passing it the event object that you provided, and
// returns the function's result (the modified event object) in the response. The
// response also contains function logs and error messages, if any exist. For more
// information about testing functions, see [Testing functions]in the Amazon CloudFront Developer
// Guide.
//
// To test a function, you provide the function's name and version ( ETag value)
// along with the event object. To get the function's name and version, you can use
// ListFunctions and DescribeFunction .
//
// [Testing functions]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/managing-functions.html#test-function
func cloudfront_TestFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.TestFunctionInput{
		// EventObject: []byte, // Required
		// IfMatch: *string, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontEventObject) > 0 {
		if err := assignInputField(input, "EventObject", _cloudfrontEventObject); err != nil {
			log.Errorf("invalid --event-object: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontStage) > 0 {
		if err := assignInputField(input, "Stage", _cloudfrontStage); err != nil {
			log.Errorf("invalid --stage: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove tags from a CloudFront resource. For more information, see [Tagging a distribution] in the
// Amazon CloudFront Developer Guide.
//
// [Tagging a distribution]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/tagging.html
func cloudfront_UntagResource(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UntagResourceInput{
		// Resource: *string, // Required
		// TagKeys: *types.TagKeys, // Required
	}

	if len(_cloudfrontResource) > 0 {
		input.Resource = aws.String(_cloudfrontResource)
	}
	if len(_cloudfrontTagKeys) > 0 {
		if err := assignInputField(input, "TagKeys", _cloudfrontTagKeys); err != nil {
			log.Errorf("invalid --tag-keys: %s", err.Error())
			return
		}
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Anycast static IP list.
func cloudfront_UpdateAnycastIpList(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateAnycastIpListInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _cloudfrontIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAnycastIpList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a cache policy configuration.
// When you update a cache policy configuration, all the fields are updated with
// the values provided in the request. You cannot update some fields independent of
// others. To update a cache policy configuration:
//
// - Use GetCachePolicyConfig to get the current configuration.
//
// - Locally modify the fields in the cache policy configuration that you want
// to update.
//
// - Call UpdateCachePolicy by providing the entire cache policy configuration,
// including the fields that you modified and those that you didn't.
//
// If your minimum TTL is greater than 0, CloudFront will cache content for at
// least the duration specified in the cache policy's minimum TTL, even if the
// Cache-Control: no-cache , no-store , or private directives are present in the
// origin headers.
func cloudfront_UpdateCachePolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateCachePolicyInput{
		// CachePolicyConfig: *types.CachePolicyConfig, // Required
		// Id: *string, // Required
	}

	if len(_cloudfrontCachePolicyConfig) > 0 {
		if err := assignInputField(input, "CachePolicyConfig", _cloudfrontCachePolicyConfig); err != nil {
			log.Errorf("invalid --cache-policy-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateCachePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an origin access identity.
func cloudfront_UpdateCloudFrontOriginAccessIdentity(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateCloudFrontOriginAccessIdentityInput{
		// CloudFrontOriginAccessIdentityConfig: *types.CloudFrontOriginAccessIdentityConfig, // Required
		// Id: *string, // Required
	}

	if len(_cloudfrontCloudFrontOriginAccessIdentityConfig) > 0 {
		if err := assignInputField(input, "CloudFrontOriginAccessIdentityConfig", _cloudfrontCloudFrontOriginAccessIdentityConfig); err != nil {
			log.Errorf("invalid --cloud-front-origin-access-identity-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateCloudFrontOriginAccessIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a connection function.
func cloudfront_UpdateConnectionFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateConnectionFunctionInput{
		// ConnectionFunctionCode: []byte, // Required
		// ConnectionFunctionConfig: *types.FunctionConfig, // Required
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontConnectionFunctionCode) > 0 {
		if err := assignInputField(input, "ConnectionFunctionCode", _cloudfrontConnectionFunctionCode); err != nil {
			log.Errorf("invalid --connection-function-code: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontConnectionFunctionConfig) > 0 {
		if err := assignInputField(input, "ConnectionFunctionConfig", _cloudfrontConnectionFunctionConfig); err != nil {
			log.Errorf("invalid --connection-function-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateConnectionFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a connection group.
func cloudfront_UpdateConnectionGroup(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateConnectionGroupInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontAnycastIpListId) > 0 {
		input.AnycastIpListId = aws.String(_cloudfrontAnycastIpListId)
	}
	if len(_cloudfrontEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _cloudfrontEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIpv6Enabled) > 0 {
		if err := assignInputField(input, "Ipv6Enabled", _cloudfrontIpv6Enabled); err != nil {
			log.Errorf("invalid --ipv6-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConnectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a continuous deployment policy. You can update a continuous deployment
// policy to enable or disable it, to change the percentage of traffic that it
// sends to the staging distribution, or to change the staging distribution that it
// sends traffic to.
//
// When you update a continuous deployment policy configuration, all the fields
// are updated with the values that are provided in the request. You cannot update
// some fields independent of others. To update a continuous deployment policy
// configuration:
//
// - Use GetContinuousDeploymentPolicyConfig to get the current configuration.
//
// - Locally modify the fields in the continuous deployment policy configuration
// that you want to update.
//
// - Use UpdateContinuousDeploymentPolicy , providing the entire continuous
// deployment policy configuration, including the fields that you modified and
// those that you didn't.
func cloudfront_UpdateContinuousDeploymentPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateContinuousDeploymentPolicyInput{
		// ContinuousDeploymentPolicyConfig: *types.ContinuousDeploymentPolicyConfig, // Required
		// Id: *string, // Required
	}

	if len(_cloudfrontContinuousDeploymentPolicyConfig) > 0 {
		if err := assignInputField(input, "ContinuousDeploymentPolicyConfig", _cloudfrontContinuousDeploymentPolicyConfig); err != nil {
			log.Errorf("invalid --continuous-deployment-policy-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateContinuousDeploymentPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration for a CloudFront distribution.
// The update process includes getting the current distribution configuration,
// updating it to make your changes, and then submitting an UpdateDistribution
// request to make the updates.
//
// # To update a web distribution using the CloudFront API
//
// - Use GetDistributionConfig to get the current configuration, including the
// version identifier ( ETag ).
//
// - Update the distribution configuration that was returned in the response.
// Note the following important requirements and restrictions:
//
// - You must copy the ETag field value from the response. (You'll use it for the
// IfMatch parameter in your request.) Then, remove the ETag field from the
// distribution configuration.
//
// - You can't change the value of CallerReference .
//
// - Submit an UpdateDistribution request, providing the updated distribution
// configuration. The new configuration replaces the existing configuration. The
// values that you specify in an UpdateDistribution request are not merged into
// your existing configuration. Make sure to include all fields: the ones that you
// modified and also the ones that you didn't.
func cloudfront_UpdateDistribution(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateDistributionInput{
		// DistributionConfig: *types.DistributionConfig, // Required
		// Id: *string, // Required
	}

	if len(_cloudfrontDistributionConfig) > 0 {
		if err := assignInputField(input, "DistributionConfig", _cloudfrontDistributionConfig); err != nil {
			log.Errorf("invalid --distribution-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a distribution tenant.
func cloudfront_UpdateDistributionTenant(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateDistributionTenantInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontConnectionGroupId) > 0 {
		input.ConnectionGroupId = aws.String(_cloudfrontConnectionGroupId)
	}
	if len(_cloudfrontCustomizations) > 0 {
		if err := assignInputField(input, "Customizations", _cloudfrontCustomizations); err != nil {
			log.Errorf("invalid --customizations: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontDistributionId) > 0 {
		input.DistributionId = aws.String(_cloudfrontDistributionId)
	}
	if len(_cloudfrontDomains) > 0 {
		if err := assignInputField(input, "Domains", _cloudfrontDomains); err != nil {
			log.Errorf("invalid --domains: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _cloudfrontEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontManagedCertificateRequest) > 0 {
		if err := assignInputField(input, "ManagedCertificateRequest", _cloudfrontManagedCertificateRequest); err != nil {
			log.Errorf("invalid --managed-certificate-request: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontParameters) > 0 {
		if err := assignInputField(input, "Parameters", _cloudfrontParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDistributionTenant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the staging distribution's configuration to its corresponding primary
// distribution. The primary distribution retains its Aliases (also known as
// alternate domain names or CNAMEs) and ContinuousDeploymentPolicyId value, but
// otherwise its configuration is overwritten to match the staging distribution.
//
// You can use this operation in a continuous deployment workflow after you have
// tested configuration changes on the staging distribution. After using a
// continuous deployment policy to move a portion of your domain name's traffic to
// the staging distribution and verifying that it works as intended, you can use
// this operation to copy the staging distribution's configuration to the primary
// distribution. This action will disable the continuous deployment policy and move
// your domain's traffic back to the primary distribution.
//
// This API operation requires the following IAM permissions:
//
// [GetDistribution]
//
// [UpdateDistribution]
//
// [GetDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistribution.html
// [UpdateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html
func cloudfront_UpdateDistributionWithStagingConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateDistributionWithStagingConfigInput{
		// Id: *string, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontStagingDistributionId) > 0 {
		input.StagingDistributionId = aws.String(_cloudfrontStagingDistributionId)
	}

	if resp, err := client.UpdateDistributionWithStagingConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// We recommend that you use the UpdateDomainAssociation API operation to move a
// domain association, as it supports both standard distributions and distribution
// tenants. [AssociateAlias]performs similar checks but only supports standard distributions.
//
// Moves a domain from its current standard distribution or distribution tenant to
// another one.
//
// You must first disable the source distribution (standard distribution or
// distribution tenant) and then separately call this operation to move the domain
// to another target distribution (standard distribution or distribution tenant).
//
// To use this operation, specify the domain and the ID of the target resource
// (standard distribution or distribution tenant). For more information, including
// how to set up the target resource, prerequisites that you must complete, and
// other restrictions, see [Moving an alternate domain name to a different standard distribution or distribution tenant]in the Amazon CloudFront Developer Guide.
//
// [AssociateAlias]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_AssociateAlias.html
// [Moving an alternate domain name to a different standard distribution or distribution tenant]: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/CNAMEs.html#alternate-domain-names-move
func cloudfront_UpdateDomainAssociation(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateDomainAssociationInput{
		// Domain: *string, // Required
		// TargetResource: *types.DistributionResourceId, // Required
	}

	if len(_cloudfrontDomain) > 0 {
		input.Domain = aws.String(_cloudfrontDomain)
	}
	if len(_cloudfrontTargetResource) > 0 {
		if err := assignInputField(input, "TargetResource", _cloudfrontTargetResource); err != nil {
			log.Errorf("invalid --target-resource: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a field-level encryption configuration.
func cloudfront_UpdateFieldLevelEncryptionConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateFieldLevelEncryptionConfigInput{
		// FieldLevelEncryptionConfig: *types.FieldLevelEncryptionConfig, // Required
		// Id: *string, // Required
	}

	if len(_cloudfrontFieldLevelEncryptionConfig) > 0 {
		if err := assignInputField(input, "FieldLevelEncryptionConfig", _cloudfrontFieldLevelEncryptionConfig); err != nil {
			log.Errorf("invalid --field-level-encryption-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateFieldLevelEncryptionConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a field-level encryption profile.
func cloudfront_UpdateFieldLevelEncryptionProfile(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateFieldLevelEncryptionProfileInput{
		// FieldLevelEncryptionProfileConfig: *types.FieldLevelEncryptionProfileConfig, // Required
		// Id: *string, // Required
	}

	if len(_cloudfrontFieldLevelEncryptionProfileConfig) > 0 {
		if err := assignInputField(input, "FieldLevelEncryptionProfileConfig", _cloudfrontFieldLevelEncryptionProfileConfig); err != nil {
			log.Errorf("invalid --field-level-encryption-profile-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateFieldLevelEncryptionProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a CloudFront function.
// You can update a function's code or the comment that describes the function.
// You cannot update a function's name.
//
// To update a function, you provide the function's name and version ( ETag value)
// along with the updated function code. To get the name and version, you can use
// ListFunctions and DescribeFunction .
func cloudfront_UpdateFunction(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateFunctionInput{
		// FunctionCode: []byte, // Required
		// FunctionConfig: *types.FunctionConfig, // Required
		// IfMatch: *string, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontFunctionCode) > 0 {
		if err := assignInputField(input, "FunctionCode", _cloudfrontFunctionCode); err != nil {
			log.Errorf("invalid --function-code: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontFunctionConfig) > 0 {
		if err := assignInputField(input, "FunctionConfig", _cloudfrontFunctionConfig); err != nil {
			log.Errorf("invalid --function-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}

	if resp, err := client.UpdateFunction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a key group.
// When you update a key group, all the fields are updated with the values
// provided in the request. You cannot update some fields independent of others. To
// update a key group:
//
// - Get the current key group with GetKeyGroup or GetKeyGroupConfig .
//
// - Locally modify the fields in the key group that you want to update. For
// example, add or remove public key IDs.
//
// - Call UpdateKeyGroup with the entire key group object, including the fields
// that you modified and those that you didn't.
func cloudfront_UpdateKeyGroup(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateKeyGroupInput{
		// Id: *string, // Required
		// KeyGroupConfig: *types.KeyGroupConfig, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontKeyGroupConfig) > 0 {
		if err := assignInputField(input, "KeyGroupConfig", _cloudfrontKeyGroupConfig); err != nil {
			log.Errorf("invalid --key-group-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateKeyGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies the key value store to update.
func cloudfront_UpdateKeyValueStore(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateKeyValueStoreInput{
		// Comment: *string, // Required
		// IfMatch: *string, // Required
		// Name: *string, // Required
	}

	if len(_cloudfrontComment) > 0 {
		input.Comment = aws.String(_cloudfrontComment)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}

	if resp, err := client.UpdateKeyValueStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a CloudFront origin access control.
func cloudfront_UpdateOriginAccessControl(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateOriginAccessControlInput{
		// Id: *string, // Required
		// OriginAccessControlConfig: *types.OriginAccessControlConfig, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontOriginAccessControlConfig) > 0 {
		if err := assignInputField(input, "OriginAccessControlConfig", _cloudfrontOriginAccessControlConfig); err != nil {
			log.Errorf("invalid --origin-access-control-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateOriginAccessControl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an origin request policy configuration.
// When you update an origin request policy configuration, all the fields are
// updated with the values provided in the request. You cannot update some fields
// independent of others. To update an origin request policy configuration:
//
// - Use GetOriginRequestPolicyConfig to get the current configuration.
//
// - Locally modify the fields in the origin request policy configuration that
// you want to update.
//
// - Call UpdateOriginRequestPolicy by providing the entire origin request policy
// configuration, including the fields that you modified and those that you didn't.
func cloudfront_UpdateOriginRequestPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateOriginRequestPolicyInput{
		// Id: *string, // Required
		// OriginRequestPolicyConfig: *types.OriginRequestPolicyConfig, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontOriginRequestPolicyConfig) > 0 {
		if err := assignInputField(input, "OriginRequestPolicyConfig", _cloudfrontOriginRequestPolicyConfig); err != nil {
			log.Errorf("invalid --origin-request-policy-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateOriginRequestPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update public key information. Note that the only value you can change is the
// comment.
func cloudfront_UpdatePublicKey(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdatePublicKeyInput{
		// Id: *string, // Required
		// PublicKeyConfig: *types.PublicKeyConfig, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontPublicKeyConfig) > 0 {
		if err := assignInputField(input, "PublicKeyConfig", _cloudfrontPublicKeyConfig); err != nil {
			log.Errorf("invalid --public-key-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdatePublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a real-time log configuration.
// When you update a real-time log configuration, all the parameters are updated
// with the values provided in the request. You cannot update some parameters
// independent of others. To update a real-time log configuration:
//
// - Call GetRealtimeLogConfig to get the current real-time log configuration.
//
// - Locally modify the parameters in the real-time log configuration that you
// want to update.
//
// - Call this API ( UpdateRealtimeLogConfig ) by providing the entire real-time
// log configuration, including the parameters that you modified and those that you
// didn't.
//
// You cannot update a real-time log configuration's Name or ARN .
func cloudfront_UpdateRealtimeLogConfig(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateRealtimeLogConfigInput{}

	if len(_cloudfrontARN) > 0 {
		input.ARN = aws.String(_cloudfrontARN)
	}
	if len(_cloudfrontEndPoints) > 0 {
		if err := assignInputField(input, "EndPoints", _cloudfrontEndPoints); err != nil {
			log.Errorf("invalid --end-points: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontFields) > 0 {
		input.Fields = append([]string(nil), _cloudfrontFields...)
	}
	if len(_cloudfrontName) > 0 {
		input.Name = aws.String(_cloudfrontName)
	}
	if len(_cloudfrontSamplingRate) > 0 {
		if err := assignInputField(input, "SamplingRate", _cloudfrontSamplingRate); err != nil {
			log.Errorf("invalid --sampling-rate: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRealtimeLogConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a response headers policy.
// When you update a response headers policy, the entire policy is replaced. You
// cannot update some policy fields independent of others. To update a response
// headers policy configuration:
//
// - Use GetResponseHeadersPolicyConfig to get the current policy's configuration.
//
// - Modify the fields in the response headers policy configuration that you
// want to update.
//
// - Call UpdateResponseHeadersPolicy , providing the entire response headers
// policy configuration, including the fields that you modified and those that you
// didn't.
func cloudfront_UpdateResponseHeadersPolicy(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateResponseHeadersPolicyInput{
		// Id: *string, // Required
		// ResponseHeadersPolicyConfig: *types.ResponseHeadersPolicyConfig, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontResponseHeadersPolicyConfig) > 0 {
		if err := assignInputField(input, "ResponseHeadersPolicyConfig", _cloudfrontResponseHeadersPolicyConfig); err != nil {
			log.Errorf("invalid --response-headers-policy-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateResponseHeadersPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a streaming distribution.
func cloudfront_UpdateStreamingDistribution(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateStreamingDistributionInput{
		// Id: *string, // Required
		// StreamingDistributionConfig: *types.StreamingDistributionConfig, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontStreamingDistributionConfig) > 0 {
		if err := assignInputField(input, "StreamingDistributionConfig", _cloudfrontStreamingDistributionConfig); err != nil {
			log.Errorf("invalid --streaming-distribution-config: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateStreamingDistribution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a trust store.
func cloudfront_UpdateTrustStore(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateTrustStoreInput{
		// CaCertificatesBundleSource: types.CaCertificatesBundleSource, // Required
		// Id: *string, // Required
		// IfMatch: *string, // Required
	}

	if len(_cloudfrontCaCertificatesBundleSource) > 0 {
		if err := assignInputField(input, "CaCertificatesBundleSource", _cloudfrontCaCertificatesBundleSource); err != nil {
			log.Errorf("invalid --ca-certificates-bundle-source: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}

	if resp, err := client.UpdateTrustStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an Amazon CloudFront VPC origin in your account.
func cloudfront_UpdateVpcOrigin(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.UpdateVpcOriginInput{
		// Id: *string, // Required
		// IfMatch: *string, // Required
		// VpcOriginEndpointConfig: *types.VpcOriginEndpointConfig, // Required
	}

	if len(_cloudfrontId) > 0 {
		input.Id = aws.String(_cloudfrontId)
	}
	if len(_cloudfrontIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontIfMatch)
	}
	if len(_cloudfrontVpcOriginEndpointConfig) > 0 {
		if err := assignInputField(input, "VpcOriginEndpointConfig", _cloudfrontVpcOriginEndpointConfig); err != nil {
			log.Errorf("invalid --vpc-origin-endpoint-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateVpcOrigin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verify the DNS configuration for your domain names. This API operation checks
// whether your domain name points to the correct routing endpoint of the
// connection group, such as d111111abcdef8.cloudfront.net. You can use this API
// operation to troubleshoot and resolve DNS configuration issues.
func cloudfront_VerifyDnsConfiguration(cfg aws.Config, client *cloudfront.Client) {
	input := &cloudfront.VerifyDnsConfigurationInput{
		// Identifier: *string, // Required
	}

	if len(_cloudfrontIdentifier) > 0 {
		input.Identifier = aws.String(_cloudfrontIdentifier)
	}
	if len(_cloudfrontDomain) > 0 {
		input.Domain = aws.String(_cloudfrontDomain)
	}

	if resp, err := client.VerifyDnsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudfrontCmd)
	_cloudfrontCmd.Flags().SortFlags = false

	_cloudfrontCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cloudfrontCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudfrontCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontAlias, "alias", "", "", "Alias")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontAnycastIpListId, "anycast-ip-list-id", "", "", "Anycast IP List ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontARN, "arn", "", "", "ARN")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontAssociationFilter, "association-filter", "", "", "Association Filter")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontCaCertificatesBundleSource, "ca-certificates-bundle-source", "", "", "Ca Certificates Bundle Source")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontCachePolicyConfig, "cache-policy-config", "", "", "Cache Policy Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontCachePolicyId, "cache-policy-id", "", "", "Cache Policy ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontCallerReference, "caller-reference", "", "", "Caller Reference")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontCertificateArn, "certificate-arn", "", "", "Certificate ARN")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontCloudFrontOriginAccessIdentityConfig, "cloud-front-origin-access-identity-config", "", "", "Cloud Front Origin Access Identity Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontComment, "comment", "", "", "Comment")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontConnectionFunctionCode, "connection-function-code", "", "", "Connection Function Code")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontConnectionFunctionConfig, "connection-function-config", "", "", "Connection Function Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontConnectionFunctionIdentifier, "connection-function-identifier", "", "", "Connection Function Identifier")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontConnectionGroupId, "connection-group-id", "", "", "Connection Group ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontConnectionMode, "connection-mode", "", "", "Connection Mode")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontConnectionObject, "connection-object", "", "", "Connection Object")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontContinuousDeploymentPolicyConfig, "continuous-deployment-policy-config", "", "", "Continuous Deployment Policy Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontCustomizations, "customizations", "", "", "Customizations")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontDistributionConfig, "distribution-config", "", "", "Distribution Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontDistributionConfigWithTags, "distribution-config-with-tags", "", "", "Distribution Config With Tags")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontDistributionId, "distribution-id", "", "", "Distribution ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontDistributionTenantId, "distribution-tenant-id", "", "", "Distribution Tenant ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontDomain, "domain", "", "", "Domain")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontDomainControlValidationResource, "domain-control-validation-resource", "", "", "Domain Control Validation Resource")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontDomains, "domains", "", "", "Domains")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontEnabled, "enabled", "", "", "Enabled")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontEndPoints, "end-points", "", "", "End Points")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontEventObject, "event-object", "", "", "Event Object")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontFieldLevelEncryptionConfig, "field-level-encryption-config", "", "", "Field Level Encryption Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontFieldLevelEncryptionProfileConfig, "field-level-encryption-profile-config", "", "", "Field Level Encryption Profile Config")
	_cloudfrontCmd.Flags().StringSliceVarP(&_cloudfrontFields, "fields", "", nil, "Fields")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontFunctionCode, "function-code", "", "", "Function Code")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontFunctionConfig, "function-config", "", "", "Function Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontId, "id", "", "", "ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontIdentifier, "identifier", "", "", "Identifier")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontIfMatch, "if-match", "", "", "If Match")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontImportSource, "import-source", "", "", "Import Source")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontInvalidationBatch, "invalidation-batch", "", "", "Invalidation Batch")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontIpAddressType, "ip-address-type", "", "", "IP Address Type")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontIpCount, "ip-count", "", "", "IP Count")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontIpamCidrConfigs, "ipam-cidr-configs", "", "", "Ipam CIDR Configs")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontIpv6Enabled, "ipv6-enabled", "", "", "IPV6 Enabled")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontKeyGroupConfig, "key-group-config", "", "", "Key Group Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontKeyGroupId, "key-group-id", "", "", "Key Group ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontManagedCertificateRequest, "managed-certificate-request", "", "", "Managed Certificate Request")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontMarker, "marker", "", "", "Marker")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontMaxItems, "max-items", "", "", "Max Items")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontMonitoringSubscription, "monitoring-subscription", "", "", "Monitoring Subscription")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontName, "name", "", "", "Name")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontOriginAccessControlConfig, "origin-access-control-config", "", "", "Origin Access Control Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontOriginRequestPolicyConfig, "origin-request-policy-config", "", "", "Origin Request Policy Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontOriginRequestPolicyId, "origin-request-policy-id", "", "", "Origin Request Policy ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontParameters, "parameters", "", "", "Parameters")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontPolicyDocument, "policy-document", "", "", "Policy Document")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontPrimaryDistributionId, "primary-distribution-id", "", "", "Primary Distribution ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontPublicKeyConfig, "public-key-config", "", "", "Public Key Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontRealtimeLogConfigArn, "realtime-log-config-arn", "", "", "Realtime Log Config ARN")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontRealtimeLogConfigName, "realtime-log-config-name", "", "", "Realtime Log Config Name")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontResource, "resource", "", "", "Resource")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontResourceArn, "resource-arn", "", "", "Resource ARN")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontResponseHeadersPolicyConfig, "response-headers-policy-config", "", "", "Response Headers Policy Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontResponseHeadersPolicyId, "response-headers-policy-id", "", "", "Response Headers Policy ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontRoutingEndpoint, "routing-endpoint", "", "", "Routing Endpoint")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontSamplingRate, "sampling-rate", "", "", "Sampling Rate")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontStage, "stage", "", "", "Stage")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontStaging, "staging", "", "", "Staging")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontStagingDistributionId, "staging-distribution-id", "", "", "Staging Distribution ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontStatus, "status", "", "", "Status")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontStreamingDistributionConfig, "streaming-distribution-config", "", "", "Streaming Distribution Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontStreamingDistributionConfigWithTags, "streaming-distribution-config-with-tags", "", "", "Streaming Distribution Config With Tags")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontTagKeys, "tag-keys", "", "", "Tag Keys")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontTags, "tags", "", "", "Tags")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontTargetDistributionId, "target-distribution-id", "", "", "Target Distribution ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontTargetResource, "target-resource", "", "", "Target Resource")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontTrustStoreIdentifier, "trust-store-identifier", "", "", "Trust Store Identifier")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontType, "type", "", "", "Type")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontVpcOriginEndpointConfig, "vpc-origin-endpoint-config", "", "", "VPC Origin Endpoint Config")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontVpcOriginId, "vpc-origin-id", "", "", "VPC Origin ID")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontWebACLArn, "web-acl-arn", "", "", "Web ACL ARN")
	_cloudfrontCmd.Flags().StringVarP(&_cloudfrontWebACLId, "web-aclid", "", "", "Web Aclid")

	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontAssociateAlias, "associate-alias", "", false, "Associate Alias")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontAssociateDistributionTenantWebACL, "associate-distribution-tenant-web-acl", "", false, "Associate Distribution Tenant Web ACL")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontAssociateDistributionWebACL, "associate-distribution-web-acl", "", false, "Associate Distribution Web ACL")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCopyDistribution, "copy-distribution", "", false, "Copy Distribution")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateAnycastIpList, "create-anycast-ip-list", "", false, "Create Anycast IP List")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateCachePolicy, "create-cache-policy", "", false, "Create Cache Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateCloudFrontOriginAccessIdentity, "create-cloud-front-origin-access-identity", "", false, "Create Cloud Front Origin Access Identity")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateConnectionFunction, "create-connection-function", "", false, "Create Connection Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateConnectionGroup, "create-connection-group", "", false, "Create Connection Group")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateContinuousDeploymentPolicy, "create-continuous-deployment-policy", "", false, "Create Continuous Deployment Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateDistribution, "create-distribution", "", false, "Create Distribution")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateDistributionTenant, "create-distribution-tenant", "", false, "Create Distribution Tenant")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateDistributionWithTags, "create-distribution-with-tags", "", false, "Create Distribution With Tags")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateFieldLevelEncryptionConfig, "create-field-level-encryption-config", "", false, "Create Field Level Encryption Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateFieldLevelEncryptionProfile, "create-field-level-encryption-profile", "", false, "Create Field Level Encryption Profile")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateFunction, "create-function", "", false, "Create Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateInvalidation, "create-invalidation", "", false, "Create Invalidation")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateInvalidationForDistributionTenant, "create-invalidation-for-distribution-tenant", "", false, "Create Invalidation For Distribution Tenant")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateKeyGroup, "create-key-group", "", false, "Create Key Group")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateKeyValueStore, "create-key-value-store", "", false, "Create Key Value Store")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateMonitoringSubscription, "create-monitoring-subscription", "", false, "Create Monitoring Subscription")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateOriginAccessControl, "create-origin-access-control", "", false, "Create Origin Access Control")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateOriginRequestPolicy, "create-origin-request-policy", "", false, "Create Origin Request Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreatePublicKey, "create-public-key", "", false, "Create Public Key")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateRealtimeLogConfig, "create-realtime-log-config", "", false, "Create Realtime Log Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateResponseHeadersPolicy, "create-response-headers-policy", "", false, "Create Response Headers Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateStreamingDistribution, "create-streaming-distribution", "", false, "Create Streaming Distribution")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateStreamingDistributionWithTags, "create-streaming-distribution-with-tags", "", false, "Create Streaming Distribution With Tags")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateTrustStore, "create-trust-store", "", false, "Create Trust Store")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontCreateVpcOrigin, "create-vpc-origin", "", false, "Create VPC Origin")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteAnycastIpList, "delete-anycast-ip-list", "", false, "Delete Anycast IP List")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteCachePolicy, "delete-cache-policy", "", false, "Delete Cache Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteCloudFrontOriginAccessIdentity, "delete-cloud-front-origin-access-identity", "", false, "Delete Cloud Front Origin Access Identity")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteConnectionFunction, "delete-connection-function", "", false, "Delete Connection Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteConnectionGroup, "delete-connection-group", "", false, "Delete Connection Group")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteContinuousDeploymentPolicy, "delete-continuous-deployment-policy", "", false, "Delete Continuous Deployment Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteDistribution, "delete-distribution", "", false, "Delete Distribution")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteDistributionTenant, "delete-distribution-tenant", "", false, "Delete Distribution Tenant")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteFieldLevelEncryptionConfig, "delete-field-level-encryption-config", "", false, "Delete Field Level Encryption Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteFieldLevelEncryptionProfile, "delete-field-level-encryption-profile", "", false, "Delete Field Level Encryption Profile")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteFunction, "delete-function", "", false, "Delete Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteKeyGroup, "delete-key-group", "", false, "Delete Key Group")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteKeyValueStore, "delete-key-value-store", "", false, "Delete Key Value Store")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteMonitoringSubscription, "delete-monitoring-subscription", "", false, "Delete Monitoring Subscription")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteOriginAccessControl, "delete-origin-access-control", "", false, "Delete Origin Access Control")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteOriginRequestPolicy, "delete-origin-request-policy", "", false, "Delete Origin Request Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeletePublicKey, "delete-public-key", "", false, "Delete Public Key")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteRealtimeLogConfig, "delete-realtime-log-config", "", false, "Delete Realtime Log Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteResponseHeadersPolicy, "delete-response-headers-policy", "", false, "Delete Response Headers Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteStreamingDistribution, "delete-streaming-distribution", "", false, "Delete Streaming Distribution")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteTrustStore, "delete-trust-store", "", false, "Delete Trust Store")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDeleteVpcOrigin, "delete-vpc-origin", "", false, "Delete VPC Origin")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDescribeConnectionFunction, "describe-connection-function", "", false, "Describe Connection Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDescribeFunction, "describe-function", "", false, "Describe Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDescribeKeyValueStore, "describe-key-value-store", "", false, "Describe Key Value Store")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDisassociateDistributionTenantWebACL, "disassociate-distribution-tenant-web-acl", "", false, "Disassociate Distribution Tenant Web ACL")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontDisassociateDistributionWebACL, "disassociate-distribution-web-acl", "", false, "Disassociate Distribution Web ACL")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetAnycastIpList, "get-anycast-ip-list", "", false, "Get Anycast IP List")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetCachePolicy, "get-cache-policy", "", false, "Get Cache Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetCachePolicyConfig, "get-cache-policy-config", "", false, "Get Cache Policy Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetCloudFrontOriginAccessIdentity, "get-cloud-front-origin-access-identity", "", false, "Get Cloud Front Origin Access Identity")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetCloudFrontOriginAccessIdentityConfig, "get-cloud-front-origin-access-identity-config", "", false, "Get Cloud Front Origin Access Identity Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetConnectionFunction, "get-connection-function", "", false, "Get Connection Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetConnectionGroup, "get-connection-group", "", false, "Get Connection Group")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetConnectionGroupByRoutingEndpoint, "get-connection-group-by-routing-endpoint", "", false, "Get Connection Group By Routing Endpoint")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetContinuousDeploymentPolicy, "get-continuous-deployment-policy", "", false, "Get Continuous Deployment Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetContinuousDeploymentPolicyConfig, "get-continuous-deployment-policy-config", "", false, "Get Continuous Deployment Policy Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetDistribution, "get-distribution", "", false, "Get Distribution")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetDistributionConfig, "get-distribution-config", "", false, "Get Distribution Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetDistributionTenant, "get-distribution-tenant", "", false, "Get Distribution Tenant")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetDistributionTenantByDomain, "get-distribution-tenant-by-domain", "", false, "Get Distribution Tenant By Domain")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetFieldLevelEncryption, "get-field-level-encryption", "", false, "Get Field Level Encryption")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetFieldLevelEncryptionConfig, "get-field-level-encryption-config", "", false, "Get Field Level Encryption Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetFieldLevelEncryptionProfile, "get-field-level-encryption-profile", "", false, "Get Field Level Encryption Profile")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetFieldLevelEncryptionProfileConfig, "get-field-level-encryption-profile-config", "", false, "Get Field Level Encryption Profile Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetFunction, "get-function", "", false, "Get Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetInvalidation, "get-invalidation", "", false, "Get Invalidation")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetInvalidationForDistributionTenant, "get-invalidation-for-distribution-tenant", "", false, "Get Invalidation For Distribution Tenant")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetKeyGroup, "get-key-group", "", false, "Get Key Group")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetKeyGroupConfig, "get-key-group-config", "", false, "Get Key Group Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetManagedCertificateDetails, "get-managed-certificate-details", "", false, "Get Managed Certificate Details")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetMonitoringSubscription, "get-monitoring-subscription", "", false, "Get Monitoring Subscription")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetOriginAccessControl, "get-origin-access-control", "", false, "Get Origin Access Control")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetOriginAccessControlConfig, "get-origin-access-control-config", "", false, "Get Origin Access Control Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetOriginRequestPolicy, "get-origin-request-policy", "", false, "Get Origin Request Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetOriginRequestPolicyConfig, "get-origin-request-policy-config", "", false, "Get Origin Request Policy Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetPublicKey, "get-public-key", "", false, "Get Public Key")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetPublicKeyConfig, "get-public-key-config", "", false, "Get Public Key Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetRealtimeLogConfig, "get-realtime-log-config", "", false, "Get Realtime Log Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetResponseHeadersPolicy, "get-response-headers-policy", "", false, "Get Response Headers Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetResponseHeadersPolicyConfig, "get-response-headers-policy-config", "", false, "Get Response Headers Policy Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetStreamingDistribution, "get-streaming-distribution", "", false, "Get Streaming Distribution")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetStreamingDistributionConfig, "get-streaming-distribution-config", "", false, "Get Streaming Distribution Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetTrustStore, "get-trust-store", "", false, "Get Trust Store")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontGetVpcOrigin, "get-vpc-origin", "", false, "Get VPC Origin")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListAnycastIpLists, "list-anycast-ip-lists", "", false, "List Anycast IP Lists")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListCachePolicies, "list-cache-policies", "", false, "List Cache Policies")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListCloudFrontOriginAccessIdentities, "list-cloud-front-origin-access-identities", "", false, "List Cloud Front Origin Access Identities")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListConflictingAliases, "list-conflicting-aliases", "", false, "List Conflicting Aliases")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListConnectionFunctions, "list-connection-functions", "", false, "List Connection Functions")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListConnectionGroups, "list-connection-groups", "", false, "List Connection Groups")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListContinuousDeploymentPolicies, "list-continuous-deployment-policies", "", false, "List Continuous Deployment Policies")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionTenants, "list-distribution-tenants", "", false, "List Distribution Tenants")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionTenantsByCustomization, "list-distribution-tenants-by-customization", "", false, "List Distribution Tenants By Customization")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributions, "list-distributions", "", false, "List Distributions")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByAnycastIpListId, "list-distributions-by-anycast-ip-list-id", "", false, "List Distributions By Anycast IP List ID")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByCachePolicyId, "list-distributions-by-cache-policy-id", "", false, "List Distributions By Cache Policy ID")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByConnectionFunction, "list-distributions-by-connection-function", "", false, "List Distributions By Connection Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByConnectionMode, "list-distributions-by-connection-mode", "", false, "List Distributions By Connection Mode")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByKeyGroup, "list-distributions-by-key-group", "", false, "List Distributions By Key Group")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByOriginRequestPolicyId, "list-distributions-by-origin-request-policy-id", "", false, "List Distributions By Origin Request Policy ID")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByOwnedResource, "list-distributions-by-owned-resource", "", false, "List Distributions By Owned Resource")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByRealtimeLogConfig, "list-distributions-by-realtime-log-config", "", false, "List Distributions By Realtime Log Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByResponseHeadersPolicyId, "list-distributions-by-response-headers-policy-id", "", false, "List Distributions By Response Headers Policy ID")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByTrustStore, "list-distributions-by-trust-store", "", false, "List Distributions By Trust Store")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByVpcOriginId, "list-distributions-by-vpc-origin-id", "", false, "List Distributions By VPC Origin ID")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDistributionsByWebACLId, "list-distributions-by-web-aclid", "", false, "List Distributions By Web Aclid")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListDomainConflicts, "list-domain-conflicts", "", false, "List Domain Conflicts")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListFieldLevelEncryptionConfigs, "list-field-level-encryption-configs", "", false, "List Field Level Encryption Configs")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListFieldLevelEncryptionProfiles, "list-field-level-encryption-profiles", "", false, "List Field Level Encryption Profiles")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListFunctions, "list-functions", "", false, "List Functions")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListInvalidations, "list-invalidations", "", false, "List Invalidations")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListInvalidationsForDistributionTenant, "list-invalidations-for-distribution-tenant", "", false, "List Invalidations For Distribution Tenant")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListKeyGroups, "list-key-groups", "", false, "List Key Groups")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListKeyValueStores, "list-key-value-stores", "", false, "List Key Value Stores")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListOriginAccessControls, "list-origin-access-controls", "", false, "List Origin Access Controls")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListOriginRequestPolicies, "list-origin-request-policies", "", false, "List Origin Request Policies")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListPublicKeys, "list-public-keys", "", false, "List Public Keys")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListRealtimeLogConfigs, "list-realtime-log-configs", "", false, "List Realtime Log Configs")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListResponseHeadersPolicies, "list-response-headers-policies", "", false, "List Response Headers Policies")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListStreamingDistributions, "list-streaming-distributions", "", false, "List Streaming Distributions")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListTrustStores, "list-trust-stores", "", false, "List Trust Stores")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontListVpcOrigins, "list-vpc-origins", "", false, "List VPC Origins")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontPublishConnectionFunction, "publish-connection-function", "", false, "Publish Connection Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontPublishFunction, "publish-function", "", false, "Publish Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontTagResource, "tag-resource", "", false, "Tag Resource")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontTestConnectionFunction, "test-connection-function", "", false, "Test Connection Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontTestFunction, "test-function", "", false, "Test Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUntagResource, "untag-resource", "", false, "Untag Resource")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateAnycastIpList, "update-anycast-ip-list", "", false, "Update Anycast IP List")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateCachePolicy, "update-cache-policy", "", false, "Update Cache Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateCloudFrontOriginAccessIdentity, "update-cloud-front-origin-access-identity", "", false, "Update Cloud Front Origin Access Identity")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateConnectionFunction, "update-connection-function", "", false, "Update Connection Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateConnectionGroup, "update-connection-group", "", false, "Update Connection Group")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateContinuousDeploymentPolicy, "update-continuous-deployment-policy", "", false, "Update Continuous Deployment Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateDistribution, "update-distribution", "", false, "Update Distribution")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateDistributionTenant, "update-distribution-tenant", "", false, "Update Distribution Tenant")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateDistributionWithStagingConfig, "update-distribution-with-staging-config", "", false, "Update Distribution With Staging Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateDomainAssociation, "update-domain-association", "", false, "Update Domain Association")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateFieldLevelEncryptionConfig, "update-field-level-encryption-config", "", false, "Update Field Level Encryption Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateFieldLevelEncryptionProfile, "update-field-level-encryption-profile", "", false, "Update Field Level Encryption Profile")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateFunction, "update-function", "", false, "Update Function")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateKeyGroup, "update-key-group", "", false, "Update Key Group")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateKeyValueStore, "update-key-value-store", "", false, "Update Key Value Store")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateOriginAccessControl, "update-origin-access-control", "", false, "Update Origin Access Control")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateOriginRequestPolicy, "update-origin-request-policy", "", false, "Update Origin Request Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdatePublicKey, "update-public-key", "", false, "Update Public Key")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateRealtimeLogConfig, "update-realtime-log-config", "", false, "Update Realtime Log Config")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateResponseHeadersPolicy, "update-response-headers-policy", "", false, "Update Response Headers Policy")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateStreamingDistribution, "update-streaming-distribution", "", false, "Update Streaming Distribution")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateTrustStore, "update-trust-store", "", false, "Update Trust Store")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontUpdateVpcOrigin, "update-vpc-origin", "", false, "Update VPC Origin")
	_cloudfrontCmd.Flags().BoolVarP(&_cloudfrontVerifyDnsConfiguration, "verify-dns-configuration", "", false, "Verify DNS Configuration")

}
