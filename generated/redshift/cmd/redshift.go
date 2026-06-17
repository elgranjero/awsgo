package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// redshiftCmd represents the redshift command
var _redshiftCmd = &cobra.Command{
	Use:   "redshift",
	Short: "AWS redshift CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := redshift.NewFromConfig(cfg)
		if _redshiftAcceptReservedNodeExchange {
			redshift_AcceptReservedNodeExchange(cfg, client)
			return
		}
		if _redshiftAddPartner {
			redshift_AddPartner(cfg, client)
			return
		}
		if _redshiftAssociateDataShareConsumer {
			redshift_AssociateDataShareConsumer(cfg, client)
			return
		}
		if _redshiftAuthorizeClusterSecurityGroupIngress {
			redshift_AuthorizeClusterSecurityGroupIngress(cfg, client)
			return
		}
		if _redshiftAuthorizeDataShare {
			redshift_AuthorizeDataShare(cfg, client)
			return
		}
		if _redshiftAuthorizeEndpointAccess {
			redshift_AuthorizeEndpointAccess(cfg, client)
			return
		}
		if _redshiftAuthorizeSnapshotAccess {
			redshift_AuthorizeSnapshotAccess(cfg, client)
			return
		}
		if _redshiftBatchDeleteClusterSnapshots {
			redshift_BatchDeleteClusterSnapshots(cfg, client)
			return
		}
		if _redshiftBatchModifyClusterSnapshots {
			redshift_BatchModifyClusterSnapshots(cfg, client)
			return
		}
		if _redshiftCancelResize {
			redshift_CancelResize(cfg, client)
			return
		}
		if _redshiftCopyClusterSnapshot {
			redshift_CopyClusterSnapshot(cfg, client)
			return
		}
		if _redshiftCreateAuthenticationProfile {
			redshift_CreateAuthenticationProfile(cfg, client)
			return
		}
		if _redshiftCreateCluster {
			redshift_CreateCluster(cfg, client)
			return
		}
		if _redshiftCreateClusterParameterGroup {
			redshift_CreateClusterParameterGroup(cfg, client)
			return
		}
		if _redshiftCreateClusterSecurityGroup {
			redshift_CreateClusterSecurityGroup(cfg, client)
			return
		}
		if _redshiftCreateClusterSnapshot {
			redshift_CreateClusterSnapshot(cfg, client)
			return
		}
		if _redshiftCreateClusterSubnetGroup {
			redshift_CreateClusterSubnetGroup(cfg, client)
			return
		}
		if _redshiftCreateCustomDomainAssociation {
			redshift_CreateCustomDomainAssociation(cfg, client)
			return
		}
		if _redshiftCreateEndpointAccess {
			redshift_CreateEndpointAccess(cfg, client)
			return
		}
		if _redshiftCreateEventSubscription {
			redshift_CreateEventSubscription(cfg, client)
			return
		}
		if _redshiftCreateHsmClientCertificate {
			redshift_CreateHsmClientCertificate(cfg, client)
			return
		}
		if _redshiftCreateHsmConfiguration {
			redshift_CreateHsmConfiguration(cfg, client)
			return
		}
		if _redshiftCreateIntegration {
			redshift_CreateIntegration(cfg, client)
			return
		}
		if _redshiftCreateRedshiftIdcApplication {
			redshift_CreateRedshiftIdcApplication(cfg, client)
			return
		}
		if _redshiftCreateScheduledAction {
			redshift_CreateScheduledAction(cfg, client)
			return
		}
		if _redshiftCreateSnapshotCopyGrant {
			redshift_CreateSnapshotCopyGrant(cfg, client)
			return
		}
		if _redshiftCreateSnapshotSchedule {
			redshift_CreateSnapshotSchedule(cfg, client)
			return
		}
		if _redshiftCreateTags {
			redshift_CreateTags(cfg, client)
			return
		}
		if _redshiftCreateUsageLimit {
			redshift_CreateUsageLimit(cfg, client)
			return
		}
		if _redshiftDeauthorizeDataShare {
			redshift_DeauthorizeDataShare(cfg, client)
			return
		}
		if _redshiftDeleteAuthenticationProfile {
			redshift_DeleteAuthenticationProfile(cfg, client)
			return
		}
		if _redshiftDeleteCluster {
			redshift_DeleteCluster(cfg, client)
			return
		}
		if _redshiftDeleteClusterParameterGroup {
			redshift_DeleteClusterParameterGroup(cfg, client)
			return
		}
		if _redshiftDeleteClusterSecurityGroup {
			redshift_DeleteClusterSecurityGroup(cfg, client)
			return
		}
		if _redshiftDeleteClusterSnapshot {
			redshift_DeleteClusterSnapshot(cfg, client)
			return
		}
		if _redshiftDeleteClusterSubnetGroup {
			redshift_DeleteClusterSubnetGroup(cfg, client)
			return
		}
		if _redshiftDeleteCustomDomainAssociation {
			redshift_DeleteCustomDomainAssociation(cfg, client)
			return
		}
		if _redshiftDeleteEndpointAccess {
			redshift_DeleteEndpointAccess(cfg, client)
			return
		}
		if _redshiftDeleteEventSubscription {
			redshift_DeleteEventSubscription(cfg, client)
			return
		}
		if _redshiftDeleteHsmClientCertificate {
			redshift_DeleteHsmClientCertificate(cfg, client)
			return
		}
		if _redshiftDeleteHsmConfiguration {
			redshift_DeleteHsmConfiguration(cfg, client)
			return
		}
		if _redshiftDeleteIntegration {
			redshift_DeleteIntegration(cfg, client)
			return
		}
		if _redshiftDeletePartner {
			redshift_DeletePartner(cfg, client)
			return
		}
		if _redshiftDeleteRedshiftIdcApplication {
			redshift_DeleteRedshiftIdcApplication(cfg, client)
			return
		}
		if _redshiftDeleteResourcePolicy {
			redshift_DeleteResourcePolicy(cfg, client)
			return
		}
		if _redshiftDeleteScheduledAction {
			redshift_DeleteScheduledAction(cfg, client)
			return
		}
		if _redshiftDeleteSnapshotCopyGrant {
			redshift_DeleteSnapshotCopyGrant(cfg, client)
			return
		}
		if _redshiftDeleteSnapshotSchedule {
			redshift_DeleteSnapshotSchedule(cfg, client)
			return
		}
		if _redshiftDeleteTags {
			redshift_DeleteTags(cfg, client)
			return
		}
		if _redshiftDeleteUsageLimit {
			redshift_DeleteUsageLimit(cfg, client)
			return
		}
		if _redshiftDeregisterNamespace {
			redshift_DeregisterNamespace(cfg, client)
			return
		}
		if _redshiftDescribeAccountAttributes {
			redshift_DescribeAccountAttributes(cfg, client)
			return
		}
		if _redshiftDescribeAuthenticationProfiles {
			redshift_DescribeAuthenticationProfiles(cfg, client)
			return
		}
		if _redshiftDescribeClusterDbRevisions {
			redshift_DescribeClusterDbRevisions(cfg, client)
			return
		}
		if _redshiftDescribeClusterParameterGroups {
			redshift_DescribeClusterParameterGroups(cfg, client)
			return
		}
		if _redshiftDescribeClusterParameters {
			redshift_DescribeClusterParameters(cfg, client)
			return
		}
		if _redshiftDescribeClusterSecurityGroups {
			redshift_DescribeClusterSecurityGroups(cfg, client)
			return
		}
		if _redshiftDescribeClusterSnapshots {
			redshift_DescribeClusterSnapshots(cfg, client)
			return
		}
		if _redshiftDescribeClusterSubnetGroups {
			redshift_DescribeClusterSubnetGroups(cfg, client)
			return
		}
		if _redshiftDescribeClusterTracks {
			redshift_DescribeClusterTracks(cfg, client)
			return
		}
		if _redshiftDescribeClusterVersions {
			redshift_DescribeClusterVersions(cfg, client)
			return
		}
		if _redshiftDescribeClusters {
			redshift_DescribeClusters(cfg, client)
			return
		}
		if _redshiftDescribeCustomDomainAssociations {
			redshift_DescribeCustomDomainAssociations(cfg, client)
			return
		}
		if _redshiftDescribeDataShares {
			redshift_DescribeDataShares(cfg, client)
			return
		}
		if _redshiftDescribeDataSharesForConsumer {
			redshift_DescribeDataSharesForConsumer(cfg, client)
			return
		}
		if _redshiftDescribeDataSharesForProducer {
			redshift_DescribeDataSharesForProducer(cfg, client)
			return
		}
		if _redshiftDescribeDefaultClusterParameters {
			redshift_DescribeDefaultClusterParameters(cfg, client)
			return
		}
		if _redshiftDescribeEndpointAccess {
			redshift_DescribeEndpointAccess(cfg, client)
			return
		}
		if _redshiftDescribeEndpointAuthorization {
			redshift_DescribeEndpointAuthorization(cfg, client)
			return
		}
		if _redshiftDescribeEventCategories {
			redshift_DescribeEventCategories(cfg, client)
			return
		}
		if _redshiftDescribeEventSubscriptions {
			redshift_DescribeEventSubscriptions(cfg, client)
			return
		}
		if _redshiftDescribeEvents {
			redshift_DescribeEvents(cfg, client)
			return
		}
		if _redshiftDescribeHsmClientCertificates {
			redshift_DescribeHsmClientCertificates(cfg, client)
			return
		}
		if _redshiftDescribeHsmConfigurations {
			redshift_DescribeHsmConfigurations(cfg, client)
			return
		}
		if _redshiftDescribeInboundIntegrations {
			redshift_DescribeInboundIntegrations(cfg, client)
			return
		}
		if _redshiftDescribeIntegrations {
			redshift_DescribeIntegrations(cfg, client)
			return
		}
		if _redshiftDescribeLoggingStatus {
			redshift_DescribeLoggingStatus(cfg, client)
			return
		}
		if _redshiftDescribeNodeConfigurationOptions {
			redshift_DescribeNodeConfigurationOptions(cfg, client)
			return
		}
		if _redshiftDescribeOrderableClusterOptions {
			redshift_DescribeOrderableClusterOptions(cfg, client)
			return
		}
		if _redshiftDescribePartners {
			redshift_DescribePartners(cfg, client)
			return
		}
		if _redshiftDescribeRedshiftIdcApplications {
			redshift_DescribeRedshiftIdcApplications(cfg, client)
			return
		}
		if _redshiftDescribeReservedNodeExchangeStatus {
			redshift_DescribeReservedNodeExchangeStatus(cfg, client)
			return
		}
		if _redshiftDescribeReservedNodeOfferings {
			redshift_DescribeReservedNodeOfferings(cfg, client)
			return
		}
		if _redshiftDescribeReservedNodes {
			redshift_DescribeReservedNodes(cfg, client)
			return
		}
		if _redshiftDescribeResize {
			redshift_DescribeResize(cfg, client)
			return
		}
		if _redshiftDescribeScheduledActions {
			redshift_DescribeScheduledActions(cfg, client)
			return
		}
		if _redshiftDescribeSnapshotCopyGrants {
			redshift_DescribeSnapshotCopyGrants(cfg, client)
			return
		}
		if _redshiftDescribeSnapshotSchedules {
			redshift_DescribeSnapshotSchedules(cfg, client)
			return
		}
		if _redshiftDescribeStorage {
			redshift_DescribeStorage(cfg, client)
			return
		}
		if _redshiftDescribeTableRestoreStatus {
			redshift_DescribeTableRestoreStatus(cfg, client)
			return
		}
		if _redshiftDescribeTags {
			redshift_DescribeTags(cfg, client)
			return
		}
		if _redshiftDescribeUsageLimits {
			redshift_DescribeUsageLimits(cfg, client)
			return
		}
		if _redshiftDisableLogging {
			redshift_DisableLogging(cfg, client)
			return
		}
		if _redshiftDisableSnapshotCopy {
			redshift_DisableSnapshotCopy(cfg, client)
			return
		}
		if _redshiftDisassociateDataShareConsumer {
			redshift_DisassociateDataShareConsumer(cfg, client)
			return
		}
		if _redshiftEnableLogging {
			redshift_EnableLogging(cfg, client)
			return
		}
		if _redshiftEnableSnapshotCopy {
			redshift_EnableSnapshotCopy(cfg, client)
			return
		}
		if _redshiftFailoverPrimaryCompute {
			redshift_FailoverPrimaryCompute(cfg, client)
			return
		}
		if _redshiftGetClusterCredentials {
			redshift_GetClusterCredentials(cfg, client)
			return
		}
		if _redshiftGetClusterCredentialsWithIAM {
			redshift_GetClusterCredentialsWithIAM(cfg, client)
			return
		}
		if _redshiftGetIdentityCenterAuthToken {
			redshift_GetIdentityCenterAuthToken(cfg, client)
			return
		}
		if _redshiftGetReservedNodeExchangeConfigurationOptions {
			redshift_GetReservedNodeExchangeConfigurationOptions(cfg, client)
			return
		}
		if _redshiftGetReservedNodeExchangeOfferings {
			redshift_GetReservedNodeExchangeOfferings(cfg, client)
			return
		}
		if _redshiftGetResourcePolicy {
			redshift_GetResourcePolicy(cfg, client)
			return
		}
		if _redshiftListRecommendations {
			redshift_ListRecommendations(cfg, client)
			return
		}
		if _redshiftModifyAquaConfiguration {
			redshift_ModifyAquaConfiguration(cfg, client)
			return
		}
		if _redshiftModifyAuthenticationProfile {
			redshift_ModifyAuthenticationProfile(cfg, client)
			return
		}
		if _redshiftModifyCluster {
			redshift_ModifyCluster(cfg, client)
			return
		}
		if _redshiftModifyClusterDbRevision {
			redshift_ModifyClusterDbRevision(cfg, client)
			return
		}
		if _redshiftModifyClusterIamRoles {
			redshift_ModifyClusterIamRoles(cfg, client)
			return
		}
		if _redshiftModifyClusterMaintenance {
			redshift_ModifyClusterMaintenance(cfg, client)
			return
		}
		if _redshiftModifyClusterParameterGroup {
			redshift_ModifyClusterParameterGroup(cfg, client)
			return
		}
		if _redshiftModifyClusterSnapshot {
			redshift_ModifyClusterSnapshot(cfg, client)
			return
		}
		if _redshiftModifyClusterSnapshotSchedule {
			redshift_ModifyClusterSnapshotSchedule(cfg, client)
			return
		}
		if _redshiftModifyClusterSubnetGroup {
			redshift_ModifyClusterSubnetGroup(cfg, client)
			return
		}
		if _redshiftModifyCustomDomainAssociation {
			redshift_ModifyCustomDomainAssociation(cfg, client)
			return
		}
		if _redshiftModifyEndpointAccess {
			redshift_ModifyEndpointAccess(cfg, client)
			return
		}
		if _redshiftModifyEventSubscription {
			redshift_ModifyEventSubscription(cfg, client)
			return
		}
		if _redshiftModifyIntegration {
			redshift_ModifyIntegration(cfg, client)
			return
		}
		if _redshiftModifyLakehouseConfiguration {
			redshift_ModifyLakehouseConfiguration(cfg, client)
			return
		}
		if _redshiftModifyRedshiftIdcApplication {
			redshift_ModifyRedshiftIdcApplication(cfg, client)
			return
		}
		if _redshiftModifyScheduledAction {
			redshift_ModifyScheduledAction(cfg, client)
			return
		}
		if _redshiftModifySnapshotCopyRetentionPeriod {
			redshift_ModifySnapshotCopyRetentionPeriod(cfg, client)
			return
		}
		if _redshiftModifySnapshotSchedule {
			redshift_ModifySnapshotSchedule(cfg, client)
			return
		}
		if _redshiftModifyUsageLimit {
			redshift_ModifyUsageLimit(cfg, client)
			return
		}
		if _redshiftPauseCluster {
			redshift_PauseCluster(cfg, client)
			return
		}
		if _redshiftPurchaseReservedNodeOffering {
			redshift_PurchaseReservedNodeOffering(cfg, client)
			return
		}
		if _redshiftPutResourcePolicy {
			redshift_PutResourcePolicy(cfg, client)
			return
		}
		if _redshiftRebootCluster {
			redshift_RebootCluster(cfg, client)
			return
		}
		if _redshiftRegisterNamespace {
			redshift_RegisterNamespace(cfg, client)
			return
		}
		if _redshiftRejectDataShare {
			redshift_RejectDataShare(cfg, client)
			return
		}
		if _redshiftResetClusterParameterGroup {
			redshift_ResetClusterParameterGroup(cfg, client)
			return
		}
		if _redshiftResizeCluster {
			redshift_ResizeCluster(cfg, client)
			return
		}
		if _redshiftRestoreFromClusterSnapshot {
			redshift_RestoreFromClusterSnapshot(cfg, client)
			return
		}
		if _redshiftRestoreTableFromClusterSnapshot {
			redshift_RestoreTableFromClusterSnapshot(cfg, client)
			return
		}
		if _redshiftResumeCluster {
			redshift_ResumeCluster(cfg, client)
			return
		}
		if _redshiftRevokeClusterSecurityGroupIngress {
			redshift_RevokeClusterSecurityGroupIngress(cfg, client)
			return
		}
		if _redshiftRevokeEndpointAccess {
			redshift_RevokeEndpointAccess(cfg, client)
			return
		}
		if _redshiftRevokeSnapshotAccess {
			redshift_RevokeSnapshotAccess(cfg, client)
			return
		}
		if _redshiftRotateEncryptionKey {
			redshift_RotateEncryptionKey(cfg, client)
			return
		}
		if _redshiftUpdatePartnerStatus {
			redshift_UpdatePartnerStatus(cfg, client)
			return
		}

	},
}

var (
	_redshiftAcceptReservedNodeExchange                  bool
	_redshiftAddPartner                                  bool
	_redshiftAssociateDataShareConsumer                  bool
	_redshiftAuthorizeClusterSecurityGroupIngress        bool
	_redshiftAuthorizeDataShare                          bool
	_redshiftAuthorizeEndpointAccess                     bool
	_redshiftAuthorizeSnapshotAccess                     bool
	_redshiftBatchDeleteClusterSnapshots                 bool
	_redshiftBatchModifyClusterSnapshots                 bool
	_redshiftCancelResize                                bool
	_redshiftCopyClusterSnapshot                         bool
	_redshiftCreateAuthenticationProfile                 bool
	_redshiftCreateCluster                               bool
	_redshiftCreateClusterParameterGroup                 bool
	_redshiftCreateClusterSecurityGroup                  bool
	_redshiftCreateClusterSnapshot                       bool
	_redshiftCreateClusterSubnetGroup                    bool
	_redshiftCreateCustomDomainAssociation               bool
	_redshiftCreateEndpointAccess                        bool
	_redshiftCreateEventSubscription                     bool
	_redshiftCreateHsmClientCertificate                  bool
	_redshiftCreateHsmConfiguration                      bool
	_redshiftCreateIntegration                           bool
	_redshiftCreateRedshiftIdcApplication                bool
	_redshiftCreateScheduledAction                       bool
	_redshiftCreateSnapshotCopyGrant                     bool
	_redshiftCreateSnapshotSchedule                      bool
	_redshiftCreateTags                                  bool
	_redshiftCreateUsageLimit                            bool
	_redshiftDeauthorizeDataShare                        bool
	_redshiftDeleteAuthenticationProfile                 bool
	_redshiftDeleteCluster                               bool
	_redshiftDeleteClusterParameterGroup                 bool
	_redshiftDeleteClusterSecurityGroup                  bool
	_redshiftDeleteClusterSnapshot                       bool
	_redshiftDeleteClusterSubnetGroup                    bool
	_redshiftDeleteCustomDomainAssociation               bool
	_redshiftDeleteEndpointAccess                        bool
	_redshiftDeleteEventSubscription                     bool
	_redshiftDeleteHsmClientCertificate                  bool
	_redshiftDeleteHsmConfiguration                      bool
	_redshiftDeleteIntegration                           bool
	_redshiftDeletePartner                               bool
	_redshiftDeleteRedshiftIdcApplication                bool
	_redshiftDeleteResourcePolicy                        bool
	_redshiftDeleteScheduledAction                       bool
	_redshiftDeleteSnapshotCopyGrant                     bool
	_redshiftDeleteSnapshotSchedule                      bool
	_redshiftDeleteTags                                  bool
	_redshiftDeleteUsageLimit                            bool
	_redshiftDeregisterNamespace                         bool
	_redshiftDescribeAccountAttributes                   bool
	_redshiftDescribeAuthenticationProfiles              bool
	_redshiftDescribeClusterDbRevisions                  bool
	_redshiftDescribeClusterParameterGroups              bool
	_redshiftDescribeClusterParameters                   bool
	_redshiftDescribeClusterSecurityGroups               bool
	_redshiftDescribeClusterSnapshots                    bool
	_redshiftDescribeClusterSubnetGroups                 bool
	_redshiftDescribeClusterTracks                       bool
	_redshiftDescribeClusterVersions                     bool
	_redshiftDescribeClusters                            bool
	_redshiftDescribeCustomDomainAssociations            bool
	_redshiftDescribeDataShares                          bool
	_redshiftDescribeDataSharesForConsumer               bool
	_redshiftDescribeDataSharesForProducer               bool
	_redshiftDescribeDefaultClusterParameters            bool
	_redshiftDescribeEndpointAccess                      bool
	_redshiftDescribeEndpointAuthorization               bool
	_redshiftDescribeEventCategories                     bool
	_redshiftDescribeEventSubscriptions                  bool
	_redshiftDescribeEvents                              bool
	_redshiftDescribeHsmClientCertificates               bool
	_redshiftDescribeHsmConfigurations                   bool
	_redshiftDescribeInboundIntegrations                 bool
	_redshiftDescribeIntegrations                        bool
	_redshiftDescribeLoggingStatus                       bool
	_redshiftDescribeNodeConfigurationOptions            bool
	_redshiftDescribeOrderableClusterOptions             bool
	_redshiftDescribePartners                            bool
	_redshiftDescribeRedshiftIdcApplications             bool
	_redshiftDescribeReservedNodeExchangeStatus          bool
	_redshiftDescribeReservedNodeOfferings               bool
	_redshiftDescribeReservedNodes                       bool
	_redshiftDescribeResize                              bool
	_redshiftDescribeScheduledActions                    bool
	_redshiftDescribeSnapshotCopyGrants                  bool
	_redshiftDescribeSnapshotSchedules                   bool
	_redshiftDescribeStorage                             bool
	_redshiftDescribeTableRestoreStatus                  bool
	_redshiftDescribeTags                                bool
	_redshiftDescribeUsageLimits                         bool
	_redshiftDisableLogging                              bool
	_redshiftDisableSnapshotCopy                         bool
	_redshiftDisassociateDataShareConsumer               bool
	_redshiftEnableLogging                               bool
	_redshiftEnableSnapshotCopy                          bool
	_redshiftFailoverPrimaryCompute                      bool
	_redshiftGetClusterCredentials                       bool
	_redshiftGetClusterCredentialsWithIAM                bool
	_redshiftGetIdentityCenterAuthToken                  bool
	_redshiftGetReservedNodeExchangeConfigurationOptions bool
	_redshiftGetReservedNodeExchangeOfferings            bool
	_redshiftGetResourcePolicy                           bool
	_redshiftListRecommendations                         bool
	_redshiftModifyAquaConfiguration                     bool
	_redshiftModifyAuthenticationProfile                 bool
	_redshiftModifyCluster                               bool
	_redshiftModifyClusterDbRevision                     bool
	_redshiftModifyClusterIamRoles                       bool
	_redshiftModifyClusterMaintenance                    bool
	_redshiftModifyClusterParameterGroup                 bool
	_redshiftModifyClusterSnapshot                       bool
	_redshiftModifyClusterSnapshotSchedule               bool
	_redshiftModifyClusterSubnetGroup                    bool
	_redshiftModifyCustomDomainAssociation               bool
	_redshiftModifyEndpointAccess                        bool
	_redshiftModifyEventSubscription                     bool
	_redshiftModifyIntegration                           bool
	_redshiftModifyLakehouseConfiguration                bool
	_redshiftModifyRedshiftIdcApplication                bool
	_redshiftModifyScheduledAction                       bool
	_redshiftModifySnapshotCopyRetentionPeriod           bool
	_redshiftModifySnapshotSchedule                      bool
	_redshiftModifyUsageLimit                            bool
	_redshiftPauseCluster                                bool
	_redshiftPurchaseReservedNodeOffering                bool
	_redshiftPutResourcePolicy                           bool
	_redshiftRebootCluster                               bool
	_redshiftRegisterNamespace                           bool
	_redshiftRejectDataShare                             bool
	_redshiftResetClusterParameterGroup                  bool
	_redshiftResizeCluster                               bool
	_redshiftRestoreFromClusterSnapshot                  bool
	_redshiftRestoreTableFromClusterSnapshot             bool
	_redshiftResumeCluster                               bool
	_redshiftRevokeClusterSecurityGroupIngress           bool
	_redshiftRevokeEndpointAccess                        bool
	_redshiftRevokeSnapshotAccess                        bool
	_redshiftRotateEncryptionKey                         bool
	_redshiftUpdatePartnerStatus                         bool

	_redshiftAccount                              string
	_redshiftAccountId                            string
	_redshiftAccountWithRestoreAccess             string
	_redshiftActionType                           string
	_redshiftActive                               string
	_redshiftAddIamRoles                          []string
	_redshiftAdditionalEncryptionContext          string
	_redshiftAdditionalInfo                       string
	_redshiftAllowVersionUpgrade                  string
	_redshiftAllowWrites                          string
	_redshiftAmount                               string
	_redshiftApplicationType                      string
	_redshiftAquaConfigurationStatus              string
	_redshiftAssociateEntireAccount               string
	_redshiftAttributeNames                       []string
	_redshiftAuthenticationProfileContent         string
	_redshiftAuthenticationProfileName            string
	_redshiftAuthorizedTokenIssuerList            string
	_redshiftAutoCreate                           string
	_redshiftAutomatedSnapshotRetentionPeriod     string
	_redshiftAvailabilityZone                     string
	_redshiftAvailabilityZoneRelocation           string
	_redshiftBreachAction                         string
	_redshiftBucketName                           string
	_redshiftCatalogName                          string
	_redshiftCIDRIP                               string
	_redshiftClassic                              string
	_redshiftClusterExists                        string
	_redshiftClusterIdentifier                    string
	_redshiftClusterIds                           []string
	_redshiftClusterParameterGroupFamily          string
	_redshiftClusterParameterGroupName            string
	_redshiftClusterSecurityGroupName             string
	_redshiftClusterSecurityGroups                []string
	_redshiftClusterSubnetGroupName               string
	_redshiftClusterType                          string
	_redshiftClusterVersion                       string
	_redshiftConsumerArn                          string
	_redshiftConsumerIdentifier                   string
	_redshiftConsumerIdentifiers                  []string
	_redshiftConsumerRegion                       string
	_redshiftCustomDomainCertificateArn           string
	_redshiftCustomDomainName                     string
	_redshiftDataShareArn                         string
	_redshiftDatabaseName                         string
	_redshiftDbGroups                             []string
	_redshiftDBName                               string
	_redshiftDbUser                               string
	_redshiftDefaultIamRoleArn                    string
	_redshiftDeferMaintenance                     string
	_redshiftDeferMaintenanceDuration             string
	_redshiftDeferMaintenanceEndTime              string
	_redshiftDeferMaintenanceIdentifier           string
	_redshiftDeferMaintenanceStartTime            string
	_redshiftDescription                          string
	_redshiftDestinationRegion                    string
	_redshiftDisassociateEntireAccount            string
	_redshiftDisassociateSchedule                 string
	_redshiftDryRun                               string
	_redshiftDuration                             string
	_redshiftDurationSeconds                      string
	_redshiftEC2SecurityGroupName                 string
	_redshiftEC2SecurityGroupOwnerId              string
	_redshiftElasticIp                            string
	_redshiftEnable                               string
	_redshiftEnableCaseSensitiveIdentifier        string
	_redshiftEnabled                              string
	_redshiftEncrypted                            string
	_redshiftEndTime                              string
	_redshiftEndpointName                         string
	_redshiftEnhancedVpcRouting                   string
	_redshiftEventCategories                      []string
	_redshiftExtraComputeForAutomaticOptimization string
	_redshiftFeatureType                          string
	_redshiftFilters                              string
	_redshiftFinalClusterSnapshotIdentifier       string
	_redshiftFinalClusterSnapshotRetentionPeriod  string
	_redshiftForce                                string
	_redshiftGrantee                              string
	_redshiftHsmClientCertificateIdentifier       string
	_redshiftHsmConfigurationIdentifier           string
	_redshiftHsmIpAddress                         string
	_redshiftHsmPartitionName                     string
	_redshiftHsmPartitionPassword                 string
	_redshiftHsmServerPublicCertificate           string
	_redshiftIamRole                              string
	_redshiftIamRoleArn                           string
	_redshiftIamRoles                             []string
	_redshiftIdcDisplayName                       string
	_redshiftIdcInstanceArn                       string
	_redshiftIdentifiers                          string
	_redshiftIdentityNamespace                    string
	_redshiftIntegrationArn                       string
	_redshiftIntegrationName                      string
	_redshiftIpAddressType                        string
	_redshiftKMSKeyId                             string
	_redshiftLakehouseIdcApplicationArn           string
	_redshiftLakehouseIdcRegistration             string
	_redshiftLakehouseRegistration                string
	_redshiftLimitType                            string
	_redshiftLoadSampleData                       string
	_redshiftLogDestinationType                   string
	_redshiftLogExports                           []string
	_redshiftMaintenanceTrackName                 string
	_redshiftManageMasterPassword                 string
	_redshiftManual                               string
	_redshiftManualSnapshotRetentionPeriod        string
	_redshiftMarker                               string
	_redshiftMasterPasswordSecretKmsKeyId         string
	_redshiftMasterUserPassword                   string
	_redshiftMasterUsername                       string
	_redshiftMaxRecords                           string
	_redshiftMultiAZ                              string
	_redshiftNamespaceArn                         string
	_redshiftNamespaceIdentifier                  string
	_redshiftNewClusterIdentifier                 string
	_redshiftNewTableName                         string
	_redshiftNextInvocations                      string
	_redshiftNodeCount                            string
	_redshiftNodeType                             string
	_redshiftNumberOfNodes                        string
	_redshiftOwnerAccount                         string
	_redshiftParameterGroupFamily                 string
	_redshiftParameterGroupName                   string
	_redshiftParameters                           string
	_redshiftPartnerName                          string
	_redshiftPeriod                               string
	_redshiftPolicy                               string
	_redshiftPort                                 string
	_redshiftPreferredMaintenanceWindow           string
	_redshiftProducerArn                          string
	_redshiftPubliclyAccessible                   string
	_redshiftRedshiftIdcApplicationArn            string
	_redshiftRedshiftIdcApplicationName           string
	_redshiftRemoveIamRoles                       []string
	_redshiftReservedNodeExchangeRequestId        string
	_redshiftReservedNodeId                       string
	_redshiftReservedNodeOfferingId               string
	_redshiftResetAllParameters                   string
	_redshiftResourceArn                          string
	_redshiftResourceName                         string
	_redshiftResourceOwner                        string
	_redshiftResourceType                         string
	_redshiftRetentionPeriod                      string
	_redshiftRevisionTarget                       string
	_redshiftS3KeyPrefix                          string
	_redshiftSchedule                             string
	_redshiftScheduleDefinitions                  []string
	_redshiftScheduleDescription                  string
	_redshiftScheduleIdentifier                   string
	_redshiftScheduledActionDescription           string
	_redshiftScheduledActionName                  string
	_redshiftServiceIntegrations                  string
	_redshiftSeverity                             string
	_redshiftSkipFinalClusterSnapshot             string
	_redshiftSnapshotArn                          string
	_redshiftSnapshotClusterIdentifier            string
	_redshiftSnapshotCopyGrantName                string
	_redshiftSnapshotIdentifier                   string
	_redshiftSnapshotIdentifierList               []string
	_redshiftSnapshotScheduleIdentifier           string
	_redshiftSnapshotType                         string
	_redshiftSnsTopicArn                          string
	_redshiftSortingEntities                      string
	_redshiftSource                               string
	_redshiftSourceArn                            string
	_redshiftSourceDatabaseName                   string
	_redshiftSourceIdentifier                     string
	_redshiftSourceIds                            []string
	_redshiftSourceSchemaName                     string
	_redshiftSourceSnapshotClusterIdentifier      string
	_redshiftSourceSnapshotIdentifier             string
	_redshiftSourceTableName                      string
	_redshiftSourceType                           string
	_redshiftSsoTagKeys                           []string
	_redshiftStartTime                            string
	_redshiftStatus                               string
	_redshiftStatusMessage                        string
	_redshiftSubnetGroupName                      string
	_redshiftSubnetIds                            []string
	_redshiftSubscriptionName                     string
	_redshiftTableRestoreRequestId                string
	_redshiftTagKeys                              []string
	_redshiftTagList                              string
	_redshiftTagValues                            []string
	_redshiftTags                                 string
	_redshiftTargetAction                         string
	_redshiftTargetActionType                     string
	_redshiftTargetArn                            string
	_redshiftTargetDatabaseName                   string
	_redshiftTargetReservedNodeOfferingId         string
	_redshiftTargetSchemaName                     string
	_redshiftTargetSnapshotIdentifier             string
	_redshiftUsageLimitId                         string
	_redshiftVpcId                                string
	_redshiftVpcIds                               []string
	_redshiftVpcSecurityGroupIds                  []string
)

// Exchanges a DC1 Reserved Node for a DC2 Reserved Node with no changes to the
// configuration (term, payment type, or number of nodes) and no additional costs.
func redshift_AcceptReservedNodeExchange(cfg aws.Config, client *redshift.Client) {
	input := &redshift.AcceptReservedNodeExchangeInput{
		// ReservedNodeId: *string, // Required
		// TargetReservedNodeOfferingId: *string, // Required
	}

	if len(_redshiftReservedNodeId) > 0 {
		input.ReservedNodeId = aws.String(_redshiftReservedNodeId)
	}
	if len(_redshiftTargetReservedNodeOfferingId) > 0 {
		input.TargetReservedNodeOfferingId = aws.String(_redshiftTargetReservedNodeOfferingId)
	}

	if resp, err := client.AcceptReservedNodeExchange(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a partner integration to a cluster. This operation authorizes a partner to
// push status updates for the specified database. To complete the integration, you
// also set up the integration on the partner website.
func redshift_AddPartner(cfg aws.Config, client *redshift.Client) {
	input := &redshift.AddPartnerInput{
		// AccountId: *string, // Required
		// ClusterIdentifier: *string, // Required
		// DatabaseName: *string, // Required
		// PartnerName: *string, // Required
	}

	if len(_redshiftAccountId) > 0 {
		input.AccountId = aws.String(_redshiftAccountId)
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftDatabaseName) > 0 {
		input.DatabaseName = aws.String(_redshiftDatabaseName)
	}
	if len(_redshiftPartnerName) > 0 {
		input.PartnerName = aws.String(_redshiftPartnerName)
	}

	if resp, err := client.AddPartner(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// From a datashare consumer account, associates a datashare with the account
// (AssociateEntireAccount) or the specified namespace (ConsumerArn). If you make
// this association, the consumer can consume the datashare.
func redshift_AssociateDataShareConsumer(cfg aws.Config, client *redshift.Client) {
	input := &redshift.AssociateDataShareConsumerInput{
		// DataShareArn: *string, // Required
	}

	if len(_redshiftDataShareArn) > 0 {
		input.DataShareArn = aws.String(_redshiftDataShareArn)
	}
	if len(_redshiftAllowWrites) > 0 {
		if err := assignInputField(input, "AllowWrites", _redshiftAllowWrites); err != nil {
			log.Errorf("invalid --allow-writes: %s", err.Error())
			return
		}
	}
	if len(_redshiftAssociateEntireAccount) > 0 {
		if err := assignInputField(input, "AssociateEntireAccount", _redshiftAssociateEntireAccount); err != nil {
			log.Errorf("invalid --associate-entire-account: %s", err.Error())
			return
		}
	}
	if len(_redshiftConsumerArn) > 0 {
		input.ConsumerArn = aws.String(_redshiftConsumerArn)
	}
	if len(_redshiftConsumerRegion) > 0 {
		input.ConsumerRegion = aws.String(_redshiftConsumerRegion)
	}

	if resp, err := client.AssociateDataShareConsumer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an inbound (ingress) rule to an Amazon Redshift security group. Depending
// on whether the application accessing your cluster is running on the Internet or
// an Amazon EC2 instance, you can authorize inbound access to either a Classless
// Interdomain Routing (CIDR)/Internet Protocol (IP) range or to an Amazon EC2
// security group. You can add as many as 20 ingress rules to an Amazon Redshift
// security group.
//
// If you authorize access to an Amazon EC2 security group, specify
// EC2SecurityGroupName and EC2SecurityGroupOwnerId. The Amazon EC2 security group
// and Amazon Redshift cluster must be in the same Amazon Web Services Region.
//
// If you authorize access to a CIDR/IP address range, specify CIDRIP. For an
// overview of CIDR blocks, see the Wikipedia article on [Classless Inter-Domain Routing].
//
// You must also associate the security group with a cluster so that clients
// running on these IP addresses or the EC2 instance are authorized to connect to
// the cluster. For information about managing security groups, go to [Working with Security Groups]in the
// Amazon Redshift Cluster Management Guide.
//
// [Classless Inter-Domain Routing]: http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing
// [Working with Security Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html
func redshift_AuthorizeClusterSecurityGroupIngress(cfg aws.Config, client *redshift.Client) {
	input := &redshift.AuthorizeClusterSecurityGroupIngressInput{
		// ClusterSecurityGroupName: *string, // Required
	}

	if len(_redshiftClusterSecurityGroupName) > 0 {
		input.ClusterSecurityGroupName = aws.String(_redshiftClusterSecurityGroupName)
	}
	if len(_redshiftCIDRIP) > 0 {
		input.CIDRIP = aws.String(_redshiftCIDRIP)
	}
	if len(_redshiftEC2SecurityGroupName) > 0 {
		input.EC2SecurityGroupName = aws.String(_redshiftEC2SecurityGroupName)
	}
	if len(_redshiftEC2SecurityGroupOwnerId) > 0 {
		input.EC2SecurityGroupOwnerId = aws.String(_redshiftEC2SecurityGroupOwnerId)
	}

	if resp, err := client.AuthorizeClusterSecurityGroupIngress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// From a data producer account, authorizes the sharing of a datashare with one or
// more consumer accounts or managing entities. To authorize a datashare for a data
// consumer, the producer account must have the correct access permissions.
func redshift_AuthorizeDataShare(cfg aws.Config, client *redshift.Client) {
	input := &redshift.AuthorizeDataShareInput{
		// ConsumerIdentifier: *string, // Required
		// DataShareArn: *string, // Required
	}

	if len(_redshiftConsumerIdentifier) > 0 {
		input.ConsumerIdentifier = aws.String(_redshiftConsumerIdentifier)
	}
	if len(_redshiftDataShareArn) > 0 {
		input.DataShareArn = aws.String(_redshiftDataShareArn)
	}
	if len(_redshiftAllowWrites) > 0 {
		if err := assignInputField(input, "AllowWrites", _redshiftAllowWrites); err != nil {
			log.Errorf("invalid --allow-writes: %s", err.Error())
			return
		}
	}

	if resp, err := client.AuthorizeDataShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Grants access to a cluster.
func redshift_AuthorizeEndpointAccess(cfg aws.Config, client *redshift.Client) {
	input := &redshift.AuthorizeEndpointAccessInput{
		// Account: *string, // Required
	}

	if len(_redshiftAccount) > 0 {
		input.Account = aws.String(_redshiftAccount)
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftVpcIds) > 0 {
		input.VpcIds = append([]string(nil), _redshiftVpcIds...)
	}

	if resp, err := client.AuthorizeEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Authorizes the specified Amazon Web Services account to restore the specified
// snapshot.
//
// For more information about working with snapshots, go to [Amazon Redshift Snapshots] in the Amazon
// Redshift Cluster Management Guide.
//
// [Amazon Redshift Snapshots]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html
func redshift_AuthorizeSnapshotAccess(cfg aws.Config, client *redshift.Client) {
	input := &redshift.AuthorizeSnapshotAccessInput{
		// AccountWithRestoreAccess: *string, // Required
	}

	if len(_redshiftAccountWithRestoreAccess) > 0 {
		input.AccountWithRestoreAccess = aws.String(_redshiftAccountWithRestoreAccess)
	}
	if len(_redshiftSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_redshiftSnapshotArn)
	}
	if len(_redshiftSnapshotClusterIdentifier) > 0 {
		input.SnapshotClusterIdentifier = aws.String(_redshiftSnapshotClusterIdentifier)
	}
	if len(_redshiftSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_redshiftSnapshotIdentifier)
	}

	if resp, err := client.AuthorizeSnapshotAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a set of cluster snapshots.
func redshift_BatchDeleteClusterSnapshots(cfg aws.Config, client *redshift.Client) {
	input := &redshift.BatchDeleteClusterSnapshotsInput{
		// Identifiers: []types.DeleteClusterSnapshotMessage, // Required
	}

	if len(_redshiftIdentifiers) > 0 {
		if err := assignInputField(input, "Identifiers", _redshiftIdentifiers); err != nil {
			log.Errorf("invalid --identifiers: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDeleteClusterSnapshots(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings for a set of cluster snapshots.
func redshift_BatchModifyClusterSnapshots(cfg aws.Config, client *redshift.Client) {
	input := &redshift.BatchModifyClusterSnapshotsInput{
		// SnapshotIdentifierList: []string, // Required
	}

	if len(_redshiftSnapshotIdentifierList) > 0 {
		input.SnapshotIdentifierList = append([]string(nil), _redshiftSnapshotIdentifierList...)
	}
	if len(_redshiftForce) > 0 {
		if err := assignInputField(input, "Force", _redshiftForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_redshiftManualSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "ManualSnapshotRetentionPeriod", _redshiftManualSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --manual-snapshot-retention-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchModifyClusterSnapshots(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a resize operation for a cluster.
func redshift_CancelResize(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CancelResizeInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}

	if resp, err := client.CancelResize(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the specified automated cluster snapshot to a new manual cluster
// snapshot. The source must be an automated snapshot and it must be in the
// available state.
//
// When you delete a cluster, Amazon Redshift deletes any automated snapshots of
// the cluster. Also, when the retention period of the snapshot expires, Amazon
// Redshift automatically deletes it. If you want to keep an automated snapshot for
// a longer period, you can make a manual copy of the snapshot. Manual snapshots
// are retained until you delete them.
//
// For more information about working with snapshots, go to [Amazon Redshift Snapshots] in the Amazon
// Redshift Cluster Management Guide.
//
// [Amazon Redshift Snapshots]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html
func redshift_CopyClusterSnapshot(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CopyClusterSnapshotInput{
		// SourceSnapshotIdentifier: *string, // Required
		// TargetSnapshotIdentifier: *string, // Required
	}

	if len(_redshiftSourceSnapshotIdentifier) > 0 {
		input.SourceSnapshotIdentifier = aws.String(_redshiftSourceSnapshotIdentifier)
	}
	if len(_redshiftTargetSnapshotIdentifier) > 0 {
		input.TargetSnapshotIdentifier = aws.String(_redshiftTargetSnapshotIdentifier)
	}
	if len(_redshiftManualSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "ManualSnapshotRetentionPeriod", _redshiftManualSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --manual-snapshot-retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftSourceSnapshotClusterIdentifier) > 0 {
		input.SourceSnapshotClusterIdentifier = aws.String(_redshiftSourceSnapshotClusterIdentifier)
	}

	if resp, err := client.CopyClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an authentication profile with the specified parameters.
func redshift_CreateAuthenticationProfile(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateAuthenticationProfileInput{
		// AuthenticationProfileContent: *string, // Required
		// AuthenticationProfileName: *string, // Required
	}

	if len(_redshiftAuthenticationProfileContent) > 0 {
		input.AuthenticationProfileContent = aws.String(_redshiftAuthenticationProfileContent)
	}
	if len(_redshiftAuthenticationProfileName) > 0 {
		input.AuthenticationProfileName = aws.String(_redshiftAuthenticationProfileName)
	}

	if resp, err := client.CreateAuthenticationProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new cluster with the specified parameters.
// To create a cluster in Virtual Private Cloud (VPC), you must provide a cluster
// subnet group name. The cluster subnet group identifies the subnets of your VPC
// that Amazon Redshift uses when creating the cluster. For more information about
// managing clusters, go to [Amazon Redshift Clusters]in the Amazon Redshift Cluster Management Guide.
//
// VPC Block Public Access (BPA) enables you to block resources in VPCs and
// subnets that you own in a Region from reaching or being reached from the
// internet through internet gateways and egress-only internet gateways. If a
// subnet group for a provisioned cluster is in an account with VPC BPA turned on,
// the following capabilities are blocked:
//
// - Creating a public cluster
//
// - Restoring a public cluster
//
// - Modifying a private cluster to be public
//
// - Adding a subnet with VPC BPA turned on to the subnet group when there's at
// least one public cluster within the group
//
// For more information about VPC BPA, see [Block public access to VPCs and subnets] in the Amazon VPC User Guide.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
func redshift_CreateCluster(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateClusterInput{
		// ClusterIdentifier: *string, // Required
		// MasterUsername: *string, // Required
		// NodeType: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftMasterUsername) > 0 {
		input.MasterUsername = aws.String(_redshiftMasterUsername)
	}
	if len(_redshiftNodeType) > 0 {
		input.NodeType = aws.String(_redshiftNodeType)
	}
	if len(_redshiftAdditionalInfo) > 0 {
		input.AdditionalInfo = aws.String(_redshiftAdditionalInfo)
	}
	if len(_redshiftAllowVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowVersionUpgrade", _redshiftAllowVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_redshiftAquaConfigurationStatus) > 0 {
		if err := assignInputField(input, "AquaConfigurationStatus", _redshiftAquaConfigurationStatus); err != nil {
			log.Errorf("invalid --aqua-configuration-status: %s", err.Error())
			return
		}
	}
	if len(_redshiftAutomatedSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "AutomatedSnapshotRetentionPeriod", _redshiftAutomatedSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --automated-snapshot-retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_redshiftAvailabilityZone)
	}
	if len(_redshiftAvailabilityZoneRelocation) > 0 {
		if err := assignInputField(input, "AvailabilityZoneRelocation", _redshiftAvailabilityZoneRelocation); err != nil {
			log.Errorf("invalid --availability-zone-relocation: %s", err.Error())
			return
		}
	}
	if len(_redshiftCatalogName) > 0 {
		input.CatalogName = aws.String(_redshiftCatalogName)
	}
	if len(_redshiftClusterParameterGroupName) > 0 {
		input.ClusterParameterGroupName = aws.String(_redshiftClusterParameterGroupName)
	}
	if len(_redshiftClusterSecurityGroups) > 0 {
		input.ClusterSecurityGroups = append([]string(nil), _redshiftClusterSecurityGroups...)
	}
	if len(_redshiftClusterSubnetGroupName) > 0 {
		input.ClusterSubnetGroupName = aws.String(_redshiftClusterSubnetGroupName)
	}
	if len(_redshiftClusterType) > 0 {
		input.ClusterType = aws.String(_redshiftClusterType)
	}
	if len(_redshiftClusterVersion) > 0 {
		input.ClusterVersion = aws.String(_redshiftClusterVersion)
	}
	if len(_redshiftDBName) > 0 {
		input.DBName = aws.String(_redshiftDBName)
	}
	if len(_redshiftDefaultIamRoleArn) > 0 {
		input.DefaultIamRoleArn = aws.String(_redshiftDefaultIamRoleArn)
	}
	if len(_redshiftElasticIp) > 0 {
		input.ElasticIp = aws.String(_redshiftElasticIp)
	}
	if len(_redshiftEncrypted) > 0 {
		if err := assignInputField(input, "Encrypted", _redshiftEncrypted); err != nil {
			log.Errorf("invalid --encrypted: %s", err.Error())
			return
		}
	}
	if len(_redshiftEnhancedVpcRouting) > 0 {
		if err := assignInputField(input, "EnhancedVpcRouting", _redshiftEnhancedVpcRouting); err != nil {
			log.Errorf("invalid --enhanced-vpc-routing: %s", err.Error())
			return
		}
	}
	if len(_redshiftExtraComputeForAutomaticOptimization) > 0 {
		if err := assignInputField(input, "ExtraComputeForAutomaticOptimization", _redshiftExtraComputeForAutomaticOptimization); err != nil {
			log.Errorf("invalid --extra-compute-for-automatic-optimization: %s", err.Error())
			return
		}
	}
	if len(_redshiftHsmClientCertificateIdentifier) > 0 {
		input.HsmClientCertificateIdentifier = aws.String(_redshiftHsmClientCertificateIdentifier)
	}
	if len(_redshiftHsmConfigurationIdentifier) > 0 {
		input.HsmConfigurationIdentifier = aws.String(_redshiftHsmConfigurationIdentifier)
	}
	if len(_redshiftIamRoles) > 0 {
		input.IamRoles = append([]string(nil), _redshiftIamRoles...)
	}
	if len(_redshiftIpAddressType) > 0 {
		input.IpAddressType = aws.String(_redshiftIpAddressType)
	}
	if len(_redshiftKMSKeyId) > 0 {
		input.KmsKeyId = aws.String(_redshiftKMSKeyId)
	}
	if len(_redshiftLoadSampleData) > 0 {
		input.LoadSampleData = aws.String(_redshiftLoadSampleData)
	}
	if len(_redshiftMaintenanceTrackName) > 0 {
		input.MaintenanceTrackName = aws.String(_redshiftMaintenanceTrackName)
	}
	if len(_redshiftManageMasterPassword) > 0 {
		if err := assignInputField(input, "ManageMasterPassword", _redshiftManageMasterPassword); err != nil {
			log.Errorf("invalid --manage-master-password: %s", err.Error())
			return
		}
	}
	if len(_redshiftManualSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "ManualSnapshotRetentionPeriod", _redshiftManualSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --manual-snapshot-retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftMasterPasswordSecretKmsKeyId) > 0 {
		input.MasterPasswordSecretKmsKeyId = aws.String(_redshiftMasterPasswordSecretKmsKeyId)
	}
	if len(_redshiftMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_redshiftMasterUserPassword)
	}
	if len(_redshiftMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _redshiftMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_redshiftNumberOfNodes) > 0 {
		if err := assignInputField(input, "NumberOfNodes", _redshiftNumberOfNodes); err != nil {
			log.Errorf("invalid --number-of-nodes: %s", err.Error())
			return
		}
	}
	if len(_redshiftPort) > 0 {
		if err := assignInputField(input, "Port", _redshiftPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_redshiftPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_redshiftPreferredMaintenanceWindow)
	}
	if len(_redshiftPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _redshiftPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_redshiftRedshiftIdcApplicationArn) > 0 {
		input.RedshiftIdcApplicationArn = aws.String(_redshiftRedshiftIdcApplicationArn)
	}
	if len(_redshiftSnapshotScheduleIdentifier) > 0 {
		input.SnapshotScheduleIdentifier = aws.String(_redshiftSnapshotScheduleIdentifier)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_redshiftVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _redshiftVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Redshift parameter group.
// Creating parameter groups is independent of creating clusters. You can
// associate a cluster with a parameter group when you create the cluster. You can
// also associate an existing cluster with a parameter group after the cluster is
// created by using ModifyCluster.
//
// Parameters in the parameter group define specific behavior that applies to the
// databases you create on the cluster. For more information about parameters and
// parameter groups, go to [Amazon Redshift Parameter Groups]in the Amazon Redshift Cluster Management Guide.
//
// [Amazon Redshift Parameter Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html
func redshift_CreateClusterParameterGroup(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateClusterParameterGroupInput{
		// Description: *string, // Required
		// ParameterGroupFamily: *string, // Required
		// ParameterGroupName: *string, // Required
	}

	if len(_redshiftDescription) > 0 {
		input.Description = aws.String(_redshiftDescription)
	}
	if len(_redshiftParameterGroupFamily) > 0 {
		input.ParameterGroupFamily = aws.String(_redshiftParameterGroupFamily)
	}
	if len(_redshiftParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_redshiftParameterGroupName)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Redshift security group. You use security groups to
// control access to non-VPC clusters.
//
// For information about managing security groups, go to [Amazon Redshift Cluster Security Groups] in the Amazon Redshift
// Cluster Management Guide.
//
// [Amazon Redshift Cluster Security Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html
func redshift_CreateClusterSecurityGroup(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateClusterSecurityGroupInput{
		// ClusterSecurityGroupName: *string, // Required
		// Description: *string, // Required
	}

	if len(_redshiftClusterSecurityGroupName) > 0 {
		input.ClusterSecurityGroupName = aws.String(_redshiftClusterSecurityGroupName)
	}
	if len(_redshiftDescription) > 0 {
		input.Description = aws.String(_redshiftDescription)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateClusterSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a manual snapshot of the specified cluster. The cluster must be in the
// available state.
//
// For more information about working with snapshots, go to [Amazon Redshift Snapshots] in the Amazon
// Redshift Cluster Management Guide.
//
// [Amazon Redshift Snapshots]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html
func redshift_CreateClusterSnapshot(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateClusterSnapshotInput{
		// ClusterIdentifier: *string, // Required
		// SnapshotIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_redshiftSnapshotIdentifier)
	}
	if len(_redshiftManualSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "ManualSnapshotRetentionPeriod", _redshiftManualSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --manual-snapshot-retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Redshift subnet group. You must provide a list of one or
// more subnets in your existing Amazon Virtual Private Cloud (Amazon VPC) when
// creating Amazon Redshift subnet group.
//
// For information about subnet groups, go to [Amazon Redshift Cluster Subnet Groups] in the Amazon Redshift Cluster
// Management Guide.
//
// [Amazon Redshift Cluster Subnet Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-cluster-subnet-groups.html
func redshift_CreateClusterSubnetGroup(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateClusterSubnetGroupInput{
		// ClusterSubnetGroupName: *string, // Required
		// Description: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_redshiftClusterSubnetGroupName) > 0 {
		input.ClusterSubnetGroupName = aws.String(_redshiftClusterSubnetGroupName)
	}
	if len(_redshiftDescription) > 0 {
		input.Description = aws.String(_redshiftDescription)
	}
	if len(_redshiftSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _redshiftSubnetIds...)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateClusterSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to create a custom domain name for a cluster. Properties include the
// custom domain name, the cluster the custom domain is associated with, and the
// certificate Amazon Resource Name (ARN).
func redshift_CreateCustomDomainAssociation(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateCustomDomainAssociationInput{
		// ClusterIdentifier: *string, // Required
		// CustomDomainCertificateArn: *string, // Required
		// CustomDomainName: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftCustomDomainCertificateArn) > 0 {
		input.CustomDomainCertificateArn = aws.String(_redshiftCustomDomainCertificateArn)
	}
	if len(_redshiftCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftCustomDomainName)
	}

	if resp, err := client.CreateCustomDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Redshift-managed VPC endpoint.
func redshift_CreateEndpointAccess(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateEndpointAccessInput{
		// EndpointName: *string, // Required
		// SubnetGroupName: *string, // Required
	}

	if len(_redshiftEndpointName) > 0 {
		input.EndpointName = aws.String(_redshiftEndpointName)
	}
	if len(_redshiftSubnetGroupName) > 0 {
		input.SubnetGroupName = aws.String(_redshiftSubnetGroupName)
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftResourceOwner) > 0 {
		input.ResourceOwner = aws.String(_redshiftResourceOwner)
	}
	if len(_redshiftVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _redshiftVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Redshift event notification subscription. This action
// requires an ARN (Amazon Resource Name) of an Amazon SNS topic created by either
// the Amazon Redshift console, the Amazon SNS console, or the Amazon SNS API. To
// obtain an ARN with Amazon SNS, you must create a topic in Amazon SNS and
// subscribe to the topic. The ARN is displayed in the SNS console.
//
// You can specify the source type, and lists of Amazon Redshift source IDs, event
// categories, and event severities. Notifications will be sent for all events you
// want that match those criteria. For example, you can specify source type =
// cluster, source ID = my-cluster-1 and mycluster2, event categories =
// Availability, Backup, and severity = ERROR. The subscription will only send
// notifications for those ERROR events in the Availability and Backup categories
// for the specified clusters.
//
// If you specify both the source type and source IDs, such as source type =
// cluster and source identifier = my-cluster-1, notifications will be sent for all
// the cluster events for my-cluster-1. If you specify a source type but do not
// specify a source identifier, you will receive notice of the events for the
// objects of that type in your Amazon Web Services account. If you do not specify
// either the SourceType nor the SourceIdentifier, you will be notified of events
// generated from all Amazon Redshift sources belonging to your Amazon Web Services
// account. You must specify a source type if you specify a source ID.
func redshift_CreateEventSubscription(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateEventSubscriptionInput{
		// SnsTopicArn: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_redshiftSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_redshiftSnsTopicArn)
	}
	if len(_redshiftSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_redshiftSubscriptionName)
	}
	if len(_redshiftEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _redshiftEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_redshiftEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _redshiftEventCategories...)
	}
	if len(_redshiftSeverity) > 0 {
		input.Severity = aws.String(_redshiftSeverity)
	}
	if len(_redshiftSourceIds) > 0 {
		input.SourceIds = append([]string(nil), _redshiftSourceIds...)
	}
	if len(_redshiftSourceType) > 0 {
		input.SourceType = aws.String(_redshiftSourceType)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an HSM client certificate that an Amazon Redshift cluster will use to
// connect to the client's HSM in order to store and retrieve the keys used to
// encrypt the cluster databases.
//
// The command returns a public key, which you must store in the HSM. In addition
// to creating the HSM certificate, you must create an Amazon Redshift HSM
// configuration that provides a cluster the information needed to store and use
// encryption keys in the HSM. For more information, go to [Hardware Security Modules]in the Amazon Redshift
// Cluster Management Guide.
//
// [Hardware Security Modules]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html#working-with-HSM
func redshift_CreateHsmClientCertificate(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateHsmClientCertificateInput{
		// HsmClientCertificateIdentifier: *string, // Required
	}

	if len(_redshiftHsmClientCertificateIdentifier) > 0 {
		input.HsmClientCertificateIdentifier = aws.String(_redshiftHsmClientCertificateIdentifier)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHsmClientCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an HSM configuration that contains the information required by an
// Amazon Redshift cluster to store and use database encryption keys in a Hardware
// Security Module (HSM). After creating the HSM configuration, you can specify it
// as a parameter when creating a cluster. The cluster will then store its
// encryption keys in the HSM.
//
// In addition to creating an HSM configuration, you must also create an HSM
// client certificate. For more information, go to [Hardware Security Modules]in the Amazon Redshift Cluster
// Management Guide.
//
// [Hardware Security Modules]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-HSM.html
func redshift_CreateHsmConfiguration(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateHsmConfigurationInput{
		// Description: *string, // Required
		// HsmConfigurationIdentifier: *string, // Required
		// HsmIpAddress: *string, // Required
		// HsmPartitionName: *string, // Required
		// HsmPartitionPassword: *string, // Required
		// HsmServerPublicCertificate: *string, // Required
	}

	if len(_redshiftDescription) > 0 {
		input.Description = aws.String(_redshiftDescription)
	}
	if len(_redshiftHsmConfigurationIdentifier) > 0 {
		input.HsmConfigurationIdentifier = aws.String(_redshiftHsmConfigurationIdentifier)
	}
	if len(_redshiftHsmIpAddress) > 0 {
		input.HsmIpAddress = aws.String(_redshiftHsmIpAddress)
	}
	if len(_redshiftHsmPartitionName) > 0 {
		input.HsmPartitionName = aws.String(_redshiftHsmPartitionName)
	}
	if len(_redshiftHsmPartitionPassword) > 0 {
		input.HsmPartitionPassword = aws.String(_redshiftHsmPartitionPassword)
	}
	if len(_redshiftHsmServerPublicCertificate) > 0 {
		input.HsmServerPublicCertificate = aws.String(_redshiftHsmServerPublicCertificate)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHsmConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a zero-ETL integration or S3 event integration with Amazon Redshift.
func redshift_CreateIntegration(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateIntegrationInput{
		// IntegrationName: *string, // Required
		// SourceArn: *string, // Required
		// TargetArn: *string, // Required
	}

	if len(_redshiftIntegrationName) > 0 {
		input.IntegrationName = aws.String(_redshiftIntegrationName)
	}
	if len(_redshiftSourceArn) > 0 {
		input.SourceArn = aws.String(_redshiftSourceArn)
	}
	if len(_redshiftTargetArn) > 0 {
		input.TargetArn = aws.String(_redshiftTargetArn)
	}
	if len(_redshiftAdditionalEncryptionContext) > 0 {
		if err := assignInputField(input, "AdditionalEncryptionContext", _redshiftAdditionalEncryptionContext); err != nil {
			log.Errorf("invalid --additional-encryption-context: %s", err.Error())
			return
		}
	}
	if len(_redshiftDescription) > 0 {
		input.Description = aws.String(_redshiftDescription)
	}
	if len(_redshiftKMSKeyId) > 0 {
		input.KMSKeyId = aws.String(_redshiftKMSKeyId)
	}
	if len(_redshiftTagList) > 0 {
		if err := assignInputField(input, "TagList", _redshiftTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Redshift application for use with IAM Identity Center.
func redshift_CreateRedshiftIdcApplication(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateRedshiftIdcApplicationInput{
		// IamRoleArn: *string, // Required
		// IdcDisplayName: *string, // Required
		// IdcInstanceArn: *string, // Required
		// RedshiftIdcApplicationName: *string, // Required
	}

	if len(_redshiftIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_redshiftIamRoleArn)
	}
	if len(_redshiftIdcDisplayName) > 0 {
		input.IdcDisplayName = aws.String(_redshiftIdcDisplayName)
	}
	if len(_redshiftIdcInstanceArn) > 0 {
		input.IdcInstanceArn = aws.String(_redshiftIdcInstanceArn)
	}
	if len(_redshiftRedshiftIdcApplicationName) > 0 {
		input.RedshiftIdcApplicationName = aws.String(_redshiftRedshiftIdcApplicationName)
	}
	if len(_redshiftApplicationType) > 0 {
		if err := assignInputField(input, "ApplicationType", _redshiftApplicationType); err != nil {
			log.Errorf("invalid --application-type: %s", err.Error())
			return
		}
	}
	if len(_redshiftAuthorizedTokenIssuerList) > 0 {
		if err := assignInputField(input, "AuthorizedTokenIssuerList", _redshiftAuthorizedTokenIssuerList); err != nil {
			log.Errorf("invalid --authorized-token-issuer-list: %s", err.Error())
			return
		}
	}
	if len(_redshiftIdentityNamespace) > 0 {
		input.IdentityNamespace = aws.String(_redshiftIdentityNamespace)
	}
	if len(_redshiftServiceIntegrations) > 0 {
		if err := assignInputField(input, "ServiceIntegrations", _redshiftServiceIntegrations); err != nil {
			log.Errorf("invalid --service-integrations: %s", err.Error())
			return
		}
	}
	if len(_redshiftSsoTagKeys) > 0 {
		input.SsoTagKeys = append([]string(nil), _redshiftSsoTagKeys...)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRedshiftIdcApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a scheduled action. A scheduled action contains a schedule and an
// Amazon Redshift API action. For example, you can create a schedule of when to
// run the ResizeCluster API operation.
func redshift_CreateScheduledAction(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateScheduledActionInput{
		// IamRole: *string, // Required
		// Schedule: *string, // Required
		// ScheduledActionName: *string, // Required
		// TargetAction: *types.ScheduledActionType, // Required
	}

	if len(_redshiftIamRole) > 0 {
		input.IamRole = aws.String(_redshiftIamRole)
	}
	if len(_redshiftSchedule) > 0 {
		input.Schedule = aws.String(_redshiftSchedule)
	}
	if len(_redshiftScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_redshiftScheduledActionName)
	}
	if len(_redshiftTargetAction) > 0 {
		if err := assignInputField(input, "TargetAction", _redshiftTargetAction); err != nil {
			log.Errorf("invalid --target-action: %s", err.Error())
			return
		}
	}
	if len(_redshiftEnable) > 0 {
		if err := assignInputField(input, "Enable", _redshiftEnable); err != nil {
			log.Errorf("invalid --enable: %s", err.Error())
			return
		}
	}
	if len(_redshiftEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _redshiftEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftScheduledActionDescription) > 0 {
		input.ScheduledActionDescription = aws.String(_redshiftScheduledActionDescription)
	}
	if len(_redshiftStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _redshiftStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot copy grant that permits Amazon Redshift to use an encrypted
// symmetric key from Key Management Service (KMS) to encrypt copied snapshots in a
// destination region.
//
// For more information about managing snapshot copy grants, go to [Amazon Redshift Database Encryption] in the Amazon
// Redshift Cluster Management Guide.
//
// [Amazon Redshift Database Encryption]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html
func redshift_CreateSnapshotCopyGrant(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateSnapshotCopyGrantInput{
		// SnapshotCopyGrantName: *string, // Required
	}

	if len(_redshiftSnapshotCopyGrantName) > 0 {
		input.SnapshotCopyGrantName = aws.String(_redshiftSnapshotCopyGrantName)
	}
	if len(_redshiftKMSKeyId) > 0 {
		input.KmsKeyId = aws.String(_redshiftKMSKeyId)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSnapshotCopyGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a snapshot schedule that can be associated to a cluster and which
// overrides the default system backup schedule.
func redshift_CreateSnapshotSchedule(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateSnapshotScheduleInput{}

	if len(_redshiftDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _redshiftDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_redshiftNextInvocations) > 0 {
		if err := assignInputField(input, "NextInvocations", _redshiftNextInvocations); err != nil {
			log.Errorf("invalid --next-invocations: %s", err.Error())
			return
		}
	}
	if len(_redshiftScheduleDefinitions) > 0 {
		input.ScheduleDefinitions = append([]string(nil), _redshiftScheduleDefinitions...)
	}
	if len(_redshiftScheduleDescription) > 0 {
		input.ScheduleDescription = aws.String(_redshiftScheduleDescription)
	}
	if len(_redshiftScheduleIdentifier) > 0 {
		input.ScheduleIdentifier = aws.String(_redshiftScheduleIdentifier)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSnapshotSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to a cluster.
// A resource can have up to 50 tags. If you try to create more than 50 tags for a
// resource, you will receive an error and the attempt will fail.
//
// If you specify a key that already exists for the resource, the value for that
// key will be updated with the new value.
func redshift_CreateTags(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateTagsInput{
		// ResourceName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_redshiftResourceName) > 0 {
		input.ResourceName = aws.String(_redshiftResourceName)
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a usage limit for a specified Amazon Redshift feature on a cluster. The
// usage limit is identified by the returned usage limit identifier.
func redshift_CreateUsageLimit(cfg aws.Config, client *redshift.Client) {
	input := &redshift.CreateUsageLimitInput{
		// Amount: *int64, // Required
		// ClusterIdentifier: *string, // Required
		// FeatureType: types.UsageLimitFeatureType, // Required
		// LimitType: types.UsageLimitLimitType, // Required
	}

	if len(_redshiftAmount) > 0 {
		if err := assignInputField(input, "Amount", _redshiftAmount); err != nil {
			log.Errorf("invalid --amount: %s", err.Error())
			return
		}
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftFeatureType) > 0 {
		if err := assignInputField(input, "FeatureType", _redshiftFeatureType); err != nil {
			log.Errorf("invalid --feature-type: %s", err.Error())
			return
		}
	}
	if len(_redshiftLimitType) > 0 {
		if err := assignInputField(input, "LimitType", _redshiftLimitType); err != nil {
			log.Errorf("invalid --limit-type: %s", err.Error())
			return
		}
	}
	if len(_redshiftBreachAction) > 0 {
		if err := assignInputField(input, "BreachAction", _redshiftBreachAction); err != nil {
			log.Errorf("invalid --breach-action: %s", err.Error())
			return
		}
	}
	if len(_redshiftPeriod) > 0 {
		if err := assignInputField(input, "Period", _redshiftPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_redshiftTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUsageLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// From a datashare producer account, removes authorization from the specified
// datashare.
func redshift_DeauthorizeDataShare(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeauthorizeDataShareInput{
		// ConsumerIdentifier: *string, // Required
		// DataShareArn: *string, // Required
	}

	if len(_redshiftConsumerIdentifier) > 0 {
		input.ConsumerIdentifier = aws.String(_redshiftConsumerIdentifier)
	}
	if len(_redshiftDataShareArn) > 0 {
		input.DataShareArn = aws.String(_redshiftDataShareArn)
	}

	if resp, err := client.DeauthorizeDataShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an authentication profile.
func redshift_DeleteAuthenticationProfile(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteAuthenticationProfileInput{
		// AuthenticationProfileName: *string, // Required
	}

	if len(_redshiftAuthenticationProfileName) > 0 {
		input.AuthenticationProfileName = aws.String(_redshiftAuthenticationProfileName)
	}

	if resp, err := client.DeleteAuthenticationProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a previously provisioned cluster without its final snapshot being
// created. A successful response from the web service indicates that the request
// was received correctly. Use DescribeClustersto monitor the status of the deletion. The delete
// operation cannot be canceled or reverted once submitted. For more information
// about managing clusters, go to [Amazon Redshift Clusters]in the Amazon Redshift Cluster Management Guide.
//
// If you want to shut down the cluster and retain it for future use, set
// SkipFinalClusterSnapshot to false and specify a name for
// FinalClusterSnapshotIdentifier. You can later restore this snapshot to resume
// using the cluster. If a final cluster snapshot is requested, the status of the
// cluster will be "final-snapshot" while the snapshot is being taken, then it's
// "deleting" once Amazon Redshift begins deleting the cluster.
//
// For more information about managing clusters, go to [Amazon Redshift Clusters] in the Amazon Redshift
// Cluster Management Guide.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
func redshift_DeleteCluster(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteClusterInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftFinalClusterSnapshotIdentifier) > 0 {
		input.FinalClusterSnapshotIdentifier = aws.String(_redshiftFinalClusterSnapshotIdentifier)
	}
	if len(_redshiftFinalClusterSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "FinalClusterSnapshotRetentionPeriod", _redshiftFinalClusterSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --final-cluster-snapshot-retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftSkipFinalClusterSnapshot) > 0 {
		if err := assignInputField(input, "SkipFinalClusterSnapshot", _redshiftSkipFinalClusterSnapshot); err != nil {
			log.Errorf("invalid --skip-final-cluster-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified Amazon Redshift parameter group.
// You cannot delete a parameter group if it is associated with a cluster.
func redshift_DeleteClusterParameterGroup(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteClusterParameterGroupInput{
		// ParameterGroupName: *string, // Required
	}

	if len(_redshiftParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_redshiftParameterGroupName)
	}

	if resp, err := client.DeleteClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Redshift security group.
// You cannot delete a security group that is associated with any clusters. You
// cannot delete the default security group.
//
// For information about managing security groups, go to [Amazon Redshift Cluster Security Groups] in the Amazon Redshift
// Cluster Management Guide.
//
// [Amazon Redshift Cluster Security Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html
func redshift_DeleteClusterSecurityGroup(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteClusterSecurityGroupInput{
		// ClusterSecurityGroupName: *string, // Required
	}

	if len(_redshiftClusterSecurityGroupName) > 0 {
		input.ClusterSecurityGroupName = aws.String(_redshiftClusterSecurityGroupName)
	}

	if resp, err := client.DeleteClusterSecurityGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified manual snapshot. The snapshot must be in the available
// state, with no other users authorized to access the snapshot.
//
// Unlike automated snapshots, manual snapshots are retained even after you delete
// your cluster. Amazon Redshift does not delete your manual snapshots. You must
// delete manual snapshot explicitly to avoid getting charged. If other accounts
// are authorized to access the snapshot, you must revoke all of the authorizations
// before you can delete the snapshot.
func redshift_DeleteClusterSnapshot(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteClusterSnapshotInput{
		// SnapshotIdentifier: *string, // Required
	}

	if len(_redshiftSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_redshiftSnapshotIdentifier)
	}
	if len(_redshiftSnapshotClusterIdentifier) > 0 {
		input.SnapshotClusterIdentifier = aws.String(_redshiftSnapshotClusterIdentifier)
	}

	if resp, err := client.DeleteClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified cluster subnet group.
func redshift_DeleteClusterSubnetGroup(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteClusterSubnetGroupInput{
		// ClusterSubnetGroupName: *string, // Required
	}

	if len(_redshiftClusterSubnetGroupName) > 0 {
		input.ClusterSubnetGroupName = aws.String(_redshiftClusterSubnetGroupName)
	}

	if resp, err := client.DeleteClusterSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Contains information about deleting a custom domain association for a cluster.
func redshift_DeleteCustomDomainAssociation(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteCustomDomainAssociationInput{
		// ClusterIdentifier: *string, // Required
		// CustomDomainName: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftCustomDomainName)
	}

	if resp, err := client.DeleteCustomDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Redshift-managed VPC endpoint.
func redshift_DeleteEndpointAccess(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteEndpointAccessInput{
		// EndpointName: *string, // Required
	}

	if len(_redshiftEndpointName) > 0 {
		input.EndpointName = aws.String(_redshiftEndpointName)
	}

	if resp, err := client.DeleteEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Redshift event notification subscription.
func redshift_DeleteEventSubscription(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteEventSubscriptionInput{
		// SubscriptionName: *string, // Required
	}

	if len(_redshiftSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_redshiftSubscriptionName)
	}

	if resp, err := client.DeleteEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified HSM client certificate.
func redshift_DeleteHsmClientCertificate(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteHsmClientCertificateInput{
		// HsmClientCertificateIdentifier: *string, // Required
	}

	if len(_redshiftHsmClientCertificateIdentifier) > 0 {
		input.HsmClientCertificateIdentifier = aws.String(_redshiftHsmClientCertificateIdentifier)
	}

	if resp, err := client.DeleteHsmClientCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Amazon Redshift HSM configuration.
func redshift_DeleteHsmConfiguration(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteHsmConfigurationInput{
		// HsmConfigurationIdentifier: *string, // Required
	}

	if len(_redshiftHsmConfigurationIdentifier) > 0 {
		input.HsmConfigurationIdentifier = aws.String(_redshiftHsmConfigurationIdentifier)
	}

	if resp, err := client.DeleteHsmConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a zero-ETL integration or S3 event integration with Amazon Redshift.
func redshift_DeleteIntegration(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteIntegrationInput{
		// IntegrationArn: *string, // Required
	}

	if len(_redshiftIntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_redshiftIntegrationArn)
	}

	if resp, err := client.DeleteIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a partner integration from a cluster. Data can still flow to the
// cluster until the integration is deleted at the partner's website.
func redshift_DeletePartner(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeletePartnerInput{
		// AccountId: *string, // Required
		// ClusterIdentifier: *string, // Required
		// DatabaseName: *string, // Required
		// PartnerName: *string, // Required
	}

	if len(_redshiftAccountId) > 0 {
		input.AccountId = aws.String(_redshiftAccountId)
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftDatabaseName) > 0 {
		input.DatabaseName = aws.String(_redshiftDatabaseName)
	}
	if len(_redshiftPartnerName) > 0 {
		input.PartnerName = aws.String(_redshiftPartnerName)
	}

	if resp, err := client.DeletePartner(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Redshift IAM Identity Center application.
func redshift_DeleteRedshiftIdcApplication(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteRedshiftIdcApplicationInput{
		// RedshiftIdcApplicationArn: *string, // Required
	}

	if len(_redshiftRedshiftIdcApplicationArn) > 0 {
		input.RedshiftIdcApplicationArn = aws.String(_redshiftRedshiftIdcApplicationArn)
	}

	if resp, err := client.DeleteRedshiftIdcApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource policy for a specified resource.
func redshift_DeleteResourcePolicy(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_redshiftResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a scheduled action.
func redshift_DeleteScheduledAction(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteScheduledActionInput{
		// ScheduledActionName: *string, // Required
	}

	if len(_redshiftScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_redshiftScheduledActionName)
	}

	if resp, err := client.DeleteScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified snapshot copy grant.
func redshift_DeleteSnapshotCopyGrant(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteSnapshotCopyGrantInput{
		// SnapshotCopyGrantName: *string, // Required
	}

	if len(_redshiftSnapshotCopyGrantName) > 0 {
		input.SnapshotCopyGrantName = aws.String(_redshiftSnapshotCopyGrantName)
	}

	if resp, err := client.DeleteSnapshotCopyGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a snapshot schedule.
func redshift_DeleteSnapshotSchedule(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteSnapshotScheduleInput{
		// ScheduleIdentifier: *string, // Required
	}

	if len(_redshiftScheduleIdentifier) > 0 {
		input.ScheduleIdentifier = aws.String(_redshiftScheduleIdentifier)
	}

	if resp, err := client.DeleteSnapshotSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes tags from a resource. You must provide the ARN of the resource from
// which you want to delete the tag or tags.
func redshift_DeleteTags(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteTagsInput{
		// ResourceName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_redshiftResourceName) > 0 {
		input.ResourceName = aws.String(_redshiftResourceName)
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}

	if resp, err := client.DeleteTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a usage limit from a cluster.
func redshift_DeleteUsageLimit(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeleteUsageLimitInput{
		// UsageLimitId: *string, // Required
	}

	if len(_redshiftUsageLimitId) > 0 {
		input.UsageLimitId = aws.String(_redshiftUsageLimitId)
	}

	if resp, err := client.DeleteUsageLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters a cluster or serverless namespace from the Amazon Web Services Glue
// Data Catalog.
func redshift_DeregisterNamespace(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DeregisterNamespaceInput{
		// ConsumerIdentifiers: []string, // Required
		// NamespaceIdentifier: types.NamespaceIdentifierUnion, // Required
	}

	if len(_redshiftConsumerIdentifiers) > 0 {
		input.ConsumerIdentifiers = append([]string(nil), _redshiftConsumerIdentifiers...)
	}
	if len(_redshiftNamespaceIdentifier) > 0 {
		if err := assignInputField(input, "NamespaceIdentifier", _redshiftNamespaceIdentifier); err != nil {
			log.Errorf("invalid --namespace-identifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeregisterNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of attributes attached to an account
func redshift_DescribeAccountAttributes(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeAccountAttributesInput{}

	if len(_redshiftAttributeNames) > 0 {
		input.AttributeNames = append([]string(nil), _redshiftAttributeNames...)
	}

	if resp, err := client.DescribeAccountAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an authentication profile.
func redshift_DescribeAuthenticationProfiles(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeAuthenticationProfilesInput{}

	if len(_redshiftAuthenticationProfileName) > 0 {
		input.AuthenticationProfileName = aws.String(_redshiftAuthenticationProfileName)
	}

	if resp, err := client.DescribeAuthenticationProfiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an array of ClusterDbRevision objects.
func redshift_DescribeClusterDbRevisions(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeClusterDbRevisionsInput{}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusterDbRevisions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeClusterDbRevisionsOutput
	p := redshift.NewDescribeClusterDbRevisionsPaginator(client, input)
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

// Returns a list of Amazon Redshift parameter groups, including parameter groups
// you created and the default parameter group. For each parameter group, the
// response includes the parameter group name, description, and parameter group
// family name. You can optionally specify a name to retrieve the description of a
// specific parameter group.
//
// For more information about parameters and parameter groups, go to [Amazon Redshift Parameter Groups] in the
// Amazon Redshift Cluster Management Guide.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all parameter groups that match any combination of the
// specified keys and values. For example, if you have owner and environment for
// tag keys, and admin and test for tag values, all parameter groups that have any
// combination of those values are returned.
//
// If both tag keys and values are omitted from the request, parameter groups are
// returned regardless of whether they have tag keys or values associated with
// them.
//
// [Amazon Redshift Parameter Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html
func redshift_DescribeClusterParameterGroups(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeClusterParameterGroupsInput{}

	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_redshiftParameterGroupName)
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusterParameterGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeClusterParameterGroupsOutput
	p := redshift.NewDescribeClusterParameterGroupsPaginator(client, input)
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

// Returns a detailed list of parameters contained within the specified Amazon
// Redshift parameter group. For each parameter the response includes information
// such as parameter name, description, data type, value, whether the parameter
// value is modifiable, and so on.
//
// You can specify source filter to retrieve parameters of only specific type. For
// example, to retrieve parameters that were modified by a user action such as from
// ModifyClusterParameterGroup, you can specify source equal to user.
//
// For more information about parameters and parameter groups, go to [Amazon Redshift Parameter Groups] in the
// Amazon Redshift Cluster Management Guide.
//
// [Amazon Redshift Parameter Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html
func redshift_DescribeClusterParameters(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeClusterParametersInput{
		// ParameterGroupName: *string, // Required
	}

	if len(_redshiftParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_redshiftParameterGroupName)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftSource) > 0 {
		input.Source = aws.String(_redshiftSource)
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusterParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeClusterParametersOutput
	p := redshift.NewDescribeClusterParametersPaginator(client, input)
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

// Returns information about Amazon Redshift security groups. If the name of a
// security group is specified, the response will contain only information about
// only that security group.
//
// For information about managing security groups, go to [Amazon Redshift Cluster Security Groups] in the Amazon Redshift
// Cluster Management Guide.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all security groups that match any combination of the specified
// keys and values. For example, if you have owner and environment for tag keys,
// and admin and test for tag values, all security groups that have any
// combination of those values are returned.
//
// If both tag keys and values are omitted from the request, security groups are
// returned regardless of whether they have tag keys or values associated with
// them.
//
// [Amazon Redshift Cluster Security Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html
func redshift_DescribeClusterSecurityGroups(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeClusterSecurityGroupsInput{}

	if len(_redshiftClusterSecurityGroupName) > 0 {
		input.ClusterSecurityGroupName = aws.String(_redshiftClusterSecurityGroupName)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusterSecurityGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeClusterSecurityGroupsOutput
	p := redshift.NewDescribeClusterSecurityGroupsPaginator(client, input)
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

// Returns one or more snapshot objects, which contain metadata about your cluster
// snapshots. By default, this operation returns information about all snapshots of
// all clusters that are owned by your Amazon Web Services account. No information
// is returned for snapshots owned by inactive Amazon Web Services accounts.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all snapshots that match any combination of the specified keys
// and values. For example, if you have owner and environment for tag keys, and
// admin and test for tag values, all snapshots that have any combination of those
// values are returned. Only snapshots that you own are returned in the response;
// shared snapshots are not returned with the tag key and tag value request
// parameters.
//
// If both tag keys and values are omitted from the request, snapshots are
// returned regardless of whether they have tag keys or values associated with
// them.
func redshift_DescribeClusterSnapshots(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeClusterSnapshotsInput{}

	if len(_redshiftClusterExists) > 0 {
		if err := assignInputField(input, "ClusterExists", _redshiftClusterExists); err != nil {
			log.Errorf("invalid --cluster-exists: %s", err.Error())
			return
		}
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _redshiftEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_redshiftOwnerAccount)
	}
	if len(_redshiftSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_redshiftSnapshotArn)
	}
	if len(_redshiftSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_redshiftSnapshotIdentifier)
	}
	if len(_redshiftSnapshotType) > 0 {
		input.SnapshotType = aws.String(_redshiftSnapshotType)
	}
	if len(_redshiftSortingEntities) > 0 {
		if err := assignInputField(input, "SortingEntities", _redshiftSortingEntities); err != nil {
			log.Errorf("invalid --sorting-entities: %s", err.Error())
			return
		}
	}
	if len(_redshiftStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _redshiftStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusterSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeClusterSnapshotsOutput
	p := redshift.NewDescribeClusterSnapshotsPaginator(client, input)
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

// Returns one or more cluster subnet group objects, which contain metadata about
// your cluster subnet groups. By default, this operation returns information about
// all cluster subnet groups that are defined in your Amazon Web Services account.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all subnet groups that match any combination of the specified
// keys and values. For example, if you have owner and environment for tag keys,
// and admin and test for tag values, all subnet groups that have any combination
// of those values are returned.
//
// If both tag keys and values are omitted from the request, subnet groups are
// returned regardless of whether they have tag keys or values associated with
// them.
func redshift_DescribeClusterSubnetGroups(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeClusterSubnetGroupsInput{}

	if len(_redshiftClusterSubnetGroupName) > 0 {
		input.ClusterSubnetGroupName = aws.String(_redshiftClusterSubnetGroupName)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusterSubnetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeClusterSubnetGroupsOutput
	p := redshift.NewDescribeClusterSubnetGroupsPaginator(client, input)
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

// Returns a list of all the available maintenance tracks.
func redshift_DescribeClusterTracks(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeClusterTracksInput{}

	if len(_redshiftMaintenanceTrackName) > 0 {
		input.MaintenanceTrackName = aws.String(_redshiftMaintenanceTrackName)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusterTracks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeClusterTracksOutput
	p := redshift.NewDescribeClusterTracksPaginator(client, input)
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

// Returns descriptions of the available Amazon Redshift cluster versions. You can
// call this operation even before creating any clusters to learn more about the
// Amazon Redshift versions.
//
// For more information about managing clusters, go to [Amazon Redshift Clusters] in the Amazon Redshift
// Cluster Management Guide.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
func redshift_DescribeClusterVersions(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeClusterVersionsInput{}

	if len(_redshiftClusterParameterGroupFamily) > 0 {
		input.ClusterParameterGroupFamily = aws.String(_redshiftClusterParameterGroupFamily)
	}
	if len(_redshiftClusterVersion) > 0 {
		input.ClusterVersion = aws.String(_redshiftClusterVersion)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
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

	var results []*redshift.DescribeClusterVersionsOutput
	p := redshift.NewDescribeClusterVersionsPaginator(client, input)
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

// Returns properties of provisioned clusters including general cluster
// properties, cluster database properties, maintenance and backup properties, and
// security and access properties. This operation supports pagination. For more
// information about managing clusters, go to [Amazon Redshift Clusters]in the Amazon Redshift Cluster
// Management Guide.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all clusters that match any combination of the specified keys
// and values. For example, if you have owner and environment for tag keys, and
// admin and test for tag values, all clusters that have any combination of those
// values are returned.
//
// If both tag keys and values are omitted from the request, clusters are returned
// regardless of whether they have tag keys or values associated with them.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
func redshift_DescribeClusters(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeClustersInput{}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeClustersOutput
	p := redshift.NewDescribeClustersPaginator(client, input)
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

// Contains information about custom domain associations for a cluster.
func redshift_DescribeCustomDomainAssociations(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeCustomDomainAssociationsInput{}

	if len(_redshiftCustomDomainCertificateArn) > 0 {
		input.CustomDomainCertificateArn = aws.String(_redshiftCustomDomainCertificateArn)
	}
	if len(_redshiftCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftCustomDomainName)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeCustomDomainAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeCustomDomainAssociationsOutput
	p := redshift.NewDescribeCustomDomainAssociationsPaginator(client, input)
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

// Shows the status of any inbound or outbound datashares available in the
// specified account.
func redshift_DescribeDataShares(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeDataSharesInput{}

	if len(_redshiftDataShareArn) > 0 {
		input.DataShareArn = aws.String(_redshiftDataShareArn)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDataShares(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeDataSharesOutput
	p := redshift.NewDescribeDataSharesPaginator(client, input)
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

// Returns a list of datashares where the account identifier being called is a
// consumer account identifier.
func redshift_DescribeDataSharesForConsumer(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeDataSharesForConsumerInput{}

	if len(_redshiftConsumerArn) > 0 {
		input.ConsumerArn = aws.String(_redshiftConsumerArn)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftStatus) > 0 {
		if err := assignInputField(input, "Status", _redshiftStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDataSharesForConsumer(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeDataSharesForConsumerOutput
	p := redshift.NewDescribeDataSharesForConsumerPaginator(client, input)
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

// Returns a list of datashares when the account identifier being called is a
// producer account identifier.
func redshift_DescribeDataSharesForProducer(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeDataSharesForProducerInput{}

	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftProducerArn) > 0 {
		input.ProducerArn = aws.String(_redshiftProducerArn)
	}
	if len(_redshiftStatus) > 0 {
		if err := assignInputField(input, "Status", _redshiftStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDataSharesForProducer(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeDataSharesForProducerOutput
	p := redshift.NewDescribeDataSharesForProducerPaginator(client, input)
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

// Returns a list of parameter settings for the specified parameter group family.
// For more information about parameters and parameter groups, go to [Amazon Redshift Parameter Groups] in the
// Amazon Redshift Cluster Management Guide.
//
// [Amazon Redshift Parameter Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html
func redshift_DescribeDefaultClusterParameters(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeDefaultClusterParametersInput{
		// ParameterGroupFamily: *string, // Required
	}

	if len(_redshiftParameterGroupFamily) > 0 {
		input.ParameterGroupFamily = aws.String(_redshiftParameterGroupFamily)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDefaultClusterParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeDefaultClusterParametersOutput
	p := redshift.NewDescribeDefaultClusterParametersPaginator(client, input)
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

// Describes a Redshift-managed VPC endpoint.
func redshift_DescribeEndpointAccess(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeEndpointAccessInput{}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftEndpointName) > 0 {
		input.EndpointName = aws.String(_redshiftEndpointName)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftResourceOwner) > 0 {
		input.ResourceOwner = aws.String(_redshiftResourceOwner)
	}
	if len(_redshiftVpcId) > 0 {
		input.VpcId = aws.String(_redshiftVpcId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEndpointAccess(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeEndpointAccessOutput
	p := redshift.NewDescribeEndpointAccessPaginator(client, input)
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

// Describes an endpoint authorization.
func redshift_DescribeEndpointAuthorization(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeEndpointAuthorizationInput{}

	if len(_redshiftAccount) > 0 {
		input.Account = aws.String(_redshiftAccount)
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftGrantee) > 0 {
		if err := assignInputField(input, "Grantee", _redshiftGrantee); err != nil {
			log.Errorf("invalid --grantee: %s", err.Error())
			return
		}
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEndpointAuthorization(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeEndpointAuthorizationOutput
	p := redshift.NewDescribeEndpointAuthorizationPaginator(client, input)
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

// Displays a list of event categories for all event source types, or for a
// specified source type. For a list of the event categories and source types, go
// to [Amazon Redshift Event Notifications].
//
// [Amazon Redshift Event Notifications]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html
func redshift_DescribeEventCategories(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeEventCategoriesInput{}

	if len(_redshiftSourceType) > 0 {
		input.SourceType = aws.String(_redshiftSourceType)
	}

	if resp, err := client.DescribeEventCategories(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists descriptions of all the Amazon Redshift event notification subscriptions
// for a customer account. If you specify a subscription name, lists the
// description for that subscription.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all event notification subscriptions that match any combination
// of the specified keys and values. For example, if you have owner and environment
// for tag keys, and admin and test for tag values, all subscriptions that have
// any combination of those values are returned.
//
// If both tag keys and values are omitted from the request, subscriptions are
// returned regardless of whether they have tag keys or values associated with
// them.
func redshift_DescribeEventSubscriptions(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeEventSubscriptionsInput{}

	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_redshiftSubscriptionName)
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEventSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeEventSubscriptionsOutput
	p := redshift.NewDescribeEventSubscriptionsPaginator(client, input)
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

// Returns events related to clusters, security groups, snapshots, and parameter
// groups for the past 14 days. Events specific to a particular cluster, security
// group, snapshot or parameter group can be obtained by providing the name as a
// parameter. By default, the past hour of events are returned.
func redshift_DescribeEvents(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeEventsInput{}

	if len(_redshiftDuration) > 0 {
		if err := assignInputField(input, "Duration", _redshiftDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_redshiftEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _redshiftEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_redshiftSourceIdentifier)
	}
	if len(_redshiftSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _redshiftSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_redshiftStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _redshiftStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeEventsOutput
	p := redshift.NewDescribeEventsPaginator(client, input)
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

// Returns information about the specified HSM client certificate. If no
// certificate ID is specified, returns information about all the HSM certificates
// owned by your Amazon Web Services account.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all HSM client certificates that match any combination of the
// specified keys and values. For example, if you have owner and environment for
// tag keys, and admin and test for tag values, all HSM client certificates that
// have any combination of those values are returned.
//
// If both tag keys and values are omitted from the request, HSM client
// certificates are returned regardless of whether they have tag keys or values
// associated with them.
func redshift_DescribeHsmClientCertificates(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeHsmClientCertificatesInput{}

	if len(_redshiftHsmClientCertificateIdentifier) > 0 {
		input.HsmClientCertificateIdentifier = aws.String(_redshiftHsmClientCertificateIdentifier)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeHsmClientCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeHsmClientCertificatesOutput
	p := redshift.NewDescribeHsmClientCertificatesPaginator(client, input)
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

// Returns information about the specified Amazon Redshift HSM configuration. If
// no configuration ID is specified, returns information about all the HSM
// configurations owned by your Amazon Web Services account.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all HSM connections that match any combination of the specified
// keys and values. For example, if you have owner and environment for tag keys,
// and admin and test for tag values, all HSM connections that have any
// combination of those values are returned.
//
// If both tag keys and values are omitted from the request, HSM connections are
// returned regardless of whether they have tag keys or values associated with
// them.
func redshift_DescribeHsmConfigurations(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeHsmConfigurationsInput{}

	if len(_redshiftHsmConfigurationIdentifier) > 0 {
		input.HsmConfigurationIdentifier = aws.String(_redshiftHsmConfigurationIdentifier)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeHsmConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeHsmConfigurationsOutput
	p := redshift.NewDescribeHsmConfigurationsPaginator(client, input)
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

// Returns a list of inbound integrations.
func redshift_DescribeInboundIntegrations(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeInboundIntegrationsInput{}

	if len(_redshiftIntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_redshiftIntegrationArn)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftTargetArn) > 0 {
		input.TargetArn = aws.String(_redshiftTargetArn)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInboundIntegrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeInboundIntegrationsOutput
	p := redshift.NewDescribeInboundIntegrationsPaginator(client, input)
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

// Describes one or more zero-ETL or S3 event integrations with Amazon Redshift.
func redshift_DescribeIntegrations(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeIntegrationsInput{}

	if len(_redshiftFilters) > 0 {
		if err := assignInputField(input, "Filters", _redshiftFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_redshiftIntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_redshiftIntegrationArn)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeIntegrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeIntegrationsOutput
	p := redshift.NewDescribeIntegrationsPaginator(client, input)
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

// Describes whether information, such as queries and connection attempts, is
// being logged for the specified Amazon Redshift cluster.
func redshift_DescribeLoggingStatus(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeLoggingStatusInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}

	if resp, err := client.DescribeLoggingStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns properties of possible node configurations such as node type, number of
// nodes, and disk usage for the specified action type.
func redshift_DescribeNodeConfigurationOptions(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeNodeConfigurationOptionsInput{
		// ActionType: types.ActionType, // Required
	}

	if len(_redshiftActionType) > 0 {
		if err := assignInputField(input, "ActionType", _redshiftActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftFilters) > 0 {
		if err := assignInputField(input, "Filters", _redshiftFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_redshiftOwnerAccount)
	}
	if len(_redshiftSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_redshiftSnapshotArn)
	}
	if len(_redshiftSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_redshiftSnapshotIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.DescribeNodeConfigurationOptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeNodeConfigurationOptionsOutput
	p := redshift.NewDescribeNodeConfigurationOptionsPaginator(client, input)
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

// Returns a list of orderable cluster options. Before you create a new cluster
// you can use this operation to find what options are available, such as the EC2
// Availability Zones (AZ) in the specific Amazon Web Services Region that you can
// specify, and the node types you can request. The node types differ by available
// storage, memory, CPU and price. With the cost involved you might want to obtain
// a list of cluster options in the specific region and specify values when
// creating a cluster. For more information about managing clusters, go to [Amazon Redshift Clusters]in the
// Amazon Redshift Cluster Management Guide.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
func redshift_DescribeOrderableClusterOptions(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeOrderableClusterOptionsInput{}

	if len(_redshiftClusterVersion) > 0 {
		input.ClusterVersion = aws.String(_redshiftClusterVersion)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftNodeType) > 0 {
		input.NodeType = aws.String(_redshiftNodeType)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOrderableClusterOptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeOrderableClusterOptionsOutput
	p := redshift.NewDescribeOrderableClusterOptionsPaginator(client, input)
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

// Returns information about the partner integrations defined for a cluster.
func redshift_DescribePartners(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribePartnersInput{
		// AccountId: *string, // Required
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftAccountId) > 0 {
		input.AccountId = aws.String(_redshiftAccountId)
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftDatabaseName) > 0 {
		input.DatabaseName = aws.String(_redshiftDatabaseName)
	}
	if len(_redshiftPartnerName) > 0 {
		input.PartnerName = aws.String(_redshiftPartnerName)
	}

	if resp, err := client.DescribePartners(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Amazon Redshift IAM Identity Center applications.
func redshift_DescribeRedshiftIdcApplications(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeRedshiftIdcApplicationsInput{}

	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftRedshiftIdcApplicationArn) > 0 {
		input.RedshiftIdcApplicationArn = aws.String(_redshiftRedshiftIdcApplicationArn)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRedshiftIdcApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeRedshiftIdcApplicationsOutput
	p := redshift.NewDescribeRedshiftIdcApplicationsPaginator(client, input)
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

// Returns exchange status details and associated metadata for a reserved-node
// exchange. Statuses include such values as in progress and requested.
func redshift_DescribeReservedNodeExchangeStatus(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeReservedNodeExchangeStatusInput{}

	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftReservedNodeExchangeRequestId) > 0 {
		input.ReservedNodeExchangeRequestId = aws.String(_redshiftReservedNodeExchangeRequestId)
	}
	if len(_redshiftReservedNodeId) > 0 {
		input.ReservedNodeId = aws.String(_redshiftReservedNodeId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedNodeExchangeStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeReservedNodeExchangeStatusOutput
	p := redshift.NewDescribeReservedNodeExchangeStatusPaginator(client, input)
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

// Returns a list of the available reserved node offerings by Amazon Redshift with
// their descriptions including the node type, the fixed and recurring costs of
// reserving the node and duration the node will be reserved for you. These
// descriptions help you determine which reserve node offering you want to
// purchase. You then use the unique offering ID in you call to PurchaseReservedNodeOfferingto reserve one or
// more nodes for your Amazon Redshift cluster.
//
// For more information about reserved node offerings, go to [Purchasing Reserved Nodes] in the Amazon
// Redshift Cluster Management Guide.
//
// [Purchasing Reserved Nodes]: https://docs.aws.amazon.com/redshift/latest/mgmt/purchase-reserved-node-instance.html
func redshift_DescribeReservedNodeOfferings(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeReservedNodeOfferingsInput{}

	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftReservedNodeOfferingId) > 0 {
		input.ReservedNodeOfferingId = aws.String(_redshiftReservedNodeOfferingId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedNodeOfferings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeReservedNodeOfferingsOutput
	p := redshift.NewDescribeReservedNodeOfferingsPaginator(client, input)
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

// Returns the descriptions of the reserved nodes.
func redshift_DescribeReservedNodes(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeReservedNodesInput{}

	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftReservedNodeId) > 0 {
		input.ReservedNodeId = aws.String(_redshiftReservedNodeId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeReservedNodesOutput
	p := redshift.NewDescribeReservedNodesPaginator(client, input)
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

// Returns information about the last resize operation for the specified cluster.
// If no resize operation has ever been initiated for the specified cluster, a
// HTTP 404 error is returned. If a resize operation was initiated and completed,
// the status of the resize remains as SUCCEEDED until the next resize.
//
// A resize operation can be requested using ModifyCluster and specifying a different number or
// type of nodes for the cluster.
func redshift_DescribeResize(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeResizeInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}

	if resp, err := client.DescribeResize(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes properties of scheduled actions.
func redshift_DescribeScheduledActions(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeScheduledActionsInput{}

	if len(_redshiftActive) > 0 {
		if err := assignInputField(input, "Active", _redshiftActive); err != nil {
			log.Errorf("invalid --active: %s", err.Error())
			return
		}
	}
	if len(_redshiftEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _redshiftEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftFilters) > 0 {
		if err := assignInputField(input, "Filters", _redshiftFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_redshiftScheduledActionName)
	}
	if len(_redshiftStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _redshiftStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftTargetActionType) > 0 {
		if err := assignInputField(input, "TargetActionType", _redshiftTargetActionType); err != nil {
			log.Errorf("invalid --target-action-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeScheduledActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeScheduledActionsOutput
	p := redshift.NewDescribeScheduledActionsPaginator(client, input)
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

// Returns a list of snapshot copy grants owned by the Amazon Web Services account
// in the destination region.
//
// For more information about managing snapshot copy grants, go to [Amazon Redshift Database Encryption] in the Amazon
// Redshift Cluster Management Guide.
//
// [Amazon Redshift Database Encryption]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html
func redshift_DescribeSnapshotCopyGrants(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeSnapshotCopyGrantsInput{}

	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftSnapshotCopyGrantName) > 0 {
		input.SnapshotCopyGrantName = aws.String(_redshiftSnapshotCopyGrantName)
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSnapshotCopyGrants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeSnapshotCopyGrantsOutput
	p := redshift.NewDescribeSnapshotCopyGrantsPaginator(client, input)
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

// Returns a list of snapshot schedules.
func redshift_DescribeSnapshotSchedules(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeSnapshotSchedulesInput{}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftScheduleIdentifier) > 0 {
		input.ScheduleIdentifier = aws.String(_redshiftScheduleIdentifier)
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSnapshotSchedules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeSnapshotSchedulesOutput
	p := redshift.NewDescribeSnapshotSchedulesPaginator(client, input)
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

// Returns account level backups storage size and provisional storage.
func redshift_DescribeStorage(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeStorageInput{}

	if resp, err := client.DescribeStorage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the status of one or more table restore requests made using the RestoreTableFromClusterSnapshot API
// action. If you don't specify a value for the TableRestoreRequestId parameter,
// then DescribeTableRestoreStatus returns the status of all table restore
// requests ordered by the date and time of the request in ascending order.
// Otherwise DescribeTableRestoreStatus returns the status of the table specified
// by TableRestoreRequestId .
func redshift_DescribeTableRestoreStatus(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeTableRestoreStatusInput{}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftTableRestoreRequestId) > 0 {
		input.TableRestoreRequestId = aws.String(_redshiftTableRestoreRequestId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTableRestoreStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeTableRestoreStatusOutput
	p := redshift.NewDescribeTableRestoreStatusPaginator(client, input)
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

// Returns a list of tags. You can return tags from a specific resource by
// specifying an ARN, or you can return all tags for a given type of resource, such
// as clusters, snapshots, and so on.
//
// The following are limitations for DescribeTags :
//
// - You cannot specify an ARN and a resource-type value together in the same
// request.
//
// - You cannot use the MaxRecords and Marker parameters together with the ARN
// parameter.
//
// - The MaxRecords parameter can be a range from 10 to 50 results to return in a
// request.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all resources that match any combination of the specified keys
// and values. For example, if you have owner and environment for tag keys, and
// admin and test for tag values, all resources that have any combination of those
// values are returned.
//
// If both tag keys and values are omitted from the request, resources are
// returned regardless of whether they have tag keys or values associated with
// them.
func redshift_DescribeTags(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeTagsInput{}

	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftResourceName) > 0 {
		input.ResourceName = aws.String(_redshiftResourceName)
	}
	if len(_redshiftResourceType) > 0 {
		input.ResourceType = aws.String(_redshiftResourceType)
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeTagsOutput
	p := redshift.NewDescribeTagsPaginator(client, input)
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

// Shows usage limits on a cluster. Results are filtered based on the combination
// of input usage limit identifier, cluster identifier, and feature type
// parameters:
//
// - If usage limit identifier, cluster identifier, and feature type are not
// provided, then all usage limit objects for the current account in the current
// region are returned.
//
// - If usage limit identifier is provided, then the corresponding usage limit
// object is returned.
//
// - If cluster identifier is provided, then all usage limit objects for the
// specified cluster are returned.
//
// - If cluster identifier and feature type are provided, then all usage limit
// objects for the combination of cluster and feature are returned.
func redshift_DescribeUsageLimits(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DescribeUsageLimitsInput{}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftFeatureType) > 0 {
		if err := assignInputField(input, "FeatureType", _redshiftFeatureType); err != nil {
			log.Errorf("invalid --feature-type: %s", err.Error())
			return
		}
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftTagKeys...)
	}
	if len(_redshiftTagValues) > 0 {
		input.TagValues = append([]string(nil), _redshiftTagValues...)
	}
	if len(_redshiftUsageLimitId) > 0 {
		input.UsageLimitId = aws.String(_redshiftUsageLimitId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeUsageLimits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.DescribeUsageLimitsOutput
	p := redshift.NewDescribeUsageLimitsPaginator(client, input)
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

// Stops logging information, such as queries and connection attempts, for the
// specified Amazon Redshift cluster.
func redshift_DisableLogging(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DisableLoggingInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}

	if resp, err := client.DisableLogging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the automatic copying of snapshots from one region to another region
// for a specified cluster.
//
// If your cluster and its snapshots are encrypted using an encrypted symmetric
// key from Key Management Service, use DeleteSnapshotCopyGrantto delete the grant that grants Amazon
// Redshift permission to the key in the destination region.
func redshift_DisableSnapshotCopy(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DisableSnapshotCopyInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}

	if resp, err := client.DisableSnapshotCopy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// From a datashare consumer account, remove association for the specified
// datashare.
func redshift_DisassociateDataShareConsumer(cfg aws.Config, client *redshift.Client) {
	input := &redshift.DisassociateDataShareConsumerInput{
		// DataShareArn: *string, // Required
	}

	if len(_redshiftDataShareArn) > 0 {
		input.DataShareArn = aws.String(_redshiftDataShareArn)
	}
	if len(_redshiftConsumerArn) > 0 {
		input.ConsumerArn = aws.String(_redshiftConsumerArn)
	}
	if len(_redshiftConsumerRegion) > 0 {
		input.ConsumerRegion = aws.String(_redshiftConsumerRegion)
	}
	if len(_redshiftDisassociateEntireAccount) > 0 {
		if err := assignInputField(input, "DisassociateEntireAccount", _redshiftDisassociateEntireAccount); err != nil {
			log.Errorf("invalid --disassociate-entire-account: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociateDataShareConsumer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts logging information, such as queries and connection attempts, for the
// specified Amazon Redshift cluster.
func redshift_EnableLogging(cfg aws.Config, client *redshift.Client) {
	input := &redshift.EnableLoggingInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftBucketName) > 0 {
		input.BucketName = aws.String(_redshiftBucketName)
	}
	if len(_redshiftLogDestinationType) > 0 {
		if err := assignInputField(input, "LogDestinationType", _redshiftLogDestinationType); err != nil {
			log.Errorf("invalid --log-destination-type: %s", err.Error())
			return
		}
	}
	if len(_redshiftLogExports) > 0 {
		input.LogExports = append([]string(nil), _redshiftLogExports...)
	}
	if len(_redshiftS3KeyPrefix) > 0 {
		input.S3KeyPrefix = aws.String(_redshiftS3KeyPrefix)
	}

	if resp, err := client.EnableLogging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the automatic copy of snapshots from one region to another region for a
// specified cluster.
func redshift_EnableSnapshotCopy(cfg aws.Config, client *redshift.Client) {
	input := &redshift.EnableSnapshotCopyInput{
		// ClusterIdentifier: *string, // Required
		// DestinationRegion: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftDestinationRegion) > 0 {
		input.DestinationRegion = aws.String(_redshiftDestinationRegion)
	}
	if len(_redshiftManualSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "ManualSnapshotRetentionPeriod", _redshiftManualSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --manual-snapshot-retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftRetentionPeriod) > 0 {
		if err := assignInputField(input, "RetentionPeriod", _redshiftRetentionPeriod); err != nil {
			log.Errorf("invalid --retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftSnapshotCopyGrantName) > 0 {
		input.SnapshotCopyGrantName = aws.String(_redshiftSnapshotCopyGrantName)
	}

	if resp, err := client.EnableSnapshotCopy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fails over the primary compute unit of the specified Multi-AZ cluster to
// another Availability Zone.
func redshift_FailoverPrimaryCompute(cfg aws.Config, client *redshift.Client) {
	input := &redshift.FailoverPrimaryComputeInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}

	if resp, err := client.FailoverPrimaryCompute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a database user name and temporary password with temporary
// authorization to log on to an Amazon Redshift database. The action returns the
// database user name prefixed with IAM: if AutoCreate is False or IAMA: if
// AutoCreate is True . You can optionally specify one or more database user groups
// that the user will join at log on. By default, the temporary credentials expire
// in 900 seconds. You can optionally specify a duration between 900 seconds (15
// minutes) and 3600 seconds (60 minutes). For more information, see [Using IAM Authentication to Generate Database User Credentials]in the Amazon
// Redshift Cluster Management Guide.
//
// The Identity and Access Management (IAM) user or role that runs
// GetClusterCredentials must have an IAM policy attached that allows access to all
// necessary actions and resources. For more information about permissions, see [Resource Policies for GetClusterCredentials]in
// the Amazon Redshift Cluster Management Guide.
//
// If the DbGroups parameter is specified, the IAM policy must allow the
// redshift:JoinGroup action with access to the listed dbgroups .
//
// In addition, if the AutoCreate parameter is set to True , then the policy must
// include the redshift:CreateClusterUser permission.
//
// If the DbName parameter is specified, the IAM policy must allow access to the
// resource dbname for the specified database name.
//
// [Using IAM Authentication to Generate Database User Credentials]: https://docs.aws.amazon.com/redshift/latest/mgmt/generating-user-credentials.html
// [Resource Policies for GetClusterCredentials]: https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html#redshift-policy-resources.getclustercredentials-resources
func redshift_GetClusterCredentials(cfg aws.Config, client *redshift.Client) {
	input := &redshift.GetClusterCredentialsInput{
		// DbUser: *string, // Required
	}

	if len(_redshiftDbUser) > 0 {
		input.DbUser = aws.String(_redshiftDbUser)
	}
	if len(_redshiftAutoCreate) > 0 {
		if err := assignInputField(input, "AutoCreate", _redshiftAutoCreate); err != nil {
			log.Errorf("invalid --auto-create: %s", err.Error())
			return
		}
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftCustomDomainName)
	}
	if len(_redshiftDbGroups) > 0 {
		input.DbGroups = append([]string(nil), _redshiftDbGroups...)
	}
	if len(_redshiftDBName) > 0 {
		input.DbName = aws.String(_redshiftDBName)
	}
	if len(_redshiftDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _redshiftDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetClusterCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a database user name and temporary password with temporary
// authorization to log in to an Amazon Redshift database. The database user is
// mapped 1:1 to the source Identity and Access Management (IAM) identity. For more
// information about IAM identities, see [IAM Identities (users, user groups, and roles)]in the Amazon Web Services Identity and
// Access Management User Guide.
//
// The Identity and Access Management (IAM) identity that runs this operation must
// have an IAM policy attached that allows access to all necessary actions and
// resources. For more information about permissions, see [Using identity-based policies (IAM policies)]in the Amazon Redshift
// Cluster Management Guide.
//
// [IAM Identities (users, user groups, and roles)]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id.html
// [Using identity-based policies (IAM policies)]: https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html
func redshift_GetClusterCredentialsWithIAM(cfg aws.Config, client *redshift.Client) {
	input := &redshift.GetClusterCredentialsWithIAMInput{}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftCustomDomainName)
	}
	if len(_redshiftDBName) > 0 {
		input.DbName = aws.String(_redshiftDBName)
	}
	if len(_redshiftDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _redshiftDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetClusterCredentialsWithIAM(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates an encrypted authentication token that propagates the caller's Amazon
// Web Services IAM Identity Center identity to Amazon Redshift clusters. This API
// extracts the Amazon Web Services IAM Identity Center identity from enhanced
// credentials and creates a secure token that Amazon Redshift drivers can use for
// authentication.
//
// The token is encrypted using Key Management Service (KMS) and can only be
// decrypted by the specified Amazon Redshift clusters. The token contains the
// caller's Amazon Web Services IAM Identity Center identity information and is
// valid for a limited time period.
//
// This API is exclusively for use with Amazon Web Services IAM Identity Center
// enhanced credentials. If the caller is not using enhanced credentials with
// embedded Amazon Web Services IAM Identity Center identity, the API will return
// an error.
func redshift_GetIdentityCenterAuthToken(cfg aws.Config, client *redshift.Client) {
	input := &redshift.GetIdentityCenterAuthTokenInput{
		// ClusterIds: []string, // Required
	}

	if len(_redshiftClusterIds) > 0 {
		input.ClusterIds = append([]string(nil), _redshiftClusterIds...)
	}

	if resp, err := client.GetIdentityCenterAuthToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the configuration options for the reserved-node exchange. These options
// include information about the source reserved node and target reserved node
// offering. Details include the node type, the price, the node count, and the
// offering type.
func redshift_GetReservedNodeExchangeConfigurationOptions(cfg aws.Config, client *redshift.Client) {
	input := &redshift.GetReservedNodeExchangeConfigurationOptionsInput{
		// ActionType: types.ReservedNodeExchangeActionType, // Required
	}

	if len(_redshiftActionType) > 0 {
		if err := assignInputField(input, "ActionType", _redshiftActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_redshiftSnapshotIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.GetReservedNodeExchangeConfigurationOptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.GetReservedNodeExchangeConfigurationOptionsOutput
	p := redshift.NewGetReservedNodeExchangeConfigurationOptionsPaginator(client, input)
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

// Returns an array of DC2 ReservedNodeOfferings that matches the payment type,
// term, and usage price of the given DC1 reserved node.
func redshift_GetReservedNodeExchangeOfferings(cfg aws.Config, client *redshift.Client) {
	input := &redshift.GetReservedNodeExchangeOfferingsInput{
		// ReservedNodeId: *string, // Required
	}

	if len(_redshiftReservedNodeId) > 0 {
		input.ReservedNodeId = aws.String(_redshiftReservedNodeId)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetReservedNodeExchangeOfferings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.GetReservedNodeExchangeOfferingsOutput
	p := redshift.NewGetReservedNodeExchangeOfferingsPaginator(client, input)
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

// Get the resource policy for a specified resource.
func redshift_GetResourcePolicy(cfg aws.Config, client *redshift.Client) {
	input := &redshift.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_redshiftResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the Amazon Redshift Advisor recommendations for one or multiple Amazon
// Redshift clusters in an Amazon Web Services account.
func redshift_ListRecommendations(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ListRecommendationsInput{}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftMarker) > 0 {
		input.Marker = aws.String(_redshiftMarker)
	}
	if len(_redshiftMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _redshiftMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_redshiftNamespaceArn) > 0 {
		input.NamespaceArn = aws.String(_redshiftNamespaceArn)
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshift.ListRecommendationsOutput
	p := redshift.NewListRecommendationsPaginator(client, input)
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

// This operation is retired. Calling this operation does not change AQUA
// configuration. Amazon Redshift automatically determines whether to use AQUA
// (Advanced Query Accelerator).
func redshift_ModifyAquaConfiguration(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyAquaConfigurationInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftAquaConfigurationStatus) > 0 {
		if err := assignInputField(input, "AquaConfigurationStatus", _redshiftAquaConfigurationStatus); err != nil {
			log.Errorf("invalid --aqua-configuration-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyAquaConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an authentication profile.
func redshift_ModifyAuthenticationProfile(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyAuthenticationProfileInput{
		// AuthenticationProfileContent: *string, // Required
		// AuthenticationProfileName: *string, // Required
	}

	if len(_redshiftAuthenticationProfileContent) > 0 {
		input.AuthenticationProfileContent = aws.String(_redshiftAuthenticationProfileContent)
	}
	if len(_redshiftAuthenticationProfileName) > 0 {
		input.AuthenticationProfileName = aws.String(_redshiftAuthenticationProfileName)
	}

	if resp, err := client.ModifyAuthenticationProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings for a cluster.
// You can also change node type and the number of nodes to scale up or down the
// cluster. When resizing a cluster, you must specify both the number of nodes and
// the node type even if one of the parameters does not change.
//
// You can add another security or parameter group, or change the admin user
// password. Resetting a cluster password or modifying the security groups
// associated with a cluster do not need a reboot. However, modifying a parameter
// group requires a reboot for parameters to take effect. For more information
// about managing clusters, go to [Amazon Redshift Clusters]in the Amazon Redshift Cluster Management Guide.
//
// VPC Block Public Access (BPA) enables you to block resources in VPCs and
// subnets that you own in a Region from reaching or being reached from the
// internet through internet gateways and egress-only internet gateways. If a
// subnet group for a provisioned cluster is in an account with VPC BPA turned on,
// the following capabilities are blocked:
//
// - Creating a public cluster
//
// - Restoring a public cluster
//
// - Modifying a private cluster to be public
//
// - Adding a subnet with VPC BPA turned on to the subnet group when there's at
// least one public cluster within the group
//
// For more information about VPC BPA, see [Block public access to VPCs and subnets] in the Amazon VPC User Guide.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
func redshift_ModifyCluster(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyClusterInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftAllowVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowVersionUpgrade", _redshiftAllowVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_redshiftAutomatedSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "AutomatedSnapshotRetentionPeriod", _redshiftAutomatedSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --automated-snapshot-retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_redshiftAvailabilityZone)
	}
	if len(_redshiftAvailabilityZoneRelocation) > 0 {
		if err := assignInputField(input, "AvailabilityZoneRelocation", _redshiftAvailabilityZoneRelocation); err != nil {
			log.Errorf("invalid --availability-zone-relocation: %s", err.Error())
			return
		}
	}
	if len(_redshiftClusterParameterGroupName) > 0 {
		input.ClusterParameterGroupName = aws.String(_redshiftClusterParameterGroupName)
	}
	if len(_redshiftClusterSecurityGroups) > 0 {
		input.ClusterSecurityGroups = append([]string(nil), _redshiftClusterSecurityGroups...)
	}
	if len(_redshiftClusterType) > 0 {
		input.ClusterType = aws.String(_redshiftClusterType)
	}
	if len(_redshiftClusterVersion) > 0 {
		input.ClusterVersion = aws.String(_redshiftClusterVersion)
	}
	if len(_redshiftElasticIp) > 0 {
		input.ElasticIp = aws.String(_redshiftElasticIp)
	}
	if len(_redshiftEncrypted) > 0 {
		if err := assignInputField(input, "Encrypted", _redshiftEncrypted); err != nil {
			log.Errorf("invalid --encrypted: %s", err.Error())
			return
		}
	}
	if len(_redshiftEnhancedVpcRouting) > 0 {
		if err := assignInputField(input, "EnhancedVpcRouting", _redshiftEnhancedVpcRouting); err != nil {
			log.Errorf("invalid --enhanced-vpc-routing: %s", err.Error())
			return
		}
	}
	if len(_redshiftExtraComputeForAutomaticOptimization) > 0 {
		if err := assignInputField(input, "ExtraComputeForAutomaticOptimization", _redshiftExtraComputeForAutomaticOptimization); err != nil {
			log.Errorf("invalid --extra-compute-for-automatic-optimization: %s", err.Error())
			return
		}
	}
	if len(_redshiftHsmClientCertificateIdentifier) > 0 {
		input.HsmClientCertificateIdentifier = aws.String(_redshiftHsmClientCertificateIdentifier)
	}
	if len(_redshiftHsmConfigurationIdentifier) > 0 {
		input.HsmConfigurationIdentifier = aws.String(_redshiftHsmConfigurationIdentifier)
	}
	if len(_redshiftIpAddressType) > 0 {
		input.IpAddressType = aws.String(_redshiftIpAddressType)
	}
	if len(_redshiftKMSKeyId) > 0 {
		input.KmsKeyId = aws.String(_redshiftKMSKeyId)
	}
	if len(_redshiftMaintenanceTrackName) > 0 {
		input.MaintenanceTrackName = aws.String(_redshiftMaintenanceTrackName)
	}
	if len(_redshiftManageMasterPassword) > 0 {
		if err := assignInputField(input, "ManageMasterPassword", _redshiftManageMasterPassword); err != nil {
			log.Errorf("invalid --manage-master-password: %s", err.Error())
			return
		}
	}
	if len(_redshiftManualSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "ManualSnapshotRetentionPeriod", _redshiftManualSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --manual-snapshot-retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftMasterPasswordSecretKmsKeyId) > 0 {
		input.MasterPasswordSecretKmsKeyId = aws.String(_redshiftMasterPasswordSecretKmsKeyId)
	}
	if len(_redshiftMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_redshiftMasterUserPassword)
	}
	if len(_redshiftMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _redshiftMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_redshiftNewClusterIdentifier) > 0 {
		input.NewClusterIdentifier = aws.String(_redshiftNewClusterIdentifier)
	}
	if len(_redshiftNodeType) > 0 {
		input.NodeType = aws.String(_redshiftNodeType)
	}
	if len(_redshiftNumberOfNodes) > 0 {
		if err := assignInputField(input, "NumberOfNodes", _redshiftNumberOfNodes); err != nil {
			log.Errorf("invalid --number-of-nodes: %s", err.Error())
			return
		}
	}
	if len(_redshiftPort) > 0 {
		if err := assignInputField(input, "Port", _redshiftPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_redshiftPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_redshiftPreferredMaintenanceWindow)
	}
	if len(_redshiftPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _redshiftPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_redshiftVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _redshiftVpcSecurityGroupIds...)
	}

	if resp, err := client.ModifyCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the database revision of a cluster. The database revision is a unique
// revision of the database running in a cluster.
func redshift_ModifyClusterDbRevision(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyClusterDbRevisionInput{
		// ClusterIdentifier: *string, // Required
		// RevisionTarget: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftRevisionTarget) > 0 {
		input.RevisionTarget = aws.String(_redshiftRevisionTarget)
	}

	if resp, err := client.ModifyClusterDbRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the list of Identity and Access Management (IAM) roles that can be
// used by the cluster to access other Amazon Web Services services.
//
// The maximum number of IAM roles that you can associate is subject to a quota.
// For more information, go to [Quotas and limits]in the Amazon Redshift Cluster Management Guide.
//
// [Quotas and limits]: https://docs.aws.amazon.com/redshift/latest/mgmt/amazon-redshift-limits.html
func redshift_ModifyClusterIamRoles(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyClusterIamRolesInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftAddIamRoles) > 0 {
		input.AddIamRoles = append([]string(nil), _redshiftAddIamRoles...)
	}
	if len(_redshiftDefaultIamRoleArn) > 0 {
		input.DefaultIamRoleArn = aws.String(_redshiftDefaultIamRoleArn)
	}
	if len(_redshiftRemoveIamRoles) > 0 {
		input.RemoveIamRoles = append([]string(nil), _redshiftRemoveIamRoles...)
	}

	if resp, err := client.ModifyClusterIamRoles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the maintenance settings of a cluster.
func redshift_ModifyClusterMaintenance(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyClusterMaintenanceInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftDeferMaintenance) > 0 {
		if err := assignInputField(input, "DeferMaintenance", _redshiftDeferMaintenance); err != nil {
			log.Errorf("invalid --defer-maintenance: %s", err.Error())
			return
		}
	}
	if len(_redshiftDeferMaintenanceDuration) > 0 {
		if err := assignInputField(input, "DeferMaintenanceDuration", _redshiftDeferMaintenanceDuration); err != nil {
			log.Errorf("invalid --defer-maintenance-duration: %s", err.Error())
			return
		}
	}
	if len(_redshiftDeferMaintenanceEndTime) > 0 {
		if err := assignInputField(input, "DeferMaintenanceEndTime", _redshiftDeferMaintenanceEndTime); err != nil {
			log.Errorf("invalid --defer-maintenance-end-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftDeferMaintenanceIdentifier) > 0 {
		input.DeferMaintenanceIdentifier = aws.String(_redshiftDeferMaintenanceIdentifier)
	}
	if len(_redshiftDeferMaintenanceStartTime) > 0 {
		if err := assignInputField(input, "DeferMaintenanceStartTime", _redshiftDeferMaintenanceStartTime); err != nil {
			log.Errorf("invalid --defer-maintenance-start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyClusterMaintenance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a parameter group. For the parameters parameter, it
// can't contain ASCII characters.
//
// For more information about parameters and parameter groups, go to [Amazon Redshift Parameter Groups] in the
// Amazon Redshift Cluster Management Guide.
//
// [Amazon Redshift Parameter Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html
func redshift_ModifyClusterParameterGroup(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyClusterParameterGroupInput{
		// ParameterGroupName: *string, // Required
		// Parameters: []types.Parameter, // Required
	}

	if len(_redshiftParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_redshiftParameterGroupName)
	}
	if len(_redshiftParameters) > 0 {
		if err := assignInputField(input, "Parameters", _redshiftParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings for a snapshot.
// This exanmple modifies the manual retention period setting for a cluster
// snapshot.
func redshift_ModifyClusterSnapshot(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyClusterSnapshotInput{
		// SnapshotIdentifier: *string, // Required
	}

	if len(_redshiftSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_redshiftSnapshotIdentifier)
	}
	if len(_redshiftForce) > 0 {
		if err := assignInputField(input, "Force", _redshiftForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_redshiftManualSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "ManualSnapshotRetentionPeriod", _redshiftManualSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --manual-snapshot-retention-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a snapshot schedule for a cluster.
func redshift_ModifyClusterSnapshotSchedule(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyClusterSnapshotScheduleInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftDisassociateSchedule) > 0 {
		if err := assignInputField(input, "DisassociateSchedule", _redshiftDisassociateSchedule); err != nil {
			log.Errorf("invalid --disassociate-schedule: %s", err.Error())
			return
		}
	}
	if len(_redshiftScheduleIdentifier) > 0 {
		input.ScheduleIdentifier = aws.String(_redshiftScheduleIdentifier)
	}

	if resp, err := client.ModifyClusterSnapshotSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a cluster subnet group to include the specified list of VPC subnets.
// The operation replaces the existing list of subnets with the new list of
// subnets.
//
// VPC Block Public Access (BPA) enables you to block resources in VPCs and
// subnets that you own in a Region from reaching or being reached from the
// internet through internet gateways and egress-only internet gateways. If a
// subnet group for a provisioned cluster is in an account with VPC BPA turned on,
// the following capabilities are blocked:
//
// - Creating a public cluster
//
// - Restoring a public cluster
//
// - Modifying a private cluster to be public
//
// - Adding a subnet with VPC BPA turned on to the subnet group when there's at
// least one public cluster within the group
//
// For more information about VPC BPA, see [Block public access to VPCs and subnets] in the Amazon VPC User Guide.
//
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
func redshift_ModifyClusterSubnetGroup(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyClusterSubnetGroupInput{
		// ClusterSubnetGroupName: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_redshiftClusterSubnetGroupName) > 0 {
		input.ClusterSubnetGroupName = aws.String(_redshiftClusterSubnetGroupName)
	}
	if len(_redshiftSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _redshiftSubnetIds...)
	}
	if len(_redshiftDescription) > 0 {
		input.Description = aws.String(_redshiftDescription)
	}

	if resp, err := client.ModifyClusterSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Contains information for changing a custom domain association.
func redshift_ModifyCustomDomainAssociation(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyCustomDomainAssociationInput{
		// ClusterIdentifier: *string, // Required
		// CustomDomainCertificateArn: *string, // Required
		// CustomDomainName: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftCustomDomainCertificateArn) > 0 {
		input.CustomDomainCertificateArn = aws.String(_redshiftCustomDomainCertificateArn)
	}
	if len(_redshiftCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftCustomDomainName)
	}

	if resp, err := client.ModifyCustomDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a Redshift-managed VPC endpoint.
func redshift_ModifyEndpointAccess(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyEndpointAccessInput{
		// EndpointName: *string, // Required
	}

	if len(_redshiftEndpointName) > 0 {
		input.EndpointName = aws.String(_redshiftEndpointName)
	}
	if len(_redshiftVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _redshiftVpcSecurityGroupIds...)
	}

	if resp, err := client.ModifyEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing Amazon Redshift event notification subscription.
func redshift_ModifyEventSubscription(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyEventSubscriptionInput{
		// SubscriptionName: *string, // Required
	}

	if len(_redshiftSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_redshiftSubscriptionName)
	}
	if len(_redshiftEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _redshiftEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_redshiftEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _redshiftEventCategories...)
	}
	if len(_redshiftSeverity) > 0 {
		input.Severity = aws.String(_redshiftSeverity)
	}
	if len(_redshiftSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_redshiftSnsTopicArn)
	}
	if len(_redshiftSourceIds) > 0 {
		input.SourceIds = append([]string(nil), _redshiftSourceIds...)
	}
	if len(_redshiftSourceType) > 0 {
		input.SourceType = aws.String(_redshiftSourceType)
	}

	if resp, err := client.ModifyEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a zero-ETL integration or S3 event integration with Amazon Redshift.
func redshift_ModifyIntegration(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyIntegrationInput{
		// IntegrationArn: *string, // Required
	}

	if len(_redshiftIntegrationArn) > 0 {
		input.IntegrationArn = aws.String(_redshiftIntegrationArn)
	}
	if len(_redshiftDescription) > 0 {
		input.Description = aws.String(_redshiftDescription)
	}
	if len(_redshiftIntegrationName) > 0 {
		input.IntegrationName = aws.String(_redshiftIntegrationName)
	}

	if resp, err := client.ModifyIntegration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the lakehouse configuration for a cluster. This operation allows you
// to manage Amazon Redshift federated permissions and Amazon Web Services IAM
// Identity Center trusted identity propagation.
func redshift_ModifyLakehouseConfiguration(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyLakehouseConfigurationInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftCatalogName) > 0 {
		input.CatalogName = aws.String(_redshiftCatalogName)
	}
	if len(_redshiftDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _redshiftDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_redshiftLakehouseIdcApplicationArn) > 0 {
		input.LakehouseIdcApplicationArn = aws.String(_redshiftLakehouseIdcApplicationArn)
	}
	if len(_redshiftLakehouseIdcRegistration) > 0 {
		if err := assignInputField(input, "LakehouseIdcRegistration", _redshiftLakehouseIdcRegistration); err != nil {
			log.Errorf("invalid --lakehouse-idc-registration: %s", err.Error())
			return
		}
	}
	if len(_redshiftLakehouseRegistration) > 0 {
		if err := assignInputField(input, "LakehouseRegistration", _redshiftLakehouseRegistration); err != nil {
			log.Errorf("invalid --lakehouse-registration: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyLakehouseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes an existing Amazon Redshift IAM Identity Center application.
func redshift_ModifyRedshiftIdcApplication(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyRedshiftIdcApplicationInput{
		// RedshiftIdcApplicationArn: *string, // Required
	}

	if len(_redshiftRedshiftIdcApplicationArn) > 0 {
		input.RedshiftIdcApplicationArn = aws.String(_redshiftRedshiftIdcApplicationArn)
	}
	if len(_redshiftAuthorizedTokenIssuerList) > 0 {
		if err := assignInputField(input, "AuthorizedTokenIssuerList", _redshiftAuthorizedTokenIssuerList); err != nil {
			log.Errorf("invalid --authorized-token-issuer-list: %s", err.Error())
			return
		}
	}
	if len(_redshiftIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_redshiftIamRoleArn)
	}
	if len(_redshiftIdcDisplayName) > 0 {
		input.IdcDisplayName = aws.String(_redshiftIdcDisplayName)
	}
	if len(_redshiftIdentityNamespace) > 0 {
		input.IdentityNamespace = aws.String(_redshiftIdentityNamespace)
	}
	if len(_redshiftServiceIntegrations) > 0 {
		if err := assignInputField(input, "ServiceIntegrations", _redshiftServiceIntegrations); err != nil {
			log.Errorf("invalid --service-integrations: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyRedshiftIdcApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a scheduled action.
func redshift_ModifyScheduledAction(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyScheduledActionInput{
		// ScheduledActionName: *string, // Required
	}

	if len(_redshiftScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_redshiftScheduledActionName)
	}
	if len(_redshiftEnable) > 0 {
		if err := assignInputField(input, "Enable", _redshiftEnable); err != nil {
			log.Errorf("invalid --enable: %s", err.Error())
			return
		}
	}
	if len(_redshiftEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _redshiftEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftIamRole) > 0 {
		input.IamRole = aws.String(_redshiftIamRole)
	}
	if len(_redshiftSchedule) > 0 {
		input.Schedule = aws.String(_redshiftSchedule)
	}
	if len(_redshiftScheduledActionDescription) > 0 {
		input.ScheduledActionDescription = aws.String(_redshiftScheduledActionDescription)
	}
	if len(_redshiftStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _redshiftStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftTargetAction) > 0 {
		if err := assignInputField(input, "TargetAction", _redshiftTargetAction); err != nil {
			log.Errorf("invalid --target-action: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the number of days to retain snapshots in the destination Amazon Web
// Services Region after they are copied from the source Amazon Web Services
// Region. By default, this operation only changes the retention period of copied
// automated snapshots. The retention periods for both new and existing copied
// automated snapshots are updated with the new retention period. You can set the
// manual option to change only the retention periods of copied manual snapshots.
// If you set this option, only newly copied manual snapshots have the new
// retention period.
func redshift_ModifySnapshotCopyRetentionPeriod(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifySnapshotCopyRetentionPeriodInput{
		// ClusterIdentifier: *string, // Required
		// RetentionPeriod: *int32, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftRetentionPeriod) > 0 {
		if err := assignInputField(input, "RetentionPeriod", _redshiftRetentionPeriod); err != nil {
			log.Errorf("invalid --retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftManual) > 0 {
		if err := assignInputField(input, "Manual", _redshiftManual); err != nil {
			log.Errorf("invalid --manual: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifySnapshotCopyRetentionPeriod(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a snapshot schedule. Any schedule associated with a cluster is
// modified asynchronously.
func redshift_ModifySnapshotSchedule(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifySnapshotScheduleInput{
		// ScheduleDefinitions: []string, // Required
		// ScheduleIdentifier: *string, // Required
	}

	if len(_redshiftScheduleDefinitions) > 0 {
		input.ScheduleDefinitions = append([]string(nil), _redshiftScheduleDefinitions...)
	}
	if len(_redshiftScheduleIdentifier) > 0 {
		input.ScheduleIdentifier = aws.String(_redshiftScheduleIdentifier)
	}

	if resp, err := client.ModifySnapshotSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a usage limit in a cluster. You can't modify the feature type or
// period of a usage limit.
func redshift_ModifyUsageLimit(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ModifyUsageLimitInput{
		// UsageLimitId: *string, // Required
	}

	if len(_redshiftUsageLimitId) > 0 {
		input.UsageLimitId = aws.String(_redshiftUsageLimitId)
	}
	if len(_redshiftAmount) > 0 {
		if err := assignInputField(input, "Amount", _redshiftAmount); err != nil {
			log.Errorf("invalid --amount: %s", err.Error())
			return
		}
	}
	if len(_redshiftBreachAction) > 0 {
		if err := assignInputField(input, "BreachAction", _redshiftBreachAction); err != nil {
			log.Errorf("invalid --breach-action: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyUsageLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Pauses a cluster.
func redshift_PauseCluster(cfg aws.Config, client *redshift.Client) {
	input := &redshift.PauseClusterInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}

	if resp, err := client.PauseCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to purchase reserved nodes. Amazon Redshift offers a predefined set
// of reserved node offerings. You can purchase one or more of the offerings. You
// can call the DescribeReservedNodeOfferingsAPI to obtain the available reserved node offerings. You can call
// this API by providing a specific reserved node offering and the number of nodes
// you want to reserve.
//
// For more information about reserved node offerings, go to [Purchasing Reserved Nodes] in the Amazon
// Redshift Cluster Management Guide.
//
// [Purchasing Reserved Nodes]: https://docs.aws.amazon.com/redshift/latest/mgmt/purchase-reserved-node-instance.html
func redshift_PurchaseReservedNodeOffering(cfg aws.Config, client *redshift.Client) {
	input := &redshift.PurchaseReservedNodeOfferingInput{
		// ReservedNodeOfferingId: *string, // Required
	}

	if len(_redshiftReservedNodeOfferingId) > 0 {
		input.ReservedNodeOfferingId = aws.String(_redshiftReservedNodeOfferingId)
	}
	if len(_redshiftNodeCount) > 0 {
		if err := assignInputField(input, "NodeCount", _redshiftNodeCount); err != nil {
			log.Errorf("invalid --node-count: %s", err.Error())
			return
		}
	}

	if resp, err := client.PurchaseReservedNodeOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the resource policy for a specified resource.
func redshift_PutResourcePolicy(cfg aws.Config, client *redshift.Client) {
	input := &redshift.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_redshiftPolicy) > 0 {
		input.Policy = aws.String(_redshiftPolicy)
	}
	if len(_redshiftResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots a cluster. This action is taken as soon as possible. It results in a
// momentary outage to the cluster, during which the cluster status is set to
// rebooting . A cluster event is created when the reboot is completed. Any pending
// cluster modifications (see ModifyCluster) are applied at this reboot. For more information
// about managing clusters, go to [Amazon Redshift Clusters]in the Amazon Redshift Cluster Management Guide.
//
// [Amazon Redshift Clusters]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html
func redshift_RebootCluster(cfg aws.Config, client *redshift.Client) {
	input := &redshift.RebootClusterInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}

	if resp, err := client.RebootCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a cluster or serverless namespace to the Amazon Web Services Glue
// Data Catalog.
func redshift_RegisterNamespace(cfg aws.Config, client *redshift.Client) {
	input := &redshift.RegisterNamespaceInput{
		// ConsumerIdentifiers: []string, // Required
		// NamespaceIdentifier: types.NamespaceIdentifierUnion, // Required
	}

	if len(_redshiftConsumerIdentifiers) > 0 {
		input.ConsumerIdentifiers = append([]string(nil), _redshiftConsumerIdentifiers...)
	}
	if len(_redshiftNamespaceIdentifier) > 0 {
		if err := assignInputField(input, "NamespaceIdentifier", _redshiftNamespaceIdentifier); err != nil {
			log.Errorf("invalid --namespace-identifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// From a datashare consumer account, rejects the specified datashare.
func redshift_RejectDataShare(cfg aws.Config, client *redshift.Client) {
	input := &redshift.RejectDataShareInput{
		// DataShareArn: *string, // Required
	}

	if len(_redshiftDataShareArn) > 0 {
		input.DataShareArn = aws.String(_redshiftDataShareArn)
	}

	if resp, err := client.RejectDataShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets one or more parameters of the specified parameter group to their default
// values and sets the source values of the parameters to "engine-default". To
// reset the entire parameter group specify the ResetAllParameters parameter. For
// parameter changes to take effect you must reboot any associated clusters.
func redshift_ResetClusterParameterGroup(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ResetClusterParameterGroupInput{
		// ParameterGroupName: *string, // Required
	}

	if len(_redshiftParameterGroupName) > 0 {
		input.ParameterGroupName = aws.String(_redshiftParameterGroupName)
	}
	if len(_redshiftParameters) > 0 {
		if err := assignInputField(input, "Parameters", _redshiftParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_redshiftResetAllParameters) > 0 {
		if err := assignInputField(input, "ResetAllParameters", _redshiftResetAllParameters); err != nil {
			log.Errorf("invalid --reset-all-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResetClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the size of the cluster. You can change the cluster's type, or change
// the number or type of nodes. The default behavior is to use the elastic resize
// method. With an elastic resize, your cluster is available for read and write
// operations more quickly than with the classic resize method.
//
// Elastic resize operations have the following restrictions:
//
// - You can only resize clusters of the following types:
//
// - dc2.large
//
// - dc2.8xlarge
//
// - ra3.large
//
// - ra3.xlplus
//
// - ra3.4xlarge
//
// - ra3.16xlarge
//
// - The type of nodes that you add must match the node type for the cluster.
func redshift_ResizeCluster(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ResizeClusterInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftClassic) > 0 {
		if err := assignInputField(input, "Classic", _redshiftClassic); err != nil {
			log.Errorf("invalid --classic: %s", err.Error())
			return
		}
	}
	if len(_redshiftClusterType) > 0 {
		input.ClusterType = aws.String(_redshiftClusterType)
	}
	if len(_redshiftNodeType) > 0 {
		input.NodeType = aws.String(_redshiftNodeType)
	}
	if len(_redshiftNumberOfNodes) > 0 {
		if err := assignInputField(input, "NumberOfNodes", _redshiftNumberOfNodes); err != nil {
			log.Errorf("invalid --number-of-nodes: %s", err.Error())
			return
		}
	}
	if len(_redshiftReservedNodeId) > 0 {
		input.ReservedNodeId = aws.String(_redshiftReservedNodeId)
	}
	if len(_redshiftTargetReservedNodeOfferingId) > 0 {
		input.TargetReservedNodeOfferingId = aws.String(_redshiftTargetReservedNodeOfferingId)
	}

	if resp, err := client.ResizeCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new cluster from a snapshot. By default, Amazon Redshift creates the
// resulting cluster with the same configuration as the original cluster from which
// the snapshot was created, except that the new cluster is created with the
// default cluster security and parameter groups. After Amazon Redshift creates the
// cluster, you can use the ModifyClusterAPI to associate a different security group and
// different parameter group with the restored cluster. If you are using a DS node
// type, you can also choose to change to another DS node type of the same size
// during restore.
//
// If you restore a cluster into a VPC, you must provide a cluster subnet group
// where you want the cluster restored.
//
// VPC Block Public Access (BPA) enables you to block resources in VPCs and
// subnets that you own in a Region from reaching or being reached from the
// internet through internet gateways and egress-only internet gateways. If a
// subnet group for a provisioned cluster is in an account with VPC BPA turned on,
// the following capabilities are blocked:
//
// - Creating a public cluster
//
// - Restoring a public cluster
//
// - Modifying a private cluster to be public
//
// - Adding a subnet with VPC BPA turned on to the subnet group when there's at
// least one public cluster within the group
//
// For more information about VPC BPA, see [Block public access to VPCs and subnets] in the Amazon VPC User Guide.
//
// For more information about working with snapshots, go to [Amazon Redshift Snapshots] in the Amazon
// Redshift Cluster Management Guide.
//
// [Amazon Redshift Snapshots]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
func redshift_RestoreFromClusterSnapshot(cfg aws.Config, client *redshift.Client) {
	input := &redshift.RestoreFromClusterSnapshotInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftAdditionalInfo) > 0 {
		input.AdditionalInfo = aws.String(_redshiftAdditionalInfo)
	}
	if len(_redshiftAllowVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowVersionUpgrade", _redshiftAllowVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_redshiftAquaConfigurationStatus) > 0 {
		if err := assignInputField(input, "AquaConfigurationStatus", _redshiftAquaConfigurationStatus); err != nil {
			log.Errorf("invalid --aqua-configuration-status: %s", err.Error())
			return
		}
	}
	if len(_redshiftAutomatedSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "AutomatedSnapshotRetentionPeriod", _redshiftAutomatedSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --automated-snapshot-retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_redshiftAvailabilityZone)
	}
	if len(_redshiftAvailabilityZoneRelocation) > 0 {
		if err := assignInputField(input, "AvailabilityZoneRelocation", _redshiftAvailabilityZoneRelocation); err != nil {
			log.Errorf("invalid --availability-zone-relocation: %s", err.Error())
			return
		}
	}
	if len(_redshiftCatalogName) > 0 {
		input.CatalogName = aws.String(_redshiftCatalogName)
	}
	if len(_redshiftClusterParameterGroupName) > 0 {
		input.ClusterParameterGroupName = aws.String(_redshiftClusterParameterGroupName)
	}
	if len(_redshiftClusterSecurityGroups) > 0 {
		input.ClusterSecurityGroups = append([]string(nil), _redshiftClusterSecurityGroups...)
	}
	if len(_redshiftClusterSubnetGroupName) > 0 {
		input.ClusterSubnetGroupName = aws.String(_redshiftClusterSubnetGroupName)
	}
	if len(_redshiftDefaultIamRoleArn) > 0 {
		input.DefaultIamRoleArn = aws.String(_redshiftDefaultIamRoleArn)
	}
	if len(_redshiftElasticIp) > 0 {
		input.ElasticIp = aws.String(_redshiftElasticIp)
	}
	if len(_redshiftEncrypted) > 0 {
		if err := assignInputField(input, "Encrypted", _redshiftEncrypted); err != nil {
			log.Errorf("invalid --encrypted: %s", err.Error())
			return
		}
	}
	if len(_redshiftEnhancedVpcRouting) > 0 {
		if err := assignInputField(input, "EnhancedVpcRouting", _redshiftEnhancedVpcRouting); err != nil {
			log.Errorf("invalid --enhanced-vpc-routing: %s", err.Error())
			return
		}
	}
	if len(_redshiftHsmClientCertificateIdentifier) > 0 {
		input.HsmClientCertificateIdentifier = aws.String(_redshiftHsmClientCertificateIdentifier)
	}
	if len(_redshiftHsmConfigurationIdentifier) > 0 {
		input.HsmConfigurationIdentifier = aws.String(_redshiftHsmConfigurationIdentifier)
	}
	if len(_redshiftIamRoles) > 0 {
		input.IamRoles = append([]string(nil), _redshiftIamRoles...)
	}
	if len(_redshiftIpAddressType) > 0 {
		input.IpAddressType = aws.String(_redshiftIpAddressType)
	}
	if len(_redshiftKMSKeyId) > 0 {
		input.KmsKeyId = aws.String(_redshiftKMSKeyId)
	}
	if len(_redshiftMaintenanceTrackName) > 0 {
		input.MaintenanceTrackName = aws.String(_redshiftMaintenanceTrackName)
	}
	if len(_redshiftManageMasterPassword) > 0 {
		if err := assignInputField(input, "ManageMasterPassword", _redshiftManageMasterPassword); err != nil {
			log.Errorf("invalid --manage-master-password: %s", err.Error())
			return
		}
	}
	if len(_redshiftManualSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "ManualSnapshotRetentionPeriod", _redshiftManualSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --manual-snapshot-retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftMasterPasswordSecretKmsKeyId) > 0 {
		input.MasterPasswordSecretKmsKeyId = aws.String(_redshiftMasterPasswordSecretKmsKeyId)
	}
	if len(_redshiftMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _redshiftMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_redshiftNodeType) > 0 {
		input.NodeType = aws.String(_redshiftNodeType)
	}
	if len(_redshiftNumberOfNodes) > 0 {
		if err := assignInputField(input, "NumberOfNodes", _redshiftNumberOfNodes); err != nil {
			log.Errorf("invalid --number-of-nodes: %s", err.Error())
			return
		}
	}
	if len(_redshiftOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_redshiftOwnerAccount)
	}
	if len(_redshiftPort) > 0 {
		if err := assignInputField(input, "Port", _redshiftPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_redshiftPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_redshiftPreferredMaintenanceWindow)
	}
	if len(_redshiftPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _redshiftPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_redshiftRedshiftIdcApplicationArn) > 0 {
		input.RedshiftIdcApplicationArn = aws.String(_redshiftRedshiftIdcApplicationArn)
	}
	if len(_redshiftReservedNodeId) > 0 {
		input.ReservedNodeId = aws.String(_redshiftReservedNodeId)
	}
	if len(_redshiftSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_redshiftSnapshotArn)
	}
	if len(_redshiftSnapshotClusterIdentifier) > 0 {
		input.SnapshotClusterIdentifier = aws.String(_redshiftSnapshotClusterIdentifier)
	}
	if len(_redshiftSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_redshiftSnapshotIdentifier)
	}
	if len(_redshiftSnapshotScheduleIdentifier) > 0 {
		input.SnapshotScheduleIdentifier = aws.String(_redshiftSnapshotScheduleIdentifier)
	}
	if len(_redshiftTargetReservedNodeOfferingId) > 0 {
		input.TargetReservedNodeOfferingId = aws.String(_redshiftTargetReservedNodeOfferingId)
	}
	if len(_redshiftVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _redshiftVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreFromClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new table from a table in an Amazon Redshift cluster snapshot. You
// must create the new table within the Amazon Redshift cluster that the snapshot
// was taken from.
//
// You cannot use RestoreTableFromClusterSnapshot to restore a table with the same
// name as an existing table in an Amazon Redshift cluster. That is, you cannot
// overwrite an existing table in a cluster with a restored table. If you want to
// replace your original table with a new, restored table, then rename or drop your
// original table before you call RestoreTableFromClusterSnapshot . When you have
// renamed your original table, then you can pass the original name of the table as
// the NewTableName parameter value in the call to RestoreTableFromClusterSnapshot
// . This way, you can replace the original table with the table created from the
// snapshot.
//
// You can't use this operation to restore tables with [interleaved sort keys].
//
// [interleaved sort keys]: https://docs.aws.amazon.com/redshift/latest/dg/t_Sorting_data.html#t_Sorting_data-interleaved
func redshift_RestoreTableFromClusterSnapshot(cfg aws.Config, client *redshift.Client) {
	input := &redshift.RestoreTableFromClusterSnapshotInput{
		// ClusterIdentifier: *string, // Required
		// NewTableName: *string, // Required
		// SnapshotIdentifier: *string, // Required
		// SourceDatabaseName: *string, // Required
		// SourceTableName: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftNewTableName) > 0 {
		input.NewTableName = aws.String(_redshiftNewTableName)
	}
	if len(_redshiftSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_redshiftSnapshotIdentifier)
	}
	if len(_redshiftSourceDatabaseName) > 0 {
		input.SourceDatabaseName = aws.String(_redshiftSourceDatabaseName)
	}
	if len(_redshiftSourceTableName) > 0 {
		input.SourceTableName = aws.String(_redshiftSourceTableName)
	}
	if len(_redshiftEnableCaseSensitiveIdentifier) > 0 {
		if err := assignInputField(input, "EnableCaseSensitiveIdentifier", _redshiftEnableCaseSensitiveIdentifier); err != nil {
			log.Errorf("invalid --enable-case-sensitive-identifier: %s", err.Error())
			return
		}
	}
	if len(_redshiftSourceSchemaName) > 0 {
		input.SourceSchemaName = aws.String(_redshiftSourceSchemaName)
	}
	if len(_redshiftTargetDatabaseName) > 0 {
		input.TargetDatabaseName = aws.String(_redshiftTargetDatabaseName)
	}
	if len(_redshiftTargetSchemaName) > 0 {
		input.TargetSchemaName = aws.String(_redshiftTargetSchemaName)
	}

	if resp, err := client.RestoreTableFromClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resumes a paused cluster.
func redshift_ResumeCluster(cfg aws.Config, client *redshift.Client) {
	input := &redshift.ResumeClusterInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}

	if resp, err := client.ResumeCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes an ingress rule in an Amazon Redshift security group for a previously
// authorized IP range or Amazon EC2 security group. To add an ingress rule, see AuthorizeClusterSecurityGroupIngress.
// For information about managing security groups, go to [Amazon Redshift Cluster Security Groups]in the Amazon Redshift
// Cluster Management Guide.
//
// [Amazon Redshift Cluster Security Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html
func redshift_RevokeClusterSecurityGroupIngress(cfg aws.Config, client *redshift.Client) {
	input := &redshift.RevokeClusterSecurityGroupIngressInput{
		// ClusterSecurityGroupName: *string, // Required
	}

	if len(_redshiftClusterSecurityGroupName) > 0 {
		input.ClusterSecurityGroupName = aws.String(_redshiftClusterSecurityGroupName)
	}
	if len(_redshiftCIDRIP) > 0 {
		input.CIDRIP = aws.String(_redshiftCIDRIP)
	}
	if len(_redshiftEC2SecurityGroupName) > 0 {
		input.EC2SecurityGroupName = aws.String(_redshiftEC2SecurityGroupName)
	}
	if len(_redshiftEC2SecurityGroupOwnerId) > 0 {
		input.EC2SecurityGroupOwnerId = aws.String(_redshiftEC2SecurityGroupOwnerId)
	}

	if resp, err := client.RevokeClusterSecurityGroupIngress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes access to a cluster.
func redshift_RevokeEndpointAccess(cfg aws.Config, client *redshift.Client) {
	input := &redshift.RevokeEndpointAccessInput{}

	if len(_redshiftAccount) > 0 {
		input.Account = aws.String(_redshiftAccount)
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftForce) > 0 {
		if err := assignInputField(input, "Force", _redshiftForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_redshiftVpcIds) > 0 {
		input.VpcIds = append([]string(nil), _redshiftVpcIds...)
	}

	if resp, err := client.RevokeEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the ability of the specified Amazon Web Services account to restore the
// specified snapshot. If the account is currently restoring the snapshot, the
// restore will run to completion.
//
// For more information about working with snapshots, go to [Amazon Redshift Snapshots] in the Amazon
// Redshift Cluster Management Guide.
//
// [Amazon Redshift Snapshots]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html
func redshift_RevokeSnapshotAccess(cfg aws.Config, client *redshift.Client) {
	input := &redshift.RevokeSnapshotAccessInput{
		// AccountWithRestoreAccess: *string, // Required
	}

	if len(_redshiftAccountWithRestoreAccess) > 0 {
		input.AccountWithRestoreAccess = aws.String(_redshiftAccountWithRestoreAccess)
	}
	if len(_redshiftSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_redshiftSnapshotArn)
	}
	if len(_redshiftSnapshotClusterIdentifier) > 0 {
		input.SnapshotClusterIdentifier = aws.String(_redshiftSnapshotClusterIdentifier)
	}
	if len(_redshiftSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_redshiftSnapshotIdentifier)
	}

	if resp, err := client.RevokeSnapshotAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rotates the encryption keys for a cluster.
func redshift_RotateEncryptionKey(cfg aws.Config, client *redshift.Client) {
	input := &redshift.RotateEncryptionKeyInput{
		// ClusterIdentifier: *string, // Required
	}

	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}

	if resp, err := client.RotateEncryptionKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of a partner integration.
func redshift_UpdatePartnerStatus(cfg aws.Config, client *redshift.Client) {
	input := &redshift.UpdatePartnerStatusInput{
		// AccountId: *string, // Required
		// ClusterIdentifier: *string, // Required
		// DatabaseName: *string, // Required
		// PartnerName: *string, // Required
		// Status: types.PartnerIntegrationStatus, // Required
	}

	if len(_redshiftAccountId) > 0 {
		input.AccountId = aws.String(_redshiftAccountId)
	}
	if len(_redshiftClusterIdentifier) > 0 {
		input.ClusterIdentifier = aws.String(_redshiftClusterIdentifier)
	}
	if len(_redshiftDatabaseName) > 0 {
		input.DatabaseName = aws.String(_redshiftDatabaseName)
	}
	if len(_redshiftPartnerName) > 0 {
		input.PartnerName = aws.String(_redshiftPartnerName)
	}
	if len(_redshiftStatus) > 0 {
		if err := assignInputField(input, "Status", _redshiftStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_redshiftStatusMessage) > 0 {
		input.StatusMessage = aws.String(_redshiftStatusMessage)
	}

	if resp, err := client.UpdatePartnerStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_redshiftCmd)
	_redshiftCmd.Flags().SortFlags = false

	_redshiftCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_redshiftCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_redshiftCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_redshiftCmd.Flags().StringVarP(&_redshiftAccount, "account", "", "", "Account")
	_redshiftCmd.Flags().StringVarP(&_redshiftAccountId, "account-id", "", "", "Account ID")
	_redshiftCmd.Flags().StringVarP(&_redshiftAccountWithRestoreAccess, "account-with-restore-access", "", "", "Account With Restore Access")
	_redshiftCmd.Flags().StringVarP(&_redshiftActionType, "action-type", "", "", "Action Type")
	_redshiftCmd.Flags().StringVarP(&_redshiftActive, "active", "", "", "Active")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftAddIamRoles, "add-iam-roles", "", nil, "Add IAM Roles")
	_redshiftCmd.Flags().StringVarP(&_redshiftAdditionalEncryptionContext, "additional-encryption-context", "", "", "Additional Encryption Context")
	_redshiftCmd.Flags().StringVarP(&_redshiftAdditionalInfo, "additional-info", "", "", "Additional Info")
	_redshiftCmd.Flags().StringVarP(&_redshiftAllowVersionUpgrade, "allow-version-upgrade", "", "", "Allow Version Upgrade")
	_redshiftCmd.Flags().StringVarP(&_redshiftAllowWrites, "allow-writes", "", "", "Allow Writes")
	_redshiftCmd.Flags().StringVarP(&_redshiftAmount, "amount", "", "", "Amount")
	_redshiftCmd.Flags().StringVarP(&_redshiftApplicationType, "application-type", "", "", "Application Type")
	_redshiftCmd.Flags().StringVarP(&_redshiftAquaConfigurationStatus, "aqua-configuration-status", "", "", "Aqua Configuration Status")
	_redshiftCmd.Flags().StringVarP(&_redshiftAssociateEntireAccount, "associate-entire-account", "", "", "Associate Entire Account")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftAttributeNames, "attribute-names", "", nil, "Attribute Names")
	_redshiftCmd.Flags().StringVarP(&_redshiftAuthenticationProfileContent, "authentication-profile-content", "", "", "Authentication Profile Content")
	_redshiftCmd.Flags().StringVarP(&_redshiftAuthenticationProfileName, "authentication-profile-name", "", "", "Authentication Profile Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftAuthorizedTokenIssuerList, "authorized-token-issuer-list", "", "", "Authorized Token Issuer List")
	_redshiftCmd.Flags().StringVarP(&_redshiftAutoCreate, "auto-create", "", "", "Auto Create")
	_redshiftCmd.Flags().StringVarP(&_redshiftAutomatedSnapshotRetentionPeriod, "automated-snapshot-retention-period", "", "", "Automated Snapshot Retention Period")
	_redshiftCmd.Flags().StringVarP(&_redshiftAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_redshiftCmd.Flags().StringVarP(&_redshiftAvailabilityZoneRelocation, "availability-zone-relocation", "", "", "Availability Zone Relocation")
	_redshiftCmd.Flags().StringVarP(&_redshiftBreachAction, "breach-action", "", "", "Breach Action")
	_redshiftCmd.Flags().StringVarP(&_redshiftBucketName, "bucket-name", "", "", "Bucket Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftCatalogName, "catalog-name", "", "", "Catalog Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftCIDRIP, "cidrip", "", "", "Cidrip")
	_redshiftCmd.Flags().StringVarP(&_redshiftClassic, "classic", "", "", "Classic")
	_redshiftCmd.Flags().StringVarP(&_redshiftClusterExists, "cluster-exists", "", "", "Cluster Exists")
	_redshiftCmd.Flags().StringVarP(&_redshiftClusterIdentifier, "cluster-identifier", "", "", "Cluster Identifier")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftClusterIds, "cluster-ids", "", nil, "Cluster Ids")
	_redshiftCmd.Flags().StringVarP(&_redshiftClusterParameterGroupFamily, "cluster-parameter-group-family", "", "", "Cluster Parameter Group Family")
	_redshiftCmd.Flags().StringVarP(&_redshiftClusterParameterGroupName, "cluster-parameter-group-name", "", "", "Cluster Parameter Group Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftClusterSecurityGroupName, "cluster-security-group-name", "", "", "Cluster Security Group Name")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftClusterSecurityGroups, "cluster-security-groups", "", nil, "Cluster Security Groups")
	_redshiftCmd.Flags().StringVarP(&_redshiftClusterSubnetGroupName, "cluster-subnet-group-name", "", "", "Cluster Subnet Group Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftClusterType, "cluster-type", "", "", "Cluster Type")
	_redshiftCmd.Flags().StringVarP(&_redshiftClusterVersion, "cluster-version", "", "", "Cluster Version")
	_redshiftCmd.Flags().StringVarP(&_redshiftConsumerArn, "consumer-arn", "", "", "Consumer ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftConsumerIdentifier, "consumer-identifier", "", "", "Consumer Identifier")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftConsumerIdentifiers, "consumer-identifiers", "", nil, "Consumer Identifiers")
	_redshiftCmd.Flags().StringVarP(&_redshiftConsumerRegion, "consumer-region", "", "", "Consumer Region")
	_redshiftCmd.Flags().StringVarP(&_redshiftCustomDomainCertificateArn, "custom-domain-certificate-arn", "", "", "Custom Domain Certificate ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftCustomDomainName, "custom-domain-name", "", "", "Custom Domain Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftDataShareArn, "data-share-arn", "", "", "Data Share ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftDatabaseName, "database-name", "", "", "Database Name")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftDbGroups, "db-groups", "", nil, "DB Groups")
	_redshiftCmd.Flags().StringVarP(&_redshiftDBName, "db-name", "", "", "DB Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftDbUser, "db-user", "", "", "DB User")
	_redshiftCmd.Flags().StringVarP(&_redshiftDefaultIamRoleArn, "default-iam-role-arn", "", "", "Default IAM Role ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftDeferMaintenance, "defer-maintenance", "", "", "Defer Maintenance")
	_redshiftCmd.Flags().StringVarP(&_redshiftDeferMaintenanceDuration, "defer-maintenance-duration", "", "", "Defer Maintenance Duration")
	_redshiftCmd.Flags().StringVarP(&_redshiftDeferMaintenanceEndTime, "defer-maintenance-end-time", "", "", "Defer Maintenance End Time")
	_redshiftCmd.Flags().StringVarP(&_redshiftDeferMaintenanceIdentifier, "defer-maintenance-identifier", "", "", "Defer Maintenance Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftDeferMaintenanceStartTime, "defer-maintenance-start-time", "", "", "Defer Maintenance Start Time")
	_redshiftCmd.Flags().StringVarP(&_redshiftDescription, "description", "", "", "Description")
	_redshiftCmd.Flags().StringVarP(&_redshiftDestinationRegion, "destination-region", "", "", "Destination Region")
	_redshiftCmd.Flags().StringVarP(&_redshiftDisassociateEntireAccount, "disassociate-entire-account", "", "", "Disassociate Entire Account")
	_redshiftCmd.Flags().StringVarP(&_redshiftDisassociateSchedule, "disassociate-schedule", "", "", "Disassociate Schedule")
	_redshiftCmd.Flags().StringVarP(&_redshiftDryRun, "dry-run", "", "", "Dry Run")
	_redshiftCmd.Flags().StringVarP(&_redshiftDuration, "duration", "", "", "Duration")
	_redshiftCmd.Flags().StringVarP(&_redshiftDurationSeconds, "duration-seconds", "", "", "Duration Seconds")
	_redshiftCmd.Flags().StringVarP(&_redshiftEC2SecurityGroupName, "ec2-security-group-name", "", "", "EC2 Security Group Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftEC2SecurityGroupOwnerId, "ec2-security-group-owner-id", "", "", "EC2 Security Group Owner ID")
	_redshiftCmd.Flags().StringVarP(&_redshiftElasticIp, "elastic-ip", "", "", "Elastic IP")
	_redshiftCmd.Flags().StringVarP(&_redshiftEnable, "enable", "", "", "Enable")
	_redshiftCmd.Flags().StringVarP(&_redshiftEnableCaseSensitiveIdentifier, "enable-case-sensitive-identifier", "", "", "Enable Case Sensitive Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftEnabled, "enabled", "", "", "Enabled")
	_redshiftCmd.Flags().StringVarP(&_redshiftEncrypted, "encrypted", "", "", "Encrypted")
	_redshiftCmd.Flags().StringVarP(&_redshiftEndTime, "end-time", "", "", "End Time")
	_redshiftCmd.Flags().StringVarP(&_redshiftEndpointName, "endpoint-name", "", "", "Endpoint Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftEnhancedVpcRouting, "enhanced-vpc-routing", "", "", "Enhanced VPC Routing")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftEventCategories, "event-categories", "", nil, "Event Categories")
	_redshiftCmd.Flags().StringVarP(&_redshiftExtraComputeForAutomaticOptimization, "extra-compute-for-automatic-optimization", "", "", "Extra Compute For Automatic Optimization")
	_redshiftCmd.Flags().StringVarP(&_redshiftFeatureType, "feature-type", "", "", "Feature Type")
	_redshiftCmd.Flags().StringVarP(&_redshiftFilters, "filters", "", "", "Filters")
	_redshiftCmd.Flags().StringVarP(&_redshiftFinalClusterSnapshotIdentifier, "final-cluster-snapshot-identifier", "", "", "Final Cluster Snapshot Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftFinalClusterSnapshotRetentionPeriod, "final-cluster-snapshot-retention-period", "", "", "Final Cluster Snapshot Retention Period")
	_redshiftCmd.Flags().StringVarP(&_redshiftForce, "force", "", "", "Force")
	_redshiftCmd.Flags().StringVarP(&_redshiftGrantee, "grantee", "", "", "Grantee")
	_redshiftCmd.Flags().StringVarP(&_redshiftHsmClientCertificateIdentifier, "hsm-client-certificate-identifier", "", "", "Hsm Client Certificate Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftHsmConfigurationIdentifier, "hsm-configuration-identifier", "", "", "Hsm Configuration Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftHsmIpAddress, "hsm-ip-address", "", "", "Hsm IP Address")
	_redshiftCmd.Flags().StringVarP(&_redshiftHsmPartitionName, "hsm-partition-name", "", "", "Hsm Partition Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftHsmPartitionPassword, "hsm-partition-password", "", "", "Hsm Partition Password")
	_redshiftCmd.Flags().StringVarP(&_redshiftHsmServerPublicCertificate, "hsm-server-public-certificate", "", "", "Hsm Server Public Certificate")
	_redshiftCmd.Flags().StringVarP(&_redshiftIamRole, "iam-role", "", "", "IAM Role")
	_redshiftCmd.Flags().StringVarP(&_redshiftIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftIamRoles, "iam-roles", "", nil, "IAM Roles")
	_redshiftCmd.Flags().StringVarP(&_redshiftIdcDisplayName, "idc-display-name", "", "", "Idc Display Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftIdcInstanceArn, "idc-instance-arn", "", "", "Idc Instance ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftIdentifiers, "identifiers", "", "", "Identifiers")
	_redshiftCmd.Flags().StringVarP(&_redshiftIdentityNamespace, "identity-namespace", "", "", "Identity Namespace")
	_redshiftCmd.Flags().StringVarP(&_redshiftIntegrationArn, "integration-arn", "", "", "Integration ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftIntegrationName, "integration-name", "", "", "Integration Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftIpAddressType, "ip-address-type", "", "", "IP Address Type")
	_redshiftCmd.Flags().StringVarP(&_redshiftKMSKeyId, "kms-key-id", "", "", "KMS Key ID")
	_redshiftCmd.Flags().StringVarP(&_redshiftLakehouseIdcApplicationArn, "lakehouse-idc-application-arn", "", "", "Lakehouse Idc Application ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftLakehouseIdcRegistration, "lakehouse-idc-registration", "", "", "Lakehouse Idc Registration")
	_redshiftCmd.Flags().StringVarP(&_redshiftLakehouseRegistration, "lakehouse-registration", "", "", "Lakehouse Registration")
	_redshiftCmd.Flags().StringVarP(&_redshiftLimitType, "limit-type", "", "", "Limit Type")
	_redshiftCmd.Flags().StringVarP(&_redshiftLoadSampleData, "load-sample-data", "", "", "Load Sample Data")
	_redshiftCmd.Flags().StringVarP(&_redshiftLogDestinationType, "log-destination-type", "", "", "Log Destination Type")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftLogExports, "log-exports", "", nil, "Log Exports")
	_redshiftCmd.Flags().StringVarP(&_redshiftMaintenanceTrackName, "maintenance-track-name", "", "", "Maintenance Track Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftManageMasterPassword, "manage-master-password", "", "", "Manage Master Password")
	_redshiftCmd.Flags().StringVarP(&_redshiftManual, "manual", "", "", "Manual")
	_redshiftCmd.Flags().StringVarP(&_redshiftManualSnapshotRetentionPeriod, "manual-snapshot-retention-period", "", "", "Manual Snapshot Retention Period")
	_redshiftCmd.Flags().StringVarP(&_redshiftMarker, "marker", "", "", "Marker")
	_redshiftCmd.Flags().StringVarP(&_redshiftMasterPasswordSecretKmsKeyId, "master-password-secret-kms-key-id", "", "", "Master Password Secret KMS Key ID")
	_redshiftCmd.Flags().StringVarP(&_redshiftMasterUserPassword, "master-user-password", "", "", "Master User Password")
	_redshiftCmd.Flags().StringVarP(&_redshiftMasterUsername, "master-username", "", "", "Master Username")
	_redshiftCmd.Flags().StringVarP(&_redshiftMaxRecords, "max-records", "", "", "Max Records")
	_redshiftCmd.Flags().StringVarP(&_redshiftMultiAZ, "multi-az", "", "", "Multi AZ")
	_redshiftCmd.Flags().StringVarP(&_redshiftNamespaceArn, "namespace-arn", "", "", "Namespace ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftNamespaceIdentifier, "namespace-identifier", "", "", "Namespace Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftNewClusterIdentifier, "new-cluster-identifier", "", "", "New Cluster Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftNewTableName, "new-table-name", "", "", "New Table Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftNextInvocations, "next-invocations", "", "", "Next Invocations")
	_redshiftCmd.Flags().StringVarP(&_redshiftNodeCount, "node-count", "", "", "Node Count")
	_redshiftCmd.Flags().StringVarP(&_redshiftNodeType, "node-type", "", "", "Node Type")
	_redshiftCmd.Flags().StringVarP(&_redshiftNumberOfNodes, "number-of-nodes", "", "", "Number Of Nodes")
	_redshiftCmd.Flags().StringVarP(&_redshiftOwnerAccount, "owner-account", "", "", "Owner Account")
	_redshiftCmd.Flags().StringVarP(&_redshiftParameterGroupFamily, "parameter-group-family", "", "", "Parameter Group Family")
	_redshiftCmd.Flags().StringVarP(&_redshiftParameterGroupName, "parameter-group-name", "", "", "Parameter Group Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftParameters, "parameters", "", "", "Parameters")
	_redshiftCmd.Flags().StringVarP(&_redshiftPartnerName, "partner-name", "", "", "Partner Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftPeriod, "period", "", "", "Period")
	_redshiftCmd.Flags().StringVarP(&_redshiftPolicy, "policy", "", "", "Policy")
	_redshiftCmd.Flags().StringVarP(&_redshiftPort, "port", "", "", "Port")
	_redshiftCmd.Flags().StringVarP(&_redshiftPreferredMaintenanceWindow, "preferred-maintenance-window", "", "", "Preferred Maintenance Window")
	_redshiftCmd.Flags().StringVarP(&_redshiftProducerArn, "producer-arn", "", "", "Producer ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftPubliclyAccessible, "publicly-accessible", "", "", "Publicly Accessible")
	_redshiftCmd.Flags().StringVarP(&_redshiftRedshiftIdcApplicationArn, "redshift-idc-application-arn", "", "", "Redshift Idc Application ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftRedshiftIdcApplicationName, "redshift-idc-application-name", "", "", "Redshift Idc Application Name")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftRemoveIamRoles, "remove-iam-roles", "", nil, "Remove IAM Roles")
	_redshiftCmd.Flags().StringVarP(&_redshiftReservedNodeExchangeRequestId, "reserved-node-exchange-request-id", "", "", "Reserved Node Exchange Request ID")
	_redshiftCmd.Flags().StringVarP(&_redshiftReservedNodeId, "reserved-node-id", "", "", "Reserved Node ID")
	_redshiftCmd.Flags().StringVarP(&_redshiftReservedNodeOfferingId, "reserved-node-offering-id", "", "", "Reserved Node Offering ID")
	_redshiftCmd.Flags().StringVarP(&_redshiftResetAllParameters, "reset-all-parameters", "", "", "Reset All Parameters")
	_redshiftCmd.Flags().StringVarP(&_redshiftResourceArn, "resource-arn", "", "", "Resource ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftResourceName, "resource-name", "", "", "Resource Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftResourceOwner, "resource-owner", "", "", "Resource Owner")
	_redshiftCmd.Flags().StringVarP(&_redshiftResourceType, "resource-type", "", "", "Resource Type")
	_redshiftCmd.Flags().StringVarP(&_redshiftRetentionPeriod, "retention-period", "", "", "Retention Period")
	_redshiftCmd.Flags().StringVarP(&_redshiftRevisionTarget, "revision-target", "", "", "Revision Target")
	_redshiftCmd.Flags().StringVarP(&_redshiftS3KeyPrefix, "s3-key-prefix", "", "", "S3 Key Prefix")
	_redshiftCmd.Flags().StringVarP(&_redshiftSchedule, "schedule", "", "", "Schedule")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftScheduleDefinitions, "schedule-definitions", "", nil, "Schedule Definitions")
	_redshiftCmd.Flags().StringVarP(&_redshiftScheduleDescription, "schedule-description", "", "", "Schedule Description")
	_redshiftCmd.Flags().StringVarP(&_redshiftScheduleIdentifier, "schedule-identifier", "", "", "Schedule Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftScheduledActionDescription, "scheduled-action-description", "", "", "Scheduled Action Description")
	_redshiftCmd.Flags().StringVarP(&_redshiftScheduledActionName, "scheduled-action-name", "", "", "Scheduled Action Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftServiceIntegrations, "service-integrations", "", "", "Service Integrations")
	_redshiftCmd.Flags().StringVarP(&_redshiftSeverity, "severity", "", "", "Severity")
	_redshiftCmd.Flags().StringVarP(&_redshiftSkipFinalClusterSnapshot, "skip-final-cluster-snapshot", "", "", "Skip Final Cluster Snapshot")
	_redshiftCmd.Flags().StringVarP(&_redshiftSnapshotArn, "snapshot-arn", "", "", "Snapshot ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftSnapshotClusterIdentifier, "snapshot-cluster-identifier", "", "", "Snapshot Cluster Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftSnapshotCopyGrantName, "snapshot-copy-grant-name", "", "", "Snapshot Copy Grant Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftSnapshotIdentifier, "snapshot-identifier", "", "", "Snapshot Identifier")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftSnapshotIdentifierList, "snapshot-identifier-list", "", nil, "Snapshot Identifier List")
	_redshiftCmd.Flags().StringVarP(&_redshiftSnapshotScheduleIdentifier, "snapshot-schedule-identifier", "", "", "Snapshot Schedule Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftSnapshotType, "snapshot-type", "", "", "Snapshot Type")
	_redshiftCmd.Flags().StringVarP(&_redshiftSnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftSortingEntities, "sorting-entities", "", "", "Sorting Entities")
	_redshiftCmd.Flags().StringVarP(&_redshiftSource, "source", "", "", "Source")
	_redshiftCmd.Flags().StringVarP(&_redshiftSourceArn, "source-arn", "", "", "Source ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftSourceDatabaseName, "source-database-name", "", "", "Source Database Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftSourceIdentifier, "source-identifier", "", "", "Source Identifier")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftSourceIds, "source-ids", "", nil, "Source Ids")
	_redshiftCmd.Flags().StringVarP(&_redshiftSourceSchemaName, "source-schema-name", "", "", "Source Schema Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftSourceSnapshotClusterIdentifier, "source-snapshot-cluster-identifier", "", "", "Source Snapshot Cluster Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftSourceSnapshotIdentifier, "source-snapshot-identifier", "", "", "Source Snapshot Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftSourceTableName, "source-table-name", "", "", "Source Table Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftSourceType, "source-type", "", "", "Source Type")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftSsoTagKeys, "sso-tag-keys", "", nil, "Sso Tag Keys")
	_redshiftCmd.Flags().StringVarP(&_redshiftStartTime, "start-time", "", "", "Start Time")
	_redshiftCmd.Flags().StringVarP(&_redshiftStatus, "status", "", "", "Status")
	_redshiftCmd.Flags().StringVarP(&_redshiftStatusMessage, "status-message", "", "", "Status Message")
	_redshiftCmd.Flags().StringVarP(&_redshiftSubnetGroupName, "subnet-group-name", "", "", "Subnet Group Name")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_redshiftCmd.Flags().StringVarP(&_redshiftSubscriptionName, "subscription-name", "", "", "Subscription Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftTableRestoreRequestId, "table-restore-request-id", "", "", "Table Restore Request ID")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftTagKeys, "tag-keys", "", nil, "Tag Keys")
	_redshiftCmd.Flags().StringVarP(&_redshiftTagList, "tag-list", "", "", "Tag List")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftTagValues, "tag-values", "", nil, "Tag Values")
	_redshiftCmd.Flags().StringVarP(&_redshiftTags, "tags", "", "", "Tags")
	_redshiftCmd.Flags().StringVarP(&_redshiftTargetAction, "target-action", "", "", "Target Action")
	_redshiftCmd.Flags().StringVarP(&_redshiftTargetActionType, "target-action-type", "", "", "Target Action Type")
	_redshiftCmd.Flags().StringVarP(&_redshiftTargetArn, "target-arn", "", "", "Target ARN")
	_redshiftCmd.Flags().StringVarP(&_redshiftTargetDatabaseName, "target-database-name", "", "", "Target Database Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftTargetReservedNodeOfferingId, "target-reserved-node-offering-id", "", "", "Target Reserved Node Offering ID")
	_redshiftCmd.Flags().StringVarP(&_redshiftTargetSchemaName, "target-schema-name", "", "", "Target Schema Name")
	_redshiftCmd.Flags().StringVarP(&_redshiftTargetSnapshotIdentifier, "target-snapshot-identifier", "", "", "Target Snapshot Identifier")
	_redshiftCmd.Flags().StringVarP(&_redshiftUsageLimitId, "usage-limit-id", "", "", "Usage Limit ID")
	_redshiftCmd.Flags().StringVarP(&_redshiftVpcId, "vpc-id", "", "", "VPC ID")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftVpcIds, "vpc-ids", "", nil, "VPC Ids")
	_redshiftCmd.Flags().StringSliceVarP(&_redshiftVpcSecurityGroupIds, "vpc-security-group-ids", "", nil, "VPC Security Group Ids")

	_redshiftCmd.Flags().BoolVarP(&_redshiftAcceptReservedNodeExchange, "accept-reserved-node-exchange", "", false, "Accept Reserved Node Exchange")
	_redshiftCmd.Flags().BoolVarP(&_redshiftAddPartner, "add-partner", "", false, "Add Partner")
	_redshiftCmd.Flags().BoolVarP(&_redshiftAssociateDataShareConsumer, "associate-data-share-consumer", "", false, "Associate Data Share Consumer")
	_redshiftCmd.Flags().BoolVarP(&_redshiftAuthorizeClusterSecurityGroupIngress, "authorize-cluster-security-group-ingress", "", false, "Authorize Cluster Security Group Ingress")
	_redshiftCmd.Flags().BoolVarP(&_redshiftAuthorizeDataShare, "authorize-data-share", "", false, "Authorize Data Share")
	_redshiftCmd.Flags().BoolVarP(&_redshiftAuthorizeEndpointAccess, "authorize-endpoint-access", "", false, "Authorize Endpoint Access")
	_redshiftCmd.Flags().BoolVarP(&_redshiftAuthorizeSnapshotAccess, "authorize-snapshot-access", "", false, "Authorize Snapshot Access")
	_redshiftCmd.Flags().BoolVarP(&_redshiftBatchDeleteClusterSnapshots, "batch-delete-cluster-snapshots", "", false, "Batch Delete Cluster Snapshots")
	_redshiftCmd.Flags().BoolVarP(&_redshiftBatchModifyClusterSnapshots, "batch-modify-cluster-snapshots", "", false, "Batch Modify Cluster Snapshots")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCancelResize, "cancel-resize", "", false, "Cancel Resize")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCopyClusterSnapshot, "copy-cluster-snapshot", "", false, "Copy Cluster Snapshot")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateAuthenticationProfile, "create-authentication-profile", "", false, "Create Authentication Profile")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateCluster, "create-cluster", "", false, "Create Cluster")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateClusterParameterGroup, "create-cluster-parameter-group", "", false, "Create Cluster Parameter Group")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateClusterSecurityGroup, "create-cluster-security-group", "", false, "Create Cluster Security Group")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateClusterSnapshot, "create-cluster-snapshot", "", false, "Create Cluster Snapshot")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateClusterSubnetGroup, "create-cluster-subnet-group", "", false, "Create Cluster Subnet Group")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateCustomDomainAssociation, "create-custom-domain-association", "", false, "Create Custom Domain Association")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateEndpointAccess, "create-endpoint-access", "", false, "Create Endpoint Access")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateEventSubscription, "create-event-subscription", "", false, "Create Event Subscription")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateHsmClientCertificate, "create-hsm-client-certificate", "", false, "Create Hsm Client Certificate")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateHsmConfiguration, "create-hsm-configuration", "", false, "Create Hsm Configuration")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateIntegration, "create-integration", "", false, "Create Integration")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateRedshiftIdcApplication, "create-redshift-idc-application", "", false, "Create Redshift Idc Application")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateScheduledAction, "create-scheduled-action", "", false, "Create Scheduled Action")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateSnapshotCopyGrant, "create-snapshot-copy-grant", "", false, "Create Snapshot Copy Grant")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateSnapshotSchedule, "create-snapshot-schedule", "", false, "Create Snapshot Schedule")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateTags, "create-tags", "", false, "Create Tags")
	_redshiftCmd.Flags().BoolVarP(&_redshiftCreateUsageLimit, "create-usage-limit", "", false, "Create Usage Limit")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeauthorizeDataShare, "deauthorize-data-share", "", false, "Deauthorize Data Share")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteAuthenticationProfile, "delete-authentication-profile", "", false, "Delete Authentication Profile")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteClusterParameterGroup, "delete-cluster-parameter-group", "", false, "Delete Cluster Parameter Group")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteClusterSecurityGroup, "delete-cluster-security-group", "", false, "Delete Cluster Security Group")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteClusterSnapshot, "delete-cluster-snapshot", "", false, "Delete Cluster Snapshot")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteClusterSubnetGroup, "delete-cluster-subnet-group", "", false, "Delete Cluster Subnet Group")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteCustomDomainAssociation, "delete-custom-domain-association", "", false, "Delete Custom Domain Association")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteEndpointAccess, "delete-endpoint-access", "", false, "Delete Endpoint Access")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteEventSubscription, "delete-event-subscription", "", false, "Delete Event Subscription")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteHsmClientCertificate, "delete-hsm-client-certificate", "", false, "Delete Hsm Client Certificate")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteHsmConfiguration, "delete-hsm-configuration", "", false, "Delete Hsm Configuration")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteIntegration, "delete-integration", "", false, "Delete Integration")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeletePartner, "delete-partner", "", false, "Delete Partner")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteRedshiftIdcApplication, "delete-redshift-idc-application", "", false, "Delete Redshift Idc Application")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteScheduledAction, "delete-scheduled-action", "", false, "Delete Scheduled Action")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteSnapshotCopyGrant, "delete-snapshot-copy-grant", "", false, "Delete Snapshot Copy Grant")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteSnapshotSchedule, "delete-snapshot-schedule", "", false, "Delete Snapshot Schedule")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteTags, "delete-tags", "", false, "Delete Tags")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeleteUsageLimit, "delete-usage-limit", "", false, "Delete Usage Limit")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDeregisterNamespace, "deregister-namespace", "", false, "Deregister Namespace")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeAccountAttributes, "describe-account-attributes", "", false, "Describe Account Attributes")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeAuthenticationProfiles, "describe-authentication-profiles", "", false, "Describe Authentication Profiles")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeClusterDbRevisions, "describe-cluster-db-revisions", "", false, "Describe Cluster DB Revisions")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeClusterParameterGroups, "describe-cluster-parameter-groups", "", false, "Describe Cluster Parameter Groups")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeClusterParameters, "describe-cluster-parameters", "", false, "Describe Cluster Parameters")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeClusterSecurityGroups, "describe-cluster-security-groups", "", false, "Describe Cluster Security Groups")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeClusterSnapshots, "describe-cluster-snapshots", "", false, "Describe Cluster Snapshots")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeClusterSubnetGroups, "describe-cluster-subnet-groups", "", false, "Describe Cluster Subnet Groups")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeClusterTracks, "describe-cluster-tracks", "", false, "Describe Cluster Tracks")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeClusterVersions, "describe-cluster-versions", "", false, "Describe Cluster Versions")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeClusters, "describe-clusters", "", false, "Describe Clusters")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeCustomDomainAssociations, "describe-custom-domain-associations", "", false, "Describe Custom Domain Associations")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeDataShares, "describe-data-shares", "", false, "Describe Data Shares")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeDataSharesForConsumer, "describe-data-shares-for-consumer", "", false, "Describe Data Shares For Consumer")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeDataSharesForProducer, "describe-data-shares-for-producer", "", false, "Describe Data Shares For Producer")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeDefaultClusterParameters, "describe-default-cluster-parameters", "", false, "Describe Default Cluster Parameters")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeEndpointAccess, "describe-endpoint-access", "", false, "Describe Endpoint Access")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeEndpointAuthorization, "describe-endpoint-authorization", "", false, "Describe Endpoint Authorization")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeEventCategories, "describe-event-categories", "", false, "Describe Event Categories")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeEventSubscriptions, "describe-event-subscriptions", "", false, "Describe Event Subscriptions")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeEvents, "describe-events", "", false, "Describe Events")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeHsmClientCertificates, "describe-hsm-client-certificates", "", false, "Describe Hsm Client Certificates")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeHsmConfigurations, "describe-hsm-configurations", "", false, "Describe Hsm Configurations")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeInboundIntegrations, "describe-inbound-integrations", "", false, "Describe Inbound Integrations")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeIntegrations, "describe-integrations", "", false, "Describe Integrations")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeLoggingStatus, "describe-logging-status", "", false, "Describe Logging Status")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeNodeConfigurationOptions, "describe-node-configuration-options", "", false, "Describe Node Configuration Options")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeOrderableClusterOptions, "describe-orderable-cluster-options", "", false, "Describe Orderable Cluster Options")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribePartners, "describe-partners", "", false, "Describe Partners")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeRedshiftIdcApplications, "describe-redshift-idc-applications", "", false, "Describe Redshift Idc Applications")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeReservedNodeExchangeStatus, "describe-reserved-node-exchange-status", "", false, "Describe Reserved Node Exchange Status")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeReservedNodeOfferings, "describe-reserved-node-offerings", "", false, "Describe Reserved Node Offerings")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeReservedNodes, "describe-reserved-nodes", "", false, "Describe Reserved Nodes")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeResize, "describe-resize", "", false, "Describe Resize")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeScheduledActions, "describe-scheduled-actions", "", false, "Describe Scheduled Actions")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeSnapshotCopyGrants, "describe-snapshot-copy-grants", "", false, "Describe Snapshot Copy Grants")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeSnapshotSchedules, "describe-snapshot-schedules", "", false, "Describe Snapshot Schedules")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeStorage, "describe-storage", "", false, "Describe Storage")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeTableRestoreStatus, "describe-table-restore-status", "", false, "Describe Table Restore Status")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeTags, "describe-tags", "", false, "Describe Tags")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDescribeUsageLimits, "describe-usage-limits", "", false, "Describe Usage Limits")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDisableLogging, "disable-logging", "", false, "Disable Logging")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDisableSnapshotCopy, "disable-snapshot-copy", "", false, "Disable Snapshot Copy")
	_redshiftCmd.Flags().BoolVarP(&_redshiftDisassociateDataShareConsumer, "disassociate-data-share-consumer", "", false, "Disassociate Data Share Consumer")
	_redshiftCmd.Flags().BoolVarP(&_redshiftEnableLogging, "enable-logging", "", false, "Enable Logging")
	_redshiftCmd.Flags().BoolVarP(&_redshiftEnableSnapshotCopy, "enable-snapshot-copy", "", false, "Enable Snapshot Copy")
	_redshiftCmd.Flags().BoolVarP(&_redshiftFailoverPrimaryCompute, "failover-primary-compute", "", false, "Failover Primary Compute")
	_redshiftCmd.Flags().BoolVarP(&_redshiftGetClusterCredentials, "get-cluster-credentials", "", false, "Get Cluster Credentials")
	_redshiftCmd.Flags().BoolVarP(&_redshiftGetClusterCredentialsWithIAM, "get-cluster-credentials-with-iam", "", false, "Get Cluster Credentials With IAM")
	_redshiftCmd.Flags().BoolVarP(&_redshiftGetIdentityCenterAuthToken, "get-identity-center-auth-token", "", false, "Get Identity Center Auth Token")
	_redshiftCmd.Flags().BoolVarP(&_redshiftGetReservedNodeExchangeConfigurationOptions, "get-reserved-node-exchange-configuration-options", "", false, "Get Reserved Node Exchange Configuration Options")
	_redshiftCmd.Flags().BoolVarP(&_redshiftGetReservedNodeExchangeOfferings, "get-reserved-node-exchange-offerings", "", false, "Get Reserved Node Exchange Offerings")
	_redshiftCmd.Flags().BoolVarP(&_redshiftGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_redshiftCmd.Flags().BoolVarP(&_redshiftListRecommendations, "list-recommendations", "", false, "List Recommendations")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyAquaConfiguration, "modify-aqua-configuration", "", false, "Modify Aqua Configuration")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyAuthenticationProfile, "modify-authentication-profile", "", false, "Modify Authentication Profile")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyCluster, "modify-cluster", "", false, "Modify Cluster")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyClusterDbRevision, "modify-cluster-db-revision", "", false, "Modify Cluster DB Revision")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyClusterIamRoles, "modify-cluster-iam-roles", "", false, "Modify Cluster IAM Roles")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyClusterMaintenance, "modify-cluster-maintenance", "", false, "Modify Cluster Maintenance")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyClusterParameterGroup, "modify-cluster-parameter-group", "", false, "Modify Cluster Parameter Group")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyClusterSnapshot, "modify-cluster-snapshot", "", false, "Modify Cluster Snapshot")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyClusterSnapshotSchedule, "modify-cluster-snapshot-schedule", "", false, "Modify Cluster Snapshot Schedule")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyClusterSubnetGroup, "modify-cluster-subnet-group", "", false, "Modify Cluster Subnet Group")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyCustomDomainAssociation, "modify-custom-domain-association", "", false, "Modify Custom Domain Association")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyEndpointAccess, "modify-endpoint-access", "", false, "Modify Endpoint Access")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyEventSubscription, "modify-event-subscription", "", false, "Modify Event Subscription")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyIntegration, "modify-integration", "", false, "Modify Integration")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyLakehouseConfiguration, "modify-lakehouse-configuration", "", false, "Modify Lakehouse Configuration")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyRedshiftIdcApplication, "modify-redshift-idc-application", "", false, "Modify Redshift Idc Application")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyScheduledAction, "modify-scheduled-action", "", false, "Modify Scheduled Action")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifySnapshotCopyRetentionPeriod, "modify-snapshot-copy-retention-period", "", false, "Modify Snapshot Copy Retention Period")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifySnapshotSchedule, "modify-snapshot-schedule", "", false, "Modify Snapshot Schedule")
	_redshiftCmd.Flags().BoolVarP(&_redshiftModifyUsageLimit, "modify-usage-limit", "", false, "Modify Usage Limit")
	_redshiftCmd.Flags().BoolVarP(&_redshiftPauseCluster, "pause-cluster", "", false, "Pause Cluster")
	_redshiftCmd.Flags().BoolVarP(&_redshiftPurchaseReservedNodeOffering, "purchase-reserved-node-offering", "", false, "Purchase Reserved Node Offering")
	_redshiftCmd.Flags().BoolVarP(&_redshiftPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_redshiftCmd.Flags().BoolVarP(&_redshiftRebootCluster, "reboot-cluster", "", false, "Reboot Cluster")
	_redshiftCmd.Flags().BoolVarP(&_redshiftRegisterNamespace, "register-namespace", "", false, "Register Namespace")
	_redshiftCmd.Flags().BoolVarP(&_redshiftRejectDataShare, "reject-data-share", "", false, "Reject Data Share")
	_redshiftCmd.Flags().BoolVarP(&_redshiftResetClusterParameterGroup, "reset-cluster-parameter-group", "", false, "Reset Cluster Parameter Group")
	_redshiftCmd.Flags().BoolVarP(&_redshiftResizeCluster, "resize-cluster", "", false, "Resize Cluster")
	_redshiftCmd.Flags().BoolVarP(&_redshiftRestoreFromClusterSnapshot, "restore-from-cluster-snapshot", "", false, "Restore From Cluster Snapshot")
	_redshiftCmd.Flags().BoolVarP(&_redshiftRestoreTableFromClusterSnapshot, "restore-table-from-cluster-snapshot", "", false, "Restore Table From Cluster Snapshot")
	_redshiftCmd.Flags().BoolVarP(&_redshiftResumeCluster, "resume-cluster", "", false, "Resume Cluster")
	_redshiftCmd.Flags().BoolVarP(&_redshiftRevokeClusterSecurityGroupIngress, "revoke-cluster-security-group-ingress", "", false, "Revoke Cluster Security Group Ingress")
	_redshiftCmd.Flags().BoolVarP(&_redshiftRevokeEndpointAccess, "revoke-endpoint-access", "", false, "Revoke Endpoint Access")
	_redshiftCmd.Flags().BoolVarP(&_redshiftRevokeSnapshotAccess, "revoke-snapshot-access", "", false, "Revoke Snapshot Access")
	_redshiftCmd.Flags().BoolVarP(&_redshiftRotateEncryptionKey, "rotate-encryption-key", "", false, "Rotate Encryption Key")
	_redshiftCmd.Flags().BoolVarP(&_redshiftUpdatePartnerStatus, "update-partner-status", "", false, "Update Partner Status")

}
