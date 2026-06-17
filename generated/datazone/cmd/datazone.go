package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datazone"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// datazoneCmd represents the datazone command
var _datazoneCmd = &cobra.Command{
	Use:   "datazone",
	Short: "AWS datazone CLI",
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
		client := datazone.NewFromConfig(cfg)
		if _datazoneAcceptPredictions {
			datazone_AcceptPredictions(cfg, client)
			return
		}
		if _datazoneAcceptSubscriptionRequest {
			datazone_AcceptSubscriptionRequest(cfg, client)
			return
		}
		if _datazoneAddEntityOwner {
			datazone_AddEntityOwner(cfg, client)
			return
		}
		if _datazoneAddPolicyGrant {
			datazone_AddPolicyGrant(cfg, client)
			return
		}
		if _datazoneAssociateEnvironmentRole {
			datazone_AssociateEnvironmentRole(cfg, client)
			return
		}
		if _datazoneAssociateGovernedTerms {
			datazone_AssociateGovernedTerms(cfg, client)
			return
		}
		if _datazoneBatchGetAttributesMetadata {
			datazone_BatchGetAttributesMetadata(cfg, client)
			return
		}
		if _datazoneBatchPutAttributesMetadata {
			datazone_BatchPutAttributesMetadata(cfg, client)
			return
		}
		if _datazoneCancelMetadataGenerationRun {
			datazone_CancelMetadataGenerationRun(cfg, client)
			return
		}
		if _datazoneCancelSubscription {
			datazone_CancelSubscription(cfg, client)
			return
		}
		if _datazoneCreateAccountPool {
			datazone_CreateAccountPool(cfg, client)
			return
		}
		if _datazoneCreateAsset {
			datazone_CreateAsset(cfg, client)
			return
		}
		if _datazoneCreateAssetFilter {
			datazone_CreateAssetFilter(cfg, client)
			return
		}
		if _datazoneCreateAssetRevision {
			datazone_CreateAssetRevision(cfg, client)
			return
		}
		if _datazoneCreateAssetType {
			datazone_CreateAssetType(cfg, client)
			return
		}
		if _datazoneCreateConnection {
			datazone_CreateConnection(cfg, client)
			return
		}
		if _datazoneCreateDataProduct {
			datazone_CreateDataProduct(cfg, client)
			return
		}
		if _datazoneCreateDataProductRevision {
			datazone_CreateDataProductRevision(cfg, client)
			return
		}
		if _datazoneCreateDataSource {
			datazone_CreateDataSource(cfg, client)
			return
		}
		if _datazoneCreateDomain {
			datazone_CreateDomain(cfg, client)
			return
		}
		if _datazoneCreateDomainUnit {
			datazone_CreateDomainUnit(cfg, client)
			return
		}
		if _datazoneCreateEnvironment {
			datazone_CreateEnvironment(cfg, client)
			return
		}
		if _datazoneCreateEnvironmentAction {
			datazone_CreateEnvironmentAction(cfg, client)
			return
		}
		if _datazoneCreateEnvironmentBlueprint {
			datazone_CreateEnvironmentBlueprint(cfg, client)
			return
		}
		if _datazoneCreateEnvironmentProfile {
			datazone_CreateEnvironmentProfile(cfg, client)
			return
		}
		if _datazoneCreateFormType {
			datazone_CreateFormType(cfg, client)
			return
		}
		if _datazoneCreateGlossary {
			datazone_CreateGlossary(cfg, client)
			return
		}
		if _datazoneCreateGlossaryTerm {
			datazone_CreateGlossaryTerm(cfg, client)
			return
		}
		if _datazoneCreateGroupProfile {
			datazone_CreateGroupProfile(cfg, client)
			return
		}
		if _datazoneCreateListingChangeSet {
			datazone_CreateListingChangeSet(cfg, client)
			return
		}
		if _datazoneCreateProject {
			datazone_CreateProject(cfg, client)
			return
		}
		if _datazoneCreateProjectMembership {
			datazone_CreateProjectMembership(cfg, client)
			return
		}
		if _datazoneCreateProjectProfile {
			datazone_CreateProjectProfile(cfg, client)
			return
		}
		if _datazoneCreateRule {
			datazone_CreateRule(cfg, client)
			return
		}
		if _datazoneCreateSubscriptionGrant {
			datazone_CreateSubscriptionGrant(cfg, client)
			return
		}
		if _datazoneCreateSubscriptionRequest {
			datazone_CreateSubscriptionRequest(cfg, client)
			return
		}
		if _datazoneCreateSubscriptionTarget {
			datazone_CreateSubscriptionTarget(cfg, client)
			return
		}
		if _datazoneCreateUserProfile {
			datazone_CreateUserProfile(cfg, client)
			return
		}
		if _datazoneDeleteAccountPool {
			datazone_DeleteAccountPool(cfg, client)
			return
		}
		if _datazoneDeleteAsset {
			datazone_DeleteAsset(cfg, client)
			return
		}
		if _datazoneDeleteAssetFilter {
			datazone_DeleteAssetFilter(cfg, client)
			return
		}
		if _datazoneDeleteAssetType {
			datazone_DeleteAssetType(cfg, client)
			return
		}
		if _datazoneDeleteConnection {
			datazone_DeleteConnection(cfg, client)
			return
		}
		if _datazoneDeleteDataExportConfiguration {
			datazone_DeleteDataExportConfiguration(cfg, client)
			return
		}
		if _datazoneDeleteDataProduct {
			datazone_DeleteDataProduct(cfg, client)
			return
		}
		if _datazoneDeleteDataSource {
			datazone_DeleteDataSource(cfg, client)
			return
		}
		if _datazoneDeleteDomain {
			datazone_DeleteDomain(cfg, client)
			return
		}
		if _datazoneDeleteDomainUnit {
			datazone_DeleteDomainUnit(cfg, client)
			return
		}
		if _datazoneDeleteEnvironment {
			datazone_DeleteEnvironment(cfg, client)
			return
		}
		if _datazoneDeleteEnvironmentAction {
			datazone_DeleteEnvironmentAction(cfg, client)
			return
		}
		if _datazoneDeleteEnvironmentBlueprint {
			datazone_DeleteEnvironmentBlueprint(cfg, client)
			return
		}
		if _datazoneDeleteEnvironmentBlueprintConfiguration {
			datazone_DeleteEnvironmentBlueprintConfiguration(cfg, client)
			return
		}
		if _datazoneDeleteEnvironmentProfile {
			datazone_DeleteEnvironmentProfile(cfg, client)
			return
		}
		if _datazoneDeleteFormType {
			datazone_DeleteFormType(cfg, client)
			return
		}
		if _datazoneDeleteGlossary {
			datazone_DeleteGlossary(cfg, client)
			return
		}
		if _datazoneDeleteGlossaryTerm {
			datazone_DeleteGlossaryTerm(cfg, client)
			return
		}
		if _datazoneDeleteListing {
			datazone_DeleteListing(cfg, client)
			return
		}
		if _datazoneDeleteProject {
			datazone_DeleteProject(cfg, client)
			return
		}
		if _datazoneDeleteProjectMembership {
			datazone_DeleteProjectMembership(cfg, client)
			return
		}
		if _datazoneDeleteProjectProfile {
			datazone_DeleteProjectProfile(cfg, client)
			return
		}
		if _datazoneDeleteRule {
			datazone_DeleteRule(cfg, client)
			return
		}
		if _datazoneDeleteSubscriptionGrant {
			datazone_DeleteSubscriptionGrant(cfg, client)
			return
		}
		if _datazoneDeleteSubscriptionRequest {
			datazone_DeleteSubscriptionRequest(cfg, client)
			return
		}
		if _datazoneDeleteSubscriptionTarget {
			datazone_DeleteSubscriptionTarget(cfg, client)
			return
		}
		if _datazoneDeleteTimeSeriesDataPoints {
			datazone_DeleteTimeSeriesDataPoints(cfg, client)
			return
		}
		if _datazoneDisassociateEnvironmentRole {
			datazone_DisassociateEnvironmentRole(cfg, client)
			return
		}
		if _datazoneDisassociateGovernedTerms {
			datazone_DisassociateGovernedTerms(cfg, client)
			return
		}
		if _datazoneGetAccountPool {
			datazone_GetAccountPool(cfg, client)
			return
		}
		if _datazoneGetAsset {
			datazone_GetAsset(cfg, client)
			return
		}
		if _datazoneGetAssetFilter {
			datazone_GetAssetFilter(cfg, client)
			return
		}
		if _datazoneGetAssetType {
			datazone_GetAssetType(cfg, client)
			return
		}
		if _datazoneGetConnection {
			datazone_GetConnection(cfg, client)
			return
		}
		if _datazoneGetDataExportConfiguration {
			datazone_GetDataExportConfiguration(cfg, client)
			return
		}
		if _datazoneGetDataProduct {
			datazone_GetDataProduct(cfg, client)
			return
		}
		if _datazoneGetDataSource {
			datazone_GetDataSource(cfg, client)
			return
		}
		if _datazoneGetDataSourceRun {
			datazone_GetDataSourceRun(cfg, client)
			return
		}
		if _datazoneGetDomain {
			datazone_GetDomain(cfg, client)
			return
		}
		if _datazoneGetDomainUnit {
			datazone_GetDomainUnit(cfg, client)
			return
		}
		if _datazoneGetEnvironment {
			datazone_GetEnvironment(cfg, client)
			return
		}
		if _datazoneGetEnvironmentAction {
			datazone_GetEnvironmentAction(cfg, client)
			return
		}
		if _datazoneGetEnvironmentBlueprint {
			datazone_GetEnvironmentBlueprint(cfg, client)
			return
		}
		if _datazoneGetEnvironmentBlueprintConfiguration {
			datazone_GetEnvironmentBlueprintConfiguration(cfg, client)
			return
		}
		if _datazoneGetEnvironmentCredentials {
			datazone_GetEnvironmentCredentials(cfg, client)
			return
		}
		if _datazoneGetEnvironmentProfile {
			datazone_GetEnvironmentProfile(cfg, client)
			return
		}
		if _datazoneGetFormType {
			datazone_GetFormType(cfg, client)
			return
		}
		if _datazoneGetGlossary {
			datazone_GetGlossary(cfg, client)
			return
		}
		if _datazoneGetGlossaryTerm {
			datazone_GetGlossaryTerm(cfg, client)
			return
		}
		if _datazoneGetGroupProfile {
			datazone_GetGroupProfile(cfg, client)
			return
		}
		if _datazoneGetIamPortalLoginUrl {
			datazone_GetIamPortalLoginUrl(cfg, client)
			return
		}
		if _datazoneGetJobRun {
			datazone_GetJobRun(cfg, client)
			return
		}
		if _datazoneGetLineageEvent {
			datazone_GetLineageEvent(cfg, client)
			return
		}
		if _datazoneGetLineageNode {
			datazone_GetLineageNode(cfg, client)
			return
		}
		if _datazoneGetListing {
			datazone_GetListing(cfg, client)
			return
		}
		if _datazoneGetMetadataGenerationRun {
			datazone_GetMetadataGenerationRun(cfg, client)
			return
		}
		if _datazoneGetProject {
			datazone_GetProject(cfg, client)
			return
		}
		if _datazoneGetProjectProfile {
			datazone_GetProjectProfile(cfg, client)
			return
		}
		if _datazoneGetRule {
			datazone_GetRule(cfg, client)
			return
		}
		if _datazoneGetSubscription {
			datazone_GetSubscription(cfg, client)
			return
		}
		if _datazoneGetSubscriptionGrant {
			datazone_GetSubscriptionGrant(cfg, client)
			return
		}
		if _datazoneGetSubscriptionRequestDetails {
			datazone_GetSubscriptionRequestDetails(cfg, client)
			return
		}
		if _datazoneGetSubscriptionTarget {
			datazone_GetSubscriptionTarget(cfg, client)
			return
		}
		if _datazoneGetTimeSeriesDataPoint {
			datazone_GetTimeSeriesDataPoint(cfg, client)
			return
		}
		if _datazoneGetUserProfile {
			datazone_GetUserProfile(cfg, client)
			return
		}
		if _datazoneListAccountPools {
			datazone_ListAccountPools(cfg, client)
			return
		}
		if _datazoneListAccountsInAccountPool {
			datazone_ListAccountsInAccountPool(cfg, client)
			return
		}
		if _datazoneListAssetFilters {
			datazone_ListAssetFilters(cfg, client)
			return
		}
		if _datazoneListAssetRevisions {
			datazone_ListAssetRevisions(cfg, client)
			return
		}
		if _datazoneListConnections {
			datazone_ListConnections(cfg, client)
			return
		}
		if _datazoneListDataProductRevisions {
			datazone_ListDataProductRevisions(cfg, client)
			return
		}
		if _datazoneListDataSourceRunActivities {
			datazone_ListDataSourceRunActivities(cfg, client)
			return
		}
		if _datazoneListDataSourceRuns {
			datazone_ListDataSourceRuns(cfg, client)
			return
		}
		if _datazoneListDataSources {
			datazone_ListDataSources(cfg, client)
			return
		}
		if _datazoneListDomainUnitsForParent {
			datazone_ListDomainUnitsForParent(cfg, client)
			return
		}
		if _datazoneListDomains {
			datazone_ListDomains(cfg, client)
			return
		}
		if _datazoneListEntityOwners {
			datazone_ListEntityOwners(cfg, client)
			return
		}
		if _datazoneListEnvironmentActions {
			datazone_ListEnvironmentActions(cfg, client)
			return
		}
		if _datazoneListEnvironmentBlueprintConfigurations {
			datazone_ListEnvironmentBlueprintConfigurations(cfg, client)
			return
		}
		if _datazoneListEnvironmentBlueprints {
			datazone_ListEnvironmentBlueprints(cfg, client)
			return
		}
		if _datazoneListEnvironmentProfiles {
			datazone_ListEnvironmentProfiles(cfg, client)
			return
		}
		if _datazoneListEnvironments {
			datazone_ListEnvironments(cfg, client)
			return
		}
		if _datazoneListJobRuns {
			datazone_ListJobRuns(cfg, client)
			return
		}
		if _datazoneListLineageEvents {
			datazone_ListLineageEvents(cfg, client)
			return
		}
		if _datazoneListLineageNodeHistory {
			datazone_ListLineageNodeHistory(cfg, client)
			return
		}
		if _datazoneListMetadataGenerationRuns {
			datazone_ListMetadataGenerationRuns(cfg, client)
			return
		}
		if _datazoneListNotifications {
			datazone_ListNotifications(cfg, client)
			return
		}
		if _datazoneListPolicyGrants {
			datazone_ListPolicyGrants(cfg, client)
			return
		}
		if _datazoneListProjectMemberships {
			datazone_ListProjectMemberships(cfg, client)
			return
		}
		if _datazoneListProjectProfiles {
			datazone_ListProjectProfiles(cfg, client)
			return
		}
		if _datazoneListProjects {
			datazone_ListProjects(cfg, client)
			return
		}
		if _datazoneListRules {
			datazone_ListRules(cfg, client)
			return
		}
		if _datazoneListSubscriptionGrants {
			datazone_ListSubscriptionGrants(cfg, client)
			return
		}
		if _datazoneListSubscriptionRequests {
			datazone_ListSubscriptionRequests(cfg, client)
			return
		}
		if _datazoneListSubscriptionTargets {
			datazone_ListSubscriptionTargets(cfg, client)
			return
		}
		if _datazoneListSubscriptions {
			datazone_ListSubscriptions(cfg, client)
			return
		}
		if _datazoneListTagsForResource {
			datazone_ListTagsForResource(cfg, client)
			return
		}
		if _datazoneListTimeSeriesDataPoints {
			datazone_ListTimeSeriesDataPoints(cfg, client)
			return
		}
		if _datazonePostLineageEvent {
			datazone_PostLineageEvent(cfg, client)
			return
		}
		if _datazonePostTimeSeriesDataPoints {
			datazone_PostTimeSeriesDataPoints(cfg, client)
			return
		}
		if _datazonePutDataExportConfiguration {
			datazone_PutDataExportConfiguration(cfg, client)
			return
		}
		if _datazonePutEnvironmentBlueprintConfiguration {
			datazone_PutEnvironmentBlueprintConfiguration(cfg, client)
			return
		}
		if _datazoneQueryGraph {
			datazone_QueryGraph(cfg, client)
			return
		}
		if _datazoneRejectPredictions {
			datazone_RejectPredictions(cfg, client)
			return
		}
		if _datazoneRejectSubscriptionRequest {
			datazone_RejectSubscriptionRequest(cfg, client)
			return
		}
		if _datazoneRemoveEntityOwner {
			datazone_RemoveEntityOwner(cfg, client)
			return
		}
		if _datazoneRemovePolicyGrant {
			datazone_RemovePolicyGrant(cfg, client)
			return
		}
		if _datazoneRevokeSubscription {
			datazone_RevokeSubscription(cfg, client)
			return
		}
		if _datazoneSearch {
			datazone_Search(cfg, client)
			return
		}
		if _datazoneSearchGroupProfiles {
			datazone_SearchGroupProfiles(cfg, client)
			return
		}
		if _datazoneSearchListings {
			datazone_SearchListings(cfg, client)
			return
		}
		if _datazoneSearchTypes {
			datazone_SearchTypes(cfg, client)
			return
		}
		if _datazoneSearchUserProfiles {
			datazone_SearchUserProfiles(cfg, client)
			return
		}
		if _datazoneStartDataSourceRun {
			datazone_StartDataSourceRun(cfg, client)
			return
		}
		if _datazoneStartMetadataGenerationRun {
			datazone_StartMetadataGenerationRun(cfg, client)
			return
		}
		if _datazoneTagResource {
			datazone_TagResource(cfg, client)
			return
		}
		if _datazoneUntagResource {
			datazone_UntagResource(cfg, client)
			return
		}
		if _datazoneUpdateAccountPool {
			datazone_UpdateAccountPool(cfg, client)
			return
		}
		if _datazoneUpdateAssetFilter {
			datazone_UpdateAssetFilter(cfg, client)
			return
		}
		if _datazoneUpdateConnection {
			datazone_UpdateConnection(cfg, client)
			return
		}
		if _datazoneUpdateDataSource {
			datazone_UpdateDataSource(cfg, client)
			return
		}
		if _datazoneUpdateDomain {
			datazone_UpdateDomain(cfg, client)
			return
		}
		if _datazoneUpdateDomainUnit {
			datazone_UpdateDomainUnit(cfg, client)
			return
		}
		if _datazoneUpdateEnvironment {
			datazone_UpdateEnvironment(cfg, client)
			return
		}
		if _datazoneUpdateEnvironmentAction {
			datazone_UpdateEnvironmentAction(cfg, client)
			return
		}
		if _datazoneUpdateEnvironmentBlueprint {
			datazone_UpdateEnvironmentBlueprint(cfg, client)
			return
		}
		if _datazoneUpdateEnvironmentProfile {
			datazone_UpdateEnvironmentProfile(cfg, client)
			return
		}
		if _datazoneUpdateGlossary {
			datazone_UpdateGlossary(cfg, client)
			return
		}
		if _datazoneUpdateGlossaryTerm {
			datazone_UpdateGlossaryTerm(cfg, client)
			return
		}
		if _datazoneUpdateGroupProfile {
			datazone_UpdateGroupProfile(cfg, client)
			return
		}
		if _datazoneUpdateProject {
			datazone_UpdateProject(cfg, client)
			return
		}
		if _datazoneUpdateProjectProfile {
			datazone_UpdateProjectProfile(cfg, client)
			return
		}
		if _datazoneUpdateRootDomainUnitOwner {
			datazone_UpdateRootDomainUnitOwner(cfg, client)
			return
		}
		if _datazoneUpdateRule {
			datazone_UpdateRule(cfg, client)
			return
		}
		if _datazoneUpdateSubscriptionGrantStatus {
			datazone_UpdateSubscriptionGrantStatus(cfg, client)
			return
		}
		if _datazoneUpdateSubscriptionRequest {
			datazone_UpdateSubscriptionRequest(cfg, client)
			return
		}
		if _datazoneUpdateSubscriptionTarget {
			datazone_UpdateSubscriptionTarget(cfg, client)
			return
		}
		if _datazoneUpdateUserProfile {
			datazone_UpdateUserProfile(cfg, client)
			return
		}

	},
}

var (
	_datazoneAcceptPredictions                       bool
	_datazoneAcceptSubscriptionRequest               bool
	_datazoneAddEntityOwner                          bool
	_datazoneAddPolicyGrant                          bool
	_datazoneAssociateEnvironmentRole                bool
	_datazoneAssociateGovernedTerms                  bool
	_datazoneBatchGetAttributesMetadata              bool
	_datazoneBatchPutAttributesMetadata              bool
	_datazoneCancelMetadataGenerationRun             bool
	_datazoneCancelSubscription                      bool
	_datazoneCreateAccountPool                       bool
	_datazoneCreateAsset                             bool
	_datazoneCreateAssetFilter                       bool
	_datazoneCreateAssetRevision                     bool
	_datazoneCreateAssetType                         bool
	_datazoneCreateConnection                        bool
	_datazoneCreateDataProduct                       bool
	_datazoneCreateDataProductRevision               bool
	_datazoneCreateDataSource                        bool
	_datazoneCreateDomain                            bool
	_datazoneCreateDomainUnit                        bool
	_datazoneCreateEnvironment                       bool
	_datazoneCreateEnvironmentAction                 bool
	_datazoneCreateEnvironmentBlueprint              bool
	_datazoneCreateEnvironmentProfile                bool
	_datazoneCreateFormType                          bool
	_datazoneCreateGlossary                          bool
	_datazoneCreateGlossaryTerm                      bool
	_datazoneCreateGroupProfile                      bool
	_datazoneCreateListingChangeSet                  bool
	_datazoneCreateProject                           bool
	_datazoneCreateProjectMembership                 bool
	_datazoneCreateProjectProfile                    bool
	_datazoneCreateRule                              bool
	_datazoneCreateSubscriptionGrant                 bool
	_datazoneCreateSubscriptionRequest               bool
	_datazoneCreateSubscriptionTarget                bool
	_datazoneCreateUserProfile                       bool
	_datazoneDeleteAccountPool                       bool
	_datazoneDeleteAsset                             bool
	_datazoneDeleteAssetFilter                       bool
	_datazoneDeleteAssetType                         bool
	_datazoneDeleteConnection                        bool
	_datazoneDeleteDataExportConfiguration           bool
	_datazoneDeleteDataProduct                       bool
	_datazoneDeleteDataSource                        bool
	_datazoneDeleteDomain                            bool
	_datazoneDeleteDomainUnit                        bool
	_datazoneDeleteEnvironment                       bool
	_datazoneDeleteEnvironmentAction                 bool
	_datazoneDeleteEnvironmentBlueprint              bool
	_datazoneDeleteEnvironmentBlueprintConfiguration bool
	_datazoneDeleteEnvironmentProfile                bool
	_datazoneDeleteFormType                          bool
	_datazoneDeleteGlossary                          bool
	_datazoneDeleteGlossaryTerm                      bool
	_datazoneDeleteListing                           bool
	_datazoneDeleteProject                           bool
	_datazoneDeleteProjectMembership                 bool
	_datazoneDeleteProjectProfile                    bool
	_datazoneDeleteRule                              bool
	_datazoneDeleteSubscriptionGrant                 bool
	_datazoneDeleteSubscriptionRequest               bool
	_datazoneDeleteSubscriptionTarget                bool
	_datazoneDeleteTimeSeriesDataPoints              bool
	_datazoneDisassociateEnvironmentRole             bool
	_datazoneDisassociateGovernedTerms               bool
	_datazoneGetAccountPool                          bool
	_datazoneGetAsset                                bool
	_datazoneGetAssetFilter                          bool
	_datazoneGetAssetType                            bool
	_datazoneGetConnection                           bool
	_datazoneGetDataExportConfiguration              bool
	_datazoneGetDataProduct                          bool
	_datazoneGetDataSource                           bool
	_datazoneGetDataSourceRun                        bool
	_datazoneGetDomain                               bool
	_datazoneGetDomainUnit                           bool
	_datazoneGetEnvironment                          bool
	_datazoneGetEnvironmentAction                    bool
	_datazoneGetEnvironmentBlueprint                 bool
	_datazoneGetEnvironmentBlueprintConfiguration    bool
	_datazoneGetEnvironmentCredentials               bool
	_datazoneGetEnvironmentProfile                   bool
	_datazoneGetFormType                             bool
	_datazoneGetGlossary                             bool
	_datazoneGetGlossaryTerm                         bool
	_datazoneGetGroupProfile                         bool
	_datazoneGetIamPortalLoginUrl                    bool
	_datazoneGetJobRun                               bool
	_datazoneGetLineageEvent                         bool
	_datazoneGetLineageNode                          bool
	_datazoneGetListing                              bool
	_datazoneGetMetadataGenerationRun                bool
	_datazoneGetProject                              bool
	_datazoneGetProjectProfile                       bool
	_datazoneGetRule                                 bool
	_datazoneGetSubscription                         bool
	_datazoneGetSubscriptionGrant                    bool
	_datazoneGetSubscriptionRequestDetails           bool
	_datazoneGetSubscriptionTarget                   bool
	_datazoneGetTimeSeriesDataPoint                  bool
	_datazoneGetUserProfile                          bool
	_datazoneListAccountPools                        bool
	_datazoneListAccountsInAccountPool               bool
	_datazoneListAssetFilters                        bool
	_datazoneListAssetRevisions                      bool
	_datazoneListConnections                         bool
	_datazoneListDataProductRevisions                bool
	_datazoneListDataSourceRunActivities             bool
	_datazoneListDataSourceRuns                      bool
	_datazoneListDataSources                         bool
	_datazoneListDomainUnitsForParent                bool
	_datazoneListDomains                             bool
	_datazoneListEntityOwners                        bool
	_datazoneListEnvironmentActions                  bool
	_datazoneListEnvironmentBlueprintConfigurations  bool
	_datazoneListEnvironmentBlueprints               bool
	_datazoneListEnvironmentProfiles                 bool
	_datazoneListEnvironments                        bool
	_datazoneListJobRuns                             bool
	_datazoneListLineageEvents                       bool
	_datazoneListLineageNodeHistory                  bool
	_datazoneListMetadataGenerationRuns              bool
	_datazoneListNotifications                       bool
	_datazoneListPolicyGrants                        bool
	_datazoneListProjectMemberships                  bool
	_datazoneListProjectProfiles                     bool
	_datazoneListProjects                            bool
	_datazoneListRules                               bool
	_datazoneListSubscriptionGrants                  bool
	_datazoneListSubscriptionRequests                bool
	_datazoneListSubscriptionTargets                 bool
	_datazoneListSubscriptions                       bool
	_datazoneListTagsForResource                     bool
	_datazoneListTimeSeriesDataPoints                bool
	_datazonePostLineageEvent                        bool
	_datazonePostTimeSeriesDataPoints                bool
	_datazonePutDataExportConfiguration              bool
	_datazonePutEnvironmentBlueprintConfiguration    bool
	_datazoneQueryGraph                              bool
	_datazoneRejectPredictions                       bool
	_datazoneRejectSubscriptionRequest               bool
	_datazoneRemoveEntityOwner                       bool
	_datazoneRemovePolicyGrant                       bool
	_datazoneRevokeSubscription                      bool
	_datazoneSearch                                  bool
	_datazoneSearchGroupProfiles                     bool
	_datazoneSearchListings                          bool
	_datazoneSearchTypes                             bool
	_datazoneSearchUserProfiles                      bool
	_datazoneStartDataSourceRun                      bool
	_datazoneStartMetadataGenerationRun              bool
	_datazoneTagResource                             bool
	_datazoneUntagResource                           bool
	_datazoneUpdateAccountPool                       bool
	_datazoneUpdateAssetFilter                       bool
	_datazoneUpdateConnection                        bool
	_datazoneUpdateDataSource                        bool
	_datazoneUpdateDomain                            bool
	_datazoneUpdateDomainUnit                        bool
	_datazoneUpdateEnvironment                       bool
	_datazoneUpdateEnvironmentAction                 bool
	_datazoneUpdateEnvironmentBlueprint              bool
	_datazoneUpdateEnvironmentProfile                bool
	_datazoneUpdateGlossary                          bool
	_datazoneUpdateGlossaryTerm                      bool
	_datazoneUpdateGroupProfile                      bool
	_datazoneUpdateProject                           bool
	_datazoneUpdateProjectProfile                    bool
	_datazoneUpdateRootDomainUnitOwner               bool
	_datazoneUpdateRule                              bool
	_datazoneUpdateSubscriptionGrantStatus           bool
	_datazoneUpdateSubscriptionRequest               bool
	_datazoneUpdateSubscriptionTarget                bool
	_datazoneUpdateUserProfile                       bool

	_datazoneAcceptChoices                     string
	_datazoneAcceptRule                        string
	_datazoneAccountSource                     string
	_datazoneAction                            string
	_datazoneAdditionalAttributes              string
	_datazoneAfterTimestamp                    string
	_datazoneAggregations                      string
	_datazoneAllowCustomProjectResourceTags    string
	_datazoneApplicableAssetTypes              []string
	_datazoneApproverProjectId                 string
	_datazoneAssetFormsInput                   string
	_datazoneAssetIdentifier                   string
	_datazoneAssetPermissions                  string
	_datazoneAssetScopes                       string
	_datazoneAssetTargetNames                  string
	_datazoneAssetTypes                        []string
	_datazoneAttributeIdentifiers              []string
	_datazoneAttributes                        string
	_datazoneAuthorizedPrincipals              []string
	_datazoneAwsAccountId                      string
	_datazoneAwsAccountRegion                  string
	_datazoneAwsLocation                       string
	_datazoneBeforeTimestamp                   string
	_datazoneBlueprintVersion                  string
	_datazoneClientToken                       string
	_datazoneConfiguration                     string
	_datazoneConnectionIdentifier              string
	_datazoneCurrentOwner                      string
	_datazoneDataProduct                       string
	_datazoneDataSourceIdentifier              string
	_datazoneDecisionComment                   string
	_datazoneDeploymentOrder                   string
	_datazoneDescription                       string
	_datazoneDesignation                       string
	_datazoneDetail                            string
	_datazoneDirection                         string
	_datazoneDomainExecutionRole               string
	_datazoneDomainIdentifier                  string
	_datazoneDomainUnitId                      string
	_datazoneDomainUnitIdentifier              string
	_datazoneDomainVersion                     string
	_datazoneEnableExport                      string
	_datazoneEnableSetting                     string
	_datazoneEnableTrustedIdentityPropagation  string
	_datazoneEnabledRegions                    []string
	_datazoneEncryptionConfiguration           string
	_datazoneEndedAt                           string
	_datazoneEntityIdentifier                  string
	_datazoneEntityRevision                    string
	_datazoneEntityType                        string
	_datazoneEnvironmentAccountIdentifier      string
	_datazoneEnvironmentAccountRegion          string
	_datazoneEnvironmentBlueprintIdentifier    string
	_datazoneEnvironmentConfigurationId        string
	_datazoneEnvironmentConfigurations         string
	_datazoneEnvironmentDeploymentDetails      string
	_datazoneEnvironmentId                     string
	_datazoneEnvironmentIdentifier             string
	_datazoneEnvironmentProfileIdentifier      string
	_datazoneEnvironmentRoleArn                string
	_datazoneEnvironmentRolePermissionBoundary string
	_datazoneEvent                             string
	_datazoneEventTimestamp                    string
	_datazoneEventTimestampGTE                 string
	_datazoneEventTimestampLTE                 string
	_datazoneExternalIdentifier                string
	_datazoneFailureCause                      string
	_datazoneFilters                           string
	_datazoneFormName                          string
	_datazoneFormTypeIdentifier                string
	_datazoneForms                             string
	_datazoneFormsInput                        string
	_datazoneGlobalParameters                  string
	_datazoneGlossaryIdentifier                string
	_datazoneGlossaryTerms                     []string
	_datazoneGovernedGlossaryTerms             []string
	_datazoneGrantIdentifier                   string
	_datazoneGrantedEntity                     string
	_datazoneGroupIdentifier                   string
	_datazoneGroupType                         string
	_datazoneIdentifier                        string
	_datazoneIncludeCascaded                   string
	_datazoneIncludeChildDomainUnits           string
	_datazoneItems                             string
	_datazoneJobIdentifier                     string
	_datazoneKmsKeyIdentifier                  string
	_datazoneListingRevision                   string
	_datazoneLongDescription                   string
	_datazoneManageAccessRole                  string
	_datazoneManageAccessRoleArn               string
	_datazoneManaged                           string
	_datazoneMatch                             string
	_datazoneMaxResults                        string
	_datazoneMember                            string
	_datazoneMetadataForms                     string
	_datazoneModel                             string
	_datazoneName                              string
	_datazoneNewOwner                          string
	_datazoneNextToken                         string
	_datazoneOwner                             string
	_datazoneOwningGroupId                     string
	_datazoneOwningIamPrincipalArn             string
	_datazoneOwningProjectId                   string
	_datazoneOwningProjectIdentifier           string
	_datazoneOwningUserId                      string
	_datazoneParameters                        string
	_datazoneParentDomainUnitIdentifier        string
	_datazonePolicyType                        string
	_datazonePredictionConfiguration           string
	_datazonePrincipal                         string
	_datazoneProcessingStatus                  string
	_datazoneProjectIdentifier                 string
	_datazoneProjectIds                        []string
	_datazoneProjectProfileId                  string
	_datazoneProjectProfileVersion             string
	_datazoneProjectResourceTags               string
	_datazoneProjectResourceTagsDescription    string
	_datazoneProps                             string
	_datazoneProvider                          string
	_datazoneProvisioningConfigurations        string
	_datazoneProvisioningProperties            string
	_datazoneProvisioningRoleArn               string
	_datazonePublishOnImport                   string
	_datazoneRecommendation                    string
	_datazoneRegionalParameters                string
	_datazoneRejectChoices                     string
	_datazoneRejectRule                        string
	_datazoneRequestReason                     string
	_datazoneResolutionStrategy                string
	_datazoneResourceArn                       string
	_datazoneResourceTags                      string
	_datazoneRetainPermissions                 string
	_datazoneRetainPermissionsOnRevokeFailure  string
	_datazoneRevision                          string
	_datazoneRuleType                          string
	_datazoneSchedule                          string
	_datazoneScope                             string
	_datazoneSearchIn                          string
	_datazoneSearchScope                       string
	_datazoneSearchText                        string
	_datazoneServiceRole                       string
	_datazoneShortDescription                  string
	_datazoneSingleSignOn                      string
	_datazoneSkipDeletionCheck                 string
	_datazoneSort                              string
	_datazoneSortBy                            string
	_datazoneSortOrder                         string
	_datazoneStartedAt                         string
	_datazoneStatus                            string
	_datazoneSubjects                          []string
	_datazoneSubscribedListingId               string
	_datazoneSubscribedListings                string
	_datazoneSubscribedPrincipals              string
	_datazoneSubscriptionGrantCreationMode     string
	_datazoneSubscriptionId                    string
	_datazoneSubscriptionRequestIdentifier     string
	_datazoneSubscriptionTargetConfig          string
	_datazoneSubscriptionTargetId              string
	_datazoneSubscriptionTargetIdentifier      string
	_datazoneTagKeys                           []string
	_datazoneTags                              string
	_datazoneTarget                            string
	_datazoneTargetIdentifier                  string
	_datazoneTargetName                        string
	_datazoneTargetType                        string
	_datazoneTaskStatus                        string
	_datazoneTermRelations                     string
	_datazoneTimestampAfter                    string
	_datazoneTimestampBefore                   string
	_datazoneType                              string
	_datazoneTypeIdentifier                    string
	_datazoneTypeRevision                      string
	_datazoneTypes                             string
	_datazoneUsageRestrictions                 string
	_datazoneUserIdentifier                    string
	_datazoneUserParameters                    string
	_datazoneUserType                          string
	_datazoneWithSecret                        string
)

// Accepts automatically generated business-friendly metadata for your Amazon
// DataZone assets.
func datazone_AcceptPredictions(cfg aws.Config, client *datazone.Client) {
	input := &datazone.AcceptPredictionsInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneAcceptChoices) > 0 {
		if err := assignInputField(input, "AcceptChoices", _datazoneAcceptChoices); err != nil {
			log.Errorf("invalid --accept-choices: %s", err.Error())
			return
		}
	}
	if len(_datazoneAcceptRule) > 0 {
		if err := assignInputField(input, "AcceptRule", _datazoneAcceptRule); err != nil {
			log.Errorf("invalid --accept-rule: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneRevision) > 0 {
		input.Revision = aws.String(_datazoneRevision)
	}

	if resp, err := client.AcceptPredictions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Accepts a subscription request to a specific asset.
func datazone_AcceptSubscriptionRequest(cfg aws.Config, client *datazone.Client) {
	input := &datazone.AcceptSubscriptionRequestInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneAssetPermissions) > 0 {
		if err := assignInputField(input, "AssetPermissions", _datazoneAssetPermissions); err != nil {
			log.Errorf("invalid --asset-permissions: %s", err.Error())
			return
		}
	}
	if len(_datazoneAssetScopes) > 0 {
		if err := assignInputField(input, "AssetScopes", _datazoneAssetScopes); err != nil {
			log.Errorf("invalid --asset-scopes: %s", err.Error())
			return
		}
	}
	if len(_datazoneDecisionComment) > 0 {
		input.DecisionComment = aws.String(_datazoneDecisionComment)
	}

	if resp, err := client.AcceptSubscriptionRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the owner of an entity (a domain unit).
func datazone_AddEntityOwner(cfg aws.Config, client *datazone.Client) {
	input := &datazone.AddEntityOwnerInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.DataZoneEntityType, // Required
		// Owner: types.OwnerProperties, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneOwner) > 0 {
		if err := assignInputField(input, "Owner", _datazoneOwner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}

	if resp, err := client.AddEntityOwner(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a policy grant (an authorization policy) to a specified entity, including
// domain units, environment blueprint configurations, or environment profiles.
func datazone_AddPolicyGrant(cfg aws.Config, client *datazone.Client) {
	input := &datazone.AddPolicyGrantInput{
		// Detail: types.PolicyGrantDetail, // Required
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.TargetEntityType, // Required
		// PolicyType: types.ManagedPolicyType, // Required
		// Principal: types.PolicyGrantPrincipal, // Required
	}

	if len(_datazoneDetail) > 0 {
		if err := assignInputField(input, "Detail", _datazoneDetail); err != nil {
			log.Errorf("invalid --detail: %s", err.Error())
			return
		}
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazonePolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _datazonePolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_datazonePrincipal) > 0 {
		if err := assignInputField(input, "Principal", _datazonePrincipal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}

	if resp, err := client.AddPolicyGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the environment role in Amazon DataZone.
func datazone_AssociateEnvironmentRole(cfg aws.Config, client *datazone.Client) {
	input := &datazone.AssociateEnvironmentRoleInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// EnvironmentRoleArn: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneEnvironmentRoleArn) > 0 {
		input.EnvironmentRoleArn = aws.String(_datazoneEnvironmentRoleArn)
	}

	if resp, err := client.AssociateEnvironmentRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates governed terms with an asset.
func datazone_AssociateGovernedTerms(cfg aws.Config, client *datazone.Client) {
	input := &datazone.AssociateGovernedTermsInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.GovernedEntityType, // Required
		// GovernedGlossaryTerms: []string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneGovernedGlossaryTerms) > 0 {
		input.GovernedGlossaryTerms = append([]string(nil), _datazoneGovernedGlossaryTerms...)
	}

	if resp, err := client.AssociateGovernedTerms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the attribute metadata.
func datazone_BatchGetAttributesMetadata(cfg aws.Config, client *datazone.Client) {
	input := &datazone.BatchGetAttributesMetadataInput{
		// AttributeIdentifiers: []string, // Required
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.AttributeEntityType, // Required
	}

	if len(_datazoneAttributeIdentifiers) > 0 {
		input.AttributeIdentifiers = append([]string(nil), _datazoneAttributeIdentifiers...)
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneEntityRevision) > 0 {
		input.EntityRevision = aws.String(_datazoneEntityRevision)
	}

	if resp, err := client.BatchGetAttributesMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Writes the attribute metadata.
func datazone_BatchPutAttributesMetadata(cfg aws.Config, client *datazone.Client) {
	input := &datazone.BatchPutAttributesMetadataInput{
		// Attributes: []types.AttributeInput, // Required
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.AttributeEntityType, // Required
	}

	if len(_datazoneAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _datazoneAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}

	if resp, err := client.BatchPutAttributesMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the metadata generation run.
// Prerequisites:
//
// - The run must exist and be in a cancelable status (e.g., SUBMITTED,
// IN_PROGRESS).
//
// - Runs in SUCCEEDED status cannot be cancelled.
//
// - User must have access to the run and cancel permissions.
func datazone_CancelMetadataGenerationRun(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CancelMetadataGenerationRunInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.CancelMetadataGenerationRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the subscription to the specified asset.
func datazone_CancelSubscription(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CancelSubscriptionInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.CancelSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an account pool.
func datazone_CreateAccountPool(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateAccountPoolInput{
		// AccountSource: types.AccountSource, // Required
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
		// ResolutionStrategy: types.ResolutionStrategy, // Required
	}

	if len(_datazoneAccountSource) > 0 {
		if err := assignInputField(input, "AccountSource", _datazoneAccountSource); err != nil {
			log.Errorf("invalid --account-source: %s", err.Error())
			return
		}
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneResolutionStrategy) > 0 {
		if err := assignInputField(input, "ResolutionStrategy", _datazoneResolutionStrategy); err != nil {
			log.Errorf("invalid --resolution-strategy: %s", err.Error())
			return
		}
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}

	if resp, err := client.CreateAccountPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an asset in Amazon DataZone catalog.
// Before creating assets, make sure that the following requirements are met:
//
// - --domain-identifier must refer to an existing domain.
//
// - --owning-project-identifier must be a valid project within the domain.
//
// - Asset type must be created beforehand using create-asset-type , or be a
// supported system-defined type. For more information, see [create-asset-type].
//
// - --type-revision (if used) must match a valid revision of the asset type.
//
// - formsInput is required when it is associated as required in the asset-type .
// For more information, see [create-form-type].
//
// - Form content must include all required fields as per the form schema (e.g.,
// bucketArn ).
//
// You must invoke the following pre-requisite commands before invoking this API:
//
// [CreateFormType]
//
// [CreateAssetType]
//
// [create-asset-type]: https://docs.aws.amazon.com/cli/latest/reference/datazone/create-asset-type.html
// [create-form-type]: https://docs.aws.amazon.com/cli/latest/reference/datazone/create-form-type.html
// [CreateFormType]: https://docs.aws.amazon.com/datazone/latest/APIReference/API_CreateFormType.html
// [CreateAssetType]: https://docs.aws.amazon.com/datazone/latest/APIReference/API_CreateAssetType.html
func datazone_CreateAsset(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateAssetInput{
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
		// OwningProjectIdentifier: *string, // Required
		// TypeIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneOwningProjectIdentifier) > 0 {
		input.OwningProjectIdentifier = aws.String(_datazoneOwningProjectIdentifier)
	}
	if len(_datazoneTypeIdentifier) > 0 {
		input.TypeIdentifier = aws.String(_datazoneTypeIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneExternalIdentifier) > 0 {
		input.ExternalIdentifier = aws.String(_datazoneExternalIdentifier)
	}
	if len(_datazoneFormsInput) > 0 {
		if err := assignInputField(input, "FormsInput", _datazoneFormsInput); err != nil {
			log.Errorf("invalid --forms-input: %s", err.Error())
			return
		}
	}
	if len(_datazoneGlossaryTerms) > 0 {
		input.GlossaryTerms = append([]string(nil), _datazoneGlossaryTerms...)
	}
	if len(_datazonePredictionConfiguration) > 0 {
		if err := assignInputField(input, "PredictionConfiguration", _datazonePredictionConfiguration); err != nil {
			log.Errorf("invalid --prediction-configuration: %s", err.Error())
			return
		}
	}
	if len(_datazoneTypeRevision) > 0 {
		input.TypeRevision = aws.String(_datazoneTypeRevision)
	}

	if resp, err := client.CreateAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data asset filter.
// Asset filters provide a sophisticated way to create controlled views of data
// assets by selecting specific columns or applying row-level filters. This
// capability is crucial for organizations that need to share data while
// maintaining security and privacy controls. For example, your database might be
// filtered to show only non-PII fields to certain users, or sales data might be
// filtered by region for different regional teams. Asset filters enable
// fine-grained access control while maintaining a single source of truth.
//
// Prerequisites:
//
// - A valid domain ( --domain-identifier ) must exist.
//
// - A data asset ( --asset-identifier ) must already be created under that
// domain.
//
// - The asset must have the referenced columns available in its schema for
// column-based filtering.
//
// - You cannot specify both ( columnConfiguration , rowConfiguration )at the
// same time.
func datazone_CreateAssetFilter(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateAssetFilterInput{
		// AssetIdentifier: *string, // Required
		// Configuration: types.AssetFilterConfiguration, // Required
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_datazoneAssetIdentifier) > 0 {
		input.AssetIdentifier = aws.String(_datazoneAssetIdentifier)
	}
	if len(_datazoneConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _datazoneConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}

	if resp, err := client.CreateAssetFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a revision of the asset.
// Asset revisions represent new versions of existing assets, capturing changes to
// either the underlying data or its metadata. They maintain a historical record of
// how assets evolve over time, who made changes, and when those changes occurred.
// This versioning capability is crucial for governance and compliance, allowing
// organizations to track changes, understand their impact, and roll back if
// necessary.
//
// Prerequisites:
//
// - Asset must already exist in the domain with identifier.
//
// - formsInput is required when asset has the form type. typeRevision should be
// the latest version of form type.
//
// - The form content must include all required fields (e.g., bucketArn for
// S3ObjectCollectionForm ).
//
// - The owning project of the original asset must still exist and be active.
//
// - User must have write access to the project and domain.
func datazone_CreateAssetRevision(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateAssetRevisionInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneFormsInput) > 0 {
		if err := assignInputField(input, "FormsInput", _datazoneFormsInput); err != nil {
			log.Errorf("invalid --forms-input: %s", err.Error())
			return
		}
	}
	if len(_datazoneGlossaryTerms) > 0 {
		input.GlossaryTerms = append([]string(nil), _datazoneGlossaryTerms...)
	}
	if len(_datazonePredictionConfiguration) > 0 {
		if err := assignInputField(input, "PredictionConfiguration", _datazonePredictionConfiguration); err != nil {
			log.Errorf("invalid --prediction-configuration: %s", err.Error())
			return
		}
	}
	if len(_datazoneTypeRevision) > 0 {
		input.TypeRevision = aws.String(_datazoneTypeRevision)
	}

	if resp, err := client.CreateAssetRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom asset type.
// Prerequisites:
//
// - The formsInput field is required, however, can be passed as empty (e.g.
// -forms-input {}) .
//
// - You must have CreateAssetType permissions.
//
// - The domain-identifier and owning-project-identifier must be valid and
// active.
//
// - The name of the asset type must be unique within the domain — duplicate
// names will cause failure.
//
// - JSON input must be valid — incorrect formatting causes Invalid JSON errors.
func datazone_CreateAssetType(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateAssetTypeInput{
		// DomainIdentifier: *string, // Required
		// FormsInput: map[string]types.FormEntryInput, // Required
		// Name: *string, // Required
		// OwningProjectIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneFormsInput) > 0 {
		if err := assignInputField(input, "FormsInput", _datazoneFormsInput); err != nil {
			log.Errorf("invalid --forms-input: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneOwningProjectIdentifier) > 0 {
		input.OwningProjectIdentifier = aws.String(_datazoneOwningProjectIdentifier)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}

	if resp, err := client.CreateAssetType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new connection. In Amazon DataZone, a connection enables you to
// connect your resources (domains, projects, and environments) to external
// resources and services.
func datazone_CreateConnection(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateConnectionInput{
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneAwsLocation) > 0 {
		if err := assignInputField(input, "AwsLocation", _datazoneAwsLocation); err != nil {
			log.Errorf("invalid --aws-location: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneEnableTrustedIdentityPropagation) > 0 {
		if err := assignInputField(input, "EnableTrustedIdentityPropagation", _datazoneEnableTrustedIdentityPropagation); err != nil {
			log.Errorf("invalid --enable-trusted-identity-propagation: %s", err.Error())
			return
		}
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneProps) > 0 {
		if err := assignInputField(input, "Props", _datazoneProps); err != nil {
			log.Errorf("invalid --props: %s", err.Error())
			return
		}
	}
	if len(_datazoneScope) > 0 {
		if err := assignInputField(input, "Scope", _datazoneScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data product.
// A data product is a comprehensive package that combines data assets with their
// associated metadata, documentation, and access controls. It's designed to serve
// specific business needs or use cases, making it easier for users to find and
// consume data appropriately. Data products include important information about
// data quality, freshness, and usage guidelines, effectively bridging the gap
// between data producers and consumers while ensuring proper governance.
//
// Prerequisites:
//
// - The domain must exist and be accessible.
//
// - The owning project must be valid and active.
//
// - The name must be unique within the domain (no existing data product with
// the same name).
//
// - User must have create permissions for data products in the project.
func datazone_CreateDataProduct(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateDataProductInput{
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
		// OwningProjectIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneOwningProjectIdentifier) > 0 {
		input.OwningProjectIdentifier = aws.String(_datazoneOwningProjectIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneFormsInput) > 0 {
		if err := assignInputField(input, "FormsInput", _datazoneFormsInput); err != nil {
			log.Errorf("invalid --forms-input: %s", err.Error())
			return
		}
	}
	if len(_datazoneGlossaryTerms) > 0 {
		input.GlossaryTerms = append([]string(nil), _datazoneGlossaryTerms...)
	}
	if len(_datazoneItems) > 0 {
		if err := assignInputField(input, "Items", _datazoneItems); err != nil {
			log.Errorf("invalid --items: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data product revision.
// Prerequisites:
//
// - The original data product must exist in the given domain.
//
// - User must have permissions on the data product.
//
// - The domain must be valid and accessible.
//
// - The new revision name must comply with naming constraints (if required).
func datazone_CreateDataProductRevision(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateDataProductRevisionInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneFormsInput) > 0 {
		if err := assignInputField(input, "FormsInput", _datazoneFormsInput); err != nil {
			log.Errorf("invalid --forms-input: %s", err.Error())
			return
		}
	}
	if len(_datazoneGlossaryTerms) > 0 {
		input.GlossaryTerms = append([]string(nil), _datazoneGlossaryTerms...)
	}
	if len(_datazoneItems) > 0 {
		if err := assignInputField(input, "Items", _datazoneItems); err != nil {
			log.Errorf("invalid --items: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataProductRevision(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon DataZone data source.
func datazone_CreateDataSource(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateDataSourceInput{
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
		// ProjectIdentifier: *string, // Required
		// Type: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneProjectIdentifier) > 0 {
		input.ProjectIdentifier = aws.String(_datazoneProjectIdentifier)
	}
	if len(_datazoneType) > 0 {
		input.Type = aws.String(_datazoneType)
	}
	if len(_datazoneAssetFormsInput) > 0 {
		if err := assignInputField(input, "AssetFormsInput", _datazoneAssetFormsInput); err != nil {
			log.Errorf("invalid --asset-forms-input: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _datazoneConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_datazoneConnectionIdentifier) > 0 {
		input.ConnectionIdentifier = aws.String(_datazoneConnectionIdentifier)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneEnableSetting) > 0 {
		if err := assignInputField(input, "EnableSetting", _datazoneEnableSetting); err != nil {
			log.Errorf("invalid --enable-setting: %s", err.Error())
			return
		}
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazonePublishOnImport) > 0 {
		if err := assignInputField(input, "PublishOnImport", _datazonePublishOnImport); err != nil {
			log.Errorf("invalid --publish-on-import: %s", err.Error())
			return
		}
	}
	if len(_datazoneRecommendation) > 0 {
		if err := assignInputField(input, "Recommendation", _datazoneRecommendation); err != nil {
			log.Errorf("invalid --recommendation: %s", err.Error())
			return
		}
	}
	if len(_datazoneSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _datazoneSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon DataZone domain.
func datazone_CreateDomain(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateDomainInput{
		// DomainExecutionRole: *string, // Required
		// Name: *string, // Required
	}

	if len(_datazoneDomainExecutionRole) > 0 {
		input.DomainExecutionRole = aws.String(_datazoneDomainExecutionRole)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneDomainVersion) > 0 {
		if err := assignInputField(input, "DomainVersion", _datazoneDomainVersion); err != nil {
			log.Errorf("invalid --domain-version: %s", err.Error())
			return
		}
	}
	if len(_datazoneKmsKeyIdentifier) > 0 {
		input.KmsKeyIdentifier = aws.String(_datazoneKmsKeyIdentifier)
	}
	if len(_datazoneServiceRole) > 0 {
		input.ServiceRole = aws.String(_datazoneServiceRole)
	}
	if len(_datazoneSingleSignOn) > 0 {
		if err := assignInputField(input, "SingleSignOn", _datazoneSingleSignOn); err != nil {
			log.Errorf("invalid --single-sign-on: %s", err.Error())
			return
		}
	}
	if len(_datazoneTags) > 0 {
		if err := assignInputField(input, "Tags", _datazoneTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a domain unit in Amazon DataZone.
func datazone_CreateDomainUnit(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateDomainUnitInput{
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
		// ParentDomainUnitIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneParentDomainUnitIdentifier) > 0 {
		input.ParentDomainUnitIdentifier = aws.String(_datazoneParentDomainUnitIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}

	if resp, err := client.CreateDomainUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an Amazon DataZone environment.
func datazone_CreateEnvironment(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateEnvironmentInput{
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
		// ProjectIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneProjectIdentifier) > 0 {
		input.ProjectIdentifier = aws.String(_datazoneProjectIdentifier)
	}
	if len(_datazoneDeploymentOrder) > 0 {
		if err := assignInputField(input, "DeploymentOrder", _datazoneDeploymentOrder); err != nil {
			log.Errorf("invalid --deployment-order: %s", err.Error())
			return
		}
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneEnvironmentAccountIdentifier) > 0 {
		input.EnvironmentAccountIdentifier = aws.String(_datazoneEnvironmentAccountIdentifier)
	}
	if len(_datazoneEnvironmentAccountRegion) > 0 {
		input.EnvironmentAccountRegion = aws.String(_datazoneEnvironmentAccountRegion)
	}
	if len(_datazoneEnvironmentBlueprintIdentifier) > 0 {
		input.EnvironmentBlueprintIdentifier = aws.String(_datazoneEnvironmentBlueprintIdentifier)
	}
	if len(_datazoneEnvironmentConfigurationId) > 0 {
		input.EnvironmentConfigurationId = aws.String(_datazoneEnvironmentConfigurationId)
	}
	if len(_datazoneEnvironmentProfileIdentifier) > 0 {
		input.EnvironmentProfileIdentifier = aws.String(_datazoneEnvironmentProfileIdentifier)
	}
	if len(_datazoneGlossaryTerms) > 0 {
		input.GlossaryTerms = append([]string(nil), _datazoneGlossaryTerms...)
	}
	if len(_datazoneUserParameters) > 0 {
		if err := assignInputField(input, "UserParameters", _datazoneUserParameters); err != nil {
			log.Errorf("invalid --user-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an action for the environment, for example, creates a console link for
// an analytics tool that is available in this environment.
func datazone_CreateEnvironmentAction(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateEnvironmentActionInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// Name: *string, // Required
		// Parameters: types.ActionParameters, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneParameters) > 0 {
		if err := assignInputField(input, "Parameters", _datazoneParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}

	if resp, err := client.CreateEnvironmentAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a Amazon DataZone blueprint.
func datazone_CreateEnvironmentBlueprint(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateEnvironmentBlueprintInput{
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
		// ProvisioningProperties: types.ProvisioningProperties, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneProvisioningProperties) > 0 {
		if err := assignInputField(input, "ProvisioningProperties", _datazoneProvisioningProperties); err != nil {
			log.Errorf("invalid --provisioning-properties: %s", err.Error())
			return
		}
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneUserParameters) > 0 {
		if err := assignInputField(input, "UserParameters", _datazoneUserParameters); err != nil {
			log.Errorf("invalid --user-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironmentBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon DataZone environment profile.
func datazone_CreateEnvironmentProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateEnvironmentProfileInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentBlueprintIdentifier: *string, // Required
		// Name: *string, // Required
		// ProjectIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentBlueprintIdentifier) > 0 {
		input.EnvironmentBlueprintIdentifier = aws.String(_datazoneEnvironmentBlueprintIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneProjectIdentifier) > 0 {
		input.ProjectIdentifier = aws.String(_datazoneProjectIdentifier)
	}
	if len(_datazoneAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_datazoneAwsAccountId)
	}
	if len(_datazoneAwsAccountRegion) > 0 {
		input.AwsAccountRegion = aws.String(_datazoneAwsAccountRegion)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneUserParameters) > 0 {
		if err := assignInputField(input, "UserParameters", _datazoneUserParameters); err != nil {
			log.Errorf("invalid --user-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEnvironmentProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a metadata form type.
// Prerequisites:
//
// - The domain must exist and be in an ENABLED state.
//
// - The owning project must exist and be accessible.
//
// - The name must be unique within the domain.
//
// For custom form types, to indicate that a field should be searchable, annotate
// it with (at)amazon.datazone#searchable . By default, searchable fields are indexed
// for semantic search, where related query terms will match the attribute value
// even if they are not stemmed or keyword matches. To indicate that a field should
// be indexed for lexical search (which disables semantic search but supports
// stemmed and partial matches), annotate it with
// (at)amazon.datazone#searchable(modes:["LEXICAL"]) . To indicate that a field should
// be indexed for technical identifier search (for more information on technical
// identifier search, see: [https://aws.amazon.com/blogs/big-data/streamline-data-discovery-with-precise-technical-identifier-search-in-amazon-sagemaker-unified-studio/]), annotate it with
// (at)amazon.datazone#searchable(modes:["TECHNICAL"]) .
//
// To denote that a field will store glossary term ids (which are filterable via
// the Search/SearchListings APIs), annotate it with
// (at)amazon.datazone#glossaryterm("${GLOSSARY_ID}") , where ${GLOSSARY_ID} is the
// id of the glossary that the glossary terms stored in the field belong to.
//
// [https://aws.amazon.com/blogs/big-data/streamline-data-discovery-with-precise-technical-identifier-search-in-amazon-sagemaker-unified-studio/]: https://aws.amazon.com/blogs/big-data/streamline-data-discovery-with-precise-technical-identifier-search-in-amazon-sagemaker-unified-studio/
func datazone_CreateFormType(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateFormTypeInput{
		// DomainIdentifier: *string, // Required
		// Model: types.Model, // Required
		// Name: *string, // Required
		// OwningProjectIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneModel) > 0 {
		if err := assignInputField(input, "Model", _datazoneModel); err != nil {
			log.Errorf("invalid --model: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneOwningProjectIdentifier) > 0 {
		input.OwningProjectIdentifier = aws.String(_datazoneOwningProjectIdentifier)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFormType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon DataZone business glossary.
// Specifies that this is a create glossary policy.
//
// A glossary serves as the central repository for business terminology and
// definitions within an organization. It helps establish and maintain a common
// language across different departments and teams, reducing miscommunication and
// ensuring consistent interpretation of business concepts. Glossaries can include
// hierarchical relationships between terms, cross-references, and links to actual
// data assets, making them invaluable for both business users and technical teams
// trying to understand and use data correctly.
//
// Prerequisites:
//
// - Domain must exist and be in an active state.
//
// - Owning project must exist and be accessible by the caller.
//
// - The glossary name must be unique within the domain.
func datazone_CreateGlossary(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateGlossaryInput{
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
		// OwningProjectIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneOwningProjectIdentifier) > 0 {
		input.OwningProjectIdentifier = aws.String(_datazoneOwningProjectIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_datazoneUsageRestrictions) > 0 {
		if err := assignInputField(input, "UsageRestrictions", _datazoneUsageRestrictions); err != nil {
			log.Errorf("invalid --usage-restrictions: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGlossary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a business glossary term.
// A glossary term represents an individual entry within the Amazon DataZone
// glossary, serving as a standardized definition for a specific business concept
// or data element. Each term can include rich metadata such as detailed
// definitions, synonyms, related terms, and usage examples. Glossary terms can be
// linked directly to data assets, providing business context to technical data
// elements. This linking capability helps users understand the business meaning of
// data fields and ensures consistent interpretation across different systems and
// teams. Terms can also have relationships with other terms, creating a semantic
// network that reflects the complexity of business concepts.
//
// Prerequisites:
//
// - Domain must exist.
//
// - Glossary must exist.
//
// - The term name must be unique within the glossary.
//
// - Ensure term does not conflict with existing terms in hierarchy.
func datazone_CreateGlossaryTerm(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateGlossaryTermInput{
		// DomainIdentifier: *string, // Required
		// GlossaryIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneGlossaryIdentifier) > 0 {
		input.GlossaryIdentifier = aws.String(_datazoneGlossaryIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneLongDescription) > 0 {
		input.LongDescription = aws.String(_datazoneLongDescription)
	}
	if len(_datazoneShortDescription) > 0 {
		input.ShortDescription = aws.String(_datazoneShortDescription)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_datazoneTermRelations) > 0 {
		if err := assignInputField(input, "TermRelations", _datazoneTermRelations); err != nil {
			log.Errorf("invalid --term-relations: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGlossaryTerm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a group profile in Amazon DataZone.
func datazone_CreateGroupProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateGroupProfileInput{
		// DomainIdentifier: *string, // Required
		// GroupIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_datazoneGroupIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}

	if resp, err := client.CreateGroupProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Publishes a listing (a record of an asset at a given time) or removes a listing
// from the catalog.
func datazone_CreateListingChangeSet(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateListingChangeSetInput{
		// Action: types.ChangeAction, // Required
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.EntityType, // Required
	}

	if len(_datazoneAction) > 0 {
		if err := assignInputField(input, "Action", _datazoneAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneEntityRevision) > 0 {
		input.EntityRevision = aws.String(_datazoneEntityRevision)
	}

	if resp, err := client.CreateListingChangeSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon DataZone project.
func datazone_CreateProject(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateProjectInput{
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneDomainUnitId) > 0 {
		input.DomainUnitId = aws.String(_datazoneDomainUnitId)
	}
	if len(_datazoneGlossaryTerms) > 0 {
		input.GlossaryTerms = append([]string(nil), _datazoneGlossaryTerms...)
	}
	if len(_datazoneProjectProfileId) > 0 {
		input.ProjectProfileId = aws.String(_datazoneProjectProfileId)
	}
	if len(_datazoneResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _datazoneResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_datazoneUserParameters) > 0 {
		if err := assignInputField(input, "UserParameters", _datazoneUserParameters); err != nil {
			log.Errorf("invalid --user-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a project membership in Amazon DataZone.
func datazone_CreateProjectMembership(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateProjectMembershipInput{
		// Designation: types.UserDesignation, // Required
		// DomainIdentifier: *string, // Required
		// Member: types.Member, // Required
		// ProjectIdentifier: *string, // Required
	}

	if len(_datazoneDesignation) > 0 {
		if err := assignInputField(input, "Designation", _datazoneDesignation); err != nil {
			log.Errorf("invalid --designation: %s", err.Error())
			return
		}
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneMember) > 0 {
		if err := assignInputField(input, "Member", _datazoneMember); err != nil {
			log.Errorf("invalid --member: %s", err.Error())
			return
		}
	}
	if len(_datazoneProjectIdentifier) > 0 {
		input.ProjectIdentifier = aws.String(_datazoneProjectIdentifier)
	}

	if resp, err := client.CreateProjectMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a project profile.
func datazone_CreateProjectProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateProjectProfileInput{
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneAllowCustomProjectResourceTags) > 0 {
		if err := assignInputField(input, "AllowCustomProjectResourceTags", _datazoneAllowCustomProjectResourceTags); err != nil {
			log.Errorf("invalid --allow-custom-project-resource-tags: %s", err.Error())
			return
		}
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneDomainUnitIdentifier) > 0 {
		input.DomainUnitIdentifier = aws.String(_datazoneDomainUnitIdentifier)
	}
	if len(_datazoneEnvironmentConfigurations) > 0 {
		if err := assignInputField(input, "EnvironmentConfigurations", _datazoneEnvironmentConfigurations); err != nil {
			log.Errorf("invalid --environment-configurations: %s", err.Error())
			return
		}
	}
	if len(_datazoneProjectResourceTags) > 0 {
		if err := assignInputField(input, "ProjectResourceTags", _datazoneProjectResourceTags); err != nil {
			log.Errorf("invalid --project-resource-tags: %s", err.Error())
			return
		}
	}
	if len(_datazoneProjectResourceTagsDescription) > 0 {
		input.ProjectResourceTagsDescription = aws.String(_datazoneProjectResourceTagsDescription)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProjectProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a rule in Amazon DataZone. A rule is a formal agreement that enforces
// specific requirements across user workflows (e.g., publishing assets to the
// catalog, requesting subscriptions, creating projects) within the Amazon DataZone
// data portal. These rules help maintain consistency, ensure compliance, and
// uphold governance standards in data management processes. For instance, a
// metadata enforcement rule can specify the required information for creating a
// subscription request or publishing a data asset to the catalog, ensuring
// alignment with organizational standards.
func datazone_CreateRule(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateRuleInput{
		// Action: types.RuleAction, // Required
		// Detail: types.RuleDetail, // Required
		// DomainIdentifier: *string, // Required
		// Name: *string, // Required
		// Scope: *types.RuleScope, // Required
		// Target: types.RuleTarget, // Required
	}

	if len(_datazoneAction) > 0 {
		if err := assignInputField(input, "Action", _datazoneAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_datazoneDetail) > 0 {
		if err := assignInputField(input, "Detail", _datazoneDetail); err != nil {
			log.Errorf("invalid --detail: %s", err.Error())
			return
		}
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneScope) > 0 {
		if err := assignInputField(input, "Scope", _datazoneScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_datazoneTarget) > 0 {
		if err := assignInputField(input, "Target", _datazoneTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}

	if resp, err := client.CreateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a subsscription grant in Amazon DataZone.
func datazone_CreateSubscriptionGrant(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateSubscriptionGrantInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// GrantedEntity: types.GrantedEntityInput, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneGrantedEntity) > 0 {
		if err := assignInputField(input, "GrantedEntity", _datazoneGrantedEntity); err != nil {
			log.Errorf("invalid --granted-entity: %s", err.Error())
			return
		}
	}
	if len(_datazoneAssetTargetNames) > 0 {
		if err := assignInputField(input, "AssetTargetNames", _datazoneAssetTargetNames); err != nil {
			log.Errorf("invalid --asset-target-names: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneSubscriptionTargetIdentifier) > 0 {
		input.SubscriptionTargetIdentifier = aws.String(_datazoneSubscriptionTargetIdentifier)
	}

	if resp, err := client.CreateSubscriptionGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a subscription request in Amazon DataZone.
func datazone_CreateSubscriptionRequest(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateSubscriptionRequestInput{
		// DomainIdentifier: *string, // Required
		// RequestReason: *string, // Required
		// SubscribedListings: []types.SubscribedListingInput, // Required
		// SubscribedPrincipals: []types.SubscribedPrincipalInput, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneRequestReason) > 0 {
		input.RequestReason = aws.String(_datazoneRequestReason)
	}
	if len(_datazoneSubscribedListings) > 0 {
		if err := assignInputField(input, "SubscribedListings", _datazoneSubscribedListings); err != nil {
			log.Errorf("invalid --subscribed-listings: %s", err.Error())
			return
		}
	}
	if len(_datazoneSubscribedPrincipals) > 0 {
		if err := assignInputField(input, "SubscribedPrincipals", _datazoneSubscribedPrincipals); err != nil {
			log.Errorf("invalid --subscribed-principals: %s", err.Error())
			return
		}
	}
	if len(_datazoneAssetPermissions) > 0 {
		if err := assignInputField(input, "AssetPermissions", _datazoneAssetPermissions); err != nil {
			log.Errorf("invalid --asset-permissions: %s", err.Error())
			return
		}
	}
	if len(_datazoneAssetScopes) > 0 {
		if err := assignInputField(input, "AssetScopes", _datazoneAssetScopes); err != nil {
			log.Errorf("invalid --asset-scopes: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneMetadataForms) > 0 {
		if err := assignInputField(input, "MetadataForms", _datazoneMetadataForms); err != nil {
			log.Errorf("invalid --metadata-forms: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSubscriptionRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a subscription target in Amazon DataZone.
func datazone_CreateSubscriptionTarget(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateSubscriptionTargetInput{
		// ApplicableAssetTypes: []string, // Required
		// AuthorizedPrincipals: []string, // Required
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// ManageAccessRole: *string, // Required
		// Name: *string, // Required
		// SubscriptionTargetConfig: []types.SubscriptionTargetForm, // Required
		// Type: *string, // Required
	}

	if len(_datazoneApplicableAssetTypes) > 0 {
		input.ApplicableAssetTypes = append([]string(nil), _datazoneApplicableAssetTypes...)
	}
	if len(_datazoneAuthorizedPrincipals) > 0 {
		input.AuthorizedPrincipals = append([]string(nil), _datazoneAuthorizedPrincipals...)
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneManageAccessRole) > 0 {
		input.ManageAccessRole = aws.String(_datazoneManageAccessRole)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneSubscriptionTargetConfig) > 0 {
		if err := assignInputField(input, "SubscriptionTargetConfig", _datazoneSubscriptionTargetConfig); err != nil {
			log.Errorf("invalid --subscription-target-config: %s", err.Error())
			return
		}
	}
	if len(_datazoneType) > 0 {
		input.Type = aws.String(_datazoneType)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneProvider) > 0 {
		input.Provider = aws.String(_datazoneProvider)
	}
	if len(_datazoneSubscriptionGrantCreationMode) > 0 {
		if err := assignInputField(input, "SubscriptionGrantCreationMode", _datazoneSubscriptionGrantCreationMode); err != nil {
			log.Errorf("invalid --subscription-grant-creation-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSubscriptionTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user profile in Amazon DataZone.
func datazone_CreateUserProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.CreateUserProfileInput{
		// DomainIdentifier: *string, // Required
		// UserIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneUserIdentifier) > 0 {
		input.UserIdentifier = aws.String(_datazoneUserIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneUserType) > 0 {
		if err := assignInputField(input, "UserType", _datazoneUserType); err != nil {
			log.Errorf("invalid --user-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUserProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an account pool.
func datazone_DeleteAccountPool(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteAccountPoolInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteAccountPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an asset in Amazon DataZone.
// - --domain-identifier must refer to a valid and existing domain.
//
// - --identifier must refer to an existing asset in the specified domain.
//
// - Asset must not be referenced in any existing asset filters.
//
// - Asset must not be linked to any draft or published data product.
//
// - User must have delete permissions for the domain and project.
func datazone_DeleteAsset(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteAssetInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an asset filter.
// Prerequisites:
//
// - The asset filter must exist.
//
// - The domain and asset must not have been deleted.
//
// - Ensure the --identifier refers to a valid filter ID.
func datazone_DeleteAssetFilter(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteAssetFilterInput{
		// AssetIdentifier: *string, // Required
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneAssetIdentifier) > 0 {
		input.AssetIdentifier = aws.String(_datazoneAssetIdentifier)
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteAssetFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an asset type in Amazon DataZone.
// Prerequisites:
//
// - The asset type must exist in the domain.
//
// - You must have DeleteAssetType permission.
//
// - The asset type must not be in use (e.g., assigned to any asset). If used,
// deletion will fail.
//
// - You should retrieve the asset type using get-asset-type to confirm its
// presence before deletion.
func datazone_DeleteAssetType(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteAssetTypeInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteAssetType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes and connection. In Amazon DataZone, a connection enables you to connect
// your resources (domains, projects, and environments) to external resources and
// services.
func datazone_DeleteConnection(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteConnectionInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes data export configuration for a domain.
// This operation does not delete the S3 table created by the
// PutDataExportConfiguration operation.
//
// To temporarily disable export without deleting the configuration, use the
// PutDataExportConfiguration operation with the --no-enable-export flag instead.
// This allows you to re-enable export for the same domain using the
// --enable-export flag without deleting S3 table.
func datazone_DeleteDataExportConfiguration(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteDataExportConfigurationInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}

	if resp, err := client.DeleteDataExportConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a data product in Amazon DataZone.
// Prerequisites:
//
// - The data product must exist and not be deleted or archived.
//
// - The user must have delete permissions for the data product.
//
// - Domain and project must be active.
func datazone_DeleteDataProduct(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteDataProductInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteDataProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a data source in Amazon DataZone.
func datazone_DeleteDataSource(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteDataSourceInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneRetainPermissionsOnRevokeFailure) > 0 {
		if err := assignInputField(input, "RetainPermissionsOnRevokeFailure", _datazoneRetainPermissionsOnRevokeFailure); err != nil {
			log.Errorf("invalid --retain-permissions-on-revoke-failure: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Amazon DataZone domain.
func datazone_DeleteDomain(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteDomainInput{
		// Identifier: *string, // Required
	}

	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneSkipDeletionCheck) > 0 {
		if err := assignInputField(input, "SkipDeletionCheck", _datazoneSkipDeletionCheck); err != nil {
			log.Errorf("invalid --skip-deletion-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a domain unit.
func datazone_DeleteDomainUnit(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteDomainUnitInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteDomainUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an environment in Amazon DataZone.
func datazone_DeleteEnvironment(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteEnvironmentInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an action for the environment, for example, deletes a console link for
// an analytics tool that is available in this environment.
func datazone_DeleteEnvironmentAction(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteEnvironmentActionInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteEnvironmentAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a blueprint in Amazon DataZone.
func datazone_DeleteEnvironmentBlueprint(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteEnvironmentBlueprintInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteEnvironmentBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the blueprint configuration in Amazon DataZone.
func datazone_DeleteEnvironmentBlueprintConfiguration(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteEnvironmentBlueprintConfigurationInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentBlueprintIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentBlueprintIdentifier) > 0 {
		input.EnvironmentBlueprintIdentifier = aws.String(_datazoneEnvironmentBlueprintIdentifier)
	}

	if resp, err := client.DeleteEnvironmentBlueprintConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an environment profile in Amazon DataZone.
func datazone_DeleteEnvironmentProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteEnvironmentProfileInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteEnvironmentProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes and metadata form type in Amazon DataZone.
// Prerequisites:
//
// - The form type must exist in the domain.
//
// - The form type must not be in use by any asset types or assets.
//
// - The domain must be valid and accessible.
//
// - User must have delete permissions on the form type.
//
// - Any dependencies (such as linked asset types) must be removed first.
func datazone_DeleteFormType(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteFormTypeInput{
		// DomainIdentifier: *string, // Required
		// FormTypeIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneFormTypeIdentifier) > 0 {
		input.FormTypeIdentifier = aws.String(_datazoneFormTypeIdentifier)
	}

	if resp, err := client.DeleteFormType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a business glossary in Amazon DataZone.
// Prerequisites:
//
// - The glossary must be in DISABLED state.
//
// - The glossary must not have any glossary terms associated with it.
//
// - The glossary must exist in the specified domain.
//
// - The caller must have the datazone:DeleteGlossary permission in the domain
// and glossary.
//
// - Glossary should not be linked to any active metadata forms.
func datazone_DeleteGlossary(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteGlossaryInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteGlossary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a business glossary term in Amazon DataZone.
// Prerequisites:
//
// - Glossary term must exist and be active.
//
// - The term must not be linked to other assets or child terms.
//
// - Caller must have delete permissions in the domain/glossary.
//
// - Ensure all associations (such as to assets or parent terms) are removed
// before deletion.
func datazone_DeleteGlossaryTerm(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteGlossaryTermInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteGlossaryTerm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a listing (a record of an asset at a given time).
func datazone_DeleteListing(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteListingInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteListing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a project in Amazon DataZone.
func datazone_DeleteProject(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteProjectInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneSkipDeletionCheck) > 0 {
		if err := assignInputField(input, "SkipDeletionCheck", _datazoneSkipDeletionCheck); err != nil {
			log.Errorf("invalid --skip-deletion-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes project membership in Amazon DataZone.
func datazone_DeleteProjectMembership(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteProjectMembershipInput{
		// DomainIdentifier: *string, // Required
		// Member: types.Member, // Required
		// ProjectIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneMember) > 0 {
		if err := assignInputField(input, "Member", _datazoneMember); err != nil {
			log.Errorf("invalid --member: %s", err.Error())
			return
		}
	}
	if len(_datazoneProjectIdentifier) > 0 {
		input.ProjectIdentifier = aws.String(_datazoneProjectIdentifier)
	}

	if resp, err := client.DeleteProjectMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a project profile.
func datazone_DeleteProjectProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteProjectProfileInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteProjectProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a rule in Amazon DataZone. A rule is a formal agreement that enforces
// specific requirements across user workflows (e.g., publishing assets to the
// catalog, requesting subscriptions, creating projects) within the Amazon DataZone
// data portal. These rules help maintain consistency, ensure compliance, and
// uphold governance standards in data management processes. For instance, a
// metadata enforcement rule can specify the required information for creating a
// subscription request or publishing a data asset to the catalog, ensuring
// alignment with organizational standards.
func datazone_DeleteRule(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteRuleInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes and subscription grant in Amazon DataZone.
func datazone_DeleteSubscriptionGrant(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteSubscriptionGrantInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteSubscriptionGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a subscription request in Amazon DataZone.
func datazone_DeleteSubscriptionRequest(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteSubscriptionRequestInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteSubscriptionRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a subscription target in Amazon DataZone.
func datazone_DeleteSubscriptionTarget(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteSubscriptionTargetInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.DeleteSubscriptionTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified time series form for the specified asset.
func datazone_DeleteTimeSeriesDataPoints(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DeleteTimeSeriesDataPointsInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.TimeSeriesEntityType, // Required
		// FormName: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneFormName) > 0 {
		input.FormName = aws.String(_datazoneFormName)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}

	if resp, err := client.DeleteTimeSeriesDataPoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the environment role in Amazon DataZone.
func datazone_DisassociateEnvironmentRole(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DisassociateEnvironmentRoleInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// EnvironmentRoleArn: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneEnvironmentRoleArn) > 0 {
		input.EnvironmentRoleArn = aws.String(_datazoneEnvironmentRoleArn)
	}

	if resp, err := client.DisassociateEnvironmentRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates restricted terms from an asset.
func datazone_DisassociateGovernedTerms(cfg aws.Config, client *datazone.Client) {
	input := &datazone.DisassociateGovernedTermsInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.GovernedEntityType, // Required
		// GovernedGlossaryTerms: []string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneGovernedGlossaryTerms) > 0 {
		input.GovernedGlossaryTerms = append([]string(nil), _datazoneGovernedGlossaryTerms...)
	}

	if resp, err := client.DisassociateGovernedTerms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of the account pool.
func datazone_GetAccountPool(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetAccountPoolInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetAccountPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon DataZone asset.
// An asset is the fundamental building block in Amazon DataZone, representing any
// data resource that needs to be cataloged and managed. It can take many forms,
// from Amazon S3 buckets and database tables to dashboards and machine learning
// models. Each asset contains comprehensive metadata about the resource, including
// its location, schema, ownership, and lineage information. Assets are essential
// for organizing and managing data resources across an organization, making them
// discoverable and usable while maintaining proper governance.
//
// Before using the Amazon DataZone GetAsset command, ensure the following
// prerequisites are met:
//
// - Domain identifier must exist and be valid
//
// - Asset identifier must exist
//
// - User must have the required permissions to perform the action
func datazone_GetAsset(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetAssetInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneRevision) > 0 {
		input.Revision = aws.String(_datazoneRevision)
	}

	if resp, err := client.GetAsset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an asset filter.
// Prerequisites:
//
// - Domain ( --domain-identifier ), asset ( --asset-identifier ), and filter (
// --identifier ) must all exist.
//
// - The asset filter should not have been deleted.
//
// - The asset must still exist (since the filter is linked to it).
func datazone_GetAssetFilter(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetAssetFilterInput{
		// AssetIdentifier: *string, // Required
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneAssetIdentifier) > 0 {
		input.AssetIdentifier = aws.String(_datazoneAssetIdentifier)
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetAssetFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon DataZone asset type.
// Asset types define the categories and characteristics of different kinds of
// data assets within Amazon DataZone.. They determine what metadata fields are
// required, what operations are possible, and how the asset integrates with other
// Amazon Web Services services. Asset types can range from built-in types like
// Amazon S3 buckets and Amazon Web Services Glue tables to custom types defined
// for specific organizational needs. Understanding asset types is crucial for
// properly organizing and managing different kinds of data resources.
//
// Prerequisites:
//
// - The asset type with identifier must exist in the domain.
// ResourceNotFoundException.
//
// - You must have the GetAssetType permission.
//
// - Ensure the domain-identifier value is correct and accessible.
func datazone_GetAssetType(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetAssetTypeInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneRevision) > 0 {
		input.Revision = aws.String(_datazoneRevision)
	}

	if resp, err := client.GetAssetType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a connection. In Amazon DataZone, a connection enables you to connect your
// resources (domains, projects, and environments) to external resources and
// services.
func datazone_GetConnection(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetConnectionInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneWithSecret) > 0 {
		if err := assignInputField(input, "WithSecret", _datazoneWithSecret); err != nil {
			log.Errorf("invalid --with-secret: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets data export configuration details.
func datazone_GetDataExportConfiguration(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetDataExportConfigurationInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}

	if resp, err := client.GetDataExportConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the data product.
// Prerequisites:
//
// - The data product ID must exist.
//
// - The domain must be valid and accessible.
//
// - User must have read or discovery permissions for the data product.
func datazone_GetDataProduct(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetDataProductInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneRevision) > 0 {
		input.Revision = aws.String(_datazoneRevision)
	}

	if resp, err := client.GetDataProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon DataZone data source.
func datazone_GetDataSource(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetDataSourceInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon DataZone data source run.
func datazone_GetDataSourceRun(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetDataSourceRunInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetDataSourceRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon DataZone domain.
func datazone_GetDomain(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetDomainInput{
		// Identifier: *string, // Required
	}

	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of the specified domain unit.
func datazone_GetDomainUnit(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetDomainUnitInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetDomainUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon DataZone environment.
func datazone_GetEnvironment(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetEnvironmentInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the specified environment action.
func datazone_GetEnvironmentAction(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetEnvironmentActionInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetEnvironmentAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon DataZone blueprint.
func datazone_GetEnvironmentBlueprint(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetEnvironmentBlueprintInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetEnvironmentBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the blueprint configuration in Amazon DataZone.
func datazone_GetEnvironmentBlueprintConfiguration(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetEnvironmentBlueprintConfigurationInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentBlueprintIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentBlueprintIdentifier) > 0 {
		input.EnvironmentBlueprintIdentifier = aws.String(_datazoneEnvironmentBlueprintIdentifier)
	}

	if resp, err := client.GetEnvironmentBlueprintConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the credentials of an environment in Amazon DataZone.
func datazone_GetEnvironmentCredentials(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetEnvironmentCredentialsInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}

	if resp, err := client.GetEnvironmentCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an evinronment profile in Amazon DataZone.
func datazone_GetEnvironmentProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetEnvironmentProfileInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetEnvironmentProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a metadata form type in Amazon DataZone.
// Form types define the structure and validation rules for collecting metadata
// about assets in Amazon DataZone. They act as templates that ensure consistent
// metadata capture across similar types of assets, while allowing for
// customization to meet specific organizational needs. Form types can include
// required fields, validation rules, and dependencies, helping maintain
// high-quality metadata that makes data assets more discoverable and usable.
//
// - The form type with the specified identifier must exist in the given domain.
//
// - The domain must be valid and active.
//
// - User must have permission on the form type.
//
// - The form type should not be deleted or in an invalid state.
//
// One use case for this API is to determine whether a form field is indexed for
// search.
//
// A searchable field will be annotated with (at)amazon.datazone#searchable . By
// default, searchable fields are indexed for semantic search, where related query
// terms will match the attribute value even if they are not stemmed or keyword
// matches. If a field is indexed technical identifier search, it will be annotated
// with (at)amazon.datazone#searchable(modes:["TECHNICAL"]) . If a field is indexed
// for lexical search (supports stemmed and prefix matches but not semantic
// matches), it will be annotated with
// (at)amazon.datazone#searchable(modes:["LEXICAL"]) .
//
// A field storing glossary term IDs (which is filterable) will be annotated with
// (at)amazon.datazone#glossaryterm("${glossaryId}") .
func datazone_GetFormType(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetFormTypeInput{
		// DomainIdentifier: *string, // Required
		// FormTypeIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneFormTypeIdentifier) > 0 {
		input.FormTypeIdentifier = aws.String(_datazoneFormTypeIdentifier)
	}
	if len(_datazoneRevision) > 0 {
		input.Revision = aws.String(_datazoneRevision)
	}

	if resp, err := client.GetFormType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a business glossary in Amazon DataZone.
// Prerequisites:
//
// - The specified glossary ID must exist and be associated with the given
// domain.
//
// - The caller must have the datazone:GetGlossary permission on the domain.
func datazone_GetGlossary(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetGlossaryInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetGlossary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a business glossary term in Amazon DataZone.
// Prerequisites:
//
// - Glossary term with identifier must exist in the domain.
//
// - User must have permission on the glossary term.
//
// - Domain must be accessible and active.
func datazone_GetGlossaryTerm(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetGlossaryTermInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetGlossaryTerm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a group profile in Amazon DataZone.
func datazone_GetGroupProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetGroupProfileInput{
		// DomainIdentifier: *string, // Required
		// GroupIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_datazoneGroupIdentifier)
	}

	if resp, err := client.GetGroupProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the data portal URL for the specified Amazon DataZone domain.
func datazone_GetIamPortalLoginUrl(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetIamPortalLoginUrlInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}

	if resp, err := client.GetIamPortalLoginUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The details of the job run.
func datazone_GetJobRun(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetJobRunInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetJobRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the lineage event.
func datazone_GetLineageEvent(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetLineageEventInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetLineageEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the data lineage node.
func datazone_GetLineageNode(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetLineageNodeInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneEventTimestamp) > 0 {
		if err := assignInputField(input, "EventTimestamp", _datazoneEventTimestamp); err != nil {
			log.Errorf("invalid --event-timestamp: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLineageNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a listing (a record of an asset at a given time). If you specify a listing
// version, only details that are specific to that version are returned.
func datazone_GetListing(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetListingInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneListingRevision) > 0 {
		input.ListingRevision = aws.String(_datazoneListingRevision)
	}

	if resp, err := client.GetListing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a metadata generation run in Amazon DataZone.
// Prerequisites:
//
// - Valid domain and run identifier.
//
// - The metadata generation run must exist.
//
// - User must have read access to the metadata run.
func datazone_GetMetadataGenerationRun(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetMetadataGenerationRunInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneType) > 0 {
		if err := assignInputField(input, "Type", _datazoneType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMetadataGenerationRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a project in Amazon DataZone.
func datazone_GetProject(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetProjectInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The details of the project profile.
func datazone_GetProjectProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetProjectProfileInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetProjectProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a rule in Amazon DataZone. A rule is a formal agreement
// that enforces specific requirements across user workflows (e.g., publishing
// assets to the catalog, requesting subscriptions, creating projects) within the
// Amazon DataZone data portal. These rules help maintain consistency, ensure
// compliance, and uphold governance standards in data management processes. For
// instance, a metadata enforcement rule can specify the required information for
// creating a subscription request or publishing a data asset to the catalog,
// ensuring alignment with organizational standards.
func datazone_GetRule(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetRuleInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneRevision) > 0 {
		input.Revision = aws.String(_datazoneRevision)
	}

	if resp, err := client.GetRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a subscription in Amazon DataZone.
func datazone_GetSubscription(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetSubscriptionInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the subscription grant in Amazon DataZone.
func datazone_GetSubscriptionGrant(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetSubscriptionGrantInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetSubscriptionGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of the specified subscription request.
func datazone_GetSubscriptionRequestDetails(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetSubscriptionRequestDetailsInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetSubscriptionRequestDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the subscription target in Amazon DataZone.
func datazone_GetSubscriptionTarget(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetSubscriptionTargetInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetSubscriptionTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the existing data point for the asset.
func datazone_GetTimeSeriesDataPoint(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetTimeSeriesDataPointInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.TimeSeriesEntityType, // Required
		// FormName: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneFormName) > 0 {
		input.FormName = aws.String(_datazoneFormName)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}

	if resp, err := client.GetTimeSeriesDataPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a user profile in Amazon DataZone.
func datazone_GetUserProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.GetUserProfileInput{
		// DomainIdentifier: *string, // Required
		// UserIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneUserIdentifier) > 0 {
		input.UserIdentifier = aws.String(_datazoneUserIdentifier)
	}
	if len(_datazoneType) > 0 {
		if err := assignInputField(input, "Type", _datazoneType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetUserProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists existing account pools.
func datazone_ListAccountPools(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListAccountPoolsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _datazoneSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAccountPools(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListAccountPoolsOutput
	p := datazone.NewListAccountPoolsPaginator(client, input)
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

// Lists the accounts in the specified account pool.
func datazone_ListAccountsInAccountPool(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListAccountsInAccountPoolInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccountsInAccountPool(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListAccountsInAccountPoolOutput
	p := datazone.NewListAccountsInAccountPoolPaginator(client, input)
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

// Lists asset filters.
// Prerequisites:
//
// - A valid domain and asset must exist.
//
// - The asset must have at least one filter created to return results.
func datazone_ListAssetFilters(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListAssetFiltersInput{
		// AssetIdentifier: *string, // Required
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneAssetIdentifier) > 0 {
		input.AssetIdentifier = aws.String(_datazoneAssetIdentifier)
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAssetFilters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListAssetFiltersOutput
	p := datazone.NewListAssetFiltersPaginator(client, input)
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

// Lists the revisions for the asset.
// Prerequisites:
//
// - The asset must exist in the domain.
//
// - There must be at least one revision of the asset (which happens
// automatically after creation).
//
// - The domain must be valid and active.
//
// - User must have permissions on the asset and domain.
func datazone_ListAssetRevisions(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListAssetRevisionsInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssetRevisions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListAssetRevisionsOutput
	p := datazone.NewListAssetRevisionsPaginator(client, input)
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

// Lists connections. In Amazon DataZone, a connection enables you to connect your
// resources (domains, projects, and environments) to external resources and
// services.
func datazone_ListConnections(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListConnectionsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneProjectIdentifier) > 0 {
		input.ProjectIdentifier = aws.String(_datazoneProjectIdentifier)
	}
	if len(_datazoneScope) > 0 {
		if err := assignInputField(input, "Scope", _datazoneScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_datazoneSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _datazoneSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_datazoneType) > 0 {
		if err := assignInputField(input, "Type", _datazoneType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListConnectionsOutput
	p := datazone.NewListConnectionsPaginator(client, input)
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

// Lists data product revisions.
// Prerequisites:
//
// - The data product ID must exist within the domain.
//
// - User must have view permissions on the data product.
//
// - The domain must be in a valid and accessible state.
func datazone_ListDataProductRevisions(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListDataProductRevisionsInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDataProductRevisions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListDataProductRevisionsOutput
	p := datazone.NewListDataProductRevisionsPaginator(client, input)
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

// Lists data source run activities.
func datazone_ListDataSourceRunActivities(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListDataSourceRunActivitiesInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDataSourceRunActivities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListDataSourceRunActivitiesOutput
	p := datazone.NewListDataSourceRunActivitiesPaginator(client, input)
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

// Lists data source runs in Amazon DataZone.
func datazone_ListDataSourceRuns(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListDataSourceRunsInput{
		// DataSourceIdentifier: *string, // Required
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDataSourceIdentifier) > 0 {
		input.DataSourceIdentifier = aws.String(_datazoneDataSourceIdentifier)
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDataSourceRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListDataSourceRunsOutput
	p := datazone.NewListDataSourceRunsPaginator(client, input)
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

// Lists data sources in Amazon DataZone.
func datazone_ListDataSources(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListDataSourcesInput{
		// DomainIdentifier: *string, // Required
		// ProjectIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneProjectIdentifier) > 0 {
		input.ProjectIdentifier = aws.String(_datazoneProjectIdentifier)
	}
	if len(_datazoneConnectionIdentifier) > 0 {
		input.ConnectionIdentifier = aws.String(_datazoneConnectionIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_datazoneType) > 0 {
		input.Type = aws.String(_datazoneType)
	}

	if disablePaginator() {
		if resp, err := client.ListDataSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListDataSourcesOutput
	p := datazone.NewListDataSourcesPaginator(client, input)
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

// Lists child domain units for the specified parent domain unit.
func datazone_ListDomainUnitsForParent(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListDomainUnitsForParentInput{
		// DomainIdentifier: *string, // Required
		// ParentDomainUnitIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneParentDomainUnitIdentifier) > 0 {
		input.ParentDomainUnitIdentifier = aws.String(_datazoneParentDomainUnitIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomainUnitsForParent(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListDomainUnitsForParentOutput
	p := datazone.NewListDomainUnitsForParentPaginator(client, input)
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

// Lists Amazon DataZone domains.
func datazone_ListDomains(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListDomainsInput{}

	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDomains(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListDomainsOutput
	p := datazone.NewListDomainsPaginator(client, input)
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

// Lists the entity (domain units) owners.
func datazone_ListEntityOwners(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListEntityOwnersInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.DataZoneEntityType, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEntityOwners(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListEntityOwnersOutput
	p := datazone.NewListEntityOwnersPaginator(client, input)
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

// Lists existing environment actions.
func datazone_ListEnvironmentActions(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListEnvironmentActionsInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListEnvironmentActionsOutput
	p := datazone.NewListEnvironmentActionsPaginator(client, input)
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

// Lists blueprint configurations for a Amazon DataZone environment.
func datazone_ListEnvironmentBlueprintConfigurations(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListEnvironmentBlueprintConfigurationsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentBlueprintConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListEnvironmentBlueprintConfigurationsOutput
	p := datazone.NewListEnvironmentBlueprintConfigurationsPaginator(client, input)
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

// Lists blueprints in an Amazon DataZone environment.
func datazone_ListEnvironmentBlueprints(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListEnvironmentBlueprintsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneManaged) > 0 {
		if err := assignInputField(input, "Managed", _datazoneManaged); err != nil {
			log.Errorf("invalid --managed: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentBlueprints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListEnvironmentBlueprintsOutput
	p := datazone.NewListEnvironmentBlueprintsPaginator(client, input)
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

// Lists Amazon DataZone environment profiles.
func datazone_ListEnvironmentProfiles(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListEnvironmentProfilesInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_datazoneAwsAccountId)
	}
	if len(_datazoneAwsAccountRegion) > 0 {
		input.AwsAccountRegion = aws.String(_datazoneAwsAccountRegion)
	}
	if len(_datazoneEnvironmentBlueprintIdentifier) > 0 {
		input.EnvironmentBlueprintIdentifier = aws.String(_datazoneEnvironmentBlueprintIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneProjectIdentifier) > 0 {
		input.ProjectIdentifier = aws.String(_datazoneProjectIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironmentProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListEnvironmentProfilesOutput
	p := datazone.NewListEnvironmentProfilesPaginator(client, input)
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

// Lists Amazon DataZone environments.
func datazone_ListEnvironments(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListEnvironmentsInput{
		// DomainIdentifier: *string, // Required
		// ProjectIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneProjectIdentifier) > 0 {
		input.ProjectIdentifier = aws.String(_datazoneProjectIdentifier)
	}
	if len(_datazoneAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_datazoneAwsAccountId)
	}
	if len(_datazoneAwsAccountRegion) > 0 {
		input.AwsAccountRegion = aws.String(_datazoneAwsAccountRegion)
	}
	if len(_datazoneEnvironmentBlueprintIdentifier) > 0 {
		input.EnvironmentBlueprintIdentifier = aws.String(_datazoneEnvironmentBlueprintIdentifier)
	}
	if len(_datazoneEnvironmentProfileIdentifier) > 0 {
		input.EnvironmentProfileIdentifier = aws.String(_datazoneEnvironmentProfileIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneProvider) > 0 {
		input.Provider = aws.String(_datazoneProvider)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListEnvironmentsOutput
	p := datazone.NewListEnvironmentsPaginator(client, input)
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

// Lists job runs.
func datazone_ListJobRuns(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListJobRunsInput{
		// DomainIdentifier: *string, // Required
		// JobIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneJobIdentifier) > 0 {
		input.JobIdentifier = aws.String(_datazoneJobIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListJobRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListJobRunsOutput
	p := datazone.NewListJobRunsPaginator(client, input)
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

// Lists lineage events.
func datazone_ListLineageEvents(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListLineageEventsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneProcessingStatus) > 0 {
		if err := assignInputField(input, "ProcessingStatus", _datazoneProcessingStatus); err != nil {
			log.Errorf("invalid --processing-status: %s", err.Error())
			return
		}
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_datazoneTimestampAfter) > 0 {
		if err := assignInputField(input, "TimestampAfter", _datazoneTimestampAfter); err != nil {
			log.Errorf("invalid --timestamp-after: %s", err.Error())
			return
		}
	}
	if len(_datazoneTimestampBefore) > 0 {
		if err := assignInputField(input, "TimestampBefore", _datazoneTimestampBefore); err != nil {
			log.Errorf("invalid --timestamp-before: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLineageEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListLineageEventsOutput
	p := datazone.NewListLineageEventsPaginator(client, input)
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

// Lists the history of the specified data lineage node.
func datazone_ListLineageNodeHistory(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListLineageNodeHistoryInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneDirection) > 0 {
		if err := assignInputField(input, "Direction", _datazoneDirection); err != nil {
			log.Errorf("invalid --direction: %s", err.Error())
			return
		}
	}
	if len(_datazoneEventTimestampGTE) > 0 {
		if err := assignInputField(input, "EventTimestampGTE", _datazoneEventTimestampGTE); err != nil {
			log.Errorf("invalid --event-timestamp-gte: %s", err.Error())
			return
		}
	}
	if len(_datazoneEventTimestampLTE) > 0 {
		if err := assignInputField(input, "EventTimestampLTE", _datazoneEventTimestampLTE); err != nil {
			log.Errorf("invalid --event-timestamp-lte: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLineageNodeHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListLineageNodeHistoryOutput
	p := datazone.NewListLineageNodeHistoryPaginator(client, input)
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

// Lists all metadata generation runs.
// Metadata generation runs represent automated processes that leverage AI/ML
// capabilities to create or enhance asset metadata at scale. This feature helps
// organizations maintain comprehensive and consistent metadata across large
// numbers of assets without manual intervention. It can automatically generate
// business descriptions, tags, and other metadata elements, significantly reducing
// the time and effort required for metadata management while improving consistency
// and completeness.
//
// Prerequisites:
//
// - Valid domain identifier.
//
// - User must have access to metadata generation runs in the domain.
func datazone_ListMetadataGenerationRuns(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListMetadataGenerationRunsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_datazoneTargetIdentifier) > 0 {
		input.TargetIdentifier = aws.String(_datazoneTargetIdentifier)
	}
	if len(_datazoneType) > 0 {
		if err := assignInputField(input, "Type", _datazoneType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMetadataGenerationRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListMetadataGenerationRunsOutput
	p := datazone.NewListMetadataGenerationRunsPaginator(client, input)
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

// Lists all Amazon DataZone notifications.
func datazone_ListNotifications(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListNotificationsInput{
		// DomainIdentifier: *string, // Required
		// Type: types.NotificationType, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneType) > 0 {
		if err := assignInputField(input, "Type", _datazoneType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_datazoneAfterTimestamp) > 0 {
		if err := assignInputField(input, "AfterTimestamp", _datazoneAfterTimestamp); err != nil {
			log.Errorf("invalid --after-timestamp: %s", err.Error())
			return
		}
	}
	if len(_datazoneBeforeTimestamp) > 0 {
		if err := assignInputField(input, "BeforeTimestamp", _datazoneBeforeTimestamp); err != nil {
			log.Errorf("invalid --before-timestamp: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSubjects) > 0 {
		input.Subjects = append([]string(nil), _datazoneSubjects...)
	}
	if len(_datazoneTaskStatus) > 0 {
		if err := assignInputField(input, "TaskStatus", _datazoneTaskStatus); err != nil {
			log.Errorf("invalid --task-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListNotifications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListNotificationsOutput
	p := datazone.NewListNotificationsPaginator(client, input)
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

// Lists policy grants.
func datazone_ListPolicyGrants(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListPolicyGrantsInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.TargetEntityType, // Required
		// PolicyType: types.ManagedPolicyType, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazonePolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _datazonePolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyGrants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListPolicyGrantsOutput
	p := datazone.NewListPolicyGrantsPaginator(client, input)
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

// Lists all members of the specified project.
func datazone_ListProjectMemberships(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListProjectMembershipsInput{
		// DomainIdentifier: *string, // Required
		// ProjectIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneProjectIdentifier) > 0 {
		input.ProjectIdentifier = aws.String(_datazoneProjectIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _datazoneSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProjectMemberships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListProjectMembershipsOutput
	p := datazone.NewListProjectMembershipsPaginator(client, input)
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

// Lists project profiles.
func datazone_ListProjectProfiles(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListProjectProfilesInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _datazoneSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProjectProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListProjectProfilesOutput
	p := datazone.NewListProjectProfilesPaginator(client, input)
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

// Lists Amazon DataZone projects.
func datazone_ListProjects(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListProjectsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_datazoneGroupIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneUserIdentifier) > 0 {
		input.UserIdentifier = aws.String(_datazoneUserIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListProjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListProjectsOutput
	p := datazone.NewListProjectsPaginator(client, input)
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

// Lists existing rules. In Amazon DataZone, a rule is a formal agreement that
// enforces specific requirements across user workflows (e.g., publishing assets to
// the catalog, requesting subscriptions, creating projects) within the Amazon
// DataZone data portal. These rules help maintain consistency, ensure compliance,
// and uphold governance standards in data management processes. For instance, a
// metadata enforcement rule can specify the required information for creating a
// subscription request or publishing a data asset to the catalog, ensuring
// alignment with organizational standards.
func datazone_ListRules(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListRulesInput{
		// DomainIdentifier: *string, // Required
		// TargetIdentifier: *string, // Required
		// TargetType: types.RuleTargetType, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneTargetIdentifier) > 0 {
		input.TargetIdentifier = aws.String(_datazoneTargetIdentifier)
	}
	if len(_datazoneTargetType) > 0 {
		if err := assignInputField(input, "TargetType", _datazoneTargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneAction) > 0 {
		if err := assignInputField(input, "Action", _datazoneAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_datazoneAssetTypes) > 0 {
		input.AssetTypes = append([]string(nil), _datazoneAssetTypes...)
	}
	if len(_datazoneDataProduct) > 0 {
		if err := assignInputField(input, "DataProduct", _datazoneDataProduct); err != nil {
			log.Errorf("invalid --data-product: %s", err.Error())
			return
		}
	}
	if len(_datazoneIncludeCascaded) > 0 {
		if err := assignInputField(input, "IncludeCascaded", _datazoneIncludeCascaded); err != nil {
			log.Errorf("invalid --include-cascaded: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneProjectIds) > 0 {
		input.ProjectIds = append([]string(nil), _datazoneProjectIds...)
	}
	if len(_datazoneRuleType) > 0 {
		if err := assignInputField(input, "RuleType", _datazoneRuleType); err != nil {
			log.Errorf("invalid --rule-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListRulesOutput
	p := datazone.NewListRulesPaginator(client, input)
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

// Lists subscription grants.
func datazone_ListSubscriptionGrants(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListSubscriptionGrantsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentId) > 0 {
		input.EnvironmentId = aws.String(_datazoneEnvironmentId)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneOwningGroupId) > 0 {
		input.OwningGroupId = aws.String(_datazoneOwningGroupId)
	}
	if len(_datazoneOwningIamPrincipalArn) > 0 {
		input.OwningIamPrincipalArn = aws.String(_datazoneOwningIamPrincipalArn)
	}
	if len(_datazoneOwningProjectId) > 0 {
		input.OwningProjectId = aws.String(_datazoneOwningProjectId)
	}
	if len(_datazoneOwningUserId) > 0 {
		input.OwningUserId = aws.String(_datazoneOwningUserId)
	}
	if len(_datazoneSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _datazoneSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_datazoneSubscribedListingId) > 0 {
		input.SubscribedListingId = aws.String(_datazoneSubscribedListingId)
	}
	if len(_datazoneSubscriptionId) > 0 {
		input.SubscriptionId = aws.String(_datazoneSubscriptionId)
	}
	if len(_datazoneSubscriptionTargetId) > 0 {
		input.SubscriptionTargetId = aws.String(_datazoneSubscriptionTargetId)
	}

	if disablePaginator() {
		if resp, err := client.ListSubscriptionGrants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListSubscriptionGrantsOutput
	p := datazone.NewListSubscriptionGrantsPaginator(client, input)
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

// Lists Amazon DataZone subscription requests.
func datazone_ListSubscriptionRequests(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListSubscriptionRequestsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneApproverProjectId) > 0 {
		input.ApproverProjectId = aws.String(_datazoneApproverProjectId)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneOwningGroupId) > 0 {
		input.OwningGroupId = aws.String(_datazoneOwningGroupId)
	}
	if len(_datazoneOwningIamPrincipalArn) > 0 {
		input.OwningIamPrincipalArn = aws.String(_datazoneOwningIamPrincipalArn)
	}
	if len(_datazoneOwningProjectId) > 0 {
		input.OwningProjectId = aws.String(_datazoneOwningProjectId)
	}
	if len(_datazoneOwningUserId) > 0 {
		input.OwningUserId = aws.String(_datazoneOwningUserId)
	}
	if len(_datazoneSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _datazoneSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_datazoneSubscribedListingId) > 0 {
		input.SubscribedListingId = aws.String(_datazoneSubscribedListingId)
	}

	if disablePaginator() {
		if resp, err := client.ListSubscriptionRequests(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListSubscriptionRequestsOutput
	p := datazone.NewListSubscriptionRequestsPaginator(client, input)
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

// Lists subscription targets in Amazon DataZone.
func datazone_ListSubscriptionTargets(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListSubscriptionTargetsInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _datazoneSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSubscriptionTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListSubscriptionTargetsOutput
	p := datazone.NewListSubscriptionTargetsPaginator(client, input)
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

// Lists subscriptions in Amazon DataZone.
func datazone_ListSubscriptions(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListSubscriptionsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneApproverProjectId) > 0 {
		input.ApproverProjectId = aws.String(_datazoneApproverProjectId)
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneOwningGroupId) > 0 {
		input.OwningGroupId = aws.String(_datazoneOwningGroupId)
	}
	if len(_datazoneOwningIamPrincipalArn) > 0 {
		input.OwningIamPrincipalArn = aws.String(_datazoneOwningIamPrincipalArn)
	}
	if len(_datazoneOwningProjectId) > 0 {
		input.OwningProjectId = aws.String(_datazoneOwningProjectId)
	}
	if len(_datazoneOwningUserId) > 0 {
		input.OwningUserId = aws.String(_datazoneOwningUserId)
	}
	if len(_datazoneSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _datazoneSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_datazoneSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _datazoneSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_datazoneSubscribedListingId) > 0 {
		input.SubscribedListingId = aws.String(_datazoneSubscribedListingId)
	}
	if len(_datazoneSubscriptionRequestIdentifier) > 0 {
		input.SubscriptionRequestIdentifier = aws.String(_datazoneSubscriptionRequestIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.ListSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListSubscriptionsOutput
	p := datazone.NewListSubscriptionsPaginator(client, input)
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

// Lists tags for the specified resource in Amazon DataZone.
func datazone_ListTagsForResource(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_datazoneResourceArn) > 0 {
		input.ResourceArn = aws.String(_datazoneResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists time series data points.
func datazone_ListTimeSeriesDataPoints(cfg aws.Config, client *datazone.Client) {
	input := &datazone.ListTimeSeriesDataPointsInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.TimeSeriesEntityType, // Required
		// FormName: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneFormName) > 0 {
		input.FormName = aws.String(_datazoneFormName)
	}
	if len(_datazoneEndedAt) > 0 {
		if err := assignInputField(input, "EndedAt", _datazoneEndedAt); err != nil {
			log.Errorf("invalid --ended-at: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneStartedAt) > 0 {
		if err := assignInputField(input, "StartedAt", _datazoneStartedAt); err != nil {
			log.Errorf("invalid --started-at: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTimeSeriesDataPoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.ListTimeSeriesDataPointsOutput
	p := datazone.NewListTimeSeriesDataPointsPaginator(client, input)
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

// Posts a data lineage event.
func datazone_PostLineageEvent(cfg aws.Config, client *datazone.Client) {
	input := &datazone.PostLineageEventInput{
		// DomainIdentifier: *string, // Required
		// Event: []byte, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEvent) > 0 {
		if err := assignInputField(input, "Event", _datazoneEvent); err != nil {
			log.Errorf("invalid --event: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}

	if resp, err := client.PostLineageEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Posts time series data points to Amazon DataZone for the specified asset.
func datazone_PostTimeSeriesDataPoints(cfg aws.Config, client *datazone.Client) {
	input := &datazone.PostTimeSeriesDataPointsInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.TimeSeriesEntityType, // Required
		// Forms: []types.TimeSeriesDataPointFormInput, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneForms) > 0 {
		if err := assignInputField(input, "Forms", _datazoneForms); err != nil {
			log.Errorf("invalid --forms: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}

	if resp, err := client.PostTimeSeriesDataPoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates data export configuration details.
// If you want to temporarily disable export and later re-enable it for the same
// domain, use the --no-enable-export flag to disable and the --enable-export flag
// to re-enable. This preserves the configuration and allows you to re-enable
// export without deleting S3 table.
//
// You can enable asset metadata export for only one domain per account per
// Region. To enable export for a different domain, complete the following steps:
//
// - Delete the export configuration for the currently enabled domain using the
// DeleteDataExportConfiguration operation.
//
// - Delete the asset S3 table under the aws-sagemaker-catalog S3 table bucket.
// We recommend backing up the S3 table before deletion.
//
// - Call the PutDataExportConfiguration API to enable export for the new domain.
func datazone_PutDataExportConfiguration(cfg aws.Config, client *datazone.Client) {
	input := &datazone.PutDataExportConfigurationInput{
		// DomainIdentifier: *string, // Required
		// EnableExport: *bool, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnableExport) > 0 {
		if err := assignInputField(input, "EnableExport", _datazoneEnableExport); err != nil {
			log.Errorf("invalid --enable-export: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _datazoneEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDataExportConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Writes the configuration for the specified environment blueprint in Amazon
// DataZone.
func datazone_PutEnvironmentBlueprintConfiguration(cfg aws.Config, client *datazone.Client) {
	input := &datazone.PutEnvironmentBlueprintConfigurationInput{
		// DomainIdentifier: *string, // Required
		// EnabledRegions: []string, // Required
		// EnvironmentBlueprintIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnabledRegions) > 0 {
		input.EnabledRegions = append([]string(nil), _datazoneEnabledRegions...)
	}
	if len(_datazoneEnvironmentBlueprintIdentifier) > 0 {
		input.EnvironmentBlueprintIdentifier = aws.String(_datazoneEnvironmentBlueprintIdentifier)
	}
	if len(_datazoneEnvironmentRolePermissionBoundary) > 0 {
		input.EnvironmentRolePermissionBoundary = aws.String(_datazoneEnvironmentRolePermissionBoundary)
	}
	if len(_datazoneGlobalParameters) > 0 {
		if err := assignInputField(input, "GlobalParameters", _datazoneGlobalParameters); err != nil {
			log.Errorf("invalid --global-parameters: %s", err.Error())
			return
		}
	}
	if len(_datazoneManageAccessRoleArn) > 0 {
		input.ManageAccessRoleArn = aws.String(_datazoneManageAccessRoleArn)
	}
	if len(_datazoneProvisioningConfigurations) > 0 {
		if err := assignInputField(input, "ProvisioningConfigurations", _datazoneProvisioningConfigurations); err != nil {
			log.Errorf("invalid --provisioning-configurations: %s", err.Error())
			return
		}
	}
	if len(_datazoneProvisioningRoleArn) > 0 {
		input.ProvisioningRoleArn = aws.String(_datazoneProvisioningRoleArn)
	}
	if len(_datazoneRegionalParameters) > 0 {
		if err := assignInputField(input, "RegionalParameters", _datazoneRegionalParameters); err != nil {
			log.Errorf("invalid --regional-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEnvironmentBlueprintConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Queries entities in the graph store.
func datazone_QueryGraph(cfg aws.Config, client *datazone.Client) {
	input := &datazone.QueryGraphInput{
		// DomainIdentifier: *string, // Required
		// Match: []types.MatchClause, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneMatch) > 0 {
		if err := assignInputField(input, "Match", _datazoneMatch); err != nil {
			log.Errorf("invalid --match: %s", err.Error())
			return
		}
	}
	if len(_datazoneAdditionalAttributes) > 0 {
		if err := assignInputField(input, "AdditionalAttributes", _datazoneAdditionalAttributes); err != nil {
			log.Errorf("invalid --additional-attributes: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}

	if disablePaginator() {
		if resp, err := client.QueryGraph(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.QueryGraphOutput
	p := datazone.NewQueryGraphPaginator(client, input)
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

// Rejects automatically generated business-friendly metadata for your Amazon
// DataZone assets.
func datazone_RejectPredictions(cfg aws.Config, client *datazone.Client) {
	input := &datazone.RejectPredictionsInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneRejectChoices) > 0 {
		if err := assignInputField(input, "RejectChoices", _datazoneRejectChoices); err != nil {
			log.Errorf("invalid --reject-choices: %s", err.Error())
			return
		}
	}
	if len(_datazoneRejectRule) > 0 {
		if err := assignInputField(input, "RejectRule", _datazoneRejectRule); err != nil {
			log.Errorf("invalid --reject-rule: %s", err.Error())
			return
		}
	}
	if len(_datazoneRevision) > 0 {
		input.Revision = aws.String(_datazoneRevision)
	}

	if resp, err := client.RejectPredictions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects the specified subscription request.
func datazone_RejectSubscriptionRequest(cfg aws.Config, client *datazone.Client) {
	input := &datazone.RejectSubscriptionRequestInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneDecisionComment) > 0 {
		input.DecisionComment = aws.String(_datazoneDecisionComment)
	}

	if resp, err := client.RejectSubscriptionRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an owner from an entity.
func datazone_RemoveEntityOwner(cfg aws.Config, client *datazone.Client) {
	input := &datazone.RemoveEntityOwnerInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.DataZoneEntityType, // Required
		// Owner: types.OwnerProperties, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneOwner) > 0 {
		if err := assignInputField(input, "Owner", _datazoneOwner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}

	if resp, err := client.RemoveEntityOwner(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a policy grant.
func datazone_RemovePolicyGrant(cfg aws.Config, client *datazone.Client) {
	input := &datazone.RemovePolicyGrantInput{
		// DomainIdentifier: *string, // Required
		// EntityIdentifier: *string, // Required
		// EntityType: types.TargetEntityType, // Required
		// PolicyType: types.ManagedPolicyType, // Required
		// Principal: types.PolicyGrantPrincipal, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEntityIdentifier) > 0 {
		input.EntityIdentifier = aws.String(_datazoneEntityIdentifier)
	}
	if len(_datazoneEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _datazoneEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_datazonePolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _datazonePolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_datazonePrincipal) > 0 {
		if err := assignInputField(input, "Principal", _datazonePrincipal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneGrantIdentifier) > 0 {
		input.GrantIdentifier = aws.String(_datazoneGrantIdentifier)
	}

	if resp, err := client.RemovePolicyGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes a specified subscription in Amazon DataZone.
func datazone_RevokeSubscription(cfg aws.Config, client *datazone.Client) {
	input := &datazone.RevokeSubscriptionInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneRetainPermissions) > 0 {
		if err := assignInputField(input, "RetainPermissions", _datazoneRetainPermissions); err != nil {
			log.Errorf("invalid --retain-permissions: %s", err.Error())
			return
		}
	}

	if resp, err := client.RevokeSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for assets in Amazon DataZone.
// Search in Amazon DataZone is a powerful capability that enables users to
// discover and explore data assets, glossary terms, and data products across their
// organization. It provides both basic and advanced search functionality, allowing
// users to find resources based on names, descriptions, metadata, and other
// attributes. Search can be scoped to specific types of resources (like assets,
// glossary terms, or data products) and can be filtered using various criteria
// such as creation date, owner, or status. The search functionality is essential
// for making the wealth of data resources in an organization discoverable and
// usable, helping users find the right data for their needs quickly and
// efficiently.
//
// Many search commands in Amazon DataZone are paginated, including search and
// search-types . When the result set is large, Amazon DataZone returns a nextToken
// in the response. This token can be used to retrieve the next page of results.
//
// Prerequisites:
//
// - The --domain-identifier must refer to an existing Amazon DataZone domain.
//
// - --search-scope must be one of: ASSET, GLOSSARY_TERM, DATA_PRODUCT, or
// GLOSSARY.
//
// - The user must have search permissions in the specified domain.
//
// - If using --filters, ensure that the JSON is well-formed and that each
// filter includes valid attribute and value keys.
//
// - For paginated results, be prepared to use --next-token to fetch additional
// pages.
//
// To run a standard free-text search, the searchText parameter must be supplied.
// By default, all searchable fields are indexed for semantic search and will
// return semantic matches for SearchListings queries. To prevent semantic search
// indexing for a custom form attribute, see the [CreateFormType API documentation]. To run a lexical search query,
// enclose the query with double quotes (""). This will disable semantic search
// even for fields that have semantic search enabled and will only return results
// that contain the keywords wrapped by double quotes (order of tokens in the query
// is not enforced). Free-text search is supported for all attributes annotated
// with (at)amazon.datazone#searchable.
//
// To run a filtered search, provide filter clause using the filters parameter. To
// filter on glossary terms, use the special attribute __DataZoneGlossaryTerms . To
// filter on an indexed numeric attribute (i.e., a numeric attribute annotated with
// (at)amazon.datazone#sortable ), provide a filter using the intValue parameter. The
// filters parameter can also be used to run more advanced free-text searches that
// target specific attributes (attributes must be annotated with
// (at)amazon.datazone#searchable for free-text search). Create/update timestamp
// filtering is supported using the special creationTime / lastUpdatedTime
// attributes. Filter types can be mixed and matched to power complex queries.
//
// To find out whether an attribute has been annotated and indexed for a given
// search type, use the GetFormType API to retrieve the form containing the
// attribute.
//
// [CreateFormType API documentation]: https://docs.aws.amazon.com/datazone/latest/APIReference/API_CreateFormType.html
func datazone_Search(cfg aws.Config, client *datazone.Client) {
	input := &datazone.SearchInput{
		// DomainIdentifier: *string, // Required
		// SearchScope: types.InventorySearchScope, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneSearchScope) > 0 {
		if err := assignInputField(input, "SearchScope", _datazoneSearchScope); err != nil {
			log.Errorf("invalid --search-scope: %s", err.Error())
			return
		}
	}
	if len(_datazoneAdditionalAttributes) > 0 {
		if err := assignInputField(input, "AdditionalAttributes", _datazoneAdditionalAttributes); err != nil {
			log.Errorf("invalid --additional-attributes: %s", err.Error())
			return
		}
	}
	if len(_datazoneFilters) > 0 {
		if err := assignInputField(input, "Filters", _datazoneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneOwningProjectIdentifier) > 0 {
		input.OwningProjectIdentifier = aws.String(_datazoneOwningProjectIdentifier)
	}
	if len(_datazoneSearchIn) > 0 {
		if err := assignInputField(input, "SearchIn", _datazoneSearchIn); err != nil {
			log.Errorf("invalid --search-in: %s", err.Error())
			return
		}
	}
	if len(_datazoneSearchText) > 0 {
		input.SearchText = aws.String(_datazoneSearchText)
	}
	if len(_datazoneSort) > 0 {
		if err := assignInputField(input, "Sort", _datazoneSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.Search(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.SearchOutput
	p := datazone.NewSearchPaginator(client, input)
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

// Searches group profiles in Amazon DataZone.
func datazone_SearchGroupProfiles(cfg aws.Config, client *datazone.Client) {
	input := &datazone.SearchGroupProfilesInput{
		// DomainIdentifier: *string, // Required
		// GroupType: types.GroupSearchType, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneGroupType) > 0 {
		if err := assignInputField(input, "GroupType", _datazoneGroupType); err != nil {
			log.Errorf("invalid --group-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSearchText) > 0 {
		input.SearchText = aws.String(_datazoneSearchText)
	}

	if disablePaginator() {
		if resp, err := client.SearchGroupProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.SearchGroupProfilesOutput
	p := datazone.NewSearchGroupProfilesPaginator(client, input)
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

// Searches listings in Amazon DataZone.
// SearchListings is a powerful capability that enables users to discover and
// explore published assets and data products across their organization. It
// provides both basic and advanced search functionality, allowing users to find
// resources based on names, descriptions, metadata, and other attributes.
// SearchListings also supports filtering using various criteria such as creation
// date, owner, or status. This API is essential for making the wealth of data
// resources in an organization discoverable and usable, helping users find the
// right data for their needs quickly and efficiently.
//
// SearchListings returns results in a paginated format. When the result set is
// large, the response will include a nextToken, which can be used to retrieve the
// next page of results.
//
// The SearchListings API gives users flexibility in specifying what kind of
// search is run.
//
// To run a standard free-text search, the searchText parameter must be supplied.
// By default, all searchable fields are indexed for semantic search and will
// return semantic matches for SearchListings queries. To prevent semantic search
// indexing for a custom form attribute, see the [CreateFormType API documentation]. To run a lexical search query,
// enclose the query with double quotes (""). This will disable semantic search
// even for fields that have semantic search enabled and will only return results
// that contain the keywords wrapped by double quotes (order of tokens in the query
// is not enforced). Free-text search is supported for all attributes annotated
// with (at)amazon.datazone#searchable.
//
// To run a filtered search, provide filter clause using the filters parameter. To
// filter on glossary terms, use the special attribute __DataZoneGlossaryTerms . To
// filter on an indexed numeric attribute (i.e., a numeric attribute annotated with
// (at)amazon.datazone#sortable ), provide a filter using the intValue parameter. The
// filters parameter can also be used to run more advanced free-text searches that
// target specific attributes (attributes must be annotated with
// (at)amazon.datazone#searchable for free-text search). Create/update timestamp
// filtering is supported using the special creationTime / lastUpdatedTime
// attributes. Filter types can be mixed and matched to power complex queries.
//
// To find out whether an attribute has been annotated and indexed for a given
// search type, use the GetFormType API to retrieve the form containing the
// attribute.
//
// [CreateFormType API documentation]: https://docs.aws.amazon.com/datazone/latest/APIReference/API_CreateFormType.html
func datazone_SearchListings(cfg aws.Config, client *datazone.Client) {
	input := &datazone.SearchListingsInput{
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneAdditionalAttributes) > 0 {
		if err := assignInputField(input, "AdditionalAttributes", _datazoneAdditionalAttributes); err != nil {
			log.Errorf("invalid --additional-attributes: %s", err.Error())
			return
		}
	}
	if len(_datazoneAggregations) > 0 {
		if err := assignInputField(input, "Aggregations", _datazoneAggregations); err != nil {
			log.Errorf("invalid --aggregations: %s", err.Error())
			return
		}
	}
	if len(_datazoneFilters) > 0 {
		if err := assignInputField(input, "Filters", _datazoneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSearchIn) > 0 {
		if err := assignInputField(input, "SearchIn", _datazoneSearchIn); err != nil {
			log.Errorf("invalid --search-in: %s", err.Error())
			return
		}
	}
	if len(_datazoneSearchText) > 0 {
		input.SearchText = aws.String(_datazoneSearchText)
	}
	if len(_datazoneSort) > 0 {
		if err := assignInputField(input, "Sort", _datazoneSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchListings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.SearchListingsOutput
	p := datazone.NewSearchListingsPaginator(client, input)
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

// Searches for types in Amazon DataZone.
// Prerequisites:
//
// - The --domain-identifier must refer to an existing Amazon DataZone domain.
//
// - --search-scope must be one of the valid values including: ASSET_TYPE,
// GLOSSARY_TERM_TYPE, DATA_PRODUCT_TYPE.
//
// - The --managed flag must be present without a value.
//
// - The user must have permissions for form or asset types in the domain.
//
// - If using --filters, ensure that the JSON is valid.
//
// - Filters contain correct structure (attribute, value, operator).
func datazone_SearchTypes(cfg aws.Config, client *datazone.Client) {
	input := &datazone.SearchTypesInput{
		// DomainIdentifier: *string, // Required
		// Managed: *bool, // Required
		// SearchScope: types.TypesSearchScope, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneManaged) > 0 {
		if err := assignInputField(input, "Managed", _datazoneManaged); err != nil {
			log.Errorf("invalid --managed: %s", err.Error())
			return
		}
	}
	if len(_datazoneSearchScope) > 0 {
		if err := assignInputField(input, "SearchScope", _datazoneSearchScope); err != nil {
			log.Errorf("invalid --search-scope: %s", err.Error())
			return
		}
	}
	if len(_datazoneFilters) > 0 {
		if err := assignInputField(input, "Filters", _datazoneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSearchIn) > 0 {
		if err := assignInputField(input, "SearchIn", _datazoneSearchIn); err != nil {
			log.Errorf("invalid --search-in: %s", err.Error())
			return
		}
	}
	if len(_datazoneSearchText) > 0 {
		input.SearchText = aws.String(_datazoneSearchText)
	}
	if len(_datazoneSort) > 0 {
		if err := assignInputField(input, "Sort", _datazoneSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.SearchTypesOutput
	p := datazone.NewSearchTypesPaginator(client, input)
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

// Searches user profiles in Amazon DataZone.
func datazone_SearchUserProfiles(cfg aws.Config, client *datazone.Client) {
	input := &datazone.SearchUserProfilesInput{
		// DomainIdentifier: *string, // Required
		// UserType: types.UserSearchType, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneUserType) > 0 {
		if err := assignInputField(input, "UserType", _datazoneUserType); err != nil {
			log.Errorf("invalid --user-type: %s", err.Error())
			return
		}
	}
	if len(_datazoneMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _datazoneMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_datazoneNextToken) > 0 {
		input.NextToken = aws.String(_datazoneNextToken)
	}
	if len(_datazoneSearchText) > 0 {
		input.SearchText = aws.String(_datazoneSearchText)
	}

	if disablePaginator() {
		if resp, err := client.SearchUserProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*datazone.SearchUserProfilesOutput
	p := datazone.NewSearchUserProfilesPaginator(client, input)
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

// Start the run of the specified data source in Amazon DataZone.
func datazone_StartDataSourceRun(cfg aws.Config, client *datazone.Client) {
	input := &datazone.StartDataSourceRunInput{
		// DataSourceIdentifier: *string, // Required
		// DomainIdentifier: *string, // Required
	}

	if len(_datazoneDataSourceIdentifier) > 0 {
		input.DataSourceIdentifier = aws.String(_datazoneDataSourceIdentifier)
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}

	if resp, err := client.StartDataSourceRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the metadata generation run.
// Prerequisites:
//
// - Asset must be created and belong to the specified domain and project.
//
// - Asset type must be supported for metadata generation (e.g., Amazon Web
// Services Glue table).
//
// - Asset must have a structured schema with valid rows and columns.
//
// - Valid values for --type: BUSINESS_DESCRIPTIONS, BUSINESS_NAMES,
// BUSINESS_GLOSSARY_ASSOCIATIONS.
//
// - The user must have permission to run metadata generation in the
// domain/project.
func datazone_StartMetadataGenerationRun(cfg aws.Config, client *datazone.Client) {
	input := &datazone.StartMetadataGenerationRunInput{
		// DomainIdentifier: *string, // Required
		// OwningProjectIdentifier: *string, // Required
		// Target: *types.MetadataGenerationRunTarget, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneOwningProjectIdentifier) > 0 {
		input.OwningProjectIdentifier = aws.String(_datazoneOwningProjectIdentifier)
	}
	if len(_datazoneTarget) > 0 {
		if err := assignInputField(input, "Target", _datazoneTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneType) > 0 {
		if err := assignInputField(input, "Type", _datazoneType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_datazoneTypes) > 0 {
		if err := assignInputField(input, "Types", _datazoneTypes); err != nil {
			log.Errorf("invalid --types: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMetadataGenerationRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a resource in Amazon DataZone.
func datazone_TagResource(cfg aws.Config, client *datazone.Client) {
	input := &datazone.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_datazoneResourceArn) > 0 {
		input.ResourceArn = aws.String(_datazoneResourceArn)
	}
	if len(_datazoneTags) > 0 {
		if err := assignInputField(input, "Tags", _datazoneTags); err != nil {
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

// Untags a resource in Amazon DataZone.
func datazone_UntagResource(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_datazoneResourceArn) > 0 {
		input.ResourceArn = aws.String(_datazoneResourceArn)
	}
	if len(_datazoneTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _datazoneTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the account pool.
func datazone_UpdateAccountPool(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateAccountPoolInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneAccountSource) > 0 {
		if err := assignInputField(input, "AccountSource", _datazoneAccountSource); err != nil {
			log.Errorf("invalid --account-source: %s", err.Error())
			return
		}
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneResolutionStrategy) > 0 {
		if err := assignInputField(input, "ResolutionStrategy", _datazoneResolutionStrategy); err != nil {
			log.Errorf("invalid --resolution-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an asset filter.
// Prerequisites:
//
// - The domain, asset, and asset filter identifier must all exist.
//
// - The asset must contain the columns being referenced in the update.
//
// - If applying a row filter, ensure the column referenced in the expression
// exists in the asset schema.
func datazone_UpdateAssetFilter(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateAssetFilterInput{
		// AssetIdentifier: *string, // Required
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneAssetIdentifier) > 0 {
		input.AssetIdentifier = aws.String(_datazoneAssetIdentifier)
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _datazoneConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}

	if resp, err := client.UpdateAssetFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a connection. In Amazon DataZone, a connection enables you to connect
// your resources (domains, projects, and environments) to external resources and
// services.
func datazone_UpdateConnection(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateConnectionInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneAwsLocation) > 0 {
		if err := assignInputField(input, "AwsLocation", _datazoneAwsLocation); err != nil {
			log.Errorf("invalid --aws-location: %s", err.Error())
			return
		}
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneProps) > 0 {
		if err := assignInputField(input, "Props", _datazoneProps); err != nil {
			log.Errorf("invalid --props: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified data source in Amazon DataZone.
func datazone_UpdateDataSource(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateDataSourceInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneAssetFormsInput) > 0 {
		if err := assignInputField(input, "AssetFormsInput", _datazoneAssetFormsInput); err != nil {
			log.Errorf("invalid --asset-forms-input: %s", err.Error())
			return
		}
	}
	if len(_datazoneConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _datazoneConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneEnableSetting) > 0 {
		if err := assignInputField(input, "EnableSetting", _datazoneEnableSetting); err != nil {
			log.Errorf("invalid --enable-setting: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazonePublishOnImport) > 0 {
		if err := assignInputField(input, "PublishOnImport", _datazonePublishOnImport); err != nil {
			log.Errorf("invalid --publish-on-import: %s", err.Error())
			return
		}
	}
	if len(_datazoneRecommendation) > 0 {
		if err := assignInputField(input, "Recommendation", _datazoneRecommendation); err != nil {
			log.Errorf("invalid --recommendation: %s", err.Error())
			return
		}
	}
	if len(_datazoneRetainPermissionsOnRevokeFailure) > 0 {
		if err := assignInputField(input, "RetainPermissionsOnRevokeFailure", _datazoneRetainPermissionsOnRevokeFailure); err != nil {
			log.Errorf("invalid --retain-permissions-on-revoke-failure: %s", err.Error())
			return
		}
	}
	if len(_datazoneSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _datazoneSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Amazon DataZone domain.
func datazone_UpdateDomain(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateDomainInput{
		// Identifier: *string, // Required
	}

	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneDomainExecutionRole) > 0 {
		input.DomainExecutionRole = aws.String(_datazoneDomainExecutionRole)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneServiceRole) > 0 {
		input.ServiceRole = aws.String(_datazoneServiceRole)
	}
	if len(_datazoneSingleSignOn) > 0 {
		if err := assignInputField(input, "SingleSignOn", _datazoneSingleSignOn); err != nil {
			log.Errorf("invalid --single-sign-on: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the domain unit.
func datazone_UpdateDomainUnit(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateDomainUnitInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}

	if resp, err := client.UpdateDomainUnit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified environment in Amazon DataZone.
func datazone_UpdateEnvironment(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateEnvironmentInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneBlueprintVersion) > 0 {
		input.BlueprintVersion = aws.String(_datazoneBlueprintVersion)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneGlossaryTerms) > 0 {
		input.GlossaryTerms = append([]string(nil), _datazoneGlossaryTerms...)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneUserParameters) > 0 {
		if err := assignInputField(input, "UserParameters", _datazoneUserParameters); err != nil {
			log.Errorf("invalid --user-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an environment action.
func datazone_UpdateEnvironmentAction(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateEnvironmentActionInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneParameters) > 0 {
		if err := assignInputField(input, "Parameters", _datazoneParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnvironmentAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an environment blueprint in Amazon DataZone.
func datazone_UpdateEnvironmentBlueprint(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateEnvironmentBlueprintInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneProvisioningProperties) > 0 {
		if err := assignInputField(input, "ProvisioningProperties", _datazoneProvisioningProperties); err != nil {
			log.Errorf("invalid --provisioning-properties: %s", err.Error())
			return
		}
	}
	if len(_datazoneUserParameters) > 0 {
		if err := assignInputField(input, "UserParameters", _datazoneUserParameters); err != nil {
			log.Errorf("invalid --user-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnvironmentBlueprint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified environment profile in Amazon DataZone.
func datazone_UpdateEnvironmentProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateEnvironmentProfileInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_datazoneAwsAccountId)
	}
	if len(_datazoneAwsAccountRegion) > 0 {
		input.AwsAccountRegion = aws.String(_datazoneAwsAccountRegion)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneUserParameters) > 0 {
		if err := assignInputField(input, "UserParameters", _datazoneUserParameters); err != nil {
			log.Errorf("invalid --user-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEnvironmentProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the business glossary in Amazon DataZone.
// Prerequisites:
//
// - The glossary must exist in the given domain.
//
// - The caller must have the datazone:UpdateGlossary permission to update it.
//
// - When updating the name, the new name must be unique within the domain.
//
// - The glossary must not be deleted or in a terminal state.
func datazone_UpdateGlossary(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateGlossaryInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGlossary(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a business glossary term in Amazon DataZone.
// Prerequisites:
//
// - Glossary term must exist in the specified domain.
//
// - New name must not conflict with existing terms in the same glossary.
//
// - User must have permissions on the term.
//
// - The term must not be in DELETED status.
func datazone_UpdateGlossaryTerm(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateGlossaryTermInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneGlossaryIdentifier) > 0 {
		input.GlossaryIdentifier = aws.String(_datazoneGlossaryIdentifier)
	}
	if len(_datazoneLongDescription) > 0 {
		input.LongDescription = aws.String(_datazoneLongDescription)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneShortDescription) > 0 {
		input.ShortDescription = aws.String(_datazoneShortDescription)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_datazoneTermRelations) > 0 {
		if err := assignInputField(input, "TermRelations", _datazoneTermRelations); err != nil {
			log.Errorf("invalid --term-relations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGlossaryTerm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified group profile in Amazon DataZone.
func datazone_UpdateGroupProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateGroupProfileInput{
		// DomainIdentifier: *string, // Required
		// GroupIdentifier: *string, // Required
		// Status: types.GroupProfileStatus, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneGroupIdentifier) > 0 {
		input.GroupIdentifier = aws.String(_datazoneGroupIdentifier)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGroupProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified project in Amazon DataZone.
func datazone_UpdateProject(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateProjectInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneDomainUnitId) > 0 {
		input.DomainUnitId = aws.String(_datazoneDomainUnitId)
	}
	if len(_datazoneEnvironmentDeploymentDetails) > 0 {
		if err := assignInputField(input, "EnvironmentDeploymentDetails", _datazoneEnvironmentDeploymentDetails); err != nil {
			log.Errorf("invalid --environment-deployment-details: %s", err.Error())
			return
		}
	}
	if len(_datazoneGlossaryTerms) > 0 {
		input.GlossaryTerms = append([]string(nil), _datazoneGlossaryTerms...)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneProjectProfileVersion) > 0 {
		input.ProjectProfileVersion = aws.String(_datazoneProjectProfileVersion)
	}
	if len(_datazoneResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _datazoneResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}
	if len(_datazoneUserParameters) > 0 {
		if err := assignInputField(input, "UserParameters", _datazoneUserParameters); err != nil {
			log.Errorf("invalid --user-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a project profile.
func datazone_UpdateProjectProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateProjectProfileInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneAllowCustomProjectResourceTags) > 0 {
		if err := assignInputField(input, "AllowCustomProjectResourceTags", _datazoneAllowCustomProjectResourceTags); err != nil {
			log.Errorf("invalid --allow-custom-project-resource-tags: %s", err.Error())
			return
		}
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneDomainUnitIdentifier) > 0 {
		input.DomainUnitIdentifier = aws.String(_datazoneDomainUnitIdentifier)
	}
	if len(_datazoneEnvironmentConfigurations) > 0 {
		if err := assignInputField(input, "EnvironmentConfigurations", _datazoneEnvironmentConfigurations); err != nil {
			log.Errorf("invalid --environment-configurations: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneProjectResourceTags) > 0 {
		if err := assignInputField(input, "ProjectResourceTags", _datazoneProjectResourceTags); err != nil {
			log.Errorf("invalid --project-resource-tags: %s", err.Error())
			return
		}
	}
	if len(_datazoneProjectResourceTagsDescription) > 0 {
		input.ProjectResourceTagsDescription = aws.String(_datazoneProjectResourceTagsDescription)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProjectProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the owner of the root domain unit.
func datazone_UpdateRootDomainUnitOwner(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateRootDomainUnitOwnerInput{
		// CurrentOwner: *string, // Required
		// DomainIdentifier: *string, // Required
		// NewOwner: *string, // Required
	}

	if len(_datazoneCurrentOwner) > 0 {
		input.CurrentOwner = aws.String(_datazoneCurrentOwner)
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneNewOwner) > 0 {
		input.NewOwner = aws.String(_datazoneNewOwner)
	}
	if len(_datazoneClientToken) > 0 {
		input.ClientToken = aws.String(_datazoneClientToken)
	}

	if resp, err := client.UpdateRootDomainUnitOwner(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a rule. In Amazon DataZone, a rule is a formal agreement that enforces
// specific requirements across user workflows (e.g., publishing assets to the
// catalog, requesting subscriptions, creating projects) within the Amazon DataZone
// data portal. These rules help maintain consistency, ensure compliance, and
// uphold governance standards in data management processes. For instance, a
// metadata enforcement rule can specify the required information for creating a
// subscription request or publishing a data asset to the catalog, ensuring
// alignment with organizational standards.
func datazone_UpdateRule(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateRuleInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneDescription) > 0 {
		input.Description = aws.String(_datazoneDescription)
	}
	if len(_datazoneDetail) > 0 {
		if err := assignInputField(input, "Detail", _datazoneDetail); err != nil {
			log.Errorf("invalid --detail: %s", err.Error())
			return
		}
	}
	if len(_datazoneIncludeChildDomainUnits) > 0 {
		if err := assignInputField(input, "IncludeChildDomainUnits", _datazoneIncludeChildDomainUnits); err != nil {
			log.Errorf("invalid --include-child-domain-units: %s", err.Error())
			return
		}
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneScope) > 0 {
		if err := assignInputField(input, "Scope", _datazoneScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of the specified subscription grant status in Amazon
// DataZone.
func datazone_UpdateSubscriptionGrantStatus(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateSubscriptionGrantStatusInput{
		// AssetIdentifier: *string, // Required
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
		// Status: types.SubscriptionGrantStatus, // Required
	}

	if len(_datazoneAssetIdentifier) > 0 {
		input.AssetIdentifier = aws.String(_datazoneAssetIdentifier)
	}
	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_datazoneFailureCause) > 0 {
		if err := assignInputField(input, "FailureCause", _datazoneFailureCause); err != nil {
			log.Errorf("invalid --failure-cause: %s", err.Error())
			return
		}
	}
	if len(_datazoneTargetName) > 0 {
		input.TargetName = aws.String(_datazoneTargetName)
	}

	if resp, err := client.UpdateSubscriptionGrantStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a specified subscription request in Amazon DataZone.
func datazone_UpdateSubscriptionRequest(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateSubscriptionRequestInput{
		// DomainIdentifier: *string, // Required
		// Identifier: *string, // Required
		// RequestReason: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneRequestReason) > 0 {
		input.RequestReason = aws.String(_datazoneRequestReason)
	}

	if resp, err := client.UpdateSubscriptionRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified subscription target in Amazon DataZone.
func datazone_UpdateSubscriptionTarget(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateSubscriptionTargetInput{
		// DomainIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_datazoneEnvironmentIdentifier)
	}
	if len(_datazoneIdentifier) > 0 {
		input.Identifier = aws.String(_datazoneIdentifier)
	}
	if len(_datazoneApplicableAssetTypes) > 0 {
		input.ApplicableAssetTypes = append([]string(nil), _datazoneApplicableAssetTypes...)
	}
	if len(_datazoneAuthorizedPrincipals) > 0 {
		input.AuthorizedPrincipals = append([]string(nil), _datazoneAuthorizedPrincipals...)
	}
	if len(_datazoneManageAccessRole) > 0 {
		input.ManageAccessRole = aws.String(_datazoneManageAccessRole)
	}
	if len(_datazoneName) > 0 {
		input.Name = aws.String(_datazoneName)
	}
	if len(_datazoneProvider) > 0 {
		input.Provider = aws.String(_datazoneProvider)
	}
	if len(_datazoneSubscriptionGrantCreationMode) > 0 {
		if err := assignInputField(input, "SubscriptionGrantCreationMode", _datazoneSubscriptionGrantCreationMode); err != nil {
			log.Errorf("invalid --subscription-grant-creation-mode: %s", err.Error())
			return
		}
	}
	if len(_datazoneSubscriptionTargetConfig) > 0 {
		if err := assignInputField(input, "SubscriptionTargetConfig", _datazoneSubscriptionTargetConfig); err != nil {
			log.Errorf("invalid --subscription-target-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSubscriptionTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified user profile in Amazon DataZone.
func datazone_UpdateUserProfile(cfg aws.Config, client *datazone.Client) {
	input := &datazone.UpdateUserProfileInput{
		// DomainIdentifier: *string, // Required
		// Status: types.UserProfileStatus, // Required
		// UserIdentifier: *string, // Required
	}

	if len(_datazoneDomainIdentifier) > 0 {
		input.DomainIdentifier = aws.String(_datazoneDomainIdentifier)
	}
	if len(_datazoneStatus) > 0 {
		if err := assignInputField(input, "Status", _datazoneStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_datazoneUserIdentifier) > 0 {
		input.UserIdentifier = aws.String(_datazoneUserIdentifier)
	}
	if len(_datazoneType) > 0 {
		if err := assignInputField(input, "Type", _datazoneType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUserProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_datazoneCmd)
	_datazoneCmd.Flags().SortFlags = false

	_datazoneCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_datazoneCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_datazoneCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_datazoneCmd.Flags().StringVarP(&_datazoneAcceptChoices, "accept-choices", "", "", "Accept Choices")
	_datazoneCmd.Flags().StringVarP(&_datazoneAcceptRule, "accept-rule", "", "", "Accept Rule")
	_datazoneCmd.Flags().StringVarP(&_datazoneAccountSource, "account-source", "", "", "Account Source")
	_datazoneCmd.Flags().StringVarP(&_datazoneAction, "action", "", "", "Action")
	_datazoneCmd.Flags().StringVarP(&_datazoneAdditionalAttributes, "additional-attributes", "", "", "Additional Attributes")
	_datazoneCmd.Flags().StringVarP(&_datazoneAfterTimestamp, "after-timestamp", "", "", "After Timestamp")
	_datazoneCmd.Flags().StringVarP(&_datazoneAggregations, "aggregations", "", "", "Aggregations")
	_datazoneCmd.Flags().StringVarP(&_datazoneAllowCustomProjectResourceTags, "allow-custom-project-resource-tags", "", "", "Allow Custom Project Resource Tags")
	_datazoneCmd.Flags().StringSliceVarP(&_datazoneApplicableAssetTypes, "applicable-asset-types", "", nil, "Applicable Asset Types")
	_datazoneCmd.Flags().StringVarP(&_datazoneApproverProjectId, "approver-project-id", "", "", "Approver Project ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneAssetFormsInput, "asset-forms-input", "", "", "Asset Forms Input")
	_datazoneCmd.Flags().StringVarP(&_datazoneAssetIdentifier, "asset-identifier", "", "", "Asset Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneAssetPermissions, "asset-permissions", "", "", "Asset Permissions")
	_datazoneCmd.Flags().StringVarP(&_datazoneAssetScopes, "asset-scopes", "", "", "Asset Scopes")
	_datazoneCmd.Flags().StringVarP(&_datazoneAssetTargetNames, "asset-target-names", "", "", "Asset Target Names")
	_datazoneCmd.Flags().StringSliceVarP(&_datazoneAssetTypes, "asset-types", "", nil, "Asset Types")
	_datazoneCmd.Flags().StringSliceVarP(&_datazoneAttributeIdentifiers, "attribute-identifiers", "", nil, "Attribute Identifiers")
	_datazoneCmd.Flags().StringVarP(&_datazoneAttributes, "attributes", "", "", "Attributes")
	_datazoneCmd.Flags().StringSliceVarP(&_datazoneAuthorizedPrincipals, "authorized-principals", "", nil, "Authorized Principals")
	_datazoneCmd.Flags().StringVarP(&_datazoneAwsAccountId, "aws-account-id", "", "", "AWS Account ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneAwsAccountRegion, "aws-account-region", "", "", "AWS Account Region")
	_datazoneCmd.Flags().StringVarP(&_datazoneAwsLocation, "aws-location", "", "", "AWS Location")
	_datazoneCmd.Flags().StringVarP(&_datazoneBeforeTimestamp, "before-timestamp", "", "", "Before Timestamp")
	_datazoneCmd.Flags().StringVarP(&_datazoneBlueprintVersion, "blueprint-version", "", "", "Blueprint Version")
	_datazoneCmd.Flags().StringVarP(&_datazoneClientToken, "client-token", "", "", "Client Token")
	_datazoneCmd.Flags().StringVarP(&_datazoneConfiguration, "configuration", "", "", "Configuration")
	_datazoneCmd.Flags().StringVarP(&_datazoneConnectionIdentifier, "connection-identifier", "", "", "Connection Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneCurrentOwner, "current-owner", "", "", "Current Owner")
	_datazoneCmd.Flags().StringVarP(&_datazoneDataProduct, "data-product", "", "", "Data Product")
	_datazoneCmd.Flags().StringVarP(&_datazoneDataSourceIdentifier, "data-source-identifier", "", "", "Data Source Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneDecisionComment, "decision-comment", "", "", "Decision Comment")
	_datazoneCmd.Flags().StringVarP(&_datazoneDeploymentOrder, "deployment-order", "", "", "Deployment Order")
	_datazoneCmd.Flags().StringVarP(&_datazoneDescription, "description", "", "", "Description")
	_datazoneCmd.Flags().StringVarP(&_datazoneDesignation, "designation", "", "", "Designation")
	_datazoneCmd.Flags().StringVarP(&_datazoneDetail, "detail", "", "", "Detail")
	_datazoneCmd.Flags().StringVarP(&_datazoneDirection, "direction", "", "", "Direction")
	_datazoneCmd.Flags().StringVarP(&_datazoneDomainExecutionRole, "domain-execution-role", "", "", "Domain Execution Role")
	_datazoneCmd.Flags().StringVarP(&_datazoneDomainIdentifier, "domain-identifier", "", "", "Domain Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneDomainUnitId, "domain-unit-id", "", "", "Domain Unit ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneDomainUnitIdentifier, "domain-unit-identifier", "", "", "Domain Unit Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneDomainVersion, "domain-version", "", "", "Domain Version")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnableExport, "enable-export", "", "", "Enable Export")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnableSetting, "enable-setting", "", "", "Enable Setting")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnableTrustedIdentityPropagation, "enable-trusted-identity-propagation", "", "", "Enable Trusted Identity Propagation")
	_datazoneCmd.Flags().StringSliceVarP(&_datazoneEnabledRegions, "enabled-regions", "", nil, "Enabled Regions")
	_datazoneCmd.Flags().StringVarP(&_datazoneEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_datazoneCmd.Flags().StringVarP(&_datazoneEndedAt, "ended-at", "", "", "Ended At")
	_datazoneCmd.Flags().StringVarP(&_datazoneEntityIdentifier, "entity-identifier", "", "", "Entity Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneEntityRevision, "entity-revision", "", "", "Entity Revision")
	_datazoneCmd.Flags().StringVarP(&_datazoneEntityType, "entity-type", "", "", "Entity Type")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentAccountIdentifier, "environment-account-identifier", "", "", "Environment Account Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentAccountRegion, "environment-account-region", "", "", "Environment Account Region")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentBlueprintIdentifier, "environment-blueprint-identifier", "", "", "Environment Blueprint Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentConfigurationId, "environment-configuration-id", "", "", "Environment Configuration ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentConfigurations, "environment-configurations", "", "", "Environment Configurations")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentDeploymentDetails, "environment-deployment-details", "", "", "Environment Deployment Details")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentId, "environment-id", "", "", "Environment ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentIdentifier, "environment-identifier", "", "", "Environment Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentProfileIdentifier, "environment-profile-identifier", "", "", "Environment Profile Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentRoleArn, "environment-role-arn", "", "", "Environment Role ARN")
	_datazoneCmd.Flags().StringVarP(&_datazoneEnvironmentRolePermissionBoundary, "environment-role-permission-boundary", "", "", "Environment Role Permission Boundary")
	_datazoneCmd.Flags().StringVarP(&_datazoneEvent, "event", "", "", "Event")
	_datazoneCmd.Flags().StringVarP(&_datazoneEventTimestamp, "event-timestamp", "", "", "Event Timestamp")
	_datazoneCmd.Flags().StringVarP(&_datazoneEventTimestampGTE, "event-timestamp-gte", "", "", "Event Timestamp Gte")
	_datazoneCmd.Flags().StringVarP(&_datazoneEventTimestampLTE, "event-timestamp-lte", "", "", "Event Timestamp Lte")
	_datazoneCmd.Flags().StringVarP(&_datazoneExternalIdentifier, "external-identifier", "", "", "External Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneFailureCause, "failure-cause", "", "", "Failure Cause")
	_datazoneCmd.Flags().StringVarP(&_datazoneFilters, "filters", "", "", "Filters")
	_datazoneCmd.Flags().StringVarP(&_datazoneFormName, "form-name", "", "", "Form Name")
	_datazoneCmd.Flags().StringVarP(&_datazoneFormTypeIdentifier, "form-type-identifier", "", "", "Form Type Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneForms, "forms", "", "", "Forms")
	_datazoneCmd.Flags().StringVarP(&_datazoneFormsInput, "forms-input", "", "", "Forms Input")
	_datazoneCmd.Flags().StringVarP(&_datazoneGlobalParameters, "global-parameters", "", "", "Global Parameters")
	_datazoneCmd.Flags().StringVarP(&_datazoneGlossaryIdentifier, "glossary-identifier", "", "", "Glossary Identifier")
	_datazoneCmd.Flags().StringSliceVarP(&_datazoneGlossaryTerms, "glossary-terms", "", nil, "Glossary Terms")
	_datazoneCmd.Flags().StringSliceVarP(&_datazoneGovernedGlossaryTerms, "governed-glossary-terms", "", nil, "Governed Glossary Terms")
	_datazoneCmd.Flags().StringVarP(&_datazoneGrantIdentifier, "grant-identifier", "", "", "Grant Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneGrantedEntity, "granted-entity", "", "", "Granted Entity")
	_datazoneCmd.Flags().StringVarP(&_datazoneGroupIdentifier, "group-identifier", "", "", "Group Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneGroupType, "group-type", "", "", "Group Type")
	_datazoneCmd.Flags().StringVarP(&_datazoneIdentifier, "identifier", "", "", "Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneIncludeCascaded, "include-cascaded", "", "", "Include Cascaded")
	_datazoneCmd.Flags().StringVarP(&_datazoneIncludeChildDomainUnits, "include-child-domain-units", "", "", "Include Child Domain Units")
	_datazoneCmd.Flags().StringVarP(&_datazoneItems, "items", "", "", "Items")
	_datazoneCmd.Flags().StringVarP(&_datazoneJobIdentifier, "job-identifier", "", "", "Job Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneKmsKeyIdentifier, "kms-key-identifier", "", "", "KMS Key Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneListingRevision, "listing-revision", "", "", "Listing Revision")
	_datazoneCmd.Flags().StringVarP(&_datazoneLongDescription, "long-description", "", "", "Long Description")
	_datazoneCmd.Flags().StringVarP(&_datazoneManageAccessRole, "manage-access-role", "", "", "Manage Access Role")
	_datazoneCmd.Flags().StringVarP(&_datazoneManageAccessRoleArn, "manage-access-role-arn", "", "", "Manage Access Role ARN")
	_datazoneCmd.Flags().StringVarP(&_datazoneManaged, "managed", "", "", "Managed")
	_datazoneCmd.Flags().StringVarP(&_datazoneMatch, "match", "", "", "Match")
	_datazoneCmd.Flags().StringVarP(&_datazoneMaxResults, "max-results", "", "", "Max Results")
	_datazoneCmd.Flags().StringVarP(&_datazoneMember, "member", "", "", "Member")
	_datazoneCmd.Flags().StringVarP(&_datazoneMetadataForms, "metadata-forms", "", "", "Metadata Forms")
	_datazoneCmd.Flags().StringVarP(&_datazoneModel, "model", "", "", "Model")
	_datazoneCmd.Flags().StringVarP(&_datazoneName, "name", "", "", "Name")
	_datazoneCmd.Flags().StringVarP(&_datazoneNewOwner, "new-owner", "", "", "New Owner")
	_datazoneCmd.Flags().StringVarP(&_datazoneNextToken, "next-token", "", "", "Next Token")
	_datazoneCmd.Flags().StringVarP(&_datazoneOwner, "owner", "", "", "Owner")
	_datazoneCmd.Flags().StringVarP(&_datazoneOwningGroupId, "owning-group-id", "", "", "Owning Group ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneOwningIamPrincipalArn, "owning-iam-principal-arn", "", "", "Owning IAM Principal ARN")
	_datazoneCmd.Flags().StringVarP(&_datazoneOwningProjectId, "owning-project-id", "", "", "Owning Project ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneOwningProjectIdentifier, "owning-project-identifier", "", "", "Owning Project Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneOwningUserId, "owning-user-id", "", "", "Owning User ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneParameters, "parameters", "", "", "Parameters")
	_datazoneCmd.Flags().StringVarP(&_datazoneParentDomainUnitIdentifier, "parent-domain-unit-identifier", "", "", "Parent Domain Unit Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazonePolicyType, "policy-type", "", "", "Policy Type")
	_datazoneCmd.Flags().StringVarP(&_datazonePredictionConfiguration, "prediction-configuration", "", "", "Prediction Configuration")
	_datazoneCmd.Flags().StringVarP(&_datazonePrincipal, "principal", "", "", "Principal")
	_datazoneCmd.Flags().StringVarP(&_datazoneProcessingStatus, "processing-status", "", "", "Processing Status")
	_datazoneCmd.Flags().StringVarP(&_datazoneProjectIdentifier, "project-identifier", "", "", "Project Identifier")
	_datazoneCmd.Flags().StringSliceVarP(&_datazoneProjectIds, "project-ids", "", nil, "Project Ids")
	_datazoneCmd.Flags().StringVarP(&_datazoneProjectProfileId, "project-profile-id", "", "", "Project Profile ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneProjectProfileVersion, "project-profile-version", "", "", "Project Profile Version")
	_datazoneCmd.Flags().StringVarP(&_datazoneProjectResourceTags, "project-resource-tags", "", "", "Project Resource Tags")
	_datazoneCmd.Flags().StringVarP(&_datazoneProjectResourceTagsDescription, "project-resource-tags-description", "", "", "Project Resource Tags Description")
	_datazoneCmd.Flags().StringVarP(&_datazoneProps, "props", "", "", "Props")
	_datazoneCmd.Flags().StringVarP(&_datazoneProvider, "provider", "", "", "Provider")
	_datazoneCmd.Flags().StringVarP(&_datazoneProvisioningConfigurations, "provisioning-configurations", "", "", "Provisioning Configurations")
	_datazoneCmd.Flags().StringVarP(&_datazoneProvisioningProperties, "provisioning-properties", "", "", "Provisioning Properties")
	_datazoneCmd.Flags().StringVarP(&_datazoneProvisioningRoleArn, "provisioning-role-arn", "", "", "Provisioning Role ARN")
	_datazoneCmd.Flags().StringVarP(&_datazonePublishOnImport, "publish-on-import", "", "", "Publish On Import")
	_datazoneCmd.Flags().StringVarP(&_datazoneRecommendation, "recommendation", "", "", "Recommendation")
	_datazoneCmd.Flags().StringVarP(&_datazoneRegionalParameters, "regional-parameters", "", "", "Regional Parameters")
	_datazoneCmd.Flags().StringVarP(&_datazoneRejectChoices, "reject-choices", "", "", "Reject Choices")
	_datazoneCmd.Flags().StringVarP(&_datazoneRejectRule, "reject-rule", "", "", "Reject Rule")
	_datazoneCmd.Flags().StringVarP(&_datazoneRequestReason, "request-reason", "", "", "Request Reason")
	_datazoneCmd.Flags().StringVarP(&_datazoneResolutionStrategy, "resolution-strategy", "", "", "Resolution Strategy")
	_datazoneCmd.Flags().StringVarP(&_datazoneResourceArn, "resource-arn", "", "", "Resource ARN")
	_datazoneCmd.Flags().StringVarP(&_datazoneResourceTags, "resource-tags", "", "", "Resource Tags")
	_datazoneCmd.Flags().StringVarP(&_datazoneRetainPermissions, "retain-permissions", "", "", "Retain Permissions")
	_datazoneCmd.Flags().StringVarP(&_datazoneRetainPermissionsOnRevokeFailure, "retain-permissions-on-revoke-failure", "", "", "Retain Permissions On Revoke Failure")
	_datazoneCmd.Flags().StringVarP(&_datazoneRevision, "revision", "", "", "Revision")
	_datazoneCmd.Flags().StringVarP(&_datazoneRuleType, "rule-type", "", "", "Rule Type")
	_datazoneCmd.Flags().StringVarP(&_datazoneSchedule, "schedule", "", "", "Schedule")
	_datazoneCmd.Flags().StringVarP(&_datazoneScope, "scope", "", "", "Scope")
	_datazoneCmd.Flags().StringVarP(&_datazoneSearchIn, "search-in", "", "", "Search In")
	_datazoneCmd.Flags().StringVarP(&_datazoneSearchScope, "search-scope", "", "", "Search Scope")
	_datazoneCmd.Flags().StringVarP(&_datazoneSearchText, "search-text", "", "", "Search Text")
	_datazoneCmd.Flags().StringVarP(&_datazoneServiceRole, "service-role", "", "", "Service Role")
	_datazoneCmd.Flags().StringVarP(&_datazoneShortDescription, "short-description", "", "", "Short Description")
	_datazoneCmd.Flags().StringVarP(&_datazoneSingleSignOn, "single-sign-on", "", "", "Single Sign On")
	_datazoneCmd.Flags().StringVarP(&_datazoneSkipDeletionCheck, "skip-deletion-check", "", "", "Skip Deletion Check")
	_datazoneCmd.Flags().StringVarP(&_datazoneSort, "sort", "", "", "Sort")
	_datazoneCmd.Flags().StringVarP(&_datazoneSortBy, "sort-by", "", "", "Sort By")
	_datazoneCmd.Flags().StringVarP(&_datazoneSortOrder, "sort-order", "", "", "Sort Order")
	_datazoneCmd.Flags().StringVarP(&_datazoneStartedAt, "started-at", "", "", "Started At")
	_datazoneCmd.Flags().StringVarP(&_datazoneStatus, "status", "", "", "Status")
	_datazoneCmd.Flags().StringSliceVarP(&_datazoneSubjects, "subjects", "", nil, "Subjects")
	_datazoneCmd.Flags().StringVarP(&_datazoneSubscribedListingId, "subscribed-listing-id", "", "", "Subscribed Listing ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneSubscribedListings, "subscribed-listings", "", "", "Subscribed Listings")
	_datazoneCmd.Flags().StringVarP(&_datazoneSubscribedPrincipals, "subscribed-principals", "", "", "Subscribed Principals")
	_datazoneCmd.Flags().StringVarP(&_datazoneSubscriptionGrantCreationMode, "subscription-grant-creation-mode", "", "", "Subscription Grant Creation Mode")
	_datazoneCmd.Flags().StringVarP(&_datazoneSubscriptionId, "subscription-id", "", "", "Subscription ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneSubscriptionRequestIdentifier, "subscription-request-identifier", "", "", "Subscription Request Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneSubscriptionTargetConfig, "subscription-target-config", "", "", "Subscription Target Config")
	_datazoneCmd.Flags().StringVarP(&_datazoneSubscriptionTargetId, "subscription-target-id", "", "", "Subscription Target ID")
	_datazoneCmd.Flags().StringVarP(&_datazoneSubscriptionTargetIdentifier, "subscription-target-identifier", "", "", "Subscription Target Identifier")
	_datazoneCmd.Flags().StringSliceVarP(&_datazoneTagKeys, "tag-keys", "", nil, "Tag Keys")
	_datazoneCmd.Flags().StringVarP(&_datazoneTags, "tags", "", "", "Tags")
	_datazoneCmd.Flags().StringVarP(&_datazoneTarget, "target", "", "", "Target")
	_datazoneCmd.Flags().StringVarP(&_datazoneTargetIdentifier, "target-identifier", "", "", "Target Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneTargetName, "target-name", "", "", "Target Name")
	_datazoneCmd.Flags().StringVarP(&_datazoneTargetType, "target-type", "", "", "Target Type")
	_datazoneCmd.Flags().StringVarP(&_datazoneTaskStatus, "task-status", "", "", "Task Status")
	_datazoneCmd.Flags().StringVarP(&_datazoneTermRelations, "term-relations", "", "", "Term Relations")
	_datazoneCmd.Flags().StringVarP(&_datazoneTimestampAfter, "timestamp-after", "", "", "Timestamp After")
	_datazoneCmd.Flags().StringVarP(&_datazoneTimestampBefore, "timestamp-before", "", "", "Timestamp Before")
	_datazoneCmd.Flags().StringVarP(&_datazoneType, "type", "", "", "Type")
	_datazoneCmd.Flags().StringVarP(&_datazoneTypeIdentifier, "type-identifier", "", "", "Type Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneTypeRevision, "type-revision", "", "", "Type Revision")
	_datazoneCmd.Flags().StringVarP(&_datazoneTypes, "types", "", "", "Types")
	_datazoneCmd.Flags().StringVarP(&_datazoneUsageRestrictions, "usage-restrictions", "", "", "Usage Restrictions")
	_datazoneCmd.Flags().StringVarP(&_datazoneUserIdentifier, "user-identifier", "", "", "User Identifier")
	_datazoneCmd.Flags().StringVarP(&_datazoneUserParameters, "user-parameters", "", "", "User Parameters")
	_datazoneCmd.Flags().StringVarP(&_datazoneUserType, "user-type", "", "", "User Type")
	_datazoneCmd.Flags().StringVarP(&_datazoneWithSecret, "with-secret", "", "", "With Secret")

	_datazoneCmd.Flags().BoolVarP(&_datazoneAcceptPredictions, "accept-predictions", "", false, "Accept Predictions")
	_datazoneCmd.Flags().BoolVarP(&_datazoneAcceptSubscriptionRequest, "accept-subscription-request", "", false, "Accept Subscription Request")
	_datazoneCmd.Flags().BoolVarP(&_datazoneAddEntityOwner, "add-entity-owner", "", false, "Add Entity Owner")
	_datazoneCmd.Flags().BoolVarP(&_datazoneAddPolicyGrant, "add-policy-grant", "", false, "Add Policy Grant")
	_datazoneCmd.Flags().BoolVarP(&_datazoneAssociateEnvironmentRole, "associate-environment-role", "", false, "Associate Environment Role")
	_datazoneCmd.Flags().BoolVarP(&_datazoneAssociateGovernedTerms, "associate-governed-terms", "", false, "Associate Governed Terms")
	_datazoneCmd.Flags().BoolVarP(&_datazoneBatchGetAttributesMetadata, "batch-get-attributes-metadata", "", false, "Batch Get Attributes Metadata")
	_datazoneCmd.Flags().BoolVarP(&_datazoneBatchPutAttributesMetadata, "batch-put-attributes-metadata", "", false, "Batch Put Attributes Metadata")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCancelMetadataGenerationRun, "cancel-metadata-generation-run", "", false, "Cancel Metadata Generation Run")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCancelSubscription, "cancel-subscription", "", false, "Cancel Subscription")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateAccountPool, "create-account-pool", "", false, "Create Account Pool")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateAsset, "create-asset", "", false, "Create Asset")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateAssetFilter, "create-asset-filter", "", false, "Create Asset Filter")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateAssetRevision, "create-asset-revision", "", false, "Create Asset Revision")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateAssetType, "create-asset-type", "", false, "Create Asset Type")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateConnection, "create-connection", "", false, "Create Connection")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateDataProduct, "create-data-product", "", false, "Create Data Product")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateDataProductRevision, "create-data-product-revision", "", false, "Create Data Product Revision")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateDataSource, "create-data-source", "", false, "Create Data Source")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateDomain, "create-domain", "", false, "Create Domain")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateDomainUnit, "create-domain-unit", "", false, "Create Domain Unit")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateEnvironment, "create-environment", "", false, "Create Environment")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateEnvironmentAction, "create-environment-action", "", false, "Create Environment Action")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateEnvironmentBlueprint, "create-environment-blueprint", "", false, "Create Environment Blueprint")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateEnvironmentProfile, "create-environment-profile", "", false, "Create Environment Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateFormType, "create-form-type", "", false, "Create Form Type")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateGlossary, "create-glossary", "", false, "Create Glossary")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateGlossaryTerm, "create-glossary-term", "", false, "Create Glossary Term")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateGroupProfile, "create-group-profile", "", false, "Create Group Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateListingChangeSet, "create-listing-change-set", "", false, "Create Listing Change Set")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateProject, "create-project", "", false, "Create Project")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateProjectMembership, "create-project-membership", "", false, "Create Project Membership")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateProjectProfile, "create-project-profile", "", false, "Create Project Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateRule, "create-rule", "", false, "Create Rule")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateSubscriptionGrant, "create-subscription-grant", "", false, "Create Subscription Grant")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateSubscriptionRequest, "create-subscription-request", "", false, "Create Subscription Request")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateSubscriptionTarget, "create-subscription-target", "", false, "Create Subscription Target")
	_datazoneCmd.Flags().BoolVarP(&_datazoneCreateUserProfile, "create-user-profile", "", false, "Create User Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteAccountPool, "delete-account-pool", "", false, "Delete Account Pool")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteAsset, "delete-asset", "", false, "Delete Asset")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteAssetFilter, "delete-asset-filter", "", false, "Delete Asset Filter")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteAssetType, "delete-asset-type", "", false, "Delete Asset Type")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteDataExportConfiguration, "delete-data-export-configuration", "", false, "Delete Data Export Configuration")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteDataProduct, "delete-data-product", "", false, "Delete Data Product")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteDataSource, "delete-data-source", "", false, "Delete Data Source")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteDomain, "delete-domain", "", false, "Delete Domain")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteDomainUnit, "delete-domain-unit", "", false, "Delete Domain Unit")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteEnvironment, "delete-environment", "", false, "Delete Environment")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteEnvironmentAction, "delete-environment-action", "", false, "Delete Environment Action")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteEnvironmentBlueprint, "delete-environment-blueprint", "", false, "Delete Environment Blueprint")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteEnvironmentBlueprintConfiguration, "delete-environment-blueprint-configuration", "", false, "Delete Environment Blueprint Configuration")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteEnvironmentProfile, "delete-environment-profile", "", false, "Delete Environment Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteFormType, "delete-form-type", "", false, "Delete Form Type")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteGlossary, "delete-glossary", "", false, "Delete Glossary")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteGlossaryTerm, "delete-glossary-term", "", false, "Delete Glossary Term")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteListing, "delete-listing", "", false, "Delete Listing")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteProject, "delete-project", "", false, "Delete Project")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteProjectMembership, "delete-project-membership", "", false, "Delete Project Membership")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteProjectProfile, "delete-project-profile", "", false, "Delete Project Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteRule, "delete-rule", "", false, "Delete Rule")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteSubscriptionGrant, "delete-subscription-grant", "", false, "Delete Subscription Grant")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteSubscriptionRequest, "delete-subscription-request", "", false, "Delete Subscription Request")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteSubscriptionTarget, "delete-subscription-target", "", false, "Delete Subscription Target")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDeleteTimeSeriesDataPoints, "delete-time-series-data-points", "", false, "Delete Time Series Data Points")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDisassociateEnvironmentRole, "disassociate-environment-role", "", false, "Disassociate Environment Role")
	_datazoneCmd.Flags().BoolVarP(&_datazoneDisassociateGovernedTerms, "disassociate-governed-terms", "", false, "Disassociate Governed Terms")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetAccountPool, "get-account-pool", "", false, "Get Account Pool")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetAsset, "get-asset", "", false, "Get Asset")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetAssetFilter, "get-asset-filter", "", false, "Get Asset Filter")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetAssetType, "get-asset-type", "", false, "Get Asset Type")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetConnection, "get-connection", "", false, "Get Connection")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetDataExportConfiguration, "get-data-export-configuration", "", false, "Get Data Export Configuration")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetDataProduct, "get-data-product", "", false, "Get Data Product")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetDataSource, "get-data-source", "", false, "Get Data Source")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetDataSourceRun, "get-data-source-run", "", false, "Get Data Source Run")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetDomain, "get-domain", "", false, "Get Domain")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetDomainUnit, "get-domain-unit", "", false, "Get Domain Unit")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetEnvironment, "get-environment", "", false, "Get Environment")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetEnvironmentAction, "get-environment-action", "", false, "Get Environment Action")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetEnvironmentBlueprint, "get-environment-blueprint", "", false, "Get Environment Blueprint")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetEnvironmentBlueprintConfiguration, "get-environment-blueprint-configuration", "", false, "Get Environment Blueprint Configuration")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetEnvironmentCredentials, "get-environment-credentials", "", false, "Get Environment Credentials")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetEnvironmentProfile, "get-environment-profile", "", false, "Get Environment Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetFormType, "get-form-type", "", false, "Get Form Type")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetGlossary, "get-glossary", "", false, "Get Glossary")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetGlossaryTerm, "get-glossary-term", "", false, "Get Glossary Term")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetGroupProfile, "get-group-profile", "", false, "Get Group Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetIamPortalLoginUrl, "get-iam-portal-login-url", "", false, "Get IAM Portal Login URL")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetJobRun, "get-job-run", "", false, "Get Job Run")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetLineageEvent, "get-lineage-event", "", false, "Get Lineage Event")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetLineageNode, "get-lineage-node", "", false, "Get Lineage Node")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetListing, "get-listing", "", false, "Get Listing")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetMetadataGenerationRun, "get-metadata-generation-run", "", false, "Get Metadata Generation Run")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetProject, "get-project", "", false, "Get Project")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetProjectProfile, "get-project-profile", "", false, "Get Project Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetRule, "get-rule", "", false, "Get Rule")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetSubscription, "get-subscription", "", false, "Get Subscription")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetSubscriptionGrant, "get-subscription-grant", "", false, "Get Subscription Grant")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetSubscriptionRequestDetails, "get-subscription-request-details", "", false, "Get Subscription Request Details")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetSubscriptionTarget, "get-subscription-target", "", false, "Get Subscription Target")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetTimeSeriesDataPoint, "get-time-series-data-point", "", false, "Get Time Series Data Point")
	_datazoneCmd.Flags().BoolVarP(&_datazoneGetUserProfile, "get-user-profile", "", false, "Get User Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListAccountPools, "list-account-pools", "", false, "List Account Pools")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListAccountsInAccountPool, "list-accounts-in-account-pool", "", false, "List Accounts In Account Pool")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListAssetFilters, "list-asset-filters", "", false, "List Asset Filters")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListAssetRevisions, "list-asset-revisions", "", false, "List Asset Revisions")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListConnections, "list-connections", "", false, "List Connections")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListDataProductRevisions, "list-data-product-revisions", "", false, "List Data Product Revisions")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListDataSourceRunActivities, "list-data-source-run-activities", "", false, "List Data Source Run Activities")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListDataSourceRuns, "list-data-source-runs", "", false, "List Data Source Runs")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListDataSources, "list-data-sources", "", false, "List Data Sources")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListDomainUnitsForParent, "list-domain-units-for-parent", "", false, "List Domain Units For Parent")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListDomains, "list-domains", "", false, "List Domains")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListEntityOwners, "list-entity-owners", "", false, "List Entity Owners")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListEnvironmentActions, "list-environment-actions", "", false, "List Environment Actions")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListEnvironmentBlueprintConfigurations, "list-environment-blueprint-configurations", "", false, "List Environment Blueprint Configurations")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListEnvironmentBlueprints, "list-environment-blueprints", "", false, "List Environment Blueprints")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListEnvironmentProfiles, "list-environment-profiles", "", false, "List Environment Profiles")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListEnvironments, "list-environments", "", false, "List Environments")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListJobRuns, "list-job-runs", "", false, "List Job Runs")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListLineageEvents, "list-lineage-events", "", false, "List Lineage Events")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListLineageNodeHistory, "list-lineage-node-history", "", false, "List Lineage Node History")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListMetadataGenerationRuns, "list-metadata-generation-runs", "", false, "List Metadata Generation Runs")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListNotifications, "list-notifications", "", false, "List Notifications")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListPolicyGrants, "list-policy-grants", "", false, "List Policy Grants")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListProjectMemberships, "list-project-memberships", "", false, "List Project Memberships")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListProjectProfiles, "list-project-profiles", "", false, "List Project Profiles")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListProjects, "list-projects", "", false, "List Projects")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListRules, "list-rules", "", false, "List Rules")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListSubscriptionGrants, "list-subscription-grants", "", false, "List Subscription Grants")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListSubscriptionRequests, "list-subscription-requests", "", false, "List Subscription Requests")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListSubscriptionTargets, "list-subscription-targets", "", false, "List Subscription Targets")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListSubscriptions, "list-subscriptions", "", false, "List Subscriptions")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_datazoneCmd.Flags().BoolVarP(&_datazoneListTimeSeriesDataPoints, "list-time-series-data-points", "", false, "List Time Series Data Points")
	_datazoneCmd.Flags().BoolVarP(&_datazonePostLineageEvent, "post-lineage-event", "", false, "Post Lineage Event")
	_datazoneCmd.Flags().BoolVarP(&_datazonePostTimeSeriesDataPoints, "post-time-series-data-points", "", false, "Post Time Series Data Points")
	_datazoneCmd.Flags().BoolVarP(&_datazonePutDataExportConfiguration, "put-data-export-configuration", "", false, "Put Data Export Configuration")
	_datazoneCmd.Flags().BoolVarP(&_datazonePutEnvironmentBlueprintConfiguration, "put-environment-blueprint-configuration", "", false, "Put Environment Blueprint Configuration")
	_datazoneCmd.Flags().BoolVarP(&_datazoneQueryGraph, "query-graph", "", false, "Query Graph")
	_datazoneCmd.Flags().BoolVarP(&_datazoneRejectPredictions, "reject-predictions", "", false, "Reject Predictions")
	_datazoneCmd.Flags().BoolVarP(&_datazoneRejectSubscriptionRequest, "reject-subscription-request", "", false, "Reject Subscription Request")
	_datazoneCmd.Flags().BoolVarP(&_datazoneRemoveEntityOwner, "remove-entity-owner", "", false, "Remove Entity Owner")
	_datazoneCmd.Flags().BoolVarP(&_datazoneRemovePolicyGrant, "remove-policy-grant", "", false, "Remove Policy Grant")
	_datazoneCmd.Flags().BoolVarP(&_datazoneRevokeSubscription, "revoke-subscription", "", false, "Revoke Subscription")
	_datazoneCmd.Flags().BoolVarP(&_datazoneSearch, "search", "", false, "Search")
	_datazoneCmd.Flags().BoolVarP(&_datazoneSearchGroupProfiles, "search-group-profiles", "", false, "Search Group Profiles")
	_datazoneCmd.Flags().BoolVarP(&_datazoneSearchListings, "search-listings", "", false, "Search Listings")
	_datazoneCmd.Flags().BoolVarP(&_datazoneSearchTypes, "search-types", "", false, "Search Types")
	_datazoneCmd.Flags().BoolVarP(&_datazoneSearchUserProfiles, "search-user-profiles", "", false, "Search User Profiles")
	_datazoneCmd.Flags().BoolVarP(&_datazoneStartDataSourceRun, "start-data-source-run", "", false, "Start Data Source Run")
	_datazoneCmd.Flags().BoolVarP(&_datazoneStartMetadataGenerationRun, "start-metadata-generation-run", "", false, "Start Metadata Generation Run")
	_datazoneCmd.Flags().BoolVarP(&_datazoneTagResource, "tag-resource", "", false, "Tag Resource")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUntagResource, "untag-resource", "", false, "Untag Resource")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateAccountPool, "update-account-pool", "", false, "Update Account Pool")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateAssetFilter, "update-asset-filter", "", false, "Update Asset Filter")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateConnection, "update-connection", "", false, "Update Connection")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateDataSource, "update-data-source", "", false, "Update Data Source")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateDomain, "update-domain", "", false, "Update Domain")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateDomainUnit, "update-domain-unit", "", false, "Update Domain Unit")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateEnvironment, "update-environment", "", false, "Update Environment")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateEnvironmentAction, "update-environment-action", "", false, "Update Environment Action")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateEnvironmentBlueprint, "update-environment-blueprint", "", false, "Update Environment Blueprint")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateEnvironmentProfile, "update-environment-profile", "", false, "Update Environment Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateGlossary, "update-glossary", "", false, "Update Glossary")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateGlossaryTerm, "update-glossary-term", "", false, "Update Glossary Term")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateGroupProfile, "update-group-profile", "", false, "Update Group Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateProject, "update-project", "", false, "Update Project")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateProjectProfile, "update-project-profile", "", false, "Update Project Profile")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateRootDomainUnitOwner, "update-root-domain-unit-owner", "", false, "Update Root Domain Unit Owner")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateRule, "update-rule", "", false, "Update Rule")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateSubscriptionGrantStatus, "update-subscription-grant-status", "", false, "Update Subscription Grant Status")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateSubscriptionRequest, "update-subscription-request", "", false, "Update Subscription Request")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateSubscriptionTarget, "update-subscription-target", "", false, "Update Subscription Target")
	_datazoneCmd.Flags().BoolVarP(&_datazoneUpdateUserProfile, "update-user-profile", "", false, "Update User Profile")

}
