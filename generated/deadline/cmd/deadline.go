package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/deadline"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deadlineCmd represents the deadline command
var _deadlineCmd = &cobra.Command{
	Use:   "deadline",
	Short: "AWS deadline CLI",
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
		client := deadline.NewFromConfig(cfg)
		if _deadlineAssociateMemberToFarm {
			deadline_AssociateMemberToFarm(cfg, client)
			return
		}
		if _deadlineAssociateMemberToFleet {
			deadline_AssociateMemberToFleet(cfg, client)
			return
		}
		if _deadlineAssociateMemberToJob {
			deadline_AssociateMemberToJob(cfg, client)
			return
		}
		if _deadlineAssociateMemberToQueue {
			deadline_AssociateMemberToQueue(cfg, client)
			return
		}
		if _deadlineAssumeFleetRoleForRead {
			deadline_AssumeFleetRoleForRead(cfg, client)
			return
		}
		if _deadlineAssumeFleetRoleForWorker {
			deadline_AssumeFleetRoleForWorker(cfg, client)
			return
		}
		if _deadlineAssumeQueueRoleForRead {
			deadline_AssumeQueueRoleForRead(cfg, client)
			return
		}
		if _deadlineAssumeQueueRoleForUser {
			deadline_AssumeQueueRoleForUser(cfg, client)
			return
		}
		if _deadlineAssumeQueueRoleForWorker {
			deadline_AssumeQueueRoleForWorker(cfg, client)
			return
		}
		if _deadlineBatchGetJobEntity {
			deadline_BatchGetJobEntity(cfg, client)
			return
		}
		if _deadlineCopyJobTemplate {
			deadline_CopyJobTemplate(cfg, client)
			return
		}
		if _deadlineCreateBudget {
			deadline_CreateBudget(cfg, client)
			return
		}
		if _deadlineCreateFarm {
			deadline_CreateFarm(cfg, client)
			return
		}
		if _deadlineCreateFleet {
			deadline_CreateFleet(cfg, client)
			return
		}
		if _deadlineCreateJob {
			deadline_CreateJob(cfg, client)
			return
		}
		if _deadlineCreateLicenseEndpoint {
			deadline_CreateLicenseEndpoint(cfg, client)
			return
		}
		if _deadlineCreateLimit {
			deadline_CreateLimit(cfg, client)
			return
		}
		if _deadlineCreateMonitor {
			deadline_CreateMonitor(cfg, client)
			return
		}
		if _deadlineCreateQueue {
			deadline_CreateQueue(cfg, client)
			return
		}
		if _deadlineCreateQueueEnvironment {
			deadline_CreateQueueEnvironment(cfg, client)
			return
		}
		if _deadlineCreateQueueFleetAssociation {
			deadline_CreateQueueFleetAssociation(cfg, client)
			return
		}
		if _deadlineCreateQueueLimitAssociation {
			deadline_CreateQueueLimitAssociation(cfg, client)
			return
		}
		if _deadlineCreateStorageProfile {
			deadline_CreateStorageProfile(cfg, client)
			return
		}
		if _deadlineCreateWorker {
			deadline_CreateWorker(cfg, client)
			return
		}
		if _deadlineDeleteBudget {
			deadline_DeleteBudget(cfg, client)
			return
		}
		if _deadlineDeleteFarm {
			deadline_DeleteFarm(cfg, client)
			return
		}
		if _deadlineDeleteFleet {
			deadline_DeleteFleet(cfg, client)
			return
		}
		if _deadlineDeleteLicenseEndpoint {
			deadline_DeleteLicenseEndpoint(cfg, client)
			return
		}
		if _deadlineDeleteLimit {
			deadline_DeleteLimit(cfg, client)
			return
		}
		if _deadlineDeleteMeteredProduct {
			deadline_DeleteMeteredProduct(cfg, client)
			return
		}
		if _deadlineDeleteMonitor {
			deadline_DeleteMonitor(cfg, client)
			return
		}
		if _deadlineDeleteQueue {
			deadline_DeleteQueue(cfg, client)
			return
		}
		if _deadlineDeleteQueueEnvironment {
			deadline_DeleteQueueEnvironment(cfg, client)
			return
		}
		if _deadlineDeleteQueueFleetAssociation {
			deadline_DeleteQueueFleetAssociation(cfg, client)
			return
		}
		if _deadlineDeleteQueueLimitAssociation {
			deadline_DeleteQueueLimitAssociation(cfg, client)
			return
		}
		if _deadlineDeleteStorageProfile {
			deadline_DeleteStorageProfile(cfg, client)
			return
		}
		if _deadlineDeleteWorker {
			deadline_DeleteWorker(cfg, client)
			return
		}
		if _deadlineDisassociateMemberFromFarm {
			deadline_DisassociateMemberFromFarm(cfg, client)
			return
		}
		if _deadlineDisassociateMemberFromFleet {
			deadline_DisassociateMemberFromFleet(cfg, client)
			return
		}
		if _deadlineDisassociateMemberFromJob {
			deadline_DisassociateMemberFromJob(cfg, client)
			return
		}
		if _deadlineDisassociateMemberFromQueue {
			deadline_DisassociateMemberFromQueue(cfg, client)
			return
		}
		if _deadlineGetBudget {
			deadline_GetBudget(cfg, client)
			return
		}
		if _deadlineGetFarm {
			deadline_GetFarm(cfg, client)
			return
		}
		if _deadlineGetFleet {
			deadline_GetFleet(cfg, client)
			return
		}
		if _deadlineGetJob {
			deadline_GetJob(cfg, client)
			return
		}
		if _deadlineGetLicenseEndpoint {
			deadline_GetLicenseEndpoint(cfg, client)
			return
		}
		if _deadlineGetLimit {
			deadline_GetLimit(cfg, client)
			return
		}
		if _deadlineGetMonitor {
			deadline_GetMonitor(cfg, client)
			return
		}
		if _deadlineGetQueue {
			deadline_GetQueue(cfg, client)
			return
		}
		if _deadlineGetQueueEnvironment {
			deadline_GetQueueEnvironment(cfg, client)
			return
		}
		if _deadlineGetQueueFleetAssociation {
			deadline_GetQueueFleetAssociation(cfg, client)
			return
		}
		if _deadlineGetQueueLimitAssociation {
			deadline_GetQueueLimitAssociation(cfg, client)
			return
		}
		if _deadlineGetSession {
			deadline_GetSession(cfg, client)
			return
		}
		if _deadlineGetSessionAction {
			deadline_GetSessionAction(cfg, client)
			return
		}
		if _deadlineGetSessionsStatisticsAggregation {
			deadline_GetSessionsStatisticsAggregation(cfg, client)
			return
		}
		if _deadlineGetStep {
			deadline_GetStep(cfg, client)
			return
		}
		if _deadlineGetStorageProfile {
			deadline_GetStorageProfile(cfg, client)
			return
		}
		if _deadlineGetStorageProfileForQueue {
			deadline_GetStorageProfileForQueue(cfg, client)
			return
		}
		if _deadlineGetTask {
			deadline_GetTask(cfg, client)
			return
		}
		if _deadlineGetWorker {
			deadline_GetWorker(cfg, client)
			return
		}
		if _deadlineListAvailableMeteredProducts {
			deadline_ListAvailableMeteredProducts(cfg, client)
			return
		}
		if _deadlineListBudgets {
			deadline_ListBudgets(cfg, client)
			return
		}
		if _deadlineListFarmMembers {
			deadline_ListFarmMembers(cfg, client)
			return
		}
		if _deadlineListFarms {
			deadline_ListFarms(cfg, client)
			return
		}
		if _deadlineListFleetMembers {
			deadline_ListFleetMembers(cfg, client)
			return
		}
		if _deadlineListFleets {
			deadline_ListFleets(cfg, client)
			return
		}
		if _deadlineListJobMembers {
			deadline_ListJobMembers(cfg, client)
			return
		}
		if _deadlineListJobParameterDefinitions {
			deadline_ListJobParameterDefinitions(cfg, client)
			return
		}
		if _deadlineListJobs {
			deadline_ListJobs(cfg, client)
			return
		}
		if _deadlineListLicenseEndpoints {
			deadline_ListLicenseEndpoints(cfg, client)
			return
		}
		if _deadlineListLimits {
			deadline_ListLimits(cfg, client)
			return
		}
		if _deadlineListMeteredProducts {
			deadline_ListMeteredProducts(cfg, client)
			return
		}
		if _deadlineListMonitors {
			deadline_ListMonitors(cfg, client)
			return
		}
		if _deadlineListQueueEnvironments {
			deadline_ListQueueEnvironments(cfg, client)
			return
		}
		if _deadlineListQueueFleetAssociations {
			deadline_ListQueueFleetAssociations(cfg, client)
			return
		}
		if _deadlineListQueueLimitAssociations {
			deadline_ListQueueLimitAssociations(cfg, client)
			return
		}
		if _deadlineListQueueMembers {
			deadline_ListQueueMembers(cfg, client)
			return
		}
		if _deadlineListQueues {
			deadline_ListQueues(cfg, client)
			return
		}
		if _deadlineListSessionActions {
			deadline_ListSessionActions(cfg, client)
			return
		}
		if _deadlineListSessions {
			deadline_ListSessions(cfg, client)
			return
		}
		if _deadlineListSessionsForWorker {
			deadline_ListSessionsForWorker(cfg, client)
			return
		}
		if _deadlineListStepConsumers {
			deadline_ListStepConsumers(cfg, client)
			return
		}
		if _deadlineListStepDependencies {
			deadline_ListStepDependencies(cfg, client)
			return
		}
		if _deadlineListSteps {
			deadline_ListSteps(cfg, client)
			return
		}
		if _deadlineListStorageProfiles {
			deadline_ListStorageProfiles(cfg, client)
			return
		}
		if _deadlineListStorageProfilesForQueue {
			deadline_ListStorageProfilesForQueue(cfg, client)
			return
		}
		if _deadlineListTagsForResource {
			deadline_ListTagsForResource(cfg, client)
			return
		}
		if _deadlineListTasks {
			deadline_ListTasks(cfg, client)
			return
		}
		if _deadlineListWorkers {
			deadline_ListWorkers(cfg, client)
			return
		}
		if _deadlinePutMeteredProduct {
			deadline_PutMeteredProduct(cfg, client)
			return
		}
		if _deadlineSearchJobs {
			deadline_SearchJobs(cfg, client)
			return
		}
		if _deadlineSearchSteps {
			deadline_SearchSteps(cfg, client)
			return
		}
		if _deadlineSearchTasks {
			deadline_SearchTasks(cfg, client)
			return
		}
		if _deadlineSearchWorkers {
			deadline_SearchWorkers(cfg, client)
			return
		}
		if _deadlineStartSessionsStatisticsAggregation {
			deadline_StartSessionsStatisticsAggregation(cfg, client)
			return
		}
		if _deadlineTagResource {
			deadline_TagResource(cfg, client)
			return
		}
		if _deadlineUntagResource {
			deadline_UntagResource(cfg, client)
			return
		}
		if _deadlineUpdateBudget {
			deadline_UpdateBudget(cfg, client)
			return
		}
		if _deadlineUpdateFarm {
			deadline_UpdateFarm(cfg, client)
			return
		}
		if _deadlineUpdateFleet {
			deadline_UpdateFleet(cfg, client)
			return
		}
		if _deadlineUpdateJob {
			deadline_UpdateJob(cfg, client)
			return
		}
		if _deadlineUpdateLimit {
			deadline_UpdateLimit(cfg, client)
			return
		}
		if _deadlineUpdateMonitor {
			deadline_UpdateMonitor(cfg, client)
			return
		}
		if _deadlineUpdateQueue {
			deadline_UpdateQueue(cfg, client)
			return
		}
		if _deadlineUpdateQueueEnvironment {
			deadline_UpdateQueueEnvironment(cfg, client)
			return
		}
		if _deadlineUpdateQueueFleetAssociation {
			deadline_UpdateQueueFleetAssociation(cfg, client)
			return
		}
		if _deadlineUpdateQueueLimitAssociation {
			deadline_UpdateQueueLimitAssociation(cfg, client)
			return
		}
		if _deadlineUpdateSession {
			deadline_UpdateSession(cfg, client)
			return
		}
		if _deadlineUpdateStep {
			deadline_UpdateStep(cfg, client)
			return
		}
		if _deadlineUpdateStorageProfile {
			deadline_UpdateStorageProfile(cfg, client)
			return
		}
		if _deadlineUpdateTask {
			deadline_UpdateTask(cfg, client)
			return
		}
		if _deadlineUpdateWorker {
			deadline_UpdateWorker(cfg, client)
			return
		}
		if _deadlineUpdateWorkerSchedule {
			deadline_UpdateWorkerSchedule(cfg, client)
			return
		}

	},
}

var (
	_deadlineAssociateMemberToFarm              bool
	_deadlineAssociateMemberToFleet             bool
	_deadlineAssociateMemberToJob               bool
	_deadlineAssociateMemberToQueue             bool
	_deadlineAssumeFleetRoleForRead             bool
	_deadlineAssumeFleetRoleForWorker           bool
	_deadlineAssumeQueueRoleForRead             bool
	_deadlineAssumeQueueRoleForUser             bool
	_deadlineAssumeQueueRoleForWorker           bool
	_deadlineBatchGetJobEntity                  bool
	_deadlineCopyJobTemplate                    bool
	_deadlineCreateBudget                       bool
	_deadlineCreateFarm                         bool
	_deadlineCreateFleet                        bool
	_deadlineCreateJob                          bool
	_deadlineCreateLicenseEndpoint              bool
	_deadlineCreateLimit                        bool
	_deadlineCreateMonitor                      bool
	_deadlineCreateQueue                        bool
	_deadlineCreateQueueEnvironment             bool
	_deadlineCreateQueueFleetAssociation        bool
	_deadlineCreateQueueLimitAssociation        bool
	_deadlineCreateStorageProfile               bool
	_deadlineCreateWorker                       bool
	_deadlineDeleteBudget                       bool
	_deadlineDeleteFarm                         bool
	_deadlineDeleteFleet                        bool
	_deadlineDeleteLicenseEndpoint              bool
	_deadlineDeleteLimit                        bool
	_deadlineDeleteMeteredProduct               bool
	_deadlineDeleteMonitor                      bool
	_deadlineDeleteQueue                        bool
	_deadlineDeleteQueueEnvironment             bool
	_deadlineDeleteQueueFleetAssociation        bool
	_deadlineDeleteQueueLimitAssociation        bool
	_deadlineDeleteStorageProfile               bool
	_deadlineDeleteWorker                       bool
	_deadlineDisassociateMemberFromFarm         bool
	_deadlineDisassociateMemberFromFleet        bool
	_deadlineDisassociateMemberFromJob          bool
	_deadlineDisassociateMemberFromQueue        bool
	_deadlineGetBudget                          bool
	_deadlineGetFarm                            bool
	_deadlineGetFleet                           bool
	_deadlineGetJob                             bool
	_deadlineGetLicenseEndpoint                 bool
	_deadlineGetLimit                           bool
	_deadlineGetMonitor                         bool
	_deadlineGetQueue                           bool
	_deadlineGetQueueEnvironment                bool
	_deadlineGetQueueFleetAssociation           bool
	_deadlineGetQueueLimitAssociation           bool
	_deadlineGetSession                         bool
	_deadlineGetSessionAction                   bool
	_deadlineGetSessionsStatisticsAggregation   bool
	_deadlineGetStep                            bool
	_deadlineGetStorageProfile                  bool
	_deadlineGetStorageProfileForQueue          bool
	_deadlineGetTask                            bool
	_deadlineGetWorker                          bool
	_deadlineListAvailableMeteredProducts       bool
	_deadlineListBudgets                        bool
	_deadlineListFarmMembers                    bool
	_deadlineListFarms                          bool
	_deadlineListFleetMembers                   bool
	_deadlineListFleets                         bool
	_deadlineListJobMembers                     bool
	_deadlineListJobParameterDefinitions        bool
	_deadlineListJobs                           bool
	_deadlineListLicenseEndpoints               bool
	_deadlineListLimits                         bool
	_deadlineListMeteredProducts                bool
	_deadlineListMonitors                       bool
	_deadlineListQueueEnvironments              bool
	_deadlineListQueueFleetAssociations         bool
	_deadlineListQueueLimitAssociations         bool
	_deadlineListQueueMembers                   bool
	_deadlineListQueues                         bool
	_deadlineListSessionActions                 bool
	_deadlineListSessions                       bool
	_deadlineListSessionsForWorker              bool
	_deadlineListStepConsumers                  bool
	_deadlineListStepDependencies               bool
	_deadlineListSteps                          bool
	_deadlineListStorageProfiles                bool
	_deadlineListStorageProfilesForQueue        bool
	_deadlineListTagsForResource                bool
	_deadlineListTasks                          bool
	_deadlineListWorkers                        bool
	_deadlinePutMeteredProduct                  bool
	_deadlineSearchJobs                         bool
	_deadlineSearchSteps                        bool
	_deadlineSearchTasks                        bool
	_deadlineSearchWorkers                      bool
	_deadlineStartSessionsStatisticsAggregation bool
	_deadlineTagResource                        bool
	_deadlineUntagResource                      bool
	_deadlineUpdateBudget                       bool
	_deadlineUpdateFarm                         bool
	_deadlineUpdateFleet                        bool
	_deadlineUpdateJob                          bool
	_deadlineUpdateLimit                        bool
	_deadlineUpdateMonitor                      bool
	_deadlineUpdateQueue                        bool
	_deadlineUpdateQueueEnvironment             bool
	_deadlineUpdateQueueFleetAssociation        bool
	_deadlineUpdateQueueLimitAssociation        bool
	_deadlineUpdateSession                      bool
	_deadlineUpdateStep                         bool
	_deadlineUpdateStorageProfile               bool
	_deadlineUpdateTask                         bool
	_deadlineUpdateWorker                       bool
	_deadlineUpdateWorkerSchedule               bool

	_deadlineActions                                 string
	_deadlineActionsToAdd                            string
	_deadlineActionsToRemove                         string
	_deadlineAggregationId                           string
	_deadlineAllowedStorageProfileIds                []string
	_deadlineAllowedStorageProfileIdsToAdd           []string
	_deadlineAllowedStorageProfileIdsToRemove        []string
	_deadlineAmountRequirementName                   string
	_deadlineApproximateDollarLimit                  string
	_deadlineAttachments                             string
	_deadlineBudgetId                                string
	_deadlineCapabilities                            string
	_deadlineClientToken                             string
	_deadlineConfiguration                           string
	_deadlineDefaultBudgetAction                     string
	_deadlineDescription                             string
	_deadlineDescriptionOverride                     string
	_deadlineDisplayName                             string
	_deadlineEndTime                                 string
	_deadlineFarmId                                  string
	_deadlineFileSystemLocations                     string
	_deadlineFileSystemLocationsToAdd                string
	_deadlineFileSystemLocationsToRemove             string
	_deadlineFilterExpressions                       string
	_deadlineFleetId                                 string
	_deadlineFleetIds                                []string
	_deadlineGroupBy                                 string
	_deadlineHostConfiguration                       string
	_deadlineHostProperties                          string
	_deadlineIdentifiers                             string
	_deadlineIdentityCenterInstanceArn               string
	_deadlineIdentityStoreId                         string
	_deadlineItemOffset                              string
	_deadlineJobAttachmentSettings                   string
	_deadlineJobId                                   string
	_deadlineJobRunAsUser                            string
	_deadlineKmsKeyArn                               string
	_deadlineLicenseEndpointId                       string
	_deadlineLifecycleStatus                         string
	_deadlineLimitId                                 string
	_deadlineMaxCount                                string
	_deadlineMaxFailedTasksCount                     string
	_deadlineMaxResults                              string
	_deadlineMaxRetriesPerTask                       string
	_deadlineMaxWorkerCount                          string
	_deadlineMembershipLevel                         string
	_deadlineMinWorkerCount                          string
	_deadlineMonitorId                               string
	_deadlineName                                    string
	_deadlineNameOverride                            string
	_deadlineNextToken                               string
	_deadlineOsFamily                                string
	_deadlinePageSize                                string
	_deadlineParameters                              string
	_deadlinePeriod                                  string
	_deadlinePrincipalId                             string
	_deadlinePrincipalType                           string
	_deadlinePriority                                string
	_deadlineProductId                               string
	_deadlineQueueEnvironmentId                      string
	_deadlineQueueId                                 string
	_deadlineQueueIds                                []string
	_deadlineRequiredFileSystemLocationNames         []string
	_deadlineRequiredFileSystemLocationNamesToAdd    []string
	_deadlineRequiredFileSystemLocationNamesToRemove []string
	_deadlineResourceArn                             string
	_deadlineResourceIds                             string
	_deadlineRoleArn                                 string
	_deadlineSchedule                                string
	_deadlineSecurityGroupIds                        []string
	_deadlineSessionActionId                         string
	_deadlineSessionId                               string
	_deadlineSortExpressions                         string
	_deadlineSourceJobId                             string
	_deadlineStartTime                               string
	_deadlineStatistics                              string
	_deadlineStatus                                  string
	_deadlineStepId                                  string
	_deadlineStorageProfileId                        string
	_deadlineSubdomain                               string
	_deadlineSubnetIds                               []string
	_deadlineTagKeys                                 []string
	_deadlineTags                                    string
	_deadlineTargetLifecycleStatus                   string
	_deadlineTargetRunStatus                         string
	_deadlineTargetS3Location                        string
	_deadlineTargetTaskRunStatus                     string
	_deadlineTaskId                                  string
	_deadlineTemplate                                string
	_deadlineTemplateType                            string
	_deadlineTimezone                                string
	_deadlineUpdatedSessionActions                   string
	_deadlineUsageTrackingResource                   string
	_deadlineVpcId                                   string
	_deadlineWorkerId                                string
)

// Assigns a farm membership level to a member.
func deadline_AssociateMemberToFarm(cfg aws.Config, client *deadline.Client) {
	input := &deadline.AssociateMemberToFarmInput{
		// FarmId: *string, // Required
		// IdentityStoreId: *string, // Required
		// MembershipLevel: types.MembershipLevel, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.DeadlinePrincipalType, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_deadlineIdentityStoreId)
	}
	if len(_deadlineMembershipLevel) > 0 {
		if err := assignInputField(input, "MembershipLevel", _deadlineMembershipLevel); err != nil {
			log.Errorf("invalid --membership-level: %s", err.Error())
			return
		}
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}
	if len(_deadlinePrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _deadlinePrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateMemberToFarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns a fleet membership level to a member.
func deadline_AssociateMemberToFleet(cfg aws.Config, client *deadline.Client) {
	input := &deadline.AssociateMemberToFleetInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// IdentityStoreId: *string, // Required
		// MembershipLevel: types.MembershipLevel, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.DeadlinePrincipalType, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_deadlineIdentityStoreId)
	}
	if len(_deadlineMembershipLevel) > 0 {
		if err := assignInputField(input, "MembershipLevel", _deadlineMembershipLevel); err != nil {
			log.Errorf("invalid --membership-level: %s", err.Error())
			return
		}
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}
	if len(_deadlinePrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _deadlinePrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateMemberToFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns a job membership level to a member
func deadline_AssociateMemberToJob(cfg aws.Config, client *deadline.Client) {
	input := &deadline.AssociateMemberToJobInput{
		// FarmId: *string, // Required
		// IdentityStoreId: *string, // Required
		// JobId: *string, // Required
		// MembershipLevel: types.MembershipLevel, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.DeadlinePrincipalType, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_deadlineIdentityStoreId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineMembershipLevel) > 0 {
		if err := assignInputField(input, "MembershipLevel", _deadlineMembershipLevel); err != nil {
			log.Errorf("invalid --membership-level: %s", err.Error())
			return
		}
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}
	if len(_deadlinePrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _deadlinePrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.AssociateMemberToJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns a queue membership level to a member
func deadline_AssociateMemberToQueue(cfg aws.Config, client *deadline.Client) {
	input := &deadline.AssociateMemberToQueueInput{
		// FarmId: *string, // Required
		// IdentityStoreId: *string, // Required
		// MembershipLevel: types.MembershipLevel, // Required
		// PrincipalId: *string, // Required
		// PrincipalType: types.DeadlinePrincipalType, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_deadlineIdentityStoreId)
	}
	if len(_deadlineMembershipLevel) > 0 {
		if err := assignInputField(input, "MembershipLevel", _deadlineMembershipLevel); err != nil {
			log.Errorf("invalid --membership-level: %s", err.Error())
			return
		}
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}
	if len(_deadlinePrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _deadlinePrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.AssociateMemberToQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get Amazon Web Services credentials from the fleet role. The IAM permissions of
// the credentials are scoped down to have read-only access.
func deadline_AssumeFleetRoleForRead(cfg aws.Config, client *deadline.Client) {
	input := &deadline.AssumeFleetRoleForReadInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}

	if resp, err := client.AssumeFleetRoleForRead(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get credentials from the fleet role for a worker.
func deadline_AssumeFleetRoleForWorker(cfg aws.Config, client *deadline.Client) {
	input := &deadline.AssumeFleetRoleForWorkerInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineWorkerId) > 0 {
		input.WorkerId = aws.String(_deadlineWorkerId)
	}

	if resp, err := client.AssumeFleetRoleForWorker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets Amazon Web Services credentials from the queue role. The IAM permissions
// of the credentials are scoped down to have read-only access.
func deadline_AssumeQueueRoleForRead(cfg aws.Config, client *deadline.Client) {
	input := &deadline.AssumeQueueRoleForReadInput{
		// FarmId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.AssumeQueueRoleForRead(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows a user to assume a role for a queue.
func deadline_AssumeQueueRoleForUser(cfg aws.Config, client *deadline.Client) {
	input := &deadline.AssumeQueueRoleForUserInput{
		// FarmId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.AssumeQueueRoleForUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows a worker to assume a queue role.
func deadline_AssumeQueueRoleForWorker(cfg aws.Config, client *deadline.Client) {
	input := &deadline.AssumeQueueRoleForWorkerInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// QueueId: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineWorkerId) > 0 {
		input.WorkerId = aws.String(_deadlineWorkerId)
	}

	if resp, err := client.AssumeQueueRoleForWorker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get batched job details for a worker.
func deadline_BatchGetJobEntity(cfg aws.Config, client *deadline.Client) {
	input := &deadline.BatchGetJobEntityInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// Identifiers: []types.JobEntityIdentifiersUnion, // Required
		// WorkerId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineIdentifiers) > 0 {
		if err := assignInputField(input, "Identifiers", _deadlineIdentifiers); err != nil {
			log.Errorf("invalid --identifiers: %s", err.Error())
			return
		}
	}
	if len(_deadlineWorkerId) > 0 {
		input.WorkerId = aws.String(_deadlineWorkerId)
	}

	if resp, err := client.BatchGetJobEntity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies a job template to an Amazon S3 bucket.
func deadline_CopyJobTemplate(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CopyJobTemplateInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// TargetS3Location: *types.S3Location, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineTargetS3Location) > 0 {
		if err := assignInputField(input, "TargetS3Location", _deadlineTargetS3Location); err != nil {
			log.Errorf("invalid --target-s3-location: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyJobTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a budget to set spending thresholds for your rendering activity.
func deadline_CreateBudget(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateBudgetInput{
		// Actions: []types.BudgetActionToAdd, // Required
		// ApproximateDollarLimit: *float32, // Required
		// DisplayName: *string, // Required
		// FarmId: *string, // Required
		// Schedule: types.BudgetSchedule, // Required
		// UsageTrackingResource: types.UsageTrackingResource, // Required
	}

	if len(_deadlineActions) > 0 {
		if err := assignInputField(input, "Actions", _deadlineActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_deadlineApproximateDollarLimit) > 0 {
		if err := assignInputField(input, "ApproximateDollarLimit", _deadlineApproximateDollarLimit); err != nil {
			log.Errorf("invalid --approximate-dollar-limit: %s", err.Error())
			return
		}
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _deadlineSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_deadlineUsageTrackingResource) > 0 {
		if err := assignInputField(input, "UsageTrackingResource", _deadlineUsageTrackingResource); err != nil {
			log.Errorf("invalid --usage-tracking-resource: %s", err.Error())
			return
		}
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}
	if len(_deadlineTags) > 0 {
		if err := assignInputField(input, "Tags", _deadlineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBudget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a farm to allow space for queues and fleets. Farms are the space where
// the components of your renders gather and are pieced together in the cloud.
// Farms contain budgets and allow you to enforce permissions. Deadline Cloud farms
// are a useful container for large projects.
func deadline_CreateFarm(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateFarmInput{
		// DisplayName: *string, // Required
	}

	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}
	if len(_deadlineKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_deadlineKmsKeyArn)
	}
	if len(_deadlineTags) > 0 {
		if err := assignInputField(input, "Tags", _deadlineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a fleet. Fleets gather information relating to compute, or capacity,
// for renders within your farms. You can choose to manage your own capacity or opt
// to have fleets fully managed by Deadline Cloud.
func deadline_CreateFleet(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateFleetInput{
		// Configuration: types.FleetConfiguration, // Required
		// DisplayName: *string, // Required
		// FarmId: *string, // Required
		// MaxWorkerCount: *int32, // Required
		// RoleArn: *string, // Required
	}

	if len(_deadlineConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _deadlineConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineMaxWorkerCount) > 0 {
		if err := assignInputField(input, "MaxWorkerCount", _deadlineMaxWorkerCount); err != nil {
			log.Errorf("invalid --max-worker-count: %s", err.Error())
			return
		}
	}
	if len(_deadlineRoleArn) > 0 {
		input.RoleArn = aws.String(_deadlineRoleArn)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}
	if len(_deadlineHostConfiguration) > 0 {
		if err := assignInputField(input, "HostConfiguration", _deadlineHostConfiguration); err != nil {
			log.Errorf("invalid --host-configuration: %s", err.Error())
			return
		}
	}
	if len(_deadlineMinWorkerCount) > 0 {
		if err := assignInputField(input, "MinWorkerCount", _deadlineMinWorkerCount); err != nil {
			log.Errorf("invalid --min-worker-count: %s", err.Error())
			return
		}
	}
	if len(_deadlineTags) > 0 {
		if err := assignInputField(input, "Tags", _deadlineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a job. A job is a set of instructions that Deadline Cloud uses to
// schedule and run work on available workers. For more information, see [Deadline Cloud jobs].
//
// [Deadline Cloud jobs]: https://docs.aws.amazon.com/deadline-cloud/latest/userguide/deadline-cloud-jobs.html
func deadline_CreateJob(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateJobInput{
		// FarmId: *string, // Required
		// Priority: *int32, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlinePriority) > 0 {
		if err := assignInputField(input, "Priority", _deadlinePriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineAttachments) > 0 {
		if err := assignInputField(input, "Attachments", _deadlineAttachments); err != nil {
			log.Errorf("invalid --attachments: %s", err.Error())
			return
		}
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineDescriptionOverride) > 0 {
		input.DescriptionOverride = aws.String(_deadlineDescriptionOverride)
	}
	if len(_deadlineMaxFailedTasksCount) > 0 {
		if err := assignInputField(input, "MaxFailedTasksCount", _deadlineMaxFailedTasksCount); err != nil {
			log.Errorf("invalid --max-failed-tasks-count: %s", err.Error())
			return
		}
	}
	if len(_deadlineMaxRetriesPerTask) > 0 {
		if err := assignInputField(input, "MaxRetriesPerTask", _deadlineMaxRetriesPerTask); err != nil {
			log.Errorf("invalid --max-retries-per-task: %s", err.Error())
			return
		}
	}
	if len(_deadlineMaxWorkerCount) > 0 {
		if err := assignInputField(input, "MaxWorkerCount", _deadlineMaxWorkerCount); err != nil {
			log.Errorf("invalid --max-worker-count: %s", err.Error())
			return
		}
	}
	if len(_deadlineNameOverride) > 0 {
		input.NameOverride = aws.String(_deadlineNameOverride)
	}
	if len(_deadlineParameters) > 0 {
		if err := assignInputField(input, "Parameters", _deadlineParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_deadlineSourceJobId) > 0 {
		input.SourceJobId = aws.String(_deadlineSourceJobId)
	}
	if len(_deadlineStorageProfileId) > 0 {
		input.StorageProfileId = aws.String(_deadlineStorageProfileId)
	}
	if len(_deadlineTags) > 0 {
		if err := assignInputField(input, "Tags", _deadlineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_deadlineTargetTaskRunStatus) > 0 {
		if err := assignInputField(input, "TargetTaskRunStatus", _deadlineTargetTaskRunStatus); err != nil {
			log.Errorf("invalid --target-task-run-status: %s", err.Error())
			return
		}
	}
	if len(_deadlineTemplate) > 0 {
		input.Template = aws.String(_deadlineTemplate)
	}
	if len(_deadlineTemplateType) > 0 {
		if err := assignInputField(input, "TemplateType", _deadlineTemplateType); err != nil {
			log.Errorf("invalid --template-type: %s", err.Error())
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

// Creates a license endpoint to integrate your various licensed software used for
// rendering on Deadline Cloud.
func deadline_CreateLicenseEndpoint(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateLicenseEndpointInput{
		// SecurityGroupIds: []string, // Required
		// SubnetIds: []string, // Required
		// VpcId: *string, // Required
	}

	if len(_deadlineSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _deadlineSecurityGroupIds...)
	}
	if len(_deadlineSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _deadlineSubnetIds...)
	}
	if len(_deadlineVpcId) > 0 {
		input.VpcId = aws.String(_deadlineVpcId)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineTags) > 0 {
		if err := assignInputField(input, "Tags", _deadlineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLicenseEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a limit that manages the distribution of shared resources, such as
// floating licenses. A limit can throttle work assignments, help manage workloads,
// and track current usage. Before you use a limit, you must associate the limit
// with one or more queues.
//
// You must add the amountRequirementName to a step in a job template to declare
// the limit requirement.
func deadline_CreateLimit(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateLimitInput{
		// AmountRequirementName: *string, // Required
		// DisplayName: *string, // Required
		// FarmId: *string, // Required
		// MaxCount: *int32, // Required
	}

	if len(_deadlineAmountRequirementName) > 0 {
		input.AmountRequirementName = aws.String(_deadlineAmountRequirementName)
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineMaxCount) > 0 {
		if err := assignInputField(input, "MaxCount", _deadlineMaxCount); err != nil {
			log.Errorf("invalid --max-count: %s", err.Error())
			return
		}
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}

	if resp, err := client.CreateLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Web Services Deadline Cloud monitor that you can use to view
// your farms, queues, and fleets. After you submit a job, you can track the
// progress of the tasks and steps that make up the job, and then download the
// job's results.
func deadline_CreateMonitor(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateMonitorInput{
		// DisplayName: *string, // Required
		// IdentityCenterInstanceArn: *string, // Required
		// RoleArn: *string, // Required
		// Subdomain: *string, // Required
	}

	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineIdentityCenterInstanceArn) > 0 {
		input.IdentityCenterInstanceArn = aws.String(_deadlineIdentityCenterInstanceArn)
	}
	if len(_deadlineRoleArn) > 0 {
		input.RoleArn = aws.String(_deadlineRoleArn)
	}
	if len(_deadlineSubdomain) > 0 {
		input.Subdomain = aws.String(_deadlineSubdomain)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineTags) > 0 {
		if err := assignInputField(input, "Tags", _deadlineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a queue to coordinate the order in which jobs run on a farm. A queue
// can also specify where to pull resources and indicate where to output completed
// jobs.
func deadline_CreateQueue(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateQueueInput{
		// DisplayName: *string, // Required
		// FarmId: *string, // Required
	}

	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineAllowedStorageProfileIds) > 0 {
		input.AllowedStorageProfileIds = append([]string(nil), _deadlineAllowedStorageProfileIds...)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineDefaultBudgetAction) > 0 {
		if err := assignInputField(input, "DefaultBudgetAction", _deadlineDefaultBudgetAction); err != nil {
			log.Errorf("invalid --default-budget-action: %s", err.Error())
			return
		}
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}
	if len(_deadlineJobAttachmentSettings) > 0 {
		if err := assignInputField(input, "JobAttachmentSettings", _deadlineJobAttachmentSettings); err != nil {
			log.Errorf("invalid --job-attachment-settings: %s", err.Error())
			return
		}
	}
	if len(_deadlineJobRunAsUser) > 0 {
		if err := assignInputField(input, "JobRunAsUser", _deadlineJobRunAsUser); err != nil {
			log.Errorf("invalid --job-run-as-user: %s", err.Error())
			return
		}
	}
	if len(_deadlineRequiredFileSystemLocationNames) > 0 {
		input.RequiredFileSystemLocationNames = append([]string(nil), _deadlineRequiredFileSystemLocationNames...)
	}
	if len(_deadlineRoleArn) > 0 {
		input.RoleArn = aws.String(_deadlineRoleArn)
	}
	if len(_deadlineTags) > 0 {
		if err := assignInputField(input, "Tags", _deadlineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an environment for a queue that defines how jobs in the queue run.
func deadline_CreateQueueEnvironment(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateQueueEnvironmentInput{
		// FarmId: *string, // Required
		// Priority: *int32, // Required
		// QueueId: *string, // Required
		// Template: *string, // Required
		// TemplateType: types.EnvironmentTemplateType, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlinePriority) > 0 {
		if err := assignInputField(input, "Priority", _deadlinePriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineTemplate) > 0 {
		input.Template = aws.String(_deadlineTemplate)
	}
	if len(_deadlineTemplateType) > 0 {
		if err := assignInputField(input, "TemplateType", _deadlineTemplateType); err != nil {
			log.Errorf("invalid --template-type: %s", err.Error())
			return
		}
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}

	if resp, err := client.CreateQueueEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an association between a queue and a fleet.
func deadline_CreateQueueFleetAssociation(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateQueueFleetAssociationInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.CreateQueueFleetAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a limit with a particular queue. After the limit is associated, all
// workers for jobs that specify the limit associated with the queue are subject to
// the limit. You can't associate two limits with the same amountRequirementName
// to the same queue.
func deadline_CreateQueueLimitAssociation(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateQueueLimitAssociationInput{
		// FarmId: *string, // Required
		// LimitId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineLimitId) > 0 {
		input.LimitId = aws.String(_deadlineLimitId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.CreateQueueLimitAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a storage profile that specifies the operating system, file type, and
// file location of resources used on a farm.
func deadline_CreateStorageProfile(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateStorageProfileInput{
		// DisplayName: *string, // Required
		// FarmId: *string, // Required
		// OsFamily: types.StorageProfileOperatingSystemFamily, // Required
	}

	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineOsFamily) > 0 {
		if err := assignInputField(input, "OsFamily", _deadlineOsFamily); err != nil {
			log.Errorf("invalid --os-family: %s", err.Error())
			return
		}
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineFileSystemLocations) > 0 {
		if err := assignInputField(input, "FileSystemLocations", _deadlineFileSystemLocations); err != nil {
			log.Errorf("invalid --file-system-locations: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStorageProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a worker. A worker tells your instance how much processing power
// (vCPU), and memory (GiB) you’ll need to assemble the digital assets held within
// a particular instance. You can specify certain instance types to use, or let the
// worker know which instances types to exclude.
//
// Deadline Cloud limits the number of workers to less than or equal to the
// fleet's maximum worker count. The service maintains eventual consistency for the
// worker count. If you make multiple rapid calls to CreateWorker before the field
// updates, you might exceed your fleet's maximum worker count. For example, if
// your maxWorkerCount is 10 and you currently have 9 workers, making two quick
// CreateWorker calls might successfully create 2 workers instead of 1, resulting
// in 11 total workers.
func deadline_CreateWorker(cfg aws.Config, client *deadline.Client) {
	input := &deadline.CreateWorkerInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineHostProperties) > 0 {
		if err := assignInputField(input, "HostProperties", _deadlineHostProperties); err != nil {
			log.Errorf("invalid --host-properties: %s", err.Error())
			return
		}
	}
	if len(_deadlineTags) > 0 {
		if err := assignInputField(input, "Tags", _deadlineTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a budget.
func deadline_DeleteBudget(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteBudgetInput{
		// BudgetId: *string, // Required
		// FarmId: *string, // Required
	}

	if len(_deadlineBudgetId) > 0 {
		input.BudgetId = aws.String(_deadlineBudgetId)
	}
	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}

	if resp, err := client.DeleteBudget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a farm.
func deadline_DeleteFarm(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteFarmInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}

	if resp, err := client.DeleteFarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a fleet.
func deadline_DeleteFleet(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteFleetInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}

	if resp, err := client.DeleteFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a license endpoint.
func deadline_DeleteLicenseEndpoint(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteLicenseEndpointInput{
		// LicenseEndpointId: *string, // Required
	}

	if len(_deadlineLicenseEndpointId) > 0 {
		input.LicenseEndpointId = aws.String(_deadlineLicenseEndpointId)
	}

	if resp, err := client.DeleteLicenseEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a limit from the specified farm. Before you delete a limit you must use
// the DeleteQueueLimitAssociation operation to remove the association with any
// queues.
func deadline_DeleteLimit(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteLimitInput{
		// FarmId: *string, // Required
		// LimitId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineLimitId) > 0 {
		input.LimitId = aws.String(_deadlineLimitId)
	}

	if resp, err := client.DeleteLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a metered product.
func deadline_DeleteMeteredProduct(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteMeteredProductInput{
		// LicenseEndpointId: *string, // Required
		// ProductId: *string, // Required
	}

	if len(_deadlineLicenseEndpointId) > 0 {
		input.LicenseEndpointId = aws.String(_deadlineLicenseEndpointId)
	}
	if len(_deadlineProductId) > 0 {
		input.ProductId = aws.String(_deadlineProductId)
	}

	if resp, err := client.DeleteMeteredProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a Deadline Cloud monitor. After you delete a monitor, you can create a
// new one and attach farms to the monitor.
func deadline_DeleteMonitor(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteMonitorInput{
		// MonitorId: *string, // Required
	}

	if len(_deadlineMonitorId) > 0 {
		input.MonitorId = aws.String(_deadlineMonitorId)
	}

	if resp, err := client.DeleteMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a queue.
// You can't recover the jobs in a queue if you delete the queue. Deleting the
// queue also deletes the jobs in that queue.
func deadline_DeleteQueue(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteQueueInput{
		// FarmId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.DeleteQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a queue environment.
func deadline_DeleteQueueEnvironment(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteQueueEnvironmentInput{
		// FarmId: *string, // Required
		// QueueEnvironmentId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueEnvironmentId) > 0 {
		input.QueueEnvironmentId = aws.String(_deadlineQueueEnvironmentId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.DeleteQueueEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a queue-fleet association.
func deadline_DeleteQueueFleetAssociation(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteQueueFleetAssociationInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.DeleteQueueFleetAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between a queue and a limit. You must use the
// UpdateQueueLimitAssociation operation to set the status to
// STOP_LIMIT_USAGE_AND_COMPLETE_TASKS or STOP_LIMIT_USAGE_AND_CANCEL_TASKS . The
// status does not change immediately. Use the GetQueueLimitAssociation operation
// to see if the status changed to STOPPED before deleting the association.
func deadline_DeleteQueueLimitAssociation(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteQueueLimitAssociationInput{
		// FarmId: *string, // Required
		// LimitId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineLimitId) > 0 {
		input.LimitId = aws.String(_deadlineLimitId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.DeleteQueueLimitAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a storage profile.
func deadline_DeleteStorageProfile(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteStorageProfileInput{
		// FarmId: *string, // Required
		// StorageProfileId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineStorageProfileId) > 0 {
		input.StorageProfileId = aws.String(_deadlineStorageProfileId)
	}

	if resp, err := client.DeleteStorageProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a worker.
func deadline_DeleteWorker(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DeleteWorkerInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineWorkerId) > 0 {
		input.WorkerId = aws.String(_deadlineWorkerId)
	}

	if resp, err := client.DeleteWorker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a member from a farm.
func deadline_DisassociateMemberFromFarm(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DisassociateMemberFromFarmInput{
		// FarmId: *string, // Required
		// PrincipalId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}

	if resp, err := client.DisassociateMemberFromFarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a member from a fleet.
func deadline_DisassociateMemberFromFleet(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DisassociateMemberFromFleetInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// PrincipalId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}

	if resp, err := client.DisassociateMemberFromFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a member from a job.
func deadline_DisassociateMemberFromJob(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DisassociateMemberFromJobInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// PrincipalId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.DisassociateMemberFromJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a member from a queue.
func deadline_DisassociateMemberFromQueue(cfg aws.Config, client *deadline.Client) {
	input := &deadline.DisassociateMemberFromQueueInput{
		// FarmId: *string, // Required
		// PrincipalId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.DisassociateMemberFromQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a budget.
func deadline_GetBudget(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetBudgetInput{
		// BudgetId: *string, // Required
		// FarmId: *string, // Required
	}

	if len(_deadlineBudgetId) > 0 {
		input.BudgetId = aws.String(_deadlineBudgetId)
	}
	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}

	if resp, err := client.GetBudget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a farm.
func deadline_GetFarm(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetFarmInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}

	if resp, err := client.GetFarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get a fleet.
func deadline_GetFleet(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetFleetInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}

	if resp, err := client.GetFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a Deadline Cloud job.
func deadline_GetJob(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetJobInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.GetJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a licence endpoint.
func deadline_GetLicenseEndpoint(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetLicenseEndpointInput{
		// LicenseEndpointId: *string, // Required
	}

	if len(_deadlineLicenseEndpointId) > 0 {
		input.LicenseEndpointId = aws.String(_deadlineLicenseEndpointId)
	}

	if resp, err := client.GetLicenseEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specific limit.
func deadline_GetLimit(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetLimitInput{
		// FarmId: *string, // Required
		// LimitId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineLimitId) > 0 {
		input.LimitId = aws.String(_deadlineLimitId)
	}

	if resp, err := client.GetLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified monitor.
func deadline_GetMonitor(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetMonitorInput{
		// MonitorId: *string, // Required
	}

	if len(_deadlineMonitorId) > 0 {
		input.MonitorId = aws.String(_deadlineMonitorId)
	}

	if resp, err := client.GetMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a queue.
func deadline_GetQueue(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetQueueInput{
		// FarmId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.GetQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a queue environment.
func deadline_GetQueueEnvironment(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetQueueEnvironmentInput{
		// FarmId: *string, // Required
		// QueueEnvironmentId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueEnvironmentId) > 0 {
		input.QueueEnvironmentId = aws.String(_deadlineQueueEnvironmentId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.GetQueueEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a queue-fleet association.
func deadline_GetQueueFleetAssociation(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetQueueFleetAssociationInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.GetQueueFleetAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specific association between a queue and a limit.
func deadline_GetQueueLimitAssociation(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetQueueLimitAssociationInput{
		// FarmId: *string, // Required
		// LimitId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineLimitId) > 0 {
		input.LimitId = aws.String(_deadlineLimitId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if resp, err := client.GetQueueLimitAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a session.
func deadline_GetSession(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetSessionInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// SessionId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineSessionId) > 0 {
		input.SessionId = aws.String(_deadlineSessionId)
	}

	if resp, err := client.GetSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a session action for the job.
func deadline_GetSessionAction(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetSessionActionInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// SessionActionId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineSessionActionId) > 0 {
		input.SessionActionId = aws.String(_deadlineSessionActionId)
	}

	if resp, err := client.GetSessionAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a set of statistics for queues or farms. Before you can call the
// GetSessionStatisticsAggregation operation, you must first call the
// StartSessionsStatisticsAggregation operation. Statistics are available for 1
// hour after you call the StartSessionsStatisticsAggregation operation.
func deadline_GetSessionsStatisticsAggregation(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetSessionsStatisticsAggregationInput{
		// AggregationId: *string, // Required
		// FarmId: *string, // Required
	}

	if len(_deadlineAggregationId) > 0 {
		input.AggregationId = aws.String(_deadlineAggregationId)
	}
	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetSessionsStatisticsAggregation(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.GetSessionsStatisticsAggregationOutput
	p := deadline.NewGetSessionsStatisticsAggregationPaginator(client, input)
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

// Gets a step.
func deadline_GetStep(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetStepInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// StepId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineStepId) > 0 {
		input.StepId = aws.String(_deadlineStepId)
	}

	if resp, err := client.GetStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a storage profile.
func deadline_GetStorageProfile(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetStorageProfileInput{
		// FarmId: *string, // Required
		// StorageProfileId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineStorageProfileId) > 0 {
		input.StorageProfileId = aws.String(_deadlineStorageProfileId)
	}

	if resp, err := client.GetStorageProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a storage profile for a queue.
func deadline_GetStorageProfileForQueue(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetStorageProfileForQueueInput{
		// FarmId: *string, // Required
		// QueueId: *string, // Required
		// StorageProfileId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineStorageProfileId) > 0 {
		input.StorageProfileId = aws.String(_deadlineStorageProfileId)
	}

	if resp, err := client.GetStorageProfileForQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a task.
func deadline_GetTask(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetTaskInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// StepId: *string, // Required
		// TaskId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineStepId) > 0 {
		input.StepId = aws.String(_deadlineStepId)
	}
	if len(_deadlineTaskId) > 0 {
		input.TaskId = aws.String(_deadlineTaskId)
	}

	if resp, err := client.GetTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a worker.
func deadline_GetWorker(cfg aws.Config, client *deadline.Client) {
	input := &deadline.GetWorkerInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineWorkerId) > 0 {
		input.WorkerId = aws.String(_deadlineWorkerId)
	}

	if resp, err := client.GetWorker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A list of the available metered products.
func deadline_ListAvailableMeteredProducts(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListAvailableMeteredProductsInput{}

	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAvailableMeteredProducts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListAvailableMeteredProductsOutput
	p := deadline.NewListAvailableMeteredProductsPaginator(client, input)
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

// A list of budgets in a farm.
func deadline_ListBudgets(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListBudgetsInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}
	if len(_deadlineStatus) > 0 {
		if err := assignInputField(input, "Status", _deadlineStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListBudgets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListBudgetsOutput
	p := deadline.NewListBudgetsPaginator(client, input)
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

// Lists the members of a farm.
func deadline_ListFarmMembers(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListFarmMembersInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFarmMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListFarmMembersOutput
	p := deadline.NewListFarmMembersPaginator(client, input)
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

// Lists farms.
func deadline_ListFarms(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListFarmsInput{}

	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}

	if disablePaginator() {
		if resp, err := client.ListFarms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListFarmsOutput
	p := deadline.NewListFarmsPaginator(client, input)
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

// Lists fleet members.
func deadline_ListFleetMembers(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListFleetMembersInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFleetMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListFleetMembersOutput
	p := deadline.NewListFleetMembersPaginator(client, input)
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

// Lists fleets.
func deadline_ListFleets(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListFleetsInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}
	if len(_deadlineStatus) > 0 {
		if err := assignInputField(input, "Status", _deadlineStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFleets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListFleetsOutput
	p := deadline.NewListFleetsPaginator(client, input)
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

// Lists members on a job.
func deadline_ListJobMembers(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListJobMembersInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListJobMembersOutput
	p := deadline.NewListJobMembersPaginator(client, input)
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

// Lists parameter definitions of a job.
func deadline_ListJobParameterDefinitions(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListJobParameterDefinitionsInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobParameterDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListJobParameterDefinitionsOutput
	p := deadline.NewListJobParameterDefinitionsPaginator(client, input)
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
func deadline_ListJobs(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListJobsInput{
		// FarmId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
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

	var results []*deadline.ListJobsOutput
	p := deadline.NewListJobsPaginator(client, input)
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

// Lists license endpoints.
func deadline_ListLicenseEndpoints(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListLicenseEndpointsInput{}

	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLicenseEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListLicenseEndpointsOutput
	p := deadline.NewListLicenseEndpointsPaginator(client, input)
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

// Gets a list of limits defined in the specified farm.
func deadline_ListLimits(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListLimitsInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLimits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListLimitsOutput
	p := deadline.NewListLimitsPaginator(client, input)
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

// Lists metered products.
func deadline_ListMeteredProducts(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListMeteredProductsInput{
		// LicenseEndpointId: *string, // Required
	}

	if len(_deadlineLicenseEndpointId) > 0 {
		input.LicenseEndpointId = aws.String(_deadlineLicenseEndpointId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMeteredProducts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListMeteredProductsOutput
	p := deadline.NewListMeteredProductsPaginator(client, input)
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

// Gets a list of your monitors in Deadline Cloud.
func deadline_ListMonitors(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListMonitorsInput{}

	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMonitors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListMonitorsOutput
	p := deadline.NewListMonitorsPaginator(client, input)
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

// Lists queue environments.
func deadline_ListQueueEnvironments(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListQueueEnvironmentsInput{
		// FarmId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListQueueEnvironments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListQueueEnvironmentsOutput
	p := deadline.NewListQueueEnvironmentsPaginator(client, input)
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

// Lists queue-fleet associations.
func deadline_ListQueueFleetAssociations(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListQueueFleetAssociationsInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if disablePaginator() {
		if resp, err := client.ListQueueFleetAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListQueueFleetAssociationsOutput
	p := deadline.NewListQueueFleetAssociationsPaginator(client, input)
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

// Gets a list of the associations between queues and limits defined in a farm.
func deadline_ListQueueLimitAssociations(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListQueueLimitAssociationsInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineLimitId) > 0 {
		input.LimitId = aws.String(_deadlineLimitId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}

	if disablePaginator() {
		if resp, err := client.ListQueueLimitAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListQueueLimitAssociationsOutput
	p := deadline.NewListQueueLimitAssociationsPaginator(client, input)
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

// Lists the members in a queue.
func deadline_ListQueueMembers(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListQueueMembersInput{
		// FarmId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListQueueMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListQueueMembersOutput
	p := deadline.NewListQueueMembersPaginator(client, input)
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

// Lists queues.
func deadline_ListQueues(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListQueuesInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}
	if len(_deadlinePrincipalId) > 0 {
		input.PrincipalId = aws.String(_deadlinePrincipalId)
	}
	if len(_deadlineStatus) > 0 {
		if err := assignInputField(input, "Status", _deadlineStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListQueues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListQueuesOutput
	p := deadline.NewListQueuesPaginator(client, input)
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

// Lists session actions.
func deadline_ListSessionActions(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListSessionActionsInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}
	if len(_deadlineSessionId) > 0 {
		input.SessionId = aws.String(_deadlineSessionId)
	}
	if len(_deadlineTaskId) > 0 {
		input.TaskId = aws.String(_deadlineTaskId)
	}

	if disablePaginator() {
		if resp, err := client.ListSessionActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListSessionActionsOutput
	p := deadline.NewListSessionActionsPaginator(client, input)
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

// Lists sessions.
func deadline_ListSessions(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListSessionsInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListSessionsOutput
	p := deadline.NewListSessionsPaginator(client, input)
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

// Lists sessions for a worker.
func deadline_ListSessionsForWorker(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListSessionsForWorkerInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineWorkerId) > 0 {
		input.WorkerId = aws.String(_deadlineWorkerId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSessionsForWorker(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListSessionsForWorkerOutput
	p := deadline.NewListSessionsForWorkerPaginator(client, input)
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

// Lists step consumers.
func deadline_ListStepConsumers(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListStepConsumersInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// StepId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineStepId) > 0 {
		input.StepId = aws.String(_deadlineStepId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStepConsumers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListStepConsumersOutput
	p := deadline.NewListStepConsumersPaginator(client, input)
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

// Lists the dependencies for a step.
func deadline_ListStepDependencies(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListStepDependenciesInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// StepId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineStepId) > 0 {
		input.StepId = aws.String(_deadlineStepId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStepDependencies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListStepDependenciesOutput
	p := deadline.NewListStepDependenciesPaginator(client, input)
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

// Lists steps for a job.
func deadline_ListSteps(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListStepsInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSteps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListStepsOutput
	p := deadline.NewListStepsPaginator(client, input)
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

// Lists storage profiles.
func deadline_ListStorageProfiles(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListStorageProfilesInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStorageProfiles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListStorageProfilesOutput
	p := deadline.NewListStorageProfilesPaginator(client, input)
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

// Lists storage profiles for a queue.
func deadline_ListStorageProfilesForQueue(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListStorageProfilesForQueueInput{
		// FarmId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStorageProfilesForQueue(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListStorageProfilesForQueueOutput
	p := deadline.NewListStorageProfilesForQueuePaginator(client, input)
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

// Lists tags for a resource.
func deadline_ListTagsForResource(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_deadlineResourceArn) > 0 {
		input.ResourceArn = aws.String(_deadlineResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists tasks for a job.
func deadline_ListTasks(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListTasksInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// StepId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineStepId) > 0 {
		input.StepId = aws.String(_deadlineStepId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTasks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListTasksOutput
	p := deadline.NewListTasksPaginator(client, input)
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

// Lists workers.
func deadline_ListWorkers(cfg aws.Config, client *deadline.Client) {
	input := &deadline.ListWorkersInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _deadlineMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_deadlineNextToken) > 0 {
		input.NextToken = aws.String(_deadlineNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*deadline.ListWorkersOutput
	p := deadline.NewListWorkersPaginator(client, input)
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

// Adds a metered product.
func deadline_PutMeteredProduct(cfg aws.Config, client *deadline.Client) {
	input := &deadline.PutMeteredProductInput{
		// LicenseEndpointId: *string, // Required
		// ProductId: *string, // Required
	}

	if len(_deadlineLicenseEndpointId) > 0 {
		input.LicenseEndpointId = aws.String(_deadlineLicenseEndpointId)
	}
	if len(_deadlineProductId) > 0 {
		input.ProductId = aws.String(_deadlineProductId)
	}

	if resp, err := client.PutMeteredProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for jobs.
func deadline_SearchJobs(cfg aws.Config, client *deadline.Client) {
	input := &deadline.SearchJobsInput{
		// FarmId: *string, // Required
		// ItemOffset: *int32, // Required
		// QueueIds: []string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineItemOffset) > 0 {
		if err := assignInputField(input, "ItemOffset", _deadlineItemOffset); err != nil {
			log.Errorf("invalid --item-offset: %s", err.Error())
			return
		}
	}
	if len(_deadlineQueueIds) > 0 {
		input.QueueIds = append([]string(nil), _deadlineQueueIds...)
	}
	if len(_deadlineFilterExpressions) > 0 {
		if err := assignInputField(input, "FilterExpressions", _deadlineFilterExpressions); err != nil {
			log.Errorf("invalid --filter-expressions: %s", err.Error())
			return
		}
	}
	if len(_deadlinePageSize) > 0 {
		if err := assignInputField(input, "PageSize", _deadlinePageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_deadlineSortExpressions) > 0 {
		if err := assignInputField(input, "SortExpressions", _deadlineSortExpressions); err != nil {
			log.Errorf("invalid --sort-expressions: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchJobs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for steps.
func deadline_SearchSteps(cfg aws.Config, client *deadline.Client) {
	input := &deadline.SearchStepsInput{
		// FarmId: *string, // Required
		// ItemOffset: *int32, // Required
		// QueueIds: []string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineItemOffset) > 0 {
		if err := assignInputField(input, "ItemOffset", _deadlineItemOffset); err != nil {
			log.Errorf("invalid --item-offset: %s", err.Error())
			return
		}
	}
	if len(_deadlineQueueIds) > 0 {
		input.QueueIds = append([]string(nil), _deadlineQueueIds...)
	}
	if len(_deadlineFilterExpressions) > 0 {
		if err := assignInputField(input, "FilterExpressions", _deadlineFilterExpressions); err != nil {
			log.Errorf("invalid --filter-expressions: %s", err.Error())
			return
		}
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlinePageSize) > 0 {
		if err := assignInputField(input, "PageSize", _deadlinePageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_deadlineSortExpressions) > 0 {
		if err := assignInputField(input, "SortExpressions", _deadlineSortExpressions); err != nil {
			log.Errorf("invalid --sort-expressions: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchSteps(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for tasks.
func deadline_SearchTasks(cfg aws.Config, client *deadline.Client) {
	input := &deadline.SearchTasksInput{
		// FarmId: *string, // Required
		// ItemOffset: *int32, // Required
		// QueueIds: []string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineItemOffset) > 0 {
		if err := assignInputField(input, "ItemOffset", _deadlineItemOffset); err != nil {
			log.Errorf("invalid --item-offset: %s", err.Error())
			return
		}
	}
	if len(_deadlineQueueIds) > 0 {
		input.QueueIds = append([]string(nil), _deadlineQueueIds...)
	}
	if len(_deadlineFilterExpressions) > 0 {
		if err := assignInputField(input, "FilterExpressions", _deadlineFilterExpressions); err != nil {
			log.Errorf("invalid --filter-expressions: %s", err.Error())
			return
		}
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlinePageSize) > 0 {
		if err := assignInputField(input, "PageSize", _deadlinePageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_deadlineSortExpressions) > 0 {
		if err := assignInputField(input, "SortExpressions", _deadlineSortExpressions); err != nil {
			log.Errorf("invalid --sort-expressions: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for workers.
func deadline_SearchWorkers(cfg aws.Config, client *deadline.Client) {
	input := &deadline.SearchWorkersInput{
		// FarmId: *string, // Required
		// FleetIds: []string, // Required
		// ItemOffset: *int32, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetIds) > 0 {
		input.FleetIds = append([]string(nil), _deadlineFleetIds...)
	}
	if len(_deadlineItemOffset) > 0 {
		if err := assignInputField(input, "ItemOffset", _deadlineItemOffset); err != nil {
			log.Errorf("invalid --item-offset: %s", err.Error())
			return
		}
	}
	if len(_deadlineFilterExpressions) > 0 {
		if err := assignInputField(input, "FilterExpressions", _deadlineFilterExpressions); err != nil {
			log.Errorf("invalid --filter-expressions: %s", err.Error())
			return
		}
	}
	if len(_deadlinePageSize) > 0 {
		if err := assignInputField(input, "PageSize", _deadlinePageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_deadlineSortExpressions) > 0 {
		if err := assignInputField(input, "SortExpressions", _deadlineSortExpressions); err != nil {
			log.Errorf("invalid --sort-expressions: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchWorkers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous request for getting aggregated statistics about queues
// and farms. Get the statistics using the GetSessionsStatisticsAggregation
// operation. You can only have one running aggregation for your Deadline Cloud
// farm. Call the GetSessionsStatisticsAggregation operation and check the status
// field to see if an aggregation is running. Statistics are available for 1 hour
// after you call the StartSessionsStatisticsAggregation operation.
func deadline_StartSessionsStatisticsAggregation(cfg aws.Config, client *deadline.Client) {
	input := &deadline.StartSessionsStatisticsAggregationInput{
		// EndTime: *time.Time, // Required
		// FarmId: *string, // Required
		// GroupBy: []types.UsageGroupByField, // Required
		// ResourceIds: types.SessionsStatisticsResources, // Required
		// StartTime: *time.Time, // Required
		// Statistics: []types.UsageStatistic, // Required
	}

	if len(_deadlineEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _deadlineEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _deadlineGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_deadlineResourceIds) > 0 {
		if err := assignInputField(input, "ResourceIds", _deadlineResourceIds); err != nil {
			log.Errorf("invalid --resource-ids: %s", err.Error())
			return
		}
	}
	if len(_deadlineStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _deadlineStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_deadlineStatistics) > 0 {
		if err := assignInputField(input, "Statistics", _deadlineStatistics); err != nil {
			log.Errorf("invalid --statistics: %s", err.Error())
			return
		}
	}
	if len(_deadlinePeriod) > 0 {
		if err := assignInputField(input, "Period", _deadlinePeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}
	if len(_deadlineTimezone) > 0 {
		input.Timezone = aws.String(_deadlineTimezone)
	}

	if resp, err := client.StartSessionsStatisticsAggregation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a resource using the resource's ARN and desired tags.
func deadline_TagResource(cfg aws.Config, client *deadline.Client) {
	input := &deadline.TagResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_deadlineResourceArn) > 0 {
		input.ResourceArn = aws.String(_deadlineResourceArn)
	}
	if len(_deadlineTags) > 0 {
		if err := assignInputField(input, "Tags", _deadlineTags); err != nil {
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

// Removes a tag from a resource using the resource's ARN and tag to remove.
func deadline_UntagResource(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_deadlineResourceArn) > 0 {
		input.ResourceArn = aws.String(_deadlineResourceArn)
	}
	if len(_deadlineTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _deadlineTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a budget that sets spending thresholds for rendering activity.
func deadline_UpdateBudget(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateBudgetInput{
		// BudgetId: *string, // Required
		// FarmId: *string, // Required
	}

	if len(_deadlineBudgetId) > 0 {
		input.BudgetId = aws.String(_deadlineBudgetId)
	}
	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineActionsToAdd) > 0 {
		if err := assignInputField(input, "ActionsToAdd", _deadlineActionsToAdd); err != nil {
			log.Errorf("invalid --actions-to-add: %s", err.Error())
			return
		}
	}
	if len(_deadlineActionsToRemove) > 0 {
		if err := assignInputField(input, "ActionsToRemove", _deadlineActionsToRemove); err != nil {
			log.Errorf("invalid --actions-to-remove: %s", err.Error())
			return
		}
	}
	if len(_deadlineApproximateDollarLimit) > 0 {
		if err := assignInputField(input, "ApproximateDollarLimit", _deadlineApproximateDollarLimit); err != nil {
			log.Errorf("invalid --approximate-dollar-limit: %s", err.Error())
			return
		}
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _deadlineSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_deadlineStatus) > 0 {
		if err := assignInputField(input, "Status", _deadlineStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBudget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a farm.
func deadline_UpdateFarm(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateFarmInput{
		// FarmId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}

	if resp, err := client.UpdateFarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a fleet.
func deadline_UpdateFleet(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateFleetInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _deadlineConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineHostConfiguration) > 0 {
		if err := assignInputField(input, "HostConfiguration", _deadlineHostConfiguration); err != nil {
			log.Errorf("invalid --host-configuration: %s", err.Error())
			return
		}
	}
	if len(_deadlineMaxWorkerCount) > 0 {
		if err := assignInputField(input, "MaxWorkerCount", _deadlineMaxWorkerCount); err != nil {
			log.Errorf("invalid --max-worker-count: %s", err.Error())
			return
		}
	}
	if len(_deadlineMinWorkerCount) > 0 {
		if err := assignInputField(input, "MinWorkerCount", _deadlineMinWorkerCount); err != nil {
			log.Errorf("invalid --min-worker-count: %s", err.Error())
			return
		}
	}
	if len(_deadlineRoleArn) > 0 {
		input.RoleArn = aws.String(_deadlineRoleArn)
	}

	if resp, err := client.UpdateFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a job.
// When you change the status of the job to ARCHIVED , the job can't be scheduled
// or archived.
//
// An archived jobs and its steps and tasks are deleted after 120 days. The job
// can't be recovered.
func deadline_UpdateJob(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateJobInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}
	if len(_deadlineLifecycleStatus) > 0 {
		if err := assignInputField(input, "LifecycleStatus", _deadlineLifecycleStatus); err != nil {
			log.Errorf("invalid --lifecycle-status: %s", err.Error())
			return
		}
	}
	if len(_deadlineMaxFailedTasksCount) > 0 {
		if err := assignInputField(input, "MaxFailedTasksCount", _deadlineMaxFailedTasksCount); err != nil {
			log.Errorf("invalid --max-failed-tasks-count: %s", err.Error())
			return
		}
	}
	if len(_deadlineMaxRetriesPerTask) > 0 {
		if err := assignInputField(input, "MaxRetriesPerTask", _deadlineMaxRetriesPerTask); err != nil {
			log.Errorf("invalid --max-retries-per-task: %s", err.Error())
			return
		}
	}
	if len(_deadlineMaxWorkerCount) > 0 {
		if err := assignInputField(input, "MaxWorkerCount", _deadlineMaxWorkerCount); err != nil {
			log.Errorf("invalid --max-worker-count: %s", err.Error())
			return
		}
	}
	if len(_deadlineName) > 0 {
		input.Name = aws.String(_deadlineName)
	}
	if len(_deadlinePriority) > 0 {
		if err := assignInputField(input, "Priority", _deadlinePriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_deadlineTargetTaskRunStatus) > 0 {
		if err := assignInputField(input, "TargetTaskRunStatus", _deadlineTargetTaskRunStatus); err != nil {
			log.Errorf("invalid --target-task-run-status: %s", err.Error())
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

// Updates the properties of the specified limit.
func deadline_UpdateLimit(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateLimitInput{
		// FarmId: *string, // Required
		// LimitId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineLimitId) > 0 {
		input.LimitId = aws.String(_deadlineLimitId)
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineMaxCount) > 0 {
		if err := assignInputField(input, "MaxCount", _deadlineMaxCount); err != nil {
			log.Errorf("invalid --max-count: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the settings for a Deadline Cloud monitor. You can modify one or all
// of the settings when you call UpdateMonitor .
func deadline_UpdateMonitor(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateMonitorInput{
		// MonitorId: *string, // Required
	}

	if len(_deadlineMonitorId) > 0 {
		input.MonitorId = aws.String(_deadlineMonitorId)
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineRoleArn) > 0 {
		input.RoleArn = aws.String(_deadlineRoleArn)
	}
	if len(_deadlineSubdomain) > 0 {
		input.Subdomain = aws.String(_deadlineSubdomain)
	}

	if resp, err := client.UpdateMonitor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a queue.
func deadline_UpdateQueue(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateQueueInput{
		// FarmId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineAllowedStorageProfileIdsToAdd) > 0 {
		input.AllowedStorageProfileIdsToAdd = append([]string(nil), _deadlineAllowedStorageProfileIdsToAdd...)
	}
	if len(_deadlineAllowedStorageProfileIdsToRemove) > 0 {
		input.AllowedStorageProfileIdsToRemove = append([]string(nil), _deadlineAllowedStorageProfileIdsToRemove...)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineDefaultBudgetAction) > 0 {
		if err := assignInputField(input, "DefaultBudgetAction", _deadlineDefaultBudgetAction); err != nil {
			log.Errorf("invalid --default-budget-action: %s", err.Error())
			return
		}
	}
	if len(_deadlineDescription) > 0 {
		input.Description = aws.String(_deadlineDescription)
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineJobAttachmentSettings) > 0 {
		if err := assignInputField(input, "JobAttachmentSettings", _deadlineJobAttachmentSettings); err != nil {
			log.Errorf("invalid --job-attachment-settings: %s", err.Error())
			return
		}
	}
	if len(_deadlineJobRunAsUser) > 0 {
		if err := assignInputField(input, "JobRunAsUser", _deadlineJobRunAsUser); err != nil {
			log.Errorf("invalid --job-run-as-user: %s", err.Error())
			return
		}
	}
	if len(_deadlineRequiredFileSystemLocationNamesToAdd) > 0 {
		input.RequiredFileSystemLocationNamesToAdd = append([]string(nil), _deadlineRequiredFileSystemLocationNamesToAdd...)
	}
	if len(_deadlineRequiredFileSystemLocationNamesToRemove) > 0 {
		input.RequiredFileSystemLocationNamesToRemove = append([]string(nil), _deadlineRequiredFileSystemLocationNamesToRemove...)
	}
	if len(_deadlineRoleArn) > 0 {
		input.RoleArn = aws.String(_deadlineRoleArn)
	}

	if resp, err := client.UpdateQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the queue environment.
func deadline_UpdateQueueEnvironment(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateQueueEnvironmentInput{
		// FarmId: *string, // Required
		// QueueEnvironmentId: *string, // Required
		// QueueId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineQueueEnvironmentId) > 0 {
		input.QueueEnvironmentId = aws.String(_deadlineQueueEnvironmentId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlinePriority) > 0 {
		if err := assignInputField(input, "Priority", _deadlinePriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_deadlineTemplate) > 0 {
		input.Template = aws.String(_deadlineTemplate)
	}
	if len(_deadlineTemplateType) > 0 {
		if err := assignInputField(input, "TemplateType", _deadlineTemplateType); err != nil {
			log.Errorf("invalid --template-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQueueEnvironment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a queue-fleet association.
func deadline_UpdateQueueFleetAssociation(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateQueueFleetAssociationInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// QueueId: *string, // Required
		// Status: types.UpdateQueueFleetAssociationStatus, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineStatus) > 0 {
		if err := assignInputField(input, "Status", _deadlineStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQueueFleetAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of the queue. If you set the status to one of the
// STOP_LIMIT_USAGE* values, there will be a delay before the status transitions to
// the STOPPED state.
func deadline_UpdateQueueLimitAssociation(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateQueueLimitAssociationInput{
		// FarmId: *string, // Required
		// LimitId: *string, // Required
		// QueueId: *string, // Required
		// Status: types.UpdateQueueLimitAssociationStatus, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineLimitId) > 0 {
		input.LimitId = aws.String(_deadlineLimitId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineStatus) > 0 {
		if err := assignInputField(input, "Status", _deadlineStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQueueLimitAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a session.
func deadline_UpdateSession(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateSessionInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// SessionId: *string, // Required
		// TargetLifecycleStatus: types.SessionLifecycleTargetStatus, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineSessionId) > 0 {
		input.SessionId = aws.String(_deadlineSessionId)
	}
	if len(_deadlineTargetLifecycleStatus) > 0 {
		if err := assignInputField(input, "TargetLifecycleStatus", _deadlineTargetLifecycleStatus); err != nil {
			log.Errorf("invalid --target-lifecycle-status: %s", err.Error())
			return
		}
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}

	if resp, err := client.UpdateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a step.
func deadline_UpdateStep(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateStepInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// StepId: *string, // Required
		// TargetTaskRunStatus: types.StepTargetTaskRunStatus, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineStepId) > 0 {
		input.StepId = aws.String(_deadlineStepId)
	}
	if len(_deadlineTargetTaskRunStatus) > 0 {
		if err := assignInputField(input, "TargetTaskRunStatus", _deadlineTargetTaskRunStatus); err != nil {
			log.Errorf("invalid --target-task-run-status: %s", err.Error())
			return
		}
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}

	if resp, err := client.UpdateStep(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a storage profile.
func deadline_UpdateStorageProfile(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateStorageProfileInput{
		// FarmId: *string, // Required
		// StorageProfileId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineStorageProfileId) > 0 {
		input.StorageProfileId = aws.String(_deadlineStorageProfileId)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}
	if len(_deadlineDisplayName) > 0 {
		input.DisplayName = aws.String(_deadlineDisplayName)
	}
	if len(_deadlineFileSystemLocationsToAdd) > 0 {
		if err := assignInputField(input, "FileSystemLocationsToAdd", _deadlineFileSystemLocationsToAdd); err != nil {
			log.Errorf("invalid --file-system-locations-to-add: %s", err.Error())
			return
		}
	}
	if len(_deadlineFileSystemLocationsToRemove) > 0 {
		if err := assignInputField(input, "FileSystemLocationsToRemove", _deadlineFileSystemLocationsToRemove); err != nil {
			log.Errorf("invalid --file-system-locations-to-remove: %s", err.Error())
			return
		}
	}
	if len(_deadlineOsFamily) > 0 {
		if err := assignInputField(input, "OsFamily", _deadlineOsFamily); err != nil {
			log.Errorf("invalid --os-family: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStorageProfile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a task.
func deadline_UpdateTask(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateTaskInput{
		// FarmId: *string, // Required
		// JobId: *string, // Required
		// QueueId: *string, // Required
		// StepId: *string, // Required
		// TargetRunStatus: types.TaskTargetRunStatus, // Required
		// TaskId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineJobId) > 0 {
		input.JobId = aws.String(_deadlineJobId)
	}
	if len(_deadlineQueueId) > 0 {
		input.QueueId = aws.String(_deadlineQueueId)
	}
	if len(_deadlineStepId) > 0 {
		input.StepId = aws.String(_deadlineStepId)
	}
	if len(_deadlineTargetRunStatus) > 0 {
		if err := assignInputField(input, "TargetRunStatus", _deadlineTargetRunStatus); err != nil {
			log.Errorf("invalid --target-run-status: %s", err.Error())
			return
		}
	}
	if len(_deadlineTaskId) > 0 {
		input.TaskId = aws.String(_deadlineTaskId)
	}
	if len(_deadlineClientToken) > 0 {
		input.ClientToken = aws.String(_deadlineClientToken)
	}

	if resp, err := client.UpdateTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a worker.
func deadline_UpdateWorker(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateWorkerInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineWorkerId) > 0 {
		input.WorkerId = aws.String(_deadlineWorkerId)
	}
	if len(_deadlineCapabilities) > 0 {
		if err := assignInputField(input, "Capabilities", _deadlineCapabilities); err != nil {
			log.Errorf("invalid --capabilities: %s", err.Error())
			return
		}
	}
	if len(_deadlineHostProperties) > 0 {
		if err := assignInputField(input, "HostProperties", _deadlineHostProperties); err != nil {
			log.Errorf("invalid --host-properties: %s", err.Error())
			return
		}
	}
	if len(_deadlineStatus) > 0 {
		if err := assignInputField(input, "Status", _deadlineStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the schedule for a worker.
func deadline_UpdateWorkerSchedule(cfg aws.Config, client *deadline.Client) {
	input := &deadline.UpdateWorkerScheduleInput{
		// FarmId: *string, // Required
		// FleetId: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_deadlineFarmId) > 0 {
		input.FarmId = aws.String(_deadlineFarmId)
	}
	if len(_deadlineFleetId) > 0 {
		input.FleetId = aws.String(_deadlineFleetId)
	}
	if len(_deadlineWorkerId) > 0 {
		input.WorkerId = aws.String(_deadlineWorkerId)
	}
	if len(_deadlineUpdatedSessionActions) > 0 {
		if err := assignInputField(input, "UpdatedSessionActions", _deadlineUpdatedSessionActions); err != nil {
			log.Errorf("invalid --updated-session-actions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWorkerSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_deadlineCmd)
	_deadlineCmd.Flags().SortFlags = false

	_deadlineCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_deadlineCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_deadlineCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_deadlineCmd.Flags().StringVarP(&_deadlineActions, "actions", "", "", "Actions")
	_deadlineCmd.Flags().StringVarP(&_deadlineActionsToAdd, "actions-to-add", "", "", "Actions To Add")
	_deadlineCmd.Flags().StringVarP(&_deadlineActionsToRemove, "actions-to-remove", "", "", "Actions To Remove")
	_deadlineCmd.Flags().StringVarP(&_deadlineAggregationId, "aggregation-id", "", "", "Aggregation ID")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineAllowedStorageProfileIds, "allowed-storage-profile-ids", "", nil, "Allowed Storage Profile Ids")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineAllowedStorageProfileIdsToAdd, "allowed-storage-profile-ids-to-add", "", nil, "Allowed Storage Profile Ids To Add")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineAllowedStorageProfileIdsToRemove, "allowed-storage-profile-ids-to-remove", "", nil, "Allowed Storage Profile Ids To Remove")
	_deadlineCmd.Flags().StringVarP(&_deadlineAmountRequirementName, "amount-requirement-name", "", "", "Amount Requirement Name")
	_deadlineCmd.Flags().StringVarP(&_deadlineApproximateDollarLimit, "approximate-dollar-limit", "", "", "Approximate Dollar Limit")
	_deadlineCmd.Flags().StringVarP(&_deadlineAttachments, "attachments", "", "", "Attachments")
	_deadlineCmd.Flags().StringVarP(&_deadlineBudgetId, "budget-id", "", "", "Budget ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineCapabilities, "capabilities", "", "", "Capabilities")
	_deadlineCmd.Flags().StringVarP(&_deadlineClientToken, "client-token", "", "", "Client Token")
	_deadlineCmd.Flags().StringVarP(&_deadlineConfiguration, "configuration", "", "", "Configuration")
	_deadlineCmd.Flags().StringVarP(&_deadlineDefaultBudgetAction, "default-budget-action", "", "", "Default Budget Action")
	_deadlineCmd.Flags().StringVarP(&_deadlineDescription, "description", "", "", "Description")
	_deadlineCmd.Flags().StringVarP(&_deadlineDescriptionOverride, "description-override", "", "", "Description Override")
	_deadlineCmd.Flags().StringVarP(&_deadlineDisplayName, "display-name", "", "", "Display Name")
	_deadlineCmd.Flags().StringVarP(&_deadlineEndTime, "end-time", "", "", "End Time")
	_deadlineCmd.Flags().StringVarP(&_deadlineFarmId, "farm-id", "", "", "Farm ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineFileSystemLocations, "file-system-locations", "", "", "File System Locations")
	_deadlineCmd.Flags().StringVarP(&_deadlineFileSystemLocationsToAdd, "file-system-locations-to-add", "", "", "File System Locations To Add")
	_deadlineCmd.Flags().StringVarP(&_deadlineFileSystemLocationsToRemove, "file-system-locations-to-remove", "", "", "File System Locations To Remove")
	_deadlineCmd.Flags().StringVarP(&_deadlineFilterExpressions, "filter-expressions", "", "", "Filter Expressions")
	_deadlineCmd.Flags().StringVarP(&_deadlineFleetId, "fleet-id", "", "", "Fleet ID")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineFleetIds, "fleet-ids", "", nil, "Fleet Ids")
	_deadlineCmd.Flags().StringVarP(&_deadlineGroupBy, "group-by", "", "", "Group By")
	_deadlineCmd.Flags().StringVarP(&_deadlineHostConfiguration, "host-configuration", "", "", "Host Configuration")
	_deadlineCmd.Flags().StringVarP(&_deadlineHostProperties, "host-properties", "", "", "Host Properties")
	_deadlineCmd.Flags().StringVarP(&_deadlineIdentifiers, "identifiers", "", "", "Identifiers")
	_deadlineCmd.Flags().StringVarP(&_deadlineIdentityCenterInstanceArn, "identity-center-instance-arn", "", "", "Identity Center Instance ARN")
	_deadlineCmd.Flags().StringVarP(&_deadlineIdentityStoreId, "identity-store-id", "", "", "Identity Store ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineItemOffset, "item-offset", "", "", "Item Offset")
	_deadlineCmd.Flags().StringVarP(&_deadlineJobAttachmentSettings, "job-attachment-settings", "", "", "Job Attachment Settings")
	_deadlineCmd.Flags().StringVarP(&_deadlineJobId, "job-id", "", "", "Job ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineJobRunAsUser, "job-run-as-user", "", "", "Job Run As User")
	_deadlineCmd.Flags().StringVarP(&_deadlineKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_deadlineCmd.Flags().StringVarP(&_deadlineLicenseEndpointId, "license-endpoint-id", "", "", "License Endpoint ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineLifecycleStatus, "lifecycle-status", "", "", "Lifecycle Status")
	_deadlineCmd.Flags().StringVarP(&_deadlineLimitId, "limit-id", "", "", "Limit ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineMaxCount, "max-count", "", "", "Max Count")
	_deadlineCmd.Flags().StringVarP(&_deadlineMaxFailedTasksCount, "max-failed-tasks-count", "", "", "Max Failed Tasks Count")
	_deadlineCmd.Flags().StringVarP(&_deadlineMaxResults, "max-results", "", "", "Max Results")
	_deadlineCmd.Flags().StringVarP(&_deadlineMaxRetriesPerTask, "max-retries-per-task", "", "", "Max Retries Per Task")
	_deadlineCmd.Flags().StringVarP(&_deadlineMaxWorkerCount, "max-worker-count", "", "", "Max Worker Count")
	_deadlineCmd.Flags().StringVarP(&_deadlineMembershipLevel, "membership-level", "", "", "Membership Level")
	_deadlineCmd.Flags().StringVarP(&_deadlineMinWorkerCount, "min-worker-count", "", "", "Min Worker Count")
	_deadlineCmd.Flags().StringVarP(&_deadlineMonitorId, "monitor-id", "", "", "Monitor ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineName, "name", "", "", "Name")
	_deadlineCmd.Flags().StringVarP(&_deadlineNameOverride, "name-override", "", "", "Name Override")
	_deadlineCmd.Flags().StringVarP(&_deadlineNextToken, "next-token", "", "", "Next Token")
	_deadlineCmd.Flags().StringVarP(&_deadlineOsFamily, "os-family", "", "", "OS Family")
	_deadlineCmd.Flags().StringVarP(&_deadlinePageSize, "page-size", "", "", "Page Size")
	_deadlineCmd.Flags().StringVarP(&_deadlineParameters, "parameters", "", "", "Parameters")
	_deadlineCmd.Flags().StringVarP(&_deadlinePeriod, "period", "", "", "Period")
	_deadlineCmd.Flags().StringVarP(&_deadlinePrincipalId, "principal-id", "", "", "Principal ID")
	_deadlineCmd.Flags().StringVarP(&_deadlinePrincipalType, "principal-type", "", "", "Principal Type")
	_deadlineCmd.Flags().StringVarP(&_deadlinePriority, "priority", "", "", "Priority")
	_deadlineCmd.Flags().StringVarP(&_deadlineProductId, "product-id", "", "", "Product ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineQueueEnvironmentId, "queue-environment-id", "", "", "Queue Environment ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineQueueId, "queue-id", "", "", "Queue ID")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineQueueIds, "queue-ids", "", nil, "Queue Ids")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineRequiredFileSystemLocationNames, "required-file-system-location-names", "", nil, "Required File System Location Names")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineRequiredFileSystemLocationNamesToAdd, "required-file-system-location-names-to-add", "", nil, "Required File System Location Names To Add")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineRequiredFileSystemLocationNamesToRemove, "required-file-system-location-names-to-remove", "", nil, "Required File System Location Names To Remove")
	_deadlineCmd.Flags().StringVarP(&_deadlineResourceArn, "resource-arn", "", "", "Resource ARN")
	_deadlineCmd.Flags().StringVarP(&_deadlineResourceIds, "resource-ids", "", "", "Resource Ids")
	_deadlineCmd.Flags().StringVarP(&_deadlineRoleArn, "role-arn", "", "", "Role ARN")
	_deadlineCmd.Flags().StringVarP(&_deadlineSchedule, "schedule", "", "", "Schedule")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_deadlineCmd.Flags().StringVarP(&_deadlineSessionActionId, "session-action-id", "", "", "Session Action ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineSessionId, "session-id", "", "", "Session ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineSortExpressions, "sort-expressions", "", "", "Sort Expressions")
	_deadlineCmd.Flags().StringVarP(&_deadlineSourceJobId, "source-job-id", "", "", "Source Job ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineStartTime, "start-time", "", "", "Start Time")
	_deadlineCmd.Flags().StringVarP(&_deadlineStatistics, "statistics", "", "", "Statistics")
	_deadlineCmd.Flags().StringVarP(&_deadlineStatus, "status", "", "", "Status")
	_deadlineCmd.Flags().StringVarP(&_deadlineStepId, "step-id", "", "", "Step ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineStorageProfileId, "storage-profile-id", "", "", "Storage Profile ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineSubdomain, "subdomain", "", "", "Subdomain")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_deadlineCmd.Flags().StringSliceVarP(&_deadlineTagKeys, "tag-keys", "", nil, "Tag Keys")
	_deadlineCmd.Flags().StringVarP(&_deadlineTags, "tags", "", "", "Tags")
	_deadlineCmd.Flags().StringVarP(&_deadlineTargetLifecycleStatus, "target-lifecycle-status", "", "", "Target Lifecycle Status")
	_deadlineCmd.Flags().StringVarP(&_deadlineTargetRunStatus, "target-run-status", "", "", "Target Run Status")
	_deadlineCmd.Flags().StringVarP(&_deadlineTargetS3Location, "target-s3-location", "", "", "Target S3 Location")
	_deadlineCmd.Flags().StringVarP(&_deadlineTargetTaskRunStatus, "target-task-run-status", "", "", "Target Task Run Status")
	_deadlineCmd.Flags().StringVarP(&_deadlineTaskId, "task-id", "", "", "Task ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineTemplate, "template", "", "", "Template")
	_deadlineCmd.Flags().StringVarP(&_deadlineTemplateType, "template-type", "", "", "Template Type")
	_deadlineCmd.Flags().StringVarP(&_deadlineTimezone, "timezone", "", "", "Timezone")
	_deadlineCmd.Flags().StringVarP(&_deadlineUpdatedSessionActions, "updated-session-actions", "", "", "Updated Session Actions")
	_deadlineCmd.Flags().StringVarP(&_deadlineUsageTrackingResource, "usage-tracking-resource", "", "", "Usage Tracking Resource")
	_deadlineCmd.Flags().StringVarP(&_deadlineVpcId, "vpc-id", "", "", "VPC ID")
	_deadlineCmd.Flags().StringVarP(&_deadlineWorkerId, "worker-id", "", "", "Worker ID")

	_deadlineCmd.Flags().BoolVarP(&_deadlineAssociateMemberToFarm, "associate-member-to-farm", "", false, "Associate Member To Farm")
	_deadlineCmd.Flags().BoolVarP(&_deadlineAssociateMemberToFleet, "associate-member-to-fleet", "", false, "Associate Member To Fleet")
	_deadlineCmd.Flags().BoolVarP(&_deadlineAssociateMemberToJob, "associate-member-to-job", "", false, "Associate Member To Job")
	_deadlineCmd.Flags().BoolVarP(&_deadlineAssociateMemberToQueue, "associate-member-to-queue", "", false, "Associate Member To Queue")
	_deadlineCmd.Flags().BoolVarP(&_deadlineAssumeFleetRoleForRead, "assume-fleet-role-for-read", "", false, "Assume Fleet Role For Read")
	_deadlineCmd.Flags().BoolVarP(&_deadlineAssumeFleetRoleForWorker, "assume-fleet-role-for-worker", "", false, "Assume Fleet Role For Worker")
	_deadlineCmd.Flags().BoolVarP(&_deadlineAssumeQueueRoleForRead, "assume-queue-role-for-read", "", false, "Assume Queue Role For Read")
	_deadlineCmd.Flags().BoolVarP(&_deadlineAssumeQueueRoleForUser, "assume-queue-role-for-user", "", false, "Assume Queue Role For User")
	_deadlineCmd.Flags().BoolVarP(&_deadlineAssumeQueueRoleForWorker, "assume-queue-role-for-worker", "", false, "Assume Queue Role For Worker")
	_deadlineCmd.Flags().BoolVarP(&_deadlineBatchGetJobEntity, "batch-get-job-entity", "", false, "Batch Get Job Entity")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCopyJobTemplate, "copy-job-template", "", false, "Copy Job Template")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateBudget, "create-budget", "", false, "Create Budget")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateFarm, "create-farm", "", false, "Create Farm")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateFleet, "create-fleet", "", false, "Create Fleet")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateJob, "create-job", "", false, "Create Job")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateLicenseEndpoint, "create-license-endpoint", "", false, "Create License Endpoint")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateLimit, "create-limit", "", false, "Create Limit")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateMonitor, "create-monitor", "", false, "Create Monitor")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateQueue, "create-queue", "", false, "Create Queue")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateQueueEnvironment, "create-queue-environment", "", false, "Create Queue Environment")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateQueueFleetAssociation, "create-queue-fleet-association", "", false, "Create Queue Fleet Association")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateQueueLimitAssociation, "create-queue-limit-association", "", false, "Create Queue Limit Association")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateStorageProfile, "create-storage-profile", "", false, "Create Storage Profile")
	_deadlineCmd.Flags().BoolVarP(&_deadlineCreateWorker, "create-worker", "", false, "Create Worker")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteBudget, "delete-budget", "", false, "Delete Budget")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteFarm, "delete-farm", "", false, "Delete Farm")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteFleet, "delete-fleet", "", false, "Delete Fleet")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteLicenseEndpoint, "delete-license-endpoint", "", false, "Delete License Endpoint")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteLimit, "delete-limit", "", false, "Delete Limit")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteMeteredProduct, "delete-metered-product", "", false, "Delete Metered Product")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteMonitor, "delete-monitor", "", false, "Delete Monitor")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteQueue, "delete-queue", "", false, "Delete Queue")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteQueueEnvironment, "delete-queue-environment", "", false, "Delete Queue Environment")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteQueueFleetAssociation, "delete-queue-fleet-association", "", false, "Delete Queue Fleet Association")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteQueueLimitAssociation, "delete-queue-limit-association", "", false, "Delete Queue Limit Association")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteStorageProfile, "delete-storage-profile", "", false, "Delete Storage Profile")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDeleteWorker, "delete-worker", "", false, "Delete Worker")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDisassociateMemberFromFarm, "disassociate-member-from-farm", "", false, "Disassociate Member From Farm")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDisassociateMemberFromFleet, "disassociate-member-from-fleet", "", false, "Disassociate Member From Fleet")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDisassociateMemberFromJob, "disassociate-member-from-job", "", false, "Disassociate Member From Job")
	_deadlineCmd.Flags().BoolVarP(&_deadlineDisassociateMemberFromQueue, "disassociate-member-from-queue", "", false, "Disassociate Member From Queue")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetBudget, "get-budget", "", false, "Get Budget")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetFarm, "get-farm", "", false, "Get Farm")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetFleet, "get-fleet", "", false, "Get Fleet")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetJob, "get-job", "", false, "Get Job")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetLicenseEndpoint, "get-license-endpoint", "", false, "Get License Endpoint")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetLimit, "get-limit", "", false, "Get Limit")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetMonitor, "get-monitor", "", false, "Get Monitor")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetQueue, "get-queue", "", false, "Get Queue")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetQueueEnvironment, "get-queue-environment", "", false, "Get Queue Environment")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetQueueFleetAssociation, "get-queue-fleet-association", "", false, "Get Queue Fleet Association")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetQueueLimitAssociation, "get-queue-limit-association", "", false, "Get Queue Limit Association")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetSession, "get-session", "", false, "Get Session")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetSessionAction, "get-session-action", "", false, "Get Session Action")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetSessionsStatisticsAggregation, "get-sessions-statistics-aggregation", "", false, "Get Sessions Statistics Aggregation")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetStep, "get-step", "", false, "Get Step")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetStorageProfile, "get-storage-profile", "", false, "Get Storage Profile")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetStorageProfileForQueue, "get-storage-profile-for-queue", "", false, "Get Storage Profile For Queue")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetTask, "get-task", "", false, "Get Task")
	_deadlineCmd.Flags().BoolVarP(&_deadlineGetWorker, "get-worker", "", false, "Get Worker")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListAvailableMeteredProducts, "list-available-metered-products", "", false, "List Available Metered Products")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListBudgets, "list-budgets", "", false, "List Budgets")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListFarmMembers, "list-farm-members", "", false, "List Farm Members")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListFarms, "list-farms", "", false, "List Farms")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListFleetMembers, "list-fleet-members", "", false, "List Fleet Members")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListFleets, "list-fleets", "", false, "List Fleets")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListJobMembers, "list-job-members", "", false, "List Job Members")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListJobParameterDefinitions, "list-job-parameter-definitions", "", false, "List Job Parameter Definitions")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListJobs, "list-jobs", "", false, "List Jobs")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListLicenseEndpoints, "list-license-endpoints", "", false, "List License Endpoints")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListLimits, "list-limits", "", false, "List Limits")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListMeteredProducts, "list-metered-products", "", false, "List Metered Products")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListMonitors, "list-monitors", "", false, "List Monitors")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListQueueEnvironments, "list-queue-environments", "", false, "List Queue Environments")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListQueueFleetAssociations, "list-queue-fleet-associations", "", false, "List Queue Fleet Associations")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListQueueLimitAssociations, "list-queue-limit-associations", "", false, "List Queue Limit Associations")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListQueueMembers, "list-queue-members", "", false, "List Queue Members")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListQueues, "list-queues", "", false, "List Queues")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListSessionActions, "list-session-actions", "", false, "List Session Actions")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListSessions, "list-sessions", "", false, "List Sessions")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListSessionsForWorker, "list-sessions-for-worker", "", false, "List Sessions For Worker")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListStepConsumers, "list-step-consumers", "", false, "List Step Consumers")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListStepDependencies, "list-step-dependencies", "", false, "List Step Dependencies")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListSteps, "list-steps", "", false, "List Steps")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListStorageProfiles, "list-storage-profiles", "", false, "List Storage Profiles")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListStorageProfilesForQueue, "list-storage-profiles-for-queue", "", false, "List Storage Profiles For Queue")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListTasks, "list-tasks", "", false, "List Tasks")
	_deadlineCmd.Flags().BoolVarP(&_deadlineListWorkers, "list-workers", "", false, "List Workers")
	_deadlineCmd.Flags().BoolVarP(&_deadlinePutMeteredProduct, "put-metered-product", "", false, "Put Metered Product")
	_deadlineCmd.Flags().BoolVarP(&_deadlineSearchJobs, "search-jobs", "", false, "Search Jobs")
	_deadlineCmd.Flags().BoolVarP(&_deadlineSearchSteps, "search-steps", "", false, "Search Steps")
	_deadlineCmd.Flags().BoolVarP(&_deadlineSearchTasks, "search-tasks", "", false, "Search Tasks")
	_deadlineCmd.Flags().BoolVarP(&_deadlineSearchWorkers, "search-workers", "", false, "Search Workers")
	_deadlineCmd.Flags().BoolVarP(&_deadlineStartSessionsStatisticsAggregation, "start-sessions-statistics-aggregation", "", false, "Start Sessions Statistics Aggregation")
	_deadlineCmd.Flags().BoolVarP(&_deadlineTagResource, "tag-resource", "", false, "Tag Resource")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUntagResource, "untag-resource", "", false, "Untag Resource")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateBudget, "update-budget", "", false, "Update Budget")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateFarm, "update-farm", "", false, "Update Farm")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateFleet, "update-fleet", "", false, "Update Fleet")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateJob, "update-job", "", false, "Update Job")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateLimit, "update-limit", "", false, "Update Limit")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateMonitor, "update-monitor", "", false, "Update Monitor")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateQueue, "update-queue", "", false, "Update Queue")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateQueueEnvironment, "update-queue-environment", "", false, "Update Queue Environment")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateQueueFleetAssociation, "update-queue-fleet-association", "", false, "Update Queue Fleet Association")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateQueueLimitAssociation, "update-queue-limit-association", "", false, "Update Queue Limit Association")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateSession, "update-session", "", false, "Update Session")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateStep, "update-step", "", false, "Update Step")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateStorageProfile, "update-storage-profile", "", false, "Update Storage Profile")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateTask, "update-task", "", false, "Update Task")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateWorker, "update-worker", "", false, "Update Worker")
	_deadlineCmd.Flags().BoolVarP(&_deadlineUpdateWorkerSchedule, "update-worker-schedule", "", false, "Update Worker Schedule")

}
