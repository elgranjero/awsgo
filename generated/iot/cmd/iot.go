package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iotCmd represents the iot command
var _iotCmd = &cobra.Command{
	Use:   "iot",
	Short: "AWS iot CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := iot.NewFromConfig(cfg)
		if _iotAcceptCertificateTransfer {
			iot_AcceptCertificateTransfer(cfg, client)
			return
		}
		if _iotAddThingToBillingGroup {
			iot_AddThingToBillingGroup(cfg, client)
			return
		}
		if _iotAddThingToThingGroup {
			iot_AddThingToThingGroup(cfg, client)
			return
		}
		if _iotAssociateSbomWithPackageVersion {
			iot_AssociateSbomWithPackageVersion(cfg, client)
			return
		}
		if _iotAssociateTargetsWithJob {
			iot_AssociateTargetsWithJob(cfg, client)
			return
		}
		if _iotAttachPolicy {
			iot_AttachPolicy(cfg, client)
			return
		}
		if _iotAttachPrincipalPolicy {
			iot_AttachPrincipalPolicy(cfg, client)
			return
		}
		if _iotAttachSecurityProfile {
			iot_AttachSecurityProfile(cfg, client)
			return
		}
		if _iotAttachThingPrincipal {
			iot_AttachThingPrincipal(cfg, client)
			return
		}
		if _iotCancelAuditMitigationActionsTask {
			iot_CancelAuditMitigationActionsTask(cfg, client)
			return
		}
		if _iotCancelAuditTask {
			iot_CancelAuditTask(cfg, client)
			return
		}
		if _iotCancelCertificateTransfer {
			iot_CancelCertificateTransfer(cfg, client)
			return
		}
		if _iotCancelDetectMitigationActionsTask {
			iot_CancelDetectMitigationActionsTask(cfg, client)
			return
		}
		if _iotCancelJob {
			iot_CancelJob(cfg, client)
			return
		}
		if _iotCancelJobExecution {
			iot_CancelJobExecution(cfg, client)
			return
		}
		if _iotClearDefaultAuthorizer {
			iot_ClearDefaultAuthorizer(cfg, client)
			return
		}
		if _iotConfirmTopicRuleDestination {
			iot_ConfirmTopicRuleDestination(cfg, client)
			return
		}
		if _iotCreateAuditSuppression {
			iot_CreateAuditSuppression(cfg, client)
			return
		}
		if _iotCreateAuthorizer {
			iot_CreateAuthorizer(cfg, client)
			return
		}
		if _iotCreateBillingGroup {
			iot_CreateBillingGroup(cfg, client)
			return
		}
		if _iotCreateCertificateFromCsr {
			iot_CreateCertificateFromCsr(cfg, client)
			return
		}
		if _iotCreateCertificateProvider {
			iot_CreateCertificateProvider(cfg, client)
			return
		}
		if _iotCreateCommand {
			iot_CreateCommand(cfg, client)
			return
		}
		if _iotCreateCustomMetric {
			iot_CreateCustomMetric(cfg, client)
			return
		}
		if _iotCreateDimension {
			iot_CreateDimension(cfg, client)
			return
		}
		if _iotCreateDomainConfiguration {
			iot_CreateDomainConfiguration(cfg, client)
			return
		}
		if _iotCreateDynamicThingGroup {
			iot_CreateDynamicThingGroup(cfg, client)
			return
		}
		if _iotCreateFleetMetric {
			iot_CreateFleetMetric(cfg, client)
			return
		}
		if _iotCreateJob {
			iot_CreateJob(cfg, client)
			return
		}
		if _iotCreateJobTemplate {
			iot_CreateJobTemplate(cfg, client)
			return
		}
		if _iotCreateKeysAndCertificate {
			iot_CreateKeysAndCertificate(cfg, client)
			return
		}
		if _iotCreateMitigationAction {
			iot_CreateMitigationAction(cfg, client)
			return
		}
		if _iotCreateOTAUpdate {
			iot_CreateOTAUpdate(cfg, client)
			return
		}
		if _iotCreatePackage {
			iot_CreatePackage(cfg, client)
			return
		}
		if _iotCreatePackageVersion {
			iot_CreatePackageVersion(cfg, client)
			return
		}
		if _iotCreatePolicy {
			iot_CreatePolicy(cfg, client)
			return
		}
		if _iotCreatePolicyVersion {
			iot_CreatePolicyVersion(cfg, client)
			return
		}
		if _iotCreateProvisioningClaim {
			iot_CreateProvisioningClaim(cfg, client)
			return
		}
		if _iotCreateProvisioningTemplate {
			iot_CreateProvisioningTemplate(cfg, client)
			return
		}
		if _iotCreateProvisioningTemplateVersion {
			iot_CreateProvisioningTemplateVersion(cfg, client)
			return
		}
		if _iotCreateRoleAlias {
			iot_CreateRoleAlias(cfg, client)
			return
		}
		if _iotCreateScheduledAudit {
			iot_CreateScheduledAudit(cfg, client)
			return
		}
		if _iotCreateSecurityProfile {
			iot_CreateSecurityProfile(cfg, client)
			return
		}
		if _iotCreateStream {
			iot_CreateStream(cfg, client)
			return
		}
		if _iotCreateThing {
			iot_CreateThing(cfg, client)
			return
		}
		if _iotCreateThingGroup {
			iot_CreateThingGroup(cfg, client)
			return
		}
		if _iotCreateThingType {
			iot_CreateThingType(cfg, client)
			return
		}
		if _iotCreateTopicRule {
			iot_CreateTopicRule(cfg, client)
			return
		}
		if _iotCreateTopicRuleDestination {
			iot_CreateTopicRuleDestination(cfg, client)
			return
		}
		if _iotDeleteAccountAuditConfiguration {
			iot_DeleteAccountAuditConfiguration(cfg, client)
			return
		}
		if _iotDeleteAuditSuppression {
			iot_DeleteAuditSuppression(cfg, client)
			return
		}
		if _iotDeleteAuthorizer {
			iot_DeleteAuthorizer(cfg, client)
			return
		}
		if _iotDeleteBillingGroup {
			iot_DeleteBillingGroup(cfg, client)
			return
		}
		if _iotDeleteCACertificate {
			iot_DeleteCACertificate(cfg, client)
			return
		}
		if _iotDeleteCertificate {
			iot_DeleteCertificate(cfg, client)
			return
		}
		if _iotDeleteCertificateProvider {
			iot_DeleteCertificateProvider(cfg, client)
			return
		}
		if _iotDeleteCommand {
			iot_DeleteCommand(cfg, client)
			return
		}
		if _iotDeleteCommandExecution {
			iot_DeleteCommandExecution(cfg, client)
			return
		}
		if _iotDeleteCustomMetric {
			iot_DeleteCustomMetric(cfg, client)
			return
		}
		if _iotDeleteDimension {
			iot_DeleteDimension(cfg, client)
			return
		}
		if _iotDeleteDomainConfiguration {
			iot_DeleteDomainConfiguration(cfg, client)
			return
		}
		if _iotDeleteDynamicThingGroup {
			iot_DeleteDynamicThingGroup(cfg, client)
			return
		}
		if _iotDeleteFleetMetric {
			iot_DeleteFleetMetric(cfg, client)
			return
		}
		if _iotDeleteJob {
			iot_DeleteJob(cfg, client)
			return
		}
		if _iotDeleteJobExecution {
			iot_DeleteJobExecution(cfg, client)
			return
		}
		if _iotDeleteJobTemplate {
			iot_DeleteJobTemplate(cfg, client)
			return
		}
		if _iotDeleteMitigationAction {
			iot_DeleteMitigationAction(cfg, client)
			return
		}
		if _iotDeleteOTAUpdate {
			iot_DeleteOTAUpdate(cfg, client)
			return
		}
		if _iotDeletePackage {
			iot_DeletePackage(cfg, client)
			return
		}
		if _iotDeletePackageVersion {
			iot_DeletePackageVersion(cfg, client)
			return
		}
		if _iotDeletePolicy {
			iot_DeletePolicy(cfg, client)
			return
		}
		if _iotDeletePolicyVersion {
			iot_DeletePolicyVersion(cfg, client)
			return
		}
		if _iotDeleteProvisioningTemplate {
			iot_DeleteProvisioningTemplate(cfg, client)
			return
		}
		if _iotDeleteProvisioningTemplateVersion {
			iot_DeleteProvisioningTemplateVersion(cfg, client)
			return
		}
		if _iotDeleteRegistrationCode {
			iot_DeleteRegistrationCode(cfg, client)
			return
		}
		if _iotDeleteRoleAlias {
			iot_DeleteRoleAlias(cfg, client)
			return
		}
		if _iotDeleteScheduledAudit {
			iot_DeleteScheduledAudit(cfg, client)
			return
		}
		if _iotDeleteSecurityProfile {
			iot_DeleteSecurityProfile(cfg, client)
			return
		}
		if _iotDeleteStream {
			iot_DeleteStream(cfg, client)
			return
		}
		if _iotDeleteThing {
			iot_DeleteThing(cfg, client)
			return
		}
		if _iotDeleteThingGroup {
			iot_DeleteThingGroup(cfg, client)
			return
		}
		if _iotDeleteThingType {
			iot_DeleteThingType(cfg, client)
			return
		}
		if _iotDeleteTopicRule {
			iot_DeleteTopicRule(cfg, client)
			return
		}
		if _iotDeleteTopicRuleDestination {
			iot_DeleteTopicRuleDestination(cfg, client)
			return
		}
		if _iotDeleteV2LoggingLevel {
			iot_DeleteV2LoggingLevel(cfg, client)
			return
		}
		if _iotDeprecateThingType {
			iot_DeprecateThingType(cfg, client)
			return
		}
		if _iotDescribeAccountAuditConfiguration {
			iot_DescribeAccountAuditConfiguration(cfg, client)
			return
		}
		if _iotDescribeAuditFinding {
			iot_DescribeAuditFinding(cfg, client)
			return
		}
		if _iotDescribeAuditMitigationActionsTask {
			iot_DescribeAuditMitigationActionsTask(cfg, client)
			return
		}
		if _iotDescribeAuditSuppression {
			iot_DescribeAuditSuppression(cfg, client)
			return
		}
		if _iotDescribeAuditTask {
			iot_DescribeAuditTask(cfg, client)
			return
		}
		if _iotDescribeAuthorizer {
			iot_DescribeAuthorizer(cfg, client)
			return
		}
		if _iotDescribeBillingGroup {
			iot_DescribeBillingGroup(cfg, client)
			return
		}
		if _iotDescribeCACertificate {
			iot_DescribeCACertificate(cfg, client)
			return
		}
		if _iotDescribeCertificate {
			iot_DescribeCertificate(cfg, client)
			return
		}
		if _iotDescribeCertificateProvider {
			iot_DescribeCertificateProvider(cfg, client)
			return
		}
		if _iotDescribeCustomMetric {
			iot_DescribeCustomMetric(cfg, client)
			return
		}
		if _iotDescribeDefaultAuthorizer {
			iot_DescribeDefaultAuthorizer(cfg, client)
			return
		}
		if _iotDescribeDetectMitigationActionsTask {
			iot_DescribeDetectMitigationActionsTask(cfg, client)
			return
		}
		if _iotDescribeDimension {
			iot_DescribeDimension(cfg, client)
			return
		}
		if _iotDescribeDomainConfiguration {
			iot_DescribeDomainConfiguration(cfg, client)
			return
		}
		if _iotDescribeEncryptionConfiguration {
			iot_DescribeEncryptionConfiguration(cfg, client)
			return
		}
		if _iotDescribeEndpoint {
			iot_DescribeEndpoint(cfg, client)
			return
		}
		if _iotDescribeEventConfigurations {
			iot_DescribeEventConfigurations(cfg, client)
			return
		}
		if _iotDescribeFleetMetric {
			iot_DescribeFleetMetric(cfg, client)
			return
		}
		if _iotDescribeIndex {
			iot_DescribeIndex(cfg, client)
			return
		}
		if _iotDescribeJob {
			iot_DescribeJob(cfg, client)
			return
		}
		if _iotDescribeJobExecution {
			iot_DescribeJobExecution(cfg, client)
			return
		}
		if _iotDescribeJobTemplate {
			iot_DescribeJobTemplate(cfg, client)
			return
		}
		if _iotDescribeManagedJobTemplate {
			iot_DescribeManagedJobTemplate(cfg, client)
			return
		}
		if _iotDescribeMitigationAction {
			iot_DescribeMitigationAction(cfg, client)
			return
		}
		if _iotDescribeProvisioningTemplate {
			iot_DescribeProvisioningTemplate(cfg, client)
			return
		}
		if _iotDescribeProvisioningTemplateVersion {
			iot_DescribeProvisioningTemplateVersion(cfg, client)
			return
		}
		if _iotDescribeRoleAlias {
			iot_DescribeRoleAlias(cfg, client)
			return
		}
		if _iotDescribeScheduledAudit {
			iot_DescribeScheduledAudit(cfg, client)
			return
		}
		if _iotDescribeSecurityProfile {
			iot_DescribeSecurityProfile(cfg, client)
			return
		}
		if _iotDescribeStream {
			iot_DescribeStream(cfg, client)
			return
		}
		if _iotDescribeThing {
			iot_DescribeThing(cfg, client)
			return
		}
		if _iotDescribeThingGroup {
			iot_DescribeThingGroup(cfg, client)
			return
		}
		if _iotDescribeThingRegistrationTask {
			iot_DescribeThingRegistrationTask(cfg, client)
			return
		}
		if _iotDescribeThingType {
			iot_DescribeThingType(cfg, client)
			return
		}
		if _iotDetachPolicy {
			iot_DetachPolicy(cfg, client)
			return
		}
		if _iotDetachPrincipalPolicy {
			iot_DetachPrincipalPolicy(cfg, client)
			return
		}
		if _iotDetachSecurityProfile {
			iot_DetachSecurityProfile(cfg, client)
			return
		}
		if _iotDetachThingPrincipal {
			iot_DetachThingPrincipal(cfg, client)
			return
		}
		if _iotDisableTopicRule {
			iot_DisableTopicRule(cfg, client)
			return
		}
		if _iotDisassociateSbomFromPackageVersion {
			iot_DisassociateSbomFromPackageVersion(cfg, client)
			return
		}
		if _iotEnableTopicRule {
			iot_EnableTopicRule(cfg, client)
			return
		}
		if _iotGetBehaviorModelTrainingSummaries {
			iot_GetBehaviorModelTrainingSummaries(cfg, client)
			return
		}
		if _iotGetBucketsAggregation {
			iot_GetBucketsAggregation(cfg, client)
			return
		}
		if _iotGetCardinality {
			iot_GetCardinality(cfg, client)
			return
		}
		if _iotGetCommand {
			iot_GetCommand(cfg, client)
			return
		}
		if _iotGetCommandExecution {
			iot_GetCommandExecution(cfg, client)
			return
		}
		if _iotGetEffectivePolicies {
			iot_GetEffectivePolicies(cfg, client)
			return
		}
		if _iotGetIndexingConfiguration {
			iot_GetIndexingConfiguration(cfg, client)
			return
		}
		if _iotGetJobDocument {
			iot_GetJobDocument(cfg, client)
			return
		}
		if _iotGetLoggingOptions {
			iot_GetLoggingOptions(cfg, client)
			return
		}
		if _iotGetOTAUpdate {
			iot_GetOTAUpdate(cfg, client)
			return
		}
		if _iotGetPackage {
			iot_GetPackage(cfg, client)
			return
		}
		if _iotGetPackageConfiguration {
			iot_GetPackageConfiguration(cfg, client)
			return
		}
		if _iotGetPackageVersion {
			iot_GetPackageVersion(cfg, client)
			return
		}
		if _iotGetPercentiles {
			iot_GetPercentiles(cfg, client)
			return
		}
		if _iotGetPolicy {
			iot_GetPolicy(cfg, client)
			return
		}
		if _iotGetPolicyVersion {
			iot_GetPolicyVersion(cfg, client)
			return
		}
		if _iotGetRegistrationCode {
			iot_GetRegistrationCode(cfg, client)
			return
		}
		if _iotGetStatistics {
			iot_GetStatistics(cfg, client)
			return
		}
		if _iotGetThingConnectivityData {
			iot_GetThingConnectivityData(cfg, client)
			return
		}
		if _iotGetTopicRule {
			iot_GetTopicRule(cfg, client)
			return
		}
		if _iotGetTopicRuleDestination {
			iot_GetTopicRuleDestination(cfg, client)
			return
		}
		if _iotGetV2LoggingOptions {
			iot_GetV2LoggingOptions(cfg, client)
			return
		}
		if _iotListActiveViolations {
			iot_ListActiveViolations(cfg, client)
			return
		}
		if _iotListAttachedPolicies {
			iot_ListAttachedPolicies(cfg, client)
			return
		}
		if _iotListAuditFindings {
			iot_ListAuditFindings(cfg, client)
			return
		}
		if _iotListAuditMitigationActionsExecutions {
			iot_ListAuditMitigationActionsExecutions(cfg, client)
			return
		}
		if _iotListAuditMitigationActionsTasks {
			iot_ListAuditMitigationActionsTasks(cfg, client)
			return
		}
		if _iotListAuditSuppressions {
			iot_ListAuditSuppressions(cfg, client)
			return
		}
		if _iotListAuditTasks {
			iot_ListAuditTasks(cfg, client)
			return
		}
		if _iotListAuthorizers {
			iot_ListAuthorizers(cfg, client)
			return
		}
		if _iotListBillingGroups {
			iot_ListBillingGroups(cfg, client)
			return
		}
		if _iotListCACertificates {
			iot_ListCACertificates(cfg, client)
			return
		}
		if _iotListCertificateProviders {
			iot_ListCertificateProviders(cfg, client)
			return
		}
		if _iotListCertificates {
			iot_ListCertificates(cfg, client)
			return
		}
		if _iotListCertificatesByCA {
			iot_ListCertificatesByCA(cfg, client)
			return
		}
		if _iotListCommandExecutions {
			iot_ListCommandExecutions(cfg, client)
			return
		}
		if _iotListCommands {
			iot_ListCommands(cfg, client)
			return
		}
		if _iotListCustomMetrics {
			iot_ListCustomMetrics(cfg, client)
			return
		}
		if _iotListDetectMitigationActionsExecutions {
			iot_ListDetectMitigationActionsExecutions(cfg, client)
			return
		}
		if _iotListDetectMitigationActionsTasks {
			iot_ListDetectMitigationActionsTasks(cfg, client)
			return
		}
		if _iotListDimensions {
			iot_ListDimensions(cfg, client)
			return
		}
		if _iotListDomainConfigurations {
			iot_ListDomainConfigurations(cfg, client)
			return
		}
		if _iotListFleetMetrics {
			iot_ListFleetMetrics(cfg, client)
			return
		}
		if _iotListIndices {
			iot_ListIndices(cfg, client)
			return
		}
		if _iotListJobExecutionsForJob {
			iot_ListJobExecutionsForJob(cfg, client)
			return
		}
		if _iotListJobExecutionsForThing {
			iot_ListJobExecutionsForThing(cfg, client)
			return
		}
		if _iotListJobTemplates {
			iot_ListJobTemplates(cfg, client)
			return
		}
		if _iotListJobs {
			iot_ListJobs(cfg, client)
			return
		}
		if _iotListManagedJobTemplates {
			iot_ListManagedJobTemplates(cfg, client)
			return
		}
		if _iotListMetricValues {
			iot_ListMetricValues(cfg, client)
			return
		}
		if _iotListMitigationActions {
			iot_ListMitigationActions(cfg, client)
			return
		}
		if _iotListOTAUpdates {
			iot_ListOTAUpdates(cfg, client)
			return
		}
		if _iotListOutgoingCertificates {
			iot_ListOutgoingCertificates(cfg, client)
			return
		}
		if _iotListPackageVersions {
			iot_ListPackageVersions(cfg, client)
			return
		}
		if _iotListPackages {
			iot_ListPackages(cfg, client)
			return
		}
		if _iotListPolicies {
			iot_ListPolicies(cfg, client)
			return
		}
		if _iotListPolicyPrincipals {
			iot_ListPolicyPrincipals(cfg, client)
			return
		}
		if _iotListPolicyVersions {
			iot_ListPolicyVersions(cfg, client)
			return
		}
		if _iotListPrincipalPolicies {
			iot_ListPrincipalPolicies(cfg, client)
			return
		}
		if _iotListPrincipalThings {
			iot_ListPrincipalThings(cfg, client)
			return
		}
		if _iotListPrincipalThingsV2 {
			iot_ListPrincipalThingsV2(cfg, client)
			return
		}
		if _iotListProvisioningTemplateVersions {
			iot_ListProvisioningTemplateVersions(cfg, client)
			return
		}
		if _iotListProvisioningTemplates {
			iot_ListProvisioningTemplates(cfg, client)
			return
		}
		if _iotListRelatedResourcesForAuditFinding {
			iot_ListRelatedResourcesForAuditFinding(cfg, client)
			return
		}
		if _iotListRoleAliases {
			iot_ListRoleAliases(cfg, client)
			return
		}
		if _iotListSbomValidationResults {
			iot_ListSbomValidationResults(cfg, client)
			return
		}
		if _iotListScheduledAudits {
			iot_ListScheduledAudits(cfg, client)
			return
		}
		if _iotListSecurityProfiles {
			iot_ListSecurityProfiles(cfg, client)
			return
		}
		if _iotListSecurityProfilesForTarget {
			iot_ListSecurityProfilesForTarget(cfg, client)
			return
		}
		if _iotListStreams {
			iot_ListStreams(cfg, client)
			return
		}
		if _iotListTagsForResource {
			iot_ListTagsForResource(cfg, client)
			return
		}
		if _iotListTargetsForPolicy {
			iot_ListTargetsForPolicy(cfg, client)
			return
		}
		if _iotListTargetsForSecurityProfile {
			iot_ListTargetsForSecurityProfile(cfg, client)
			return
		}
		if _iotListThingGroups {
			iot_ListThingGroups(cfg, client)
			return
		}
		if _iotListThingGroupsForThing {
			iot_ListThingGroupsForThing(cfg, client)
			return
		}
		if _iotListThingPrincipals {
			iot_ListThingPrincipals(cfg, client)
			return
		}
		if _iotListThingPrincipalsV2 {
			iot_ListThingPrincipalsV2(cfg, client)
			return
		}
		if _iotListThingRegistrationTaskReports {
			iot_ListThingRegistrationTaskReports(cfg, client)
			return
		}
		if _iotListThingRegistrationTasks {
			iot_ListThingRegistrationTasks(cfg, client)
			return
		}
		if _iotListThingTypes {
			iot_ListThingTypes(cfg, client)
			return
		}
		if _iotListThings {
			iot_ListThings(cfg, client)
			return
		}
		if _iotListThingsInBillingGroup {
			iot_ListThingsInBillingGroup(cfg, client)
			return
		}
		if _iotListThingsInThingGroup {
			iot_ListThingsInThingGroup(cfg, client)
			return
		}
		if _iotListTopicRuleDestinations {
			iot_ListTopicRuleDestinations(cfg, client)
			return
		}
		if _iotListTopicRules {
			iot_ListTopicRules(cfg, client)
			return
		}
		if _iotListV2LoggingLevels {
			iot_ListV2LoggingLevels(cfg, client)
			return
		}
		if _iotListViolationEvents {
			iot_ListViolationEvents(cfg, client)
			return
		}
		if _iotPutVerificationStateOnViolation {
			iot_PutVerificationStateOnViolation(cfg, client)
			return
		}
		if _iotRegisterCACertificate {
			iot_RegisterCACertificate(cfg, client)
			return
		}
		if _iotRegisterCertificate {
			iot_RegisterCertificate(cfg, client)
			return
		}
		if _iotRegisterCertificateWithoutCA {
			iot_RegisterCertificateWithoutCA(cfg, client)
			return
		}
		if _iotRegisterThing {
			iot_RegisterThing(cfg, client)
			return
		}
		if _iotRejectCertificateTransfer {
			iot_RejectCertificateTransfer(cfg, client)
			return
		}
		if _iotRemoveThingFromBillingGroup {
			iot_RemoveThingFromBillingGroup(cfg, client)
			return
		}
		if _iotRemoveThingFromThingGroup {
			iot_RemoveThingFromThingGroup(cfg, client)
			return
		}
		if _iotReplaceTopicRule {
			iot_ReplaceTopicRule(cfg, client)
			return
		}
		if _iotSearchIndex {
			iot_SearchIndex(cfg, client)
			return
		}
		if _iotSetDefaultAuthorizer {
			iot_SetDefaultAuthorizer(cfg, client)
			return
		}
		if _iotSetDefaultPolicyVersion {
			iot_SetDefaultPolicyVersion(cfg, client)
			return
		}
		if _iotSetLoggingOptions {
			iot_SetLoggingOptions(cfg, client)
			return
		}
		if _iotSetV2LoggingLevel {
			iot_SetV2LoggingLevel(cfg, client)
			return
		}
		if _iotSetV2LoggingOptions {
			iot_SetV2LoggingOptions(cfg, client)
			return
		}
		if _iotStartAuditMitigationActionsTask {
			iot_StartAuditMitigationActionsTask(cfg, client)
			return
		}
		if _iotStartDetectMitigationActionsTask {
			iot_StartDetectMitigationActionsTask(cfg, client)
			return
		}
		if _iotStartOnDemandAuditTask {
			iot_StartOnDemandAuditTask(cfg, client)
			return
		}
		if _iotStartThingRegistrationTask {
			iot_StartThingRegistrationTask(cfg, client)
			return
		}
		if _iotStopThingRegistrationTask {
			iot_StopThingRegistrationTask(cfg, client)
			return
		}
		if _iotTagResource {
			iot_TagResource(cfg, client)
			return
		}
		if _iotTestAuthorization {
			iot_TestAuthorization(cfg, client)
			return
		}
		if _iotTestInvokeAuthorizer {
			iot_TestInvokeAuthorizer(cfg, client)
			return
		}
		if _iotTransferCertificate {
			iot_TransferCertificate(cfg, client)
			return
		}
		if _iotUntagResource {
			iot_UntagResource(cfg, client)
			return
		}
		if _iotUpdateAccountAuditConfiguration {
			iot_UpdateAccountAuditConfiguration(cfg, client)
			return
		}
		if _iotUpdateAuditSuppression {
			iot_UpdateAuditSuppression(cfg, client)
			return
		}
		if _iotUpdateAuthorizer {
			iot_UpdateAuthorizer(cfg, client)
			return
		}
		if _iotUpdateBillingGroup {
			iot_UpdateBillingGroup(cfg, client)
			return
		}
		if _iotUpdateCACertificate {
			iot_UpdateCACertificate(cfg, client)
			return
		}
		if _iotUpdateCertificate {
			iot_UpdateCertificate(cfg, client)
			return
		}
		if _iotUpdateCertificateProvider {
			iot_UpdateCertificateProvider(cfg, client)
			return
		}
		if _iotUpdateCommand {
			iot_UpdateCommand(cfg, client)
			return
		}
		if _iotUpdateCustomMetric {
			iot_UpdateCustomMetric(cfg, client)
			return
		}
		if _iotUpdateDimension {
			iot_UpdateDimension(cfg, client)
			return
		}
		if _iotUpdateDomainConfiguration {
			iot_UpdateDomainConfiguration(cfg, client)
			return
		}
		if _iotUpdateDynamicThingGroup {
			iot_UpdateDynamicThingGroup(cfg, client)
			return
		}
		if _iotUpdateEncryptionConfiguration {
			iot_UpdateEncryptionConfiguration(cfg, client)
			return
		}
		if _iotUpdateEventConfigurations {
			iot_UpdateEventConfigurations(cfg, client)
			return
		}
		if _iotUpdateFleetMetric {
			iot_UpdateFleetMetric(cfg, client)
			return
		}
		if _iotUpdateIndexingConfiguration {
			iot_UpdateIndexingConfiguration(cfg, client)
			return
		}
		if _iotUpdateJob {
			iot_UpdateJob(cfg, client)
			return
		}
		if _iotUpdateMitigationAction {
			iot_UpdateMitigationAction(cfg, client)
			return
		}
		if _iotUpdatePackage {
			iot_UpdatePackage(cfg, client)
			return
		}
		if _iotUpdatePackageConfiguration {
			iot_UpdatePackageConfiguration(cfg, client)
			return
		}
		if _iotUpdatePackageVersion {
			iot_UpdatePackageVersion(cfg, client)
			return
		}
		if _iotUpdateProvisioningTemplate {
			iot_UpdateProvisioningTemplate(cfg, client)
			return
		}
		if _iotUpdateRoleAlias {
			iot_UpdateRoleAlias(cfg, client)
			return
		}
		if _iotUpdateScheduledAudit {
			iot_UpdateScheduledAudit(cfg, client)
			return
		}
		if _iotUpdateSecurityProfile {
			iot_UpdateSecurityProfile(cfg, client)
			return
		}
		if _iotUpdateStream {
			iot_UpdateStream(cfg, client)
			return
		}
		if _iotUpdateThing {
			iot_UpdateThing(cfg, client)
			return
		}
		if _iotUpdateThingGroup {
			iot_UpdateThingGroup(cfg, client)
			return
		}
		if _iotUpdateThingGroupsForThing {
			iot_UpdateThingGroupsForThing(cfg, client)
			return
		}
		if _iotUpdateThingType {
			iot_UpdateThingType(cfg, client)
			return
		}
		if _iotUpdateTopicRuleDestination {
			iot_UpdateTopicRuleDestination(cfg, client)
			return
		}
		if _iotValidateSecurityProfileBehaviors {
			iot_ValidateSecurityProfileBehaviors(cfg, client)
			return
		}

	},
}

var (
	_iotAcceptCertificateTransfer             bool
	_iotAddThingToBillingGroup                bool
	_iotAddThingToThingGroup                  bool
	_iotAssociateSbomWithPackageVersion       bool
	_iotAssociateTargetsWithJob               bool
	_iotAttachPolicy                          bool
	_iotAttachPrincipalPolicy                 bool
	_iotAttachSecurityProfile                 bool
	_iotAttachThingPrincipal                  bool
	_iotCancelAuditMitigationActionsTask      bool
	_iotCancelAuditTask                       bool
	_iotCancelCertificateTransfer             bool
	_iotCancelDetectMitigationActionsTask     bool
	_iotCancelJob                             bool
	_iotCancelJobExecution                    bool
	_iotClearDefaultAuthorizer                bool
	_iotConfirmTopicRuleDestination           bool
	_iotCreateAuditSuppression                bool
	_iotCreateAuthorizer                      bool
	_iotCreateBillingGroup                    bool
	_iotCreateCertificateFromCsr              bool
	_iotCreateCertificateProvider             bool
	_iotCreateCommand                         bool
	_iotCreateCustomMetric                    bool
	_iotCreateDimension                       bool
	_iotCreateDomainConfiguration             bool
	_iotCreateDynamicThingGroup               bool
	_iotCreateFleetMetric                     bool
	_iotCreateJob                             bool
	_iotCreateJobTemplate                     bool
	_iotCreateKeysAndCertificate              bool
	_iotCreateMitigationAction                bool
	_iotCreateOTAUpdate                       bool
	_iotCreatePackage                         bool
	_iotCreatePackageVersion                  bool
	_iotCreatePolicy                          bool
	_iotCreatePolicyVersion                   bool
	_iotCreateProvisioningClaim               bool
	_iotCreateProvisioningTemplate            bool
	_iotCreateProvisioningTemplateVersion     bool
	_iotCreateRoleAlias                       bool
	_iotCreateScheduledAudit                  bool
	_iotCreateSecurityProfile                 bool
	_iotCreateStream                          bool
	_iotCreateThing                           bool
	_iotCreateThingGroup                      bool
	_iotCreateThingType                       bool
	_iotCreateTopicRule                       bool
	_iotCreateTopicRuleDestination            bool
	_iotDeleteAccountAuditConfiguration       bool
	_iotDeleteAuditSuppression                bool
	_iotDeleteAuthorizer                      bool
	_iotDeleteBillingGroup                    bool
	_iotDeleteCACertificate                   bool
	_iotDeleteCertificate                     bool
	_iotDeleteCertificateProvider             bool
	_iotDeleteCommand                         bool
	_iotDeleteCommandExecution                bool
	_iotDeleteCustomMetric                    bool
	_iotDeleteDimension                       bool
	_iotDeleteDomainConfiguration             bool
	_iotDeleteDynamicThingGroup               bool
	_iotDeleteFleetMetric                     bool
	_iotDeleteJob                             bool
	_iotDeleteJobExecution                    bool
	_iotDeleteJobTemplate                     bool
	_iotDeleteMitigationAction                bool
	_iotDeleteOTAUpdate                       bool
	_iotDeletePackage                         bool
	_iotDeletePackageVersion                  bool
	_iotDeletePolicy                          bool
	_iotDeletePolicyVersion                   bool
	_iotDeleteProvisioningTemplate            bool
	_iotDeleteProvisioningTemplateVersion     bool
	_iotDeleteRegistrationCode                bool
	_iotDeleteRoleAlias                       bool
	_iotDeleteScheduledAudit                  bool
	_iotDeleteSecurityProfile                 bool
	_iotDeleteStream                          bool
	_iotDeleteThing                           bool
	_iotDeleteThingGroup                      bool
	_iotDeleteThingType                       bool
	_iotDeleteTopicRule                       bool
	_iotDeleteTopicRuleDestination            bool
	_iotDeleteV2LoggingLevel                  bool
	_iotDeprecateThingType                    bool
	_iotDescribeAccountAuditConfiguration     bool
	_iotDescribeAuditFinding                  bool
	_iotDescribeAuditMitigationActionsTask    bool
	_iotDescribeAuditSuppression              bool
	_iotDescribeAuditTask                     bool
	_iotDescribeAuthorizer                    bool
	_iotDescribeBillingGroup                  bool
	_iotDescribeCACertificate                 bool
	_iotDescribeCertificate                   bool
	_iotDescribeCertificateProvider           bool
	_iotDescribeCustomMetric                  bool
	_iotDescribeDefaultAuthorizer             bool
	_iotDescribeDetectMitigationActionsTask   bool
	_iotDescribeDimension                     bool
	_iotDescribeDomainConfiguration           bool
	_iotDescribeEncryptionConfiguration       bool
	_iotDescribeEndpoint                      bool
	_iotDescribeEventConfigurations           bool
	_iotDescribeFleetMetric                   bool
	_iotDescribeIndex                         bool
	_iotDescribeJob                           bool
	_iotDescribeJobExecution                  bool
	_iotDescribeJobTemplate                   bool
	_iotDescribeManagedJobTemplate            bool
	_iotDescribeMitigationAction              bool
	_iotDescribeProvisioningTemplate          bool
	_iotDescribeProvisioningTemplateVersion   bool
	_iotDescribeRoleAlias                     bool
	_iotDescribeScheduledAudit                bool
	_iotDescribeSecurityProfile               bool
	_iotDescribeStream                        bool
	_iotDescribeThing                         bool
	_iotDescribeThingGroup                    bool
	_iotDescribeThingRegistrationTask         bool
	_iotDescribeThingType                     bool
	_iotDetachPolicy                          bool
	_iotDetachPrincipalPolicy                 bool
	_iotDetachSecurityProfile                 bool
	_iotDetachThingPrincipal                  bool
	_iotDisableTopicRule                      bool
	_iotDisassociateSbomFromPackageVersion    bool
	_iotEnableTopicRule                       bool
	_iotGetBehaviorModelTrainingSummaries     bool
	_iotGetBucketsAggregation                 bool
	_iotGetCardinality                        bool
	_iotGetCommand                            bool
	_iotGetCommandExecution                   bool
	_iotGetEffectivePolicies                  bool
	_iotGetIndexingConfiguration              bool
	_iotGetJobDocument                        bool
	_iotGetLoggingOptions                     bool
	_iotGetOTAUpdate                          bool
	_iotGetPackage                            bool
	_iotGetPackageConfiguration               bool
	_iotGetPackageVersion                     bool
	_iotGetPercentiles                        bool
	_iotGetPolicy                             bool
	_iotGetPolicyVersion                      bool
	_iotGetRegistrationCode                   bool
	_iotGetStatistics                         bool
	_iotGetThingConnectivityData              bool
	_iotGetTopicRule                          bool
	_iotGetTopicRuleDestination               bool
	_iotGetV2LoggingOptions                   bool
	_iotListActiveViolations                  bool
	_iotListAttachedPolicies                  bool
	_iotListAuditFindings                     bool
	_iotListAuditMitigationActionsExecutions  bool
	_iotListAuditMitigationActionsTasks       bool
	_iotListAuditSuppressions                 bool
	_iotListAuditTasks                        bool
	_iotListAuthorizers                       bool
	_iotListBillingGroups                     bool
	_iotListCACertificates                    bool
	_iotListCertificateProviders              bool
	_iotListCertificates                      bool
	_iotListCertificatesByCA                  bool
	_iotListCommandExecutions                 bool
	_iotListCommands                          bool
	_iotListCustomMetrics                     bool
	_iotListDetectMitigationActionsExecutions bool
	_iotListDetectMitigationActionsTasks      bool
	_iotListDimensions                        bool
	_iotListDomainConfigurations              bool
	_iotListFleetMetrics                      bool
	_iotListIndices                           bool
	_iotListJobExecutionsForJob               bool
	_iotListJobExecutionsForThing             bool
	_iotListJobTemplates                      bool
	_iotListJobs                              bool
	_iotListManagedJobTemplates               bool
	_iotListMetricValues                      bool
	_iotListMitigationActions                 bool
	_iotListOTAUpdates                        bool
	_iotListOutgoingCertificates              bool
	_iotListPackageVersions                   bool
	_iotListPackages                          bool
	_iotListPolicies                          bool
	_iotListPolicyPrincipals                  bool
	_iotListPolicyVersions                    bool
	_iotListPrincipalPolicies                 bool
	_iotListPrincipalThings                   bool
	_iotListPrincipalThingsV2                 bool
	_iotListProvisioningTemplateVersions      bool
	_iotListProvisioningTemplates             bool
	_iotListRelatedResourcesForAuditFinding   bool
	_iotListRoleAliases                       bool
	_iotListSbomValidationResults             bool
	_iotListScheduledAudits                   bool
	_iotListSecurityProfiles                  bool
	_iotListSecurityProfilesForTarget         bool
	_iotListStreams                           bool
	_iotListTagsForResource                   bool
	_iotListTargetsForPolicy                  bool
	_iotListTargetsForSecurityProfile         bool
	_iotListThingGroups                       bool
	_iotListThingGroupsForThing               bool
	_iotListThingPrincipals                   bool
	_iotListThingPrincipalsV2                 bool
	_iotListThingRegistrationTaskReports      bool
	_iotListThingRegistrationTasks            bool
	_iotListThingTypes                        bool
	_iotListThings                            bool
	_iotListThingsInBillingGroup              bool
	_iotListThingsInThingGroup                bool
	_iotListTopicRuleDestinations             bool
	_iotListTopicRules                        bool
	_iotListV2LoggingLevels                   bool
	_iotListViolationEvents                   bool
	_iotPutVerificationStateOnViolation       bool
	_iotRegisterCACertificate                 bool
	_iotRegisterCertificate                   bool
	_iotRegisterCertificateWithoutCA          bool
	_iotRegisterThing                         bool
	_iotRejectCertificateTransfer             bool
	_iotRemoveThingFromBillingGroup           bool
	_iotRemoveThingFromThingGroup             bool
	_iotReplaceTopicRule                      bool
	_iotSearchIndex                           bool
	_iotSetDefaultAuthorizer                  bool
	_iotSetDefaultPolicyVersion               bool
	_iotSetLoggingOptions                     bool
	_iotSetV2LoggingLevel                     bool
	_iotSetV2LoggingOptions                   bool
	_iotStartAuditMitigationActionsTask       bool
	_iotStartDetectMitigationActionsTask      bool
	_iotStartOnDemandAuditTask                bool
	_iotStartThingRegistrationTask            bool
	_iotStopThingRegistrationTask             bool
	_iotTagResource                           bool
	_iotTestAuthorization                     bool
	_iotTestInvokeAuthorizer                  bool
	_iotTransferCertificate                   bool
	_iotUntagResource                         bool
	_iotUpdateAccountAuditConfiguration       bool
	_iotUpdateAuditSuppression                bool
	_iotUpdateAuthorizer                      bool
	_iotUpdateBillingGroup                    bool
	_iotUpdateCACertificate                   bool
	_iotUpdateCertificate                     bool
	_iotUpdateCertificateProvider             bool
	_iotUpdateCommand                         bool
	_iotUpdateCustomMetric                    bool
	_iotUpdateDimension                       bool
	_iotUpdateDomainConfiguration             bool
	_iotUpdateDynamicThingGroup               bool
	_iotUpdateEncryptionConfiguration         bool
	_iotUpdateEventConfigurations             bool
	_iotUpdateFleetMetric                     bool
	_iotUpdateIndexingConfiguration           bool
	_iotUpdateJob                             bool
	_iotUpdateMitigationAction                bool
	_iotUpdatePackage                         bool
	_iotUpdatePackageConfiguration            bool
	_iotUpdatePackageVersion                  bool
	_iotUpdateProvisioningTemplate            bool
	_iotUpdateRoleAlias                       bool
	_iotUpdateScheduledAudit                  bool
	_iotUpdateSecurityProfile                 bool
	_iotUpdateStream                          bool
	_iotUpdateThing                           bool
	_iotUpdateThingGroup                      bool
	_iotUpdateThingGroupsForThing             bool
	_iotUpdateThingType                       bool
	_iotUpdateTopicRuleDestination            bool
	_iotValidateSecurityProfileBehaviors      bool

	_iotAbortConfig                           string
	_iotAccountDefaultForOperations           string
	_iotAction                                string
	_iotActionName                            string
	_iotActionParams                          string
	_iotActionStatus                          string
	_iotActionType                            string
	_iotActions                               []string
	_iotAdditionalMetricsToRetain             []string
	_iotAdditionalMetricsToRetainV2           string
	_iotAdditionalParameters                  string
	_iotAggregationField                      string
	_iotAggregationType                       string
	_iotAlertTargets                          string
	_iotAllowAutoRegistration                 string
	_iotApplicationProtocol                   string
	_iotArn                                   string
	_iotArtifact                              string
	_iotAscendingOrder                        string
	_iotAttributeName                         string
	_iotAttributePayload                      string
	_iotAttributeValue                        string
	_iotAttributes                            string
	_iotAuditCheckConfigurations              string
	_iotAuditCheckToActionsMapping            string
	_iotAuditNotificationTargetConfigurations string
	_iotAuditTaskId                           string
	_iotAuthInfos                             string
	_iotAuthenticationType                    string
	_iotAuthorizerConfig                      string
	_iotAuthorizerFunctionArn                 string
	_iotAuthorizerName                        string
	_iotAwsJobAbortConfig                     string
	_iotAwsJobExecutionsRolloutConfig         string
	_iotAwsJobPresignedUrlConfig              string
	_iotAwsJobTimeoutConfig                   string
	_iotBeforeSubstitution                    string
	_iotBehaviorCriteriaType                  string
	_iotBehaviors                             string
	_iotBillingGroupArn                       string
	_iotBillingGroupName                      string
	_iotBillingGroupProperties                string
	_iotBucketsAggregationType                string
	_iotCaCertificate                         string
	_iotCaCertificateId                       string
	_iotCaCertificatePem                      string
	_iotCertificateId                         string
	_iotCertificateMode                       string
	_iotCertificatePem                        string
	_iotCertificateProviderName               string
	_iotCertificateSigningRequest             string
	_iotCheckName                             string
	_iotClientCertificateConfig               string
	_iotClientId                              string
	_iotClientRequestToken                    string
	_iotClientToken                           string
	_iotCognitoIdentityPoolId                 string
	_iotCommandArn                            string
	_iotCommandId                             string
	_iotCommandParameterName                  string
	_iotComment                               string
	_iotCompletedTimeFilter                   string
	_iotConfirmationToken                     string
	_iotCredentialDurationSeconds             string
	_iotDayOfMonth                            string
	_iotDayOfWeek                             string
	_iotDefaultLogLevel                       string
	_iotDefaultVersionId                      string
	_iotDefaultVersionName                    string
	_iotDeleteAdditionalMetricsToRetain       string
	_iotDeleteAlertTargets                    string
	_iotDeleteBehaviors                       string
	_iotDeleteMetricsExportConfig             string
	_iotDeleteScheduledAudits                 string
	_iotDeprecated                            string
	_iotDescription                           string
	_iotDestinationConfiguration              string
	_iotDestinationPackageVersions            []string
	_iotDimensionName                         string
	_iotDimensionValueOperator                string
	_iotDisableAllLogs                        string
	_iotDisplayName                           string
	_iotDocument                              string
	_iotDocumentParameters                    string
	_iotDocumentSource                        string
	_iotDomainConfigurationName               string
	_iotDomainConfigurationStatus             string
	_iotDomainName                            string
	_iotEnableCachingForHttp                  string
	_iotEnabled                               string
	_iotEncryptionType                        string
	_iotEndTime                               string
	_iotEndpointType                          string
	_iotEventConfigurations                   string
	_iotExecutionId                           string
	_iotExecutionNumber                       string
	_iotExpectedVersion                       string
	_iotExpirationDate                        string
	_iotFiles                                 string
	_iotFindingId                             string
	_iotForce                                 string
	_iotForceDelete                           string
	_iotForceDeleteAWSJob                     string
	_iotFrequency                             string
	_iotHttpContext                           string
	_iotIncludeOnlyActiveViolations           string
	_iotIncludeResult                         string
	_iotIncludeSuppressedAlerts               string
	_iotIndexName                             string
	_iotInputFileBucket                       string
	_iotInputFileKey                          string
	_iotJobArn                                string
	_iotJobExecutionsRetryConfig              string
	_iotJobExecutionsRolloutConfig            string
	_iotJobId                                 string
	_iotJobTemplateArn                        string
	_iotJobTemplateId                         string
	_iotKmsAccessRoleArn                      string
	_iotKmsKeyArn                             string
	_iotLambdaFunctionArn                     string
	_iotListSuppressedAlerts                  string
	_iotListSuppressedFindings                string
	_iotLogLevel                              string
	_iotLogTarget                             string
	_iotLoggingOptionsPayload                 string
	_iotMaintenanceWindows                    string
	_iotMandatoryParameters                   string
	_iotMarker                                string
	_iotMaxResults                            string
	_iotMetricName                            string
	_iotMetricType                            string
	_iotMetricsExportConfig                   string
	_iotMqttContext                           string
	_iotName                                  string
	_iotNamePrefixFilter                      string
	_iotNamespace                             string
	_iotNamespaceId                           string
	_iotNewAutoRegistrationStatus             string
	_iotNewStatus                             string
	_iotNextToken                             string
	_iotOtaUpdateId                           string
	_iotOtaUpdateStatus                       string
	_iotOverrideDynamicGroups                 string
	_iotPackageName                           string
	_iotPageSize                              string
	_iotParameters                            string
	_iotParentGroup                           string
	_iotParentGroupName                       string
	_iotPayload                               string
	_iotPayloadTemplate                       string
	_iotPercents                              string
	_iotPeriod                                string
	_iotPolicyDocument                        string
	_iotPolicyName                            string
	_iotPolicyNamesToAdd                      []string
	_iotPolicyNamesToSkip                     []string
	_iotPolicyVersionId                       string
	_iotPreProvisioningHook                   string
	_iotPreprocessor                          string
	_iotPresignedUrlConfig                    string
	_iotPrincipal                             string
	_iotProtocols                             string
	_iotProvisioningRoleArn                   string
	_iotQueryString                           string
	_iotQueryVersion                          string
	_iotReasonCode                            string
	_iotRecipe                                string
	_iotRecursive                             string
	_iotRegistrationConfig                    string
	_iotRejectReason                          string
	_iotRemoveAuthorizerConfig                string
	_iotRemoveAutoRegistration                string
	_iotRemovePreProvisioningHook             string
	_iotRemoveThingType                       string
	_iotReportType                            string
	_iotResourceArn                           string
	_iotResourceIdentifier                    string
	_iotRoleAlias                             string
	_iotRoleArn                               string
	_iotRuleDisabled                          string
	_iotRuleName                              string
	_iotSbom                                  string
	_iotScheduledAuditName                    string
	_iotSchedulingConfig                      string
	_iotSecurityProfileDescription            string
	_iotSecurityProfileName                   string
	_iotSecurityProfileTargetArn              string
	_iotServerCertificateArns                 []string
	_iotServerCertificateConfig               string
	_iotServiceType                           string
	_iotSetAsActive                           string
	_iotSetAsDefault                          string
	_iotSigningDisabled                       string
	_iotSortOrder                             string
	_iotStartTime                             string
	_iotStartedTimeFilter                     string
	_iotStatus                                string
	_iotStatusDetails                         string
	_iotStreamId                              string
	_iotStringValues                          []string
	_iotSuppressIndefinitely                  string
	_iotTagKeys                               []string
	_iotTags                                  string
	_iotTarget                                string
	_iotTargetArn                             string
	_iotTargetAwsAccount                      string
	_iotTargetCheckNames                      []string
	_iotTargetName                            string
	_iotTargetSelection                       string
	_iotTargetType                            string
	_iotTargets                               []string
	_iotTaskId                                string
	_iotTaskStatus                            string
	_iotTaskType                              string
	_iotTemplateBody                          string
	_iotTemplateName                          string
	_iotTemplateVersion                       string
	_iotThingArn                              string
	_iotThingGroupArn                         string
	_iotThingGroupId                          string
	_iotThingGroupIndexingConfiguration       string
	_iotThingGroupName                        string
	_iotThingGroupProperties                  string
	_iotThingGroupsToAdd                      []string
	_iotThingGroupsToRemove                   []string
	_iotThingIndexingConfiguration            string
	_iotThingName                             string
	_iotThingPrincipalType                    string
	_iotThingTypeName                         string
	_iotThingTypeProperties                   string
	_iotTimeoutConfig                         string
	_iotTlsConfig                             string
	_iotTlsContext                            string
	_iotToken                                 string
	_iotTokenKeyName                          string
	_iotTokenSignature                        string
	_iotTokenSigningPublicKeys                string
	_iotTopic                                 string
	_iotTopicRulePayload                      string
	_iotTransferMessage                       string
	_iotType                                  string
	_iotUndoDeprecate                         string
	_iotUnit                                  string
	_iotUnsetDefaultVersion                   string
	_iotUsePrefixAttributeValue               string
	_iotValidationCertificateArn              string
	_iotValidationResult                      string
	_iotVerbose                               string
	_iotVerificationCertificate               string
	_iotVerificationState                     string
	_iotVerificationStateDescription          string
	_iotVersionId                             string
	_iotVersionName                           string
	_iotVersionUpdateByJobsConfig             string
	_iotViolationEventOccurrenceRange         string
	_iotViolationId                           string
)

// Accepts a pending certificate transfer. The default state of the certificate is
// INACTIVE.
//
// To check for pending certificate transfers, call ListCertificates to enumerate your
// certificates.
//
// Requires permission to access the [AcceptCertificateTransfer] action.
//
// [AcceptCertificateTransfer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_AcceptCertificateTransfer(cfg aws.Config, client *iot.Client) {
	input := &iot.AcceptCertificateTransferInput{
		// CertificateId: *string, // Required
	}

	if len(_iotCertificateId) > 0 {
		input.CertificateId = aws.String(_iotCertificateId)
	}
	if len(_iotSetAsActive) > 0 {
		if err := assignInputField(input, "SetAsActive", _iotSetAsActive); err != nil {
			log.Errorf("invalid --set-as-active: %s", err.Error())
			return
		}
	}

	if resp, err := client.AcceptCertificateTransfer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a thing to a billing group.
// Requires permission to access the [AddThingToBillingGroup] action.
//
// [AddThingToBillingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_AddThingToBillingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.AddThingToBillingGroupInput{}

	if len(_iotBillingGroupArn) > 0 {
		input.BillingGroupArn = aws.String(_iotBillingGroupArn)
	}
	if len(_iotBillingGroupName) > 0 {
		input.BillingGroupName = aws.String(_iotBillingGroupName)
	}
	if len(_iotThingArn) > 0 {
		input.ThingArn = aws.String(_iotThingArn)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}

	if resp, err := client.AddThingToBillingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a thing to a thing group.
// Requires permission to access the [AddThingToThingGroup] action.
//
// [AddThingToThingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_AddThingToThingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.AddThingToThingGroupInput{}

	if len(_iotOverrideDynamicGroups) > 0 {
		if err := assignInputField(input, "OverrideDynamicGroups", _iotOverrideDynamicGroups); err != nil {
			log.Errorf("invalid --override-dynamic-groups: %s", err.Error())
			return
		}
	}
	if len(_iotThingArn) > 0 {
		input.ThingArn = aws.String(_iotThingArn)
	}
	if len(_iotThingGroupArn) > 0 {
		input.ThingGroupArn = aws.String(_iotThingGroupArn)
	}
	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}

	if resp, err := client.AddThingToThingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the selected software bill of materials (SBOM) with a specific
// software package version.
//
// Requires permission to access the [AssociateSbomWithPackageVersion] action.
//
// [AssociateSbomWithPackageVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_AssociateSbomWithPackageVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.AssociateSbomWithPackageVersionInput{
		// PackageName: *string, // Required
		// Sbom: *types.Sbom, // Required
		// VersionName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotSbom) > 0 {
		if err := assignInputField(input, "Sbom", _iotSbom); err != nil {
			log.Errorf("invalid --sbom: %s", err.Error())
			return
		}
	}
	if len(_iotVersionName) > 0 {
		input.VersionName = aws.String(_iotVersionName)
	}
	if len(_iotClientToken) > 0 {
		input.ClientToken = aws.String(_iotClientToken)
	}

	if resp, err := client.AssociateSbomWithPackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a group with a continuous job. The following criteria must be met:
// - The job must have been created with the targetSelection field set to
// "CONTINUOUS".
//
// - The job status must currently be "IN_PROGRESS".
//
// - The total number of targets associated with a job must not exceed 100.
//
// Requires permission to access the [AssociateTargetsWithJob] action.
//
// [AssociateTargetsWithJob]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_AssociateTargetsWithJob(cfg aws.Config, client *iot.Client) {
	input := &iot.AssociateTargetsWithJobInput{
		// JobId: *string, // Required
		// Targets: []string, // Required
	}

	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotTargets) > 0 {
		input.Targets = append([]string(nil), _iotTargets...)
	}
	if len(_iotComment) > 0 {
		input.Comment = aws.String(_iotComment)
	}
	if len(_iotNamespaceId) > 0 {
		input.NamespaceId = aws.String(_iotNamespaceId)
	}

	if resp, err := client.AssociateTargetsWithJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches the specified policy to the specified principal (certificate or other
// credential).
//
// Requires permission to access the [AttachPolicy] action.
//
// [AttachPolicy]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_AttachPolicy(cfg aws.Config, client *iot.Client) {
	input := &iot.AttachPolicyInput{
		// PolicyName: *string, // Required
		// Target: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotTarget) > 0 {
		input.Target = aws.String(_iotTarget)
	}

	if resp, err := client.AttachPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches the specified policy to the specified principal (certificate or other
// credential).
//
// Note: This action is deprecated and works as expected for backward
// compatibility, but we won't add enhancements. Use AttachPolicyinstead.
//
// Requires permission to access the [AttachPrincipalPolicy] action.
//
// Deprecated: This operation has been deprecated.
//
// [AttachPrincipalPolicy]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_AttachPrincipalPolicy(cfg aws.Config, client *iot.Client) {
	input := &iot.AttachPrincipalPolicyInput{
		// PolicyName: *string, // Required
		// Principal: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotPrincipal) > 0 {
		input.Principal = aws.String(_iotPrincipal)
	}

	if resp, err := client.AttachPrincipalPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a Device Defender security profile with a thing group or this
// account. Each thing group or account can have up to five security profiles
// associated with it.
//
// Requires permission to access the [AttachSecurityProfile] action.
//
// [AttachSecurityProfile]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_AttachSecurityProfile(cfg aws.Config, client *iot.Client) {
	input := &iot.AttachSecurityProfileInput{
		// SecurityProfileName: *string, // Required
		// SecurityProfileTargetArn: *string, // Required
	}

	if len(_iotSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_iotSecurityProfileName)
	}
	if len(_iotSecurityProfileTargetArn) > 0 {
		input.SecurityProfileTargetArn = aws.String(_iotSecurityProfileTargetArn)
	}

	if resp, err := client.AttachSecurityProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches the specified principal to the specified thing. A principal can be
// X.509 certificates, Amazon Cognito identities or federated identities.
//
// Requires permission to access the [AttachThingPrincipal] action.
//
// [AttachThingPrincipal]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_AttachThingPrincipal(cfg aws.Config, client *iot.Client) {
	input := &iot.AttachThingPrincipalInput{
		// Principal: *string, // Required
		// ThingName: *string, // Required
	}

	if len(_iotPrincipal) > 0 {
		input.Principal = aws.String(_iotPrincipal)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotThingPrincipalType) > 0 {
		if err := assignInputField(input, "ThingPrincipalType", _iotThingPrincipalType); err != nil {
			log.Errorf("invalid --thing-principal-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.AttachThingPrincipal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a mitigation action task that is in progress. If the task is not in
// progress, an InvalidRequestException occurs.
//
// Requires permission to access the [CancelAuditMitigationActionsTask] action.
//
// [CancelAuditMitigationActionsTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CancelAuditMitigationActionsTask(cfg aws.Config, client *iot.Client) {
	input := &iot.CancelAuditMitigationActionsTaskInput{
		// TaskId: *string, // Required
	}

	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}

	if resp, err := client.CancelAuditMitigationActionsTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an audit that is in progress. The audit can be either scheduled or on
// demand. If the audit isn't in progress, an "InvalidRequestException" occurs.
//
// Requires permission to access the [CancelAuditTask] action.
//
// [CancelAuditTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CancelAuditTask(cfg aws.Config, client *iot.Client) {
	input := &iot.CancelAuditTaskInput{
		// TaskId: *string, // Required
	}

	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}

	if resp, err := client.CancelAuditTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a pending transfer for the specified certificate.
// Note Only the transfer source account can use this operation to cancel a
// transfer. (Transfer destinations can use RejectCertificateTransferinstead.) After transfer, IoT returns
// the certificate to the source account in the INACTIVE state. After the
// destination account has accepted the transfer, the transfer cannot be cancelled.
//
// After a certificate transfer is cancelled, the status of the certificate
// changes from PENDING_TRANSFER to INACTIVE.
//
// Requires permission to access the [CancelCertificateTransfer] action.
//
// [CancelCertificateTransfer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CancelCertificateTransfer(cfg aws.Config, client *iot.Client) {
	input := &iot.CancelCertificateTransferInput{
		// CertificateId: *string, // Required
	}

	if len(_iotCertificateId) > 0 {
		input.CertificateId = aws.String(_iotCertificateId)
	}

	if resp, err := client.CancelCertificateTransfer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a Device Defender ML Detect mitigation action.
// Requires permission to access the [CancelDetectMitigationActionsTask] action.
//
// [CancelDetectMitigationActionsTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CancelDetectMitigationActionsTask(cfg aws.Config, client *iot.Client) {
	input := &iot.CancelDetectMitigationActionsTaskInput{
		// TaskId: *string, // Required
	}

	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}

	if resp, err := client.CancelDetectMitigationActionsTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a job.
// Requires permission to access the [CancelJob] action.
//
// [CancelJob]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CancelJob(cfg aws.Config, client *iot.Client) {
	input := &iot.CancelJobInput{
		// JobId: *string, // Required
	}

	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotComment) > 0 {
		input.Comment = aws.String(_iotComment)
	}
	if len(_iotForce) > 0 {
		if err := assignInputField(input, "Force", _iotForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_iotReasonCode) > 0 {
		input.ReasonCode = aws.String(_iotReasonCode)
	}

	if resp, err := client.CancelJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the execution of a job for a given thing.
// Requires permission to access the [CancelJobExecution] action.
//
// [CancelJobExecution]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CancelJobExecution(cfg aws.Config, client *iot.Client) {
	input := &iot.CancelJobExecutionInput{
		// JobId: *string, // Required
		// ThingName: *string, // Required
	}

	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}
	if len(_iotForce) > 0 {
		if err := assignInputField(input, "Force", _iotForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_iotStatusDetails) > 0 {
		if err := assignInputField(input, "StatusDetails", _iotStatusDetails); err != nil {
			log.Errorf("invalid --status-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.CancelJobExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Clears the default authorizer.
// Requires permission to access the [ClearDefaultAuthorizer] action.
//
// [ClearDefaultAuthorizer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ClearDefaultAuthorizer(cfg aws.Config, client *iot.Client) {
	input := &iot.ClearDefaultAuthorizerInput{}

	if resp, err := client.ClearDefaultAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Confirms a topic rule destination. When you create a rule requiring a
// destination, IoT sends a confirmation message to the endpoint or base address
// you specify. The message includes a token which you pass back when calling
// ConfirmTopicRuleDestination to confirm that you own or have access to the
// endpoint.
//
// Requires permission to access the [ConfirmTopicRuleDestination] action.
//
// [ConfirmTopicRuleDestination]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ConfirmTopicRuleDestination(cfg aws.Config, client *iot.Client) {
	input := &iot.ConfirmTopicRuleDestinationInput{
		// ConfirmationToken: *string, // Required
	}

	if len(_iotConfirmationToken) > 0 {
		input.ConfirmationToken = aws.String(_iotConfirmationToken)
	}

	if resp, err := client.ConfirmTopicRuleDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Device Defender audit suppression.
// Requires permission to access the [CreateAuditSuppression] action.
//
// [CreateAuditSuppression]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateAuditSuppression(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateAuditSuppressionInput{
		// CheckName: *string, // Required
		// ClientRequestToken: *string, // Required
		// ResourceIdentifier: *types.ResourceIdentifier, // Required
	}

	if len(_iotCheckName) > 0 {
		input.CheckName = aws.String(_iotCheckName)
	}
	if len(_iotClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotClientRequestToken)
	}
	if len(_iotResourceIdentifier) > 0 {
		if err := assignInputField(input, "ResourceIdentifier", _iotResourceIdentifier); err != nil {
			log.Errorf("invalid --resource-identifier: %s", err.Error())
			return
		}
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotExpirationDate) > 0 {
		if err := assignInputField(input, "ExpirationDate", _iotExpirationDate); err != nil {
			log.Errorf("invalid --expiration-date: %s", err.Error())
			return
		}
	}
	if len(_iotSuppressIndefinitely) > 0 {
		if err := assignInputField(input, "SuppressIndefinitely", _iotSuppressIndefinitely); err != nil {
			log.Errorf("invalid --suppress-indefinitely: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAuditSuppression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an authorizer.
// Requires permission to access the [CreateAuthorizer] action.
//
// [CreateAuthorizer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateAuthorizer(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateAuthorizerInput{
		// AuthorizerFunctionArn: *string, // Required
		// AuthorizerName: *string, // Required
	}

	if len(_iotAuthorizerFunctionArn) > 0 {
		input.AuthorizerFunctionArn = aws.String(_iotAuthorizerFunctionArn)
	}
	if len(_iotAuthorizerName) > 0 {
		input.AuthorizerName = aws.String(_iotAuthorizerName)
	}
	if len(_iotEnableCachingForHttp) > 0 {
		if err := assignInputField(input, "EnableCachingForHttp", _iotEnableCachingForHttp); err != nil {
			log.Errorf("invalid --enable-caching-for-http: %s", err.Error())
			return
		}
	}
	if len(_iotSigningDisabled) > 0 {
		if err := assignInputField(input, "SigningDisabled", _iotSigningDisabled); err != nil {
			log.Errorf("invalid --signing-disabled: %s", err.Error())
			return
		}
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotTokenKeyName) > 0 {
		input.TokenKeyName = aws.String(_iotTokenKeyName)
	}
	if len(_iotTokenSigningPublicKeys) > 0 {
		if err := assignInputField(input, "TokenSigningPublicKeys", _iotTokenSigningPublicKeys); err != nil {
			log.Errorf("invalid --token-signing-public-keys: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a billing group. If this call is made multiple times using the same
// billing group name and configuration, the call will succeed. If this call is
// made with the same billing group name but different configuration a
// ResourceAlreadyExistsException is thrown.
//
// Requires permission to access the [CreateBillingGroup] action.
//
// [CreateBillingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateBillingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateBillingGroupInput{
		// BillingGroupName: *string, // Required
	}

	if len(_iotBillingGroupName) > 0 {
		input.BillingGroupName = aws.String(_iotBillingGroupName)
	}
	if len(_iotBillingGroupProperties) > 0 {
		if err := assignInputField(input, "BillingGroupProperties", _iotBillingGroupProperties); err != nil {
			log.Errorf("invalid --billing-group-properties: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBillingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an X.509 certificate using the specified certificate signing request.
// Requires permission to access the [CreateCertificateFromCsr] action.
//
// The CSR must include a public key that is either an RSA key with a length of at
// least 2048 bits or an ECC key from NIST P-256, NIST P-384, or NIST P-521 curves.
// For supported certificates, consult [Certificate signing algorithms supported by IoT].
//
// Reusing the same certificate signing request (CSR) results in a distinct
// certificate.
//
// You can create multiple certificates in a batch by creating a directory,
// copying multiple .csr files into that directory, and then specifying that
// directory on the command line. The following commands show how to create a batch
// of certificates given a batch of CSRs. In the following commands, we assume that
// a set of CSRs are located inside of the directory my-csr-directory:
//
// On Linux and OS X, the command is:
//
// $ ls my-csr-directory/ | xargs -I {} aws iot create-certificate-from-csr
// --certificate-signing-request file://my-csr-directory/{}
//
// This command lists all of the CSRs in my-csr-directory and pipes each CSR file
// name to the aws iot create-certificate-from-csr Amazon Web Services CLI command
// to create a certificate for the corresponding CSR.
//
// You can also run the aws iot create-certificate-from-csr part of the command in
// parallel to speed up the certificate creation process:
//
// $ ls my-csr-directory/ | xargs -P 10 -I {} aws iot create-certificate-from-csr
// --certificate-signing-request file://my-csr-directory/{}
//
// On Windows PowerShell, the command to create certificates for all CSRs in
// my-csr-directory is:
//
// > ls -Name my-csr-directory | %{aws iot create-certificate-from-csr
// --certificate-signing-request file://my-csr-directory/$_}
//
// On a Windows command prompt, the command to create certificates for all CSRs in
// my-csr-directory is:
//
// > forfiles /p my-csr-directory /c "cmd /c aws iot create-certificate-from-csr
// --certificate-signing-request file://(at)path"
//
// [CreateCertificateFromCsr]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [Certificate signing algorithms supported by IoT]: https://docs.aws.amazon.com/iot/latest/developerguide/x509-client-certs.html#x509-cert-algorithms
func iot_CreateCertificateFromCsr(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateCertificateFromCsrInput{
		// CertificateSigningRequest: *string, // Required
	}

	if len(_iotCertificateSigningRequest) > 0 {
		input.CertificateSigningRequest = aws.String(_iotCertificateSigningRequest)
	}
	if len(_iotSetAsActive) > 0 {
		if err := assignInputField(input, "SetAsActive", _iotSetAsActive); err != nil {
			log.Errorf("invalid --set-as-active: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCertificateFromCsr(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services IoT Core certificate provider. You can use
// Amazon Web Services IoT Core certificate provider to customize how to sign a
// certificate signing request (CSR) in IoT fleet provisioning. For more
// information, see [Customizing certificate signing using Amazon Web Services IoT Core certificate provider]from Amazon Web Services IoT Core Developer Guide.
//
// Requires permission to access the [CreateCertificateProvider] action.
//
// After you create a certificate provider, the behavior of [CreateCertificateFromCsr API for fleet provisioning]
// CreateCertificateFromCsr will change and all API calls to
// CreateCertificateFromCsr will invoke the certificate provider to create the
// certificates. It can take up to a few minutes for this behavior to change after
// a certificate provider is created.
//
// [CreateCertificateProvider]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [CreateCertificateFromCsr API for fleet provisioning]: https://docs.aws.amazon.com/iot/latest/developerguide/fleet-provision-api.html#create-cert-csr
// [Customizing certificate signing using Amazon Web Services IoT Core certificate provider]: https://docs.aws.amazon.com/iot/latest/developerguide/provisioning-cert-provider.html
func iot_CreateCertificateProvider(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateCertificateProviderInput{
		// AccountDefaultForOperations: []types.CertificateProviderOperation, // Required
		// CertificateProviderName: *string, // Required
		// LambdaFunctionArn: *string, // Required
	}

	if len(_iotAccountDefaultForOperations) > 0 {
		if err := assignInputField(input, "AccountDefaultForOperations", _iotAccountDefaultForOperations); err != nil {
			log.Errorf("invalid --account-default-for-operations: %s", err.Error())
			return
		}
	}
	if len(_iotCertificateProviderName) > 0 {
		input.CertificateProviderName = aws.String(_iotCertificateProviderName)
	}
	if len(_iotLambdaFunctionArn) > 0 {
		input.LambdaFunctionArn = aws.String(_iotLambdaFunctionArn)
	}
	if len(_iotClientToken) > 0 {
		input.ClientToken = aws.String(_iotClientToken)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCertificateProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a command. A command contains reusable configurations that can be
// applied before they are sent to the devices.
func iot_CreateCommand(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateCommandInput{
		// CommandId: *string, // Required
	}

	if len(_iotCommandId) > 0 {
		input.CommandId = aws.String(_iotCommandId)
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotDisplayName) > 0 {
		input.DisplayName = aws.String(_iotDisplayName)
	}
	if len(_iotMandatoryParameters) > 0 {
		if err := assignInputField(input, "MandatoryParameters", _iotMandatoryParameters); err != nil {
			log.Errorf("invalid --mandatory-parameters: %s", err.Error())
			return
		}
	}
	if len(_iotNamespace) > 0 {
		if err := assignInputField(input, "Namespace", _iotNamespace); err != nil {
			log.Errorf("invalid --namespace: %s", err.Error())
			return
		}
	}
	if len(_iotPayload) > 0 {
		if err := assignInputField(input, "Payload", _iotPayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_iotPayloadTemplate) > 0 {
		input.PayloadTemplate = aws.String(_iotPayloadTemplate)
	}
	if len(_iotPreprocessor) > 0 {
		if err := assignInputField(input, "Preprocessor", _iotPreprocessor); err != nil {
			log.Errorf("invalid --preprocessor: %s", err.Error())
			return
		}
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCommand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this API to define a Custom Metric published by your devices to Device
// Defender.
//
// Requires permission to access the [CreateCustomMetric] action.
//
// [CreateCustomMetric]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateCustomMetric(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateCustomMetricInput{
		// ClientRequestToken: *string, // Required
		// MetricName: *string, // Required
		// MetricType: types.CustomMetricType, // Required
	}

	if len(_iotClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotClientRequestToken)
	}
	if len(_iotMetricName) > 0 {
		input.MetricName = aws.String(_iotMetricName)
	}
	if len(_iotMetricType) > 0 {
		if err := assignInputField(input, "MetricType", _iotMetricType); err != nil {
			log.Errorf("invalid --metric-type: %s", err.Error())
			return
		}
	}
	if len(_iotDisplayName) > 0 {
		input.DisplayName = aws.String(_iotDisplayName)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomMetric(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a dimension that you can use to limit the scope of a metric used in a
// security profile for IoT Device Defender. For example, using a TOPIC_FILTER
// dimension, you can narrow down the scope of the metric only to MQTT topics whose
// name match the pattern specified in the dimension.
//
// Requires permission to access the [CreateDimension] action.
//
// [CreateDimension]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateDimension(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateDimensionInput{
		// ClientRequestToken: *string, // Required
		// Name: *string, // Required
		// StringValues: []string, // Required
		// Type: types.DimensionType, // Required
	}

	if len(_iotClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotClientRequestToken)
	}
	if len(_iotName) > 0 {
		input.Name = aws.String(_iotName)
	}
	if len(_iotStringValues) > 0 {
		input.StringValues = append([]string(nil), _iotStringValues...)
	}
	if len(_iotType) > 0 {
		if err := assignInputField(input, "Type", _iotType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDimension(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a domain configuration.
// Requires permission to access the [CreateDomainConfiguration] action.
//
// [CreateDomainConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateDomainConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateDomainConfigurationInput{
		// DomainConfigurationName: *string, // Required
	}

	if len(_iotDomainConfigurationName) > 0 {
		input.DomainConfigurationName = aws.String(_iotDomainConfigurationName)
	}
	if len(_iotApplicationProtocol) > 0 {
		if err := assignInputField(input, "ApplicationProtocol", _iotApplicationProtocol); err != nil {
			log.Errorf("invalid --application-protocol: %s", err.Error())
			return
		}
	}
	if len(_iotAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _iotAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_iotAuthorizerConfig) > 0 {
		if err := assignInputField(input, "AuthorizerConfig", _iotAuthorizerConfig); err != nil {
			log.Errorf("invalid --authorizer-config: %s", err.Error())
			return
		}
	}
	if len(_iotClientCertificateConfig) > 0 {
		if err := assignInputField(input, "ClientCertificateConfig", _iotClientCertificateConfig); err != nil {
			log.Errorf("invalid --client-certificate-config: %s", err.Error())
			return
		}
	}
	if len(_iotDomainName) > 0 {
		input.DomainName = aws.String(_iotDomainName)
	}
	if len(_iotServerCertificateArns) > 0 {
		input.ServerCertificateArns = append([]string(nil), _iotServerCertificateArns...)
	}
	if len(_iotServerCertificateConfig) > 0 {
		if err := assignInputField(input, "ServerCertificateConfig", _iotServerCertificateConfig); err != nil {
			log.Errorf("invalid --server-certificate-config: %s", err.Error())
			return
		}
	}
	if len(_iotServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _iotServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotTlsConfig) > 0 {
		if err := assignInputField(input, "TlsConfig", _iotTlsConfig); err != nil {
			log.Errorf("invalid --tls-config: %s", err.Error())
			return
		}
	}
	if len(_iotValidationCertificateArn) > 0 {
		input.ValidationCertificateArn = aws.String(_iotValidationCertificateArn)
	}

	if resp, err := client.CreateDomainConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a dynamic thing group.
// Requires permission to access the [CreateDynamicThingGroup] action.
//
// [CreateDynamicThingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateDynamicThingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateDynamicThingGroupInput{
		// QueryString: *string, // Required
		// ThingGroupName: *string, // Required
	}

	if len(_iotQueryString) > 0 {
		input.QueryString = aws.String(_iotQueryString)
	}
	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}
	if len(_iotIndexName) > 0 {
		input.IndexName = aws.String(_iotIndexName)
	}
	if len(_iotQueryVersion) > 0 {
		input.QueryVersion = aws.String(_iotQueryVersion)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotThingGroupProperties) > 0 {
		if err := assignInputField(input, "ThingGroupProperties", _iotThingGroupProperties); err != nil {
			log.Errorf("invalid --thing-group-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDynamicThingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a fleet metric.
// Requires permission to access the [CreateFleetMetric] action.
//
// [CreateFleetMetric]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateFleetMetric(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateFleetMetricInput{
		// AggregationField: *string, // Required
		// AggregationType: *types.AggregationType, // Required
		// MetricName: *string, // Required
		// Period: *int32, // Required
		// QueryString: *string, // Required
	}

	if len(_iotAggregationField) > 0 {
		input.AggregationField = aws.String(_iotAggregationField)
	}
	if len(_iotAggregationType) > 0 {
		if err := assignInputField(input, "AggregationType", _iotAggregationType); err != nil {
			log.Errorf("invalid --aggregation-type: %s", err.Error())
			return
		}
	}
	if len(_iotMetricName) > 0 {
		input.MetricName = aws.String(_iotMetricName)
	}
	if len(_iotPeriod) > 0 {
		if err := assignInputField(input, "Period", _iotPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_iotQueryString) > 0 {
		input.QueryString = aws.String(_iotQueryString)
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotIndexName) > 0 {
		input.IndexName = aws.String(_iotIndexName)
	}
	if len(_iotQueryVersion) > 0 {
		input.QueryVersion = aws.String(_iotQueryVersion)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotUnit) > 0 {
		if err := assignInputField(input, "Unit", _iotUnit); err != nil {
			log.Errorf("invalid --unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFleetMetric(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a job.
// Requires permission to access the [CreateJob] action.
//
// [CreateJob]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateJob(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateJobInput{
		// JobId: *string, // Required
		// Targets: []string, // Required
	}

	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotTargets) > 0 {
		input.Targets = append([]string(nil), _iotTargets...)
	}
	if len(_iotAbortConfig) > 0 {
		if err := assignInputField(input, "AbortConfig", _iotAbortConfig); err != nil {
			log.Errorf("invalid --abort-config: %s", err.Error())
			return
		}
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotDestinationPackageVersions) > 0 {
		input.DestinationPackageVersions = append([]string(nil), _iotDestinationPackageVersions...)
	}
	if len(_iotDocument) > 0 {
		input.Document = aws.String(_iotDocument)
	}
	if len(_iotDocumentParameters) > 0 {
		if err := assignInputField(input, "DocumentParameters", _iotDocumentParameters); err != nil {
			log.Errorf("invalid --document-parameters: %s", err.Error())
			return
		}
	}
	if len(_iotDocumentSource) > 0 {
		input.DocumentSource = aws.String(_iotDocumentSource)
	}
	if len(_iotJobExecutionsRetryConfig) > 0 {
		if err := assignInputField(input, "JobExecutionsRetryConfig", _iotJobExecutionsRetryConfig); err != nil {
			log.Errorf("invalid --job-executions-retry-config: %s", err.Error())
			return
		}
	}
	if len(_iotJobExecutionsRolloutConfig) > 0 {
		if err := assignInputField(input, "JobExecutionsRolloutConfig", _iotJobExecutionsRolloutConfig); err != nil {
			log.Errorf("invalid --job-executions-rollout-config: %s", err.Error())
			return
		}
	}
	if len(_iotJobTemplateArn) > 0 {
		input.JobTemplateArn = aws.String(_iotJobTemplateArn)
	}
	if len(_iotNamespaceId) > 0 {
		input.NamespaceId = aws.String(_iotNamespaceId)
	}
	if len(_iotPresignedUrlConfig) > 0 {
		if err := assignInputField(input, "PresignedUrlConfig", _iotPresignedUrlConfig); err != nil {
			log.Errorf("invalid --presigned-url-config: %s", err.Error())
			return
		}
	}
	if len(_iotSchedulingConfig) > 0 {
		if err := assignInputField(input, "SchedulingConfig", _iotSchedulingConfig); err != nil {
			log.Errorf("invalid --scheduling-config: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotTargetSelection) > 0 {
		if err := assignInputField(input, "TargetSelection", _iotTargetSelection); err != nil {
			log.Errorf("invalid --target-selection: %s", err.Error())
			return
		}
	}
	if len(_iotTimeoutConfig) > 0 {
		if err := assignInputField(input, "TimeoutConfig", _iotTimeoutConfig); err != nil {
			log.Errorf("invalid --timeout-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a job template.
// Requires permission to access the [CreateJobTemplate] action.
//
// [CreateJobTemplate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateJobTemplate(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateJobTemplateInput{
		// Description: *string, // Required
		// JobTemplateId: *string, // Required
	}

	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotJobTemplateId) > 0 {
		input.JobTemplateId = aws.String(_iotJobTemplateId)
	}
	if len(_iotAbortConfig) > 0 {
		if err := assignInputField(input, "AbortConfig", _iotAbortConfig); err != nil {
			log.Errorf("invalid --abort-config: %s", err.Error())
			return
		}
	}
	if len(_iotDestinationPackageVersions) > 0 {
		input.DestinationPackageVersions = append([]string(nil), _iotDestinationPackageVersions...)
	}
	if len(_iotDocument) > 0 {
		input.Document = aws.String(_iotDocument)
	}
	if len(_iotDocumentSource) > 0 {
		input.DocumentSource = aws.String(_iotDocumentSource)
	}
	if len(_iotJobArn) > 0 {
		input.JobArn = aws.String(_iotJobArn)
	}
	if len(_iotJobExecutionsRetryConfig) > 0 {
		if err := assignInputField(input, "JobExecutionsRetryConfig", _iotJobExecutionsRetryConfig); err != nil {
			log.Errorf("invalid --job-executions-retry-config: %s", err.Error())
			return
		}
	}
	if len(_iotJobExecutionsRolloutConfig) > 0 {
		if err := assignInputField(input, "JobExecutionsRolloutConfig", _iotJobExecutionsRolloutConfig); err != nil {
			log.Errorf("invalid --job-executions-rollout-config: %s", err.Error())
			return
		}
	}
	if len(_iotMaintenanceWindows) > 0 {
		if err := assignInputField(input, "MaintenanceWindows", _iotMaintenanceWindows); err != nil {
			log.Errorf("invalid --maintenance-windows: %s", err.Error())
			return
		}
	}
	if len(_iotPresignedUrlConfig) > 0 {
		if err := assignInputField(input, "PresignedUrlConfig", _iotPresignedUrlConfig); err != nil {
			log.Errorf("invalid --presigned-url-config: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotTimeoutConfig) > 0 {
		if err := assignInputField(input, "TimeoutConfig", _iotTimeoutConfig); err != nil {
			log.Errorf("invalid --timeout-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a 2048-bit RSA key pair and issues an X.509 certificate using the
// issued public key. You can also call CreateKeysAndCertificate over MQTT from a
// device, for more information, see [Provisioning MQTT API].
//
// Note This is the only time IoT issues the private key for this certificate, so
// it is important to keep it in a secure location.
//
// Requires permission to access the [CreateKeysAndCertificate] action.
//
// [CreateKeysAndCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [Provisioning MQTT API]: https://docs.aws.amazon.com/iot/latest/developerguide/provision-wo-cert.html#provision-mqtt-api
func iot_CreateKeysAndCertificate(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateKeysAndCertificateInput{}

	if len(_iotSetAsActive) > 0 {
		if err := assignInputField(input, "SetAsActive", _iotSetAsActive); err != nil {
			log.Errorf("invalid --set-as-active: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKeysAndCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines an action that can be applied to audit findings by using
// StartAuditMitigationActionsTask. Only certain types of mitigation actions can be
// applied to specific check names. For more information, see [Mitigation actions]. Each mitigation
// action can apply only one type of change.
//
// Requires permission to access the [CreateMitigationAction] action.
//
// [Mitigation actions]: https://docs.aws.amazon.com/iot/latest/developerguide/device-defender-mitigation-actions.html
// [CreateMitigationAction]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateMitigationAction(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateMitigationActionInput{
		// ActionName: *string, // Required
		// ActionParams: *types.MitigationActionParams, // Required
		// RoleArn: *string, // Required
	}

	if len(_iotActionName) > 0 {
		input.ActionName = aws.String(_iotActionName)
	}
	if len(_iotActionParams) > 0 {
		if err := assignInputField(input, "ActionParams", _iotActionParams); err != nil {
			log.Errorf("invalid --action-params: %s", err.Error())
			return
		}
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMitigationAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IoT OTA update on a target group of things or groups.
// Requires permission to access the [CreateOTAUpdate] action.
//
// [CreateOTAUpdate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateOTAUpdate(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateOTAUpdateInput{
		// Files: []types.OTAUpdateFile, // Required
		// OtaUpdateId: *string, // Required
		// RoleArn: *string, // Required
		// Targets: []string, // Required
	}

	if len(_iotFiles) > 0 {
		if err := assignInputField(input, "Files", _iotFiles); err != nil {
			log.Errorf("invalid --files: %s", err.Error())
			return
		}
	}
	if len(_iotOtaUpdateId) > 0 {
		input.OtaUpdateId = aws.String(_iotOtaUpdateId)
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}
	if len(_iotTargets) > 0 {
		input.Targets = append([]string(nil), _iotTargets...)
	}
	if len(_iotAdditionalParameters) > 0 {
		if err := assignInputField(input, "AdditionalParameters", _iotAdditionalParameters); err != nil {
			log.Errorf("invalid --additional-parameters: %s", err.Error())
			return
		}
	}
	if len(_iotAwsJobAbortConfig) > 0 {
		if err := assignInputField(input, "AwsJobAbortConfig", _iotAwsJobAbortConfig); err != nil {
			log.Errorf("invalid --aws-job-abort-config: %s", err.Error())
			return
		}
	}
	if len(_iotAwsJobExecutionsRolloutConfig) > 0 {
		if err := assignInputField(input, "AwsJobExecutionsRolloutConfig", _iotAwsJobExecutionsRolloutConfig); err != nil {
			log.Errorf("invalid --aws-job-executions-rollout-config: %s", err.Error())
			return
		}
	}
	if len(_iotAwsJobPresignedUrlConfig) > 0 {
		if err := assignInputField(input, "AwsJobPresignedUrlConfig", _iotAwsJobPresignedUrlConfig); err != nil {
			log.Errorf("invalid --aws-job-presigned-url-config: %s", err.Error())
			return
		}
	}
	if len(_iotAwsJobTimeoutConfig) > 0 {
		if err := assignInputField(input, "AwsJobTimeoutConfig", _iotAwsJobTimeoutConfig); err != nil {
			log.Errorf("invalid --aws-job-timeout-config: %s", err.Error())
			return
		}
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotProtocols) > 0 {
		if err := assignInputField(input, "Protocols", _iotProtocols); err != nil {
			log.Errorf("invalid --protocols: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotTargetSelection) > 0 {
		if err := assignInputField(input, "TargetSelection", _iotTargetSelection); err != nil {
			log.Errorf("invalid --target-selection: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOTAUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IoT software package that can be deployed to your fleet.
// Requires permission to access the [CreatePackage] and [GetIndexingConfiguration] actions.
//
// [CreatePackage]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [GetIndexingConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreatePackage(cfg aws.Config, client *iot.Client) {
	input := &iot.CreatePackageInput{
		// PackageName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotClientToken) > 0 {
		input.ClientToken = aws.String(_iotClientToken)
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version for an existing IoT software package.
// Requires permission to access the [CreatePackageVersion] and [GetIndexingConfiguration] actions.
//
// [CreatePackageVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [GetIndexingConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreatePackageVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.CreatePackageVersionInput{
		// PackageName: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotVersionName) > 0 {
		input.VersionName = aws.String(_iotVersionName)
	}
	if len(_iotArtifact) > 0 {
		if err := assignInputField(input, "Artifact", _iotArtifact); err != nil {
			log.Errorf("invalid --artifact: %s", err.Error())
			return
		}
	}
	if len(_iotAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _iotAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_iotClientToken) > 0 {
		input.ClientToken = aws.String(_iotClientToken)
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotRecipe) > 0 {
		input.Recipe = aws.String(_iotRecipe)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IoT policy.
// The created policy is the default version for the policy. This operation
// creates a policy version with a version identifier of 1 and sets 1 as the
// policy's default version.
//
// Requires permission to access the [CreatePolicy] action.
//
// [CreatePolicy]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreatePolicy(cfg aws.Config, client *iot.Client) {
	input := &iot.CreatePolicyInput{
		// PolicyDocument: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_iotPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_iotPolicyDocument)
	}
	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of the specified IoT policy. To update a policy, create a
// new policy version. A managed policy can have up to five versions. If the policy
// has five versions, you must use DeletePolicyVersionto delete an existing version before you create
// a new one.
//
// Optionally, you can set the new version as the policy's default version. The
// default version is the operative version (that is, the version that is in effect
// for the certificates to which the policy is attached).
//
// Requires permission to access the [CreatePolicyVersion] action.
//
// [CreatePolicyVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreatePolicyVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.CreatePolicyVersionInput{
		// PolicyDocument: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_iotPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_iotPolicyDocument)
	}
	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotSetAsDefault) > 0 {
		if err := assignInputField(input, "SetAsDefault", _iotSetAsDefault); err != nil {
			log.Errorf("invalid --set-as-default: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a provisioning claim.
// Requires permission to access the [CreateProvisioningClaim] action.
//
// [CreateProvisioningClaim]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateProvisioningClaim(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateProvisioningClaimInput{
		// TemplateName: *string, // Required
	}

	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}

	if resp, err := client.CreateProvisioningClaim(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a provisioning template.
// Requires permission to access the [CreateProvisioningTemplate] action.
//
// [CreateProvisioningTemplate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateProvisioningTemplate(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateProvisioningTemplateInput{
		// ProvisioningRoleArn: *string, // Required
		// TemplateBody: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_iotProvisioningRoleArn) > 0 {
		input.ProvisioningRoleArn = aws.String(_iotProvisioningRoleArn)
	}
	if len(_iotTemplateBody) > 0 {
		input.TemplateBody = aws.String(_iotTemplateBody)
	}
	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _iotEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_iotPreProvisioningHook) > 0 {
		if err := assignInputField(input, "PreProvisioningHook", _iotPreProvisioningHook); err != nil {
			log.Errorf("invalid --pre-provisioning-hook: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotType) > 0 {
		if err := assignInputField(input, "Type", _iotType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProvisioningTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of a provisioning template.
// Requires permission to access the [CreateProvisioningTemplateVersion] action.
//
// [CreateProvisioningTemplateVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateProvisioningTemplateVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateProvisioningTemplateVersionInput{
		// TemplateBody: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_iotTemplateBody) > 0 {
		input.TemplateBody = aws.String(_iotTemplateBody)
	}
	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}
	if len(_iotSetAsDefault) > 0 {
		if err := assignInputField(input, "SetAsDefault", _iotSetAsDefault); err != nil {
			log.Errorf("invalid --set-as-default: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProvisioningTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a role alias.
// Requires permission to access the [CreateRoleAlias] action.
//
// The value of [credentialDurationSeconds]credentialDurationSeconds must be less than or equal to the
// maximum session duration of the IAM role that the role alias references. For
// more information, see [Modifying a role maximum session duration (Amazon Web Services API)]from the Amazon Web Services Identity and Access
// Management User Guide.
//
// [CreateRoleAlias]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [Modifying a role maximum session duration (Amazon Web Services API)]: https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-managingrole-editing-api.html#roles-modify_max-session-duration-api
// [credentialDurationSeconds]: https://docs.aws.amazon.com/iot/latest/apireference/API_CreateRoleAlias.html#iot-CreateRoleAlias-request-credentialDurationSeconds
func iot_CreateRoleAlias(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateRoleAliasInput{
		// RoleAlias: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_iotRoleAlias) > 0 {
		input.RoleAlias = aws.String(_iotRoleAlias)
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}
	if len(_iotCredentialDurationSeconds) > 0 {
		if err := assignInputField(input, "CredentialDurationSeconds", _iotCredentialDurationSeconds); err != nil {
			log.Errorf("invalid --credential-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRoleAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a scheduled audit that is run at a specified time interval.
// Requires permission to access the [CreateScheduledAudit] action.
//
// [CreateScheduledAudit]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateScheduledAudit(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateScheduledAuditInput{
		// Frequency: types.AuditFrequency, // Required
		// ScheduledAuditName: *string, // Required
		// TargetCheckNames: []string, // Required
	}

	if len(_iotFrequency) > 0 {
		if err := assignInputField(input, "Frequency", _iotFrequency); err != nil {
			log.Errorf("invalid --frequency: %s", err.Error())
			return
		}
	}
	if len(_iotScheduledAuditName) > 0 {
		input.ScheduledAuditName = aws.String(_iotScheduledAuditName)
	}
	if len(_iotTargetCheckNames) > 0 {
		input.TargetCheckNames = append([]string(nil), _iotTargetCheckNames...)
	}
	if len(_iotDayOfMonth) > 0 {
		input.DayOfMonth = aws.String(_iotDayOfMonth)
	}
	if len(_iotDayOfWeek) > 0 {
		if err := assignInputField(input, "DayOfWeek", _iotDayOfWeek); err != nil {
			log.Errorf("invalid --day-of-week: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateScheduledAudit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Device Defender security profile.
// Requires permission to access the [CreateSecurityProfile] action.
//
// [CreateSecurityProfile]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateSecurityProfile(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateSecurityProfileInput{
		// SecurityProfileName: *string, // Required
	}

	if len(_iotSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_iotSecurityProfileName)
	}
	if len(_iotAdditionalMetricsToRetain) > 0 {
		input.AdditionalMetricsToRetain = append([]string(nil), _iotAdditionalMetricsToRetain...)
	}
	if len(_iotAdditionalMetricsToRetainV2) > 0 {
		if err := assignInputField(input, "AdditionalMetricsToRetainV2", _iotAdditionalMetricsToRetainV2); err != nil {
			log.Errorf("invalid --additional-metrics-to-retain-v2: %s", err.Error())
			return
		}
	}
	if len(_iotAlertTargets) > 0 {
		if err := assignInputField(input, "AlertTargets", _iotAlertTargets); err != nil {
			log.Errorf("invalid --alert-targets: %s", err.Error())
			return
		}
	}
	if len(_iotBehaviors) > 0 {
		if err := assignInputField(input, "Behaviors", _iotBehaviors); err != nil {
			log.Errorf("invalid --behaviors: %s", err.Error())
			return
		}
	}
	if len(_iotMetricsExportConfig) > 0 {
		if err := assignInputField(input, "MetricsExportConfig", _iotMetricsExportConfig); err != nil {
			log.Errorf("invalid --metrics-export-config: %s", err.Error())
			return
		}
	}
	if len(_iotSecurityProfileDescription) > 0 {
		input.SecurityProfileDescription = aws.String(_iotSecurityProfileDescription)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSecurityProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a stream for delivering one or more large files in chunks over MQTT. A
// stream transports data bytes in chunks or blocks packaged as MQTT messages from
// a source like S3. You can have one or more files associated with a stream.
//
// Requires permission to access the [CreateStream] action.
//
// [CreateStream]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateStream(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateStreamInput{
		// Files: []types.StreamFile, // Required
		// RoleArn: *string, // Required
		// StreamId: *string, // Required
	}

	if len(_iotFiles) > 0 {
		if err := assignInputField(input, "Files", _iotFiles); err != nil {
			log.Errorf("invalid --files: %s", err.Error())
			return
		}
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}
	if len(_iotStreamId) > 0 {
		input.StreamId = aws.String(_iotStreamId)
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a thing record in the registry. If this call is made multiple times
// using the same thing name and configuration, the call will succeed. If this call
// is made with the same thing name but different configuration a
// ResourceAlreadyExistsException is thrown.
//
// This is a control plane operation. See [Authorization] for information about authorizing
// control plane actions.
//
// Requires permission to access the [CreateThing] action.
//
// [Authorization]: https://docs.aws.amazon.com/iot/latest/developerguide/iot-authorization.html
// [CreateThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateThing(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateThingInput{
		// ThingName: *string, // Required
	}

	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotAttributePayload) > 0 {
		if err := assignInputField(input, "AttributePayload", _iotAttributePayload); err != nil {
			log.Errorf("invalid --attribute-payload: %s", err.Error())
			return
		}
	}
	if len(_iotBillingGroupName) > 0 {
		input.BillingGroupName = aws.String(_iotBillingGroupName)
	}
	if len(_iotThingTypeName) > 0 {
		input.ThingTypeName = aws.String(_iotThingTypeName)
	}

	if resp, err := client.CreateThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a thing group.
// This is a control plane operation. See [Authorization] for information about authorizing
// control plane actions.
//
// If the ThingGroup that you create has the exact same attributes as an existing
// ThingGroup , you will get a 200 success response.
//
// Requires permission to access the [CreateThingGroup] action.
//
// [Authorization]: https://docs.aws.amazon.com/iot/latest/developerguide/iot-authorization.html
// [CreateThingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateThingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateThingGroupInput{
		// ThingGroupName: *string, // Required
	}

	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}
	if len(_iotParentGroupName) > 0 {
		input.ParentGroupName = aws.String(_iotParentGroupName)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotThingGroupProperties) > 0 {
		if err := assignInputField(input, "ThingGroupProperties", _iotThingGroupProperties); err != nil {
			log.Errorf("invalid --thing-group-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateThingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new thing type. If this call is made multiple times using the same
// thing type name and configuration, the call will succeed. If this call is made
// with the same thing type name but different configuration a
// ResourceAlreadyExistsException is thrown.
//
// Requires permission to access the [CreateThingType] action.
//
// [CreateThingType]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateThingType(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateThingTypeInput{
		// ThingTypeName: *string, // Required
	}

	if len(_iotThingTypeName) > 0 {
		input.ThingTypeName = aws.String(_iotThingTypeName)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotThingTypeProperties) > 0 {
		if err := assignInputField(input, "ThingTypeProperties", _iotThingTypeProperties); err != nil {
			log.Errorf("invalid --thing-type-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateThingType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a rule. Creating rules is an administrator-level action. Any user who
// has permission to create rules will be able to access data processed by the
// rule.
//
// Requires permission to access the [CreateTopicRule] action.
//
// [CreateTopicRule]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateTopicRule(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateTopicRuleInput{
		// RuleName: *string, // Required
		// TopicRulePayload: *types.TopicRulePayload, // Required
	}

	if len(_iotRuleName) > 0 {
		input.RuleName = aws.String(_iotRuleName)
	}
	if len(_iotTopicRulePayload) > 0 {
		if err := assignInputField(input, "TopicRulePayload", _iotTopicRulePayload); err != nil {
			log.Errorf("invalid --topic-rule-payload: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		input.Tags = aws.String(_iotTags)
	}

	if resp, err := client.CreateTopicRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a topic rule destination. The destination must be confirmed prior to
// use.
//
// Requires permission to access the [CreateTopicRuleDestination] action.
//
// [CreateTopicRuleDestination]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_CreateTopicRuleDestination(cfg aws.Config, client *iot.Client) {
	input := &iot.CreateTopicRuleDestinationInput{
		// DestinationConfiguration: *types.TopicRuleDestinationConfiguration, // Required
	}

	if len(_iotDestinationConfiguration) > 0 {
		if err := assignInputField(input, "DestinationConfiguration", _iotDestinationConfiguration); err != nil {
			log.Errorf("invalid --destination-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTopicRuleDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores the default settings for Device Defender audits for this account. Any
// configuration data you entered is deleted and all audit checks are reset to
// disabled.
//
// Requires permission to access the [DeleteAccountAuditConfiguration] action.
//
// [DeleteAccountAuditConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteAccountAuditConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteAccountAuditConfigurationInput{}

	if len(_iotDeleteScheduledAudits) > 0 {
		if err := assignInputField(input, "DeleteScheduledAudits", _iotDeleteScheduledAudits); err != nil {
			log.Errorf("invalid --delete-scheduled-audits: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAccountAuditConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Device Defender audit suppression.
// Requires permission to access the [DeleteAuditSuppression] action.
//
// [DeleteAuditSuppression]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteAuditSuppression(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteAuditSuppressionInput{
		// CheckName: *string, // Required
		// ResourceIdentifier: *types.ResourceIdentifier, // Required
	}

	if len(_iotCheckName) > 0 {
		input.CheckName = aws.String(_iotCheckName)
	}
	if len(_iotResourceIdentifier) > 0 {
		if err := assignInputField(input, "ResourceIdentifier", _iotResourceIdentifier); err != nil {
			log.Errorf("invalid --resource-identifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAuditSuppression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an authorizer.
// Requires permission to access the [DeleteAuthorizer] action.
//
// [DeleteAuthorizer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteAuthorizer(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteAuthorizerInput{
		// AuthorizerName: *string, // Required
	}

	if len(_iotAuthorizerName) > 0 {
		input.AuthorizerName = aws.String(_iotAuthorizerName)
	}

	if resp, err := client.DeleteAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the billing group.
// Requires permission to access the [DeleteBillingGroup] action.
//
// [DeleteBillingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteBillingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteBillingGroupInput{
		// BillingGroupName: *string, // Required
	}

	if len(_iotBillingGroupName) > 0 {
		input.BillingGroupName = aws.String(_iotBillingGroupName)
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteBillingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a registered CA certificate.
// Requires permission to access the [DeleteCACertificate] action.
//
// [DeleteCACertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteCACertificate(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteCACertificateInput{
		// CertificateId: *string, // Required
	}

	if len(_iotCertificateId) > 0 {
		input.CertificateId = aws.String(_iotCertificateId)
	}

	if resp, err := client.DeleteCACertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified certificate.
// A certificate cannot be deleted if it has a policy or IoT thing attached to it
// or if its status is set to ACTIVE. To delete a certificate, first use the DetachPolicy
// action to detach all policies. Next, use the UpdateCertificateaction to set the certificate to
// the INACTIVE status.
//
// Requires permission to access the [DeleteCertificate] action.
//
// [DeleteCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteCertificate(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteCertificateInput{
		// CertificateId: *string, // Required
	}

	if len(_iotCertificateId) > 0 {
		input.CertificateId = aws.String(_iotCertificateId)
	}
	if len(_iotForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _iotForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a certificate provider.
// Requires permission to access the [DeleteCertificateProvider] action.
//
// If you delete the certificate provider resource, the behavior of
// CreateCertificateFromCsr will resume, and IoT will create certificates signed by
// IoT from a certificate signing request (CSR).
//
// [DeleteCertificateProvider]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteCertificateProvider(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteCertificateProviderInput{
		// CertificateProviderName: *string, // Required
	}

	if len(_iotCertificateProviderName) > 0 {
		input.CertificateProviderName = aws.String(_iotCertificateProviderName)
	}

	if resp, err := client.DeleteCertificateProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a command resource.
func iot_DeleteCommand(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteCommandInput{
		// CommandId: *string, // Required
	}

	if len(_iotCommandId) > 0 {
		input.CommandId = aws.String(_iotCommandId)
	}

	if resp, err := client.DeleteCommand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a command execution.
// Only command executions that enter a terminal state can be deleted from your
// account.
func iot_DeleteCommandExecution(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteCommandExecutionInput{
		// ExecutionId: *string, // Required
		// TargetArn: *string, // Required
	}

	if len(_iotExecutionId) > 0 {
		input.ExecutionId = aws.String(_iotExecutionId)
	}
	if len(_iotTargetArn) > 0 {
		input.TargetArn = aws.String(_iotTargetArn)
	}

	if resp, err := client.DeleteCommandExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Device Defender detect custom metric.
// Requires permission to access the [DeleteCustomMetric] action.
//
// Before you can delete a custom metric, you must first remove the custom metric
// from all security profiles it's a part of. The security profile associated with
// the custom metric can be found using the [ListSecurityProfiles]API with metricName set to your custom
// metric name.
//
// [DeleteCustomMetric]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [ListSecurityProfiles]: https://docs.aws.amazon.com/iot/latest/apireference/API_ListSecurityProfiles.html
func iot_DeleteCustomMetric(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteCustomMetricInput{
		// MetricName: *string, // Required
	}

	if len(_iotMetricName) > 0 {
		input.MetricName = aws.String(_iotMetricName)
	}

	if resp, err := client.DeleteCustomMetric(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified dimension from your Amazon Web Services accounts.
// Requires permission to access the [DeleteDimension] action.
//
// [DeleteDimension]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteDimension(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteDimensionInput{
		// Name: *string, // Required
	}

	if len(_iotName) > 0 {
		input.Name = aws.String(_iotName)
	}

	if resp, err := client.DeleteDimension(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified domain configuration.
// Requires permission to access the [DeleteDomainConfiguration] action.
//
// [DeleteDomainConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteDomainConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteDomainConfigurationInput{
		// DomainConfigurationName: *string, // Required
	}

	if len(_iotDomainConfigurationName) > 0 {
		input.DomainConfigurationName = aws.String(_iotDomainConfigurationName)
	}

	if resp, err := client.DeleteDomainConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a dynamic thing group.
// Requires permission to access the [DeleteDynamicThingGroup] action.
//
// [DeleteDynamicThingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteDynamicThingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteDynamicThingGroupInput{
		// ThingGroupName: *string, // Required
	}

	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDynamicThingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified fleet metric. Returns successfully with no error if the
// deletion is successful or you specify a fleet metric that doesn't exist.
//
// Requires permission to access the [DeleteFleetMetric] action.
//
// [DeleteFleetMetric]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteFleetMetric(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteFleetMetricInput{
		// MetricName: *string, // Required
	}

	if len(_iotMetricName) > 0 {
		input.MetricName = aws.String(_iotMetricName)
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteFleetMetric(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a job and its related job executions.
// Deleting a job may take time, depending on the number of job executions created
// for the job and various other factors. While the job is being deleted, the
// status of the job will be shown as "DELETION_IN_PROGRESS". Attempting to delete
// or cancel a job whose status is already "DELETION_IN_PROGRESS" will result in an
// error.
//
// Only 10 jobs may have status "DELETION_IN_PROGRESS" at the same time, or a
// LimitExceededException will occur.
//
// Requires permission to access the [DeleteJob] action.
//
// [DeleteJob]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteJob(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteJobInput{
		// JobId: *string, // Required
	}

	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotForce) > 0 {
		if err := assignInputField(input, "Force", _iotForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_iotNamespaceId) > 0 {
		input.NamespaceId = aws.String(_iotNamespaceId)
	}

	if resp, err := client.DeleteJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a job execution.
// Requires permission to access the [DeleteJobExecution] action.
//
// [DeleteJobExecution]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteJobExecution(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteJobExecutionInput{
		// ExecutionNumber: *int64, // Required
		// JobId: *string, // Required
		// ThingName: *string, // Required
	}

	if len(_iotExecutionNumber) > 0 {
		if err := assignInputField(input, "ExecutionNumber", _iotExecutionNumber); err != nil {
			log.Errorf("invalid --execution-number: %s", err.Error())
			return
		}
	}
	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotForce) > 0 {
		if err := assignInputField(input, "Force", _iotForce); err != nil {
			log.Errorf("invalid --force: %s", err.Error())
			return
		}
	}
	if len(_iotNamespaceId) > 0 {
		input.NamespaceId = aws.String(_iotNamespaceId)
	}

	if resp, err := client.DeleteJobExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified job template.
func iot_DeleteJobTemplate(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteJobTemplateInput{
		// JobTemplateId: *string, // Required
	}

	if len(_iotJobTemplateId) > 0 {
		input.JobTemplateId = aws.String(_iotJobTemplateId)
	}

	if resp, err := client.DeleteJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a defined mitigation action from your Amazon Web Services accounts.
// Requires permission to access the [DeleteMitigationAction] action.
//
// [DeleteMitigationAction]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteMitigationAction(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteMitigationActionInput{
		// ActionName: *string, // Required
	}

	if len(_iotActionName) > 0 {
		input.ActionName = aws.String(_iotActionName)
	}

	if resp, err := client.DeleteMitigationAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an OTA update.
// Requires permission to access the [DeleteOTAUpdate] action.
//
// [DeleteOTAUpdate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteOTAUpdate(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteOTAUpdateInput{
		// OtaUpdateId: *string, // Required
	}

	if len(_iotOtaUpdateId) > 0 {
		input.OtaUpdateId = aws.String(_iotOtaUpdateId)
	}
	if len(_iotForceDeleteAWSJob) > 0 {
		if err := assignInputField(input, "ForceDeleteAWSJob", _iotForceDeleteAWSJob); err != nil {
			log.Errorf("invalid --force-delete-aws-job: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteOTAUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific version from a software package.
// Note: All package versions must be deleted before deleting the software package.
//
// Requires permission to access the [DeletePackageVersion] action.
//
// [DeletePackageVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeletePackage(cfg aws.Config, client *iot.Client) {
	input := &iot.DeletePackageInput{
		// PackageName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotClientToken) > 0 {
		input.ClientToken = aws.String(_iotClientToken)
	}

	if resp, err := client.DeletePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specific version from a software package.
// Note: If a package version is designated as default, you must remove the
// designation from the software package using the UpdatePackageaction.
func iot_DeletePackageVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.DeletePackageVersionInput{
		// PackageName: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotVersionName) > 0 {
		input.VersionName = aws.String(_iotVersionName)
	}
	if len(_iotClientToken) > 0 {
		input.ClientToken = aws.String(_iotClientToken)
	}

	if resp, err := client.DeletePackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified policy.
// A policy cannot be deleted if it has non-default versions or it is attached to
// any certificate.
//
// To delete a policy, use the DeletePolicyVersion action to delete all non-default versions of the
// policy; use the DetachPolicyaction to detach the policy from any certificate; and then use
// the DeletePolicy action to delete the policy.
//
// When a policy is deleted using DeletePolicy, its default version is deleted
// with it.
//
// Because of the distributed nature of Amazon Web Services, it can take up to
// five minutes after a policy is detached before it's ready to be deleted.
//
// Requires permission to access the [DeletePolicy] action.
//
// [DeletePolicy]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeletePolicy(cfg aws.Config, client *iot.Client) {
	input := &iot.DeletePolicyInput{
		// PolicyName: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}

	if resp, err := client.DeletePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified version of the specified policy. You cannot delete the
// default version of a policy using this action. To delete the default version of
// a policy, use DeletePolicy. To find out which version of a policy is marked as the default
// version, use ListPolicyVersions.
//
// Requires permission to access the [DeletePolicyVersion] action.
//
// [DeletePolicyVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeletePolicyVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.DeletePolicyVersionInput{
		// PolicyName: *string, // Required
		// PolicyVersionId: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotPolicyVersionId) > 0 {
		input.PolicyVersionId = aws.String(_iotPolicyVersionId)
	}

	if resp, err := client.DeletePolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a provisioning template.
// Requires permission to access the [DeleteProvisioningTemplate] action.
//
// [DeleteProvisioningTemplate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteProvisioningTemplate(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteProvisioningTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}

	if resp, err := client.DeleteProvisioningTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a provisioning template version.
// Requires permission to access the [DeleteProvisioningTemplateVersion] action.
//
// [DeleteProvisioningTemplateVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteProvisioningTemplateVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteProvisioningTemplateVersionInput{
		// TemplateName: *string, // Required
		// VersionId: *int32, // Required
	}

	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}
	if len(_iotVersionId) > 0 {
		if err := assignInputField(input, "VersionId", _iotVersionId); err != nil {
			log.Errorf("invalid --version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteProvisioningTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a CA certificate registration code.
// Requires permission to access the [DeleteRegistrationCode] action.
//
// [DeleteRegistrationCode]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteRegistrationCode(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteRegistrationCodeInput{}

	if resp, err := client.DeleteRegistrationCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a role alias
// Requires permission to access the [DeleteRoleAlias] action.
//
// [DeleteRoleAlias]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteRoleAlias(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteRoleAliasInput{
		// RoleAlias: *string, // Required
	}

	if len(_iotRoleAlias) > 0 {
		input.RoleAlias = aws.String(_iotRoleAlias)
	}

	if resp, err := client.DeleteRoleAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a scheduled audit.
// Requires permission to access the [DeleteScheduledAudit] action.
//
// [DeleteScheduledAudit]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteScheduledAudit(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteScheduledAuditInput{
		// ScheduledAuditName: *string, // Required
	}

	if len(_iotScheduledAuditName) > 0 {
		input.ScheduledAuditName = aws.String(_iotScheduledAuditName)
	}

	if resp, err := client.DeleteScheduledAudit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Device Defender security profile.
// Requires permission to access the [DeleteSecurityProfile] action.
//
// [DeleteSecurityProfile]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteSecurityProfile(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteSecurityProfileInput{
		// SecurityProfileName: *string, // Required
	}

	if len(_iotSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_iotSecurityProfileName)
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteSecurityProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a stream.
// Requires permission to access the [DeleteStream] action.
//
// [DeleteStream]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteStream(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteStreamInput{
		// StreamId: *string, // Required
	}

	if len(_iotStreamId) > 0 {
		input.StreamId = aws.String(_iotStreamId)
	}

	if resp, err := client.DeleteStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified thing. Returns successfully with no error if the deletion
// is successful or you specify a thing that doesn't exist.
//
// Requires permission to access the [DeleteThing] action.
//
// [DeleteThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteThing(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteThingInput{
		// ThingName: *string, // Required
	}

	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a thing group.
// Requires permission to access the [DeleteThingGroup] action.
//
// [DeleteThingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteThingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteThingGroupInput{
		// ThingGroupName: *string, // Required
	}

	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteThingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified thing type. You cannot delete a thing type if it has
// things associated with it. To delete a thing type, first mark it as deprecated
// by calling DeprecateThingType, then remove any associated things by calling UpdateThing to change the thing
// type on any associated thing, and finally use DeleteThingTypeto delete the thing type.
//
// Requires permission to access the [DeleteThingType] action.
//
// [DeleteThingType]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteThingType(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteThingTypeInput{
		// ThingTypeName: *string, // Required
	}

	if len(_iotThingTypeName) > 0 {
		input.ThingTypeName = aws.String(_iotThingTypeName)
	}

	if resp, err := client.DeleteThingType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the rule.
// Requires permission to access the [DeleteTopicRule] action.
//
// [DeleteTopicRule]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteTopicRule(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteTopicRuleInput{
		// RuleName: *string, // Required
	}

	if len(_iotRuleName) > 0 {
		input.RuleName = aws.String(_iotRuleName)
	}

	if resp, err := client.DeleteTopicRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a topic rule destination.
// Requires permission to access the [DeleteTopicRuleDestination] action.
//
// [DeleteTopicRuleDestination]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteTopicRuleDestination(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteTopicRuleDestinationInput{
		// Arn: *string, // Required
	}

	if len(_iotArn) > 0 {
		input.Arn = aws.String(_iotArn)
	}

	if resp, err := client.DeleteTopicRuleDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a logging level.
// Requires permission to access the [DeleteV2LoggingLevel] action.
//
// [DeleteV2LoggingLevel]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeleteV2LoggingLevel(cfg aws.Config, client *iot.Client) {
	input := &iot.DeleteV2LoggingLevelInput{
		// TargetName: *string, // Required
		// TargetType: types.LogTargetType, // Required
	}

	if len(_iotTargetName) > 0 {
		input.TargetName = aws.String(_iotTargetName)
	}
	if len(_iotTargetType) > 0 {
		if err := assignInputField(input, "TargetType", _iotTargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteV2LoggingLevel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecates a thing type. You can not associate new things with deprecated thing
// type.
//
// Requires permission to access the [DeprecateThingType] action.
//
// [DeprecateThingType]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DeprecateThingType(cfg aws.Config, client *iot.Client) {
	input := &iot.DeprecateThingTypeInput{
		// ThingTypeName: *string, // Required
	}

	if len(_iotThingTypeName) > 0 {
		input.ThingTypeName = aws.String(_iotThingTypeName)
	}
	if len(_iotUndoDeprecate) > 0 {
		if err := assignInputField(input, "UndoDeprecate", _iotUndoDeprecate); err != nil {
			log.Errorf("invalid --undo-deprecate: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeprecateThingType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the Device Defender audit settings for this account.
// Settings include how audit notifications are sent and which audit checks are
// enabled or disabled.
//
// Requires permission to access the [DescribeAccountAuditConfiguration] action.
//
// [DescribeAccountAuditConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeAccountAuditConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeAccountAuditConfigurationInput{}

	if resp, err := client.DescribeAccountAuditConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a single audit finding. Properties include the reason
// for noncompliance, the severity of the issue, and the start time when the audit
// that returned the finding.
//
// Requires permission to access the [DescribeAuditFinding] action.
//
// [DescribeAuditFinding]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeAuditFinding(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeAuditFindingInput{
		// FindingId: *string, // Required
	}

	if len(_iotFindingId) > 0 {
		input.FindingId = aws.String(_iotFindingId)
	}

	if resp, err := client.DescribeAuditFinding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about an audit mitigation task that is used to apply
// mitigation actions to a set of audit findings. Properties include the actions
// being applied, the audit checks to which they're being applied, the task status,
// and aggregated task statistics.
func iot_DescribeAuditMitigationActionsTask(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeAuditMitigationActionsTaskInput{
		// TaskId: *string, // Required
	}

	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}

	if resp, err := client.DescribeAuditMitigationActionsTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Device Defender audit suppression.
func iot_DescribeAuditSuppression(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeAuditSuppressionInput{
		// CheckName: *string, // Required
		// ResourceIdentifier: *types.ResourceIdentifier, // Required
	}

	if len(_iotCheckName) > 0 {
		input.CheckName = aws.String(_iotCheckName)
	}
	if len(_iotResourceIdentifier) > 0 {
		if err := assignInputField(input, "ResourceIdentifier", _iotResourceIdentifier); err != nil {
			log.Errorf("invalid --resource-identifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeAuditSuppression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Device Defender audit.
// Requires permission to access the [DescribeAuditTask] action.
//
// [DescribeAuditTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeAuditTask(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeAuditTaskInput{
		// TaskId: *string, // Required
	}

	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}

	if resp, err := client.DescribeAuditTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an authorizer.
// Requires permission to access the [DescribeAuthorizer] action.
//
// [DescribeAuthorizer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeAuthorizer(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeAuthorizerInput{
		// AuthorizerName: *string, // Required
	}

	if len(_iotAuthorizerName) > 0 {
		input.AuthorizerName = aws.String(_iotAuthorizerName)
	}

	if resp, err := client.DescribeAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a billing group.
// Requires permission to access the [DescribeBillingGroup] action.
//
// [DescribeBillingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeBillingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeBillingGroupInput{
		// BillingGroupName: *string, // Required
	}

	if len(_iotBillingGroupName) > 0 {
		input.BillingGroupName = aws.String(_iotBillingGroupName)
	}

	if resp, err := client.DescribeBillingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a registered CA certificate.
// Requires permission to access the [DescribeCACertificate] action.
//
// [DescribeCACertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeCACertificate(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeCACertificateInput{
		// CertificateId: *string, // Required
	}

	if len(_iotCertificateId) > 0 {
		input.CertificateId = aws.String(_iotCertificateId)
	}

	if resp, err := client.DescribeCACertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified certificate.
// Requires permission to access the [DescribeCertificate] action.
//
// [DescribeCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeCertificate(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeCertificateInput{
		// CertificateId: *string, // Required
	}

	if len(_iotCertificateId) > 0 {
		input.CertificateId = aws.String(_iotCertificateId)
	}

	if resp, err := client.DescribeCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a certificate provider.
// Requires permission to access the [DescribeCertificateProvider] action.
//
// [DescribeCertificateProvider]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeCertificateProvider(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeCertificateProviderInput{
		// CertificateProviderName: *string, // Required
	}

	if len(_iotCertificateProviderName) > 0 {
		input.CertificateProviderName = aws.String(_iotCertificateProviderName)
	}

	if resp, err := client.DescribeCertificateProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Device Defender detect custom metric.
// Requires permission to access the [DescribeCustomMetric] action.
//
// [DescribeCustomMetric]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeCustomMetric(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeCustomMetricInput{
		// MetricName: *string, // Required
	}

	if len(_iotMetricName) > 0 {
		input.MetricName = aws.String(_iotMetricName)
	}

	if resp, err := client.DescribeCustomMetric(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the default authorizer.
// Requires permission to access the [DescribeDefaultAuthorizer] action.
//
// [DescribeDefaultAuthorizer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeDefaultAuthorizer(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeDefaultAuthorizerInput{}

	if resp, err := client.DescribeDefaultAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Device Defender ML Detect mitigation action.
// Requires permission to access the [DescribeDetectMitigationActionsTask] action.
//
// [DescribeDetectMitigationActionsTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeDetectMitigationActionsTask(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeDetectMitigationActionsTaskInput{
		// TaskId: *string, // Required
	}

	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}

	if resp, err := client.DescribeDetectMitigationActionsTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about a dimension that is defined in your Amazon Web Services
// accounts.
//
// Requires permission to access the [DescribeDimension] action.
//
// [DescribeDimension]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeDimension(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeDimensionInput{
		// Name: *string, // Required
	}

	if len(_iotName) > 0 {
		input.Name = aws.String(_iotName)
	}

	if resp, err := client.DescribeDimension(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets summary information about a domain configuration.
// Requires permission to access the [DescribeDomainConfiguration] action.
//
// [DescribeDomainConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeDomainConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeDomainConfigurationInput{
		// DomainConfigurationName: *string, // Required
	}

	if len(_iotDomainConfigurationName) > 0 {
		input.DomainConfigurationName = aws.String(_iotDomainConfigurationName)
	}

	if resp, err := client.DescribeDomainConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the encryption configuration for resources and data of your Amazon
// Web Services account in Amazon Web Services IoT Core. For more information, see [Data encryption at rest]
// in the Amazon Web Services IoT Core Developer Guide.
//
// [Data encryption at rest]: https://docs.aws.amazon.com/iot/latest/developerguide/encryption-at-rest.html
func iot_DescribeEncryptionConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeEncryptionConfigurationInput{}

	if resp, err := client.DescribeEncryptionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns or creates a unique endpoint specific to the Amazon Web Services
// account making the call.
//
// The first time DescribeEndpoint is called, an endpoint is created. All
// subsequent calls to DescribeEndpoint return the same endpoint.
//
// Requires permission to access the [DescribeEndpoint] action.
//
// [DescribeEndpoint]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeEndpoint(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeEndpointInput{}

	if len(_iotEndpointType) > 0 {
		input.EndpointType = aws.String(_iotEndpointType)
	}

	if resp, err := client.DescribeEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes event configurations.
// Requires permission to access the [DescribeEventConfigurations] action.
//
// [DescribeEventConfigurations]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeEventConfigurations(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeEventConfigurationsInput{}

	if resp, err := client.DescribeEventConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified fleet metric.
// Requires permission to access the [DescribeFleetMetric] action.
//
// [DescribeFleetMetric]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeFleetMetric(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeFleetMetricInput{
		// MetricName: *string, // Required
	}

	if len(_iotMetricName) > 0 {
		input.MetricName = aws.String(_iotMetricName)
	}

	if resp, err := client.DescribeFleetMetric(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a search index.
// Requires permission to access the [DescribeIndex] action.
//
// [DescribeIndex]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeIndex(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeIndexInput{
		// IndexName: *string, // Required
	}

	if len(_iotIndexName) > 0 {
		input.IndexName = aws.String(_iotIndexName)
	}

	if resp, err := client.DescribeIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a job.
// Requires permission to access the [DescribeJob] action.
//
// [DescribeJob]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeJob(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeJobInput{
		// JobId: *string, // Required
	}

	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotBeforeSubstitution) > 0 {
		if err := assignInputField(input, "BeforeSubstitution", _iotBeforeSubstitution); err != nil {
			log.Errorf("invalid --before-substitution: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a job execution.
// Requires permission to access the [DescribeJobExecution] action.
//
// [DescribeJobExecution]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeJobExecution(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeJobExecutionInput{
		// JobId: *string, // Required
		// ThingName: *string, // Required
	}

	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotExecutionNumber) > 0 {
		if err := assignInputField(input, "ExecutionNumber", _iotExecutionNumber); err != nil {
			log.Errorf("invalid --execution-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeJobExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a job template.
func iot_DescribeJobTemplate(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeJobTemplateInput{
		// JobTemplateId: *string, // Required
	}

	if len(_iotJobTemplateId) > 0 {
		input.JobTemplateId = aws.String(_iotJobTemplateId)
	}

	if resp, err := client.DescribeJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// View details of a managed job template.
func iot_DescribeManagedJobTemplate(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeManagedJobTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}
	if len(_iotTemplateVersion) > 0 {
		input.TemplateVersion = aws.String(_iotTemplateVersion)
	}

	if resp, err := client.DescribeManagedJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a mitigation action.
// Requires permission to access the [DescribeMitigationAction] action.
//
// [DescribeMitigationAction]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeMitigationAction(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeMitigationActionInput{
		// ActionName: *string, // Required
	}

	if len(_iotActionName) > 0 {
		input.ActionName = aws.String(_iotActionName)
	}

	if resp, err := client.DescribeMitigationAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a provisioning template.
// Requires permission to access the [DescribeProvisioningTemplate] action.
//
// [DescribeProvisioningTemplate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeProvisioningTemplate(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeProvisioningTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}

	if resp, err := client.DescribeProvisioningTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a provisioning template version.
// Requires permission to access the [DescribeProvisioningTemplateVersion] action.
//
// [DescribeProvisioningTemplateVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeProvisioningTemplateVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeProvisioningTemplateVersionInput{
		// TemplateName: *string, // Required
		// VersionId: *int32, // Required
	}

	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}
	if len(_iotVersionId) > 0 {
		if err := assignInputField(input, "VersionId", _iotVersionId); err != nil {
			log.Errorf("invalid --version-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeProvisioningTemplateVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a role alias.
// Requires permission to access the [DescribeRoleAlias] action.
//
// [DescribeRoleAlias]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeRoleAlias(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeRoleAliasInput{
		// RoleAlias: *string, // Required
	}

	if len(_iotRoleAlias) > 0 {
		input.RoleAlias = aws.String(_iotRoleAlias)
	}

	if resp, err := client.DescribeRoleAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a scheduled audit.
// Requires permission to access the [DescribeScheduledAudit] action.
//
// [DescribeScheduledAudit]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeScheduledAudit(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeScheduledAuditInput{
		// ScheduledAuditName: *string, // Required
	}

	if len(_iotScheduledAuditName) > 0 {
		input.ScheduledAuditName = aws.String(_iotScheduledAuditName)
	}

	if resp, err := client.DescribeScheduledAudit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a Device Defender security profile.
// Requires permission to access the [DescribeSecurityProfile] action.
//
// [DescribeSecurityProfile]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeSecurityProfile(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeSecurityProfileInput{
		// SecurityProfileName: *string, // Required
	}

	if len(_iotSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_iotSecurityProfileName)
	}

	if resp, err := client.DescribeSecurityProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a stream.
// Requires permission to access the [DescribeStream] action.
//
// [DescribeStream]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeStream(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeStreamInput{
		// StreamId: *string, // Required
	}

	if len(_iotStreamId) > 0 {
		input.StreamId = aws.String(_iotStreamId)
	}

	if resp, err := client.DescribeStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified thing.
// Requires permission to access the [DescribeThing] action.
//
// [DescribeThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeThing(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeThingInput{
		// ThingName: *string, // Required
	}

	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}

	if resp, err := client.DescribeThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe a thing group.
// Requires permission to access the [DescribeThingGroup] action.
//
// [DescribeThingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeThingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeThingGroupInput{
		// ThingGroupName: *string, // Required
	}

	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}

	if resp, err := client.DescribeThingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a bulk thing provisioning task.
// Requires permission to access the [DescribeThingRegistrationTask] action.
//
// [DescribeThingRegistrationTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeThingRegistrationTask(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeThingRegistrationTaskInput{
		// TaskId: *string, // Required
	}

	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}

	if resp, err := client.DescribeThingRegistrationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified thing type.
// Requires permission to access the [DescribeThingType] action.
//
// [DescribeThingType]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DescribeThingType(cfg aws.Config, client *iot.Client) {
	input := &iot.DescribeThingTypeInput{
		// ThingTypeName: *string, // Required
	}

	if len(_iotThingTypeName) > 0 {
		input.ThingTypeName = aws.String(_iotThingTypeName)
	}

	if resp, err := client.DescribeThingType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches a policy from the specified target.
// Because of the distributed nature of Amazon Web Services, it can take up to
// five minutes after a policy is detached before it's ready to be deleted.
//
// Requires permission to access the [DetachPolicy] action.
//
// [DetachPolicy]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DetachPolicy(cfg aws.Config, client *iot.Client) {
	input := &iot.DetachPolicyInput{
		// PolicyName: *string, // Required
		// Target: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotTarget) > 0 {
		input.Target = aws.String(_iotTarget)
	}

	if resp, err := client.DetachPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified policy from the specified certificate.
// Note: This action is deprecated and works as expected for backward
// compatibility, but we won't add enhancements. Use DetachPolicyinstead.
//
// Requires permission to access the [DetachPrincipalPolicy] action.
//
// Deprecated: This operation has been deprecated.
//
// [DetachPrincipalPolicy]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DetachPrincipalPolicy(cfg aws.Config, client *iot.Client) {
	input := &iot.DetachPrincipalPolicyInput{
		// PolicyName: *string, // Required
		// Principal: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotPrincipal) > 0 {
		input.Principal = aws.String(_iotPrincipal)
	}

	if resp, err := client.DetachPrincipalPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a Device Defender security profile from a thing group or from
// this account.
//
// Requires permission to access the [DetachSecurityProfile] action.
//
// [DetachSecurityProfile]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DetachSecurityProfile(cfg aws.Config, client *iot.Client) {
	input := &iot.DetachSecurityProfileInput{
		// SecurityProfileName: *string, // Required
		// SecurityProfileTargetArn: *string, // Required
	}

	if len(_iotSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_iotSecurityProfileName)
	}
	if len(_iotSecurityProfileTargetArn) > 0 {
		input.SecurityProfileTargetArn = aws.String(_iotSecurityProfileTargetArn)
	}

	if resp, err := client.DetachSecurityProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches the specified principal from the specified thing. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities.
//
// This call is asynchronous. It might take several seconds for the detachment to
// propagate.
//
// Requires permission to access the [DetachThingPrincipal] action.
//
// [DetachThingPrincipal]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DetachThingPrincipal(cfg aws.Config, client *iot.Client) {
	input := &iot.DetachThingPrincipalInput{
		// Principal: *string, // Required
		// ThingName: *string, // Required
	}

	if len(_iotPrincipal) > 0 {
		input.Principal = aws.String(_iotPrincipal)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}

	if resp, err := client.DetachThingPrincipal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the rule.
// Requires permission to access the [DisableTopicRule] action.
//
// [DisableTopicRule]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DisableTopicRule(cfg aws.Config, client *iot.Client) {
	input := &iot.DisableTopicRuleInput{
		// RuleName: *string, // Required
	}

	if len(_iotRuleName) > 0 {
		input.RuleName = aws.String(_iotRuleName)
	}

	if resp, err := client.DisableTopicRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the selected software bill of materials (SBOM) from a specific
// software package version.
//
// Requires permission to access the [DisassociateSbomWithPackageVersion] action.
//
// [DisassociateSbomWithPackageVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_DisassociateSbomFromPackageVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.DisassociateSbomFromPackageVersionInput{
		// PackageName: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotVersionName) > 0 {
		input.VersionName = aws.String(_iotVersionName)
	}
	if len(_iotClientToken) > 0 {
		input.ClientToken = aws.String(_iotClientToken)
	}

	if resp, err := client.DisassociateSbomFromPackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the rule.
// Requires permission to access the [EnableTopicRule] action.
//
// [EnableTopicRule]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_EnableTopicRule(cfg aws.Config, client *iot.Client) {
	input := &iot.EnableTopicRuleInput{
		// RuleName: *string, // Required
	}

	if len(_iotRuleName) > 0 {
		input.RuleName = aws.String(_iotRuleName)
	}

	if resp, err := client.EnableTopicRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a Device Defender's ML Detect Security Profile training model's
// status.
//
// Requires permission to access the [GetBehaviorModelTrainingSummaries] action.
//
// [GetBehaviorModelTrainingSummaries]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetBehaviorModelTrainingSummaries(cfg aws.Config, client *iot.Client) {
	input := &iot.GetBehaviorModelTrainingSummariesInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_iotSecurityProfileName)
	}

	if disablePaginator() {
		if resp, err := client.GetBehaviorModelTrainingSummaries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.GetBehaviorModelTrainingSummariesOutput
	p := iot.NewGetBehaviorModelTrainingSummariesPaginator(client, input)
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

// Aggregates on indexed data with search queries pertaining to particular fields.
// Requires permission to access the [GetBucketsAggregation] action.
//
// [GetBucketsAggregation]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetBucketsAggregation(cfg aws.Config, client *iot.Client) {
	input := &iot.GetBucketsAggregationInput{
		// AggregationField: *string, // Required
		// BucketsAggregationType: *types.BucketsAggregationType, // Required
		// QueryString: *string, // Required
	}

	if len(_iotAggregationField) > 0 {
		input.AggregationField = aws.String(_iotAggregationField)
	}
	if len(_iotBucketsAggregationType) > 0 {
		if err := assignInputField(input, "BucketsAggregationType", _iotBucketsAggregationType); err != nil {
			log.Errorf("invalid --buckets-aggregation-type: %s", err.Error())
			return
		}
	}
	if len(_iotQueryString) > 0 {
		input.QueryString = aws.String(_iotQueryString)
	}
	if len(_iotIndexName) > 0 {
		input.IndexName = aws.String(_iotIndexName)
	}
	if len(_iotQueryVersion) > 0 {
		input.QueryVersion = aws.String(_iotQueryVersion)
	}

	if resp, err := client.GetBucketsAggregation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the approximate count of unique values that match the query.
// Requires permission to access the [GetCardinality] action.
//
// [GetCardinality]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetCardinality(cfg aws.Config, client *iot.Client) {
	input := &iot.GetCardinalityInput{
		// QueryString: *string, // Required
	}

	if len(_iotQueryString) > 0 {
		input.QueryString = aws.String(_iotQueryString)
	}
	if len(_iotAggregationField) > 0 {
		input.AggregationField = aws.String(_iotAggregationField)
	}
	if len(_iotIndexName) > 0 {
		input.IndexName = aws.String(_iotIndexName)
	}
	if len(_iotQueryVersion) > 0 {
		input.QueryVersion = aws.String(_iotQueryVersion)
	}

	if resp, err := client.GetCardinality(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified command.
func iot_GetCommand(cfg aws.Config, client *iot.Client) {
	input := &iot.GetCommandInput{
		// CommandId: *string, // Required
	}

	if len(_iotCommandId) > 0 {
		input.CommandId = aws.String(_iotCommandId)
	}

	if resp, err := client.GetCommand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specific command execution on a single device.
func iot_GetCommandExecution(cfg aws.Config, client *iot.Client) {
	input := &iot.GetCommandExecutionInput{
		// ExecutionId: *string, // Required
		// TargetArn: *string, // Required
	}

	if len(_iotExecutionId) > 0 {
		input.ExecutionId = aws.String(_iotExecutionId)
	}
	if len(_iotTargetArn) > 0 {
		input.TargetArn = aws.String(_iotTargetArn)
	}
	if len(_iotIncludeResult) > 0 {
		if err := assignInputField(input, "IncludeResult", _iotIncludeResult); err != nil {
			log.Errorf("invalid --include-result: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetCommandExecution(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of the policies that have an effect on the authorization behavior
// of the specified device when it connects to the IoT device gateway.
//
// Requires permission to access the [GetEffectivePolicies] action.
//
// [GetEffectivePolicies]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetEffectivePolicies(cfg aws.Config, client *iot.Client) {
	input := &iot.GetEffectivePoliciesInput{}

	if len(_iotCognitoIdentityPoolId) > 0 {
		input.CognitoIdentityPoolId = aws.String(_iotCognitoIdentityPoolId)
	}
	if len(_iotPrincipal) > 0 {
		input.Principal = aws.String(_iotPrincipal)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}

	if resp, err := client.GetEffectivePolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the indexing configuration.
// Requires permission to access the [GetIndexingConfiguration] action.
//
// [GetIndexingConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetIndexingConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.GetIndexingConfigurationInput{}

	if resp, err := client.GetIndexingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a job document.
// Requires permission to access the [GetJobDocument] action.
//
// [GetJobDocument]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetJobDocument(cfg aws.Config, client *iot.Client) {
	input := &iot.GetJobDocumentInput{
		// JobId: *string, // Required
	}

	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotBeforeSubstitution) > 0 {
		if err := assignInputField(input, "BeforeSubstitution", _iotBeforeSubstitution); err != nil {
			log.Errorf("invalid --before-substitution: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetJobDocument(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the logging options.
// NOTE: use of this command is not recommended. Use GetV2LoggingOptions instead.
//
// Requires permission to access the [GetLoggingOptions] action.
//
// [GetLoggingOptions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetLoggingOptions(cfg aws.Config, client *iot.Client) {
	input := &iot.GetLoggingOptionsInput{}

	if resp, err := client.GetLoggingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an OTA update.
// Requires permission to access the [GetOTAUpdate] action.
//
// [GetOTAUpdate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetOTAUpdate(cfg aws.Config, client *iot.Client) {
	input := &iot.GetOTAUpdateInput{
		// OtaUpdateId: *string, // Required
	}

	if len(_iotOtaUpdateId) > 0 {
		input.OtaUpdateId = aws.String(_iotOtaUpdateId)
	}

	if resp, err := client.GetOTAUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified software package.
// Requires permission to access the [GetPackage] action.
//
// [GetPackage]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetPackage(cfg aws.Config, client *iot.Client) {
	input := &iot.GetPackageInput{
		// PackageName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}

	if resp, err := client.GetPackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified software package's configuration.
// Requires permission to access the [GetPackageConfiguration] action.
//
// [GetPackageConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetPackageConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.GetPackageConfigurationInput{}

	if resp, err := client.GetPackageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified package version.
// Requires permission to access the [GetPackageVersion] action.
//
// [GetPackageVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetPackageVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.GetPackageVersionInput{
		// PackageName: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotVersionName) > 0 {
		input.VersionName = aws.String(_iotVersionName)
	}

	if resp, err := client.GetPackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Groups the aggregated values that match the query into percentile groupings.
// The default percentile groupings are: 1,5,25,50,75,95,99, although you can
// specify your own when you call GetPercentiles . This function returns a value
// for each percentile group specified (or the default percentile groupings). The
// percentile group "1" contains the aggregated field value that occurs in
// approximately one percent of the values that match the query. The percentile
// group "5" contains the aggregated field value that occurs in approximately five
// percent of the values that match the query, and so on. The result is an
// approximation, the more values that match the query, the more accurate the
// percentile values.
//
// Requires permission to access the [GetPercentiles] action.
//
// [GetPercentiles]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetPercentiles(cfg aws.Config, client *iot.Client) {
	input := &iot.GetPercentilesInput{
		// QueryString: *string, // Required
	}

	if len(_iotQueryString) > 0 {
		input.QueryString = aws.String(_iotQueryString)
	}
	if len(_iotAggregationField) > 0 {
		input.AggregationField = aws.String(_iotAggregationField)
	}
	if len(_iotIndexName) > 0 {
		input.IndexName = aws.String(_iotIndexName)
	}
	if len(_iotPercents) > 0 {
		if err := assignInputField(input, "Percents", _iotPercents); err != nil {
			log.Errorf("invalid --percents: %s", err.Error())
			return
		}
	}
	if len(_iotQueryVersion) > 0 {
		input.QueryVersion = aws.String(_iotQueryVersion)
	}

	if resp, err := client.GetPercentiles(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified policy with the policy document of the
// default version.
//
// Requires permission to access the [GetPolicy] action.
//
// [GetPolicy]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetPolicy(cfg aws.Config, client *iot.Client) {
	input := &iot.GetPolicyInput{
		// PolicyName: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified policy version.
// Requires permission to access the [GetPolicyVersion] action.
//
// [GetPolicyVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetPolicyVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.GetPolicyVersionInput{
		// PolicyName: *string, // Required
		// PolicyVersionId: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotPolicyVersionId) > 0 {
		input.PolicyVersionId = aws.String(_iotPolicyVersionId)
	}

	if resp, err := client.GetPolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a registration code used to register a CA certificate with IoT.
// IoT will create a registration code as part of this API call if the
// registration code doesn't exist or has been deleted. If you already have a
// registration code, this API call will return the same registration code.
//
// Requires permission to access the [GetRegistrationCode] action.
//
// [GetRegistrationCode]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetRegistrationCode(cfg aws.Config, client *iot.Client) {
	input := &iot.GetRegistrationCodeInput{}

	if resp, err := client.GetRegistrationCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the count, average, sum, minimum, maximum, sum of squares, variance,
// and standard deviation for the specified aggregated field. If the aggregation
// field is of type String , only the count statistic is returned.
//
// Requires permission to access the [GetStatistics] action.
//
// [GetStatistics]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetStatistics(cfg aws.Config, client *iot.Client) {
	input := &iot.GetStatisticsInput{
		// QueryString: *string, // Required
	}

	if len(_iotQueryString) > 0 {
		input.QueryString = aws.String(_iotQueryString)
	}
	if len(_iotAggregationField) > 0 {
		input.AggregationField = aws.String(_iotAggregationField)
	}
	if len(_iotIndexName) > 0 {
		input.IndexName = aws.String(_iotIndexName)
	}
	if len(_iotQueryVersion) > 0 {
		input.QueryVersion = aws.String(_iotQueryVersion)
	}

	if resp, err := client.GetStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the live connectivity status per device.
func iot_GetThingConnectivityData(cfg aws.Config, client *iot.Client) {
	input := &iot.GetThingConnectivityDataInput{
		// ThingName: *string, // Required
	}

	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}

	if resp, err := client.GetThingConnectivityData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the rule.
// Requires permission to access the [GetTopicRule] action.
//
// [GetTopicRule]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetTopicRule(cfg aws.Config, client *iot.Client) {
	input := &iot.GetTopicRuleInput{
		// RuleName: *string, // Required
	}

	if len(_iotRuleName) > 0 {
		input.RuleName = aws.String(_iotRuleName)
	}

	if resp, err := client.GetTopicRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a topic rule destination.
// Requires permission to access the [GetTopicRuleDestination] action.
//
// [GetTopicRuleDestination]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetTopicRuleDestination(cfg aws.Config, client *iot.Client) {
	input := &iot.GetTopicRuleDestinationInput{
		// Arn: *string, // Required
	}

	if len(_iotArn) > 0 {
		input.Arn = aws.String(_iotArn)
	}

	if resp, err := client.GetTopicRuleDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the fine grained logging options.
// Requires permission to access the [GetV2LoggingOptions] action.
//
// [GetV2LoggingOptions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_GetV2LoggingOptions(cfg aws.Config, client *iot.Client) {
	input := &iot.GetV2LoggingOptionsInput{}

	if len(_iotVerbose) > 0 {
		if err := assignInputField(input, "Verbose", _iotVerbose); err != nil {
			log.Errorf("invalid --verbose: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetV2LoggingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the active violations for a given Device Defender security profile.
// Requires permission to access the [ListActiveViolations] action.
//
// [ListActiveViolations]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListActiveViolations(cfg aws.Config, client *iot.Client) {
	input := &iot.ListActiveViolationsInput{}

	if len(_iotBehaviorCriteriaType) > 0 {
		if err := assignInputField(input, "BehaviorCriteriaType", _iotBehaviorCriteriaType); err != nil {
			log.Errorf("invalid --behavior-criteria-type: %s", err.Error())
			return
		}
	}
	if len(_iotListSuppressedAlerts) > 0 {
		if err := assignInputField(input, "ListSuppressedAlerts", _iotListSuppressedAlerts); err != nil {
			log.Errorf("invalid --list-suppressed-alerts: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_iotSecurityProfileName)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotVerificationState) > 0 {
		if err := assignInputField(input, "VerificationState", _iotVerificationState); err != nil {
			log.Errorf("invalid --verification-state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListActiveViolations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListActiveViolationsOutput
	p := iot.NewListActiveViolationsPaginator(client, input)
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

// Lists the policies attached to the specified thing group.
// Requires permission to access the [ListAttachedPolicies] action.
//
// [ListAttachedPolicies]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListAttachedPolicies(cfg aws.Config, client *iot.Client) {
	input := &iot.ListAttachedPoliciesInput{
		// Target: *string, // Required
	}

	if len(_iotTarget) > 0 {
		input.Target = aws.String(_iotTarget)
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_iotRecursive) > 0 {
		if err := assignInputField(input, "Recursive", _iotRecursive); err != nil {
			log.Errorf("invalid --recursive: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAttachedPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListAttachedPoliciesOutput
	p := iot.NewListAttachedPoliciesPaginator(client, input)
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

// Lists the findings (results) of a Device Defender audit or of the audits
// performed during a specified time period. (Findings are retained for 90 days.)
//
// Requires permission to access the [ListAuditFindings] action.
//
// [ListAuditFindings]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListAuditFindings(cfg aws.Config, client *iot.Client) {
	input := &iot.ListAuditFindingsInput{}

	if len(_iotCheckName) > 0 {
		input.CheckName = aws.String(_iotCheckName)
	}
	if len(_iotEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _iotEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_iotListSuppressedFindings) > 0 {
		if err := assignInputField(input, "ListSuppressedFindings", _iotListSuppressedFindings); err != nil {
			log.Errorf("invalid --list-suppressed-findings: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotResourceIdentifier) > 0 {
		if err := assignInputField(input, "ResourceIdentifier", _iotResourceIdentifier); err != nil {
			log.Errorf("invalid --resource-identifier: %s", err.Error())
			return
		}
	}
	if len(_iotStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _iotStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}

	if disablePaginator() {
		if resp, err := client.ListAuditFindings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListAuditFindingsOutput
	p := iot.NewListAuditFindingsPaginator(client, input)
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

// Gets the status of audit mitigation action tasks that were executed.
// Requires permission to access the [ListAuditMitigationActionsExecutions] action.
//
// [ListAuditMitigationActionsExecutions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListAuditMitigationActionsExecutions(cfg aws.Config, client *iot.Client) {
	input := &iot.ListAuditMitigationActionsExecutionsInput{
		// FindingId: *string, // Required
		// TaskId: *string, // Required
	}

	if len(_iotFindingId) > 0 {
		input.FindingId = aws.String(_iotFindingId)
	}
	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}
	if len(_iotActionStatus) > 0 {
		if err := assignInputField(input, "ActionStatus", _iotActionStatus); err != nil {
			log.Errorf("invalid --action-status: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAuditMitigationActionsExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListAuditMitigationActionsExecutionsOutput
	p := iot.NewListAuditMitigationActionsExecutionsPaginator(client, input)
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

// Gets a list of audit mitigation action tasks that match the specified filters.
// Requires permission to access the [ListAuditMitigationActionsTasks] action.
//
// [ListAuditMitigationActionsTasks]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListAuditMitigationActionsTasks(cfg aws.Config, client *iot.Client) {
	input := &iot.ListAuditMitigationActionsTasksInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_iotEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _iotEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_iotStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _iotStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_iotAuditTaskId) > 0 {
		input.AuditTaskId = aws.String(_iotAuditTaskId)
	}
	if len(_iotFindingId) > 0 {
		input.FindingId = aws.String(_iotFindingId)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotTaskStatus) > 0 {
		if err := assignInputField(input, "TaskStatus", _iotTaskStatus); err != nil {
			log.Errorf("invalid --task-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAuditMitigationActionsTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListAuditMitigationActionsTasksOutput
	p := iot.NewListAuditMitigationActionsTasksPaginator(client, input)
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

// Lists your Device Defender audit listings.
// Requires permission to access the [ListAuditSuppressions] action.
//
// [ListAuditSuppressions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListAuditSuppressions(cfg aws.Config, client *iot.Client) {
	input := &iot.ListAuditSuppressionsInput{}

	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotCheckName) > 0 {
		input.CheckName = aws.String(_iotCheckName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotResourceIdentifier) > 0 {
		if err := assignInputField(input, "ResourceIdentifier", _iotResourceIdentifier); err != nil {
			log.Errorf("invalid --resource-identifier: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAuditSuppressions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListAuditSuppressionsOutput
	p := iot.NewListAuditSuppressionsPaginator(client, input)
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

// Lists the Device Defender audits that have been performed during a given time
// period.
//
// Requires permission to access the [ListAuditTasks] action.
//
// [ListAuditTasks]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListAuditTasks(cfg aws.Config, client *iot.Client) {
	input := &iot.ListAuditTasksInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_iotEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _iotEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_iotStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _iotStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotTaskStatus) > 0 {
		if err := assignInputField(input, "TaskStatus", _iotTaskStatus); err != nil {
			log.Errorf("invalid --task-status: %s", err.Error())
			return
		}
	}
	if len(_iotTaskType) > 0 {
		if err := assignInputField(input, "TaskType", _iotTaskType); err != nil {
			log.Errorf("invalid --task-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAuditTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListAuditTasksOutput
	p := iot.NewListAuditTasksPaginator(client, input)
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

// Lists the authorizers registered in your account.
// Requires permission to access the [ListAuthorizers] action.
//
// [ListAuthorizers]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListAuthorizers(cfg aws.Config, client *iot.Client) {
	input := &iot.ListAuthorizersInput{}

	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAuthorizers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListAuthorizersOutput
	p := iot.NewListAuthorizersPaginator(client, input)
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

// Lists the billing groups you have created.
// Requires permission to access the [ListBillingGroups] action.
//
// [ListBillingGroups]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListBillingGroups(cfg aws.Config, client *iot.Client) {
	input := &iot.ListBillingGroupsInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNamePrefixFilter) > 0 {
		input.NamePrefixFilter = aws.String(_iotNamePrefixFilter)
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBillingGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListBillingGroupsOutput
	p := iot.NewListBillingGroupsPaginator(client, input)
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

// Lists the CA certificates registered for your Amazon Web Services account.
// The results are paginated with a default page size of 25. You can use the
// returned marker to retrieve additional results.
//
// Requires permission to access the [ListCACertificates] action.
//
// [ListCACertificates]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListCACertificates(cfg aws.Config, client *iot.Client) {
	input := &iot.ListCACertificatesInput{}

	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}

	if disablePaginator() {
		if resp, err := client.ListCACertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListCACertificatesOutput
	p := iot.NewListCACertificatesPaginator(client, input)
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

// Lists all your certificate providers in your Amazon Web Services account.
// Requires permission to access the [ListCertificateProviders] action.
//
// [ListCertificateProviders]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListCertificateProviders(cfg aws.Config, client *iot.Client) {
	input := &iot.ListCertificateProvidersInput{}

	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if resp, err := client.ListCertificateProviders(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the certificates registered in your Amazon Web Services account.
// The results are paginated with a default page size of 25. You can use the
// returned marker to retrieve additional results.
//
// Requires permission to access the [ListCertificates] action.
//
// [ListCertificates]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListCertificates(cfg aws.Config, client *iot.Client) {
	input := &iot.ListCertificatesInput{}

	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListCertificatesOutput
	p := iot.NewListCertificatesPaginator(client, input)
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

// List the device certificates signed by the specified CA certificate.
// Requires permission to access the [ListCertificatesByCA] action.
//
// [ListCertificatesByCA]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListCertificatesByCA(cfg aws.Config, client *iot.Client) {
	input := &iot.ListCertificatesByCAInput{
		// CaCertificateId: *string, // Required
	}

	if len(_iotCaCertificateId) > 0 {
		input.CaCertificateId = aws.String(_iotCaCertificateId)
	}
	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCertificatesByCA(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListCertificatesByCAOutput
	p := iot.NewListCertificatesByCAPaginator(client, input)
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

// List all command executions.
// - You must provide only the startedTimeFilter or the completedTimeFilter
// information. If you provide both time filters, the API will generate an error.
// You can use this information to retrieve a list of command executions within a
// specific timeframe.
//
// - You must provide only the commandArn or the thingArn information depending
// on whether you want to list executions for a specific command or an IoT thing.
// If you provide both fields, the API will generate an error.
//
// For more information about considerations for using this API, see [List command executions in your account (CLI)].
//
// [List command executions in your account (CLI)]: https://docs.aws.amazon.com/iot/latest/developerguide/iot-remote-command-execution-start-monitor.html#iot-remote-command-execution-list-cli
func iot_ListCommandExecutions(cfg aws.Config, client *iot.Client) {
	input := &iot.ListCommandExecutionsInput{}

	if len(_iotCommandArn) > 0 {
		input.CommandArn = aws.String(_iotCommandArn)
	}
	if len(_iotCompletedTimeFilter) > 0 {
		if err := assignInputField(input, "CompletedTimeFilter", _iotCompletedTimeFilter); err != nil {
			log.Errorf("invalid --completed-time-filter: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNamespace) > 0 {
		if err := assignInputField(input, "Namespace", _iotNamespace); err != nil {
			log.Errorf("invalid --namespace: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _iotSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_iotStartedTimeFilter) > 0 {
		if err := assignInputField(input, "StartedTimeFilter", _iotStartedTimeFilter); err != nil {
			log.Errorf("invalid --started-time-filter: %s", err.Error())
			return
		}
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_iotTargetArn) > 0 {
		input.TargetArn = aws.String(_iotTargetArn)
	}

	if disablePaginator() {
		if resp, err := client.ListCommandExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListCommandExecutionsOutput
	p := iot.NewListCommandExecutionsPaginator(client, input)
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

// List all commands in your account.
func iot_ListCommands(cfg aws.Config, client *iot.Client) {
	input := &iot.ListCommandsInput{}

	if len(_iotCommandParameterName) > 0 {
		input.CommandParameterName = aws.String(_iotCommandParameterName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNamespace) > 0 {
		if err := assignInputField(input, "Namespace", _iotNamespace); err != nil {
			log.Errorf("invalid --namespace: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _iotSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListCommands(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListCommandsOutput
	p := iot.NewListCommandsPaginator(client, input)
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

// Lists your Device Defender detect custom metrics.
// Requires permission to access the [ListCustomMetrics] action.
//
// [ListCustomMetrics]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListCustomMetrics(cfg aws.Config, client *iot.Client) {
	input := &iot.ListCustomMetricsInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListCustomMetricsOutput
	p := iot.NewListCustomMetricsPaginator(client, input)
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

// Lists mitigation actions executions for a Device Defender ML Detect Security
// Profile.
//
// Requires permission to access the [ListDetectMitigationActionsExecutions] action.
//
// [ListDetectMitigationActionsExecutions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListDetectMitigationActionsExecutions(cfg aws.Config, client *iot.Client) {
	input := &iot.ListDetectMitigationActionsExecutionsInput{}

	if len(_iotEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _iotEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _iotStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotViolationId) > 0 {
		input.ViolationId = aws.String(_iotViolationId)
	}

	if disablePaginator() {
		if resp, err := client.ListDetectMitigationActionsExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListDetectMitigationActionsExecutionsOutput
	p := iot.NewListDetectMitigationActionsExecutionsPaginator(client, input)
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

// List of Device Defender ML Detect mitigation actions tasks.
// Requires permission to access the [ListDetectMitigationActionsTasks] action.
//
// [ListDetectMitigationActionsTasks]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListDetectMitigationActionsTasks(cfg aws.Config, client *iot.Client) {
	input := &iot.ListDetectMitigationActionsTasksInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_iotEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _iotEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_iotStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _iotStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDetectMitigationActionsTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListDetectMitigationActionsTasksOutput
	p := iot.NewListDetectMitigationActionsTasksPaginator(client, input)
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

// List the set of dimensions that are defined for your Amazon Web Services
// accounts.
//
// Requires permission to access the [ListDimensions] action.
//
// [ListDimensions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListDimensions(cfg aws.Config, client *iot.Client) {
	input := &iot.ListDimensionsInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDimensions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListDimensionsOutput
	p := iot.NewListDimensionsPaginator(client, input)
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

// Gets a list of domain configurations for the user. This list is sorted
// alphabetically by domain configuration name.
//
// Requires permission to access the [ListDomainConfigurations] action.
//
// [ListDomainConfigurations]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListDomainConfigurations(cfg aws.Config, client *iot.Client) {
	input := &iot.ListDomainConfigurationsInput{}

	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_iotServiceType) > 0 {
		if err := assignInputField(input, "ServiceType", _iotServiceType); err != nil {
			log.Errorf("invalid --service-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDomainConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListDomainConfigurationsOutput
	p := iot.NewListDomainConfigurationsPaginator(client, input)
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

// Lists all your fleet metrics.
// Requires permission to access the [ListFleetMetrics] action.
//
// [ListFleetMetrics]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListFleetMetrics(cfg aws.Config, client *iot.Client) {
	input := &iot.ListFleetMetricsInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFleetMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListFleetMetricsOutput
	p := iot.NewListFleetMetricsPaginator(client, input)
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

// Lists the search indices.
// Requires permission to access the [ListIndices] action.
//
// [ListIndices]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListIndices(cfg aws.Config, client *iot.Client) {
	input := &iot.ListIndicesInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIndices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListIndicesOutput
	p := iot.NewListIndicesPaginator(client, input)
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

// Lists the job executions for a job.
// Requires permission to access the [ListJobExecutionsForJob] action.
//
// [ListJobExecutionsForJob]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListJobExecutionsForJob(cfg aws.Config, client *iot.Client) {
	input := &iot.ListJobExecutionsForJobInput{
		// JobId: *string, // Required
	}

	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListJobExecutionsForJob(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListJobExecutionsForJobOutput
	p := iot.NewListJobExecutionsForJobPaginator(client, input)
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

// Lists the job executions for the specified thing.
// Requires permission to access the [ListJobExecutionsForThing] action.
//
// [ListJobExecutionsForThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListJobExecutionsForThing(cfg aws.Config, client *iot.Client) {
	input := &iot.ListJobExecutionsForThingInput{
		// ThingName: *string, // Required
	}

	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNamespaceId) > 0 {
		input.NamespaceId = aws.String(_iotNamespaceId)
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListJobExecutionsForThing(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListJobExecutionsForThingOutput
	p := iot.NewListJobExecutionsForThingPaginator(client, input)
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

// Returns a list of job templates.
// Requires permission to access the [ListJobTemplates] action.
//
// [ListJobTemplates]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListJobTemplates(cfg aws.Config, client *iot.Client) {
	input := &iot.ListJobTemplatesInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListJobTemplatesOutput
	p := iot.NewListJobTemplatesPaginator(client, input)
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

// Lists jobs.
// Requires permission to access the [ListJobs] action.
//
// [ListJobs]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListJobs(cfg aws.Config, client *iot.Client) {
	input := &iot.ListJobsInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNamespaceId) > 0 {
		input.NamespaceId = aws.String(_iotNamespaceId)
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_iotTargetSelection) > 0 {
		if err := assignInputField(input, "TargetSelection", _iotTargetSelection); err != nil {
			log.Errorf("invalid --target-selection: %s", err.Error())
			return
		}
	}
	if len(_iotThingGroupId) > 0 {
		input.ThingGroupId = aws.String(_iotThingGroupId)
	}
	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}

	if disablePaginator() {
		if resp, err := client.ListJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListJobsOutput
	p := iot.NewListJobsPaginator(client, input)
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

// Returns a list of managed job templates.
func iot_ListManagedJobTemplates(cfg aws.Config, client *iot.Client) {
	input := &iot.ListManagedJobTemplatesInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedJobTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListManagedJobTemplatesOutput
	p := iot.NewListManagedJobTemplatesPaginator(client, input)
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

// Lists the values reported for an IoT Device Defender metric (device-side
// metric, cloud-side metric, or custom metric) by the given thing during the
// specified time period.
func iot_ListMetricValues(cfg aws.Config, client *iot.Client) {
	input := &iot.ListMetricValuesInput{
		// EndTime: *time.Time, // Required
		// MetricName: *string, // Required
		// StartTime: *time.Time, // Required
		// ThingName: *string, // Required
	}

	if len(_iotEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _iotEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_iotMetricName) > 0 {
		input.MetricName = aws.String(_iotMetricName)
	}
	if len(_iotStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _iotStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotDimensionName) > 0 {
		input.DimensionName = aws.String(_iotDimensionName)
	}
	if len(_iotDimensionValueOperator) > 0 {
		if err := assignInputField(input, "DimensionValueOperator", _iotDimensionValueOperator); err != nil {
			log.Errorf("invalid --dimension-value-operator: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMetricValues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListMetricValuesOutput
	p := iot.NewListMetricValuesPaginator(client, input)
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

// Gets a list of all mitigation actions that match the specified filter criteria.
// Requires permission to access the [ListMitigationActions] action.
//
// [ListMitigationActions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListMitigationActions(cfg aws.Config, client *iot.Client) {
	input := &iot.ListMitigationActionsInput{}

	if len(_iotActionType) > 0 {
		if err := assignInputField(input, "ActionType", _iotActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMitigationActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListMitigationActionsOutput
	p := iot.NewListMitigationActionsPaginator(client, input)
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

// Lists OTA updates.
// Requires permission to access the [ListOTAUpdates] action.
//
// [ListOTAUpdates]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListOTAUpdates(cfg aws.Config, client *iot.Client) {
	input := &iot.ListOTAUpdatesInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotOtaUpdateStatus) > 0 {
		if err := assignInputField(input, "OtaUpdateStatus", _iotOtaUpdateStatus); err != nil {
			log.Errorf("invalid --ota-update-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOTAUpdates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListOTAUpdatesOutput
	p := iot.NewListOTAUpdatesPaginator(client, input)
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

// Lists certificates that are being transferred but not yet accepted.
// Requires permission to access the [ListOutgoingCertificates] action.
//
// [ListOutgoingCertificates]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListOutgoingCertificates(cfg aws.Config, client *iot.Client) {
	input := &iot.ListOutgoingCertificatesInput{}

	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOutgoingCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListOutgoingCertificatesOutput
	p := iot.NewListOutgoingCertificatesPaginator(client, input)
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

// Lists the software package versions associated to the account.
// Requires permission to access the [ListPackageVersions] action.
//
// [ListPackageVersions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListPackageVersions(cfg aws.Config, client *iot.Client) {
	input := &iot.ListPackageVersionsInput{
		// PackageName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPackageVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListPackageVersionsOutput
	p := iot.NewListPackageVersionsPaginator(client, input)
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

// Lists the software packages associated to the account.
// Requires permission to access the [ListPackages] action.
//
// [ListPackages]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListPackages(cfg aws.Config, client *iot.Client) {
	input := &iot.ListPackagesInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListPackagesOutput
	p := iot.NewListPackagesPaginator(client, input)
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

// Lists your policies.
// Requires permission to access the [ListPolicies] action.
//
// [ListPolicies]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListPolicies(cfg aws.Config, client *iot.Client) {
	input := &iot.ListPoliciesInput{}

	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListPoliciesOutput
	p := iot.NewListPoliciesPaginator(client, input)
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

// Lists the principals associated with the specified policy.
// Note: This action is deprecated and works as expected for backward
// compatibility, but we won't add enhancements. Use ListTargetsForPolicyinstead.
//
// Requires permission to access the [ListPolicyPrincipals] action.
//
// Deprecated: This operation has been deprecated.
//
// [ListPolicyPrincipals]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListPolicyPrincipals(cfg aws.Config, client *iot.Client) {
	input := &iot.ListPolicyPrincipalsInput{
		// PolicyName: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyPrincipals(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListPolicyPrincipalsOutput
	p := iot.NewListPolicyPrincipalsPaginator(client, input)
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

// Lists the versions of the specified policy and identifies the default version.
// Requires permission to access the [ListPolicyVersions] action.
//
// [ListPolicyVersions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListPolicyVersions(cfg aws.Config, client *iot.Client) {
	input := &iot.ListPolicyVersionsInput{
		// PolicyName: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}

	if resp, err := client.ListPolicyVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the policies attached to the specified principal. If you use an Cognito
// identity, the ID must be in [AmazonCognito Identity format].
//
// Note: This action is deprecated and works as expected for backward
// compatibility, but we won't add enhancements. Use ListAttachedPoliciesinstead.
//
// Requires permission to access the [ListPrincipalPolicies] action.
//
// Deprecated: This operation has been deprecated.
//
// [ListPrincipalPolicies]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [AmazonCognito Identity format]: https://docs.aws.amazon.com/cognitoidentity/latest/APIReference/API_GetCredentialsForIdentity.html#API_GetCredentialsForIdentity_RequestSyntax
func iot_ListPrincipalPolicies(cfg aws.Config, client *iot.Client) {
	input := &iot.ListPrincipalPoliciesInput{
		// Principal: *string, // Required
	}

	if len(_iotPrincipal) > 0 {
		input.Principal = aws.String(_iotPrincipal)
	}
	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPrincipalPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListPrincipalPoliciesOutput
	p := iot.NewListPrincipalPoliciesPaginator(client, input)
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

// Lists the things associated with the specified principal. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities.
//
// Requires permission to access the [ListPrincipalThings] action.
//
// [ListPrincipalThings]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListPrincipalThings(cfg aws.Config, client *iot.Client) {
	input := &iot.ListPrincipalThingsInput{
		// Principal: *string, // Required
	}

	if len(_iotPrincipal) > 0 {
		input.Principal = aws.String(_iotPrincipal)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPrincipalThings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListPrincipalThingsOutput
	p := iot.NewListPrincipalThingsPaginator(client, input)
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

// Lists the things associated with the specified principal. A principal can be an
// X.509 certificate or an Amazon Cognito ID.
//
// Requires permission to access the [ListPrincipalThings] action.
//
// [ListPrincipalThings]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListPrincipalThingsV2(cfg aws.Config, client *iot.Client) {
	input := &iot.ListPrincipalThingsV2Input{
		// Principal: *string, // Required
	}

	if len(_iotPrincipal) > 0 {
		input.Principal = aws.String(_iotPrincipal)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotThingPrincipalType) > 0 {
		if err := assignInputField(input, "ThingPrincipalType", _iotThingPrincipalType); err != nil {
			log.Errorf("invalid --thing-principal-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPrincipalThingsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListPrincipalThingsV2Output
	p := iot.NewListPrincipalThingsV2Paginator(client, input)
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

// A list of provisioning template versions.
// Requires permission to access the [ListProvisioningTemplateVersions] action.
//
// [ListProvisioningTemplateVersions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListProvisioningTemplateVersions(cfg aws.Config, client *iot.Client) {
	input := &iot.ListProvisioningTemplateVersionsInput{
		// TemplateName: *string, // Required
	}

	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProvisioningTemplateVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListProvisioningTemplateVersionsOutput
	p := iot.NewListProvisioningTemplateVersionsPaginator(client, input)
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

// Lists the provisioning templates in your Amazon Web Services account.
// Requires permission to access the [ListProvisioningTemplates] action.
//
// [ListProvisioningTemplates]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListProvisioningTemplates(cfg aws.Config, client *iot.Client) {
	input := &iot.ListProvisioningTemplatesInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProvisioningTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListProvisioningTemplatesOutput
	p := iot.NewListProvisioningTemplatesPaginator(client, input)
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

// The related resources of an Audit finding. The following resources can be
// returned from calling this API:
//
// - DEVICE_CERTIFICATE
//
// - CA_CERTIFICATE
//
// - IOT_POLICY
//
// - COGNITO_IDENTITY_POOL
//
// - CLIENT_ID
//
// - ACCOUNT_SETTINGS
//
// - ROLE_ALIAS
//
// - IAM_ROLE
//
// - ISSUER_CERTIFICATE
//
// This API is similar to DescribeAuditFinding's [RelatedResources] but provides pagination and is
// not limited to 10 resources. When calling [DescribeAuditFinding]for the intermediate CA revoked for
// active device certificates check, RelatedResources will not be populated. You
// must use this API, ListRelatedResourcesForAuditFinding, to list the
// certificates.
//
// [RelatedResources]: https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeAuditFinding.html
// [DescribeAuditFinding]: https://docs.aws.amazon.com/iot/latest/apireference/API_DescribeAuditFinding.html
func iot_ListRelatedResourcesForAuditFinding(cfg aws.Config, client *iot.Client) {
	input := &iot.ListRelatedResourcesForAuditFindingInput{
		// FindingId: *string, // Required
	}

	if len(_iotFindingId) > 0 {
		input.FindingId = aws.String(_iotFindingId)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRelatedResourcesForAuditFinding(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListRelatedResourcesForAuditFindingOutput
	p := iot.NewListRelatedResourcesForAuditFindingPaginator(client, input)
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

// Lists the role aliases registered in your account.
// Requires permission to access the [ListRoleAliases] action.
//
// [ListRoleAliases]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListRoleAliases(cfg aws.Config, client *iot.Client) {
	input := &iot.ListRoleAliasesInput{}

	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRoleAliases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListRoleAliasesOutput
	p := iot.NewListRoleAliasesPaginator(client, input)
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

// The validation results for all software bill of materials (SBOM) attached to a
// specific software package version.
//
// Requires permission to access the [ListSbomValidationResults] action.
//
// [ListSbomValidationResults]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListSbomValidationResults(cfg aws.Config, client *iot.Client) {
	input := &iot.ListSbomValidationResultsInput{
		// PackageName: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotVersionName) > 0 {
		input.VersionName = aws.String(_iotVersionName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotValidationResult) > 0 {
		if err := assignInputField(input, "ValidationResult", _iotValidationResult); err != nil {
			log.Errorf("invalid --validation-result: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSbomValidationResults(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListSbomValidationResultsOutput
	p := iot.NewListSbomValidationResultsPaginator(client, input)
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

// Lists all of your scheduled audits.
// Requires permission to access the [ListScheduledAudits] action.
//
// [ListScheduledAudits]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListScheduledAudits(cfg aws.Config, client *iot.Client) {
	input := &iot.ListScheduledAuditsInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListScheduledAudits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListScheduledAuditsOutput
	p := iot.NewListScheduledAuditsPaginator(client, input)
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

// Lists the Device Defender security profiles you've created. You can filter
// security profiles by dimension or custom metric.
//
// Requires permission to access the [ListSecurityProfiles] action.
//
// dimensionName and metricName cannot be used in the same request.
//
// [ListSecurityProfiles]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListSecurityProfiles(cfg aws.Config, client *iot.Client) {
	input := &iot.ListSecurityProfilesInput{}

	if len(_iotDimensionName) > 0 {
		input.DimensionName = aws.String(_iotDimensionName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotMetricName) > 0 {
		input.MetricName = aws.String(_iotMetricName)
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListSecurityProfilesOutput
	p := iot.NewListSecurityProfilesPaginator(client, input)
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

// Lists the Device Defender security profiles attached to a target (thing group).
// Requires permission to access the [ListSecurityProfilesForTarget] action.
//
// [ListSecurityProfilesForTarget]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListSecurityProfilesForTarget(cfg aws.Config, client *iot.Client) {
	input := &iot.ListSecurityProfilesForTargetInput{
		// SecurityProfileTargetArn: *string, // Required
	}

	if len(_iotSecurityProfileTargetArn) > 0 {
		input.SecurityProfileTargetArn = aws.String(_iotSecurityProfileTargetArn)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotRecursive) > 0 {
		if err := assignInputField(input, "Recursive", _iotRecursive); err != nil {
			log.Errorf("invalid --recursive: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityProfilesForTarget(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListSecurityProfilesForTargetOutput
	p := iot.NewListSecurityProfilesForTargetPaginator(client, input)
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

// Lists all of the streams in your Amazon Web Services account.
// Requires permission to access the [ListStreams] action.
//
// [ListStreams]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListStreams(cfg aws.Config, client *iot.Client) {
	input := &iot.ListStreamsInput{}

	if len(_iotAscendingOrder) > 0 {
		if err := assignInputField(input, "AscendingOrder", _iotAscendingOrder); err != nil {
			log.Errorf("invalid --ascending-order: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStreams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListStreamsOutput
	p := iot.NewListStreamsPaginator(client, input)
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

// Lists the tags (metadata) you have assigned to the resource.
// Requires permission to access the [ListTagsForResource] action.
//
// [ListTagsForResource]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListTagsForResource(cfg aws.Config, client *iot.Client) {
	input := &iot.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_iotResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotResourceArn)
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
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

	var results []*iot.ListTagsForResourceOutput
	p := iot.NewListTagsForResourcePaginator(client, input)
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

// List targets for the specified policy.
// Requires permission to access the [ListTargetsForPolicy] action.
//
// [ListTargetsForPolicy]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListTargetsForPolicy(cfg aws.Config, client *iot.Client) {
	input := &iot.ListTargetsForPolicyInput{
		// PolicyName: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotMarker) > 0 {
		input.Marker = aws.String(_iotMarker)
	}
	if len(_iotPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _iotPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTargetsForPolicy(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListTargetsForPolicyOutput
	p := iot.NewListTargetsForPolicyPaginator(client, input)
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

// Lists the targets (thing groups) associated with a given Device Defender
// security profile.
//
// Requires permission to access the [ListTargetsForSecurityProfile] action.
//
// [ListTargetsForSecurityProfile]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListTargetsForSecurityProfile(cfg aws.Config, client *iot.Client) {
	input := &iot.ListTargetsForSecurityProfileInput{
		// SecurityProfileName: *string, // Required
	}

	if len(_iotSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_iotSecurityProfileName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTargetsForSecurityProfile(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListTargetsForSecurityProfileOutput
	p := iot.NewListTargetsForSecurityProfilePaginator(client, input)
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

// List the thing groups in your account.
// Requires permission to access the [ListThingGroups] action.
//
// [ListThingGroups]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListThingGroups(cfg aws.Config, client *iot.Client) {
	input := &iot.ListThingGroupsInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNamePrefixFilter) > 0 {
		input.NamePrefixFilter = aws.String(_iotNamePrefixFilter)
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotParentGroup) > 0 {
		input.ParentGroup = aws.String(_iotParentGroup)
	}
	if len(_iotRecursive) > 0 {
		if err := assignInputField(input, "Recursive", _iotRecursive); err != nil {
			log.Errorf("invalid --recursive: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListThingGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListThingGroupsOutput
	p := iot.NewListThingGroupsPaginator(client, input)
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

// List the thing groups to which the specified thing belongs.
// Requires permission to access the [ListThingGroupsForThing] action.
//
// [ListThingGroupsForThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListThingGroupsForThing(cfg aws.Config, client *iot.Client) {
	input := &iot.ListThingGroupsForThingInput{
		// ThingName: *string, // Required
	}

	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListThingGroupsForThing(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListThingGroupsForThingOutput
	p := iot.NewListThingGroupsForThingPaginator(client, input)
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

// Lists the principals associated with the specified thing. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities.
//
// Requires permission to access the [ListThingPrincipals] action.
//
// [ListThingPrincipals]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListThingPrincipals(cfg aws.Config, client *iot.Client) {
	input := &iot.ListThingPrincipalsInput{
		// ThingName: *string, // Required
	}

	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListThingPrincipals(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListThingPrincipalsOutput
	p := iot.NewListThingPrincipalsPaginator(client, input)
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

// Lists the principals associated with the specified thing. A principal can be an
// X.509 certificate or an Amazon Cognito ID.
//
// Requires permission to access the [ListThingPrincipals] action.
//
// [ListThingPrincipals]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListThingPrincipalsV2(cfg aws.Config, client *iot.Client) {
	input := &iot.ListThingPrincipalsV2Input{
		// ThingName: *string, // Required
	}

	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotThingPrincipalType) > 0 {
		if err := assignInputField(input, "ThingPrincipalType", _iotThingPrincipalType); err != nil {
			log.Errorf("invalid --thing-principal-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListThingPrincipalsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListThingPrincipalsV2Output
	p := iot.NewListThingPrincipalsV2Paginator(client, input)
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

// Information about the thing registration tasks.
func iot_ListThingRegistrationTaskReports(cfg aws.Config, client *iot.Client) {
	input := &iot.ListThingRegistrationTaskReportsInput{
		// ReportType: types.ReportType, // Required
		// TaskId: *string, // Required
	}

	if len(_iotReportType) > 0 {
		if err := assignInputField(input, "ReportType", _iotReportType); err != nil {
			log.Errorf("invalid --report-type: %s", err.Error())
			return
		}
	}
	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListThingRegistrationTaskReports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListThingRegistrationTaskReportsOutput
	p := iot.NewListThingRegistrationTaskReportsPaginator(client, input)
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

// List bulk thing provisioning tasks.
// Requires permission to access the [ListThingRegistrationTasks] action.
//
// [ListThingRegistrationTasks]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListThingRegistrationTasks(cfg aws.Config, client *iot.Client) {
	input := &iot.ListThingRegistrationTasksInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListThingRegistrationTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListThingRegistrationTasksOutput
	p := iot.NewListThingRegistrationTasksPaginator(client, input)
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

// Lists the existing thing types.
// Requires permission to access the [ListThingTypes] action.
//
// [ListThingTypes]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListThingTypes(cfg aws.Config, client *iot.Client) {
	input := &iot.ListThingTypesInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotThingTypeName) > 0 {
		input.ThingTypeName = aws.String(_iotThingTypeName)
	}

	if disablePaginator() {
		if resp, err := client.ListThingTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListThingTypesOutput
	p := iot.NewListThingTypesPaginator(client, input)
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

// Lists your things. Use the attributeName and attributeValue parameters to
// filter your things. For example, calling ListThings with attributeName=Color
// and attributeValue=Red retrieves all things in the registry that contain an
// attribute Color with the value Red. For more information, see [List Things]from the Amazon
// Web Services IoT Core Developer Guide.
//
// Requires permission to access the [ListThings] action.
//
// You will not be charged for calling this API if an Access denied error is
// returned. You will also not be charged if no attributes or pagination token was
// provided in request and no pagination token and no results were returned.
//
// [List Things]: https://docs.aws.amazon.com/iot/latest/developerguide/thing-registry.html#list-things
// [ListThings]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListThings(cfg aws.Config, client *iot.Client) {
	input := &iot.ListThingsInput{}

	if len(_iotAttributeName) > 0 {
		input.AttributeName = aws.String(_iotAttributeName)
	}
	if len(_iotAttributeValue) > 0 {
		input.AttributeValue = aws.String(_iotAttributeValue)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotThingTypeName) > 0 {
		input.ThingTypeName = aws.String(_iotThingTypeName)
	}
	if len(_iotUsePrefixAttributeValue) > 0 {
		if err := assignInputField(input, "UsePrefixAttributeValue", _iotUsePrefixAttributeValue); err != nil {
			log.Errorf("invalid --use-prefix-attribute-value: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListThings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListThingsOutput
	p := iot.NewListThingsPaginator(client, input)
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

// Lists the things you have added to the given billing group.
// Requires permission to access the [ListThingsInBillingGroup] action.
//
// [ListThingsInBillingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListThingsInBillingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.ListThingsInBillingGroupInput{
		// BillingGroupName: *string, // Required
	}

	if len(_iotBillingGroupName) > 0 {
		input.BillingGroupName = aws.String(_iotBillingGroupName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListThingsInBillingGroup(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListThingsInBillingGroupOutput
	p := iot.NewListThingsInBillingGroupPaginator(client, input)
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

// Lists the things in the specified group.
// Requires permission to access the [ListThingsInThingGroup] action.
//
// [ListThingsInThingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListThingsInThingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.ListThingsInThingGroupInput{
		// ThingGroupName: *string, // Required
	}

	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotRecursive) > 0 {
		if err := assignInputField(input, "Recursive", _iotRecursive); err != nil {
			log.Errorf("invalid --recursive: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListThingsInThingGroup(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListThingsInThingGroupOutput
	p := iot.NewListThingsInThingGroupPaginator(client, input)
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

// Lists all the topic rule destinations in your Amazon Web Services account.
// Requires permission to access the [ListTopicRuleDestinations] action.
//
// [ListTopicRuleDestinations]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListTopicRuleDestinations(cfg aws.Config, client *iot.Client) {
	input := &iot.ListTopicRuleDestinationsInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTopicRuleDestinations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListTopicRuleDestinationsOutput
	p := iot.NewListTopicRuleDestinationsPaginator(client, input)
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

// Lists the rules for the specific topic.
// Requires permission to access the [ListTopicRules] action.
//
// [ListTopicRules]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListTopicRules(cfg aws.Config, client *iot.Client) {
	input := &iot.ListTopicRulesInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotRuleDisabled) > 0 {
		if err := assignInputField(input, "RuleDisabled", _iotRuleDisabled); err != nil {
			log.Errorf("invalid --rule-disabled: %s", err.Error())
			return
		}
	}
	if len(_iotTopic) > 0 {
		input.Topic = aws.String(_iotTopic)
	}

	if disablePaginator() {
		if resp, err := client.ListTopicRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListTopicRulesOutput
	p := iot.NewListTopicRulesPaginator(client, input)
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

// Lists logging levels.
// Requires permission to access the [ListV2LoggingLevels] action.
//
// [ListV2LoggingLevels]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListV2LoggingLevels(cfg aws.Config, client *iot.Client) {
	input := &iot.ListV2LoggingLevelsInput{}

	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotTargetType) > 0 {
		if err := assignInputField(input, "TargetType", _iotTargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListV2LoggingLevels(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListV2LoggingLevelsOutput
	p := iot.NewListV2LoggingLevelsPaginator(client, input)
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

// Lists the Device Defender security profile violations discovered during the
// given time period. You can use filters to limit the results to those alerts
// issued for a particular security profile, behavior, or thing (device).
//
// Requires permission to access the [ListViolationEvents] action.
//
// [ListViolationEvents]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ListViolationEvents(cfg aws.Config, client *iot.Client) {
	input := &iot.ListViolationEventsInput{
		// EndTime: *time.Time, // Required
		// StartTime: *time.Time, // Required
	}

	if len(_iotEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _iotEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_iotStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _iotStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_iotBehaviorCriteriaType) > 0 {
		if err := assignInputField(input, "BehaviorCriteriaType", _iotBehaviorCriteriaType); err != nil {
			log.Errorf("invalid --behavior-criteria-type: %s", err.Error())
			return
		}
	}
	if len(_iotListSuppressedAlerts) > 0 {
		if err := assignInputField(input, "ListSuppressedAlerts", _iotListSuppressedAlerts); err != nil {
			log.Errorf("invalid --list-suppressed-alerts: %s", err.Error())
			return
		}
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_iotSecurityProfileName)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotVerificationState) > 0 {
		if err := assignInputField(input, "VerificationState", _iotVerificationState); err != nil {
			log.Errorf("invalid --verification-state: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListViolationEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iot.ListViolationEventsOutput
	p := iot.NewListViolationEventsPaginator(client, input)
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

// Set a verification state and provide a description of that verification state
// on a violation (detect alarm).
func iot_PutVerificationStateOnViolation(cfg aws.Config, client *iot.Client) {
	input := &iot.PutVerificationStateOnViolationInput{
		// VerificationState: types.VerificationState, // Required
		// ViolationId: *string, // Required
	}

	if len(_iotVerificationState) > 0 {
		if err := assignInputField(input, "VerificationState", _iotVerificationState); err != nil {
			log.Errorf("invalid --verification-state: %s", err.Error())
			return
		}
	}
	if len(_iotViolationId) > 0 {
		input.ViolationId = aws.String(_iotViolationId)
	}
	if len(_iotVerificationStateDescription) > 0 {
		input.VerificationStateDescription = aws.String(_iotVerificationStateDescription)
	}

	if resp, err := client.PutVerificationStateOnViolation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a CA certificate with Amazon Web Services IoT Core. There is no limit
// to the number of CA certificates you can register in your Amazon Web Services
// account. You can register up to 10 CA certificates with the same CA subject
// field per Amazon Web Services account.
//
// Requires permission to access the [RegisterCACertificate] action.
//
// [RegisterCACertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_RegisterCACertificate(cfg aws.Config, client *iot.Client) {
	input := &iot.RegisterCACertificateInput{
		// CaCertificate: *string, // Required
	}

	if len(_iotCaCertificate) > 0 {
		input.CaCertificate = aws.String(_iotCaCertificate)
	}
	if len(_iotAllowAutoRegistration) > 0 {
		if err := assignInputField(input, "AllowAutoRegistration", _iotAllowAutoRegistration); err != nil {
			log.Errorf("invalid --allow-auto-registration: %s", err.Error())
			return
		}
	}
	if len(_iotCertificateMode) > 0 {
		if err := assignInputField(input, "CertificateMode", _iotCertificateMode); err != nil {
			log.Errorf("invalid --certificate-mode: %s", err.Error())
			return
		}
	}
	if len(_iotRegistrationConfig) > 0 {
		if err := assignInputField(input, "RegistrationConfig", _iotRegistrationConfig); err != nil {
			log.Errorf("invalid --registration-config: %s", err.Error())
			return
		}
	}
	if len(_iotSetAsActive) > 0 {
		if err := assignInputField(input, "SetAsActive", _iotSetAsActive); err != nil {
			log.Errorf("invalid --set-as-active: %s", err.Error())
			return
		}
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotVerificationCertificate) > 0 {
		input.VerificationCertificate = aws.String(_iotVerificationCertificate)
	}

	if resp, err := client.RegisterCACertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers a device certificate with IoT in the same [certificate mode] as the signing CA. If you
// have more than one CA certificate that has the same subject field, you must
// specify the CA certificate that was used to sign the device certificate being
// registered.
//
// Requires permission to access the [RegisterCertificate] action.
//
// [RegisterCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [certificate mode]: https://docs.aws.amazon.com/iot/latest/apireference/API_CertificateDescription.html#iot-Type-CertificateDescription-certificateMode
func iot_RegisterCertificate(cfg aws.Config, client *iot.Client) {
	input := &iot.RegisterCertificateInput{
		// CertificatePem: *string, // Required
	}

	if len(_iotCertificatePem) > 0 {
		input.CertificatePem = aws.String(_iotCertificatePem)
	}
	if len(_iotCaCertificatePem) > 0 {
		input.CaCertificatePem = aws.String(_iotCaCertificatePem)
	}
	if len(_iotSetAsActive) > 0 {
		if err := assignInputField(input, "SetAsActive", _iotSetAsActive); err != nil {
			log.Errorf("invalid --set-as-active: %s", err.Error())
			return
		}
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Register a certificate that does not have a certificate authority (CA). For
// supported certificates, consult [Certificate signing algorithms supported by IoT].
//
// [Certificate signing algorithms supported by IoT]: https://docs.aws.amazon.com/iot/latest/developerguide/x509-client-certs.html#x509-cert-algorithms
func iot_RegisterCertificateWithoutCA(cfg aws.Config, client *iot.Client) {
	input := &iot.RegisterCertificateWithoutCAInput{
		// CertificatePem: *string, // Required
	}

	if len(_iotCertificatePem) > 0 {
		input.CertificatePem = aws.String(_iotCertificatePem)
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterCertificateWithoutCA(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provisions a thing in the device registry. RegisterThing calls other IoT
// control plane APIs. These calls might exceed your account level [IoT Throttling Limits]and cause
// throttle errors. Please contact [Amazon Web Services Customer Support]to raise your throttling limits if necessary.
//
// Requires permission to access the [RegisterThing] action.
//
// [IoT Throttling Limits]: https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_iot
// [RegisterThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [Amazon Web Services Customer Support]: https://console.aws.amazon.com/support/home
func iot_RegisterThing(cfg aws.Config, client *iot.Client) {
	input := &iot.RegisterThingInput{
		// TemplateBody: *string, // Required
	}

	if len(_iotTemplateBody) > 0 {
		input.TemplateBody = aws.String(_iotTemplateBody)
	}
	if len(_iotParameters) > 0 {
		if err := assignInputField(input, "Parameters", _iotParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects a pending certificate transfer. After IoT rejects a certificate
// transfer, the certificate status changes from PENDING_TRANSFER to INACTIVE.
//
// To check for pending certificate transfers, call ListCertificates to enumerate your
// certificates.
//
// This operation can only be called by the transfer destination. After it is
// called, the certificate will be returned to the source's account in the INACTIVE
// state.
//
// Requires permission to access the [RejectCertificateTransfer] action.
//
// [RejectCertificateTransfer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_RejectCertificateTransfer(cfg aws.Config, client *iot.Client) {
	input := &iot.RejectCertificateTransferInput{
		// CertificateId: *string, // Required
	}

	if len(_iotCertificateId) > 0 {
		input.CertificateId = aws.String(_iotCertificateId)
	}
	if len(_iotRejectReason) > 0 {
		input.RejectReason = aws.String(_iotRejectReason)
	}

	if resp, err := client.RejectCertificateTransfer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the given thing from the billing group.
// Requires permission to access the [RemoveThingFromBillingGroup] action.
//
// This call is asynchronous. It might take several seconds for the detachment to
// propagate.
//
// [RemoveThingFromBillingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_RemoveThingFromBillingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.RemoveThingFromBillingGroupInput{}

	if len(_iotBillingGroupArn) > 0 {
		input.BillingGroupArn = aws.String(_iotBillingGroupArn)
	}
	if len(_iotBillingGroupName) > 0 {
		input.BillingGroupName = aws.String(_iotBillingGroupName)
	}
	if len(_iotThingArn) > 0 {
		input.ThingArn = aws.String(_iotThingArn)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}

	if resp, err := client.RemoveThingFromBillingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove the specified thing from the specified group.
// You must specify either a thingGroupArn or a thingGroupName to identify the
// thing group and either a thingArn or a thingName to identify the thing to
// remove from the thing group.
//
// Requires permission to access the [RemoveThingFromThingGroup] action.
//
// [RemoveThingFromThingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_RemoveThingFromThingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.RemoveThingFromThingGroupInput{}

	if len(_iotThingArn) > 0 {
		input.ThingArn = aws.String(_iotThingArn)
	}
	if len(_iotThingGroupArn) > 0 {
		input.ThingGroupArn = aws.String(_iotThingGroupArn)
	}
	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}

	if resp, err := client.RemoveThingFromThingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the rule. You must specify all parameters for the new rule. Creating
// rules is an administrator-level action. Any user who has permission to create
// rules will be able to access data processed by the rule.
//
// Requires permission to access the [ReplaceTopicRule] action.
//
// [ReplaceTopicRule]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ReplaceTopicRule(cfg aws.Config, client *iot.Client) {
	input := &iot.ReplaceTopicRuleInput{
		// RuleName: *string, // Required
		// TopicRulePayload: *types.TopicRulePayload, // Required
	}

	if len(_iotRuleName) > 0 {
		input.RuleName = aws.String(_iotRuleName)
	}
	if len(_iotTopicRulePayload) > 0 {
		if err := assignInputField(input, "TopicRulePayload", _iotTopicRulePayload); err != nil {
			log.Errorf("invalid --topic-rule-payload: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReplaceTopicRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The query search index.
// Requires permission to access the [SearchIndex] action.
//
// [SearchIndex]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_SearchIndex(cfg aws.Config, client *iot.Client) {
	input := &iot.SearchIndexInput{
		// QueryString: *string, // Required
	}

	if len(_iotQueryString) > 0 {
		input.QueryString = aws.String(_iotQueryString)
	}
	if len(_iotIndexName) > 0 {
		input.IndexName = aws.String(_iotIndexName)
	}
	if len(_iotMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotNextToken) > 0 {
		input.NextToken = aws.String(_iotNextToken)
	}
	if len(_iotQueryVersion) > 0 {
		input.QueryVersion = aws.String(_iotQueryVersion)
	}

	if resp, err := client.SearchIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the default authorizer. This will be used if a websocket connection is
// made without specifying an authorizer.
//
// Requires permission to access the [SetDefaultAuthorizer] action.
//
// [SetDefaultAuthorizer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_SetDefaultAuthorizer(cfg aws.Config, client *iot.Client) {
	input := &iot.SetDefaultAuthorizerInput{
		// AuthorizerName: *string, // Required
	}

	if len(_iotAuthorizerName) > 0 {
		input.AuthorizerName = aws.String(_iotAuthorizerName)
	}

	if resp, err := client.SetDefaultAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the specified version of the specified policy as the policy's default
// (operative) version. This action affects all certificates to which the policy is
// attached. To list the principals the policy is attached to, use the ListPrincipalPoliciesaction.
//
// Requires permission to access the [SetDefaultPolicyVersion] action.
//
// [SetDefaultPolicyVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_SetDefaultPolicyVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.SetDefaultPolicyVersionInput{
		// PolicyName: *string, // Required
		// PolicyVersionId: *string, // Required
	}

	if len(_iotPolicyName) > 0 {
		input.PolicyName = aws.String(_iotPolicyName)
	}
	if len(_iotPolicyVersionId) > 0 {
		input.PolicyVersionId = aws.String(_iotPolicyVersionId)
	}

	if resp, err := client.SetDefaultPolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the logging options.
// NOTE: use of this command is not recommended. Use SetV2LoggingOptions instead.
//
// Requires permission to access the [SetLoggingOptions] action.
//
// [SetLoggingOptions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_SetLoggingOptions(cfg aws.Config, client *iot.Client) {
	input := &iot.SetLoggingOptionsInput{
		// LoggingOptionsPayload: *types.LoggingOptionsPayload, // Required
	}

	if len(_iotLoggingOptionsPayload) > 0 {
		if err := assignInputField(input, "LoggingOptionsPayload", _iotLoggingOptionsPayload); err != nil {
			log.Errorf("invalid --logging-options-payload: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetLoggingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the logging level.
// Requires permission to access the [SetV2LoggingLevel] action.
//
// [SetV2LoggingLevel]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_SetV2LoggingLevel(cfg aws.Config, client *iot.Client) {
	input := &iot.SetV2LoggingLevelInput{
		// LogLevel: types.LogLevel, // Required
		// LogTarget: *types.LogTarget, // Required
	}

	if len(_iotLogLevel) > 0 {
		if err := assignInputField(input, "LogLevel", _iotLogLevel); err != nil {
			log.Errorf("invalid --log-level: %s", err.Error())
			return
		}
	}
	if len(_iotLogTarget) > 0 {
		if err := assignInputField(input, "LogTarget", _iotLogTarget); err != nil {
			log.Errorf("invalid --log-target: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetV2LoggingLevel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the logging options for the V2 logging service.
// Requires permission to access the [SetV2LoggingOptions] action.
//
// [SetV2LoggingOptions]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_SetV2LoggingOptions(cfg aws.Config, client *iot.Client) {
	input := &iot.SetV2LoggingOptionsInput{}

	if len(_iotDefaultLogLevel) > 0 {
		if err := assignInputField(input, "DefaultLogLevel", _iotDefaultLogLevel); err != nil {
			log.Errorf("invalid --default-log-level: %s", err.Error())
			return
		}
	}
	if len(_iotDisableAllLogs) > 0 {
		if err := assignInputField(input, "DisableAllLogs", _iotDisableAllLogs); err != nil {
			log.Errorf("invalid --disable-all-logs: %s", err.Error())
			return
		}
	}
	if len(_iotEventConfigurations) > 0 {
		if err := assignInputField(input, "EventConfigurations", _iotEventConfigurations); err != nil {
			log.Errorf("invalid --event-configurations: %s", err.Error())
			return
		}
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}

	if resp, err := client.SetV2LoggingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a task that applies a set of mitigation actions to the specified target.
// Requires permission to access the [StartAuditMitigationActionsTask] action.
//
// [StartAuditMitigationActionsTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_StartAuditMitigationActionsTask(cfg aws.Config, client *iot.Client) {
	input := &iot.StartAuditMitigationActionsTaskInput{
		// AuditCheckToActionsMapping: map[string][]string, // Required
		// ClientRequestToken: *string, // Required
		// Target: *types.AuditMitigationActionsTaskTarget, // Required
		// TaskId: *string, // Required
	}

	if len(_iotAuditCheckToActionsMapping) > 0 {
		if err := assignInputField(input, "AuditCheckToActionsMapping", _iotAuditCheckToActionsMapping); err != nil {
			log.Errorf("invalid --audit-check-to-actions-mapping: %s", err.Error())
			return
		}
	}
	if len(_iotClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotClientRequestToken)
	}
	if len(_iotTarget) > 0 {
		if err := assignInputField(input, "Target", _iotTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}

	if resp, err := client.StartAuditMitigationActionsTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a Device Defender ML Detect mitigation actions task.
// Requires permission to access the [StartDetectMitigationActionsTask] action.
//
// [StartDetectMitigationActionsTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_StartDetectMitigationActionsTask(cfg aws.Config, client *iot.Client) {
	input := &iot.StartDetectMitigationActionsTaskInput{
		// Actions: []string, // Required
		// ClientRequestToken: *string, // Required
		// Target: *types.DetectMitigationActionsTaskTarget, // Required
		// TaskId: *string, // Required
	}

	if len(_iotActions) > 0 {
		input.Actions = append([]string(nil), _iotActions...)
	}
	if len(_iotClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_iotClientRequestToken)
	}
	if len(_iotTarget) > 0 {
		if err := assignInputField(input, "Target", _iotTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}
	if len(_iotIncludeOnlyActiveViolations) > 0 {
		if err := assignInputField(input, "IncludeOnlyActiveViolations", _iotIncludeOnlyActiveViolations); err != nil {
			log.Errorf("invalid --include-only-active-violations: %s", err.Error())
			return
		}
	}
	if len(_iotIncludeSuppressedAlerts) > 0 {
		if err := assignInputField(input, "IncludeSuppressedAlerts", _iotIncludeSuppressedAlerts); err != nil {
			log.Errorf("invalid --include-suppressed-alerts: %s", err.Error())
			return
		}
	}
	if len(_iotViolationEventOccurrenceRange) > 0 {
		if err := assignInputField(input, "ViolationEventOccurrenceRange", _iotViolationEventOccurrenceRange); err != nil {
			log.Errorf("invalid --violation-event-occurrence-range: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartDetectMitigationActionsTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an on-demand Device Defender audit.
// Requires permission to access the [StartOnDemandAuditTask] action.
//
// [StartOnDemandAuditTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_StartOnDemandAuditTask(cfg aws.Config, client *iot.Client) {
	input := &iot.StartOnDemandAuditTaskInput{
		// TargetCheckNames: []string, // Required
	}

	if len(_iotTargetCheckNames) > 0 {
		input.TargetCheckNames = append([]string(nil), _iotTargetCheckNames...)
	}

	if resp, err := client.StartOnDemandAuditTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a bulk thing provisioning task.
// Requires permission to access the [StartThingRegistrationTask] action.
//
// [StartThingRegistrationTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_StartThingRegistrationTask(cfg aws.Config, client *iot.Client) {
	input := &iot.StartThingRegistrationTaskInput{
		// InputFileBucket: *string, // Required
		// InputFileKey: *string, // Required
		// RoleArn: *string, // Required
		// TemplateBody: *string, // Required
	}

	if len(_iotInputFileBucket) > 0 {
		input.InputFileBucket = aws.String(_iotInputFileBucket)
	}
	if len(_iotInputFileKey) > 0 {
		input.InputFileKey = aws.String(_iotInputFileKey)
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}
	if len(_iotTemplateBody) > 0 {
		input.TemplateBody = aws.String(_iotTemplateBody)
	}

	if resp, err := client.StartThingRegistrationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a bulk thing provisioning task.
// Requires permission to access the [StopThingRegistrationTask] action.
//
// [StopThingRegistrationTask]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_StopThingRegistrationTask(cfg aws.Config, client *iot.Client) {
	input := &iot.StopThingRegistrationTaskInput{
		// TaskId: *string, // Required
	}

	if len(_iotTaskId) > 0 {
		input.TaskId = aws.String(_iotTaskId)
	}

	if resp, err := client.StopThingRegistrationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds to or modifies the tags of the given resource. Tags are metadata which can
// be used to manage a resource.
//
// Requires permission to access the [TagResource] action.
//
// [TagResource]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_TagResource(cfg aws.Config, client *iot.Client) {
	input := &iot.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iotResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotResourceArn)
	}
	if len(_iotTags) > 0 {
		if err := assignInputField(input, "Tags", _iotTags); err != nil {
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

// Tests if a specified principal is authorized to perform an IoT action on a
// specified resource. Use this to test and debug the authorization behavior of
// devices that connect to the IoT device gateway.
//
// Requires permission to access the [TestAuthorization] action.
//
// [TestAuthorization]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_TestAuthorization(cfg aws.Config, client *iot.Client) {
	input := &iot.TestAuthorizationInput{
		// AuthInfos: []types.AuthInfo, // Required
	}

	if len(_iotAuthInfos) > 0 {
		if err := assignInputField(input, "AuthInfos", _iotAuthInfos); err != nil {
			log.Errorf("invalid --auth-infos: %s", err.Error())
			return
		}
	}
	if len(_iotClientId) > 0 {
		input.ClientId = aws.String(_iotClientId)
	}
	if len(_iotCognitoIdentityPoolId) > 0 {
		input.CognitoIdentityPoolId = aws.String(_iotCognitoIdentityPoolId)
	}
	if len(_iotPolicyNamesToAdd) > 0 {
		input.PolicyNamesToAdd = append([]string(nil), _iotPolicyNamesToAdd...)
	}
	if len(_iotPolicyNamesToSkip) > 0 {
		input.PolicyNamesToSkip = append([]string(nil), _iotPolicyNamesToSkip...)
	}
	if len(_iotPrincipal) > 0 {
		input.Principal = aws.String(_iotPrincipal)
	}

	if resp, err := client.TestAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tests a custom authorization behavior by invoking a specified custom
// authorizer. Use this to test and debug the custom authorization behavior of
// devices that connect to the IoT device gateway.
//
// Requires permission to access the [TestInvokeAuthorizer] action.
//
// [TestInvokeAuthorizer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_TestInvokeAuthorizer(cfg aws.Config, client *iot.Client) {
	input := &iot.TestInvokeAuthorizerInput{
		// AuthorizerName: *string, // Required
	}

	if len(_iotAuthorizerName) > 0 {
		input.AuthorizerName = aws.String(_iotAuthorizerName)
	}
	if len(_iotHttpContext) > 0 {
		if err := assignInputField(input, "HttpContext", _iotHttpContext); err != nil {
			log.Errorf("invalid --http-context: %s", err.Error())
			return
		}
	}
	if len(_iotMqttContext) > 0 {
		if err := assignInputField(input, "MqttContext", _iotMqttContext); err != nil {
			log.Errorf("invalid --mqtt-context: %s", err.Error())
			return
		}
	}
	if len(_iotTlsContext) > 0 {
		if err := assignInputField(input, "TlsContext", _iotTlsContext); err != nil {
			log.Errorf("invalid --tls-context: %s", err.Error())
			return
		}
	}
	if len(_iotToken) > 0 {
		input.Token = aws.String(_iotToken)
	}
	if len(_iotTokenSignature) > 0 {
		input.TokenSignature = aws.String(_iotTokenSignature)
	}

	if resp, err := client.TestInvokeAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Transfers the specified certificate to the specified Amazon Web Services
// account.
//
// Requires permission to access the [TransferCertificate] action.
//
// You can cancel the transfer until it is accepted by the recipient.
//
// No notification is sent to the transfer destination's account. The caller is
// responsible for notifying the transfer target.
//
// The certificate being transferred must not be in the ACTIVE state. You can use
// the UpdateCertificateaction to deactivate it.
//
// The certificate must not have any policies attached to it. You can use the DetachPolicy
// action to detach them.
//
// Customer managed key behavior: When you use a customer managed key to encrypt
// your data and then transfer the certificate to a customer in a different account
// using the TransferCertificate operation, the certificates will no longer be
// encrypted by their customer managed key configuration. During the transfer
// process, certificates are encrypted using Amazon Web Services IoT Core owned
// keys.
//
// While a certificate is in the PENDING_TRANSFER state, it's always protected by
// Amazon Web Services IoT Core owned keys, regardless of the customer managed key
// configuration of either the source or destination account.
//
// Once the transfer is completed through AcceptCertificateTransfer, RejectCertificateTransfer, or CancelCertificateTransfer, the certificate will be
// protected by the customer managed key configuration of the account that owns the
// certificate after the transfer operation:
//
// - If the transfer is accepted: The certificate is encrypted by the target
// account's customer managed key configuration.
//
// - If the transfer is rejected or cancelled: The certificate is protected by
// the source account's customer managed key configuration.
//
// [TransferCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_TransferCertificate(cfg aws.Config, client *iot.Client) {
	input := &iot.TransferCertificateInput{
		// CertificateId: *string, // Required
		// TargetAwsAccount: *string, // Required
	}

	if len(_iotCertificateId) > 0 {
		input.CertificateId = aws.String(_iotCertificateId)
	}
	if len(_iotTargetAwsAccount) > 0 {
		input.TargetAwsAccount = aws.String(_iotTargetAwsAccount)
	}
	if len(_iotTransferMessage) > 0 {
		input.TransferMessage = aws.String(_iotTransferMessage)
	}

	if resp, err := client.TransferCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the given tags (metadata) from the resource.
// Requires permission to access the [UntagResource] action.
//
// [UntagResource]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UntagResource(cfg aws.Config, client *iot.Client) {
	input := &iot.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iotResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotResourceArn)
	}
	if len(_iotTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iotTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures or reconfigures the Device Defender audit settings for this account.
// Settings include how audit notifications are sent and which audit checks are
// enabled or disabled.
//
// Requires permission to access the [UpdateAccountAuditConfiguration] action.
//
// [UpdateAccountAuditConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateAccountAuditConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateAccountAuditConfigurationInput{}

	if len(_iotAuditCheckConfigurations) > 0 {
		if err := assignInputField(input, "AuditCheckConfigurations", _iotAuditCheckConfigurations); err != nil {
			log.Errorf("invalid --audit-check-configurations: %s", err.Error())
			return
		}
	}
	if len(_iotAuditNotificationTargetConfigurations) > 0 {
		if err := assignInputField(input, "AuditNotificationTargetConfigurations", _iotAuditNotificationTargetConfigurations); err != nil {
			log.Errorf("invalid --audit-notification-target-configurations: %s", err.Error())
			return
		}
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}

	if resp, err := client.UpdateAccountAuditConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Device Defender audit suppression.
func iot_UpdateAuditSuppression(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateAuditSuppressionInput{
		// CheckName: *string, // Required
		// ResourceIdentifier: *types.ResourceIdentifier, // Required
	}

	if len(_iotCheckName) > 0 {
		input.CheckName = aws.String(_iotCheckName)
	}
	if len(_iotResourceIdentifier) > 0 {
		if err := assignInputField(input, "ResourceIdentifier", _iotResourceIdentifier); err != nil {
			log.Errorf("invalid --resource-identifier: %s", err.Error())
			return
		}
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotExpirationDate) > 0 {
		if err := assignInputField(input, "ExpirationDate", _iotExpirationDate); err != nil {
			log.Errorf("invalid --expiration-date: %s", err.Error())
			return
		}
	}
	if len(_iotSuppressIndefinitely) > 0 {
		if err := assignInputField(input, "SuppressIndefinitely", _iotSuppressIndefinitely); err != nil {
			log.Errorf("invalid --suppress-indefinitely: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAuditSuppression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an authorizer.
// Requires permission to access the [UpdateAuthorizer] action.
//
// [UpdateAuthorizer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateAuthorizer(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateAuthorizerInput{
		// AuthorizerName: *string, // Required
	}

	if len(_iotAuthorizerName) > 0 {
		input.AuthorizerName = aws.String(_iotAuthorizerName)
	}
	if len(_iotAuthorizerFunctionArn) > 0 {
		input.AuthorizerFunctionArn = aws.String(_iotAuthorizerFunctionArn)
	}
	if len(_iotEnableCachingForHttp) > 0 {
		if err := assignInputField(input, "EnableCachingForHttp", _iotEnableCachingForHttp); err != nil {
			log.Errorf("invalid --enable-caching-for-http: %s", err.Error())
			return
		}
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_iotTokenKeyName) > 0 {
		input.TokenKeyName = aws.String(_iotTokenKeyName)
	}
	if len(_iotTokenSigningPublicKeys) > 0 {
		if err := assignInputField(input, "TokenSigningPublicKeys", _iotTokenSigningPublicKeys); err != nil {
			log.Errorf("invalid --token-signing-public-keys: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAuthorizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information about the billing group.
// Requires permission to access the [UpdateBillingGroup] action.
//
// [UpdateBillingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateBillingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateBillingGroupInput{
		// BillingGroupName: *string, // Required
		// BillingGroupProperties: *types.BillingGroupProperties, // Required
	}

	if len(_iotBillingGroupName) > 0 {
		input.BillingGroupName = aws.String(_iotBillingGroupName)
	}
	if len(_iotBillingGroupProperties) > 0 {
		if err := assignInputField(input, "BillingGroupProperties", _iotBillingGroupProperties); err != nil {
			log.Errorf("invalid --billing-group-properties: %s", err.Error())
			return
		}
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBillingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a registered CA certificate.
// Requires permission to access the [UpdateCACertificate] action.
//
// [UpdateCACertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateCACertificate(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateCACertificateInput{
		// CertificateId: *string, // Required
	}

	if len(_iotCertificateId) > 0 {
		input.CertificateId = aws.String(_iotCertificateId)
	}
	if len(_iotNewAutoRegistrationStatus) > 0 {
		if err := assignInputField(input, "NewAutoRegistrationStatus", _iotNewAutoRegistrationStatus); err != nil {
			log.Errorf("invalid --new-auto-registration-status: %s", err.Error())
			return
		}
	}
	if len(_iotNewStatus) > 0 {
		if err := assignInputField(input, "NewStatus", _iotNewStatus); err != nil {
			log.Errorf("invalid --new-status: %s", err.Error())
			return
		}
	}
	if len(_iotRegistrationConfig) > 0 {
		if err := assignInputField(input, "RegistrationConfig", _iotRegistrationConfig); err != nil {
			log.Errorf("invalid --registration-config: %s", err.Error())
			return
		}
	}
	if len(_iotRemoveAutoRegistration) > 0 {
		if err := assignInputField(input, "RemoveAutoRegistration", _iotRemoveAutoRegistration); err != nil {
			log.Errorf("invalid --remove-auto-registration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCACertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of the specified certificate. This operation is idempotent.
// Requires permission to access the [UpdateCertificate] action.
//
// Certificates must be in the ACTIVE state to authenticate devices that use a
// certificate to connect to IoT.
//
// Within a few minutes of updating a certificate from the ACTIVE state to any
// other state, IoT disconnects all devices that used that certificate to connect.
// Devices cannot use a certificate that is not in the ACTIVE state to reconnect.
//
// [UpdateCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateCertificate(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateCertificateInput{
		// CertificateId: *string, // Required
		// NewStatus: types.CertificateStatus, // Required
	}

	if len(_iotCertificateId) > 0 {
		input.CertificateId = aws.String(_iotCertificateId)
	}
	if len(_iotNewStatus) > 0 {
		if err := assignInputField(input, "NewStatus", _iotNewStatus); err != nil {
			log.Errorf("invalid --new-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCertificate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a certificate provider.
// Requires permission to access the [UpdateCertificateProvider] action.
//
// [UpdateCertificateProvider]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateCertificateProvider(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateCertificateProviderInput{
		// CertificateProviderName: *string, // Required
	}

	if len(_iotCertificateProviderName) > 0 {
		input.CertificateProviderName = aws.String(_iotCertificateProviderName)
	}
	if len(_iotAccountDefaultForOperations) > 0 {
		if err := assignInputField(input, "AccountDefaultForOperations", _iotAccountDefaultForOperations); err != nil {
			log.Errorf("invalid --account-default-for-operations: %s", err.Error())
			return
		}
	}
	if len(_iotLambdaFunctionArn) > 0 {
		input.LambdaFunctionArn = aws.String(_iotLambdaFunctionArn)
	}

	if resp, err := client.UpdateCertificateProvider(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update information about a command or mark a command for deprecation.
func iot_UpdateCommand(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateCommandInput{
		// CommandId: *string, // Required
	}

	if len(_iotCommandId) > 0 {
		input.CommandId = aws.String(_iotCommandId)
	}
	if len(_iotDeprecated) > 0 {
		if err := assignInputField(input, "Deprecated", _iotDeprecated); err != nil {
			log.Errorf("invalid --deprecated: %s", err.Error())
			return
		}
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotDisplayName) > 0 {
		input.DisplayName = aws.String(_iotDisplayName)
	}

	if resp, err := client.UpdateCommand(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Device Defender detect custom metric.
// Requires permission to access the [UpdateCustomMetric] action.
//
// [UpdateCustomMetric]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateCustomMetric(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateCustomMetricInput{
		// DisplayName: *string, // Required
		// MetricName: *string, // Required
	}

	if len(_iotDisplayName) > 0 {
		input.DisplayName = aws.String(_iotDisplayName)
	}
	if len(_iotMetricName) > 0 {
		input.MetricName = aws.String(_iotMetricName)
	}

	if resp, err := client.UpdateCustomMetric(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the definition for a dimension. You cannot change the type of a
// dimension after it is created (you can delete it and recreate it).
//
// Requires permission to access the [UpdateDimension] action.
//
// [UpdateDimension]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateDimension(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateDimensionInput{
		// Name: *string, // Required
		// StringValues: []string, // Required
	}

	if len(_iotName) > 0 {
		input.Name = aws.String(_iotName)
	}
	if len(_iotStringValues) > 0 {
		input.StringValues = append([]string(nil), _iotStringValues...)
	}

	if resp, err := client.UpdateDimension(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates values stored in the domain configuration. Domain configurations for
// default endpoints can't be updated.
//
// Requires permission to access the [UpdateDomainConfiguration] action.
//
// [UpdateDomainConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateDomainConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateDomainConfigurationInput{
		// DomainConfigurationName: *string, // Required
	}

	if len(_iotDomainConfigurationName) > 0 {
		input.DomainConfigurationName = aws.String(_iotDomainConfigurationName)
	}
	if len(_iotApplicationProtocol) > 0 {
		if err := assignInputField(input, "ApplicationProtocol", _iotApplicationProtocol); err != nil {
			log.Errorf("invalid --application-protocol: %s", err.Error())
			return
		}
	}
	if len(_iotAuthenticationType) > 0 {
		if err := assignInputField(input, "AuthenticationType", _iotAuthenticationType); err != nil {
			log.Errorf("invalid --authentication-type: %s", err.Error())
			return
		}
	}
	if len(_iotAuthorizerConfig) > 0 {
		if err := assignInputField(input, "AuthorizerConfig", _iotAuthorizerConfig); err != nil {
			log.Errorf("invalid --authorizer-config: %s", err.Error())
			return
		}
	}
	if len(_iotClientCertificateConfig) > 0 {
		if err := assignInputField(input, "ClientCertificateConfig", _iotClientCertificateConfig); err != nil {
			log.Errorf("invalid --client-certificate-config: %s", err.Error())
			return
		}
	}
	if len(_iotDomainConfigurationStatus) > 0 {
		if err := assignInputField(input, "DomainConfigurationStatus", _iotDomainConfigurationStatus); err != nil {
			log.Errorf("invalid --domain-configuration-status: %s", err.Error())
			return
		}
	}
	if len(_iotRemoveAuthorizerConfig) > 0 {
		if err := assignInputField(input, "RemoveAuthorizerConfig", _iotRemoveAuthorizerConfig); err != nil {
			log.Errorf("invalid --remove-authorizer-config: %s", err.Error())
			return
		}
	}
	if len(_iotServerCertificateConfig) > 0 {
		if err := assignInputField(input, "ServerCertificateConfig", _iotServerCertificateConfig); err != nil {
			log.Errorf("invalid --server-certificate-config: %s", err.Error())
			return
		}
	}
	if len(_iotTlsConfig) > 0 {
		if err := assignInputField(input, "TlsConfig", _iotTlsConfig); err != nil {
			log.Errorf("invalid --tls-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDomainConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a dynamic thing group.
// Requires permission to access the [UpdateDynamicThingGroup] action.
//
// [UpdateDynamicThingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateDynamicThingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateDynamicThingGroupInput{
		// ThingGroupName: *string, // Required
		// ThingGroupProperties: *types.ThingGroupProperties, // Required
	}

	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}
	if len(_iotThingGroupProperties) > 0 {
		if err := assignInputField(input, "ThingGroupProperties", _iotThingGroupProperties); err != nil {
			log.Errorf("invalid --thing-group-properties: %s", err.Error())
			return
		}
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}
	if len(_iotIndexName) > 0 {
		input.IndexName = aws.String(_iotIndexName)
	}
	if len(_iotQueryString) > 0 {
		input.QueryString = aws.String(_iotQueryString)
	}
	if len(_iotQueryVersion) > 0 {
		input.QueryVersion = aws.String(_iotQueryVersion)
	}

	if resp, err := client.UpdateDynamicThingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the encryption configuration. By default, Amazon Web Services IoT Core
// encrypts your data at rest using Amazon Web Services owned keys. Amazon Web
// Services IoT Core also supports symmetric customer managed keys from Key
// Management Service (KMS). With customer managed keys, you create, own, and
// manage the KMS keys in your Amazon Web Services account.
//
// Before using this API, you must set up permissions for Amazon Web Services IoT
// Core to access KMS. For more information, see [Data encryption at rest]in the Amazon Web Services IoT
// Core Developer Guide.
//
// [Data encryption at rest]: https://docs.aws.amazon.com/iot/latest/developerguide/encryption-at-rest.html
func iot_UpdateEncryptionConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateEncryptionConfigurationInput{
		// EncryptionType: types.EncryptionType, // Required
	}

	if len(_iotEncryptionType) > 0 {
		if err := assignInputField(input, "EncryptionType", _iotEncryptionType); err != nil {
			log.Errorf("invalid --encryption-type: %s", err.Error())
			return
		}
	}
	if len(_iotKmsAccessRoleArn) > 0 {
		input.KmsAccessRoleArn = aws.String(_iotKmsAccessRoleArn)
	}
	if len(_iotKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_iotKmsKeyArn)
	}

	if resp, err := client.UpdateEncryptionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the event configurations.
// Requires permission to access the [UpdateEventConfigurations] action.
//
// [UpdateEventConfigurations]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateEventConfigurations(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateEventConfigurationsInput{}

	if len(_iotEventConfigurations) > 0 {
		if err := assignInputField(input, "EventConfigurations", _iotEventConfigurations); err != nil {
			log.Errorf("invalid --event-configurations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEventConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the data for a fleet metric.
// Requires permission to access the [UpdateFleetMetric] action.
//
// [UpdateFleetMetric]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateFleetMetric(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateFleetMetricInput{
		// IndexName: *string, // Required
		// MetricName: *string, // Required
	}

	if len(_iotIndexName) > 0 {
		input.IndexName = aws.String(_iotIndexName)
	}
	if len(_iotMetricName) > 0 {
		input.MetricName = aws.String(_iotMetricName)
	}
	if len(_iotAggregationField) > 0 {
		input.AggregationField = aws.String(_iotAggregationField)
	}
	if len(_iotAggregationType) > 0 {
		if err := assignInputField(input, "AggregationType", _iotAggregationType); err != nil {
			log.Errorf("invalid --aggregation-type: %s", err.Error())
			return
		}
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}
	if len(_iotPeriod) > 0 {
		if err := assignInputField(input, "Period", _iotPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_iotQueryString) > 0 {
		input.QueryString = aws.String(_iotQueryString)
	}
	if len(_iotQueryVersion) > 0 {
		input.QueryVersion = aws.String(_iotQueryVersion)
	}
	if len(_iotUnit) > 0 {
		if err := assignInputField(input, "Unit", _iotUnit); err != nil {
			log.Errorf("invalid --unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFleetMetric(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the search configuration.
// Requires permission to access the [UpdateIndexingConfiguration] action.
//
// [UpdateIndexingConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateIndexingConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateIndexingConfigurationInput{}

	if len(_iotThingGroupIndexingConfiguration) > 0 {
		if err := assignInputField(input, "ThingGroupIndexingConfiguration", _iotThingGroupIndexingConfiguration); err != nil {
			log.Errorf("invalid --thing-group-indexing-configuration: %s", err.Error())
			return
		}
	}
	if len(_iotThingIndexingConfiguration) > 0 {
		if err := assignInputField(input, "ThingIndexingConfiguration", _iotThingIndexingConfiguration); err != nil {
			log.Errorf("invalid --thing-indexing-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIndexingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates supported fields of the specified job.
// Requires permission to access the [UpdateJob] action.
//
// [UpdateJob]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateJob(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateJobInput{
		// JobId: *string, // Required
	}

	if len(_iotJobId) > 0 {
		input.JobId = aws.String(_iotJobId)
	}
	if len(_iotAbortConfig) > 0 {
		if err := assignInputField(input, "AbortConfig", _iotAbortConfig); err != nil {
			log.Errorf("invalid --abort-config: %s", err.Error())
			return
		}
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotJobExecutionsRetryConfig) > 0 {
		if err := assignInputField(input, "JobExecutionsRetryConfig", _iotJobExecutionsRetryConfig); err != nil {
			log.Errorf("invalid --job-executions-retry-config: %s", err.Error())
			return
		}
	}
	if len(_iotJobExecutionsRolloutConfig) > 0 {
		if err := assignInputField(input, "JobExecutionsRolloutConfig", _iotJobExecutionsRolloutConfig); err != nil {
			log.Errorf("invalid --job-executions-rollout-config: %s", err.Error())
			return
		}
	}
	if len(_iotNamespaceId) > 0 {
		input.NamespaceId = aws.String(_iotNamespaceId)
	}
	if len(_iotPresignedUrlConfig) > 0 {
		if err := assignInputField(input, "PresignedUrlConfig", _iotPresignedUrlConfig); err != nil {
			log.Errorf("invalid --presigned-url-config: %s", err.Error())
			return
		}
	}
	if len(_iotTimeoutConfig) > 0 {
		if err := assignInputField(input, "TimeoutConfig", _iotTimeoutConfig); err != nil {
			log.Errorf("invalid --timeout-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the definition for the specified mitigation action.
// Requires permission to access the [UpdateMitigationAction] action.
//
// [UpdateMitigationAction]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateMitigationAction(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateMitigationActionInput{
		// ActionName: *string, // Required
	}

	if len(_iotActionName) > 0 {
		input.ActionName = aws.String(_iotActionName)
	}
	if len(_iotActionParams) > 0 {
		if err := assignInputField(input, "ActionParams", _iotActionParams); err != nil {
			log.Errorf("invalid --action-params: %s", err.Error())
			return
		}
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}

	if resp, err := client.UpdateMitigationAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the supported fields for a specific software package.
// Requires permission to access the [UpdatePackage] and [GetIndexingConfiguration] actions.
//
// [UpdatePackage]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [GetIndexingConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdatePackage(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdatePackageInput{
		// PackageName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotClientToken) > 0 {
		input.ClientToken = aws.String(_iotClientToken)
	}
	if len(_iotDefaultVersionName) > 0 {
		input.DefaultVersionName = aws.String(_iotDefaultVersionName)
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotUnsetDefaultVersion) > 0 {
		if err := assignInputField(input, "UnsetDefaultVersion", _iotUnsetDefaultVersion); err != nil {
			log.Errorf("invalid --unset-default-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the software package configuration.
// Requires permission to access the [UpdatePackageConfiguration] and [iam:PassRole] actions.
//
// [iam:PassRole]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html
// [UpdatePackageConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdatePackageConfiguration(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdatePackageConfigurationInput{}

	if len(_iotClientToken) > 0 {
		input.ClientToken = aws.String(_iotClientToken)
	}
	if len(_iotVersionUpdateByJobsConfig) > 0 {
		if err := assignInputField(input, "VersionUpdateByJobsConfig", _iotVersionUpdateByJobsConfig); err != nil {
			log.Errorf("invalid --version-update-by-jobs-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePackageConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the supported fields for a specific package version.
// Requires permission to access the [UpdatePackageVersion] and [GetIndexingConfiguration] actions.
//
// [UpdatePackageVersion]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [GetIndexingConfiguration]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdatePackageVersion(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdatePackageVersionInput{
		// PackageName: *string, // Required
		// VersionName: *string, // Required
	}

	if len(_iotPackageName) > 0 {
		input.PackageName = aws.String(_iotPackageName)
	}
	if len(_iotVersionName) > 0 {
		input.VersionName = aws.String(_iotVersionName)
	}
	if len(_iotAction) > 0 {
		if err := assignInputField(input, "Action", _iotAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_iotArtifact) > 0 {
		if err := assignInputField(input, "Artifact", _iotArtifact); err != nil {
			log.Errorf("invalid --artifact: %s", err.Error())
			return
		}
	}
	if len(_iotAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _iotAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_iotClientToken) > 0 {
		input.ClientToken = aws.String(_iotClientToken)
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotRecipe) > 0 {
		input.Recipe = aws.String(_iotRecipe)
	}

	if resp, err := client.UpdatePackageVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a provisioning template.
// Requires permission to access the [UpdateProvisioningTemplate] action.
//
// [UpdateProvisioningTemplate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateProvisioningTemplate(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateProvisioningTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_iotTemplateName) > 0 {
		input.TemplateName = aws.String(_iotTemplateName)
	}
	if len(_iotDefaultVersionId) > 0 {
		if err := assignInputField(input, "DefaultVersionId", _iotDefaultVersionId); err != nil {
			log.Errorf("invalid --default-version-id: %s", err.Error())
			return
		}
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _iotEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_iotPreProvisioningHook) > 0 {
		if err := assignInputField(input, "PreProvisioningHook", _iotPreProvisioningHook); err != nil {
			log.Errorf("invalid --pre-provisioning-hook: %s", err.Error())
			return
		}
	}
	if len(_iotProvisioningRoleArn) > 0 {
		input.ProvisioningRoleArn = aws.String(_iotProvisioningRoleArn)
	}
	if len(_iotRemovePreProvisioningHook) > 0 {
		if err := assignInputField(input, "RemovePreProvisioningHook", _iotRemovePreProvisioningHook); err != nil {
			log.Errorf("invalid --remove-pre-provisioning-hook: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProvisioningTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a role alias.
// Requires permission to access the [UpdateRoleAlias] action.
//
// The value of [credentialDurationSeconds]credentialDurationSeconds must be less than or equal to the
// maximum session duration of the IAM role that the role alias references. For
// more information, see [Modifying a role maximum session duration (Amazon Web Services API)]from the Amazon Web Services Identity and Access
// Management User Guide.
//
// [Modifying a role maximum session duration (Amazon Web Services API)]: https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-managingrole-editing-api.html#roles-modify_max-session-duration-api
// [credentialDurationSeconds]: https://docs.aws.amazon.com/iot/latest/apireference/API_UpdateRoleAlias.html#iot-UpdateRoleAlias-request-credentialDurationSeconds
// [UpdateRoleAlias]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateRoleAlias(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateRoleAliasInput{
		// RoleAlias: *string, // Required
	}

	if len(_iotRoleAlias) > 0 {
		input.RoleAlias = aws.String(_iotRoleAlias)
	}
	if len(_iotCredentialDurationSeconds) > 0 {
		if err := assignInputField(input, "CredentialDurationSeconds", _iotCredentialDurationSeconds); err != nil {
			log.Errorf("invalid --credential-duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}

	if resp, err := client.UpdateRoleAlias(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a scheduled audit, including which checks are performed and how often
// the audit takes place.
//
// Requires permission to access the [UpdateScheduledAudit] action.
//
// [UpdateScheduledAudit]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateScheduledAudit(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateScheduledAuditInput{
		// ScheduledAuditName: *string, // Required
	}

	if len(_iotScheduledAuditName) > 0 {
		input.ScheduledAuditName = aws.String(_iotScheduledAuditName)
	}
	if len(_iotDayOfMonth) > 0 {
		input.DayOfMonth = aws.String(_iotDayOfMonth)
	}
	if len(_iotDayOfWeek) > 0 {
		if err := assignInputField(input, "DayOfWeek", _iotDayOfWeek); err != nil {
			log.Errorf("invalid --day-of-week: %s", err.Error())
			return
		}
	}
	if len(_iotFrequency) > 0 {
		if err := assignInputField(input, "Frequency", _iotFrequency); err != nil {
			log.Errorf("invalid --frequency: %s", err.Error())
			return
		}
	}
	if len(_iotTargetCheckNames) > 0 {
		input.TargetCheckNames = append([]string(nil), _iotTargetCheckNames...)
	}

	if resp, err := client.UpdateScheduledAudit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Device Defender security profile.
// Requires permission to access the [UpdateSecurityProfile] action.
//
// [UpdateSecurityProfile]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateSecurityProfile(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateSecurityProfileInput{
		// SecurityProfileName: *string, // Required
	}

	if len(_iotSecurityProfileName) > 0 {
		input.SecurityProfileName = aws.String(_iotSecurityProfileName)
	}
	if len(_iotAdditionalMetricsToRetain) > 0 {
		input.AdditionalMetricsToRetain = append([]string(nil), _iotAdditionalMetricsToRetain...)
	}
	if len(_iotAdditionalMetricsToRetainV2) > 0 {
		if err := assignInputField(input, "AdditionalMetricsToRetainV2", _iotAdditionalMetricsToRetainV2); err != nil {
			log.Errorf("invalid --additional-metrics-to-retain-v2: %s", err.Error())
			return
		}
	}
	if len(_iotAlertTargets) > 0 {
		if err := assignInputField(input, "AlertTargets", _iotAlertTargets); err != nil {
			log.Errorf("invalid --alert-targets: %s", err.Error())
			return
		}
	}
	if len(_iotBehaviors) > 0 {
		if err := assignInputField(input, "Behaviors", _iotBehaviors); err != nil {
			log.Errorf("invalid --behaviors: %s", err.Error())
			return
		}
	}
	if len(_iotDeleteAdditionalMetricsToRetain) > 0 {
		if err := assignInputField(input, "DeleteAdditionalMetricsToRetain", _iotDeleteAdditionalMetricsToRetain); err != nil {
			log.Errorf("invalid --delete-additional-metrics-to-retain: %s", err.Error())
			return
		}
	}
	if len(_iotDeleteAlertTargets) > 0 {
		if err := assignInputField(input, "DeleteAlertTargets", _iotDeleteAlertTargets); err != nil {
			log.Errorf("invalid --delete-alert-targets: %s", err.Error())
			return
		}
	}
	if len(_iotDeleteBehaviors) > 0 {
		if err := assignInputField(input, "DeleteBehaviors", _iotDeleteBehaviors); err != nil {
			log.Errorf("invalid --delete-behaviors: %s", err.Error())
			return
		}
	}
	if len(_iotDeleteMetricsExportConfig) > 0 {
		if err := assignInputField(input, "DeleteMetricsExportConfig", _iotDeleteMetricsExportConfig); err != nil {
			log.Errorf("invalid --delete-metrics-export-config: %s", err.Error())
			return
		}
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}
	if len(_iotMetricsExportConfig) > 0 {
		if err := assignInputField(input, "MetricsExportConfig", _iotMetricsExportConfig); err != nil {
			log.Errorf("invalid --metrics-export-config: %s", err.Error())
			return
		}
	}
	if len(_iotSecurityProfileDescription) > 0 {
		input.SecurityProfileDescription = aws.String(_iotSecurityProfileDescription)
	}

	if resp, err := client.UpdateSecurityProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing stream. The stream version will be incremented by one.
// Requires permission to access the [UpdateStream] action.
//
// [UpdateStream]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateStream(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateStreamInput{
		// StreamId: *string, // Required
	}

	if len(_iotStreamId) > 0 {
		input.StreamId = aws.String(_iotStreamId)
	}
	if len(_iotDescription) > 0 {
		input.Description = aws.String(_iotDescription)
	}
	if len(_iotFiles) > 0 {
		if err := assignInputField(input, "Files", _iotFiles); err != nil {
			log.Errorf("invalid --files: %s", err.Error())
			return
		}
	}
	if len(_iotRoleArn) > 0 {
		input.RoleArn = aws.String(_iotRoleArn)
	}

	if resp, err := client.UpdateStream(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the data for a thing.
// Requires permission to access the [UpdateThing] action.
//
// [UpdateThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateThing(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateThingInput{
		// ThingName: *string, // Required
	}

	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}
	if len(_iotAttributePayload) > 0 {
		if err := assignInputField(input, "AttributePayload", _iotAttributePayload); err != nil {
			log.Errorf("invalid --attribute-payload: %s", err.Error())
			return
		}
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}
	if len(_iotRemoveThingType) > 0 {
		if err := assignInputField(input, "RemoveThingType", _iotRemoveThingType); err != nil {
			log.Errorf("invalid --remove-thing-type: %s", err.Error())
			return
		}
	}
	if len(_iotThingTypeName) > 0 {
		input.ThingTypeName = aws.String(_iotThingTypeName)
	}

	if resp, err := client.UpdateThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a thing group.
// Requires permission to access the [UpdateThingGroup] action.
//
// [UpdateThingGroup]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateThingGroup(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateThingGroupInput{
		// ThingGroupName: *string, // Required
		// ThingGroupProperties: *types.ThingGroupProperties, // Required
	}

	if len(_iotThingGroupName) > 0 {
		input.ThingGroupName = aws.String(_iotThingGroupName)
	}
	if len(_iotThingGroupProperties) > 0 {
		if err := assignInputField(input, "ThingGroupProperties", _iotThingGroupProperties); err != nil {
			log.Errorf("invalid --thing-group-properties: %s", err.Error())
			return
		}
	}
	if len(_iotExpectedVersion) > 0 {
		if err := assignInputField(input, "ExpectedVersion", _iotExpectedVersion); err != nil {
			log.Errorf("invalid --expected-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateThingGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the groups to which the thing belongs.
// Requires permission to access the [UpdateThingGroupsForThing] action.
//
// [UpdateThingGroupsForThing]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateThingGroupsForThing(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateThingGroupsForThingInput{}

	if len(_iotOverrideDynamicGroups) > 0 {
		if err := assignInputField(input, "OverrideDynamicGroups", _iotOverrideDynamicGroups); err != nil {
			log.Errorf("invalid --override-dynamic-groups: %s", err.Error())
			return
		}
	}
	if len(_iotThingGroupsToAdd) > 0 {
		input.ThingGroupsToAdd = append([]string(nil), _iotThingGroupsToAdd...)
	}
	if len(_iotThingGroupsToRemove) > 0 {
		input.ThingGroupsToRemove = append([]string(nil), _iotThingGroupsToRemove...)
	}
	if len(_iotThingName) > 0 {
		input.ThingName = aws.String(_iotThingName)
	}

	if resp, err := client.UpdateThingGroupsForThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a thing type.
func iot_UpdateThingType(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateThingTypeInput{
		// ThingTypeName: *string, // Required
	}

	if len(_iotThingTypeName) > 0 {
		input.ThingTypeName = aws.String(_iotThingTypeName)
	}
	if len(_iotThingTypeProperties) > 0 {
		if err := assignInputField(input, "ThingTypeProperties", _iotThingTypeProperties); err != nil {
			log.Errorf("invalid --thing-type-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateThingType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a topic rule destination. You use this to change the status, endpoint
// URL, or confirmation URL of the destination.
//
// Requires permission to access the [UpdateTopicRuleDestination] action.
//
// [UpdateTopicRuleDestination]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_UpdateTopicRuleDestination(cfg aws.Config, client *iot.Client) {
	input := &iot.UpdateTopicRuleDestinationInput{
		// Arn: *string, // Required
		// Status: types.TopicRuleDestinationStatus, // Required
	}

	if len(_iotArn) > 0 {
		input.Arn = aws.String(_iotArn)
	}
	if len(_iotStatus) > 0 {
		if err := assignInputField(input, "Status", _iotStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTopicRuleDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates a Device Defender security profile behaviors specification.
// Requires permission to access the [ValidateSecurityProfileBehaviors] action.
//
// [ValidateSecurityProfileBehaviors]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
func iot_ValidateSecurityProfileBehaviors(cfg aws.Config, client *iot.Client) {
	input := &iot.ValidateSecurityProfileBehaviorsInput{
		// Behaviors: []types.Behavior, // Required
	}

	if len(_iotBehaviors) > 0 {
		if err := assignInputField(input, "Behaviors", _iotBehaviors); err != nil {
			log.Errorf("invalid --behaviors: %s", err.Error())
			return
		}
	}

	if resp, err := client.ValidateSecurityProfileBehaviors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iotCmd)
	_iotCmd.Flags().SortFlags = false

	_iotCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_iotCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iotCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_iotCmd.Flags().StringVarP(&_iotAbortConfig, "abort-config", "", "", "Abort Config")
	_iotCmd.Flags().StringVarP(&_iotAccountDefaultForOperations, "account-default-for-operations", "", "", "Account Default For Operations")
	_iotCmd.Flags().StringVarP(&_iotAction, "action", "", "", "Action")
	_iotCmd.Flags().StringVarP(&_iotActionName, "action-name", "", "", "Action Name")
	_iotCmd.Flags().StringVarP(&_iotActionParams, "action-params", "", "", "Action Params")
	_iotCmd.Flags().StringVarP(&_iotActionStatus, "action-status", "", "", "Action Status")
	_iotCmd.Flags().StringVarP(&_iotActionType, "action-type", "", "", "Action Type")
	_iotCmd.Flags().StringSliceVarP(&_iotActions, "actions", "", nil, "Actions")
	_iotCmd.Flags().StringSliceVarP(&_iotAdditionalMetricsToRetain, "additional-metrics-to-retain", "", nil, "Additional Metrics To Retain")
	_iotCmd.Flags().StringVarP(&_iotAdditionalMetricsToRetainV2, "additional-metrics-to-retain-v2", "", "", "Additional Metrics To Retain V2")
	_iotCmd.Flags().StringVarP(&_iotAdditionalParameters, "additional-parameters", "", "", "Additional Parameters")
	_iotCmd.Flags().StringVarP(&_iotAggregationField, "aggregation-field", "", "", "Aggregation Field")
	_iotCmd.Flags().StringVarP(&_iotAggregationType, "aggregation-type", "", "", "Aggregation Type")
	_iotCmd.Flags().StringVarP(&_iotAlertTargets, "alert-targets", "", "", "Alert Targets")
	_iotCmd.Flags().StringVarP(&_iotAllowAutoRegistration, "allow-auto-registration", "", "", "Allow Auto Registration")
	_iotCmd.Flags().StringVarP(&_iotApplicationProtocol, "application-protocol", "", "", "Application Protocol")
	_iotCmd.Flags().StringVarP(&_iotArn, "arn", "", "", "ARN")
	_iotCmd.Flags().StringVarP(&_iotArtifact, "artifact", "", "", "Artifact")
	_iotCmd.Flags().StringVarP(&_iotAscendingOrder, "ascending-order", "", "", "Ascending Order")
	_iotCmd.Flags().StringVarP(&_iotAttributeName, "attribute-name", "", "", "Attribute Name")
	_iotCmd.Flags().StringVarP(&_iotAttributePayload, "attribute-payload", "", "", "Attribute Payload")
	_iotCmd.Flags().StringVarP(&_iotAttributeValue, "attribute-value", "", "", "Attribute Value")
	_iotCmd.Flags().StringVarP(&_iotAttributes, "attributes", "", "", "Attributes")
	_iotCmd.Flags().StringVarP(&_iotAuditCheckConfigurations, "audit-check-configurations", "", "", "Audit Check Configurations")
	_iotCmd.Flags().StringVarP(&_iotAuditCheckToActionsMapping, "audit-check-to-actions-mapping", "", "", "Audit Check To Actions Mapping")
	_iotCmd.Flags().StringVarP(&_iotAuditNotificationTargetConfigurations, "audit-notification-target-configurations", "", "", "Audit Notification Target Configurations")
	_iotCmd.Flags().StringVarP(&_iotAuditTaskId, "audit-task-id", "", "", "Audit Task ID")
	_iotCmd.Flags().StringVarP(&_iotAuthInfos, "auth-infos", "", "", "Auth Infos")
	_iotCmd.Flags().StringVarP(&_iotAuthenticationType, "authentication-type", "", "", "Authentication Type")
	_iotCmd.Flags().StringVarP(&_iotAuthorizerConfig, "authorizer-config", "", "", "Authorizer Config")
	_iotCmd.Flags().StringVarP(&_iotAuthorizerFunctionArn, "authorizer-function-arn", "", "", "Authorizer Function ARN")
	_iotCmd.Flags().StringVarP(&_iotAuthorizerName, "authorizer-name", "", "", "Authorizer Name")
	_iotCmd.Flags().StringVarP(&_iotAwsJobAbortConfig, "aws-job-abort-config", "", "", "AWS Job Abort Config")
	_iotCmd.Flags().StringVarP(&_iotAwsJobExecutionsRolloutConfig, "aws-job-executions-rollout-config", "", "", "AWS Job Executions Rollout Config")
	_iotCmd.Flags().StringVarP(&_iotAwsJobPresignedUrlConfig, "aws-job-presigned-url-config", "", "", "AWS Job Presigned URL Config")
	_iotCmd.Flags().StringVarP(&_iotAwsJobTimeoutConfig, "aws-job-timeout-config", "", "", "AWS Job Timeout Config")
	_iotCmd.Flags().StringVarP(&_iotBeforeSubstitution, "before-substitution", "", "", "Before Substitution")
	_iotCmd.Flags().StringVarP(&_iotBehaviorCriteriaType, "behavior-criteria-type", "", "", "Behavior Criteria Type")
	_iotCmd.Flags().StringVarP(&_iotBehaviors, "behaviors", "", "", "Behaviors")
	_iotCmd.Flags().StringVarP(&_iotBillingGroupArn, "billing-group-arn", "", "", "Billing Group ARN")
	_iotCmd.Flags().StringVarP(&_iotBillingGroupName, "billing-group-name", "", "", "Billing Group Name")
	_iotCmd.Flags().StringVarP(&_iotBillingGroupProperties, "billing-group-properties", "", "", "Billing Group Properties")
	_iotCmd.Flags().StringVarP(&_iotBucketsAggregationType, "buckets-aggregation-type", "", "", "Buckets Aggregation Type")
	_iotCmd.Flags().StringVarP(&_iotCaCertificate, "ca-certificate", "", "", "Ca Certificate")
	_iotCmd.Flags().StringVarP(&_iotCaCertificateId, "ca-certificate-id", "", "", "Ca Certificate ID")
	_iotCmd.Flags().StringVarP(&_iotCaCertificatePem, "ca-certificate-pem", "", "", "Ca Certificate Pem")
	_iotCmd.Flags().StringVarP(&_iotCertificateId, "certificate-id", "", "", "Certificate ID")
	_iotCmd.Flags().StringVarP(&_iotCertificateMode, "certificate-mode", "", "", "Certificate Mode")
	_iotCmd.Flags().StringVarP(&_iotCertificatePem, "certificate-pem", "", "", "Certificate Pem")
	_iotCmd.Flags().StringVarP(&_iotCertificateProviderName, "certificate-provider-name", "", "", "Certificate Provider Name")
	_iotCmd.Flags().StringVarP(&_iotCertificateSigningRequest, "certificate-signing-request", "", "", "Certificate Signing Request")
	_iotCmd.Flags().StringVarP(&_iotCheckName, "check-name", "", "", "Check Name")
	_iotCmd.Flags().StringVarP(&_iotClientCertificateConfig, "client-certificate-config", "", "", "Client Certificate Config")
	_iotCmd.Flags().StringVarP(&_iotClientId, "client-id", "", "", "Client ID")
	_iotCmd.Flags().StringVarP(&_iotClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_iotCmd.Flags().StringVarP(&_iotClientToken, "client-token", "", "", "Client Token")
	_iotCmd.Flags().StringVarP(&_iotCognitoIdentityPoolId, "cognito-identity-pool-id", "", "", "Cognito Identity Pool ID")
	_iotCmd.Flags().StringVarP(&_iotCommandArn, "command-arn", "", "", "Command ARN")
	_iotCmd.Flags().StringVarP(&_iotCommandId, "command-id", "", "", "Command ID")
	_iotCmd.Flags().StringVarP(&_iotCommandParameterName, "command-parameter-name", "", "", "Command Parameter Name")
	_iotCmd.Flags().StringVarP(&_iotComment, "comment", "", "", "Comment")
	_iotCmd.Flags().StringVarP(&_iotCompletedTimeFilter, "completed-time-filter", "", "", "Completed Time Filter")
	_iotCmd.Flags().StringVarP(&_iotConfirmationToken, "confirmation-token", "", "", "Confirmation Token")
	_iotCmd.Flags().StringVarP(&_iotCredentialDurationSeconds, "credential-duration-seconds", "", "", "Credential Duration Seconds")
	_iotCmd.Flags().StringVarP(&_iotDayOfMonth, "day-of-month", "", "", "Day Of Month")
	_iotCmd.Flags().StringVarP(&_iotDayOfWeek, "day-of-week", "", "", "Day Of Week")
	_iotCmd.Flags().StringVarP(&_iotDefaultLogLevel, "default-log-level", "", "", "Default Log Level")
	_iotCmd.Flags().StringVarP(&_iotDefaultVersionId, "default-version-id", "", "", "Default Version ID")
	_iotCmd.Flags().StringVarP(&_iotDefaultVersionName, "default-version-name", "", "", "Default Version Name")
	_iotCmd.Flags().StringVarP(&_iotDeleteAdditionalMetricsToRetain, "delete-additional-metrics-to-retain", "", "", "Delete Additional Metrics To Retain")
	_iotCmd.Flags().StringVarP(&_iotDeleteAlertTargets, "delete-alert-targets", "", "", "Delete Alert Targets")
	_iotCmd.Flags().StringVarP(&_iotDeleteBehaviors, "delete-behaviors", "", "", "Delete Behaviors")
	_iotCmd.Flags().StringVarP(&_iotDeleteMetricsExportConfig, "delete-metrics-export-config", "", "", "Delete Metrics Export Config")
	_iotCmd.Flags().StringVarP(&_iotDeleteScheduledAudits, "delete-scheduled-audits", "", "", "Delete Scheduled Audits")
	_iotCmd.Flags().StringVarP(&_iotDeprecated, "deprecated", "", "", "Deprecated")
	_iotCmd.Flags().StringVarP(&_iotDescription, "description", "", "", "Description")
	_iotCmd.Flags().StringVarP(&_iotDestinationConfiguration, "destination-configuration", "", "", "Destination Configuration")
	_iotCmd.Flags().StringSliceVarP(&_iotDestinationPackageVersions, "destination-package-versions", "", nil, "Destination Package Versions")
	_iotCmd.Flags().StringVarP(&_iotDimensionName, "dimension-name", "", "", "Dimension Name")
	_iotCmd.Flags().StringVarP(&_iotDimensionValueOperator, "dimension-value-operator", "", "", "Dimension Value Operator")
	_iotCmd.Flags().StringVarP(&_iotDisableAllLogs, "disable-all-logs", "", "", "Disable All Logs")
	_iotCmd.Flags().StringVarP(&_iotDisplayName, "display-name", "", "", "Display Name")
	_iotCmd.Flags().StringVarP(&_iotDocument, "document", "", "", "Document")
	_iotCmd.Flags().StringVarP(&_iotDocumentParameters, "document-parameters", "", "", "Document Parameters")
	_iotCmd.Flags().StringVarP(&_iotDocumentSource, "document-source", "", "", "Document Source")
	_iotCmd.Flags().StringVarP(&_iotDomainConfigurationName, "domain-configuration-name", "", "", "Domain Configuration Name")
	_iotCmd.Flags().StringVarP(&_iotDomainConfigurationStatus, "domain-configuration-status", "", "", "Domain Configuration Status")
	_iotCmd.Flags().StringVarP(&_iotDomainName, "domain-name", "", "", "Domain Name")
	_iotCmd.Flags().StringVarP(&_iotEnableCachingForHttp, "enable-caching-for-http", "", "", "Enable Caching For HTTP")
	_iotCmd.Flags().StringVarP(&_iotEnabled, "enabled", "", "", "Enabled")
	_iotCmd.Flags().StringVarP(&_iotEncryptionType, "encryption-type", "", "", "Encryption Type")
	_iotCmd.Flags().StringVarP(&_iotEndTime, "end-time", "", "", "End Time")
	_iotCmd.Flags().StringVarP(&_iotEndpointType, "endpoint-type", "", "", "Endpoint Type")
	_iotCmd.Flags().StringVarP(&_iotEventConfigurations, "event-configurations", "", "", "Event Configurations")
	_iotCmd.Flags().StringVarP(&_iotExecutionId, "execution-id", "", "", "Execution ID")
	_iotCmd.Flags().StringVarP(&_iotExecutionNumber, "execution-number", "", "", "Execution Number")
	_iotCmd.Flags().StringVarP(&_iotExpectedVersion, "expected-version", "", "", "Expected Version")
	_iotCmd.Flags().StringVarP(&_iotExpirationDate, "expiration-date", "", "", "Expiration Date")
	_iotCmd.Flags().StringVarP(&_iotFiles, "files", "", "", "Files")
	_iotCmd.Flags().StringVarP(&_iotFindingId, "finding-id", "", "", "Finding ID")
	_iotCmd.Flags().StringVarP(&_iotForce, "force", "", "", "Force")
	_iotCmd.Flags().StringVarP(&_iotForceDelete, "force-delete", "", "", "Force Delete")
	_iotCmd.Flags().StringVarP(&_iotForceDeleteAWSJob, "force-delete-aws-job", "", "", "Force Delete AWS Job")
	_iotCmd.Flags().StringVarP(&_iotFrequency, "frequency", "", "", "Frequency")
	_iotCmd.Flags().StringVarP(&_iotHttpContext, "http-context", "", "", "HTTP Context")
	_iotCmd.Flags().StringVarP(&_iotIncludeOnlyActiveViolations, "include-only-active-violations", "", "", "Include Only Active Violations")
	_iotCmd.Flags().StringVarP(&_iotIncludeResult, "include-result", "", "", "Include Result")
	_iotCmd.Flags().StringVarP(&_iotIncludeSuppressedAlerts, "include-suppressed-alerts", "", "", "Include Suppressed Alerts")
	_iotCmd.Flags().StringVarP(&_iotIndexName, "index-name", "", "", "Index Name")
	_iotCmd.Flags().StringVarP(&_iotInputFileBucket, "input-file-bucket", "", "", "Input File Bucket")
	_iotCmd.Flags().StringVarP(&_iotInputFileKey, "input-file-key", "", "", "Input File Key")
	_iotCmd.Flags().StringVarP(&_iotJobArn, "job-arn", "", "", "Job ARN")
	_iotCmd.Flags().StringVarP(&_iotJobExecutionsRetryConfig, "job-executions-retry-config", "", "", "Job Executions Retry Config")
	_iotCmd.Flags().StringVarP(&_iotJobExecutionsRolloutConfig, "job-executions-rollout-config", "", "", "Job Executions Rollout Config")
	_iotCmd.Flags().StringVarP(&_iotJobId, "job-id", "", "", "Job ID")
	_iotCmd.Flags().StringVarP(&_iotJobTemplateArn, "job-template-arn", "", "", "Job Template ARN")
	_iotCmd.Flags().StringVarP(&_iotJobTemplateId, "job-template-id", "", "", "Job Template ID")
	_iotCmd.Flags().StringVarP(&_iotKmsAccessRoleArn, "kms-access-role-arn", "", "", "KMS Access Role ARN")
	_iotCmd.Flags().StringVarP(&_iotKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_iotCmd.Flags().StringVarP(&_iotLambdaFunctionArn, "lambda-function-arn", "", "", "Lambda Function ARN")
	_iotCmd.Flags().StringVarP(&_iotListSuppressedAlerts, "list-suppressed-alerts", "", "", "List Suppressed Alerts")
	_iotCmd.Flags().StringVarP(&_iotListSuppressedFindings, "list-suppressed-findings", "", "", "List Suppressed Findings")
	_iotCmd.Flags().StringVarP(&_iotLogLevel, "log-level", "", "", "Log Level")
	_iotCmd.Flags().StringVarP(&_iotLogTarget, "log-target", "", "", "Log Target")
	_iotCmd.Flags().StringVarP(&_iotLoggingOptionsPayload, "logging-options-payload", "", "", "Logging Options Payload")
	_iotCmd.Flags().StringVarP(&_iotMaintenanceWindows, "maintenance-windows", "", "", "Maintenance Windows")
	_iotCmd.Flags().StringVarP(&_iotMandatoryParameters, "mandatory-parameters", "", "", "Mandatory Parameters")
	_iotCmd.Flags().StringVarP(&_iotMarker, "marker", "", "", "Marker")
	_iotCmd.Flags().StringVarP(&_iotMaxResults, "max-results", "", "", "Max Results")
	_iotCmd.Flags().StringVarP(&_iotMetricName, "metric-name", "", "", "Metric Name")
	_iotCmd.Flags().StringVarP(&_iotMetricType, "metric-type", "", "", "Metric Type")
	_iotCmd.Flags().StringVarP(&_iotMetricsExportConfig, "metrics-export-config", "", "", "Metrics Export Config")
	_iotCmd.Flags().StringVarP(&_iotMqttContext, "mqtt-context", "", "", "Mqtt Context")
	_iotCmd.Flags().StringVarP(&_iotName, "name", "", "", "Name")
	_iotCmd.Flags().StringVarP(&_iotNamePrefixFilter, "name-prefix-filter", "", "", "Name Prefix Filter")
	_iotCmd.Flags().StringVarP(&_iotNamespace, "namespace", "", "", "Namespace")
	_iotCmd.Flags().StringVarP(&_iotNamespaceId, "namespace-id", "", "", "Namespace ID")
	_iotCmd.Flags().StringVarP(&_iotNewAutoRegistrationStatus, "new-auto-registration-status", "", "", "New Auto Registration Status")
	_iotCmd.Flags().StringVarP(&_iotNewStatus, "new-status", "", "", "New Status")
	_iotCmd.Flags().StringVarP(&_iotNextToken, "next-token", "", "", "Next Token")
	_iotCmd.Flags().StringVarP(&_iotOtaUpdateId, "ota-update-id", "", "", "Ota Update ID")
	_iotCmd.Flags().StringVarP(&_iotOtaUpdateStatus, "ota-update-status", "", "", "Ota Update Status")
	_iotCmd.Flags().StringVarP(&_iotOverrideDynamicGroups, "override-dynamic-groups", "", "", "Override Dynamic Groups")
	_iotCmd.Flags().StringVarP(&_iotPackageName, "package-name", "", "", "Package Name")
	_iotCmd.Flags().StringVarP(&_iotPageSize, "page-size", "", "", "Page Size")
	_iotCmd.Flags().StringVarP(&_iotParameters, "parameters", "", "", "Parameters")
	_iotCmd.Flags().StringVarP(&_iotParentGroup, "parent-group", "", "", "Parent Group")
	_iotCmd.Flags().StringVarP(&_iotParentGroupName, "parent-group-name", "", "", "Parent Group Name")
	_iotCmd.Flags().StringVarP(&_iotPayload, "payload", "", "", "Payload")
	_iotCmd.Flags().StringVarP(&_iotPayloadTemplate, "payload-template", "", "", "Payload Template")
	_iotCmd.Flags().StringVarP(&_iotPercents, "percents", "", "", "Percents")
	_iotCmd.Flags().StringVarP(&_iotPeriod, "period", "", "", "Period")
	_iotCmd.Flags().StringVarP(&_iotPolicyDocument, "policy-document", "", "", "Policy Document")
	_iotCmd.Flags().StringVarP(&_iotPolicyName, "policy-name", "", "", "Policy Name")
	_iotCmd.Flags().StringSliceVarP(&_iotPolicyNamesToAdd, "policy-names-to-add", "", nil, "Policy Names To Add")
	_iotCmd.Flags().StringSliceVarP(&_iotPolicyNamesToSkip, "policy-names-to-skip", "", nil, "Policy Names To Skip")
	_iotCmd.Flags().StringVarP(&_iotPolicyVersionId, "policy-version-id", "", "", "Policy Version ID")
	_iotCmd.Flags().StringVarP(&_iotPreProvisioningHook, "pre-provisioning-hook", "", "", "Pre Provisioning Hook")
	_iotCmd.Flags().StringVarP(&_iotPreprocessor, "preprocessor", "", "", "Preprocessor")
	_iotCmd.Flags().StringVarP(&_iotPresignedUrlConfig, "presigned-url-config", "", "", "Presigned URL Config")
	_iotCmd.Flags().StringVarP(&_iotPrincipal, "principal", "", "", "Principal")
	_iotCmd.Flags().StringVarP(&_iotProtocols, "protocols", "", "", "Protocols")
	_iotCmd.Flags().StringVarP(&_iotProvisioningRoleArn, "provisioning-role-arn", "", "", "Provisioning Role ARN")
	_iotCmd.Flags().StringVarP(&_iotQueryString, "query-string", "", "", "Query String")
	_iotCmd.Flags().StringVarP(&_iotQueryVersion, "query-version", "", "", "Query Version")
	_iotCmd.Flags().StringVarP(&_iotReasonCode, "reason-code", "", "", "Reason Code")
	_iotCmd.Flags().StringVarP(&_iotRecipe, "recipe", "", "", "Recipe")
	_iotCmd.Flags().StringVarP(&_iotRecursive, "recursive", "", "", "Recursive")
	_iotCmd.Flags().StringVarP(&_iotRegistrationConfig, "registration-config", "", "", "Registration Config")
	_iotCmd.Flags().StringVarP(&_iotRejectReason, "reject-reason", "", "", "Reject Reason")
	_iotCmd.Flags().StringVarP(&_iotRemoveAuthorizerConfig, "remove-authorizer-config", "", "", "Remove Authorizer Config")
	_iotCmd.Flags().StringVarP(&_iotRemoveAutoRegistration, "remove-auto-registration", "", "", "Remove Auto Registration")
	_iotCmd.Flags().StringVarP(&_iotRemovePreProvisioningHook, "remove-pre-provisioning-hook", "", "", "Remove Pre Provisioning Hook")
	_iotCmd.Flags().StringVarP(&_iotRemoveThingType, "remove-thing-type", "", "", "Remove Thing Type")
	_iotCmd.Flags().StringVarP(&_iotReportType, "report-type", "", "", "Report Type")
	_iotCmd.Flags().StringVarP(&_iotResourceArn, "resource-arn", "", "", "Resource ARN")
	_iotCmd.Flags().StringVarP(&_iotResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_iotCmd.Flags().StringVarP(&_iotRoleAlias, "role-alias", "", "", "Role Alias")
	_iotCmd.Flags().StringVarP(&_iotRoleArn, "role-arn", "", "", "Role ARN")
	_iotCmd.Flags().StringVarP(&_iotRuleDisabled, "rule-disabled", "", "", "Rule Disabled")
	_iotCmd.Flags().StringVarP(&_iotRuleName, "rule-name", "", "", "Rule Name")
	_iotCmd.Flags().StringVarP(&_iotSbom, "sbom", "", "", "Sbom")
	_iotCmd.Flags().StringVarP(&_iotScheduledAuditName, "scheduled-audit-name", "", "", "Scheduled Audit Name")
	_iotCmd.Flags().StringVarP(&_iotSchedulingConfig, "scheduling-config", "", "", "Scheduling Config")
	_iotCmd.Flags().StringVarP(&_iotSecurityProfileDescription, "security-profile-description", "", "", "Security Profile Description")
	_iotCmd.Flags().StringVarP(&_iotSecurityProfileName, "security-profile-name", "", "", "Security Profile Name")
	_iotCmd.Flags().StringVarP(&_iotSecurityProfileTargetArn, "security-profile-target-arn", "", "", "Security Profile Target ARN")
	_iotCmd.Flags().StringSliceVarP(&_iotServerCertificateArns, "server-certificate-arns", "", nil, "Server Certificate Arns")
	_iotCmd.Flags().StringVarP(&_iotServerCertificateConfig, "server-certificate-config", "", "", "Server Certificate Config")
	_iotCmd.Flags().StringVarP(&_iotServiceType, "service-type", "", "", "Service Type")
	_iotCmd.Flags().StringVarP(&_iotSetAsActive, "set-as-active", "", "", "Set As Active")
	_iotCmd.Flags().StringVarP(&_iotSetAsDefault, "set-as-default", "", "", "Set As Default")
	_iotCmd.Flags().StringVarP(&_iotSigningDisabled, "signing-disabled", "", "", "Signing Disabled")
	_iotCmd.Flags().StringVarP(&_iotSortOrder, "sort-order", "", "", "Sort Order")
	_iotCmd.Flags().StringVarP(&_iotStartTime, "start-time", "", "", "Start Time")
	_iotCmd.Flags().StringVarP(&_iotStartedTimeFilter, "started-time-filter", "", "", "Started Time Filter")
	_iotCmd.Flags().StringVarP(&_iotStatus, "status", "", "", "Status")
	_iotCmd.Flags().StringVarP(&_iotStatusDetails, "status-details", "", "", "Status Details")
	_iotCmd.Flags().StringVarP(&_iotStreamId, "stream-id", "", "", "Stream ID")
	_iotCmd.Flags().StringSliceVarP(&_iotStringValues, "string-values", "", nil, "String Values")
	_iotCmd.Flags().StringVarP(&_iotSuppressIndefinitely, "suppress-indefinitely", "", "", "Suppress Indefinitely")
	_iotCmd.Flags().StringSliceVarP(&_iotTagKeys, "tag-keys", "", nil, "Tag Keys")
	_iotCmd.Flags().StringVarP(&_iotTags, "tags", "", "", "Tags")
	_iotCmd.Flags().StringVarP(&_iotTarget, "target", "", "", "Target")
	_iotCmd.Flags().StringVarP(&_iotTargetArn, "target-arn", "", "", "Target ARN")
	_iotCmd.Flags().StringVarP(&_iotTargetAwsAccount, "target-aws-account", "", "", "Target AWS Account")
	_iotCmd.Flags().StringSliceVarP(&_iotTargetCheckNames, "target-check-names", "", nil, "Target Check Names")
	_iotCmd.Flags().StringVarP(&_iotTargetName, "target-name", "", "", "Target Name")
	_iotCmd.Flags().StringVarP(&_iotTargetSelection, "target-selection", "", "", "Target Selection")
	_iotCmd.Flags().StringVarP(&_iotTargetType, "target-type", "", "", "Target Type")
	_iotCmd.Flags().StringSliceVarP(&_iotTargets, "targets", "", nil, "Targets")
	_iotCmd.Flags().StringVarP(&_iotTaskId, "task-id", "", "", "Task ID")
	_iotCmd.Flags().StringVarP(&_iotTaskStatus, "task-status", "", "", "Task Status")
	_iotCmd.Flags().StringVarP(&_iotTaskType, "task-type", "", "", "Task Type")
	_iotCmd.Flags().StringVarP(&_iotTemplateBody, "template-body", "", "", "Template Body")
	_iotCmd.Flags().StringVarP(&_iotTemplateName, "template-name", "", "", "Template Name")
	_iotCmd.Flags().StringVarP(&_iotTemplateVersion, "template-version", "", "", "Template Version")
	_iotCmd.Flags().StringVarP(&_iotThingArn, "thing-arn", "", "", "Thing ARN")
	_iotCmd.Flags().StringVarP(&_iotThingGroupArn, "thing-group-arn", "", "", "Thing Group ARN")
	_iotCmd.Flags().StringVarP(&_iotThingGroupId, "thing-group-id", "", "", "Thing Group ID")
	_iotCmd.Flags().StringVarP(&_iotThingGroupIndexingConfiguration, "thing-group-indexing-configuration", "", "", "Thing Group Indexing Configuration")
	_iotCmd.Flags().StringVarP(&_iotThingGroupName, "thing-group-name", "", "", "Thing Group Name")
	_iotCmd.Flags().StringVarP(&_iotThingGroupProperties, "thing-group-properties", "", "", "Thing Group Properties")
	_iotCmd.Flags().StringSliceVarP(&_iotThingGroupsToAdd, "thing-groups-to-add", "", nil, "Thing Groups To Add")
	_iotCmd.Flags().StringSliceVarP(&_iotThingGroupsToRemove, "thing-groups-to-remove", "", nil, "Thing Groups To Remove")
	_iotCmd.Flags().StringVarP(&_iotThingIndexingConfiguration, "thing-indexing-configuration", "", "", "Thing Indexing Configuration")
	_iotCmd.Flags().StringVarP(&_iotThingName, "thing-name", "", "", "Thing Name")
	_iotCmd.Flags().StringVarP(&_iotThingPrincipalType, "thing-principal-type", "", "", "Thing Principal Type")
	_iotCmd.Flags().StringVarP(&_iotThingTypeName, "thing-type-name", "", "", "Thing Type Name")
	_iotCmd.Flags().StringVarP(&_iotThingTypeProperties, "thing-type-properties", "", "", "Thing Type Properties")
	_iotCmd.Flags().StringVarP(&_iotTimeoutConfig, "timeout-config", "", "", "Timeout Config")
	_iotCmd.Flags().StringVarP(&_iotTlsConfig, "tls-config", "", "", "TLS Config")
	_iotCmd.Flags().StringVarP(&_iotTlsContext, "tls-context", "", "", "TLS Context")
	_iotCmd.Flags().StringVarP(&_iotToken, "token", "", "", "Token")
	_iotCmd.Flags().StringVarP(&_iotTokenKeyName, "token-key-name", "", "", "Token Key Name")
	_iotCmd.Flags().StringVarP(&_iotTokenSignature, "token-signature", "", "", "Token Signature")
	_iotCmd.Flags().StringVarP(&_iotTokenSigningPublicKeys, "token-signing-public-keys", "", "", "Token Signing Public Keys")
	_iotCmd.Flags().StringVarP(&_iotTopic, "topic", "", "", "Topic")
	_iotCmd.Flags().StringVarP(&_iotTopicRulePayload, "topic-rule-payload", "", "", "Topic Rule Payload")
	_iotCmd.Flags().StringVarP(&_iotTransferMessage, "transfer-message", "", "", "Transfer Message")
	_iotCmd.Flags().StringVarP(&_iotType, "type", "", "", "Type")
	_iotCmd.Flags().StringVarP(&_iotUndoDeprecate, "undo-deprecate", "", "", "Undo Deprecate")
	_iotCmd.Flags().StringVarP(&_iotUnit, "unit", "", "", "Unit")
	_iotCmd.Flags().StringVarP(&_iotUnsetDefaultVersion, "unset-default-version", "", "", "Unset Default Version")
	_iotCmd.Flags().StringVarP(&_iotUsePrefixAttributeValue, "use-prefix-attribute-value", "", "", "Use Prefix Attribute Value")
	_iotCmd.Flags().StringVarP(&_iotValidationCertificateArn, "validation-certificate-arn", "", "", "Validation Certificate ARN")
	_iotCmd.Flags().StringVarP(&_iotValidationResult, "validation-result", "", "", "Validation Result")
	_iotCmd.Flags().StringVarP(&_iotVerbose, "verbose", "", "", "Verbose")
	_iotCmd.Flags().StringVarP(&_iotVerificationCertificate, "verification-certificate", "", "", "Verification Certificate")
	_iotCmd.Flags().StringVarP(&_iotVerificationState, "verification-state", "", "", "Verification State")
	_iotCmd.Flags().StringVarP(&_iotVerificationStateDescription, "verification-state-description", "", "", "Verification State Description")
	_iotCmd.Flags().StringVarP(&_iotVersionId, "version-id", "", "", "Version ID")
	_iotCmd.Flags().StringVarP(&_iotVersionName, "version-name", "", "", "Version Name")
	_iotCmd.Flags().StringVarP(&_iotVersionUpdateByJobsConfig, "version-update-by-jobs-config", "", "", "Version Update By Jobs Config")
	_iotCmd.Flags().StringVarP(&_iotViolationEventOccurrenceRange, "violation-event-occurrence-range", "", "", "Violation Event Occurrence Range")
	_iotCmd.Flags().StringVarP(&_iotViolationId, "violation-id", "", "", "Violation ID")

	_iotCmd.Flags().BoolVarP(&_iotAcceptCertificateTransfer, "accept-certificate-transfer", "", false, "Accept Certificate Transfer")
	_iotCmd.Flags().BoolVarP(&_iotAddThingToBillingGroup, "add-thing-to-billing-group", "", false, "Add Thing To Billing Group")
	_iotCmd.Flags().BoolVarP(&_iotAddThingToThingGroup, "add-thing-to-thing-group", "", false, "Add Thing To Thing Group")
	_iotCmd.Flags().BoolVarP(&_iotAssociateSbomWithPackageVersion, "associate-sbom-with-package-version", "", false, "Associate Sbom With Package Version")
	_iotCmd.Flags().BoolVarP(&_iotAssociateTargetsWithJob, "associate-targets-with-job", "", false, "Associate Targets With Job")
	_iotCmd.Flags().BoolVarP(&_iotAttachPolicy, "attach-policy", "", false, "Attach Policy")
	_iotCmd.Flags().BoolVarP(&_iotAttachPrincipalPolicy, "attach-principal-policy", "", false, "Attach Principal Policy")
	_iotCmd.Flags().BoolVarP(&_iotAttachSecurityProfile, "attach-security-profile", "", false, "Attach Security Profile")
	_iotCmd.Flags().BoolVarP(&_iotAttachThingPrincipal, "attach-thing-principal", "", false, "Attach Thing Principal")
	_iotCmd.Flags().BoolVarP(&_iotCancelAuditMitigationActionsTask, "cancel-audit-mitigation-actions-task", "", false, "Cancel Audit Mitigation Actions Task")
	_iotCmd.Flags().BoolVarP(&_iotCancelAuditTask, "cancel-audit-task", "", false, "Cancel Audit Task")
	_iotCmd.Flags().BoolVarP(&_iotCancelCertificateTransfer, "cancel-certificate-transfer", "", false, "Cancel Certificate Transfer")
	_iotCmd.Flags().BoolVarP(&_iotCancelDetectMitigationActionsTask, "cancel-detect-mitigation-actions-task", "", false, "Cancel Detect Mitigation Actions Task")
	_iotCmd.Flags().BoolVarP(&_iotCancelJob, "cancel-job", "", false, "Cancel Job")
	_iotCmd.Flags().BoolVarP(&_iotCancelJobExecution, "cancel-job-execution", "", false, "Cancel Job Execution")
	_iotCmd.Flags().BoolVarP(&_iotClearDefaultAuthorizer, "clear-default-authorizer", "", false, "Clear Default Authorizer")
	_iotCmd.Flags().BoolVarP(&_iotConfirmTopicRuleDestination, "confirm-topic-rule-destination", "", false, "Confirm Topic Rule Destination")
	_iotCmd.Flags().BoolVarP(&_iotCreateAuditSuppression, "create-audit-suppression", "", false, "Create Audit Suppression")
	_iotCmd.Flags().BoolVarP(&_iotCreateAuthorizer, "create-authorizer", "", false, "Create Authorizer")
	_iotCmd.Flags().BoolVarP(&_iotCreateBillingGroup, "create-billing-group", "", false, "Create Billing Group")
	_iotCmd.Flags().BoolVarP(&_iotCreateCertificateFromCsr, "create-certificate-from-csr", "", false, "Create Certificate From Csr")
	_iotCmd.Flags().BoolVarP(&_iotCreateCertificateProvider, "create-certificate-provider", "", false, "Create Certificate Provider")
	_iotCmd.Flags().BoolVarP(&_iotCreateCommand, "create-command", "", false, "Create Command")
	_iotCmd.Flags().BoolVarP(&_iotCreateCustomMetric, "create-custom-metric", "", false, "Create Custom Metric")
	_iotCmd.Flags().BoolVarP(&_iotCreateDimension, "create-dimension", "", false, "Create Dimension")
	_iotCmd.Flags().BoolVarP(&_iotCreateDomainConfiguration, "create-domain-configuration", "", false, "Create Domain Configuration")
	_iotCmd.Flags().BoolVarP(&_iotCreateDynamicThingGroup, "create-dynamic-thing-group", "", false, "Create Dynamic Thing Group")
	_iotCmd.Flags().BoolVarP(&_iotCreateFleetMetric, "create-fleet-metric", "", false, "Create Fleet Metric")
	_iotCmd.Flags().BoolVarP(&_iotCreateJob, "create-job", "", false, "Create Job")
	_iotCmd.Flags().BoolVarP(&_iotCreateJobTemplate, "create-job-template", "", false, "Create Job Template")
	_iotCmd.Flags().BoolVarP(&_iotCreateKeysAndCertificate, "create-keys-and-certificate", "", false, "Create Keys And Certificate")
	_iotCmd.Flags().BoolVarP(&_iotCreateMitigationAction, "create-mitigation-action", "", false, "Create Mitigation Action")
	_iotCmd.Flags().BoolVarP(&_iotCreateOTAUpdate, "create-ota-update", "", false, "Create Ota Update")
	_iotCmd.Flags().BoolVarP(&_iotCreatePackage, "create-package", "", false, "Create Package")
	_iotCmd.Flags().BoolVarP(&_iotCreatePackageVersion, "create-package-version", "", false, "Create Package Version")
	_iotCmd.Flags().BoolVarP(&_iotCreatePolicy, "create-policy", "", false, "Create Policy")
	_iotCmd.Flags().BoolVarP(&_iotCreatePolicyVersion, "create-policy-version", "", false, "Create Policy Version")
	_iotCmd.Flags().BoolVarP(&_iotCreateProvisioningClaim, "create-provisioning-claim", "", false, "Create Provisioning Claim")
	_iotCmd.Flags().BoolVarP(&_iotCreateProvisioningTemplate, "create-provisioning-template", "", false, "Create Provisioning Template")
	_iotCmd.Flags().BoolVarP(&_iotCreateProvisioningTemplateVersion, "create-provisioning-template-version", "", false, "Create Provisioning Template Version")
	_iotCmd.Flags().BoolVarP(&_iotCreateRoleAlias, "create-role-alias", "", false, "Create Role Alias")
	_iotCmd.Flags().BoolVarP(&_iotCreateScheduledAudit, "create-scheduled-audit", "", false, "Create Scheduled Audit")
	_iotCmd.Flags().BoolVarP(&_iotCreateSecurityProfile, "create-security-profile", "", false, "Create Security Profile")
	_iotCmd.Flags().BoolVarP(&_iotCreateStream, "create-stream", "", false, "Create Stream")
	_iotCmd.Flags().BoolVarP(&_iotCreateThing, "create-thing", "", false, "Create Thing")
	_iotCmd.Flags().BoolVarP(&_iotCreateThingGroup, "create-thing-group", "", false, "Create Thing Group")
	_iotCmd.Flags().BoolVarP(&_iotCreateThingType, "create-thing-type", "", false, "Create Thing Type")
	_iotCmd.Flags().BoolVarP(&_iotCreateTopicRule, "create-topic-rule", "", false, "Create Topic Rule")
	_iotCmd.Flags().BoolVarP(&_iotCreateTopicRuleDestination, "create-topic-rule-destination", "", false, "Create Topic Rule Destination")
	_iotCmd.Flags().BoolVarP(&_iotDeleteAccountAuditConfiguration, "delete-account-audit-configuration", "", false, "Delete Account Audit Configuration")
	_iotCmd.Flags().BoolVarP(&_iotDeleteAuditSuppression, "delete-audit-suppression", "", false, "Delete Audit Suppression")
	_iotCmd.Flags().BoolVarP(&_iotDeleteAuthorizer, "delete-authorizer", "", false, "Delete Authorizer")
	_iotCmd.Flags().BoolVarP(&_iotDeleteBillingGroup, "delete-billing-group", "", false, "Delete Billing Group")
	_iotCmd.Flags().BoolVarP(&_iotDeleteCACertificate, "delete-ca-certificate", "", false, "Delete Ca Certificate")
	_iotCmd.Flags().BoolVarP(&_iotDeleteCertificate, "delete-certificate", "", false, "Delete Certificate")
	_iotCmd.Flags().BoolVarP(&_iotDeleteCertificateProvider, "delete-certificate-provider", "", false, "Delete Certificate Provider")
	_iotCmd.Flags().BoolVarP(&_iotDeleteCommand, "delete-command", "", false, "Delete Command")
	_iotCmd.Flags().BoolVarP(&_iotDeleteCommandExecution, "delete-command-execution", "", false, "Delete Command Execution")
	_iotCmd.Flags().BoolVarP(&_iotDeleteCustomMetric, "delete-custom-metric", "", false, "Delete Custom Metric")
	_iotCmd.Flags().BoolVarP(&_iotDeleteDimension, "delete-dimension", "", false, "Delete Dimension")
	_iotCmd.Flags().BoolVarP(&_iotDeleteDomainConfiguration, "delete-domain-configuration", "", false, "Delete Domain Configuration")
	_iotCmd.Flags().BoolVarP(&_iotDeleteDynamicThingGroup, "delete-dynamic-thing-group", "", false, "Delete Dynamic Thing Group")
	_iotCmd.Flags().BoolVarP(&_iotDeleteFleetMetric, "delete-fleet-metric", "", false, "Delete Fleet Metric")
	_iotCmd.Flags().BoolVarP(&_iotDeleteJob, "delete-job", "", false, "Delete Job")
	_iotCmd.Flags().BoolVarP(&_iotDeleteJobExecution, "delete-job-execution", "", false, "Delete Job Execution")
	_iotCmd.Flags().BoolVarP(&_iotDeleteJobTemplate, "delete-job-template", "", false, "Delete Job Template")
	_iotCmd.Flags().BoolVarP(&_iotDeleteMitigationAction, "delete-mitigation-action", "", false, "Delete Mitigation Action")
	_iotCmd.Flags().BoolVarP(&_iotDeleteOTAUpdate, "delete-ota-update", "", false, "Delete Ota Update")
	_iotCmd.Flags().BoolVarP(&_iotDeletePackage, "delete-package", "", false, "Delete Package")
	_iotCmd.Flags().BoolVarP(&_iotDeletePackageVersion, "delete-package-version", "", false, "Delete Package Version")
	_iotCmd.Flags().BoolVarP(&_iotDeletePolicy, "delete-policy", "", false, "Delete Policy")
	_iotCmd.Flags().BoolVarP(&_iotDeletePolicyVersion, "delete-policy-version", "", false, "Delete Policy Version")
	_iotCmd.Flags().BoolVarP(&_iotDeleteProvisioningTemplate, "delete-provisioning-template", "", false, "Delete Provisioning Template")
	_iotCmd.Flags().BoolVarP(&_iotDeleteProvisioningTemplateVersion, "delete-provisioning-template-version", "", false, "Delete Provisioning Template Version")
	_iotCmd.Flags().BoolVarP(&_iotDeleteRegistrationCode, "delete-registration-code", "", false, "Delete Registration Code")
	_iotCmd.Flags().BoolVarP(&_iotDeleteRoleAlias, "delete-role-alias", "", false, "Delete Role Alias")
	_iotCmd.Flags().BoolVarP(&_iotDeleteScheduledAudit, "delete-scheduled-audit", "", false, "Delete Scheduled Audit")
	_iotCmd.Flags().BoolVarP(&_iotDeleteSecurityProfile, "delete-security-profile", "", false, "Delete Security Profile")
	_iotCmd.Flags().BoolVarP(&_iotDeleteStream, "delete-stream", "", false, "Delete Stream")
	_iotCmd.Flags().BoolVarP(&_iotDeleteThing, "delete-thing", "", false, "Delete Thing")
	_iotCmd.Flags().BoolVarP(&_iotDeleteThingGroup, "delete-thing-group", "", false, "Delete Thing Group")
	_iotCmd.Flags().BoolVarP(&_iotDeleteThingType, "delete-thing-type", "", false, "Delete Thing Type")
	_iotCmd.Flags().BoolVarP(&_iotDeleteTopicRule, "delete-topic-rule", "", false, "Delete Topic Rule")
	_iotCmd.Flags().BoolVarP(&_iotDeleteTopicRuleDestination, "delete-topic-rule-destination", "", false, "Delete Topic Rule Destination")
	_iotCmd.Flags().BoolVarP(&_iotDeleteV2LoggingLevel, "delete-v2-logging-level", "", false, "Delete V2 Logging Level")
	_iotCmd.Flags().BoolVarP(&_iotDeprecateThingType, "deprecate-thing-type", "", false, "Deprecate Thing Type")
	_iotCmd.Flags().BoolVarP(&_iotDescribeAccountAuditConfiguration, "describe-account-audit-configuration", "", false, "Describe Account Audit Configuration")
	_iotCmd.Flags().BoolVarP(&_iotDescribeAuditFinding, "describe-audit-finding", "", false, "Describe Audit Finding")
	_iotCmd.Flags().BoolVarP(&_iotDescribeAuditMitigationActionsTask, "describe-audit-mitigation-actions-task", "", false, "Describe Audit Mitigation Actions Task")
	_iotCmd.Flags().BoolVarP(&_iotDescribeAuditSuppression, "describe-audit-suppression", "", false, "Describe Audit Suppression")
	_iotCmd.Flags().BoolVarP(&_iotDescribeAuditTask, "describe-audit-task", "", false, "Describe Audit Task")
	_iotCmd.Flags().BoolVarP(&_iotDescribeAuthorizer, "describe-authorizer", "", false, "Describe Authorizer")
	_iotCmd.Flags().BoolVarP(&_iotDescribeBillingGroup, "describe-billing-group", "", false, "Describe Billing Group")
	_iotCmd.Flags().BoolVarP(&_iotDescribeCACertificate, "describe-ca-certificate", "", false, "Describe Ca Certificate")
	_iotCmd.Flags().BoolVarP(&_iotDescribeCertificate, "describe-certificate", "", false, "Describe Certificate")
	_iotCmd.Flags().BoolVarP(&_iotDescribeCertificateProvider, "describe-certificate-provider", "", false, "Describe Certificate Provider")
	_iotCmd.Flags().BoolVarP(&_iotDescribeCustomMetric, "describe-custom-metric", "", false, "Describe Custom Metric")
	_iotCmd.Flags().BoolVarP(&_iotDescribeDefaultAuthorizer, "describe-default-authorizer", "", false, "Describe Default Authorizer")
	_iotCmd.Flags().BoolVarP(&_iotDescribeDetectMitigationActionsTask, "describe-detect-mitigation-actions-task", "", false, "Describe Detect Mitigation Actions Task")
	_iotCmd.Flags().BoolVarP(&_iotDescribeDimension, "describe-dimension", "", false, "Describe Dimension")
	_iotCmd.Flags().BoolVarP(&_iotDescribeDomainConfiguration, "describe-domain-configuration", "", false, "Describe Domain Configuration")
	_iotCmd.Flags().BoolVarP(&_iotDescribeEncryptionConfiguration, "describe-encryption-configuration", "", false, "Describe Encryption Configuration")
	_iotCmd.Flags().BoolVarP(&_iotDescribeEndpoint, "describe-endpoint", "", false, "Describe Endpoint")
	_iotCmd.Flags().BoolVarP(&_iotDescribeEventConfigurations, "describe-event-configurations", "", false, "Describe Event Configurations")
	_iotCmd.Flags().BoolVarP(&_iotDescribeFleetMetric, "describe-fleet-metric", "", false, "Describe Fleet Metric")
	_iotCmd.Flags().BoolVarP(&_iotDescribeIndex, "describe-index", "", false, "Describe Index")
	_iotCmd.Flags().BoolVarP(&_iotDescribeJob, "describe-job", "", false, "Describe Job")
	_iotCmd.Flags().BoolVarP(&_iotDescribeJobExecution, "describe-job-execution", "", false, "Describe Job Execution")
	_iotCmd.Flags().BoolVarP(&_iotDescribeJobTemplate, "describe-job-template", "", false, "Describe Job Template")
	_iotCmd.Flags().BoolVarP(&_iotDescribeManagedJobTemplate, "describe-managed-job-template", "", false, "Describe Managed Job Template")
	_iotCmd.Flags().BoolVarP(&_iotDescribeMitigationAction, "describe-mitigation-action", "", false, "Describe Mitigation Action")
	_iotCmd.Flags().BoolVarP(&_iotDescribeProvisioningTemplate, "describe-provisioning-template", "", false, "Describe Provisioning Template")
	_iotCmd.Flags().BoolVarP(&_iotDescribeProvisioningTemplateVersion, "describe-provisioning-template-version", "", false, "Describe Provisioning Template Version")
	_iotCmd.Flags().BoolVarP(&_iotDescribeRoleAlias, "describe-role-alias", "", false, "Describe Role Alias")
	_iotCmd.Flags().BoolVarP(&_iotDescribeScheduledAudit, "describe-scheduled-audit", "", false, "Describe Scheduled Audit")
	_iotCmd.Flags().BoolVarP(&_iotDescribeSecurityProfile, "describe-security-profile", "", false, "Describe Security Profile")
	_iotCmd.Flags().BoolVarP(&_iotDescribeStream, "describe-stream", "", false, "Describe Stream")
	_iotCmd.Flags().BoolVarP(&_iotDescribeThing, "describe-thing", "", false, "Describe Thing")
	_iotCmd.Flags().BoolVarP(&_iotDescribeThingGroup, "describe-thing-group", "", false, "Describe Thing Group")
	_iotCmd.Flags().BoolVarP(&_iotDescribeThingRegistrationTask, "describe-thing-registration-task", "", false, "Describe Thing Registration Task")
	_iotCmd.Flags().BoolVarP(&_iotDescribeThingType, "describe-thing-type", "", false, "Describe Thing Type")
	_iotCmd.Flags().BoolVarP(&_iotDetachPolicy, "detach-policy", "", false, "Detach Policy")
	_iotCmd.Flags().BoolVarP(&_iotDetachPrincipalPolicy, "detach-principal-policy", "", false, "Detach Principal Policy")
	_iotCmd.Flags().BoolVarP(&_iotDetachSecurityProfile, "detach-security-profile", "", false, "Detach Security Profile")
	_iotCmd.Flags().BoolVarP(&_iotDetachThingPrincipal, "detach-thing-principal", "", false, "Detach Thing Principal")
	_iotCmd.Flags().BoolVarP(&_iotDisableTopicRule, "disable-topic-rule", "", false, "Disable Topic Rule")
	_iotCmd.Flags().BoolVarP(&_iotDisassociateSbomFromPackageVersion, "disassociate-sbom-from-package-version", "", false, "Disassociate Sbom From Package Version")
	_iotCmd.Flags().BoolVarP(&_iotEnableTopicRule, "enable-topic-rule", "", false, "Enable Topic Rule")
	_iotCmd.Flags().BoolVarP(&_iotGetBehaviorModelTrainingSummaries, "get-behavior-model-training-summaries", "", false, "Get Behavior Model Training Summaries")
	_iotCmd.Flags().BoolVarP(&_iotGetBucketsAggregation, "get-buckets-aggregation", "", false, "Get Buckets Aggregation")
	_iotCmd.Flags().BoolVarP(&_iotGetCardinality, "get-cardinality", "", false, "Get Cardinality")
	_iotCmd.Flags().BoolVarP(&_iotGetCommand, "get-command", "", false, "Get Command")
	_iotCmd.Flags().BoolVarP(&_iotGetCommandExecution, "get-command-execution", "", false, "Get Command Execution")
	_iotCmd.Flags().BoolVarP(&_iotGetEffectivePolicies, "get-effective-policies", "", false, "Get Effective Policies")
	_iotCmd.Flags().BoolVarP(&_iotGetIndexingConfiguration, "get-indexing-configuration", "", false, "Get Indexing Configuration")
	_iotCmd.Flags().BoolVarP(&_iotGetJobDocument, "get-job-document", "", false, "Get Job Document")
	_iotCmd.Flags().BoolVarP(&_iotGetLoggingOptions, "get-logging-options", "", false, "Get Logging Options")
	_iotCmd.Flags().BoolVarP(&_iotGetOTAUpdate, "get-ota-update", "", false, "Get Ota Update")
	_iotCmd.Flags().BoolVarP(&_iotGetPackage, "get-package", "", false, "Get Package")
	_iotCmd.Flags().BoolVarP(&_iotGetPackageConfiguration, "get-package-configuration", "", false, "Get Package Configuration")
	_iotCmd.Flags().BoolVarP(&_iotGetPackageVersion, "get-package-version", "", false, "Get Package Version")
	_iotCmd.Flags().BoolVarP(&_iotGetPercentiles, "get-percentiles", "", false, "Get Percentiles")
	_iotCmd.Flags().BoolVarP(&_iotGetPolicy, "get-policy", "", false, "Get Policy")
	_iotCmd.Flags().BoolVarP(&_iotGetPolicyVersion, "get-policy-version", "", false, "Get Policy Version")
	_iotCmd.Flags().BoolVarP(&_iotGetRegistrationCode, "get-registration-code", "", false, "Get Registration Code")
	_iotCmd.Flags().BoolVarP(&_iotGetStatistics, "get-statistics", "", false, "Get Statistics")
	_iotCmd.Flags().BoolVarP(&_iotGetThingConnectivityData, "get-thing-connectivity-data", "", false, "Get Thing Connectivity Data")
	_iotCmd.Flags().BoolVarP(&_iotGetTopicRule, "get-topic-rule", "", false, "Get Topic Rule")
	_iotCmd.Flags().BoolVarP(&_iotGetTopicRuleDestination, "get-topic-rule-destination", "", false, "Get Topic Rule Destination")
	_iotCmd.Flags().BoolVarP(&_iotGetV2LoggingOptions, "get-v2-logging-options", "", false, "Get V2 Logging Options")
	_iotCmd.Flags().BoolVarP(&_iotListActiveViolations, "list-active-violations", "", false, "List Active Violations")
	_iotCmd.Flags().BoolVarP(&_iotListAttachedPolicies, "list-attached-policies", "", false, "List Attached Policies")
	_iotCmd.Flags().BoolVarP(&_iotListAuditFindings, "list-audit-findings", "", false, "List Audit Findings")
	_iotCmd.Flags().BoolVarP(&_iotListAuditMitigationActionsExecutions, "list-audit-mitigation-actions-executions", "", false, "List Audit Mitigation Actions Executions")
	_iotCmd.Flags().BoolVarP(&_iotListAuditMitigationActionsTasks, "list-audit-mitigation-actions-tasks", "", false, "List Audit Mitigation Actions Tasks")
	_iotCmd.Flags().BoolVarP(&_iotListAuditSuppressions, "list-audit-suppressions", "", false, "List Audit Suppressions")
	_iotCmd.Flags().BoolVarP(&_iotListAuditTasks, "list-audit-tasks", "", false, "List Audit Tasks")
	_iotCmd.Flags().BoolVarP(&_iotListAuthorizers, "list-authorizers", "", false, "List Authorizers")
	_iotCmd.Flags().BoolVarP(&_iotListBillingGroups, "list-billing-groups", "", false, "List Billing Groups")
	_iotCmd.Flags().BoolVarP(&_iotListCACertificates, "list-ca-certificates", "", false, "List Ca Certificates")
	_iotCmd.Flags().BoolVarP(&_iotListCertificateProviders, "list-certificate-providers", "", false, "List Certificate Providers")
	_iotCmd.Flags().BoolVarP(&_iotListCertificates, "list-certificates", "", false, "List Certificates")
	_iotCmd.Flags().BoolVarP(&_iotListCertificatesByCA, "list-certificates-by-ca", "", false, "List Certificates By Ca")
	_iotCmd.Flags().BoolVarP(&_iotListCommandExecutions, "list-command-executions", "", false, "List Command Executions")
	_iotCmd.Flags().BoolVarP(&_iotListCommands, "list-commands", "", false, "List Commands")
	_iotCmd.Flags().BoolVarP(&_iotListCustomMetrics, "list-custom-metrics", "", false, "List Custom Metrics")
	_iotCmd.Flags().BoolVarP(&_iotListDetectMitigationActionsExecutions, "list-detect-mitigation-actions-executions", "", false, "List Detect Mitigation Actions Executions")
	_iotCmd.Flags().BoolVarP(&_iotListDetectMitigationActionsTasks, "list-detect-mitigation-actions-tasks", "", false, "List Detect Mitigation Actions Tasks")
	_iotCmd.Flags().BoolVarP(&_iotListDimensions, "list-dimensions", "", false, "List Dimensions")
	_iotCmd.Flags().BoolVarP(&_iotListDomainConfigurations, "list-domain-configurations", "", false, "List Domain Configurations")
	_iotCmd.Flags().BoolVarP(&_iotListFleetMetrics, "list-fleet-metrics", "", false, "List Fleet Metrics")
	_iotCmd.Flags().BoolVarP(&_iotListIndices, "list-indices", "", false, "List Indices")
	_iotCmd.Flags().BoolVarP(&_iotListJobExecutionsForJob, "list-job-executions-for-job", "", false, "List Job Executions For Job")
	_iotCmd.Flags().BoolVarP(&_iotListJobExecutionsForThing, "list-job-executions-for-thing", "", false, "List Job Executions For Thing")
	_iotCmd.Flags().BoolVarP(&_iotListJobTemplates, "list-job-templates", "", false, "List Job Templates")
	_iotCmd.Flags().BoolVarP(&_iotListJobs, "list-jobs", "", false, "List Jobs")
	_iotCmd.Flags().BoolVarP(&_iotListManagedJobTemplates, "list-managed-job-templates", "", false, "List Managed Job Templates")
	_iotCmd.Flags().BoolVarP(&_iotListMetricValues, "list-metric-values", "", false, "List Metric Values")
	_iotCmd.Flags().BoolVarP(&_iotListMitigationActions, "list-mitigation-actions", "", false, "List Mitigation Actions")
	_iotCmd.Flags().BoolVarP(&_iotListOTAUpdates, "list-ota-updates", "", false, "List Ota Updates")
	_iotCmd.Flags().BoolVarP(&_iotListOutgoingCertificates, "list-outgoing-certificates", "", false, "List Outgoing Certificates")
	_iotCmd.Flags().BoolVarP(&_iotListPackageVersions, "list-package-versions", "", false, "List Package Versions")
	_iotCmd.Flags().BoolVarP(&_iotListPackages, "list-packages", "", false, "List Packages")
	_iotCmd.Flags().BoolVarP(&_iotListPolicies, "list-policies", "", false, "List Policies")
	_iotCmd.Flags().BoolVarP(&_iotListPolicyPrincipals, "list-policy-principals", "", false, "List Policy Principals")
	_iotCmd.Flags().BoolVarP(&_iotListPolicyVersions, "list-policy-versions", "", false, "List Policy Versions")
	_iotCmd.Flags().BoolVarP(&_iotListPrincipalPolicies, "list-principal-policies", "", false, "List Principal Policies")
	_iotCmd.Flags().BoolVarP(&_iotListPrincipalThings, "list-principal-things", "", false, "List Principal Things")
	_iotCmd.Flags().BoolVarP(&_iotListPrincipalThingsV2, "list-principal-things-v2", "", false, "List Principal Things V2")
	_iotCmd.Flags().BoolVarP(&_iotListProvisioningTemplateVersions, "list-provisioning-template-versions", "", false, "List Provisioning Template Versions")
	_iotCmd.Flags().BoolVarP(&_iotListProvisioningTemplates, "list-provisioning-templates", "", false, "List Provisioning Templates")
	_iotCmd.Flags().BoolVarP(&_iotListRelatedResourcesForAuditFinding, "list-related-resources-for-audit-finding", "", false, "List Related Resources For Audit Finding")
	_iotCmd.Flags().BoolVarP(&_iotListRoleAliases, "list-role-aliases", "", false, "List Role Aliases")
	_iotCmd.Flags().BoolVarP(&_iotListSbomValidationResults, "list-sbom-validation-results", "", false, "List Sbom Validation Results")
	_iotCmd.Flags().BoolVarP(&_iotListScheduledAudits, "list-scheduled-audits", "", false, "List Scheduled Audits")
	_iotCmd.Flags().BoolVarP(&_iotListSecurityProfiles, "list-security-profiles", "", false, "List Security Profiles")
	_iotCmd.Flags().BoolVarP(&_iotListSecurityProfilesForTarget, "list-security-profiles-for-target", "", false, "List Security Profiles For Target")
	_iotCmd.Flags().BoolVarP(&_iotListStreams, "list-streams", "", false, "List Streams")
	_iotCmd.Flags().BoolVarP(&_iotListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_iotCmd.Flags().BoolVarP(&_iotListTargetsForPolicy, "list-targets-for-policy", "", false, "List Targets For Policy")
	_iotCmd.Flags().BoolVarP(&_iotListTargetsForSecurityProfile, "list-targets-for-security-profile", "", false, "List Targets For Security Profile")
	_iotCmd.Flags().BoolVarP(&_iotListThingGroups, "list-thing-groups", "", false, "List Thing Groups")
	_iotCmd.Flags().BoolVarP(&_iotListThingGroupsForThing, "list-thing-groups-for-thing", "", false, "List Thing Groups For Thing")
	_iotCmd.Flags().BoolVarP(&_iotListThingPrincipals, "list-thing-principals", "", false, "List Thing Principals")
	_iotCmd.Flags().BoolVarP(&_iotListThingPrincipalsV2, "list-thing-principals-v2", "", false, "List Thing Principals V2")
	_iotCmd.Flags().BoolVarP(&_iotListThingRegistrationTaskReports, "list-thing-registration-task-reports", "", false, "List Thing Registration Task Reports")
	_iotCmd.Flags().BoolVarP(&_iotListThingRegistrationTasks, "list-thing-registration-tasks", "", false, "List Thing Registration Tasks")
	_iotCmd.Flags().BoolVarP(&_iotListThingTypes, "list-thing-types", "", false, "List Thing Types")
	_iotCmd.Flags().BoolVarP(&_iotListThings, "list-things", "", false, "List Things")
	_iotCmd.Flags().BoolVarP(&_iotListThingsInBillingGroup, "list-things-in-billing-group", "", false, "List Things In Billing Group")
	_iotCmd.Flags().BoolVarP(&_iotListThingsInThingGroup, "list-things-in-thing-group", "", false, "List Things In Thing Group")
	_iotCmd.Flags().BoolVarP(&_iotListTopicRuleDestinations, "list-topic-rule-destinations", "", false, "List Topic Rule Destinations")
	_iotCmd.Flags().BoolVarP(&_iotListTopicRules, "list-topic-rules", "", false, "List Topic Rules")
	_iotCmd.Flags().BoolVarP(&_iotListV2LoggingLevels, "list-v2-logging-levels", "", false, "List V2 Logging Levels")
	_iotCmd.Flags().BoolVarP(&_iotListViolationEvents, "list-violation-events", "", false, "List Violation Events")
	_iotCmd.Flags().BoolVarP(&_iotPutVerificationStateOnViolation, "put-verification-state-on-violation", "", false, "Put Verification State On Violation")
	_iotCmd.Flags().BoolVarP(&_iotRegisterCACertificate, "register-ca-certificate", "", false, "Register Ca Certificate")
	_iotCmd.Flags().BoolVarP(&_iotRegisterCertificate, "register-certificate", "", false, "Register Certificate")
	_iotCmd.Flags().BoolVarP(&_iotRegisterCertificateWithoutCA, "register-certificate-without-ca", "", false, "Register Certificate Without Ca")
	_iotCmd.Flags().BoolVarP(&_iotRegisterThing, "register-thing", "", false, "Register Thing")
	_iotCmd.Flags().BoolVarP(&_iotRejectCertificateTransfer, "reject-certificate-transfer", "", false, "Reject Certificate Transfer")
	_iotCmd.Flags().BoolVarP(&_iotRemoveThingFromBillingGroup, "remove-thing-from-billing-group", "", false, "Remove Thing From Billing Group")
	_iotCmd.Flags().BoolVarP(&_iotRemoveThingFromThingGroup, "remove-thing-from-thing-group", "", false, "Remove Thing From Thing Group")
	_iotCmd.Flags().BoolVarP(&_iotReplaceTopicRule, "replace-topic-rule", "", false, "Replace Topic Rule")
	_iotCmd.Flags().BoolVarP(&_iotSearchIndex, "search-index", "", false, "Search Index")
	_iotCmd.Flags().BoolVarP(&_iotSetDefaultAuthorizer, "set-default-authorizer", "", false, "Set Default Authorizer")
	_iotCmd.Flags().BoolVarP(&_iotSetDefaultPolicyVersion, "set-default-policy-version", "", false, "Set Default Policy Version")
	_iotCmd.Flags().BoolVarP(&_iotSetLoggingOptions, "set-logging-options", "", false, "Set Logging Options")
	_iotCmd.Flags().BoolVarP(&_iotSetV2LoggingLevel, "set-v2-logging-level", "", false, "Set V2 Logging Level")
	_iotCmd.Flags().BoolVarP(&_iotSetV2LoggingOptions, "set-v2-logging-options", "", false, "Set V2 Logging Options")
	_iotCmd.Flags().BoolVarP(&_iotStartAuditMitigationActionsTask, "start-audit-mitigation-actions-task", "", false, "Start Audit Mitigation Actions Task")
	_iotCmd.Flags().BoolVarP(&_iotStartDetectMitigationActionsTask, "start-detect-mitigation-actions-task", "", false, "Start Detect Mitigation Actions Task")
	_iotCmd.Flags().BoolVarP(&_iotStartOnDemandAuditTask, "start-on-demand-audit-task", "", false, "Start On Demand Audit Task")
	_iotCmd.Flags().BoolVarP(&_iotStartThingRegistrationTask, "start-thing-registration-task", "", false, "Start Thing Registration Task")
	_iotCmd.Flags().BoolVarP(&_iotStopThingRegistrationTask, "stop-thing-registration-task", "", false, "Stop Thing Registration Task")
	_iotCmd.Flags().BoolVarP(&_iotTagResource, "tag-resource", "", false, "Tag Resource")
	_iotCmd.Flags().BoolVarP(&_iotTestAuthorization, "test-authorization", "", false, "Test Authorization")
	_iotCmd.Flags().BoolVarP(&_iotTestInvokeAuthorizer, "test-invoke-authorizer", "", false, "Test Invoke Authorizer")
	_iotCmd.Flags().BoolVarP(&_iotTransferCertificate, "transfer-certificate", "", false, "Transfer Certificate")
	_iotCmd.Flags().BoolVarP(&_iotUntagResource, "untag-resource", "", false, "Untag Resource")
	_iotCmd.Flags().BoolVarP(&_iotUpdateAccountAuditConfiguration, "update-account-audit-configuration", "", false, "Update Account Audit Configuration")
	_iotCmd.Flags().BoolVarP(&_iotUpdateAuditSuppression, "update-audit-suppression", "", false, "Update Audit Suppression")
	_iotCmd.Flags().BoolVarP(&_iotUpdateAuthorizer, "update-authorizer", "", false, "Update Authorizer")
	_iotCmd.Flags().BoolVarP(&_iotUpdateBillingGroup, "update-billing-group", "", false, "Update Billing Group")
	_iotCmd.Flags().BoolVarP(&_iotUpdateCACertificate, "update-ca-certificate", "", false, "Update Ca Certificate")
	_iotCmd.Flags().BoolVarP(&_iotUpdateCertificate, "update-certificate", "", false, "Update Certificate")
	_iotCmd.Flags().BoolVarP(&_iotUpdateCertificateProvider, "update-certificate-provider", "", false, "Update Certificate Provider")
	_iotCmd.Flags().BoolVarP(&_iotUpdateCommand, "update-command", "", false, "Update Command")
	_iotCmd.Flags().BoolVarP(&_iotUpdateCustomMetric, "update-custom-metric", "", false, "Update Custom Metric")
	_iotCmd.Flags().BoolVarP(&_iotUpdateDimension, "update-dimension", "", false, "Update Dimension")
	_iotCmd.Flags().BoolVarP(&_iotUpdateDomainConfiguration, "update-domain-configuration", "", false, "Update Domain Configuration")
	_iotCmd.Flags().BoolVarP(&_iotUpdateDynamicThingGroup, "update-dynamic-thing-group", "", false, "Update Dynamic Thing Group")
	_iotCmd.Flags().BoolVarP(&_iotUpdateEncryptionConfiguration, "update-encryption-configuration", "", false, "Update Encryption Configuration")
	_iotCmd.Flags().BoolVarP(&_iotUpdateEventConfigurations, "update-event-configurations", "", false, "Update Event Configurations")
	_iotCmd.Flags().BoolVarP(&_iotUpdateFleetMetric, "update-fleet-metric", "", false, "Update Fleet Metric")
	_iotCmd.Flags().BoolVarP(&_iotUpdateIndexingConfiguration, "update-indexing-configuration", "", false, "Update Indexing Configuration")
	_iotCmd.Flags().BoolVarP(&_iotUpdateJob, "update-job", "", false, "Update Job")
	_iotCmd.Flags().BoolVarP(&_iotUpdateMitigationAction, "update-mitigation-action", "", false, "Update Mitigation Action")
	_iotCmd.Flags().BoolVarP(&_iotUpdatePackage, "update-package", "", false, "Update Package")
	_iotCmd.Flags().BoolVarP(&_iotUpdatePackageConfiguration, "update-package-configuration", "", false, "Update Package Configuration")
	_iotCmd.Flags().BoolVarP(&_iotUpdatePackageVersion, "update-package-version", "", false, "Update Package Version")
	_iotCmd.Flags().BoolVarP(&_iotUpdateProvisioningTemplate, "update-provisioning-template", "", false, "Update Provisioning Template")
	_iotCmd.Flags().BoolVarP(&_iotUpdateRoleAlias, "update-role-alias", "", false, "Update Role Alias")
	_iotCmd.Flags().BoolVarP(&_iotUpdateScheduledAudit, "update-scheduled-audit", "", false, "Update Scheduled Audit")
	_iotCmd.Flags().BoolVarP(&_iotUpdateSecurityProfile, "update-security-profile", "", false, "Update Security Profile")
	_iotCmd.Flags().BoolVarP(&_iotUpdateStream, "update-stream", "", false, "Update Stream")
	_iotCmd.Flags().BoolVarP(&_iotUpdateThing, "update-thing", "", false, "Update Thing")
	_iotCmd.Flags().BoolVarP(&_iotUpdateThingGroup, "update-thing-group", "", false, "Update Thing Group")
	_iotCmd.Flags().BoolVarP(&_iotUpdateThingGroupsForThing, "update-thing-groups-for-thing", "", false, "Update Thing Groups For Thing")
	_iotCmd.Flags().BoolVarP(&_iotUpdateThingType, "update-thing-type", "", false, "Update Thing Type")
	_iotCmd.Flags().BoolVarP(&_iotUpdateTopicRuleDestination, "update-topic-rule-destination", "", false, "Update Topic Rule Destination")
	_iotCmd.Flags().BoolVarP(&_iotValidateSecurityProfileBehaviors, "validate-security-profile-behaviors", "", false, "Validate Security Profile Behaviors")

}
